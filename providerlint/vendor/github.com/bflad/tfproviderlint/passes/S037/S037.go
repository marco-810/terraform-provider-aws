package S037

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

var Analyzer = analysisutils.SchemaAttributeReferencesAnalyzer("S037", schema.SchemaFieldExactlyOneOf)
