package resourcedatagetchangeassignstmt

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

var Analyzer = analysisutils.ReceiverMethodAssignStmtAnalyzer(
	"resourcedatagetchangeassignstmt",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"GetChange",
)
