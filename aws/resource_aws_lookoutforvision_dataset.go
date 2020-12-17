package aws

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lookoutforvision"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/service/lookoutforvision/finder"
)

func resourceAwsLookoutForVisionDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsLookoutForVisionDatasetCreate,
		Read:   resourceAwsLookoutForVisionDatasetRead,
		Delete: resourceAwsLookoutForVisionDatasetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 255),
					validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9](_*-*[a-zA-Z0-9])*$`), "Valid characters are a-z, A-Z, 0-9, - (hyphen) and _ (underscore). Name must begin with an alphanumeric character."),
				),
			},
			"dataset_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"train",
					"test",
				}, false),
			},
			"source": {
				Type:     schema.TypeSet,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringLenBetween(3, 63),
						},
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 1024),
						},
						"version_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAwsLookoutForVisionDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).lookoutforvisionconn

	projectName := d.Get("project").(string)
	datasetType := d.Get("dataset_type").(string)

	input := &lookoutforvision.CreateDatasetInput{
		ProjectName: aws.String(projectName),
		DatasetType: aws.String(datasetType),
		ClientToken: aws.String(resource.UniqueId()),
	}

	// Set dataset source
	if v, ok := d.GetOk("source"); ok {
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			bd := v.(map[string]interface{})
			bucket := bd["bucket"].(string)
			key := bd["key"].(string)
			manifest := &lookoutforvision.InputS3Object{
				Bucket: &bucket,
				Key:    &key,
			}
			if versionId := bd["version_id"].(string); versionId != "" {
				manifest.VersionId = &versionId
			}
			input.DatasetSource = &lookoutforvision.DatasetSource{
				GroundTruthManifest: &lookoutforvision.DatasetGroundTruthManifest{
					S3Object: manifest,
				},
			}
		}
	}

	log.Printf("[DEBUG] Amazon Lookout for Vision dataset create config: %#v", *input)
	_, err := conn.CreateDataset(input)
	if err != nil {
		return fmt.Errorf("Error creating Amazon Lookout for Vision dataset: %w", err)
	}

	d.SetId(strings.Join([]string{projectName, datasetType}, "/"))

	return resourceAwsLookoutForVisionDatasetRead(d, meta)
}

func resourceAwsLookoutForVisionDatasetRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).lookoutforvisionconn
	projectName := d.Get("project").(string)
	datasetType := d.Get("dataset_type").(string)

	_, err := finder.DatasetByProjectAndType(conn, projectName, datasetType)
	if err != nil {
		if isAWSErr(err, "ValidationException", "Cannot find dataset") {
			d.SetId("")
			log.Printf("[WARN] Unable to find Amazon Lookout for Vision dataset (Project: %s, Type: %s); removing from state", projectName, datasetType)
			return nil
		}
		return fmt.Errorf("error reading Amazon Lookout for Vision dataset (Project: %s, Type: %s): %w", projectName, datasetType, err)

	}

	return nil
}

func resourceAwsLookoutForVisionDatasetDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).lookoutforvisionconn

	projectName := d.Get("project").(string)
	datasetType := d.Get("dataset_type").(string)

	input := &lookoutforvision.DeleteDatasetInput{
		ProjectName: aws.String(projectName),
		DatasetType: aws.String(datasetType),
	}

	if _, err := conn.DeleteDataset(input); err != nil {
		if isAWSErr(err, "ValidationException", "Cannot find dataset") {
			return nil
		}
		return fmt.Errorf("error deleting Lookout for Vision dataset (%s): %w", d.Id(), err)
	}

	return nil
}
