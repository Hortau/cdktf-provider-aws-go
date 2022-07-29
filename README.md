# ⛔️ DEPRECATED
# use the official: https://github.com/hashicorp/cdktf-provider-aws-go

## cdktf-provider-aws-go
- Terraform [aws Provider](https://registry.terraform.io/providers/hashicorp/aws/latest) v3.74.3
- Generated with [CDK for Terraform](https://github.com/hashicorp/terraform-cdk) v0.9.1

### To Add in you project:

    go get github.com/hortau/cdktf-provider-aws-go

### Example:
```go
package main

import (
	"os"

	"github.com/hortau/cdktf-provider-aws-go/vpc"
	aws "github.com/hortau/cdktf-provider-aws-go"

	constructs "github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, jsii.String("aws"), &aws.AwsProviderConfig{
		Region: jsii.String(os.Getenv("AWS_REGION")),
		AccessKey: jsii.String(os.Getenv("AWS_ACCESS_KEY")),
		SecretKey: jsii.String(os.Getenv("AWS_SECRET_KEY")),
	})


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