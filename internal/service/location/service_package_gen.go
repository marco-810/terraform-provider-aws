// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package location

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
	return map[string]func() *schema.Resource{
		"aws_location_geofence_collection":  DataSourceGeofenceCollection,
		"aws_location_map":                  DataSourceMap,
		"aws_location_place_index":          DataSourcePlaceIndex,
		"aws_location_route_calculator":     DataSourceRouteCalculator,
		"aws_location_tracker":              DataSourceTracker,
		"aws_location_tracker_association":  DataSourceTrackerAssociation,
		"aws_location_tracker_associations": DataSourceTrackerAssociations,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_location_geofence_collection": ResourceGeofenceCollection,
		"aws_location_map":                 ResourceMap,
		"aws_location_place_index":         ResourcePlaceIndex,
		"aws_location_route_calculator":    ResourceRouteCalculator,
		"aws_location_tracker":             ResourceTracker,
		"aws_location_tracker_association": ResourceTrackerAssociation,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Location
}

var ServicePackage = &servicePackage{}
