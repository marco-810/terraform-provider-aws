// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package logs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listLogGroupTags lists logs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listLogGroupTags(ctx context.Context, conn cloudwatchlogsiface.CloudWatchLogsAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &cloudwatchlogs.ListTagsLogGroupInput{
		LogGroupName: aws.String(identifier),
	}

	output, err := conn.ListTagsLogGroupWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// updateLogGroupTags updates logs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateLogGroupTags(ctx context.Context, conn cloudwatchlogsiface.CloudWatchLogsAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.Logs)
	if len(removedTags) > 0 {
		input := &cloudwatchlogs.UntagLogGroupInput{
			LogGroupName: aws.String(identifier),
			Tags:         aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.UntagLogGroupWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.Logs)
	if len(updatedTags) > 0 {
		input := &cloudwatchlogs.TagLogGroupInput{
			LogGroupName: aws.String(identifier),
			Tags:         Tags(updatedTags),
		}

		_, err := conn.TagLogGroupWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
