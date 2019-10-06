// +build ignore

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
)

const filename = `update_tags_gen.go`

var serviceNames = []string{
	"amplify",
	"apigateway",
	"apigatewayv2",
	"appmesh",
	"appstream",
	"appsync",
	"athena",
	"backup",
	"cloudhsmv2",
	"cloudwatch",
	"cloudwatchevents",
	"codecommit",
	"codedeploy",
	"codepipeline",
	"cognitoidentity",
	"cognitoidentityprovider",
	"configservice",
	"databasemigrationservice",
	"datapipeline",
	"datasync",
	"dax",
	"devicefarm",
	"directconnect",
	"directoryservice",
	"docdb",
	"dynamodb",
	"ecr",
	"ecs",
	"efs",
	"eks",
	"elasticache",
	"elasticsearchservice",
	"emr",
	"firehose",
	"fsx",
	"glue",
	"guardduty",
	"iot",
	"iotanalytics",
	"iotevents",
	"kafka",
	"kinesisanalytics",
	"kinesisanalyticsv2",
	"kms",
	"lambda",
	"licensemanager",
	"lightsail",
	"mediaconnect",
	"mediaconvert",
	"medialive",
	"mediapackage",
	"mediastore",
	"mq",
	"neptune",
	"opsworks",
	"organizations",
	"qldb",
	"ram",
	"rds",
	"redshift",
	"route53resolver",
	"secretsmanager",
	"securityhub",
	"sfn",
	"sns",
	"ssm",
	"storagegateway",
	"swf",
	"transfer",
	"waf",
	"workspaces",
}

type TemplateData struct {
	ServiceNames []string
}

func main() {
	// Always sort to reduce any potential generation churn
	sort.Strings(serviceNames)

	templateData := TemplateData{
		ServiceNames: serviceNames,
	}
	templateFuncMap := template.FuncMap{
		"ClientType":                keyvaluetags.ServiceClientType,
		"TagFunction":               ServiceTagFunction,
		"TagInputIdentifierField":   ServiceTagInputIdentifierField,
		"TagInputResourceTypeField": ServiceTagInputResourceTypeField,
		"TagInputTagsField":         ServiceTagInputTagsField,
		"Title":                     strings.Title,
		"UntagFunction":             ServiceUntagFunction,
		"UntagInputTagsField":       ServiceUntagInputTagsField,
	}

	tmpl, err := template.New("updatetags").Funcs(templateFuncMap).Parse(templateBody)

	if err != nil {
		log.Fatalf("error parsing template: %s", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		log.Fatalf("error executing template: %s", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		log.Fatalf("error formatting generated file: %s", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		log.Fatalf("error creating file (%s): %s", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		log.Fatalf("error writing to file (%s): %s", filename, err)
	}
}

var templateBody = `
// Code generated by generators/updatetags/main.go; DO NOT EDIT.

package keyvaluetags

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
{{- range .ServiceNames }}
	"github.com/aws/aws-sdk-go/service/{{ . }}"
{{- end }}
)
{{ range .ServiceNames }}

// {{ . | Title }}UpdateTags updates {{ . }} service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func {{ . | Title }}UpdateTags(conn {{ . | ClientType }}, identifier string{{ if . | TagInputResourceTypeField }}, resourceType string{{ end }}, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := New(oldTagsMap)
	newTags := New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &{{ . }}.{{ . | UntagFunction }}Input{
			{{ . | TagInputIdentifierField }}:   aws.String(identifier),
			{{- if . | TagInputResourceTypeField }}
			{{ . | TagInputResourceTypeField }}: aws.String(resourceType),
			{{- end }}
			{{ . | UntagInputTagsField }}:       aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.{{ . | UntagFunction }}(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %s", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &{{ . }}.{{ . | TagFunction }}Input{
			{{ . | TagInputIdentifierField }}:   aws.String(identifier),
			{{- if . | TagInputResourceTypeField }}
			{{ . | TagInputResourceTypeField }}: aws.String(resourceType),
			{{- end }}
			{{ . | TagInputTagsField }}:         updatedTags.IgnoreAws().{{ . | Title }}Tags(),
		}

		_, err := conn.{{ . | TagFunction }}(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %s", identifier, err)
		}
	}

	return nil
}
{{- end }}
`

// ServiceTagFunction determines the service tagging function.
func ServiceTagFunction(serviceName string) string {
	switch serviceName {
	case "databasemigrationservice":
		return "AddTagsToResource"
	case "datapipeline":
		return "AddTags"
	case "directoryservice":
		return "AddTagsToResource"
	case "docdb":
		return "AddTagsToResource"
	case "efs":
		return "CreateTags"
	case "elasticache":
		return "AddTagsToResource"
	case "elasticsearchservice":
		return "AddTags"
	case "emr":
		return "AddTags"
	case "firehose":
		return "TagDeliveryStream"
	case "medialive":
		return "CreateTags"
	case "mq":
		return "CreateTags"
	case "neptune":
		return "AddTagsToResource"
	case "rds":
		return "AddTagsToResource"
	case "redshift":
		return "CreateTags"
	case "sagemaker":
		return "AddTags"
	case "ssm":
		return "AddTagsToResource"
	case "storagegateway":
		return "AddTagsToResource"
	case "workspaces":
		return "CreateTags"
	default:
		return "TagResource"
	}
}

// ServiceTagInputIdentifierField determines the service tag identifier field.
func ServiceTagInputIdentifierField(serviceName string) string {
	switch serviceName {
	case "athena":
		return "ResourceARN"
	case "cloudhsmv2":
		return "ResourceId"
	case "cloudwatch":
		return "ResourceARN"
	case "cloudwatchevents":
		return "ResourceARN"
	case "datapipeline":
		return "PipelineId"
	case "dax":
		return "ResourceName"
	case "devicefarm":
		return "ResourceARN"
	case "directoryservice":
		return "ResourceId"
	case "docdb":
		return "ResourceName"
	case "efs":
		return "FileSystemId"
	case "elasticache":
		return "ResourceName"
	case "elasticsearchservice":
		return "ARN"
	case "emr":
		return "ResourceId"
	case "firehose":
		return "DeliveryStreamName"
	case "fsx":
		return "ResourceARN"
	case "kinesisanalytics":
		return "ResourceARN"
	case "kinesisanalyticsv2":
		return "ResourceARN"
	case "kms":
		return "KeyId"
	case "lambda":
		return "Resource"
	case "lightsail":
		return "ResourceName"
	case "mediaconvert":
		return "Arn"
	case "mediastore":
		return "Resource"
	case "neptune":
		return "ResourceName"
	case "organizations":
		return "ResourceId"
	case "ram":
		return "ResourceShareArn"
	case "rds":
		return "ResourceName"
	case "redshift":
		return "ResourceName"
	case "secretsmanager":
		return "SecretId"
	case "ssm":
		return "ResourceId"
	case "storagegateway":
		return "ResourceARN"
	case "transfer":
		return "Arn"
	case "waf":
		return "ResourceARN"
	case "workspaces":
		return "ResourceId"
	default:
		return "ResourceArn"
	}
}

// ServiceTagInputTagsField determines the service tagging tags field.
func ServiceTagInputTagsField(serviceName string) string {
	switch serviceName {
	case "cloudhsmv2":
		return "TagList"
	case "elasticsearchservice":
		return "TagList"
	case "glue":
		return "TagsToAdd"
	default:
		return "Tags"
	}
}

// ServiceTagInputResourceTypeField determines the service tagging resource type field.
func ServiceTagInputResourceTypeField(serviceName string) string {
	switch serviceName {
	case "ssm":
		return "ResourceType"
	default:
		return ""
	}
}

// ServiceUntagFunction determines the service untagging function.
func ServiceUntagFunction(serviceName string) string {
	switch serviceName {
	case "databasemigrationservice":
		return "RemoveTagsFromResource"
	case "datapipeline":
		return "RemoveTags"
	case "directoryservice":
		return "RemoveTagsFromResource"
	case "docdb":
		return "RemoveTagsFromResource"
	case "efs":
		return "DeleteTags"
	case "elasticache":
		return "RemoveTagsFromResource"
	case "elasticsearchservice":
		return "RemoveTags"
	case "emr":
		return "RemoveTags"
	case "firehose":
		return "UntagDeliveryStream"
	case "medialive":
		return "DeleteTags"
	case "mq":
		return "DeleteTags"
	case "neptune":
		return "RemoveTagsFromResource"
	case "rds":
		return "RemoveTagsFromResource"
	case "redshift":
		return "DeleteTags"
	case "sagemaker":
		return "DeleteTags"
	case "ssm":
		return "RemoveTagsFromResource"
	case "storagegateway":
		return "RemoveTagsFromResource"
	case "workspaces":
		return "DeleteTags"
	default:
		return "UntagResource"
	}
}

// ServiceUntagInputTagsField determines the service untagging tags field.
func ServiceUntagInputTagsField(serviceName string) string {
	switch serviceName {
	case "backup":
		return "TagKeyList"
	case "cloudhsmv2":
		return "TagKeyList"
	case "datasync":
		return "Keys"
	case "glue":
		return "TagsToRemove"
	default:
		return "TagKeys"
	}
}
