// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_config_aggregate_authorization":       ResourceAggregateAuthorization,
		"aws_config_config_rule":                   ResourceConfigRule,
		"aws_config_configuration_aggregator":      ResourceConfigurationAggregator,
		"aws_config_configuration_recorder":        ResourceConfigurationRecorder,
		"aws_config_configuration_recorder_status": ResourceConfigurationRecorderStatus,
		"aws_config_conformance_pack":              ResourceConformancePack,
		"aws_config_delivery_channel":              ResourceDeliveryChannel,
		"aws_config_organization_conformance_pack": ResourceOrganizationConformancePack,
		"aws_config_organization_custom_rule":      ResourceOrganizationCustomRule,
		"aws_config_organization_managed_rule":     ResourceOrganizationManagedRule,
		"aws_config_remediation_configuration":     ResourceRemediationConfiguration,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ConfigService
}

var ServicePackage = &servicePackage{}
