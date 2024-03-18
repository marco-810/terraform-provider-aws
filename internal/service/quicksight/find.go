// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
)

func FindGroupMembership(ctx context.Context, conn *quicksight.Client, listInput *quicksight.ListGroupMembershipsInput, userName string) (bool, error) {
	found := false

	for {
		resp, err := conn.ListGroupMemberships(ctx, listInput)
		if err != nil {
			return false, err
		}

		for _, member := range resp.GroupMemberList {
			if aws.ToString(member.MemberName) == userName {
				found = true
				break
			}
		}

		if found || resp.NextToken == nil {
			break
		}

		listInput.NextToken = resp.NextToken
	}

	return found, nil
}
