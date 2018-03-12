package aws

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsWafRegionalWebAclAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsWafRegionalWebAclAssociationCreate,
		Read:   resourceAwsWafRegionalWebAclAssociationRead,
		Update: resourceAwsWafRegionalWebAclAssociationUpdate,
		Delete: resourceAwsWafRegionalWebAclAssociationDelete,

		Schema: map[string]*schema.Schema{
			"web_acl_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsWafRegionalWebAclAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafregionalconn

	log.Printf(
		"[INFO] Creating WAF Regional Web ACL association: %s => %s",
		d.Get("web_acl_id").(string),
		d.Get("resource_arn").(string))

	params := &wafregional.AssociateWebACLInput{
		WebACLId:    aws.String(d.Get("web_acl_id").(string)),
		ResourceArn: aws.String(d.Get("resource_arn").(string)),
	}

	// create association and wait on retryable error
	// no response body
	var err error
	err = resource.Retry(2*time.Minute, func() *resource.RetryError {
		_, err = conn.AssociateWebACL(params)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "WAFUnavailableEntityException" {
					return resource.RetryableError(awsErr)
				}
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Store association id
	d.SetId(fmt.Sprintf("%s:%s", *params.WebACLId, *params.ResourceArn))

	return nil
}

func resourceAwsWafRegionalWebAclAssociationRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafregionalconn

	webAclId, resourceArn := resourceAwsWafRegionalWebAclAssociationParseId(d.Id())

	// List all resources for Web ACL and see if we get a match
	params := &wafregional.ListResourcesForWebACLInput{
		WebACLId: aws.String(webAclId),
	}

	resp, err := conn.ListResourcesForWebACL(params)
	if err != nil {
		return err
	}

	// Find match
	found := false
	for _, listResourceArn := range resp.ResourceArns {
		if resourceArn == *listResourceArn {
			found = true
			break
		}
	}
	if !found {
		// It seems it doesn't exist anymore, so clear the ID
		d.SetId("")
	}

	return nil
}

func resourceAwsWafRegionalWebAclAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsWafRegionalWebAclAssociationRead(d, meta)
}

func resourceAwsWafRegionalWebAclAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafregionalconn

	_, resourceArn := resourceAwsWafRegionalWebAclAssociationParseId(d.Id())

	log.Printf("[INFO] Deleting WAF Regional Web ACL association: %s", resourceArn)

	params := &wafregional.DisassociateWebACLInput{
		ResourceArn: aws.String(resourceArn),
	}

	// If action successful HTTP 200 response with an empty body
	_, err := conn.DisassociateWebACL(params)
	if err != nil {
		return err
	}

	return nil
}

func resourceAwsWafRegionalWebAclAssociationParseId(id string) (webAclId, resourceArn string) {
	parts := strings.SplitN(id, ":", 2)
	webAclId = parts[0]
	resourceArn = parts[1]
	return
}
