package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/lambdafunction/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lambda_alias aws_lambda_alias}.
type DataAwsLambdaAlias interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionVersion() *string
	Id() *string
	InvokeArn() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLambdaAlias
type jsiiProxy_DataAwsLambdaAlias struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLambdaAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_alias aws_lambda_alias} Data Source.
func NewDataAwsLambdaAlias(scope constructs.Construct, id *string, config *DataAwsLambdaAliasConfig) DataAwsLambdaAlias {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaAlias{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_alias aws_lambda_alias} Data Source.
func NewDataAwsLambdaAlias_Override(d DataAwsLambdaAlias, scope constructs.Construct, id *string, config *DataAwsLambdaAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaAlias",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaAlias) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLambdaAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.DataAwsLambdaAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLambdaAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.DataAwsLambdaAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsLambdaAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type DataAwsLambdaAliasConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_alias#function_name DataAwsLambdaAlias#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_alias#name DataAwsLambdaAlias#name}.
	Name *string `json:"name" yaml:"name"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lambda_code_signing_config aws_lambda_code_signing_config}.
type DataAwsLambdaCodeSigningConfig interface {
	cdktf.TerraformDataSource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ConfigId() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastModified() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	AllowedPublishers(index *string) DataAwsLambdaCodeSigningConfigAllowedPublishers
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	Policies(index *string) DataAwsLambdaCodeSigningConfigPolicies
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLambdaCodeSigningConfig
type jsiiProxy_DataAwsLambdaCodeSigningConfig struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) ConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_code_signing_config aws_lambda_code_signing_config} Data Source.
func NewDataAwsLambdaCodeSigningConfig(scope constructs.Construct, id *string, config *DataAwsLambdaCodeSigningConfigConfig) DataAwsLambdaCodeSigningConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaCodeSigningConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_code_signing_config aws_lambda_code_signing_config} Data Source.
func NewDataAwsLambdaCodeSigningConfig_Override(d DataAwsLambdaCodeSigningConfig, scope constructs.Construct, id *string, config *DataAwsLambdaCodeSigningConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLambdaCodeSigningConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLambdaCodeSigningConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) AllowedPublishers(index *string) DataAwsLambdaCodeSigningConfigAllowedPublishers {
	var returns DataAwsLambdaCodeSigningConfigAllowedPublishers

	_jsii_.Invoke(
		d,
		"allowedPublishers",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) Policies(index *string) DataAwsLambdaCodeSigningConfigPolicies {
	var returns DataAwsLambdaCodeSigningConfigPolicies

	_jsii_.Invoke(
		d,
		"policies",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsLambdaCodeSigningConfigAllowedPublishers interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	SigningProfileVersionArns() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaCodeSigningConfigAllowedPublishers
type jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) SigningProfileVersionArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingProfileVersionArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaCodeSigningConfigAllowedPublishers(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaCodeSigningConfigAllowedPublishers {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfigAllowedPublishers",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaCodeSigningConfigAllowedPublishers_Override(d DataAwsLambdaCodeSigningConfigAllowedPublishers, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfigAllowedPublishers",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigAllowedPublishers) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// AWS Lambda.
type DataAwsLambdaCodeSigningConfigConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_code_signing_config#arn DataAwsLambdaCodeSigningConfig#arn}.
	Arn *string `json:"arn" yaml:"arn"`
}

type DataAwsLambdaCodeSigningConfigPolicies interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedArtifactOnDeployment() *string
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaCodeSigningConfigPolicies
type jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) UntrustedArtifactOnDeployment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedArtifactOnDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaCodeSigningConfigPolicies(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaCodeSigningConfigPolicies {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfigPolicies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaCodeSigningConfigPolicies_Override(d DataAwsLambdaCodeSigningConfigPolicies, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaCodeSigningConfigPolicies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaCodeSigningConfigPolicies) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lambda_function aws_lambda_function}.
type DataAwsLambdaFunction interface {
	cdktf.TerraformDataSource
	Architectures() *[]*string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CodeSigningConfigArn() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Handler() *string
	Id() *string
	ImageUri() *string
	InvokeArn() *string
	KmsKeyArn() *string
	LastModified() *string
	Layers() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MemorySize() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	QualifiedArn() *string
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	ReservedConcurrentExecutions() *float64
	Role() *string
	Runtime() *string
	SigningJobArn() *string
	SigningProfileVersionArn() *string
	SourceCodeHash() *string
	SourceCodeSize() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() *float64
	Version() *string
	AddOverride(path *string, value interface{})
	DeadLetterConfig(index *string) DataAwsLambdaFunctionDeadLetterConfig
	Environment(index *string) DataAwsLambdaFunctionEnvironment
	FileSystemConfig(index *string) DataAwsLambdaFunctionFileSystemConfig
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetQualifier()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	TracingConfig(index *string) DataAwsLambdaFunctionTracingConfig
	VpcConfig(index *string) DataAwsLambdaFunctionVpcConfig
}

// The jsii proxy struct for DataAwsLambdaFunction
type jsiiProxy_DataAwsLambdaFunction struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) QualifiedArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_function aws_lambda_function} Data Source.
func NewDataAwsLambdaFunction(scope constructs.Construct, id *string, config *DataAwsLambdaFunctionConfig) DataAwsLambdaFunction {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunction{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_function aws_lambda_function} Data Source.
func NewDataAwsLambdaFunction_Override(d DataAwsLambdaFunction, scope constructs.Construct, id *string, config *DataAwsLambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunction",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunction) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLambdaFunction) DeadLetterConfig(index *string) DataAwsLambdaFunctionDeadLetterConfig {
	var returns DataAwsLambdaFunctionDeadLetterConfig

	_jsii_.Invoke(
		d,
		"deadLetterConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLambdaFunction) Environment(index *string) DataAwsLambdaFunctionEnvironment {
	var returns DataAwsLambdaFunctionEnvironment

	_jsii_.Invoke(
		d,
		"environment",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLambdaFunction) FileSystemConfig(index *string) DataAwsLambdaFunctionFileSystemConfig {
	var returns DataAwsLambdaFunctionFileSystemConfig

	_jsii_.Invoke(
		d,
		"fileSystemConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaFunction) ResetQualifier() {
	_jsii_.InvokeVoid(
		d,
		"resetQualifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsLambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLambdaFunction) TracingConfig(index *string) DataAwsLambdaFunctionTracingConfig {
	var returns DataAwsLambdaFunctionTracingConfig

	_jsii_.Invoke(
		d,
		"tracingConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLambdaFunction) VpcConfig(index *string) DataAwsLambdaFunctionVpcConfig {
	var returns DataAwsLambdaFunctionVpcConfig

	_jsii_.Invoke(
		d,
		"vpcConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// AWS Lambda.
type DataAwsLambdaFunctionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_function#function_name DataAwsLambdaFunction#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_function#qualifier DataAwsLambdaFunction#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_function#tags DataAwsLambdaFunction#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsLambdaFunctionDeadLetterConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TargetArn() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaFunctionDeadLetterConfig
type jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaFunctionDeadLetterConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaFunctionDeadLetterConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionDeadLetterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaFunctionDeadLetterConfig_Override(d DataAwsLambdaFunctionDeadLetterConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionDeadLetterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionDeadLetterConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsLambdaFunctionEnvironment interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variables() *map[string]*string
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaFunctionEnvironment
type jsiiProxy_DataAwsLambdaFunctionEnvironment struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaFunctionEnvironment(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaFunctionEnvironment {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunctionEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionEnvironment",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaFunctionEnvironment_Override(d DataAwsLambdaFunctionEnvironment, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionEnvironment",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionEnvironment) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionEnvironment) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsLambdaFunctionFileSystemConfig interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	LocalMountPath() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaFunctionFileSystemConfig
type jsiiProxy_DataAwsLambdaFunctionFileSystemConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) LocalMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaFunctionFileSystemConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaFunctionFileSystemConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunctionFileSystemConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionFileSystemConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaFunctionFileSystemConfig_Override(d DataAwsLambdaFunctionFileSystemConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionFileSystemConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionFileSystemConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsLambdaFunctionTracingConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Mode() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaFunctionTracingConfig
type jsiiProxy_DataAwsLambdaFunctionTracingConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaFunctionTracingConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaFunctionTracingConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunctionTracingConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionTracingConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaFunctionTracingConfig_Override(d DataAwsLambdaFunctionTracingConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionTracingConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionTracingConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionTracingConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsLambdaFunctionVpcConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	SecurityGroupIds() *[]*string
	SubnetIds() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsLambdaFunctionVpcConfig
type jsiiProxy_DataAwsLambdaFunctionVpcConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLambdaFunctionVpcConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLambdaFunctionVpcConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaFunctionVpcConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionVpcConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLambdaFunctionVpcConfig_Override(d DataAwsLambdaFunctionVpcConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaFunctionVpcConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation aws_lambda_invocation}.
type DataAwsLambdaInvocation interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Id() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	Result() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetQualifier()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLambdaInvocation
type jsiiProxy_DataAwsLambdaInvocation struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLambdaInvocation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaInvocation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation aws_lambda_invocation} Data Source.
func NewDataAwsLambdaInvocation(scope constructs.Construct, id *string, config *DataAwsLambdaInvocationConfig) DataAwsLambdaInvocation {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaInvocation{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaInvocation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation aws_lambda_invocation} Data Source.
func NewDataAwsLambdaInvocation_Override(d DataAwsLambdaInvocation, scope constructs.Construct, id *string, config *DataAwsLambdaInvocationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaInvocation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetInput(val *string) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaInvocation) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLambdaInvocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.DataAwsLambdaInvocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLambdaInvocation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.DataAwsLambdaInvocation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaInvocation) ResetQualifier() {
	_jsii_.InvokeVoid(
		d,
		"resetQualifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaInvocation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsLambdaInvocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaInvocation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type DataAwsLambdaInvocationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation#function_name DataAwsLambdaInvocation#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation#input DataAwsLambdaInvocation#input}.
	Input *string `json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_invocation#qualifier DataAwsLambdaInvocation#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version aws_lambda_layer_version}.
type DataAwsLambdaLayerVersion interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CompatibleArchitecture() *string
	SetCompatibleArchitecture(val *string)
	CompatibleArchitectureInput() *string
	CompatibleArchitectures() *[]*string
	CompatibleRuntime() *string
	SetCompatibleRuntime(val *string)
	CompatibleRuntimeInput() *string
	CompatibleRuntimes() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LayerArn() *string
	LayerName() *string
	SetLayerName(val *string)
	LayerNameInput() *string
	LicenseInfo() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SigningJobArn() *string
	SigningProfileVersionArn() *string
	SourceCodeHash() *string
	SourceCodeSize() *float64
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetCompatibleArchitecture()
	ResetCompatibleRuntime()
	ResetOverrideLogicalId()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLambdaLayerVersion
type jsiiProxy_DataAwsLambdaLayerVersion struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleArchitecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibleArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibleArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleArchitectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleArchitectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibleRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleRuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibleRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CompatibleRuntimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) LayerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) LayerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) LayerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) LicenseInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version aws_lambda_layer_version} Data Source.
func NewDataAwsLambdaLayerVersion(scope constructs.Construct, id *string, config *DataAwsLambdaLayerVersionConfig) DataAwsLambdaLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLambdaLayerVersion{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaLayerVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version aws_lambda_layer_version} Data Source.
func NewDataAwsLambdaLayerVersion_Override(d DataAwsLambdaLayerVersion, scope constructs.Construct, id *string, config *DataAwsLambdaLayerVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.DataAwsLambdaLayerVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetCompatibleArchitecture(val *string) {
	_jsii_.Set(
		j,
		"compatibleArchitecture",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetCompatibleRuntime(val *string) {
	_jsii_.Set(
		j,
		"compatibleRuntime",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetLayerName(val *string) {
	_jsii_.Set(
		j,
		"layerName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLambdaLayerVersion) SetVersion(val *float64) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLambdaLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.DataAwsLambdaLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLambdaLayerVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.DataAwsLambdaLayerVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLambdaLayerVersion) ResetCompatibleArchitecture() {
	_jsii_.InvokeVoid(
		d,
		"resetCompatibleArchitecture",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaLayerVersion) ResetCompatibleRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetCompatibleRuntime",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaLayerVersion) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLambdaLayerVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsLambdaLayerVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type DataAwsLambdaLayerVersionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version#layer_name DataAwsLambdaLayerVersion#layer_name}.
	LayerName *string `json:"layerName" yaml:"layerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version#compatible_architecture DataAwsLambdaLayerVersion#compatible_architecture}.
	CompatibleArchitecture *string `json:"compatibleArchitecture" yaml:"compatibleArchitecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version#compatible_runtime DataAwsLambdaLayerVersion#compatible_runtime}.
	CompatibleRuntime *string `json:"compatibleRuntime" yaml:"compatibleRuntime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lambda_layer_version#version DataAwsLambdaLayerVersion#version}.
	Version *float64 `json:"version" yaml:"version"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias aws_lambda_alias}.
type LambdaAlias interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionVersion() *string
	SetFunctionVersion(val *string)
	FunctionVersionInput() *string
	Id() *string
	InvokeArn() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoutingConfig() LambdaAliasRoutingConfigOutputReference
	RoutingConfigInput() *LambdaAliasRoutingConfig
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutRoutingConfig(value *LambdaAliasRoutingConfig)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetRoutingConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaAlias
type jsiiProxy_LambdaAlias struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) FunctionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) RoutingConfig() LambdaAliasRoutingConfigOutputReference {
	var returns LambdaAliasRoutingConfigOutputReference
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) RoutingConfigInput() *LambdaAliasRoutingConfig {
	var returns *LambdaAliasRoutingConfig
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias aws_lambda_alias} Resource.
func NewLambdaAlias(scope constructs.Construct, id *string, config *LambdaAliasConfig) LambdaAlias {
	_init_.Initialize()

	j := jsiiProxy_LambdaAlias{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias aws_lambda_alias} Resource.
func NewLambdaAlias_Override(l LambdaAlias, scope constructs.Construct, id *string, config *LambdaAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaAlias",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaAlias) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetFunctionVersion(val *string) {
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LambdaAlias) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaAlias) PutRoutingConfig(value *LambdaAliasRoutingConfig) {
	_jsii_.InvokeVoid(
		l,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaAlias) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaAlias) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaAliasConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#function_name LambdaAlias#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#function_version LambdaAlias#function_version}.
	FunctionVersion *string `json:"functionVersion" yaml:"functionVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#name LambdaAlias#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#description LambdaAlias#description}.
	Description *string `json:"description" yaml:"description"`
	// routing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#routing_config LambdaAlias#routing_config}
	RoutingConfig *LambdaAliasRoutingConfig `json:"routingConfig" yaml:"routingConfig"`
}

type LambdaAliasRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#additional_version_weights LambdaAlias#additional_version_weights}.
	AdditionalVersionWeights *map[string]*float64 `json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

type LambdaAliasRoutingConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVersionWeights() *map[string]*float64
	SetAdditionalVersionWeights(val *map[string]*float64)
	AdditionalVersionWeightsInput() *map[string]*float64
	InternalValue() *LambdaAliasRoutingConfig
	SetInternalValue(val *LambdaAliasRoutingConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAdditionalVersionWeights()
}

// The jsii proxy struct for LambdaAliasRoutingConfigOutputReference
type jsiiProxy_LambdaAliasRoutingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) AdditionalVersionWeights() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"additionalVersionWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) AdditionalVersionWeightsInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"additionalVersionWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) InternalValue() *LambdaAliasRoutingConfig {
	var returns *LambdaAliasRoutingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaAliasRoutingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaAliasRoutingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaAliasRoutingConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaAliasRoutingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaAliasRoutingConfigOutputReference_Override(l LambdaAliasRoutingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaAliasRoutingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) SetAdditionalVersionWeights(val *map[string]*float64) {
	_jsii_.Set(
		j,
		"additionalVersionWeights",
		val,
	)
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) SetInternalValue(val *LambdaAliasRoutingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaAliasRoutingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaAliasRoutingConfigOutputReference) ResetAdditionalVersionWeights() {
	_jsii_.InvokeVoid(
		l,
		"resetAdditionalVersionWeights",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config aws_lambda_code_signing_config}.
type LambdaCodeSigningConfig interface {
	cdktf.TerraformResource
	AllowedPublishers() LambdaCodeSigningConfigAllowedPublishersOutputReference
	AllowedPublishersInput() *LambdaCodeSigningConfigAllowedPublishers
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfigId() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastModified() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Policies() LambdaCodeSigningConfigPoliciesOutputReference
	PoliciesInput() *LambdaCodeSigningConfigPolicies
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAllowedPublishers(value *LambdaCodeSigningConfigAllowedPublishers)
	PutPolicies(value *LambdaCodeSigningConfigPolicies)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetPolicies()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaCodeSigningConfig
type jsiiProxy_LambdaCodeSigningConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaCodeSigningConfig) AllowedPublishers() LambdaCodeSigningConfigAllowedPublishersOutputReference {
	var returns LambdaCodeSigningConfigAllowedPublishersOutputReference
	_jsii_.Get(
		j,
		"allowedPublishers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) AllowedPublishersInput() *LambdaCodeSigningConfigAllowedPublishers {
	var returns *LambdaCodeSigningConfigAllowedPublishers
	_jsii_.Get(
		j,
		"allowedPublishersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) ConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Policies() LambdaCodeSigningConfigPoliciesOutputReference {
	var returns LambdaCodeSigningConfigPoliciesOutputReference
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) PoliciesInput() *LambdaCodeSigningConfigPolicies {
	var returns *LambdaCodeSigningConfigPolicies
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config aws_lambda_code_signing_config} Resource.
func NewLambdaCodeSigningConfig(scope constructs.Construct, id *string, config *LambdaCodeSigningConfigConfig) LambdaCodeSigningConfig {
	_init_.Initialize()

	j := jsiiProxy_LambdaCodeSigningConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config aws_lambda_code_signing_config} Resource.
func NewLambdaCodeSigningConfig_Override(l LambdaCodeSigningConfig, scope constructs.Construct, id *string, config *LambdaCodeSigningConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfig",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfig) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfig) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaCodeSigningConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaCodeSigningConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaCodeSigningConfig) PutAllowedPublishers(value *LambdaCodeSigningConfigAllowedPublishers) {
	_jsii_.InvokeVoid(
		l,
		"putAllowedPublishers",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaCodeSigningConfig) PutPolicies(value *LambdaCodeSigningConfigPolicies) {
	_jsii_.InvokeVoid(
		l,
		"putPolicies",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaCodeSigningConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCodeSigningConfig) ResetPolicies() {
	_jsii_.InvokeVoid(
		l,
		"resetPolicies",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCodeSigningConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaCodeSigningConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaCodeSigningConfigAllowedPublishers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#signing_profile_version_arns LambdaCodeSigningConfig#signing_profile_version_arns}.
	SigningProfileVersionArns *[]*string `json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

type LambdaCodeSigningConfigAllowedPublishersOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaCodeSigningConfigAllowedPublishers
	SetInternalValue(val *LambdaCodeSigningConfigAllowedPublishers)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SigningProfileVersionArns() *[]*string
	SetSigningProfileVersionArns(val *[]*string)
	SigningProfileVersionArnsInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaCodeSigningConfigAllowedPublishersOutputReference
type jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) InternalValue() *LambdaCodeSigningConfigAllowedPublishers {
	var returns *LambdaCodeSigningConfigAllowedPublishers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SigningProfileVersionArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingProfileVersionArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SigningProfileVersionArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingProfileVersionArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaCodeSigningConfigAllowedPublishersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaCodeSigningConfigAllowedPublishersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfigAllowedPublishersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaCodeSigningConfigAllowedPublishersOutputReference_Override(l LambdaCodeSigningConfigAllowedPublishersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfigAllowedPublishersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SetInternalValue(val *LambdaCodeSigningConfigAllowedPublishers) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SetSigningProfileVersionArns(val *[]*string) {
	_jsii_.Set(
		j,
		"signingProfileVersionArns",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigAllowedPublishersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaCodeSigningConfigConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// allowed_publishers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#allowed_publishers LambdaCodeSigningConfig#allowed_publishers}
	AllowedPublishers *LambdaCodeSigningConfigAllowedPublishers `json:"allowedPublishers" yaml:"allowedPublishers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#description LambdaCodeSigningConfig#description}.
	Description *string `json:"description" yaml:"description"`
	// policies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#policies LambdaCodeSigningConfig#policies}
	Policies *LambdaCodeSigningConfigPolicies `json:"policies" yaml:"policies"`
}

type LambdaCodeSigningConfigPolicies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#untrusted_artifact_on_deployment LambdaCodeSigningConfig#untrusted_artifact_on_deployment}.
	UntrustedArtifactOnDeployment *string `json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

type LambdaCodeSigningConfigPoliciesOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaCodeSigningConfigPolicies
	SetInternalValue(val *LambdaCodeSigningConfigPolicies)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedArtifactOnDeployment() *string
	SetUntrustedArtifactOnDeployment(val *string)
	UntrustedArtifactOnDeploymentInput() *string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaCodeSigningConfigPoliciesOutputReference
type jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) InternalValue() *LambdaCodeSigningConfigPolicies {
	var returns *LambdaCodeSigningConfigPolicies
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) UntrustedArtifactOnDeployment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedArtifactOnDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) UntrustedArtifactOnDeploymentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedArtifactOnDeploymentInput",
		&returns,
	)
	return returns
}

func NewLambdaCodeSigningConfigPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaCodeSigningConfigPoliciesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfigPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaCodeSigningConfigPoliciesOutputReference_Override(l LambdaCodeSigningConfigPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaCodeSigningConfigPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) SetInternalValue(val *LambdaCodeSigningConfigPolicies) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) SetUntrustedArtifactOnDeployment(val *string) {
	_jsii_.Set(
		j,
		"untrustedArtifactOnDeployment",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaCodeSigningConfigPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping aws_lambda_event_source_mapping}.
type LambdaEventSourceMapping interface {
	cdktf.TerraformResource
	BatchSize() *float64
	SetBatchSize(val *float64)
	BatchSizeInput() *float64
	BisectBatchOnFunctionError() interface{}
	SetBisectBatchOnFunctionError(val interface{})
	BisectBatchOnFunctionErrorInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DestinationConfig() LambdaEventSourceMappingDestinationConfigOutputReference
	DestinationConfigInput() *LambdaEventSourceMappingDestinationConfig
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	EventSourceArnInput() *string
	FilterCriteria() LambdaEventSourceMappingFilterCriteriaOutputReference
	FilterCriteriaInput() *LambdaEventSourceMappingFilterCriteria
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionArn() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionResponseTypes() *[]*string
	SetFunctionResponseTypes(val *[]*string)
	FunctionResponseTypesInput() *[]*string
	Id() *string
	LastModified() *string
	LastProcessingResult() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	MaximumBatchingWindowInSecondsInput() *float64
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	MaximumRecordAgeInSecondsInput() *float64
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	MaximumRetryAttemptsInput() *float64
	Node() constructs.Node
	ParallelizationFactor() *float64
	SetParallelizationFactor(val *float64)
	ParallelizationFactorInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Queues() *[]*string
	SetQueues(val *[]*string)
	QueuesInput() *[]*string
	RawOverrides() interface{}
	SelfManagedEventSource() LambdaEventSourceMappingSelfManagedEventSourceOutputReference
	SelfManagedEventSourceInput() *LambdaEventSourceMappingSelfManagedEventSource
	SourceAccessConfiguration() interface{}
	SetSourceAccessConfiguration(val interface{})
	SourceAccessConfigurationInput() interface{}
	StartingPosition() *string
	SetStartingPosition(val *string)
	StartingPositionInput() *string
	StartingPositionTimestamp() *string
	SetStartingPositionTimestamp(val *string)
	StartingPositionTimestampInput() *string
	State() *string
	StateTransitionReason() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Topics() *[]*string
	SetTopics(val *[]*string)
	TopicsInput() *[]*string
	TumblingWindowInSeconds() *float64
	SetTumblingWindowInSeconds(val *float64)
	TumblingWindowInSecondsInput() *float64
	Uuid() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDestinationConfig(value *LambdaEventSourceMappingDestinationConfig)
	PutFilterCriteria(value *LambdaEventSourceMappingFilterCriteria)
	PutSelfManagedEventSource(value *LambdaEventSourceMappingSelfManagedEventSource)
	ResetBatchSize()
	ResetBisectBatchOnFunctionError()
	ResetDestinationConfig()
	ResetEnabled()
	ResetEventSourceArn()
	ResetFilterCriteria()
	ResetFunctionResponseTypes()
	ResetMaximumBatchingWindowInSeconds()
	ResetMaximumRecordAgeInSeconds()
	ResetMaximumRetryAttempts()
	ResetOverrideLogicalId()
	ResetParallelizationFactor()
	ResetQueues()
	ResetSelfManagedEventSource()
	ResetSourceAccessConfiguration()
	ResetStartingPosition()
	ResetStartingPositionTimestamp()
	ResetTopics()
	ResetTumblingWindowInSeconds()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaEventSourceMapping
type jsiiProxy_LambdaEventSourceMapping struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BisectBatchOnFunctionErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DestinationConfig() LambdaEventSourceMappingDestinationConfigOutputReference {
	var returns LambdaEventSourceMappingDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DestinationConfigInput() *LambdaEventSourceMappingDestinationConfig {
	var returns *LambdaEventSourceMappingDestinationConfig
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EventSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FilterCriteria() LambdaEventSourceMappingFilterCriteriaOutputReference {
	var returns LambdaEventSourceMappingFilterCriteriaOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FilterCriteriaInput() *LambdaEventSourceMappingFilterCriteria {
	var returns *LambdaEventSourceMappingFilterCriteria
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) LastProcessingResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProcessingResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) QueuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedEventSource() LambdaEventSourceMappingSelfManagedEventSourceOutputReference {
	var returns LambdaEventSourceMappingSelfManagedEventSourceOutputReference
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedEventSourceInput() *LambdaEventSourceMappingSelfManagedEventSource {
	var returns *LambdaEventSourceMappingSelfManagedEventSource
	_jsii_.Get(
		j,
		"selfManagedEventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SourceAccessConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SourceAccessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StateTransitionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateTransitionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TumblingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
func NewLambdaEventSourceMapping(scope constructs.Construct, id *string, config *LambdaEventSourceMappingConfig) LambdaEventSourceMapping {
	_init_.Initialize()

	j := jsiiProxy_LambdaEventSourceMapping{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
func NewLambdaEventSourceMapping_Override(l LambdaEventSourceMapping, scope constructs.Construct, id *string, config *LambdaEventSourceMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetBisectBatchOnFunctionError(val interface{}) {
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetEventSourceArn(val *string) {
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetFunctionResponseTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetMaximumBatchingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetMaximumRecordAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetMaximumRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetParallelizationFactor(val *float64) {
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetQueues(val *[]*string) {
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetSourceAccessConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sourceAccessConfiguration",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetStartingPosition(val *string) {
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetStartingPositionTimestamp(val *string) {
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetTopics(val *[]*string) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping) SetTumblingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaEventSourceMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutDestinationConfig(value *LambdaEventSourceMappingDestinationConfig) {
	_jsii_.InvokeVoid(
		l,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutFilterCriteria(value *LambdaEventSourceMappingFilterCriteria) {
	_jsii_.InvokeVoid(
		l,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutSelfManagedEventSource(value *LambdaEventSourceMappingSelfManagedEventSource) {
	_jsii_.InvokeVoid(
		l,
		"putSelfManagedEventSource",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetBatchSize() {
	_jsii_.InvokeVoid(
		l,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetBisectBatchOnFunctionError() {
	_jsii_.InvokeVoid(
		l,
		"resetBisectBatchOnFunctionError",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetEventSourceArn() {
	_jsii_.InvokeVoid(
		l,
		"resetEventSourceArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		l,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetFunctionResponseTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionResponseTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		l,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetQueues() {
	_jsii_.InvokeVoid(
		l,
		"resetQueues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetSelfManagedEventSource() {
	_jsii_.InvokeVoid(
		l,
		"resetSelfManagedEventSource",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetSourceAccessConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceAccessConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		l,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetStartingPositionTimestamp() {
	_jsii_.InvokeVoid(
		l,
		"resetStartingPositionTimestamp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetTopics() {
	_jsii_.InvokeVoid(
		l,
		"resetTopics",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetTumblingWindowInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetTumblingWindowInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaEventSourceMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaEventSourceMappingConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#function_name LambdaEventSourceMapping#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#batch_size LambdaEventSourceMapping#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#bisect_batch_on_function_error LambdaEventSourceMapping#bisect_batch_on_function_error}.
	BisectBatchOnFunctionError interface{} `json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#destination_config LambdaEventSourceMapping#destination_config}
	DestinationConfig *LambdaEventSourceMappingDestinationConfig `json:"destinationConfig" yaml:"destinationConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#enabled LambdaEventSourceMapping#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#event_source_arn LambdaEventSourceMapping#event_source_arn}.
	EventSourceArn *string `json:"eventSourceArn" yaml:"eventSourceArn"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#filter_criteria LambdaEventSourceMapping#filter_criteria}
	FilterCriteria *LambdaEventSourceMappingFilterCriteria `json:"filterCriteria" yaml:"filterCriteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#function_response_types LambdaEventSourceMapping#function_response_types}.
	FunctionResponseTypes *[]*string `json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#maximum_batching_window_in_seconds LambdaEventSourceMapping#maximum_batching_window_in_seconds}.
	MaximumBatchingWindowInSeconds *float64 `json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#maximum_record_age_in_seconds LambdaEventSourceMapping#maximum_record_age_in_seconds}.
	MaximumRecordAgeInSeconds *float64 `json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#maximum_retry_attempts LambdaEventSourceMapping#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#parallelization_factor LambdaEventSourceMapping#parallelization_factor}.
	ParallelizationFactor *float64 `json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#queues LambdaEventSourceMapping#queues}.
	Queues *[]*string `json:"queues" yaml:"queues"`
	// self_managed_event_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#self_managed_event_source LambdaEventSourceMapping#self_managed_event_source}
	SelfManagedEventSource *LambdaEventSourceMappingSelfManagedEventSource `json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// source_access_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#source_access_configuration LambdaEventSourceMapping#source_access_configuration}
	SourceAccessConfiguration interface{} `json:"sourceAccessConfiguration" yaml:"sourceAccessConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#starting_position LambdaEventSourceMapping#starting_position}.
	StartingPosition *string `json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#starting_position_timestamp LambdaEventSourceMapping#starting_position_timestamp}.
	StartingPositionTimestamp *string `json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#topics LambdaEventSourceMapping#topics}.
	Topics *[]*string `json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#tumbling_window_in_seconds LambdaEventSourceMapping#tumbling_window_in_seconds}.
	TumblingWindowInSeconds *float64 `json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}

type LambdaEventSourceMappingDestinationConfig struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#on_failure LambdaEventSourceMapping#on_failure}
	OnFailure *LambdaEventSourceMappingDestinationConfigOnFailure `json:"onFailure" yaml:"onFailure"`
}

type LambdaEventSourceMappingDestinationConfigOnFailure struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#destination_arn LambdaEventSourceMapping#destination_arn}.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
}

type LambdaEventSourceMappingDestinationConfigOnFailureOutputReference interface {
	cdktf.ComplexObject
	DestinationArn() *string
	SetDestinationArn(val *string)
	DestinationArnInput() *string
	InternalValue() *LambdaEventSourceMappingDestinationConfigOnFailure
	SetInternalValue(val *LambdaEventSourceMappingDestinationConfigOnFailure)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaEventSourceMappingDestinationConfigOnFailureOutputReference
type jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) DestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) InternalValue() *LambdaEventSourceMappingDestinationConfigOnFailure {
	var returns *LambdaEventSourceMappingDestinationConfigOnFailure
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaEventSourceMappingDestinationConfigOnFailureOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaEventSourceMappingDestinationConfigOnFailureOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingDestinationConfigOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaEventSourceMappingDestinationConfigOnFailureOutputReference_Override(l LambdaEventSourceMappingDestinationConfigOnFailureOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingDestinationConfigOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) SetDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationArn",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) SetInternalValue(val *LambdaEventSourceMappingDestinationConfigOnFailure) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOnFailureOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaEventSourceMappingDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaEventSourceMappingDestinationConfig
	SetInternalValue(val *LambdaEventSourceMappingDestinationConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OnFailure() LambdaEventSourceMappingDestinationConfigOnFailureOutputReference
	OnFailureInput() *LambdaEventSourceMappingDestinationConfigOnFailure
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutOnFailure(value *LambdaEventSourceMappingDestinationConfigOnFailure)
	ResetOnFailure()
}

// The jsii proxy struct for LambdaEventSourceMappingDestinationConfigOutputReference
type jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) InternalValue() *LambdaEventSourceMappingDestinationConfig {
	var returns *LambdaEventSourceMappingDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) OnFailure() LambdaEventSourceMappingDestinationConfigOnFailureOutputReference {
	var returns LambdaEventSourceMappingDestinationConfigOnFailureOutputReference
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) OnFailureInput() *LambdaEventSourceMappingDestinationConfigOnFailure {
	var returns *LambdaEventSourceMappingDestinationConfigOnFailure
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaEventSourceMappingDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaEventSourceMappingDestinationConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaEventSourceMappingDestinationConfigOutputReference_Override(l LambdaEventSourceMappingDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) SetInternalValue(val *LambdaEventSourceMappingDestinationConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) PutOnFailure(value *LambdaEventSourceMappingDestinationConfigOnFailure) {
	_jsii_.InvokeVoid(
		l,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMappingDestinationConfigOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		l,
		"resetOnFailure",
		nil, // no parameters
	)
}

type LambdaEventSourceMappingFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#filter LambdaEventSourceMapping#filter}
	Filter interface{} `json:"filter" yaml:"filter"`
}

type LambdaEventSourceMappingFilterCriteriaFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#pattern LambdaEventSourceMapping#pattern}.
	Pattern *string `json:"pattern" yaml:"pattern"`
}

type LambdaEventSourceMappingFilterCriteriaOutputReference interface {
	cdktf.ComplexObject
	Filter() interface{}
	SetFilter(val interface{})
	FilterInput() interface{}
	InternalValue() *LambdaEventSourceMappingFilterCriteria
	SetInternalValue(val *LambdaEventSourceMappingFilterCriteria)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetFilter()
}

// The jsii proxy struct for LambdaEventSourceMappingFilterCriteriaOutputReference
type jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) Filter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) InternalValue() *LambdaEventSourceMappingFilterCriteria {
	var returns *LambdaEventSourceMappingFilterCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaEventSourceMappingFilterCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaEventSourceMappingFilterCriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaEventSourceMappingFilterCriteriaOutputReference_Override(l LambdaEventSourceMappingFilterCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) SetFilter(val interface{}) {
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) SetInternalValue(val *LambdaEventSourceMappingFilterCriteria) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMappingFilterCriteriaOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetFilter",
		nil, // no parameters
	)
}

type LambdaEventSourceMappingSelfManagedEventSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}.
	Endpoints *map[string]*string `json:"endpoints" yaml:"endpoints"`
}

type LambdaEventSourceMappingSelfManagedEventSourceOutputReference interface {
	cdktf.ComplexObject
	Endpoints() *map[string]*string
	SetEndpoints(val *map[string]*string)
	EndpointsInput() *map[string]*string
	InternalValue() *LambdaEventSourceMappingSelfManagedEventSource
	SetInternalValue(val *LambdaEventSourceMappingSelfManagedEventSource)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaEventSourceMappingSelfManagedEventSourceOutputReference
type jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) Endpoints() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) EndpointsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) InternalValue() *LambdaEventSourceMappingSelfManagedEventSource {
	var returns *LambdaEventSourceMappingSelfManagedEventSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaEventSourceMappingSelfManagedEventSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaEventSourceMappingSelfManagedEventSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingSelfManagedEventSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaEventSourceMappingSelfManagedEventSourceOutputReference_Override(l LambdaEventSourceMappingSelfManagedEventSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaEventSourceMappingSelfManagedEventSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) SetEndpoints(val *map[string]*string) {
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) SetInternalValue(val *LambdaEventSourceMappingSelfManagedEventSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaEventSourceMappingSelfManagedEventSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaEventSourceMappingSourceAccessConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#type LambdaEventSourceMapping#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#uri LambdaEventSourceMapping#uri}.
	Uri *string `json:"uri" yaml:"uri"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_function aws_lambda_function}.
type LambdaFunction interface {
	cdktf.TerraformResource
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	CodeSigningConfigArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference
	DeadLetterConfigInput() *LambdaFunctionDeadLetterConfig
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Environment() LambdaFunctionEnvironmentOutputReference
	EnvironmentInput() *LambdaFunctionEnvironment
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	FileSystemConfig() LambdaFunctionFileSystemConfigOutputReference
	FileSystemConfigInput() *LambdaFunctionFileSystemConfig
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	ImageConfig() LambdaFunctionImageConfigOutputReference
	ImageConfigInput() *LambdaFunctionImageConfig
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InvokeArn() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	LastModified() *string
	Layers() *[]*string
	SetLayers(val *[]*string)
	LayersInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MemorySize() *float64
	SetMemorySize(val *float64)
	MemorySizeInput() *float64
	Node() constructs.Node
	PackageType() *string
	SetPackageType(val *string)
	PackageTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Publish() interface{}
	SetPublish(val interface{})
	PublishInput() interface{}
	QualifiedArn() *string
	RawOverrides() interface{}
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	ReservedConcurrentExecutionsInput() *float64
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3ObjectVersion() *string
	SetS3ObjectVersion(val *string)
	S3ObjectVersionInput() *string
	SigningJobArn() *string
	SigningProfileVersionArn() *string
	SourceCodeHash() *string
	SetSourceCodeHash(val *string)
	SourceCodeHashInput() *string
	SourceCodeSize() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() LambdaFunctionTimeoutsOutputReference
	TimeoutsInput() *LambdaFunctionTimeouts
	TracingConfig() LambdaFunctionTracingConfigOutputReference
	TracingConfigInput() *LambdaFunctionTracingConfig
	Version() *string
	VpcConfig() LambdaFunctionVpcConfigOutputReference
	VpcConfigInput() *LambdaFunctionVpcConfig
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig)
	PutEnvironment(value *LambdaFunctionEnvironment)
	PutFileSystemConfig(value *LambdaFunctionFileSystemConfig)
	PutImageConfig(value *LambdaFunctionImageConfig)
	PutTimeouts(value *LambdaFunctionTimeouts)
	PutTracingConfig(value *LambdaFunctionTracingConfig)
	PutVpcConfig(value *LambdaFunctionVpcConfig)
	ResetArchitectures()
	ResetCodeSigningConfigArn()
	ResetDeadLetterConfig()
	ResetDescription()
	ResetEnvironment()
	ResetFilename()
	ResetFileSystemConfig()
	ResetHandler()
	ResetImageConfig()
	ResetImageUri()
	ResetKmsKeyArn()
	ResetLayers()
	ResetMemorySize()
	ResetOverrideLogicalId()
	ResetPackageType()
	ResetPublish()
	ResetReservedConcurrentExecutions()
	ResetRuntime()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3ObjectVersion()
	ResetSourceCodeHash()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
	ResetTimeouts()
	ResetTracingConfig()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference {
	var returns LambdaFunctionDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfigInput() *LambdaFunctionDeadLetterConfig {
	var returns *LambdaFunctionDeadLetterConfig
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Environment() LambdaFunctionEnvironmentOutputReference {
	var returns LambdaFunctionEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EnvironmentInput() *LambdaFunctionEnvironment {
	var returns *LambdaFunctionEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfig() LambdaFunctionFileSystemConfigOutputReference {
	var returns LambdaFunctionFileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigInput() *LambdaFunctionFileSystemConfig {
	var returns *LambdaFunctionFileSystemConfig
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfig() LambdaFunctionImageConfigOutputReference {
	var returns LambdaFunctionImageConfigOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfigInput() *LambdaFunctionImageConfig {
	var returns *LambdaFunctionImageConfig
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Publish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) QualifiedArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeouts() LambdaFunctionTimeoutsOutputReference {
	var returns LambdaFunctionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutsInput() *LambdaFunctionTimeouts {
	var returns *LambdaFunctionTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfig() LambdaFunctionTracingConfigOutputReference {
	var returns LambdaFunctionTracingConfigOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfigInput() *LambdaFunctionTracingConfig {
	var returns *LambdaFunctionTracingConfig
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfig() LambdaFunctionVpcConfigOutputReference {
	var returns LambdaFunctionVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfigInput() *LambdaFunctionVpcConfig {
	var returns *LambdaFunctionVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_function aws_lambda_function} Resource.
func NewLambdaFunction(scope constructs.Construct, id *string, config *LambdaFunctionConfig) LambdaFunction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_function aws_lambda_function} Resource.
func NewLambdaFunction_Override(l LambdaFunction, scope constructs.Construct, id *string, config *LambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunction",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaFunction) SetArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetCodeSigningConfigArn(val *string) {
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetHandler(val *string) {
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetLayers(val *[]*string) {
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetMemorySize(val *float64) {
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetPackageType(val *string) {
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetPublish(val interface{}) {
	_jsii_.Set(
		j,
		"publish",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetReservedConcurrentExecutions(val *float64) {
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetS3Key(val *string) {
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetSourceCodeHash(val *string) {
	_jsii_.Set(
		j,
		"sourceCodeHash",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaFunction) PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig) {
	_jsii_.InvokeVoid(
		l,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEnvironment(value *LambdaFunctionEnvironment) {
	_jsii_.InvokeVoid(
		l,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutFileSystemConfig(value *LambdaFunctionFileSystemConfig) {
	_jsii_.InvokeVoid(
		l,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutImageConfig(value *LambdaFunctionImageConfig) {
	_jsii_.InvokeVoid(
		l,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTimeouts(value *LambdaFunctionTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTracingConfig(value *LambdaFunctionTracingConfig) {
	_jsii_.InvokeVoid(
		l,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutVpcConfig(value *LambdaFunctionVpcConfig) {
	_jsii_.InvokeVoid(
		l,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		l,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		l,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFilename() {
	_jsii_.InvokeVoid(
		l,
		"resetFilename",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		l,
		"resetHandler",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageUri() {
	_jsii_.InvokeVoid(
		l,
		"resetImageUri",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		l,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		l,
		"resetLayers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		l,
		"resetMemorySize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		l,
		"resetPackageType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPublish() {
	_jsii_.InvokeVoid(
		l,
		"resetPublish",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		l,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3Key() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Key",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSourceCodeHash() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceCodeHash",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaFunctionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#function_name LambdaFunction#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#role LambdaFunction#role}.
	Role *string `json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#architectures LambdaFunction#architectures}.
	Architectures *[]*string `json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#code_signing_config_arn LambdaFunction#code_signing_config_arn}.
	CodeSigningConfigArn *string `json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#dead_letter_config LambdaFunction#dead_letter_config}
	DeadLetterConfig *LambdaFunctionDeadLetterConfig `json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#description LambdaFunction#description}.
	Description *string `json:"description" yaml:"description"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#environment LambdaFunction#environment}
	Environment *LambdaFunctionEnvironment `json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#filename LambdaFunction#filename}.
	Filename *string `json:"filename" yaml:"filename"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#file_system_config LambdaFunction#file_system_config}
	FileSystemConfig *LambdaFunctionFileSystemConfig `json:"fileSystemConfig" yaml:"fileSystemConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#handler LambdaFunction#handler}.
	Handler *string `json:"handler" yaml:"handler"`
	// image_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#image_config LambdaFunction#image_config}
	ImageConfig *LambdaFunctionImageConfig `json:"imageConfig" yaml:"imageConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#image_uri LambdaFunction#image_uri}.
	ImageUri *string `json:"imageUri" yaml:"imageUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#kms_key_arn LambdaFunction#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#layers LambdaFunction#layers}.
	Layers *[]*string `json:"layers" yaml:"layers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#memory_size LambdaFunction#memory_size}.
	MemorySize *float64 `json:"memorySize" yaml:"memorySize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#package_type LambdaFunction#package_type}.
	PackageType *string `json:"packageType" yaml:"packageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#publish LambdaFunction#publish}.
	Publish interface{} `json:"publish" yaml:"publish"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#reserved_concurrent_executions LambdaFunction#reserved_concurrent_executions}.
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#runtime LambdaFunction#runtime}.
	Runtime *string `json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#s3_bucket LambdaFunction#s3_bucket}.
	S3Bucket *string `json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#s3_key LambdaFunction#s3_key}.
	S3Key *string `json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#s3_object_version LambdaFunction#s3_object_version}.
	S3ObjectVersion *string `json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#source_code_hash LambdaFunction#source_code_hash}.
	SourceCodeHash *string `json:"sourceCodeHash" yaml:"sourceCodeHash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#tags LambdaFunction#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#tags_all LambdaFunction#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#timeout LambdaFunction#timeout}.
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#timeouts LambdaFunction#timeouts}
	Timeouts *LambdaFunctionTimeouts `json:"timeouts" yaml:"timeouts"`
	// tracing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#tracing_config LambdaFunction#tracing_config}
	TracingConfig *LambdaFunctionTracingConfig `json:"tracingConfig" yaml:"tracingConfig"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#vpc_config LambdaFunction#vpc_config}
	VpcConfig *LambdaFunctionVpcConfig `json:"vpcConfig" yaml:"vpcConfig"`
}

type LambdaFunctionDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#target_arn LambdaFunction#target_arn}.
	TargetArn *string `json:"targetArn" yaml:"targetArn"`
}

type LambdaFunctionDeadLetterConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaFunctionDeadLetterConfig
	SetInternalValue(val *LambdaFunctionDeadLetterConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TargetArn() *string
	SetTargetArn(val *string)
	TargetArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionDeadLetterConfigOutputReference
type jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) InternalValue() *LambdaFunctionDeadLetterConfig {
	var returns *LambdaFunctionDeadLetterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) TargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionDeadLetterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionDeadLetterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionDeadLetterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionDeadLetterConfigOutputReference_Override(l LambdaFunctionDeadLetterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionDeadLetterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) SetInternalValue(val *LambdaFunctionDeadLetterConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) SetTargetArn(val *string) {
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionDeadLetterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaFunctionEnvironment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#variables LambdaFunction#variables}.
	Variables *map[string]*string `json:"variables" yaml:"variables"`
}

type LambdaFunctionEnvironmentOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaFunctionEnvironment
	SetInternalValue(val *LambdaFunctionEnvironment)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetVariables()
}

// The jsii proxy struct for LambdaFunctionEnvironmentOutputReference
type jsiiProxy_LambdaFunctionEnvironmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) InternalValue() *LambdaFunctionEnvironment {
	var returns *LambdaFunctionEnvironment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}

func NewLambdaFunctionEnvironmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionEnvironmentOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionEnvironmentOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionEnvironmentOutputReference_Override(l LambdaFunctionEnvironmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) SetInternalValue(val *LambdaFunctionEnvironment) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEnvironmentOutputReference) SetVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnvironmentOutputReference) ResetVariables() {
	_jsii_.InvokeVoid(
		l,
		"resetVariables",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config aws_lambda_function_event_invoke_config}.
type LambdaFunctionEventInvokeConfig interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DestinationConfig() LambdaFunctionEventInvokeConfigDestinationConfigOutputReference
	DestinationConfigInput() *LambdaFunctionEventInvokeConfigDestinationConfig
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumEventAgeInSeconds() *float64
	SetMaximumEventAgeInSeconds(val *float64)
	MaximumEventAgeInSecondsInput() *float64
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	MaximumRetryAttemptsInput() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDestinationConfig(value *LambdaFunctionEventInvokeConfigDestinationConfig)
	ResetDestinationConfig()
	ResetMaximumEventAgeInSeconds()
	ResetMaximumRetryAttempts()
	ResetOverrideLogicalId()
	ResetQualifier()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaFunctionEventInvokeConfig
type jsiiProxy_LambdaFunctionEventInvokeConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) DestinationConfig() LambdaFunctionEventInvokeConfigDestinationConfigOutputReference {
	var returns LambdaFunctionEventInvokeConfigDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) DestinationConfigInput() *LambdaFunctionEventInvokeConfigDestinationConfig {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfig
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) MaximumEventAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEventAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) MaximumEventAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEventAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config aws_lambda_function_event_invoke_config} Resource.
func NewLambdaFunctionEventInvokeConfig(scope constructs.Construct, id *string, config *LambdaFunctionEventInvokeConfigConfig) LambdaFunctionEventInvokeConfig {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionEventInvokeConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config aws_lambda_function_event_invoke_config} Resource.
func NewLambdaFunctionEventInvokeConfig_Override(l LambdaFunctionEventInvokeConfig, scope constructs.Construct, id *string, config *LambdaFunctionEventInvokeConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfig",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetMaximumEventAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumEventAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetMaximumRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfig) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaFunctionEventInvokeConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaFunctionEventInvokeConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) PutDestinationConfig(value *LambdaFunctionEventInvokeConfigDestinationConfig) {
	_jsii_.InvokeVoid(
		l,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ResetMaximumEventAgeInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumEventAgeInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ResetQualifier() {
	_jsii_.InvokeVoid(
		l,
		"resetQualifier",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaFunctionEventInvokeConfigConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#function_name LambdaFunctionEventInvokeConfig#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#destination_config LambdaFunctionEventInvokeConfig#destination_config}
	DestinationConfig *LambdaFunctionEventInvokeConfigDestinationConfig `json:"destinationConfig" yaml:"destinationConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#maximum_event_age_in_seconds LambdaFunctionEventInvokeConfig#maximum_event_age_in_seconds}.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#maximum_retry_attempts LambdaFunctionEventInvokeConfig#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#qualifier LambdaFunctionEventInvokeConfig#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
}

type LambdaFunctionEventInvokeConfigDestinationConfig struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#on_failure LambdaFunctionEventInvokeConfig#on_failure}
	OnFailure *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure `json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#on_success LambdaFunctionEventInvokeConfig#on_success}
	OnSuccess *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess `json:"onSuccess" yaml:"onSuccess"`
}

type LambdaFunctionEventInvokeConfigDestinationConfigOnFailure struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#destination LambdaFunctionEventInvokeConfig#destination}.
	Destination *string `json:"destination" yaml:"destination"`
}

type LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference interface {
	cdktf.ComplexObject
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure
	SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference
type jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference_Override(l LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function_event_invoke_config#destination LambdaFunctionEventInvokeConfig#destination}.
	Destination *string `json:"destination" yaml:"destination"`
}

type LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference interface {
	cdktf.ComplexObject
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess
	SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference
type jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference_Override(l LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaFunctionEventInvokeConfigDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfig
	SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OnFailure() LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference
	OnFailureInput() *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure
	OnSuccess() LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference
	OnSuccessInput() *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutOnFailure(value *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure)
	PutOnSuccess(value *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess)
	ResetOnFailure()
	ResetOnSuccess()
}

// The jsii proxy struct for LambdaFunctionEventInvokeConfigDestinationConfigOutputReference
type jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) InternalValue() *LambdaFunctionEventInvokeConfigDestinationConfig {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) OnFailure() LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference {
	var returns LambdaFunctionEventInvokeConfigDestinationConfigOnFailureOutputReference
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) OnFailureInput() *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) OnSuccess() LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference {
	var returns LambdaFunctionEventInvokeConfigDestinationConfigOnSuccessOutputReference
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) OnSuccessInput() *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess {
	var returns *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionEventInvokeConfigDestinationConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionEventInvokeConfigDestinationConfigOutputReference_Override(l LambdaFunctionEventInvokeConfigDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionEventInvokeConfigDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) SetInternalValue(val *LambdaFunctionEventInvokeConfigDestinationConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) PutOnFailure(value *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure) {
	_jsii_.InvokeVoid(
		l,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) PutOnSuccess(value *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess) {
	_jsii_.InvokeVoid(
		l,
		"putOnSuccess",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		l,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionEventInvokeConfigDestinationConfigOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		l,
		"resetOnSuccess",
		nil, // no parameters
	)
}

type LambdaFunctionFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#arn LambdaFunction#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#local_mount_path LambdaFunction#local_mount_path}.
	LocalMountPath *string `json:"localMountPath" yaml:"localMountPath"`
}

type LambdaFunctionFileSystemConfigOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	InternalValue() *LambdaFunctionFileSystemConfig
	SetInternalValue(val *LambdaFunctionFileSystemConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LocalMountPath() *string
	SetLocalMountPath(val *string)
	LocalMountPathInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionFileSystemConfigOutputReference
type jsiiProxy_LambdaFunctionFileSystemConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) InternalValue() *LambdaFunctionFileSystemConfig {
	var returns *LambdaFunctionFileSystemConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) LocalMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) LocalMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionFileSystemConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionFileSystemConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionFileSystemConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionFileSystemConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionFileSystemConfigOutputReference_Override(l LambdaFunctionFileSystemConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionFileSystemConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetInternalValue(val *LambdaFunctionFileSystemConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetLocalMountPath(val *string) {
	_jsii_.Set(
		j,
		"localMountPath",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionFileSystemConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaFunctionImageConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#command LambdaFunction#command}.
	Command *[]*string `json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#entry_point LambdaFunction#entry_point}.
	EntryPoint *[]*string `json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#working_directory LambdaFunction#working_directory}.
	WorkingDirectory *string `json:"workingDirectory" yaml:"workingDirectory"`
}

type LambdaFunctionImageConfigOutputReference interface {
	cdktf.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
	EntryPoint() *[]*string
	SetEntryPoint(val *[]*string)
	EntryPointInput() *[]*string
	InternalValue() *LambdaFunctionImageConfig
	SetInternalValue(val *LambdaFunctionImageConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCommand()
	ResetEntryPoint()
	ResetWorkingDirectory()
}

// The jsii proxy struct for LambdaFunctionImageConfigOutputReference
type jsiiProxy_LambdaFunctionImageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) InternalValue() *LambdaFunctionImageConfig {
	var returns *LambdaFunctionImageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}

func NewLambdaFunctionImageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionImageConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionImageConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionImageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionImageConfigOutputReference_Override(l LambdaFunctionImageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionImageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetCommand(val *[]*string) {
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetEntryPoint(val *[]*string) {
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetInternalValue(val *LambdaFunctionImageConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionImageConfigOutputReference) SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		l,
		"resetCommand",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		l,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionImageConfigOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

type LambdaFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#create LambdaFunction#create}.
	Create *string `json:"create" yaml:"create"`
}

type LambdaFunctionTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	InternalValue() *LambdaFunctionTimeouts
	SetInternalValue(val *LambdaFunctionTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
}

// The jsii proxy struct for LambdaFunctionTimeoutsOutputReference
type jsiiProxy_LambdaFunctionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) InternalValue() *LambdaFunctionTimeouts {
	var returns *LambdaFunctionTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionTimeoutsOutputReference_Override(l LambdaFunctionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) SetInternalValue(val *LambdaFunctionTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

type LambdaFunctionTracingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#mode LambdaFunction#mode}.
	Mode *string `json:"mode" yaml:"mode"`
}

type LambdaFunctionTracingConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaFunctionTracingConfig
	SetInternalValue(val *LambdaFunctionTracingConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionTracingConfigOutputReference
type jsiiProxy_LambdaFunctionTracingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) InternalValue() *LambdaFunctionTracingConfig {
	var returns *LambdaFunctionTracingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLambdaFunctionTracingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionTracingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionTracingConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionTracingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionTracingConfigOutputReference_Override(l LambdaFunctionTracingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionTracingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) SetInternalValue(val *LambdaFunctionTracingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionTracingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionTracingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LambdaFunctionVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#security_group_ids LambdaFunction#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#subnet_ids LambdaFunction#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type LambdaFunctionVpcConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LambdaFunctionVpcConfig
	SetInternalValue(val *LambdaFunctionVpcConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for LambdaFunctionVpcConfigOutputReference
type jsiiProxy_LambdaFunctionVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) InternalValue() *LambdaFunctionVpcConfig {
	var returns *LambdaFunctionVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func NewLambdaFunctionVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaFunctionVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaFunctionVpcConfigOutputReference_Override(l LambdaFunctionVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaFunctionVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetInternalValue(val *LambdaFunctionVpcConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaFunctionVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation aws_lambda_invocation}.
type LambdaInvocation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Id() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	Result() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetQualifier()
	ResetTriggers()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaInvocation
type jsiiProxy_LambdaInvocation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaInvocation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvocation) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation aws_lambda_invocation} Resource.
func NewLambdaInvocation(scope constructs.Construct, id *string, config *LambdaInvocationConfig) LambdaInvocation {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvocation{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaInvocation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation aws_lambda_invocation} Resource.
func NewLambdaInvocation_Override(l LambdaInvocation, scope constructs.Construct, id *string, config *LambdaInvocationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaInvocation",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetInput(val *string) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

func (j *jsiiProxy_LambdaInvocation) SetTriggers(val *map[string]*string) {
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaInvocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaInvocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaInvocation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaInvocation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaInvocation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaInvocation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvocation) ResetQualifier() {
	_jsii_.InvokeVoid(
		l,
		"resetQualifier",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvocation) ResetTriggers() {
	_jsii_.InvokeVoid(
		l,
		"resetTriggers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvocation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaInvocation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaInvocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaInvocation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaInvocationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation#function_name LambdaInvocation#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation#input LambdaInvocation#input}.
	Input *string `json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation#qualifier LambdaInvocation#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_invocation#triggers LambdaInvocation#triggers}.
	Triggers *map[string]*string `json:"triggers" yaml:"triggers"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version aws_lambda_layer_version}.
type LambdaLayerVersion interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CompatibleArchitectures() *[]*string
	SetCompatibleArchitectures(val *[]*string)
	CompatibleArchitecturesInput() *[]*string
	CompatibleRuntimes() *[]*string
	SetCompatibleRuntimes(val *[]*string)
	CompatibleRuntimesInput() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LayerArn() *string
	LayerName() *string
	SetLayerName(val *string)
	LayerNameInput() *string
	LicenseInfo() *string
	SetLicenseInfo(val *string)
	LicenseInfoInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3ObjectVersion() *string
	SetS3ObjectVersion(val *string)
	S3ObjectVersionInput() *string
	SigningJobArn() *string
	SigningProfileVersionArn() *string
	SkipDestroy() interface{}
	SetSkipDestroy(val interface{})
	SkipDestroyInput() interface{}
	SourceCodeHash() *string
	SetSourceCodeHash(val *string)
	SourceCodeHashInput() *string
	SourceCodeSize() *float64
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetCompatibleArchitectures()
	ResetCompatibleRuntimes()
	ResetDescription()
	ResetFilename()
	ResetLicenseInfo()
	ResetOverrideLogicalId()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3ObjectVersion()
	ResetSkipDestroy()
	ResetSourceCodeHash()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaLayerVersion
type jsiiProxy_LambdaLayerVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaLayerVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CompatibleArchitectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleArchitectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CompatibleArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleArchitecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CompatibleRuntimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CompatibleRuntimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleRuntimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) LayerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) LayerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) LayerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) LicenseInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) LicenseInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SourceCodeHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version aws_lambda_layer_version} Resource.
func NewLambdaLayerVersion(scope constructs.Construct, id *string, config *LambdaLayerVersionConfig) LambdaLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_LambdaLayerVersion{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaLayerVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version aws_lambda_layer_version} Resource.
func NewLambdaLayerVersion_Override(l LambdaLayerVersion, scope constructs.Construct, id *string, config *LambdaLayerVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaLayerVersion",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetCompatibleArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleArchitectures",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetCompatibleRuntimes(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleRuntimes",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetLayerName(val *string) {
	_jsii_.Set(
		j,
		"layerName",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetLicenseInfo(val *string) {
	_jsii_.Set(
		j,
		"licenseInfo",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetS3Key(val *string) {
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetSkipDestroy(val interface{}) {
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersion) SetSourceCodeHash(val *string) {
	_jsii_.Set(
		j,
		"sourceCodeHash",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaLayerVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaLayerVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetCompatibleArchitectures() {
	_jsii_.InvokeVoid(
		l,
		"resetCompatibleArchitectures",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetCompatibleRuntimes() {
	_jsii_.InvokeVoid(
		l,
		"resetCompatibleRuntimes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetFilename() {
	_jsii_.InvokeVoid(
		l,
		"resetFilename",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetLicenseInfo() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseInfo",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetS3Key() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Key",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		l,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) ResetSourceCodeHash() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceCodeHash",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaLayerVersionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#layer_name LambdaLayerVersion#layer_name}.
	LayerName *string `json:"layerName" yaml:"layerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#compatible_architectures LambdaLayerVersion#compatible_architectures}.
	CompatibleArchitectures *[]*string `json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#compatible_runtimes LambdaLayerVersion#compatible_runtimes}.
	CompatibleRuntimes *[]*string `json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#description LambdaLayerVersion#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#filename LambdaLayerVersion#filename}.
	Filename *string `json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#license_info LambdaLayerVersion#license_info}.
	LicenseInfo *string `json:"licenseInfo" yaml:"licenseInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#s3_bucket LambdaLayerVersion#s3_bucket}.
	S3Bucket *string `json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#s3_key LambdaLayerVersion#s3_key}.
	S3Key *string `json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#s3_object_version LambdaLayerVersion#s3_object_version}.
	S3ObjectVersion *string `json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#skip_destroy LambdaLayerVersion#skip_destroy}.
	SkipDestroy interface{} `json:"skipDestroy" yaml:"skipDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version#source_code_hash LambdaLayerVersion#source_code_hash}.
	SourceCodeHash *string `json:"sourceCodeHash" yaml:"sourceCodeHash"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission aws_lambda_layer_version_permission}.
type LambdaLayerVersionPermission interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LayerName() *string
	SetLayerName(val *string)
	LayerNameInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
	Policy() *string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RevisionId() *string
	StatementId() *string
	SetStatementId(val *string)
	StatementIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VersionNumber() *float64
	SetVersionNumber(val *float64)
	VersionNumberInput() *float64
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOrganizationId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaLayerVersionPermission
type jsiiProxy_LambdaLayerVersionPermission struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) LayerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) LayerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) RevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) StatementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) StatementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) VersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaLayerVersionPermission) VersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionNumberInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission aws_lambda_layer_version_permission} Resource.
func NewLambdaLayerVersionPermission(scope constructs.Construct, id *string, config *LambdaLayerVersionPermissionConfig) LambdaLayerVersionPermission {
	_init_.Initialize()

	j := jsiiProxy_LambdaLayerVersionPermission{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaLayerVersionPermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission aws_lambda_layer_version_permission} Resource.
func NewLambdaLayerVersionPermission_Override(l LambdaLayerVersionPermission, scope constructs.Construct, id *string, config *LambdaLayerVersionPermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaLayerVersionPermission",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetLayerName(val *string) {
	_jsii_.Set(
		j,
		"layerName",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetOrganizationId(val *string) {
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetStatementId(val *string) {
	_jsii_.Set(
		j,
		"statementId",
		val,
	)
}

func (j *jsiiProxy_LambdaLayerVersionPermission) SetVersionNumber(val *float64) {
	_jsii_.Set(
		j,
		"versionNumber",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaLayerVersionPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaLayerVersionPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaLayerVersionPermission_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaLayerVersionPermission",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaLayerVersionPermission) ResetOrganizationId() {
	_jsii_.InvokeVoid(
		l,
		"resetOrganizationId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaLayerVersionPermission) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaLayerVersionPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaLayerVersionPermission) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaLayerVersionPermissionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#action LambdaLayerVersionPermission#action}.
	Action *string `json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#layer_name LambdaLayerVersionPermission#layer_name}.
	LayerName *string `json:"layerName" yaml:"layerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#principal LambdaLayerVersionPermission#principal}.
	Principal *string `json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#statement_id LambdaLayerVersionPermission#statement_id}.
	StatementId *string `json:"statementId" yaml:"statementId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#version_number LambdaLayerVersionPermission#version_number}.
	VersionNumber *float64 `json:"versionNumber" yaml:"versionNumber"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_layer_version_permission#organization_id LambdaLayerVersionPermission#organization_id}.
	OrganizationId *string `json:"organizationId" yaml:"organizationId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission aws_lambda_permission}.
type LambdaPermission interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EventSourceToken() *string
	SetEventSourceToken(val *string)
	EventSourceTokenInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	SourceAccount() *string
	SetSourceAccount(val *string)
	SourceAccountInput() *string
	SourceArn() *string
	SetSourceArn(val *string)
	SourceArnInput() *string
	StatementId() *string
	SetStatementId(val *string)
	StatementIdInput() *string
	StatementIdPrefix() *string
	SetStatementIdPrefix(val *string)
	StatementIdPrefixInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetEventSourceToken()
	ResetOverrideLogicalId()
	ResetQualifier()
	ResetSourceAccount()
	ResetSourceArn()
	ResetStatementId()
	ResetStatementIdPrefix()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaPermission
type jsiiProxy_LambdaPermission struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaPermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) EventSourceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) EventSourceTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission aws_lambda_permission} Resource.
func NewLambdaPermission(scope constructs.Construct, id *string, config *LambdaPermissionConfig) LambdaPermission {
	_init_.Initialize()

	j := jsiiProxy_LambdaPermission{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaPermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission aws_lambda_permission} Resource.
func NewLambdaPermission_Override(l LambdaPermission, scope constructs.Construct, id *string, config *LambdaPermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaPermission",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaPermission) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetEventSourceToken(val *string) {
	_jsii_.Set(
		j,
		"eventSourceToken",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetSourceAccount(val *string) {
	_jsii_.Set(
		j,
		"sourceAccount",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetSourceArn(val *string) {
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetStatementId(val *string) {
	_jsii_.Set(
		j,
		"statementId",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission) SetStatementIdPrefix(val *string) {
	_jsii_.Set(
		j,
		"statementIdPrefix",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaPermission_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaPermission",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaPermission) ResetEventSourceToken() {
	_jsii_.InvokeVoid(
		l,
		"resetEventSourceToken",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaPermission) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetQualifier() {
	_jsii_.InvokeVoid(
		l,
		"resetQualifier",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetSourceAccount() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceAccount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetSourceArn() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetStatementId() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetStatementIdPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementIdPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaPermission) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaPermission) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaPermissionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#action LambdaPermission#action}.
	Action *string `json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#function_name LambdaPermission#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#principal LambdaPermission#principal}.
	Principal *string `json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#event_source_token LambdaPermission#event_source_token}.
	EventSourceToken *string `json:"eventSourceToken" yaml:"eventSourceToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#qualifier LambdaPermission#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#source_account LambdaPermission#source_account}.
	SourceAccount *string `json:"sourceAccount" yaml:"sourceAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#source_arn LambdaPermission#source_arn}.
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#statement_id LambdaPermission#statement_id}.
	StatementId *string `json:"statementId" yaml:"statementId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_permission#statement_id_prefix LambdaPermission#statement_id_prefix}.
	StatementIdPrefix *string `json:"statementIdPrefix" yaml:"statementIdPrefix"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config aws_lambda_provisioned_concurrency_config}.
type LambdaProvisionedConcurrencyConfig interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedConcurrentExecutions() *float64
	SetProvisionedConcurrentExecutions(val *float64)
	ProvisionedConcurrentExecutionsInput() *float64
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() LambdaProvisionedConcurrencyConfigTimeoutsOutputReference
	TimeoutsInput() *LambdaProvisionedConcurrencyConfigTimeouts
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *LambdaProvisionedConcurrencyConfigTimeouts)
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaProvisionedConcurrencyConfig
type jsiiProxy_LambdaProvisionedConcurrencyConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) ProvisionedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) ProvisionedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) Timeouts() LambdaProvisionedConcurrencyConfigTimeoutsOutputReference {
	var returns LambdaProvisionedConcurrencyConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) TimeoutsInput() *LambdaProvisionedConcurrencyConfigTimeouts {
	var returns *LambdaProvisionedConcurrencyConfigTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config aws_lambda_provisioned_concurrency_config} Resource.
func NewLambdaProvisionedConcurrencyConfig(scope constructs.Construct, id *string, config *LambdaProvisionedConcurrencyConfigConfig) LambdaProvisionedConcurrencyConfig {
	_init_.Initialize()

	j := jsiiProxy_LambdaProvisionedConcurrencyConfig{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config aws_lambda_provisioned_concurrency_config} Resource.
func NewLambdaProvisionedConcurrencyConfig_Override(l LambdaProvisionedConcurrencyConfig, scope constructs.Construct, id *string, config *LambdaProvisionedConcurrencyConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfig",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetProvisionedConcurrentExecutions(val *float64) {
	_jsii_.Set(
		j,
		"provisionedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfig) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaProvisionedConcurrencyConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaProvisionedConcurrencyConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) PutTimeouts(value *LambdaProvisionedConcurrencyConfigTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lambda.
type LambdaProvisionedConcurrencyConfigConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#function_name LambdaProvisionedConcurrencyConfig#function_name}.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#provisioned_concurrent_executions LambdaProvisionedConcurrencyConfig#provisioned_concurrent_executions}.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#qualifier LambdaProvisionedConcurrencyConfig#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#timeouts LambdaProvisionedConcurrencyConfig#timeouts}
	Timeouts *LambdaProvisionedConcurrencyConfigTimeouts `json:"timeouts" yaml:"timeouts"`
}

type LambdaProvisionedConcurrencyConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#create LambdaProvisionedConcurrencyConfig#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_provisioned_concurrency_config#update LambdaProvisionedConcurrencyConfig#update}.
	Update *string `json:"update" yaml:"update"`
}

type LambdaProvisionedConcurrencyConfigTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	InternalValue() *LambdaProvisionedConcurrencyConfigTimeouts
	SetInternalValue(val *LambdaProvisionedConcurrencyConfigTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetUpdate()
}

// The jsii proxy struct for LambdaProvisionedConcurrencyConfigTimeoutsOutputReference
type jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) InternalValue() *LambdaProvisionedConcurrencyConfigTimeouts {
	var returns *LambdaProvisionedConcurrencyConfigTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewLambdaProvisionedConcurrencyConfigTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LambdaProvisionedConcurrencyConfigTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfigTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLambdaProvisionedConcurrencyConfigTimeoutsOutputReference_Override(l LambdaProvisionedConcurrencyConfigTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lambdafunction.LambdaProvisionedConcurrencyConfigTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetInternalValue(val *LambdaProvisionedConcurrencyConfigTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaProvisionedConcurrencyConfigTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdate",
		nil, // no parameters
	)
}
