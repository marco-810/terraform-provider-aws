// The awsproviderlint command is a static checker for the Terraform AWS Provider.
package main

import (
	tfpasses "github.com/bflad/tfproviderlint/passes"
	tfxpasses "github.com/bflad/tfproviderlint/xpasses"
	awspasses "github.com/hashicorp/terraform-provider-aws/awsproviderlint/passes"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func main() {
	var analyzers []*analysis.Analyzer
	analyzers = append(analyzers, tfpasses.AllChecks...)
	analyzers = append(analyzers, tfxpasses.AllChecks...)
	analyzers = append(analyzers, awspasses.AllChecks...)
	multichecker.Main(analyzers...)
}
