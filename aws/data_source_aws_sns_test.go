package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceAwsSnsTopic_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { testAccPreCheck(t) },
		ErrorCheck: testAccErrorCheck(t, sns.EndpointsID),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsSnsTopicConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceAwsSnsTopicCheck("data.aws_sns_topic.by_name"),
				),
			},
		},
	})
}

func testAccDataSourceAwsSnsTopicCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		snsTopicRs, ok := s.RootModule().Resources["aws_sns_topic.tf_test"]
		if !ok {
			return fmt.Errorf("can't find aws_sns_topic.tf_test in state")
		}

		attr := rs.Primary.Attributes

		if attr["name"] != snsTopicRs.Primary.Attributes["name"] {
			return fmt.Errorf(
				"name is %s; want %s",
				attr["name"],
				snsTopicRs.Primary.Attributes["name"],
			)
		}

		return nil
	}
}

const testAccDataSourceAwsSnsTopicConfig = `
resource "aws_sns_topic" "tf_wrong1" {
  name = "wrong1"
}

resource "aws_sns_topic" "tf_test" {
  name = "tf_test"
}

resource "aws_sns_topic" "tf_wrong2" {
  name = "wrong2"
}

data "aws_sns_topic" "by_name" {
  name = aws_sns_topic.tf_test.name
}
`
