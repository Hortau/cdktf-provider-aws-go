# cdktf-provider-aws-go
Terraform CDK aws Provider v3.64.2

    go get github.com/hortau/cdktf-provider-aws-go

Example:
```go
package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/hortau/cdktf-provider-aws-go/vpc"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	vpc.NewVpc(stack, jsii.String("main"), &vpc.VpcConfig{
		CidrBlock: jsii.String("10.0.0.0/16"),
		InstanceTenancy: jsii.String("default"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-go")

	app.Synth()
}
```