package apigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/apigateway/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_account aws_api_gateway_account}.
type ApiGatewayAccount interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	CloudwatchRoleArn() *string
	SetCloudwatchRoleArn(val *string)
	CloudwatchRoleArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
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
	ResetCloudwatchRoleArn()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ThrottleSettings(index *string) ApiGatewayAccountThrottleSettings
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayAccount
type jsiiProxy_ApiGatewayAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) CloudwatchRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) CloudwatchRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_account aws_api_gateway_account} Resource.
func NewApiGatewayAccount(scope constructs.Construct, id *string, config *ApiGatewayAccountConfig) ApiGatewayAccount {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayAccount{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_account aws_api_gateway_account} Resource.
func NewApiGatewayAccount_Override(a ApiGatewayAccount, scope constructs.Construct, id *string, config *ApiGatewayAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAccount",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayAccount) SetCloudwatchRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudwatchRoleArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccount) SetProvider(val cdktf.TerraformProvider) {
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
func ApiGatewayAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayAccount) ResetCloudwatchRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchRoleArn",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayAccount) ThrottleSettings(index *string) ApiGatewayAccountThrottleSettings {
	var returns ApiGatewayAccountThrottleSettings

	_jsii_.Invoke(
		a,
		"throttleSettings",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayAccountConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_account#cloudwatch_role_arn ApiGatewayAccount#cloudwatch_role_arn}.
	CloudwatchRoleArn *string `json:"cloudwatchRoleArn" yaml:"cloudwatchRoleArn"`
}

type ApiGatewayAccountThrottleSettings interface {
	cdktf.ComplexComputedList
	BurstLimit() *float64
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	RateLimit() *float64
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

// The jsii proxy struct for ApiGatewayAccountThrottleSettings
type jsiiProxy_ApiGatewayAccountThrottleSettings struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) BurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) RateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewApiGatewayAccountThrottleSettings(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) ApiGatewayAccountThrottleSettings {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayAccountThrottleSettings{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAccountThrottleSettings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayAccountThrottleSettings_Override(a ApiGatewayAccountThrottleSettings, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAccountThrottleSettings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAccountThrottleSettings) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAccountThrottleSettings) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key aws_api_gateway_api_key}.
type ApiGatewayApiKey interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetDescription()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetValue()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayApiKey
type jsiiProxy_ApiGatewayApiKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayApiKey) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayApiKey) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key aws_api_gateway_api_key} Resource.
func NewApiGatewayApiKey(scope constructs.Construct, id *string, config *ApiGatewayApiKeyConfig) ApiGatewayApiKey {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayApiKey{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayApiKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key aws_api_gateway_api_key} Resource.
func NewApiGatewayApiKey_Override(a ApiGatewayApiKey, scope constructs.Construct, id *string, config *ApiGatewayApiKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayApiKey",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayApiKey) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayApiKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayApiKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) ResetValue() {
	_jsii_.InvokeVoid(
		a,
		"resetValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayApiKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayApiKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayApiKeyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#name ApiGatewayApiKey#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#description ApiGatewayApiKey#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#enabled ApiGatewayApiKey#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#tags ApiGatewayApiKey#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#tags_all ApiGatewayApiKey#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_api_key#value ApiGatewayApiKey#value}.
	Value *string `json:"value" yaml:"value"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer aws_api_gateway_authorizer}.
type ApiGatewayAuthorizer interface {
	cdktf.TerraformResource
	AuthorizerCredentials() *string
	SetAuthorizerCredentials(val *string)
	AuthorizerCredentialsInput() *string
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerResultTtlInSecondsInput() *float64
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	AuthorizerUriInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentitySource() *string
	SetIdentitySource(val *string)
	IdentitySourceInput() *string
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	IdentityValidationExpressionInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProviderArns() *[]*string
	SetProviderArns(val *[]*string)
	ProviderArnsInput() *[]*string
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAuthorizerCredentials()
	ResetAuthorizerResultTtlInSeconds()
	ResetAuthorizerUri()
	ResetIdentitySource()
	ResetIdentityValidationExpression()
	ResetOverrideLogicalId()
	ResetProviderArns()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayAuthorizer
type jsiiProxy_ApiGatewayAuthorizer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) IdentitySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) IdentitySourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) IdentityValidationExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) ProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"providerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) ProviderArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"providerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayAuthorizer) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer aws_api_gateway_authorizer} Resource.
func NewApiGatewayAuthorizer(scope constructs.Construct, id *string, config *ApiGatewayAuthorizerConfig) ApiGatewayAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayAuthorizer{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAuthorizer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer aws_api_gateway_authorizer} Resource.
func NewApiGatewayAuthorizer_Override(a ApiGatewayAuthorizer, scope constructs.Construct, id *string, config *ApiGatewayAuthorizerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayAuthorizer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetAuthorizerCredentials(val *string) {
	_jsii_.Set(
		j,
		"authorizerCredentials",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetIdentitySource(val *string) {
	_jsii_.Set(
		j,
		"identitySource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetIdentityValidationExpression(val *string) {
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"providerArns",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayAuthorizer) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayAuthorizer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayAuthorizer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetAuthorizerCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetAuthorizerUri() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetIdentitySource() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentitySource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetIdentityValidationExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityValidationExpression",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetProviderArns() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayAuthorizer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayAuthorizer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayAuthorizerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#name ApiGatewayAuthorizer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#rest_api_id ApiGatewayAuthorizer#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#authorizer_credentials ApiGatewayAuthorizer#authorizer_credentials}.
	AuthorizerCredentials *string `json:"authorizerCredentials" yaml:"authorizerCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#authorizer_result_ttl_in_seconds ApiGatewayAuthorizer#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#authorizer_uri ApiGatewayAuthorizer#authorizer_uri}.
	AuthorizerUri *string `json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#identity_source ApiGatewayAuthorizer#identity_source}.
	IdentitySource *string `json:"identitySource" yaml:"identitySource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#identity_validation_expression ApiGatewayAuthorizer#identity_validation_expression}.
	IdentityValidationExpression *string `json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#provider_arns ApiGatewayAuthorizer#provider_arns}.
	ProviderArns *[]*string `json:"providerArns" yaml:"providerArns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_authorizer#type ApiGatewayAuthorizer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping aws_api_gateway_base_path_mapping}.
type ApiGatewayBasePathMapping interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	BasePath() *string
	SetBasePath(val *string)
	BasePathInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
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
	ResetBasePath()
	ResetOverrideLogicalId()
	ResetStageName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayBasePathMapping
type jsiiProxy_ApiGatewayBasePathMapping struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) BasePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) BasePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping aws_api_gateway_base_path_mapping} Resource.
func NewApiGatewayBasePathMapping(scope constructs.Construct, id *string, config *ApiGatewayBasePathMappingConfig) ApiGatewayBasePathMapping {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayBasePathMapping{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayBasePathMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping aws_api_gateway_base_path_mapping} Resource.
func NewApiGatewayBasePathMapping_Override(a ApiGatewayBasePathMapping, scope constructs.Construct, id *string, config *ApiGatewayBasePathMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayBasePathMapping",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetBasePath(val *string) {
	_jsii_.Set(
		j,
		"basePath",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayBasePathMapping) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayBasePathMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayBasePathMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayBasePathMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayBasePathMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayBasePathMapping) ResetBasePath() {
	_jsii_.InvokeVoid(
		a,
		"resetBasePath",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayBasePathMapping) ResetStageName() {
	_jsii_.InvokeVoid(
		a,
		"resetStageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayBasePathMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayBasePathMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayBasePathMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayBasePathMappingConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping#api_id ApiGatewayBasePathMapping#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping#domain_name ApiGatewayBasePathMapping#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping#base_path ApiGatewayBasePathMapping#base_path}.
	BasePath *string `json:"basePath" yaml:"basePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_base_path_mapping#stage_name ApiGatewayBasePathMapping#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate aws_api_gateway_client_certificate}.
type ApiGatewayClientCertificate interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExpirationDate() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PemEncodedCertificate() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayClientCertificate
type jsiiProxy_ApiGatewayClientCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) PemEncodedCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemEncodedCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayClientCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate aws_api_gateway_client_certificate} Resource.
func NewApiGatewayClientCertificate(scope constructs.Construct, id *string, config *ApiGatewayClientCertificateConfig) ApiGatewayClientCertificate {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayClientCertificate{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayClientCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate aws_api_gateway_client_certificate} Resource.
func NewApiGatewayClientCertificate_Override(a ApiGatewayClientCertificate, scope constructs.Construct, id *string, config *ApiGatewayClientCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayClientCertificate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayClientCertificate) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayClientCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayClientCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayClientCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayClientCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayClientCertificate) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayClientCertificate) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayClientCertificate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayClientCertificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayClientCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayClientCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayClientCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate#description ApiGatewayClientCertificate#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate#tags ApiGatewayClientCertificate#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_client_certificate#tags_all ApiGatewayClientCertificate#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment aws_api_gateway_deployment}.
type ApiGatewayDeployment interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExecutionArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InvokeUrl() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StageDescription() *string
	SetStageDescription(val *string)
	StageDescriptionInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetStageDescription()
	ResetStageName()
	ResetTriggers()
	ResetVariables()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayDeployment
type jsiiProxy_ApiGatewayDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) InvokeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) StageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) StageDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDeployment) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment aws_api_gateway_deployment} Resource.
func NewApiGatewayDeployment(scope constructs.Construct, id *string, config *ApiGatewayDeploymentConfig) ApiGatewayDeployment {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDeployment{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment aws_api_gateway_deployment} Resource.
func NewApiGatewayDeployment_Override(a ApiGatewayDeployment, scope constructs.Construct, id *string, config *ApiGatewayDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDeployment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetStageDescription(val *string) {
	_jsii_.Set(
		j,
		"stageDescription",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetTriggers(val *map[string]*string) {
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDeployment) SetVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) ResetStageDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetStageDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) ResetStageName() {
	_jsii_.InvokeVoid(
		a,
		"resetStageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) ResetTriggers() {
	_jsii_.InvokeVoid(
		a,
		"resetTriggers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) ResetVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayDeploymentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#rest_api_id ApiGatewayDeployment#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#description ApiGatewayDeployment#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#stage_description ApiGatewayDeployment#stage_description}.
	StageDescription *string `json:"stageDescription" yaml:"stageDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#stage_name ApiGatewayDeployment#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#triggers ApiGatewayDeployment#triggers}.
	Triggers *map[string]*string `json:"triggers" yaml:"triggers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_deployment#variables ApiGatewayDeployment#variables}.
	Variables *map[string]*string `json:"variables" yaml:"variables"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part aws_api_gateway_documentation_part}.
type ApiGatewayDocumentationPart interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() ApiGatewayDocumentationPartLocationOutputReference
	LocationInput() *ApiGatewayDocumentationPartLocation
	Node() constructs.Node
	Properties() *string
	SetProperties(val *string)
	PropertiesInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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
	PutLocation(value *ApiGatewayDocumentationPartLocation)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayDocumentationPart
type jsiiProxy_ApiGatewayDocumentationPart struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Location() ApiGatewayDocumentationPartLocationOutputReference {
	var returns ApiGatewayDocumentationPartLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) LocationInput() *ApiGatewayDocumentationPartLocation {
	var returns *ApiGatewayDocumentationPartLocation
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) PropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part aws_api_gateway_documentation_part} Resource.
func NewApiGatewayDocumentationPart(scope constructs.Construct, id *string, config *ApiGatewayDocumentationPartConfig) ApiGatewayDocumentationPart {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDocumentationPart{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPart",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part aws_api_gateway_documentation_part} Resource.
func NewApiGatewayDocumentationPart_Override(a ApiGatewayDocumentationPart, scope constructs.Construct, id *string, config *ApiGatewayDocumentationPartConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPart",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetProperties(val *string) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPart) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayDocumentationPart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPart",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayDocumentationPart_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPart",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationPart) PutLocation(value *ApiGatewayDocumentationPartLocation) {
	_jsii_.InvokeVoid(
		a,
		"putLocation",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationPart) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayDocumentationPart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPart) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayDocumentationPartConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#location ApiGatewayDocumentationPart#location}
	Location *ApiGatewayDocumentationPartLocation `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#properties ApiGatewayDocumentationPart#properties}.
	Properties *string `json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#rest_api_id ApiGatewayDocumentationPart#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

type ApiGatewayDocumentationPartLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#type ApiGatewayDocumentationPart#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#method ApiGatewayDocumentationPart#method}.
	Method *string `json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#name ApiGatewayDocumentationPart#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#path ApiGatewayDocumentationPart#path}.
	Path *string `json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_part#status_code ApiGatewayDocumentationPart#status_code}.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
}

type ApiGatewayDocumentationPartLocationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ApiGatewayDocumentationPartLocation
	SetInternalValue(val *ApiGatewayDocumentationPartLocation)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	StatusCode() *string
	SetStatusCode(val *string)
	StatusCodeInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetMethod()
	ResetName()
	ResetPath()
	ResetStatusCode()
}

// The jsii proxy struct for ApiGatewayDocumentationPartLocationOutputReference
type jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) InternalValue() *ApiGatewayDocumentationPartLocation {
	var returns *ApiGatewayDocumentationPartLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) StatusCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewApiGatewayDocumentationPartLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayDocumentationPartLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPartLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayDocumentationPartLocationOutputReference_Override(a ApiGatewayDocumentationPartLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationPartLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetInternalValue(val *ApiGatewayDocumentationPartLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetMethod(val *string) {
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetStatusCode(val *string) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationPartLocationOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		a,
		"resetStatusCode",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version aws_api_gateway_documentation_version}.
type ApiGatewayDocumentationVersion interface {
	cdktf.TerraformResource
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
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetDescription()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayDocumentationVersion
type jsiiProxy_ApiGatewayDocumentationVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version aws_api_gateway_documentation_version} Resource.
func NewApiGatewayDocumentationVersion(scope constructs.Construct, id *string, config *ApiGatewayDocumentationVersionConfig) ApiGatewayDocumentationVersion {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDocumentationVersion{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version aws_api_gateway_documentation_version} Resource.
func NewApiGatewayDocumentationVersion_Override(a ApiGatewayDocumentationVersion, scope constructs.Construct, id *string, config *ApiGatewayDocumentationVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationVersion",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDocumentationVersion) SetVersion(val *string) {
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
func ApiGatewayDocumentationVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayDocumentationVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayDocumentationVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationVersion) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDocumentationVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayDocumentationVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayDocumentationVersionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version#rest_api_id ApiGatewayDocumentationVersion#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version#version ApiGatewayDocumentationVersion#version}.
	Version *string `json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_documentation_version#description ApiGatewayDocumentationVersion#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name aws_api_gateway_domain_name}.
type ApiGatewayDomainName interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	CertificateBody() *string
	SetCertificateBody(val *string)
	CertificateBodyInput() *string
	CertificateChain() *string
	SetCertificateChain(val *string)
	CertificateChainInput() *string
	CertificateName() *string
	SetCertificateName(val *string)
	CertificateNameInput() *string
	CertificatePrivateKey() *string
	SetCertificatePrivateKey(val *string)
	CertificatePrivateKeyInput() *string
	CertificateUploadDate() *string
	CloudfrontDomainName() *string
	CloudfrontZoneId() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EndpointConfiguration() ApiGatewayDomainNameEndpointConfigurationOutputReference
	EndpointConfigurationInput() *ApiGatewayDomainNameEndpointConfiguration
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MutualTlsAuthentication() ApiGatewayDomainNameMutualTlsAuthenticationOutputReference
	MutualTlsAuthenticationInput() *ApiGatewayDomainNameMutualTlsAuthentication
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegionalCertificateArn() *string
	SetRegionalCertificateArn(val *string)
	RegionalCertificateArnInput() *string
	RegionalCertificateName() *string
	SetRegionalCertificateName(val *string)
	RegionalCertificateNameInput() *string
	RegionalDomainName() *string
	RegionalZoneId() *string
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	PutEndpointConfiguration(value *ApiGatewayDomainNameEndpointConfiguration)
	PutMutualTlsAuthentication(value *ApiGatewayDomainNameMutualTlsAuthentication)
	ResetCertificateArn()
	ResetCertificateBody()
	ResetCertificateChain()
	ResetCertificateName()
	ResetCertificatePrivateKey()
	ResetEndpointConfiguration()
	ResetMutualTlsAuthentication()
	ResetOverrideLogicalId()
	ResetRegionalCertificateArn()
	ResetRegionalCertificateName()
	ResetSecurityPolicy()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayDomainName
type jsiiProxy_ApiGatewayDomainName struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayDomainName) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificatePrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificatePrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateUploadDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUploadDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CloudfrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CloudfrontZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) EndpointConfiguration() ApiGatewayDomainNameEndpointConfigurationOutputReference {
	var returns ApiGatewayDomainNameEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) EndpointConfigurationInput() *ApiGatewayDomainNameEndpointConfiguration {
	var returns *ApiGatewayDomainNameEndpointConfiguration
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) MutualTlsAuthentication() ApiGatewayDomainNameMutualTlsAuthenticationOutputReference {
	var returns ApiGatewayDomainNameMutualTlsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) MutualTlsAuthenticationInput() *ApiGatewayDomainNameMutualTlsAuthentication {
	var returns *ApiGatewayDomainNameMutualTlsAuthentication
	_jsii_.Get(
		j,
		"mutualTlsAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
func NewApiGatewayDomainName(scope constructs.Construct, id *string, config *ApiGatewayDomainNameConfig) ApiGatewayDomainName {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomainName{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
func NewApiGatewayDomainName_Override(a ApiGatewayDomainName, scope constructs.Construct, id *string, config *ApiGatewayDomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainName",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCertificateBody(val *string) {
	_jsii_.Set(
		j,
		"certificateBody",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCertificateChain(val *string) {
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCertificateName(val *string) {
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCertificatePrivateKey(val *string) {
	_jsii_.Set(
		j,
		"certificatePrivateKey",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetRegionalCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"regionalCertificateArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetRegionalCertificateName(val *string) {
	_jsii_.Set(
		j,
		"regionalCertificateName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetSecurityPolicy(val *string) {
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayDomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayDomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) PutEndpointConfiguration(value *ApiGatewayDomainNameEndpointConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) PutMutualTlsAuthentication(value *ApiGatewayDomainNameMutualTlsAuthentication) {
	_jsii_.InvokeVoid(
		a,
		"putMutualTlsAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateBody() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateChain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificatePrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetMutualTlsAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualTlsAuthentication",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetRegionalCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetRegionalCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayDomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayDomainNameConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#domain_name ApiGatewayDomainName#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#certificate_arn ApiGatewayDomainName#certificate_arn}.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#certificate_body ApiGatewayDomainName#certificate_body}.
	CertificateBody *string `json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#certificate_chain ApiGatewayDomainName#certificate_chain}.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#certificate_name ApiGatewayDomainName#certificate_name}.
	CertificateName *string `json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#certificate_private_key ApiGatewayDomainName#certificate_private_key}.
	CertificatePrivateKey *string `json:"certificatePrivateKey" yaml:"certificatePrivateKey"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#endpoint_configuration ApiGatewayDomainName#endpoint_configuration}
	EndpointConfiguration *ApiGatewayDomainNameEndpointConfiguration `json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// mutual_tls_authentication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#mutual_tls_authentication ApiGatewayDomainName#mutual_tls_authentication}
	MutualTlsAuthentication *ApiGatewayDomainNameMutualTlsAuthentication `json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#regional_certificate_arn ApiGatewayDomainName#regional_certificate_arn}.
	RegionalCertificateArn *string `json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#regional_certificate_name ApiGatewayDomainName#regional_certificate_name}.
	RegionalCertificateName *string `json:"regionalCertificateName" yaml:"regionalCertificateName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#security_policy ApiGatewayDomainName#security_policy}.
	SecurityPolicy *string `json:"securityPolicy" yaml:"securityPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#tags ApiGatewayDomainName#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#tags_all ApiGatewayDomainName#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type ApiGatewayDomainNameEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#types ApiGatewayDomainName#types}.
	Types *[]*string `json:"types" yaml:"types"`
}

type ApiGatewayDomainNameEndpointConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ApiGatewayDomainNameEndpointConfiguration
	SetInternalValue(val *ApiGatewayDomainNameEndpointConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
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

// The jsii proxy struct for ApiGatewayDomainNameEndpointConfigurationOutputReference
type jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) InternalValue() *ApiGatewayDomainNameEndpointConfiguration {
	var returns *ApiGatewayDomainNameEndpointConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func NewApiGatewayDomainNameEndpointConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayDomainNameEndpointConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainNameEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayDomainNameEndpointConfigurationOutputReference_Override(a ApiGatewayDomainNameEndpointConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainNameEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) SetInternalValue(val *ApiGatewayDomainNameEndpointConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) SetTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameEndpointConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ApiGatewayDomainNameMutualTlsAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#truststore_uri ApiGatewayDomainName#truststore_uri}.
	TruststoreUri *string `json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_domain_name#truststore_version ApiGatewayDomainName#truststore_version}.
	TruststoreVersion *string `json:"truststoreVersion" yaml:"truststoreVersion"`
}

type ApiGatewayDomainNameMutualTlsAuthenticationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ApiGatewayDomainNameMutualTlsAuthentication
	SetInternalValue(val *ApiGatewayDomainNameMutualTlsAuthentication)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TruststoreUri() *string
	SetTruststoreUri(val *string)
	TruststoreUriInput() *string
	TruststoreVersion() *string
	SetTruststoreVersion(val *string)
	TruststoreVersionInput() *string
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
	ResetTruststoreVersion()
}

// The jsii proxy struct for ApiGatewayDomainNameMutualTlsAuthenticationOutputReference
type jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) InternalValue() *ApiGatewayDomainNameMutualTlsAuthentication {
	var returns *ApiGatewayDomainNameMutualTlsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TruststoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TruststoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TruststoreVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) TruststoreVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreVersionInput",
		&returns,
	)
	return returns
}

func NewApiGatewayDomainNameMutualTlsAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayDomainNameMutualTlsAuthenticationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainNameMutualTlsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayDomainNameMutualTlsAuthenticationOutputReference_Override(a ApiGatewayDomainNameMutualTlsAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayDomainNameMutualTlsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetInternalValue(val *ApiGatewayDomainNameMutualTlsAuthentication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetTruststoreUri(val *string) {
	_jsii_.Set(
		j,
		"truststoreUri",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) SetTruststoreVersion(val *string) {
	_jsii_.Set(
		j,
		"truststoreVersion",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainNameMutualTlsAuthenticationOutputReference) ResetTruststoreVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetTruststoreVersion",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response aws_api_gateway_gateway_response}.
type ApiGatewayGatewayResponse interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResponseParameters() *map[string]*string
	SetResponseParameters(val *map[string]*string)
	ResponseParametersInput() *map[string]*string
	ResponseTemplates() *map[string]*string
	SetResponseTemplates(val *map[string]*string)
	ResponseTemplatesInput() *map[string]*string
	ResponseType() *string
	SetResponseType(val *string)
	ResponseTypeInput() *string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StatusCode() *string
	SetStatusCode(val *string)
	StatusCodeInput() *string
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
	ResetResponseParameters()
	ResetResponseTemplates()
	ResetStatusCode()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayGatewayResponse
type jsiiProxy_ApiGatewayGatewayResponse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) ResponseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) StatusCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response aws_api_gateway_gateway_response} Resource.
func NewApiGatewayGatewayResponse(scope constructs.Construct, id *string, config *ApiGatewayGatewayResponseConfig) ApiGatewayGatewayResponse {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayGatewayResponse{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayGatewayResponse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response aws_api_gateway_gateway_response} Resource.
func NewApiGatewayGatewayResponse_Override(a ApiGatewayGatewayResponse, scope constructs.Construct, id *string, config *ApiGatewayGatewayResponseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayGatewayResponse",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetResponseParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetResponseTemplates(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetResponseType(val *string) {
	_jsii_.Set(
		j,
		"responseType",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayGatewayResponse) SetStatusCode(val *string) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayGatewayResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayGatewayResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayGatewayResponse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayGatewayResponse",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayGatewayResponse) ResetResponseParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayGatewayResponse) ResetResponseTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayGatewayResponse) ResetStatusCode() {
	_jsii_.InvokeVoid(
		a,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayGatewayResponse) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayGatewayResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayGatewayResponse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayGatewayResponseConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response#response_type ApiGatewayGatewayResponse#response_type}.
	ResponseType *string `json:"responseType" yaml:"responseType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response#rest_api_id ApiGatewayGatewayResponse#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response#response_parameters ApiGatewayGatewayResponse#response_parameters}.
	ResponseParameters *map[string]*string `json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response#response_templates ApiGatewayGatewayResponse#response_templates}.
	ResponseTemplates *map[string]*string `json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_gateway_response#status_code ApiGatewayGatewayResponse#status_code}.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration aws_api_gateway_integration}.
type ApiGatewayIntegration interface {
	cdktf.TerraformResource
	CacheKeyParameters() *[]*string
	SetCacheKeyParameters(val *[]*string)
	CacheKeyParametersInput() *[]*string
	CacheNamespace() *string
	SetCacheNamespace(val *string)
	CacheNamespaceInput() *string
	CdktfStack() cdktf.TerraformStack
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	ContentHandling() *string
	SetContentHandling(val *string)
	ContentHandlingInput() *string
	Count() *float64
	SetCount(val *float64)
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	IntegrationHttpMethod() *string
	SetIntegrationHttpMethod(val *string)
	IntegrationHttpMethodInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	PassthroughBehaviorInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestParameters() *map[string]*string
	SetRequestParameters(val *map[string]*string)
	RequestParametersInput() *map[string]*string
	RequestTemplates() *map[string]*string
	SetRequestTemplates(val *map[string]*string)
	RequestTemplatesInput() *map[string]*string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TimeoutMilliseconds() *float64
	SetTimeoutMilliseconds(val *float64)
	TimeoutMillisecondsInput() *float64
	TlsConfig() ApiGatewayIntegrationTlsConfigOutputReference
	TlsConfigInput() *ApiGatewayIntegrationTlsConfig
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutTlsConfig(value *ApiGatewayIntegrationTlsConfig)
	ResetCacheKeyParameters()
	ResetCacheNamespace()
	ResetConnectionId()
	ResetConnectionType()
	ResetContentHandling()
	ResetCredentials()
	ResetIntegrationHttpMethod()
	ResetOverrideLogicalId()
	ResetPassthroughBehavior()
	ResetRequestParameters()
	ResetRequestTemplates()
	ResetTimeoutMilliseconds()
	ResetTlsConfig()
	ResetUri()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayIntegration
type jsiiProxy_ApiGatewayIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayIntegration) CacheKeyParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) CacheKeyParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) CacheNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) CacheNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ContentHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ContentHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) IntegrationHttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) IntegrationHttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) PassthroughBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RequestParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RequestParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RequestTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RequestTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TlsConfig() ApiGatewayIntegrationTlsConfigOutputReference {
	var returns ApiGatewayIntegrationTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TlsConfigInput() *ApiGatewayIntegrationTlsConfig {
	var returns *ApiGatewayIntegrationTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegration) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration aws_api_gateway_integration} Resource.
func NewApiGatewayIntegration(scope constructs.Construct, id *string, config *ApiGatewayIntegrationConfig) ApiGatewayIntegration {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayIntegration{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration aws_api_gateway_integration} Resource.
func NewApiGatewayIntegration_Override(a ApiGatewayIntegration, scope constructs.Construct, id *string, config *ApiGatewayIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetCacheKeyParameters(val *[]*string) {
	_jsii_.Set(
		j,
		"cacheKeyParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetCacheNamespace(val *string) {
	_jsii_.Set(
		j,
		"cacheNamespace",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetConnectionId(val *string) {
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetContentHandling(val *string) {
	_jsii_.Set(
		j,
		"contentHandling",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetCredentials(val *string) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetIntegrationHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"integrationHttpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetPassthroughBehavior(val *string) {
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetRequestParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetRequestTemplates(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetTimeoutMilliseconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutMilliseconds",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegration) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) PutTlsConfig(value *ApiGatewayIntegrationTlsConfig) {
	_jsii_.InvokeVoid(
		a,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetCacheKeyParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheKeyParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetCacheNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetConnectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetContentHandling() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetIntegrationHttpMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationHttpMethod",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetPassthroughBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPassthroughBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetRequestTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) ResetUri() {
	_jsii_.InvokeVoid(
		a,
		"resetUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayIntegrationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#http_method ApiGatewayIntegration#http_method}.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#resource_id ApiGatewayIntegration#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#rest_api_id ApiGatewayIntegration#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#type ApiGatewayIntegration#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#cache_key_parameters ApiGatewayIntegration#cache_key_parameters}.
	CacheKeyParameters *[]*string `json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#cache_namespace ApiGatewayIntegration#cache_namespace}.
	CacheNamespace *string `json:"cacheNamespace" yaml:"cacheNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#connection_id ApiGatewayIntegration#connection_id}.
	ConnectionId *string `json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#connection_type ApiGatewayIntegration#connection_type}.
	ConnectionType *string `json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#content_handling ApiGatewayIntegration#content_handling}.
	ContentHandling *string `json:"contentHandling" yaml:"contentHandling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#credentials ApiGatewayIntegration#credentials}.
	Credentials *string `json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#integration_http_method ApiGatewayIntegration#integration_http_method}.
	IntegrationHttpMethod *string `json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#passthrough_behavior ApiGatewayIntegration#passthrough_behavior}.
	PassthroughBehavior *string `json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#request_parameters ApiGatewayIntegration#request_parameters}.
	RequestParameters *map[string]*string `json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#request_templates ApiGatewayIntegration#request_templates}.
	RequestTemplates *map[string]*string `json:"requestTemplates" yaml:"requestTemplates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#timeout_milliseconds ApiGatewayIntegration#timeout_milliseconds}.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#tls_config ApiGatewayIntegration#tls_config}
	TlsConfig *ApiGatewayIntegrationTlsConfig `json:"tlsConfig" yaml:"tlsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#uri ApiGatewayIntegration#uri}.
	Uri *string `json:"uri" yaml:"uri"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response aws_api_gateway_integration_response}.
type ApiGatewayIntegrationResponse interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentHandling() *string
	SetContentHandling(val *string)
	ContentHandlingInput() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResponseParameters() *map[string]*string
	SetResponseParameters(val *map[string]*string)
	ResponseParametersInput() *map[string]*string
	ResponseTemplates() *map[string]*string
	SetResponseTemplates(val *map[string]*string)
	ResponseTemplatesInput() *map[string]*string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	SelectionPattern() *string
	SetSelectionPattern(val *string)
	SelectionPatternInput() *string
	StatusCode() *string
	SetStatusCode(val *string)
	StatusCodeInput() *string
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
	ResetContentHandling()
	ResetOverrideLogicalId()
	ResetResponseParameters()
	ResetResponseTemplates()
	ResetSelectionPattern()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayIntegrationResponse
type jsiiProxy_ApiGatewayIntegrationResponse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ContentHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ContentHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResponseParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResponseParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResponseTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) ResponseTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SelectionPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SelectionPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) StatusCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response aws_api_gateway_integration_response} Resource.
func NewApiGatewayIntegrationResponse(scope constructs.Construct, id *string, config *ApiGatewayIntegrationResponseConfig) ApiGatewayIntegrationResponse {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayIntegrationResponse{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationResponse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response aws_api_gateway_integration_response} Resource.
func NewApiGatewayIntegrationResponse_Override(a ApiGatewayIntegrationResponse, scope constructs.Construct, id *string, config *ApiGatewayIntegrationResponseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationResponse",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetContentHandling(val *string) {
	_jsii_.Set(
		j,
		"contentHandling",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetResponseParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetResponseTemplates(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetSelectionPattern(val *string) {
	_jsii_.Set(
		j,
		"selectionPattern",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationResponse) SetStatusCode(val *string) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayIntegrationResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayIntegrationResponse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationResponse",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayIntegrationResponse) ResetContentHandling() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandling",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegrationResponse) ResetResponseParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegrationResponse) ResetResponseTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegrationResponse) ResetSelectionPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetSelectionPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayIntegrationResponse) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationResponse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayIntegrationResponseConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#http_method ApiGatewayIntegrationResponse#http_method}.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#resource_id ApiGatewayIntegrationResponse#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#rest_api_id ApiGatewayIntegrationResponse#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#status_code ApiGatewayIntegrationResponse#status_code}.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#content_handling ApiGatewayIntegrationResponse#content_handling}.
	ContentHandling *string `json:"contentHandling" yaml:"contentHandling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#response_parameters ApiGatewayIntegrationResponse#response_parameters}.
	ResponseParameters *map[string]*string `json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#response_templates ApiGatewayIntegrationResponse#response_templates}.
	ResponseTemplates *map[string]*string `json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration_response#selection_pattern ApiGatewayIntegrationResponse#selection_pattern}.
	SelectionPattern *string `json:"selectionPattern" yaml:"selectionPattern"`
}

type ApiGatewayIntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_integration#insecure_skip_verification ApiGatewayIntegration#insecure_skip_verification}.
	InsecureSkipVerification interface{} `json:"insecureSkipVerification" yaml:"insecureSkipVerification"`
}

type ApiGatewayIntegrationTlsConfigOutputReference interface {
	cdktf.ComplexObject
	InsecureSkipVerification() interface{}
	SetInsecureSkipVerification(val interface{})
	InsecureSkipVerificationInput() interface{}
	InternalValue() *ApiGatewayIntegrationTlsConfig
	SetInternalValue(val *ApiGatewayIntegrationTlsConfig)
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
	ResetInsecureSkipVerification()
}

// The jsii proxy struct for ApiGatewayIntegrationTlsConfigOutputReference
type jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) InsecureSkipVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) InsecureSkipVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) InternalValue() *ApiGatewayIntegrationTlsConfig {
	var returns *ApiGatewayIntegrationTlsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApiGatewayIntegrationTlsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayIntegrationTlsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayIntegrationTlsConfigOutputReference_Override(a ApiGatewayIntegrationTlsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayIntegrationTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) SetInsecureSkipVerification(val interface{}) {
	_jsii_.Set(
		j,
		"insecureSkipVerification",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) SetInternalValue(val *ApiGatewayIntegrationTlsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference) ResetInsecureSkipVerification() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureSkipVerification",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method}.
type ApiGatewayMethod interface {
	cdktf.TerraformResource
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	ApiKeyRequiredInput() interface{}
	Authorization() *string
	SetAuthorization(val *string)
	AuthorizationInput() *string
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	AuthorizationScopesInput() *[]*string
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestModels() *map[string]*string
	SetRequestModels(val *map[string]*string)
	RequestModelsInput() *map[string]*string
	RequestParameters() *map[string]interface{}
	SetRequestParameters(val *map[string]interface{})
	RequestParametersInput() *map[string]interface{}
	RequestValidatorId() *string
	SetRequestValidatorId(val *string)
	RequestValidatorIdInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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
	ResetApiKeyRequired()
	ResetAuthorizationScopes()
	ResetAuthorizerId()
	ResetOperationName()
	ResetOverrideLogicalId()
	ResetRequestModels()
	ResetRequestParameters()
	ResetRequestValidatorId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayMethod
type jsiiProxy_ApiGatewayMethod struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayMethod) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ApiKeyRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Authorization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizationScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestParameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestParametersInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RequestValidatorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethod) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method} Resource.
func NewApiGatewayMethod(scope constructs.Construct, id *string, config *ApiGatewayMethodConfig) ApiGatewayMethod {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayMethod{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethod",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method aws_api_gateway_method} Resource.
func NewApiGatewayMethod_Override(a ApiGatewayMethod, scope constructs.Construct, id *string, config *ApiGatewayMethodConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethod",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetApiKeyRequired(val interface{}) {
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetAuthorization(val *string) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetAuthorizationScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetAuthorizerId(val *string) {
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetOperationName(val *string) {
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetRequestModels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetRequestParameters(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetRequestValidatorId(val *string) {
	_jsii_.Set(
		j,
		"requestValidatorId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethod) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayMethod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayMethod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayMethod_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayMethod",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetApiKeyRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeyRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetAuthorizationScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetOperationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestModels() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) ResetRequestValidatorId() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestValidatorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethod) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethod) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayMethodConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#authorization ApiGatewayMethod#authorization}.
	Authorization *string `json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#http_method ApiGatewayMethod#http_method}.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#resource_id ApiGatewayMethod#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#rest_api_id ApiGatewayMethod#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#api_key_required ApiGatewayMethod#api_key_required}.
	ApiKeyRequired interface{} `json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#authorization_scopes ApiGatewayMethod#authorization_scopes}.
	AuthorizationScopes *[]*string `json:"authorizationScopes" yaml:"authorizationScopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#authorizer_id ApiGatewayMethod#authorizer_id}.
	AuthorizerId *string `json:"authorizerId" yaml:"authorizerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#operation_name ApiGatewayMethod#operation_name}.
	OperationName *string `json:"operationName" yaml:"operationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#request_models ApiGatewayMethod#request_models}.
	RequestModels *map[string]*string `json:"requestModels" yaml:"requestModels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#request_parameters ApiGatewayMethod#request_parameters}.
	RequestParameters *map[string]interface{} `json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method#request_validator_id ApiGatewayMethod#request_validator_id}.
	RequestValidatorId *string `json:"requestValidatorId" yaml:"requestValidatorId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response aws_api_gateway_method_response}.
type ApiGatewayMethodResponse interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResponseModels() *map[string]*string
	SetResponseModels(val *map[string]*string)
	ResponseModelsInput() *map[string]*string
	ResponseParameters() *map[string]interface{}
	SetResponseParameters(val *map[string]interface{})
	ResponseParametersInput() *map[string]interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StatusCode() *string
	SetStatusCode(val *string)
	StatusCodeInput() *string
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
	ResetResponseModels()
	ResetResponseParameters()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayMethodResponse
type jsiiProxy_ApiGatewayMethodResponse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayMethodResponse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResponseModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResponseModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResponseParameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) ResponseParametersInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"responseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) StatusCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodResponse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response aws_api_gateway_method_response} Resource.
func NewApiGatewayMethodResponse(scope constructs.Construct, id *string, config *ApiGatewayMethodResponseConfig) ApiGatewayMethodResponse {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayMethodResponse{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodResponse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response aws_api_gateway_method_response} Resource.
func NewApiGatewayMethodResponse_Override(a ApiGatewayMethodResponse, scope constructs.Construct, id *string, config *ApiGatewayMethodResponseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodResponse",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetResponseModels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseModels",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetResponseParameters(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodResponse) SetStatusCode(val *string) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayMethodResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayMethodResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayMethodResponse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayMethodResponse",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodResponse) ResetResponseModels() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodResponse) ResetResponseParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodResponse) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayMethodResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodResponse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayMethodResponseConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#http_method ApiGatewayMethodResponse#http_method}.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#resource_id ApiGatewayMethodResponse#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#rest_api_id ApiGatewayMethodResponse#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#status_code ApiGatewayMethodResponse#status_code}.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#response_models ApiGatewayMethodResponse#response_models}.
	ResponseModels *map[string]*string `json:"responseModels" yaml:"responseModels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_response#response_parameters ApiGatewayMethodResponse#response_parameters}.
	ResponseParameters *map[string]interface{} `json:"responseParameters" yaml:"responseParameters"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings aws_api_gateway_method_settings}.
type ApiGatewayMethodSettings interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MethodPath() *string
	SetMethodPath(val *string)
	MethodPathInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	Settings() ApiGatewayMethodSettingsSettingsOutputReference
	SettingsInput() *ApiGatewayMethodSettingsSettings
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
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
	PutSettings(value *ApiGatewayMethodSettingsSettings)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayMethodSettings
type jsiiProxy_ApiGatewayMethodSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayMethodSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) MethodPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) MethodPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) Settings() ApiGatewayMethodSettingsSettingsOutputReference {
	var returns ApiGatewayMethodSettingsSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SettingsInput() *ApiGatewayMethodSettingsSettings {
	var returns *ApiGatewayMethodSettingsSettings
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings aws_api_gateway_method_settings} Resource.
func NewApiGatewayMethodSettings(scope constructs.Construct, id *string, config *ApiGatewayMethodSettingsConfig) ApiGatewayMethodSettings {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayMethodSettings{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings aws_api_gateway_method_settings} Resource.
func NewApiGatewayMethodSettings_Override(a ApiGatewayMethodSettings, scope constructs.Construct, id *string, config *ApiGatewayMethodSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettings",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetMethodPath(val *string) {
	_jsii_.Set(
		j,
		"methodPath",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettings) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayMethodSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayMethodSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettings) PutSettings(value *ApiGatewayMethodSettingsSettings) {
	_jsii_.InvokeVoid(
		a,
		"putSettings",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayMethodSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayMethodSettingsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#method_path ApiGatewayMethodSettings#method_path}.
	MethodPath *string `json:"methodPath" yaml:"methodPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#rest_api_id ApiGatewayMethodSettings#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#settings ApiGatewayMethodSettings#settings}
	Settings *ApiGatewayMethodSettingsSettings `json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#stage_name ApiGatewayMethodSettings#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
}

type ApiGatewayMethodSettingsSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#cache_data_encrypted ApiGatewayMethodSettings#cache_data_encrypted}.
	CacheDataEncrypted interface{} `json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#cache_ttl_in_seconds ApiGatewayMethodSettings#cache_ttl_in_seconds}.
	CacheTtlInSeconds *float64 `json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#caching_enabled ApiGatewayMethodSettings#caching_enabled}.
	CachingEnabled interface{} `json:"cachingEnabled" yaml:"cachingEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#data_trace_enabled ApiGatewayMethodSettings#data_trace_enabled}.
	DataTraceEnabled interface{} `json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#logging_level ApiGatewayMethodSettings#logging_level}.
	LoggingLevel *string `json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#metrics_enabled ApiGatewayMethodSettings#metrics_enabled}.
	MetricsEnabled interface{} `json:"metricsEnabled" yaml:"metricsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#require_authorization_for_cache_control ApiGatewayMethodSettings#require_authorization_for_cache_control}.
	RequireAuthorizationForCacheControl interface{} `json:"requireAuthorizationForCacheControl" yaml:"requireAuthorizationForCacheControl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#throttling_burst_limit ApiGatewayMethodSettings#throttling_burst_limit}.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#throttling_rate_limit ApiGatewayMethodSettings#throttling_rate_limit}.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_method_settings#unauthorized_cache_control_header_strategy ApiGatewayMethodSettings#unauthorized_cache_control_header_strategy}.
	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy" yaml:"unauthorizedCacheControlHeaderStrategy"`
}

type ApiGatewayMethodSettingsSettingsOutputReference interface {
	cdktf.ComplexObject
	CacheDataEncrypted() interface{}
	SetCacheDataEncrypted(val interface{})
	CacheDataEncryptedInput() interface{}
	CacheTtlInSeconds() *float64
	SetCacheTtlInSeconds(val *float64)
	CacheTtlInSecondsInput() *float64
	CachingEnabled() interface{}
	SetCachingEnabled(val interface{})
	CachingEnabledInput() interface{}
	DataTraceEnabled() interface{}
	SetDataTraceEnabled(val interface{})
	DataTraceEnabledInput() interface{}
	InternalValue() *ApiGatewayMethodSettingsSettings
	SetInternalValue(val *ApiGatewayMethodSettingsSettings)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
	MetricsEnabled() interface{}
	SetMetricsEnabled(val interface{})
	MetricsEnabledInput() interface{}
	RequireAuthorizationForCacheControl() interface{}
	SetRequireAuthorizationForCacheControl(val interface{})
	RequireAuthorizationForCacheControlInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThrottlingBurstLimit() *float64
	SetThrottlingBurstLimit(val *float64)
	ThrottlingBurstLimitInput() *float64
	ThrottlingRateLimit() *float64
	SetThrottlingRateLimit(val *float64)
	ThrottlingRateLimitInput() *float64
	UnauthorizedCacheControlHeaderStrategy() *string
	SetUnauthorizedCacheControlHeaderStrategy(val *string)
	UnauthorizedCacheControlHeaderStrategyInput() *string
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
	ResetCacheDataEncrypted()
	ResetCacheTtlInSeconds()
	ResetCachingEnabled()
	ResetDataTraceEnabled()
	ResetLoggingLevel()
	ResetMetricsEnabled()
	ResetRequireAuthorizationForCacheControl()
	ResetThrottlingBurstLimit()
	ResetThrottlingRateLimit()
	ResetUnauthorizedCacheControlHeaderStrategy()
}

// The jsii proxy struct for ApiGatewayMethodSettingsSettingsOutputReference
type jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheDataEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheDataEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CachingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CachingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InternalValue() *ApiGatewayMethodSettingsSettings {
	var returns *ApiGatewayMethodSettingsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) RequireAuthorizationForCacheControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) RequireAuthorizationForCacheControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) UnauthorizedCacheControlHeaderStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) UnauthorizedCacheControlHeaderStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategyInput",
		&returns,
	)
	return returns
}

func NewApiGatewayMethodSettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayMethodSettingsSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayMethodSettingsSettingsOutputReference_Override(a ApiGatewayMethodSettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayMethodSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetCacheDataEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"cacheDataEncrypted",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetCacheTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"cacheTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetCachingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cachingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetDataTraceEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetInternalValue(val *ApiGatewayMethodSettingsSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetLoggingLevel(val *string) {
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetMetricsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetRequireAuthorizationForCacheControl(val interface{}) {
	_jsii_.Set(
		j,
		"requireAuthorizationForCacheControl",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetThrottlingBurstLimit(val *float64) {
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetThrottlingRateLimit(val *float64) {
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) SetUnauthorizedCacheControlHeaderStrategy(val *string) {
	_jsii_.Set(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCacheDataEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheDataEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCacheTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCachingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetRequireAuthorizationForCacheControl() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireAuthorizationForCacheControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetUnauthorizedCacheControlHeaderStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthorizedCacheControlHeaderStrategy",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model aws_api_gateway_model}.
type ApiGatewayModel interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetSchema()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayModel
type jsiiProxy_ApiGatewayModel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayModel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayModel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model aws_api_gateway_model} Resource.
func NewApiGatewayModel(scope constructs.Construct, id *string, config *ApiGatewayModelConfig) ApiGatewayModel {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayModel{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayModel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model aws_api_gateway_model} Resource.
func NewApiGatewayModel_Override(a ApiGatewayModel, scope constructs.Construct, id *string, config *ApiGatewayModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayModel",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayModel) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayModel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayModel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayModel) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayModel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayModel) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayModel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayModel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayModel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayModelConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model#content_type ApiGatewayModel#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model#name ApiGatewayModel#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model#rest_api_id ApiGatewayModel#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model#description ApiGatewayModel#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_model#schema ApiGatewayModel#schema}.
	Schema *string `json:"schema" yaml:"schema"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator aws_api_gateway_request_validator}.
type ApiGatewayRequestValidator interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ValidateRequestBody() interface{}
	SetValidateRequestBody(val interface{})
	ValidateRequestBodyInput() interface{}
	ValidateRequestParameters() interface{}
	SetValidateRequestParameters(val interface{})
	ValidateRequestParametersInput() interface{}
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
	ResetValidateRequestBody()
	ResetValidateRequestParameters()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayRequestValidator
type jsiiProxy_ApiGatewayRequestValidator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayRequestValidator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) ValidateRequestBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) ValidateRequestBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) ValidateRequestParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRequestValidator) ValidateRequestParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestParametersInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator aws_api_gateway_request_validator} Resource.
func NewApiGatewayRequestValidator(scope constructs.Construct, id *string, config *ApiGatewayRequestValidatorConfig) ApiGatewayRequestValidator {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayRequestValidator{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRequestValidator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator aws_api_gateway_request_validator} Resource.
func NewApiGatewayRequestValidator_Override(a ApiGatewayRequestValidator, scope constructs.Construct, id *string, config *ApiGatewayRequestValidatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRequestValidator",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetValidateRequestBody(val interface{}) {
	_jsii_.Set(
		j,
		"validateRequestBody",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRequestValidator) SetValidateRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"validateRequestParameters",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayRequestValidator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayRequestValidator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayRequestValidator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayRequestValidator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRequestValidator) ResetValidateRequestBody() {
	_jsii_.InvokeVoid(
		a,
		"resetValidateRequestBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRequestValidator) ResetValidateRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetValidateRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRequestValidator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayRequestValidator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayRequestValidator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayRequestValidatorConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator#name ApiGatewayRequestValidator#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator#rest_api_id ApiGatewayRequestValidator#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator#validate_request_body ApiGatewayRequestValidator#validate_request_body}.
	ValidateRequestBody interface{} `json:"validateRequestBody" yaml:"validateRequestBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_request_validator#validate_request_parameters ApiGatewayRequestValidator#validate_request_parameters}.
	ValidateRequestParameters interface{} `json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource aws_api_gateway_resource}.
type ApiGatewayResource interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	Path() *string
	PathPart() *string
	SetPathPart(val *string)
	PathPartInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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

// The jsii proxy struct for ApiGatewayResource
type jsiiProxy_ApiGatewayResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) PathPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) PathPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource aws_api_gateway_resource} Resource.
func NewApiGatewayResource(scope constructs.Construct, id *string, config *ApiGatewayResourceConfig) ApiGatewayResource {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayResource{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource aws_api_gateway_resource} Resource.
func NewApiGatewayResource_Override(a ApiGatewayResource, scope constructs.Construct, id *string, config *ApiGatewayResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayResource",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetParentId(val *string) {
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetPathPart(val *string) {
	_jsii_.Set(
		j,
		"pathPart",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayResource) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayResourceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource#parent_id ApiGatewayResource#parent_id}.
	ParentId *string `json:"parentId" yaml:"parentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource#path_part ApiGatewayResource#path_part}.
	PathPart *string `json:"pathPart" yaml:"pathPart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_resource#rest_api_id ApiGatewayResource#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api aws_api_gateway_rest_api}.
type ApiGatewayRestApi interface {
	cdktf.TerraformResource
	ApiKeySource() *string
	SetApiKeySource(val *string)
	ApiKeySourceInput() *string
	Arn() *string
	BinaryMediaTypes() *[]*string
	SetBinaryMediaTypes(val *[]*string)
	BinaryMediaTypesInput() *[]*string
	Body() *string
	SetBody(val *string)
	BodyInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	DisableExecuteApiEndpointInput() interface{}
	EndpointConfiguration() ApiGatewayRestApiEndpointConfigurationOutputReference
	EndpointConfigurationInput() *ApiGatewayRestApiEndpointConfiguration
	ExecutionArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinimumCompressionSize() *float64
	SetMinimumCompressionSize(val *float64)
	MinimumCompressionSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootResourceId() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	PutEndpointConfiguration(value *ApiGatewayRestApiEndpointConfiguration)
	ResetApiKeySource()
	ResetBinaryMediaTypes()
	ResetBody()
	ResetDescription()
	ResetDisableExecuteApiEndpoint()
	ResetEndpointConfiguration()
	ResetMinimumCompressionSize()
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPolicy()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayRestApi
type jsiiProxy_ApiGatewayRestApi struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayRestApi) ApiKeySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) ApiKeySourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) BinaryMediaTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) DisableExecuteApiEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) EndpointConfiguration() ApiGatewayRestApiEndpointConfigurationOutputReference {
	var returns ApiGatewayRestApiEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) EndpointConfigurationInput() *ApiGatewayRestApiEndpointConfiguration {
	var returns *ApiGatewayRestApiEndpointConfiguration
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) MinimumCompressionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) RootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api aws_api_gateway_rest_api} Resource.
func NewApiGatewayRestApi(scope constructs.Construct, id *string, config *ApiGatewayRestApiConfig) ApiGatewayRestApi {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayRestApi{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api aws_api_gateway_rest_api} Resource.
func NewApiGatewayRestApi_Override(a ApiGatewayRestApi, scope constructs.Construct, id *string, config *ApiGatewayRestApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetApiKeySource(val *string) {
	_jsii_.Set(
		j,
		"apiKeySource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetBinaryMediaTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetBody(val *string) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetMinimumCompressionSize(val *float64) {
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApi) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayRestApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayRestApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) PutEndpointConfiguration(value *ApiGatewayRestApiEndpointConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetApiKeySource() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeySource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetBinaryMediaTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetBinaryMediaTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetBody() {
	_jsii_.InvokeVoid(
		a,
		"resetBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetDisableExecuteApiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableExecuteApiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetMinimumCompressionSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumCompressionSize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayRestApiConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#name ApiGatewayRestApi#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#api_key_source ApiGatewayRestApi#api_key_source}.
	ApiKeySource *string `json:"apiKeySource" yaml:"apiKeySource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#binary_media_types ApiGatewayRestApi#binary_media_types}.
	BinaryMediaTypes *[]*string `json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#body ApiGatewayRestApi#body}.
	Body *string `json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#description ApiGatewayRestApi#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#disable_execute_api_endpoint ApiGatewayRestApi#disable_execute_api_endpoint}.
	DisableExecuteApiEndpoint interface{} `json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#endpoint_configuration ApiGatewayRestApi#endpoint_configuration}
	EndpointConfiguration *ApiGatewayRestApiEndpointConfiguration `json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#minimum_compression_size ApiGatewayRestApi#minimum_compression_size}.
	MinimumCompressionSize *float64 `json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#parameters ApiGatewayRestApi#parameters}.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#policy ApiGatewayRestApi#policy}.
	Policy *string `json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#tags ApiGatewayRestApi#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#tags_all ApiGatewayRestApi#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type ApiGatewayRestApiEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#types ApiGatewayRestApi#types}.
	Types *[]*string `json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api#vpc_endpoint_ids ApiGatewayRestApi#vpc_endpoint_ids}.
	VpcEndpointIds *[]*string `json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

type ApiGatewayRestApiEndpointConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ApiGatewayRestApiEndpointConfiguration
	SetInternalValue(val *ApiGatewayRestApiEndpointConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
	VpcEndpointIds() *[]*string
	SetVpcEndpointIds(val *[]*string)
	VpcEndpointIdsInput() *[]*string
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
	ResetVpcEndpointIds()
}

// The jsii proxy struct for ApiGatewayRestApiEndpointConfigurationOutputReference
type jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) InternalValue() *ApiGatewayRestApiEndpointConfiguration {
	var returns *ApiGatewayRestApiEndpointConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) VpcEndpointIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcEndpointIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) VpcEndpointIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcEndpointIdsInput",
		&returns,
	)
	return returns
}

func NewApiGatewayRestApiEndpointConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayRestApiEndpointConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApiEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayRestApiEndpointConfigurationOutputReference_Override(a ApiGatewayRestApiEndpointConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApiEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetInternalValue(val *ApiGatewayRestApiEndpointConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) SetVpcEndpointIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcEndpointIds",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayRestApiEndpointConfigurationOutputReference) ResetVpcEndpointIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcEndpointIds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api_policy aws_api_gateway_rest_api_policy}.
type ApiGatewayRestApiPolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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

// The jsii proxy struct for ApiGatewayRestApiPolicy
type jsiiProxy_ApiGatewayRestApiPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api_policy aws_api_gateway_rest_api_policy} Resource.
func NewApiGatewayRestApiPolicy(scope constructs.Construct, id *string, config *ApiGatewayRestApiPolicyConfig) ApiGatewayRestApiPolicy {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayRestApiPolicy{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApiPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api_policy aws_api_gateway_rest_api_policy} Resource.
func NewApiGatewayRestApiPolicy_Override(a ApiGatewayRestApiPolicy, scope constructs.Construct, id *string, config *ApiGatewayRestApiPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayRestApiPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayRestApiPolicy) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayRestApiPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayRestApiPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayRestApiPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayRestApiPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayRestApiPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayRestApiPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayRestApiPolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api_policy#policy ApiGatewayRestApiPolicy#policy}.
	Policy *string `json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_rest_api_policy#rest_api_id ApiGatewayRestApiPolicy#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage aws_api_gateway_stage}.
type ApiGatewayStage interface {
	cdktf.TerraformResource
	AccessLogSettings() ApiGatewayStageAccessLogSettingsOutputReference
	AccessLogSettingsInput() *ApiGatewayStageAccessLogSettings
	Arn() *string
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	CacheClusterEnabledInput() interface{}
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	CacheClusterSizeInput() *string
	CdktfStack() cdktf.TerraformStack
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	ClientCertificateIdInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentId() *string
	SetDeploymentId(val *string)
	DeploymentIdInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentationVersion() *string
	SetDocumentationVersion(val *string)
	DocumentationVersionInput() *string
	ExecutionArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InvokeUrl() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
	WebAclArn() *string
	XrayTracingEnabled() interface{}
	SetXrayTracingEnabled(val interface{})
	XrayTracingEnabledInput() interface{}
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
	PutAccessLogSettings(value *ApiGatewayStageAccessLogSettings)
	ResetAccessLogSettings()
	ResetCacheClusterEnabled()
	ResetCacheClusterSize()
	ResetClientCertificateId()
	ResetDescription()
	ResetDocumentationVersion()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetVariables()
	ResetXrayTracingEnabled()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayStage
type jsiiProxy_ApiGatewayStage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayStage) AccessLogSettings() ApiGatewayStageAccessLogSettingsOutputReference {
	var returns ApiGatewayStageAccessLogSettingsOutputReference
	_jsii_.Get(
		j,
		"accessLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) AccessLogSettingsInput() *ApiGatewayStageAccessLogSettings {
	var returns *ApiGatewayStageAccessLogSettings
	_jsii_.Get(
		j,
		"accessLogSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) CacheClusterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) CacheClusterSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) ClientCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DocumentationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) DocumentationVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) InvokeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) WebAclArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) XrayTracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayTracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStage) XrayTracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayTracingEnabledInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage aws_api_gateway_stage} Resource.
func NewApiGatewayStage(scope constructs.Construct, id *string, config *ApiGatewayStageConfig) ApiGatewayStage {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayStage{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayStage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage aws_api_gateway_stage} Resource.
func NewApiGatewayStage_Override(a ApiGatewayStage, scope constructs.Construct, id *string, config *ApiGatewayStageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayStage",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetCacheClusterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetCacheClusterSize(val *string) {
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetClientCertificateId(val *string) {
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetDeploymentId(val *string) {
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetDocumentationVersion(val *string) {
	_jsii_.Set(
		j,
		"documentationVersion",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStage) SetXrayTracingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"xrayTracingEnabled",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayStage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayStage",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayStage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayStage) PutAccessLogSettings(value *ApiGatewayStageAccessLogSettings) {
	_jsii_.InvokeVoid(
		a,
		"putAccessLogSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetAccessLogSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetCacheClusterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetCacheClusterSize() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheClusterSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetClientCertificateId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetDocumentationVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentationVersion",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayStage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) ResetXrayTracingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetXrayTracingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayStage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayStageAccessLogSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#destination_arn ApiGatewayStage#destination_arn}.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#format ApiGatewayStage#format}.
	Format *string `json:"format" yaml:"format"`
}

type ApiGatewayStageAccessLogSettingsOutputReference interface {
	cdktf.ComplexObject
	DestinationArn() *string
	SetDestinationArn(val *string)
	DestinationArnInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	InternalValue() *ApiGatewayStageAccessLogSettings
	SetInternalValue(val *ApiGatewayStageAccessLogSettings)
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

// The jsii proxy struct for ApiGatewayStageAccessLogSettingsOutputReference
type jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) DestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) InternalValue() *ApiGatewayStageAccessLogSettings {
	var returns *ApiGatewayStageAccessLogSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApiGatewayStageAccessLogSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayStageAccessLogSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayStageAccessLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayStageAccessLogSettingsOutputReference_Override(a ApiGatewayStageAccessLogSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayStageAccessLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetInternalValue(val *ApiGatewayStageAccessLogSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayStageAccessLogSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayStageConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#deployment_id ApiGatewayStage#deployment_id}.
	DeploymentId *string `json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#rest_api_id ApiGatewayStage#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#stage_name ApiGatewayStage#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
	// access_log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#access_log_settings ApiGatewayStage#access_log_settings}
	AccessLogSettings *ApiGatewayStageAccessLogSettings `json:"accessLogSettings" yaml:"accessLogSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#cache_cluster_enabled ApiGatewayStage#cache_cluster_enabled}.
	CacheClusterEnabled interface{} `json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#cache_cluster_size ApiGatewayStage#cache_cluster_size}.
	CacheClusterSize *string `json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#client_certificate_id ApiGatewayStage#client_certificate_id}.
	ClientCertificateId *string `json:"clientCertificateId" yaml:"clientCertificateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#description ApiGatewayStage#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#documentation_version ApiGatewayStage#documentation_version}.
	DocumentationVersion *string `json:"documentationVersion" yaml:"documentationVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#tags ApiGatewayStage#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#tags_all ApiGatewayStage#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#variables ApiGatewayStage#variables}.
	Variables *map[string]*string `json:"variables" yaml:"variables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#xray_tracing_enabled ApiGatewayStage#xray_tracing_enabled}.
	XrayTracingEnabled interface{} `json:"xrayTracingEnabled" yaml:"xrayTracingEnabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan aws_api_gateway_usage_plan}.
type ApiGatewayUsagePlan interface {
	cdktf.TerraformResource
	ApiStages() interface{}
	SetApiStages(val interface{})
	ApiStagesInput() interface{}
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
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ProductCode() *string
	SetProductCode(val *string)
	ProductCodeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	QuotaSettings() ApiGatewayUsagePlanQuotaSettingsOutputReference
	QuotaSettingsInput() *ApiGatewayUsagePlanQuotaSettings
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ThrottleSettings() ApiGatewayUsagePlanThrottleSettingsOutputReference
	ThrottleSettingsInput() *ApiGatewayUsagePlanThrottleSettings
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
	PutQuotaSettings(value *ApiGatewayUsagePlanQuotaSettings)
	PutThrottleSettings(value *ApiGatewayUsagePlanThrottleSettings)
	ResetApiStages()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetProductCode()
	ResetQuotaSettings()
	ResetTags()
	ResetTagsAll()
	ResetThrottleSettings()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayUsagePlan
type jsiiProxy_ApiGatewayUsagePlan struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ApiStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ApiStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ProductCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ProductCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) QuotaSettings() ApiGatewayUsagePlanQuotaSettingsOutputReference {
	var returns ApiGatewayUsagePlanQuotaSettingsOutputReference
	_jsii_.Get(
		j,
		"quotaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) QuotaSettingsInput() *ApiGatewayUsagePlanQuotaSettings {
	var returns *ApiGatewayUsagePlanQuotaSettings
	_jsii_.Get(
		j,
		"quotaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ThrottleSettings() ApiGatewayUsagePlanThrottleSettingsOutputReference {
	var returns ApiGatewayUsagePlanThrottleSettingsOutputReference
	_jsii_.Get(
		j,
		"throttleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlan) ThrottleSettingsInput() *ApiGatewayUsagePlanThrottleSettings {
	var returns *ApiGatewayUsagePlanThrottleSettings
	_jsii_.Get(
		j,
		"throttleSettingsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan aws_api_gateway_usage_plan} Resource.
func NewApiGatewayUsagePlan(scope constructs.Construct, id *string, config *ApiGatewayUsagePlanConfig) ApiGatewayUsagePlan {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayUsagePlan{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlan",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan aws_api_gateway_usage_plan} Resource.
func NewApiGatewayUsagePlan_Override(a ApiGatewayUsagePlan, scope constructs.Construct, id *string, config *ApiGatewayUsagePlanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlan",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetApiStages(val interface{}) {
	_jsii_.Set(
		j,
		"apiStages",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetProductCode(val *string) {
	_jsii_.Set(
		j,
		"productCode",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlan) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayUsagePlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayUsagePlan_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlan",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) PutQuotaSettings(value *ApiGatewayUsagePlanQuotaSettings) {
	_jsii_.InvokeVoid(
		a,
		"putQuotaSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) PutThrottleSettings(value *ApiGatewayUsagePlanThrottleSettings) {
	_jsii_.InvokeVoid(
		a,
		"putThrottleSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetApiStages() {
	_jsii_.InvokeVoid(
		a,
		"resetApiStages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetProductCode() {
	_jsii_.InvokeVoid(
		a,
		"resetProductCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetQuotaSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetQuotaSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) ResetThrottleSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlan) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayUsagePlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlan) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayUsagePlanApiStages struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#api_id ApiGatewayUsagePlan#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#stage ApiGatewayUsagePlan#stage}.
	Stage *string `json:"stage" yaml:"stage"`
	// throttle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#throttle ApiGatewayUsagePlan#throttle}
	Throttle interface{} `json:"throttle" yaml:"throttle"`
}

type ApiGatewayUsagePlanApiStagesThrottle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#path ApiGatewayUsagePlan#path}.
	Path *string `json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#burst_limit ApiGatewayUsagePlan#burst_limit}.
	BurstLimit *float64 `json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#rate_limit ApiGatewayUsagePlan#rate_limit}.
	RateLimit *float64 `json:"rateLimit" yaml:"rateLimit"`
}

// API Gateway.
type ApiGatewayUsagePlanConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#name ApiGatewayUsagePlan#name}.
	Name *string `json:"name" yaml:"name"`
	// api_stages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#api_stages ApiGatewayUsagePlan#api_stages}
	ApiStages interface{} `json:"apiStages" yaml:"apiStages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#description ApiGatewayUsagePlan#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#product_code ApiGatewayUsagePlan#product_code}.
	ProductCode *string `json:"productCode" yaml:"productCode"`
	// quota_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#quota_settings ApiGatewayUsagePlan#quota_settings}
	QuotaSettings *ApiGatewayUsagePlanQuotaSettings `json:"quotaSettings" yaml:"quotaSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#tags ApiGatewayUsagePlan#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#tags_all ApiGatewayUsagePlan#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// throttle_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#throttle_settings ApiGatewayUsagePlan#throttle_settings}
	ThrottleSettings *ApiGatewayUsagePlanThrottleSettings `json:"throttleSettings" yaml:"throttleSettings"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key aws_api_gateway_usage_plan_key}.
type ApiGatewayUsagePlanKey interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UsagePlanId() *string
	SetUsagePlanId(val *string)
	UsagePlanIdInput() *string
	Value() *string
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

// The jsii proxy struct for ApiGatewayUsagePlanKey
type jsiiProxy_ApiGatewayUsagePlanKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) UsagePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key aws_api_gateway_usage_plan_key} Resource.
func NewApiGatewayUsagePlanKey(scope constructs.Construct, id *string, config *ApiGatewayUsagePlanKeyConfig) ApiGatewayUsagePlanKey {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayUsagePlanKey{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key aws_api_gateway_usage_plan_key} Resource.
func NewApiGatewayUsagePlanKey_Override(a ApiGatewayUsagePlanKey, scope constructs.Construct, id *string, config *ApiGatewayUsagePlanKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanKey",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetKeyType(val *string) {
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanKey) SetUsagePlanId(val *string) {
	_jsii_.Set(
		j,
		"usagePlanId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayUsagePlanKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayUsagePlanKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlanKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayUsagePlanKeyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key#key_id ApiGatewayUsagePlanKey#key_id}.
	KeyId *string `json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key#key_type ApiGatewayUsagePlanKey#key_type}.
	KeyType *string `json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan_key#usage_plan_id ApiGatewayUsagePlanKey#usage_plan_id}.
	UsagePlanId *string `json:"usagePlanId" yaml:"usagePlanId"`
}

type ApiGatewayUsagePlanQuotaSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#limit ApiGatewayUsagePlan#limit}.
	Limit *float64 `json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#period ApiGatewayUsagePlan#period}.
	Period *string `json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#offset ApiGatewayUsagePlan#offset}.
	Offset *float64 `json:"offset" yaml:"offset"`
}

type ApiGatewayUsagePlanQuotaSettingsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ApiGatewayUsagePlanQuotaSettings
	SetInternalValue(val *ApiGatewayUsagePlanQuotaSettings)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	Offset() *float64
	SetOffset(val *float64)
	OffsetInput() *float64
	Period() *string
	SetPeriod(val *string)
	PeriodInput() *string
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
	ResetOffset()
}

// The jsii proxy struct for ApiGatewayUsagePlanQuotaSettingsOutputReference
type jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) InternalValue() *ApiGatewayUsagePlanQuotaSettings {
	var returns *ApiGatewayUsagePlanQuotaSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) Offset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) OffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApiGatewayUsagePlanQuotaSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayUsagePlanQuotaSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanQuotaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayUsagePlanQuotaSettingsOutputReference_Override(a ApiGatewayUsagePlanQuotaSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanQuotaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetInternalValue(val *ApiGatewayUsagePlanQuotaSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetLimit(val *float64) {
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetOffset(val *float64) {
	_jsii_.Set(
		j,
		"offset",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetPeriod(val *string) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayUsagePlanQuotaSettingsOutputReference) ResetOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetOffset",
		nil, // no parameters
	)
}

type ApiGatewayUsagePlanThrottleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#burst_limit ApiGatewayUsagePlan#burst_limit}.
	BurstLimit *float64 `json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#rate_limit ApiGatewayUsagePlan#rate_limit}.
	RateLimit *float64 `json:"rateLimit" yaml:"rateLimit"`
}

type ApiGatewayUsagePlanThrottleSettingsOutputReference interface {
	cdktf.ComplexObject
	BurstLimit() *float64
	SetBurstLimit(val *float64)
	BurstLimitInput() *float64
	InternalValue() *ApiGatewayUsagePlanThrottleSettings
	SetInternalValue(val *ApiGatewayUsagePlanThrottleSettings)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RateLimit() *float64
	SetRateLimit(val *float64)
	RateLimitInput() *float64
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
	ResetBurstLimit()
	ResetRateLimit()
}

// The jsii proxy struct for ApiGatewayUsagePlanThrottleSettingsOutputReference
type jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) BurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) BurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) InternalValue() *ApiGatewayUsagePlanThrottleSettings {
	var returns *ApiGatewayUsagePlanThrottleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) RateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) RateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApiGatewayUsagePlanThrottleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ApiGatewayUsagePlanThrottleSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanThrottleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApiGatewayUsagePlanThrottleSettingsOutputReference_Override(a ApiGatewayUsagePlanThrottleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayUsagePlanThrottleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetBurstLimit(val *float64) {
	_jsii_.Set(
		j,
		"burstLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetInternalValue(val *ApiGatewayUsagePlanThrottleSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetRateLimit(val *float64) {
	_jsii_.Set(
		j,
		"rateLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) ResetBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayUsagePlanThrottleSettingsOutputReference) ResetRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimit",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link aws_api_gateway_vpc_link}.
type ApiGatewayVpcLink interface {
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
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetArns() *[]*string
	SetTargetArns(val *[]*string)
	TargetArnsInput() *[]*string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiGatewayVpcLink
type jsiiProxy_ApiGatewayVpcLink struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayVpcLink) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TargetArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayVpcLink) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link aws_api_gateway_vpc_link} Resource.
func NewApiGatewayVpcLink(scope constructs.Construct, id *string, config *ApiGatewayVpcLinkConfig) ApiGatewayVpcLink {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayVpcLink{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayVpcLink",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link aws_api_gateway_vpc_link} Resource.
func NewApiGatewayVpcLink_Override(a ApiGatewayVpcLink, scope constructs.Construct, id *string, config *ApiGatewayVpcLinkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.ApiGatewayVpcLink",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayVpcLink) SetTargetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetArns",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayVpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.ApiGatewayVpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayVpcLink_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.ApiGatewayVpcLink",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayVpcLink) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayVpcLink) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayVpcLink) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayVpcLink) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayVpcLink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ApiGatewayVpcLink) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type ApiGatewayVpcLinkConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link#name ApiGatewayVpcLink#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link#target_arns ApiGatewayVpcLink#target_arns}.
	TargetArns *[]*string `json:"targetArns" yaml:"targetArns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link#description ApiGatewayVpcLink#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link#tags ApiGatewayVpcLink#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_vpc_link#tags_all ApiGatewayVpcLink#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_api_key aws_api_gateway_api_key}.
type DataAwsApiGatewayApiKey interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Enabled() cdktf.IResolvable
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Value() *string
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
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewayApiKey
type jsiiProxy_DataAwsApiGatewayApiKey struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_api_key aws_api_gateway_api_key} Data Source.
func NewDataAwsApiGatewayApiKey(scope constructs.Construct, id *string, config *DataAwsApiGatewayApiKeyConfig) DataAwsApiGatewayApiKey {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayApiKey{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayApiKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_api_key aws_api_gateway_api_key} Data Source.
func NewDataAwsApiGatewayApiKey_Override(d DataAwsApiGatewayApiKey, scope constructs.Construct, id *string, config *DataAwsApiGatewayApiKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayApiKey",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayApiKey) SetTags(val *map[string]*string) {
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
func DataAwsApiGatewayApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayApiKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayApiKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayApiKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayApiKey) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayApiKey) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayApiKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayApiKeyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_api_key#id DataAwsApiGatewayApiKey#id}.
	Id *string `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_api_key#tags DataAwsApiGatewayApiKey#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_domain_name aws_api_gateway_domain_name}.
type DataAwsApiGatewayDomainName interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	CertificateName() *string
	CertificateUploadDate() *string
	CloudfrontDomainName() *string
	CloudfrontZoneId() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegionalCertificateArn() *string
	RegionalCertificateName() *string
	RegionalDomainName() *string
	RegionalZoneId() *string
	SecurityPolicy() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	EndpointConfiguration(index *string) DataAwsApiGatewayDomainNameEndpointConfiguration
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
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewayDomainName
type jsiiProxy_DataAwsApiGatewayDomainName struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CertificateUploadDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUploadDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CloudfrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) CloudfrontZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) RegionalCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) RegionalCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) RegionalZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_domain_name aws_api_gateway_domain_name} Data Source.
func NewDataAwsApiGatewayDomainName(scope constructs.Construct, id *string, config *DataAwsApiGatewayDomainNameConfig) DataAwsApiGatewayDomainName {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayDomainName{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_domain_name aws_api_gateway_domain_name} Data Source.
func NewDataAwsApiGatewayDomainName_Override(d DataAwsApiGatewayDomainName, scope constructs.Construct, id *string, config *DataAwsApiGatewayDomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainName",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainName) SetTags(val *map[string]*string) {
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
func DataAwsApiGatewayDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayDomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsApiGatewayDomainName) EndpointConfiguration(index *string) DataAwsApiGatewayDomainNameEndpointConfiguration {
	var returns DataAwsApiGatewayDomainNameEndpointConfiguration

	_jsii_.Invoke(
		d,
		"endpointConfiguration",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayDomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayDomainName) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayDomainName) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayDomainNameConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_domain_name#domain_name DataAwsApiGatewayDomainName#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_domain_name#tags DataAwsApiGatewayDomainName#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsApiGatewayDomainNameEndpointConfiguration interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Types() *[]*string
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

// The jsii proxy struct for DataAwsApiGatewayDomainNameEndpointConfiguration
type jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsApiGatewayDomainNameEndpointConfiguration(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsApiGatewayDomainNameEndpointConfiguration {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainNameEndpointConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsApiGatewayDomainNameEndpointConfiguration_Override(d DataAwsApiGatewayDomainNameEndpointConfiguration, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayDomainNameEndpointConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfiguration) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export aws_api_gateway_export}.
type DataAwsApiGatewayExport interface {
	cdktf.TerraformDataSource
	Accepts() *string
	SetAccepts(val *string)
	AcceptsInput() *string
	Body() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentDisposition() *string
	ContentType() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExportType() *string
	SetExportType(val *string)
	ExportTypeInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
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
	ResetAccepts()
	ResetOverrideLogicalId()
	ResetParameters()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewayExport
type jsiiProxy_DataAwsApiGatewayExport struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Accepts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accepts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) AcceptsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ExportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ExportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayExport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export aws_api_gateway_export} Data Source.
func NewDataAwsApiGatewayExport(scope constructs.Construct, id *string, config *DataAwsApiGatewayExportConfig) DataAwsApiGatewayExport {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayExport{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayExport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export aws_api_gateway_export} Data Source.
func NewDataAwsApiGatewayExport_Override(d DataAwsApiGatewayExport, scope constructs.Construct, id *string, config *DataAwsApiGatewayExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayExport",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetAccepts(val *string) {
	_jsii_.Set(
		j,
		"accepts",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetExportType(val *string) {
	_jsii_.Set(
		j,
		"exportType",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayExport) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsApiGatewayExport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayExport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayExport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayExport",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayExport) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayExport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsApiGatewayExport) ResetAccepts() {
	_jsii_.InvokeVoid(
		d,
		"resetAccepts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayExport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayExport) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayExport) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayExport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayExportConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export#export_type DataAwsApiGatewayExport#export_type}.
	ExportType *string `json:"exportType" yaml:"exportType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export#rest_api_id DataAwsApiGatewayExport#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export#stage_name DataAwsApiGatewayExport#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export#accepts DataAwsApiGatewayExport#accepts}.
	Accepts *string `json:"accepts" yaml:"accepts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_export#parameters DataAwsApiGatewayExport#parameters}.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_resource aws_api_gateway_resource}.
type DataAwsApiGatewayResource interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	ParentId() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PathPart() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
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

// The jsii proxy struct for DataAwsApiGatewayResource
type jsiiProxy_DataAwsApiGatewayResource struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) PathPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_resource aws_api_gateway_resource} Data Source.
func NewDataAwsApiGatewayResource(scope constructs.Construct, id *string, config *DataAwsApiGatewayResourceConfig) DataAwsApiGatewayResource {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayResource{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_resource aws_api_gateway_resource} Data Source.
func NewDataAwsApiGatewayResource_Override(d DataAwsApiGatewayResource, scope constructs.Construct, id *string, config *DataAwsApiGatewayResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayResource",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayResource) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsApiGatewayResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayResource) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayResourceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_resource#path DataAwsApiGatewayResource#path}.
	Path *string `json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_resource#rest_api_id DataAwsApiGatewayResource#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_rest_api aws_api_gateway_rest_api}.
type DataAwsApiGatewayRestApi interface {
	cdktf.TerraformDataSource
	ApiKeySource() *string
	Arn() *string
	BinaryMediaTypes() *[]*string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	ExecutionArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinimumCompressionSize() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Policy() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootResourceId() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	EndpointConfiguration(index *string) DataAwsApiGatewayRestApiEndpointConfiguration
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
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewayRestApi
type jsiiProxy_DataAwsApiGatewayRestApi struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) ApiKeySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) RootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_rest_api aws_api_gateway_rest_api} Data Source.
func NewDataAwsApiGatewayRestApi(scope constructs.Construct, id *string, config *DataAwsApiGatewayRestApiConfig) DataAwsApiGatewayRestApi {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayRestApi{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_rest_api aws_api_gateway_rest_api} Data Source.
func NewDataAwsApiGatewayRestApi_Override(d DataAwsApiGatewayRestApi, scope constructs.Construct, id *string, config *DataAwsApiGatewayRestApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApi",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApi) SetTags(val *map[string]*string) {
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
func DataAwsApiGatewayRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayRestApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayRestApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsApiGatewayRestApi) EndpointConfiguration(index *string) DataAwsApiGatewayRestApiEndpointConfiguration {
	var returns DataAwsApiGatewayRestApiEndpointConfiguration

	_jsii_.Invoke(
		d,
		"endpointConfiguration",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayRestApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayRestApi) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayRestApi) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayRestApiConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_rest_api#name DataAwsApiGatewayRestApi#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_rest_api#tags DataAwsApiGatewayRestApi#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsApiGatewayRestApiEndpointConfiguration interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Types() *[]*string
	VpcEndpointIds() *[]*string
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

// The jsii proxy struct for DataAwsApiGatewayRestApiEndpointConfiguration
type jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) VpcEndpointIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcEndpointIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsApiGatewayRestApiEndpointConfiguration(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsApiGatewayRestApiEndpointConfiguration {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApiEndpointConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsApiGatewayRestApiEndpointConfiguration_Override(d DataAwsApiGatewayRestApiEndpointConfiguration, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayRestApiEndpointConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayRestApiEndpointConfiguration) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk aws_api_gateway_sdk}.
type DataAwsApiGatewaySdk interface {
	cdktf.TerraformDataSource
	Body() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentDisposition() *string
	ContentType() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RestApiId() *string
	SetRestApiId(val *string)
	RestApiIdInput() *string
	SdkType() *string
	SetSdkType(val *string)
	SdkTypeInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
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
	ResetParameters()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewaySdk
type jsiiProxy_DataAwsApiGatewaySdk struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SdkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SdkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk aws_api_gateway_sdk} Data Source.
func NewDataAwsApiGatewaySdk(scope constructs.Construct, id *string, config *DataAwsApiGatewaySdkConfig) DataAwsApiGatewaySdk {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewaySdk{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewaySdk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk aws_api_gateway_sdk} Data Source.
func NewDataAwsApiGatewaySdk_Override(d DataAwsApiGatewaySdk, scope constructs.Construct, id *string, config *DataAwsApiGatewaySdkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewaySdk",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetSdkType(val *string) {
	_jsii_.Set(
		j,
		"sdkType",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewaySdk) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsApiGatewaySdk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewaySdk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewaySdk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewaySdk",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewaySdk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewaySdk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewaySdk) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewaySdk) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewaySdk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewaySdkConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk#rest_api_id DataAwsApiGatewaySdk#rest_api_id}.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk#sdk_type DataAwsApiGatewaySdk#sdk_type}.
	SdkType *string `json:"sdkType" yaml:"sdkType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk#stage_name DataAwsApiGatewaySdk#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_sdk#parameters DataAwsApiGatewaySdk#parameters}.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_vpc_link aws_api_gateway_vpc_link}.
type DataAwsApiGatewayVpcLink interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	StatusMessage() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetArns() *[]*string
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
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApiGatewayVpcLink
type jsiiProxy_DataAwsApiGatewayVpcLink struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_vpc_link aws_api_gateway_vpc_link} Data Source.
func NewDataAwsApiGatewayVpcLink(scope constructs.Construct, id *string, config *DataAwsApiGatewayVpcLinkConfig) DataAwsApiGatewayVpcLink {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApiGatewayVpcLink{}

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayVpcLink",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_vpc_link aws_api_gateway_vpc_link} Data Source.
func NewDataAwsApiGatewayVpcLink_Override(d DataAwsApiGatewayVpcLink, scope constructs.Construct, id *string, config *DataAwsApiGatewayVpcLinkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigateway.DataAwsApiGatewayVpcLink",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApiGatewayVpcLink) SetTags(val *map[string]*string) {
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
func DataAwsApiGatewayVpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigateway.DataAwsApiGatewayVpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApiGatewayVpcLink_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigateway.DataAwsApiGatewayVpcLink",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayVpcLink) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApiGatewayVpcLink) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) ToString() *string {
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
func (d *jsiiProxy_DataAwsApiGatewayVpcLink) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API Gateway.
type DataAwsApiGatewayVpcLinkConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_vpc_link#name DataAwsApiGatewayVpcLink#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/api_gateway_vpc_link#tags DataAwsApiGatewayVpcLink#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}
