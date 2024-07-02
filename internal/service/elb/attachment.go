// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	awstypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKResource("aws_elb_attachment")
func ResourceAttachment() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceAttachmentCreate,
		ReadWithoutTimeout:   resourceAttachmentRead,
		DeleteWithoutTimeout: resourceAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"elb": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"instance": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceAttachmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ELBClient(ctx)
	elbName := d.Get("elb").(string)

	instance := d.Get("instance").(string)

	registerInstancesOpts := elasticloadbalancing.RegisterInstancesWithLoadBalancerInput{
		LoadBalancerName: aws.String(elbName),
		Instances:        []awstypes.Instance{{InstanceId: aws.String(instance)}},
	}

	log.Printf("[INFO] registering instance %s with ELB %s", instance, elbName)

	err := retry.RetryContext(ctx, 10*time.Minute, func() *retry.RetryError {
		_, err := conn.RegisterInstancesWithLoadBalancer(ctx, &registerInstancesOpts)

		if tfawserr.ErrCodeEquals(err, "InvalidTarget") {
			return retry.RetryableError(fmt.Errorf("attaching instance to ELB, retrying: %s", err))
		}

		if err != nil {
			return retry.NonRetryableError(err)
		}

		return nil
	})
	if tfresource.TimedOut(err) {
		_, err = conn.RegisterInstancesWithLoadBalancer(ctx, &registerInstancesOpts)
	}
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "Failure registering instances with ELB: %s", err)
	}

	//lintignore:R016 // Allow legacy unstable ID usage in managed resource
	d.SetId(id.PrefixedUniqueId(fmt.Sprintf("%s-", elbName)))

	return diags
}

func resourceAttachmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ELBClient(ctx)
	elbName := d.Get("elb").(string)

	// only add the instance that was previously defined for this resource
	expected := d.Get("instance").(string)

	// Retrieve the ELB properties to get a list of attachments
	describeElbOpts := &elasticloadbalancing.DescribeLoadBalancersInput{
		LoadBalancerNames: []string{elbName},
	}

	resp, err := conn.DescribeLoadBalancers(ctx, describeElbOpts)
	if err != nil {
		if !d.IsNewResource() && errs.IsA[*awstypes.AccessPointNotFoundException](err) {
			log.Printf("[WARN] ELB Classic LB (%s) not found, removing from state", elbName)
			d.SetId("")
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "retrieving ELB Classic LB (%s): %s", elbName, err)
	}
	if !d.IsNewResource() && len(resp.LoadBalancerDescriptions) != 1 {
		log.Printf("[WARN] ELB Classic LB (%s) not found, removing from state", elbName)
		d.SetId("")
		return diags
	}

	// only set the instance Id that this resource manages
	found := false
	for _, i := range resp.LoadBalancerDescriptions[0].Instances {
		if expected == aws.ToString(i.InstanceId) {
			d.Set("instance", expected)
			found = true
		}
	}

	if !d.IsNewResource() && !found {
		log.Printf("[WARN] instance %s not found in elb attachments", expected)
		d.SetId("")
	}

	return diags
}

func resourceAttachmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ELBClient(ctx)
	elbName := d.Get("elb").(string)

	instance := d.Get("instance").(string)

	log.Printf("[INFO] Deleting Attachment %s from: %s", instance, elbName)

	deRegisterInstancesOpts := elasticloadbalancing.DeregisterInstancesFromLoadBalancerInput{
		LoadBalancerName: aws.String(elbName),
		Instances:        []awstypes.Instance{{InstanceId: aws.String(instance)}},
	}

	_, err := conn.DeregisterInstancesFromLoadBalancer(ctx, &deRegisterInstancesOpts)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "Failure deregistering instances from ELB: %s", err)
	}

	return diags
}
