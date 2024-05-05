// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_rds_cluster_activity_stream")
func ResourceClusterActivityStream() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceClusterActivityStreamCreate,
		ReadWithoutTimeout:   resourceClusterActivityStreamRead,
		DeleteWithoutTimeout: resourceClusterActivityStreamDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"engine_native_audit_fields_included": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"kinesis_stream_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(rds.ActivityStreamMode_Values(), false),
			},
			"resource_arn": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidARN,
			},
		},
	}
}

func resourceClusterActivityStreamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).RDSConn(ctx)

	arn := d.Get("resource_arn").(string)
	input := &rds.StartActivityStreamInput{
		ApplyImmediately:                aws.Bool(true),
		EngineNativeAuditFieldsIncluded: aws.Bool(d.Get("engine_native_audit_fields_included").(bool)),
		KmsKeyId:                        aws.String(d.Get("kms_key_id").(string)),
		Mode:                            aws.String(d.Get("mode").(string)),
		ResourceArn:                     aws.String(arn),
	}

	_, err := conn.StartActivityStreamWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("creating RDS Database Activity Stream (%s): %s", arn, err)
	}

	d.SetId(arn)

	if err := waitActivityStreamStarted(ctx, conn, d.Id()); err != nil {
		return diag.Errorf("waiting for RDS Database Activity Stream (%s) start: %s", d.Id(), err)
	}

	return resourceClusterActivityStreamRead(ctx, d, meta)
}

func resourceClusterActivityStreamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).RDSConn(ctx)

	dbClusterARN, err := IsDBClusterARN(d.Id())
	if err != nil {
		return diag.Errorf("reading RDS Database Activity Stream (%s): %s", d.Id(), err)
	}

	if dbClusterARN {
		output, err := FindDBClusterWithActivityStream(ctx, conn, d.Id())

		if !d.IsNewResource() && tfresource.NotFound(err) {
			log.Printf("[WARN] Database Activity Stream (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		if err != nil {
			return diag.Errorf("reading Database Activity Stream (%s): %s", d.Id(), err)
		}

		d.Set("engine_native_audit_fields_included", false)
		d.Set("kinesis_stream_name", output.ActivityStreamKinesisStreamName)
		d.Set("kms_key_id", output.ActivityStreamKmsKeyId)
		d.Set("mode", output.ActivityStreamMode)
		d.Set("resource_arn", output.DBClusterArn)
	} else {
		output, err := FindDBInstanceWithActivityStream(ctx, conn, d.Id())

		if !d.IsNewResource() && tfresource.NotFound(err) {
			log.Printf("[WARN] Database Activity Stream (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		if err != nil {
			return diag.Errorf("reading RDS Database Activity Stream (%s): %s", d.Id(), err)
		}

		d.Set("engine_native_audit_fields_included", output.ActivityStreamEngineNativeAuditFieldsIncluded)
		d.Set("kinesis_stream_name", output.ActivityStreamKinesisStreamName)
		d.Set("kms_key_id", output.ActivityStreamKmsKeyId)
		d.Set("mode", output.ActivityStreamMode)
		d.Set("resource_arn", output.DBInstanceArn)
	}

	return nil
}

func resourceClusterActivityStreamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).RDSConn(ctx)

	log.Printf("[DEBUG] Deleting RDS Database Activity Stream: %s", d.Id())
	_, err := conn.StopActivityStreamWithContext(ctx, &rds.StopActivityStreamInput{
		ApplyImmediately: aws.Bool(true),
		ResourceArn:      aws.String(d.Id()),
	})

	if tfawserr.ErrMessageContains(err, "InvalidParameterCombination", "Activity Streams feature expected to be started, but is stopped") {
		return nil
	}

	if err != nil {
		return diag.Errorf("stopping RDS Database Activity Stream (%s): %s", d.Id(), err)
	}

	if err := waitActivityStreamStopped(ctx, conn, d.Id()); err != nil {
		return diag.Errorf("waiting for RDS Database Activity Stream (%s) stop: %s", d.Id(), err)
	}

	return nil
}

func FindDBClusterWithActivityStream(ctx context.Context, conn *rds.RDS, arn string) (*rds.DBCluster, error) {
	output, err := FindDBClusterByID(ctx, conn, arn)
	if err != nil {
		return nil, err
	}

	if status := aws.StringValue(output.ActivityStreamStatus); status == rds.ActivityStreamStatusStopped {
		return nil, &retry.NotFoundError{
			Message: status,
		}
	}

	return output, nil
}

func FindDBInstanceWithActivityStream(ctx context.Context, conn *rds.RDS, arn string) (*rds.DBInstance, error) {
	output, err := findDBInstanceByIDSDKv1(ctx, conn, arn)
	if err != nil {
		return nil, err
	}

	if status := aws.StringValue(output.ActivityStreamStatus); status == rds.ActivityStreamStatusStopped {
		return nil, &retry.NotFoundError{
			Message: status,
		}
	}

	return output, nil
}

func IsDBClusterARN(s string) (bool, error) {
	parsedArn, err := arn.Parse(s)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(parsedArn.Resource, "cluster:"), nil
}

func statusDBClusterActivityStream(ctx context.Context, conn *rds.RDS, arn string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindDBClusterByID(ctx, conn, arn)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ActivityStreamStatus), nil
	}
}

func statusDBInstanceActivityStream(ctx context.Context, conn *rds.RDS, arn string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := findDBInstanceByIDSDKv1(ctx, conn, arn)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ActivityStreamStatus), nil
	}
}

const (
	dbClusterActivityStreamStartedTimeout = 30 * time.Minute
	dbClusterActivityStreamStoppedTimeout = 30 * time.Minute
)

func waitActivityStreamStarted(ctx context.Context, conn *rds.RDS, arn string) error {
	dbClusterARN, err := IsDBClusterARN(arn)
	if err != nil {
		return err
	}

	stateConf := &retry.StateChangeConf{
		Pending:    []string{rds.ActivityStreamStatusStarting},
		Target:     []string{rds.ActivityStreamStatusStarted},
		Timeout:    dbClusterActivityStreamStartedTimeout,
		MinTimeout: 10 * time.Second,
		Delay:      30 * time.Second,
	}

	if dbClusterARN {
		stateConf.Refresh = statusDBClusterActivityStream(ctx, conn, arn)
	} else {
		stateConf.Refresh = statusDBInstanceActivityStream(ctx, conn, arn)
	}

	_, err = stateConf.WaitForStateContext(ctx)

	return err
}

func waitActivityStreamStopped(ctx context.Context, conn *rds.RDS, arn string) error {
	dbClusterARN, err := IsDBClusterARN(arn)
	if err != nil {
		return err
	}

	stateConf := &retry.StateChangeConf{
		Pending:    []string{rds.ActivityStreamStatusStopping},
		Target:     []string{rds.ActivityStreamStatusStopped},
		Timeout:    dbClusterActivityStreamStoppedTimeout,
		MinTimeout: 10 * time.Second,
		Delay:      30 * time.Second,
	}

	if dbClusterARN {
		stateConf.Refresh = statusDBClusterActivityStream(ctx, conn, arn)
	} else {
		stateConf.Refresh = statusDBInstanceActivityStream(ctx, conn, arn)
	}

	_, err = stateConf.WaitForStateContext(ctx)

	return err
}
