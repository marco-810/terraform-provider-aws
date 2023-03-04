// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package opsworks

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
		"aws_opsworks_application":       ResourceApplication,
		"aws_opsworks_custom_layer":      ResourceCustomLayer,
		"aws_opsworks_ecs_cluster_layer": ResourceECSClusterLayer,
		"aws_opsworks_ganglia_layer":     ResourceGangliaLayer,
		"aws_opsworks_haproxy_layer":     ResourceHAProxyLayer,
		"aws_opsworks_instance":          ResourceInstance,
		"aws_opsworks_java_app_layer":    ResourceJavaAppLayer,
		"aws_opsworks_memcached_layer":   ResourceMemcachedLayer,
		"aws_opsworks_mysql_layer":       ResourceMySQLLayer,
		"aws_opsworks_nodejs_app_layer":  ResourceNodejsAppLayer,
		"aws_opsworks_permission":        ResourcePermission,
		"aws_opsworks_php_app_layer":     ResourcePHPAppLayer,
		"aws_opsworks_rails_app_layer":   ResourceRailsAppLayer,
		"aws_opsworks_rds_db_instance":   ResourceRDSDBInstance,
		"aws_opsworks_stack":             ResourceStack,
		"aws_opsworks_static_web_layer":  ResourceStaticWebLayer,
		"aws_opsworks_user_profile":      ResourceUserProfile,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpsWorks
}

var ServicePackage = &servicePackage{}
