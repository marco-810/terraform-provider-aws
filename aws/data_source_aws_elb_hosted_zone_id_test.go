package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSElbHostedZoneId_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsElbHostedZoneIdConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_elb_hosted_zone_id.main", "id", "Z1H1FL5HABSF5"),
				),
			},
			{
				Config: testAccCheckAwsElbHostedZoneIdExplicitRegionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_elb_hosted_zone_id.regional", "id", "Z1UDT6IFJ4EJM"),
				),
			},
		},
	})
}

const testAccCheckAwsElbHostedZoneIdConfig = `
data "aws_elb_hosted_zone_id" "main" { 
	elb_type = "application"
}
`

const testAccCheckAwsElbHostedZoneIdExplicitRegionConfig = `
data "aws_elb_hosted_zone_id" "regional" {
	region = "eu-west-1"
	elb_type = "network"
}
`
