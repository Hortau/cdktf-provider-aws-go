package apigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/apigatewayv2/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api aws_apigatewayv2_api}.
type Apigatewayv2Api interface {
	cdktf.TerraformResource
	ApiEndpoint() *string
	ApiKeySelectionExpression() *string
	SetApiKeySelectionExpression(val *string)
	ApiKeySelectionExpressionInput() *string
	Arn() *string
	Body() *string
	SetBody(val *string)
	BodyInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfiguration() Apigatewayv2ApiCorsConfigurationOutputReference
	CorsConfigurationInput() *Apigatewayv2ApiCorsConfiguration
	Count() *float64
	SetCount(val *float64)
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	CredentialsArnInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	DisableExecuteApiEndpointInput() interface{}
	ExecutionArn() *string
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
	FailOnWarningsInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RouteKey() *string
	SetRouteKey(val *string)
	RouteKeyInput() *string
	RouteSelectionExpression() *string
	SetRouteSelectionExpression(val *string)
	RouteSelectionExpressionInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	PutCorsConfiguration(value *Apigatewayv2ApiCorsConfiguration)
	ResetApiKeySelectionExpression()
	ResetBody()
	ResetCorsConfiguration()
	ResetCredentialsArn()
	ResetDescription()
	ResetDisableExecuteApiEndpoint()
	ResetFailOnWarnings()
	ResetOverrideLogicalId()
	ResetRouteKey()
	ResetRouteSelectionExpression()
	ResetTags()
	ResetTagsAll()
	ResetTarget()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Api
type jsiiProxy_Apigatewayv2Api struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Api) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ApiKeySelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ApiKeySelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) CorsConfiguration() Apigatewayv2ApiCorsConfigurationOutputReference {
	var returns Apigatewayv2ApiCorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) CorsConfigurationInput() *Apigatewayv2ApiCorsConfiguration {
	var returns *Apigatewayv2ApiCorsConfiguration
	_jsii_.Get(
		j,
		"corsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) CredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) DisableExecuteApiEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) FailOnWarningsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) RouteKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) RouteSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) RouteSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Api) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api aws_apigatewayv2_api} Resource.
func NewApigatewayv2Api(scope constructs.Construct, id *string, config *Apigatewayv2ApiConfig) Apigatewayv2Api {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Api{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Api",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api aws_apigatewayv2_api} Resource.
func NewApigatewayv2Api_Override(a Apigatewayv2Api, scope constructs.Construct, id *string, config *Apigatewayv2ApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Api",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetApiKeySelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"apiKeySelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetBody(val *string) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"credentialsArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetFailOnWarnings(val interface{}) {
	_jsii_.Set(
		j,
		"failOnWarnings",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetProtocolType(val *string) {
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetRouteKey(val *string) {
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetRouteSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"routeSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Api) SetVersion(val *string) {
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
func Apigatewayv2Api_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Api",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Api_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Api",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Api) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Api) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Api) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Api) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Api) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Api) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Api) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Api) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Api) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Api) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Api) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Api) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Api) PutCorsConfiguration(value *Apigatewayv2ApiCorsConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCorsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetApiKeySelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeySelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetBody() {
	_jsii_.InvokeVoid(
		a,
		"resetBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetCorsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCorsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetDisableExecuteApiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableExecuteApiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetFailOnWarnings() {
	_jsii_.InvokeVoid(
		a,
		"resetFailOnWarnings",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Api) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetRouteKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetRouteSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Api) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Api) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Api) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Api) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2ApiConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#name Apigatewayv2Api#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#protocol_type Apigatewayv2Api#protocol_type}.
	ProtocolType *string `json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#api_key_selection_expression Apigatewayv2Api#api_key_selection_expression}.
	ApiKeySelectionExpression *string `json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#body Apigatewayv2Api#body}.
	Body *string `json:"body" yaml:"body"`
	// cors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#cors_configuration Apigatewayv2Api#cors_configuration}
	CorsConfiguration *Apigatewayv2ApiCorsConfiguration `json:"corsConfiguration" yaml:"corsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#credentials_arn Apigatewayv2Api#credentials_arn}.
	CredentialsArn *string `json:"credentialsArn" yaml:"credentialsArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#description Apigatewayv2Api#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#disable_execute_api_endpoint Apigatewayv2Api#disable_execute_api_endpoint}.
	DisableExecuteApiEndpoint interface{} `json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#fail_on_warnings Apigatewayv2Api#fail_on_warnings}.
	FailOnWarnings interface{} `json:"failOnWarnings" yaml:"failOnWarnings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#route_key Apigatewayv2Api#route_key}.
	RouteKey *string `json:"routeKey" yaml:"routeKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#route_selection_expression Apigatewayv2Api#route_selection_expression}.
	RouteSelectionExpression *string `json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#tags Apigatewayv2Api#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#tags_all Apigatewayv2Api#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#target Apigatewayv2Api#target}.
	Target *string `json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#version Apigatewayv2Api#version}.
	Version *string `json:"version" yaml:"version"`
}

type Apigatewayv2ApiCorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#allow_credentials Apigatewayv2Api#allow_credentials}.
	AllowCredentials interface{} `json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#allow_headers Apigatewayv2Api#allow_headers}.
	AllowHeaders *[]*string `json:"allowHeaders" yaml:"allowHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#allow_methods Apigatewayv2Api#allow_methods}.
	AllowMethods *[]*string `json:"allowMethods" yaml:"allowMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#allow_origins Apigatewayv2Api#allow_origins}.
	AllowOrigins *[]*string `json:"allowOrigins" yaml:"allowOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#expose_headers Apigatewayv2Api#expose_headers}.
	ExposeHeaders *[]*string `json:"exposeHeaders" yaml:"exposeHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api#max_age Apigatewayv2Api#max_age}.
	MaxAge *float64 `json:"maxAge" yaml:"maxAge"`
}

type Apigatewayv2ApiCorsConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowHeaders() *[]*string
	SetAllowHeaders(val *[]*string)
	AllowHeadersInput() *[]*string
	AllowMethods() *[]*string
	SetAllowMethods(val *[]*string)
	AllowMethodsInput() *[]*string
	AllowOrigins() *[]*string
	SetAllowOrigins(val *[]*string)
	AllowOriginsInput() *[]*string
	ExposeHeaders() *[]*string
	SetExposeHeaders(val *[]*string)
	ExposeHeadersInput() *[]*string
	InternalValue() *Apigatewayv2ApiCorsConfiguration
	SetInternalValue(val *Apigatewayv2ApiCorsConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
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
	ResetAllowCredentials()
	ResetAllowHeaders()
	ResetAllowMethods()
	ResetAllowOrigins()
	ResetExposeHeaders()
	ResetMaxAge()
}

// The jsii proxy struct for Apigatewayv2ApiCorsConfigurationOutputReference
type jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) AllowOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ExposeHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ExposeHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) InternalValue() *Apigatewayv2ApiCorsConfiguration {
	var returns *Apigatewayv2ApiCorsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApigatewayv2ApiCorsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2ApiCorsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiCorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2ApiCorsConfigurationOutputReference_Override(a Apigatewayv2ApiCorsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiCorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetAllowCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetAllowHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowHeaders",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetAllowMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowMethods",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetAllowOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowOrigins",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetExposeHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"exposeHeaders",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetInternalValue(val *Apigatewayv2ApiCorsConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetMaxAge(val *float64) {
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetAllowHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetAllowMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetAllowOrigins() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowOrigins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetExposeHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetExposeHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiCorsConfigurationOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAge",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping aws_apigatewayv2_api_mapping}.
type Apigatewayv2ApiMapping interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	ApiMappingKey() *string
	SetApiMappingKey(val *string)
	ApiMappingKeyInput() *string
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
	Stage() *string
	SetStage(val *string)
	StageInput() *string
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
	ResetApiMappingKey()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2ApiMapping
type jsiiProxy_Apigatewayv2ApiMapping struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) ApiMappingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) ApiMappingKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping aws_apigatewayv2_api_mapping} Resource.
func NewApigatewayv2ApiMapping(scope constructs.Construct, id *string, config *Apigatewayv2ApiMappingConfig) Apigatewayv2ApiMapping {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2ApiMapping{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping aws_apigatewayv2_api_mapping} Resource.
func NewApigatewayv2ApiMapping_Override(a Apigatewayv2ApiMapping, scope constructs.Construct, id *string, config *Apigatewayv2ApiMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiMapping",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetApiMappingKey(val *string) {
	_jsii_.Set(
		j,
		"apiMappingKey",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2ApiMapping) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Apigatewayv2ApiMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2ApiMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2ApiMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2ApiMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2ApiMapping) ResetApiMappingKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiMappingKey",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2ApiMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2ApiMapping) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2ApiMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2ApiMappingConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping#api_id Apigatewayv2ApiMapping#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping#domain_name Apigatewayv2ApiMapping#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping#stage Apigatewayv2ApiMapping#stage}.
	Stage *string `json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api_mapping#api_mapping_key Apigatewayv2ApiMapping#api_mapping_key}.
	ApiMappingKey *string `json:"apiMappingKey" yaml:"apiMappingKey"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer aws_apigatewayv2_authorizer}.
type Apigatewayv2Authorizer interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	AuthorizerCredentialsArn() *string
	SetAuthorizerCredentialsArn(val *string)
	AuthorizerCredentialsArnInput() *string
	AuthorizerPayloadFormatVersion() *string
	SetAuthorizerPayloadFormatVersion(val *string)
	AuthorizerPayloadFormatVersionInput() *string
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerResultTtlInSecondsInput() *float64
	AuthorizerType() *string
	SetAuthorizerType(val *string)
	AuthorizerTypeInput() *string
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	AuthorizerUriInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnableSimpleResponses() interface{}
	SetEnableSimpleResponses(val interface{})
	EnableSimpleResponsesInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentitySources() *[]*string
	SetIdentitySources(val *[]*string)
	IdentitySourcesInput() *[]*string
	JwtConfiguration() Apigatewayv2AuthorizerJwtConfigurationOutputReference
	JwtConfigurationInput() *Apigatewayv2AuthorizerJwtConfiguration
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
	PutJwtConfiguration(value *Apigatewayv2AuthorizerJwtConfiguration)
	ResetAuthorizerCredentialsArn()
	ResetAuthorizerPayloadFormatVersion()
	ResetAuthorizerResultTtlInSeconds()
	ResetAuthorizerUri()
	ResetEnableSimpleResponses()
	ResetIdentitySources()
	ResetJwtConfiguration()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Authorizer
type jsiiProxy_Apigatewayv2Authorizer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerCredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerPayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerPayloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerPayloadFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerPayloadFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) EnableSimpleResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSimpleResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) EnableSimpleResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSimpleResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentitySources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) IdentitySourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) JwtConfiguration() Apigatewayv2AuthorizerJwtConfigurationOutputReference {
	var returns Apigatewayv2AuthorizerJwtConfigurationOutputReference
	_jsii_.Get(
		j,
		"jwtConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) JwtConfigurationInput() *Apigatewayv2AuthorizerJwtConfiguration {
	var returns *Apigatewayv2AuthorizerJwtConfiguration
	_jsii_.Get(
		j,
		"jwtConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Authorizer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer aws_apigatewayv2_authorizer} Resource.
func NewApigatewayv2Authorizer(scope constructs.Construct, id *string, config *Apigatewayv2AuthorizerConfig) Apigatewayv2Authorizer {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Authorizer{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Authorizer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer aws_apigatewayv2_authorizer} Resource.
func NewApigatewayv2Authorizer_Override(a Apigatewayv2Authorizer, scope constructs.Construct, id *string, config *Apigatewayv2AuthorizerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Authorizer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetAuthorizerCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"authorizerCredentialsArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetAuthorizerPayloadFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"authorizerPayloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetAuthorizerType(val *string) {
	_jsii_.Set(
		j,
		"authorizerType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetEnableSimpleResponses(val interface{}) {
	_jsii_.Set(
		j,
		"enableSimpleResponses",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetIdentitySources(val *[]*string) {
	_jsii_.Set(
		j,
		"identitySources",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Authorizer) SetProvider(val cdktf.TerraformProvider) {
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
func Apigatewayv2Authorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Authorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Authorizer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Authorizer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Authorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Authorizer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) PutJwtConfiguration(value *Apigatewayv2AuthorizerJwtConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putJwtConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerPayloadFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerPayloadFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetAuthorizerUri() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetEnableSimpleResponses() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSimpleResponses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetIdentitySources() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentitySources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) ResetJwtConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtConfiguration",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Authorizer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Authorizer) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Authorizer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2AuthorizerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#api_id Apigatewayv2Authorizer#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#authorizer_type Apigatewayv2Authorizer#authorizer_type}.
	AuthorizerType *string `json:"authorizerType" yaml:"authorizerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#name Apigatewayv2Authorizer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#authorizer_credentials_arn Apigatewayv2Authorizer#authorizer_credentials_arn}.
	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn" yaml:"authorizerCredentialsArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#authorizer_payload_format_version Apigatewayv2Authorizer#authorizer_payload_format_version}.
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion" yaml:"authorizerPayloadFormatVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#authorizer_result_ttl_in_seconds Apigatewayv2Authorizer#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#authorizer_uri Apigatewayv2Authorizer#authorizer_uri}.
	AuthorizerUri *string `json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#enable_simple_responses Apigatewayv2Authorizer#enable_simple_responses}.
	EnableSimpleResponses interface{} `json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#identity_sources Apigatewayv2Authorizer#identity_sources}.
	IdentitySources *[]*string `json:"identitySources" yaml:"identitySources"`
	// jwt_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#jwt_configuration Apigatewayv2Authorizer#jwt_configuration}
	JwtConfiguration *Apigatewayv2AuthorizerJwtConfiguration `json:"jwtConfiguration" yaml:"jwtConfiguration"`
}

type Apigatewayv2AuthorizerJwtConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#audience Apigatewayv2Authorizer#audience}.
	Audience *[]*string `json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#issuer Apigatewayv2Authorizer#issuer}.
	Issuer *string `json:"issuer" yaml:"issuer"`
}

type Apigatewayv2AuthorizerJwtConfigurationOutputReference interface {
	cdktf.ComplexObject
	Audience() *[]*string
	SetAudience(val *[]*string)
	AudienceInput() *[]*string
	InternalValue() *Apigatewayv2AuthorizerJwtConfiguration
	SetInternalValue(val *Apigatewayv2AuthorizerJwtConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
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
	ResetAudience()
	ResetIssuer()
}

// The jsii proxy struct for Apigatewayv2AuthorizerJwtConfigurationOutputReference
type jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) Audience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) AudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) InternalValue() *Apigatewayv2AuthorizerJwtConfiguration {
	var returns *Apigatewayv2AuthorizerJwtConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApigatewayv2AuthorizerJwtConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2AuthorizerJwtConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2AuthorizerJwtConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2AuthorizerJwtConfigurationOutputReference_Override(a Apigatewayv2AuthorizerJwtConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2AuthorizerJwtConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetAudience(val *[]*string) {
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetInternalValue(val *Apigatewayv2AuthorizerJwtConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2AuthorizerJwtConfigurationOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment aws_apigatewayv2_deployment}.
type Apigatewayv2Deployment interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	AutoDeployed() cdktf.IResolvable
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTriggers()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Deployment
type jsiiProxy_Apigatewayv2Deployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Deployment) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) AutoDeployed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoDeployed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Deployment) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment aws_apigatewayv2_deployment} Resource.
func NewApigatewayv2Deployment(scope constructs.Construct, id *string, config *Apigatewayv2DeploymentConfig) Apigatewayv2Deployment {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Deployment{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Deployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment aws_apigatewayv2_deployment} Resource.
func NewApigatewayv2Deployment_Override(a Apigatewayv2Deployment, scope constructs.Construct, id *string, config *Apigatewayv2DeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Deployment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Deployment) SetTriggers(val *map[string]*string) {
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
func Apigatewayv2Deployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Deployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Deployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Deployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Deployment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Deployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Deployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Deployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Deployment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Deployment) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Deployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Deployment) ResetTriggers() {
	_jsii_.InvokeVoid(
		a,
		"resetTriggers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Deployment) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Deployment) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Deployment) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Deployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2DeploymentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment#api_id Apigatewayv2Deployment#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment#description Apigatewayv2Deployment#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment#triggers Apigatewayv2Deployment#triggers}.
	Triggers *map[string]*string `json:"triggers" yaml:"triggers"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name aws_apigatewayv2_domain_name}.
type Apigatewayv2DomainName interface {
	cdktf.TerraformResource
	ApiMappingSelectionExpression() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameConfiguration() Apigatewayv2DomainNameDomainNameConfigurationOutputReference
	DomainNameConfigurationInput() *Apigatewayv2DomainNameDomainNameConfiguration
	DomainNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MutualTlsAuthentication() Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference
	MutualTlsAuthenticationInput() *Apigatewayv2DomainNameMutualTlsAuthentication
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
	Timeouts() Apigatewayv2DomainNameTimeoutsOutputReference
	TimeoutsInput() *Apigatewayv2DomainNameTimeouts
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
	PutDomainNameConfiguration(value *Apigatewayv2DomainNameDomainNameConfiguration)
	PutMutualTlsAuthentication(value *Apigatewayv2DomainNameMutualTlsAuthentication)
	PutTimeouts(value *Apigatewayv2DomainNameTimeouts)
	ResetMutualTlsAuthentication()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2DomainName
type jsiiProxy_Apigatewayv2DomainName struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2DomainName) ApiMappingSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) DomainNameConfiguration() Apigatewayv2DomainNameDomainNameConfigurationOutputReference {
	var returns Apigatewayv2DomainNameDomainNameConfigurationOutputReference
	_jsii_.Get(
		j,
		"domainNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) DomainNameConfigurationInput() *Apigatewayv2DomainNameDomainNameConfiguration {
	var returns *Apigatewayv2DomainNameDomainNameConfiguration
	_jsii_.Get(
		j,
		"domainNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) MutualTlsAuthentication() Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference {
	var returns Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) MutualTlsAuthenticationInput() *Apigatewayv2DomainNameMutualTlsAuthentication {
	var returns *Apigatewayv2DomainNameMutualTlsAuthentication
	_jsii_.Get(
		j,
		"mutualTlsAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) Timeouts() Apigatewayv2DomainNameTimeoutsOutputReference {
	var returns Apigatewayv2DomainNameTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainName) TimeoutsInput() *Apigatewayv2DomainNameTimeouts {
	var returns *Apigatewayv2DomainNameTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name aws_apigatewayv2_domain_name} Resource.
func NewApigatewayv2DomainName(scope constructs.Construct, id *string, config *Apigatewayv2DomainNameConfig) Apigatewayv2DomainName {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2DomainName{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name aws_apigatewayv2_domain_name} Resource.
func NewApigatewayv2DomainName_Override(a Apigatewayv2DomainName, scope constructs.Construct, id *string, config *Apigatewayv2DomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainName",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainName) SetTagsAll(val *map[string]*string) {
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
func Apigatewayv2DomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2DomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2DomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainName) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) PutDomainNameConfiguration(value *Apigatewayv2DomainNameDomainNameConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putDomainNameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) PutMutualTlsAuthentication(value *Apigatewayv2DomainNameMutualTlsAuthentication) {
	_jsii_.InvokeVoid(
		a,
		"putMutualTlsAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) PutTimeouts(value *Apigatewayv2DomainNameTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) ResetMutualTlsAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualTlsAuthentication",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2DomainName) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainName) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainName) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2DomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2DomainNameConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#domain_name Apigatewayv2DomainName#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// domain_name_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#domain_name_configuration Apigatewayv2DomainName#domain_name_configuration}
	DomainNameConfiguration *Apigatewayv2DomainNameDomainNameConfiguration `json:"domainNameConfiguration" yaml:"domainNameConfiguration"`
	// mutual_tls_authentication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#mutual_tls_authentication Apigatewayv2DomainName#mutual_tls_authentication}
	MutualTlsAuthentication *Apigatewayv2DomainNameMutualTlsAuthentication `json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#tags Apigatewayv2DomainName#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#tags_all Apigatewayv2DomainName#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#timeouts Apigatewayv2DomainName#timeouts}
	Timeouts *Apigatewayv2DomainNameTimeouts `json:"timeouts" yaml:"timeouts"`
}

type Apigatewayv2DomainNameDomainNameConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#certificate_arn Apigatewayv2DomainName#certificate_arn}.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#endpoint_type Apigatewayv2DomainName#endpoint_type}.
	EndpointType *string `json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#security_policy Apigatewayv2DomainName#security_policy}.
	SecurityPolicy *string `json:"securityPolicy" yaml:"securityPolicy"`
}

type Apigatewayv2DomainNameDomainNameConfigurationOutputReference interface {
	cdktf.ComplexObject
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	HostedZoneId() *string
	InternalValue() *Apigatewayv2DomainNameDomainNameConfiguration
	SetInternalValue(val *Apigatewayv2DomainNameDomainNameConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	TargetDomainName() *string
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

// The jsii proxy struct for Apigatewayv2DomainNameDomainNameConfigurationOutputReference
type jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) InternalValue() *Apigatewayv2DomainNameDomainNameConfiguration {
	var returns *Apigatewayv2DomainNameDomainNameConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) TargetDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApigatewayv2DomainNameDomainNameConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2DomainNameDomainNameConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameDomainNameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2DomainNameDomainNameConfigurationOutputReference_Override(a Apigatewayv2DomainNameDomainNameConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameDomainNameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetInternalValue(val *Apigatewayv2DomainNameDomainNameConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetSecurityPolicy(val *string) {
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameDomainNameConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type Apigatewayv2DomainNameMutualTlsAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#truststore_uri Apigatewayv2DomainName#truststore_uri}.
	TruststoreUri *string `json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#truststore_version Apigatewayv2DomainName#truststore_version}.
	TruststoreVersion *string `json:"truststoreVersion" yaml:"truststoreVersion"`
}

type Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *Apigatewayv2DomainNameMutualTlsAuthentication
	SetInternalValue(val *Apigatewayv2DomainNameMutualTlsAuthentication)
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

// The jsii proxy struct for Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference
type jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) InternalValue() *Apigatewayv2DomainNameMutualTlsAuthentication {
	var returns *Apigatewayv2DomainNameMutualTlsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TruststoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TruststoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TruststoreVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) TruststoreVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreVersionInput",
		&returns,
	)
	return returns
}

func NewApigatewayv2DomainNameMutualTlsAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2DomainNameMutualTlsAuthenticationOutputReference_Override(a Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetInternalValue(val *Apigatewayv2DomainNameMutualTlsAuthentication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetTruststoreUri(val *string) {
	_jsii_.Set(
		j,
		"truststoreUri",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) SetTruststoreVersion(val *string) {
	_jsii_.Set(
		j,
		"truststoreVersion",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2DomainNameMutualTlsAuthenticationOutputReference) ResetTruststoreVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetTruststoreVersion",
		nil, // no parameters
	)
}

type Apigatewayv2DomainNameTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#create Apigatewayv2DomainName#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_domain_name#update Apigatewayv2DomainName#update}.
	Update *string `json:"update" yaml:"update"`
}

type Apigatewayv2DomainNameTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	InternalValue() *Apigatewayv2DomainNameTimeouts
	SetInternalValue(val *Apigatewayv2DomainNameTimeouts)
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

// The jsii proxy struct for Apigatewayv2DomainNameTimeoutsOutputReference
type jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) InternalValue() *Apigatewayv2DomainNameTimeouts {
	var returns *Apigatewayv2DomainNameTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewApigatewayv2DomainNameTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2DomainNameTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2DomainNameTimeoutsOutputReference_Override(a Apigatewayv2DomainNameTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2DomainNameTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetInternalValue(val *Apigatewayv2DomainNameTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2DomainNameTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration aws_apigatewayv2_integration}.
type Apigatewayv2Integration interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	ContentHandlingStrategyInput() *string
	Count() *float64
	SetCount(val *float64)
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	CredentialsArnInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IntegrationMethod() *string
	SetIntegrationMethod(val *string)
	IntegrationMethodInput() *string
	IntegrationResponseSelectionExpression() *string
	IntegrationSubtype() *string
	SetIntegrationSubtype(val *string)
	IntegrationSubtypeInput() *string
	IntegrationType() *string
	SetIntegrationType(val *string)
	IntegrationTypeInput() *string
	IntegrationUri() *string
	SetIntegrationUri(val *string)
	IntegrationUriInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	PassthroughBehaviorInput() *string
	PayloadFormatVersion() *string
	SetPayloadFormatVersion(val *string)
	PayloadFormatVersionInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestParameters() *map[string]*string
	SetRequestParameters(val *map[string]*string)
	RequestParametersInput() *map[string]*string
	RequestTemplates() *map[string]*string
	SetRequestTemplates(val *map[string]*string)
	RequestTemplatesInput() *map[string]*string
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	ResponseParametersInput() interface{}
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
	TemplateSelectionExpressionInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TimeoutMilliseconds() *float64
	SetTimeoutMilliseconds(val *float64)
	TimeoutMillisecondsInput() *float64
	TlsConfig() Apigatewayv2IntegrationTlsConfigOutputReference
	TlsConfigInput() *Apigatewayv2IntegrationTlsConfig
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
	PutTlsConfig(value *Apigatewayv2IntegrationTlsConfig)
	ResetConnectionId()
	ResetConnectionType()
	ResetContentHandlingStrategy()
	ResetCredentialsArn()
	ResetDescription()
	ResetIntegrationMethod()
	ResetIntegrationSubtype()
	ResetIntegrationUri()
	ResetOverrideLogicalId()
	ResetPassthroughBehavior()
	ResetPayloadFormatVersion()
	ResetRequestParameters()
	ResetRequestTemplates()
	ResetResponseParameters()
	ResetTemplateSelectionExpression()
	ResetTimeoutMilliseconds()
	ResetTlsConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Integration
type jsiiProxy_Apigatewayv2Integration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Integration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ContentHandlingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) CredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationResponseSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationSubtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationSubtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationSubtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationSubtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) IntegrationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PassthroughBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) PayloadFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) RequestTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) ResponseParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TemplateSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TlsConfig() Apigatewayv2IntegrationTlsConfigOutputReference {
	var returns Apigatewayv2IntegrationTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Integration) TlsConfigInput() *Apigatewayv2IntegrationTlsConfig {
	var returns *Apigatewayv2IntegrationTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration aws_apigatewayv2_integration} Resource.
func NewApigatewayv2Integration(scope constructs.Construct, id *string, config *Apigatewayv2IntegrationConfig) Apigatewayv2Integration {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Integration{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Integration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration aws_apigatewayv2_integration} Resource.
func NewApigatewayv2Integration_Override(a Apigatewayv2Integration, scope constructs.Construct, id *string, config *Apigatewayv2IntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Integration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetConnectionId(val *string) {
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"credentialsArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetIntegrationMethod(val *string) {
	_jsii_.Set(
		j,
		"integrationMethod",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetIntegrationSubtype(val *string) {
	_jsii_.Set(
		j,
		"integrationSubtype",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetIntegrationType(val *string) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetIntegrationUri(val *string) {
	_jsii_.Set(
		j,
		"integrationUri",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetPassthroughBehavior(val *string) {
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetPayloadFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetRequestParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetRequestTemplates(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Integration) SetTimeoutMilliseconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutMilliseconds",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Apigatewayv2Integration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Integration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Integration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Integration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Integration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Integration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Integration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Integration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Integration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) PutTlsConfig(value *Apigatewayv2IntegrationTlsConfig) {
	_jsii_.InvokeVoid(
		a,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetConnectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetContentHandlingStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandlingStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationSubtype() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationSubtype",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetIntegrationUri() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationUri",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Integration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetPassthroughBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPassthroughBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetPayloadFormatVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadFormatVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetRequestTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetResponseParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTemplateSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Integration) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Integration) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Integration) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Integration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2IntegrationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#api_id Apigatewayv2Integration#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#integration_type Apigatewayv2Integration#integration_type}.
	IntegrationType *string `json:"integrationType" yaml:"integrationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#connection_id Apigatewayv2Integration#connection_id}.
	ConnectionId *string `json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#connection_type Apigatewayv2Integration#connection_type}.
	ConnectionType *string `json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#content_handling_strategy Apigatewayv2Integration#content_handling_strategy}.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#credentials_arn Apigatewayv2Integration#credentials_arn}.
	CredentialsArn *string `json:"credentialsArn" yaml:"credentialsArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#description Apigatewayv2Integration#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#integration_method Apigatewayv2Integration#integration_method}.
	IntegrationMethod *string `json:"integrationMethod" yaml:"integrationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#integration_subtype Apigatewayv2Integration#integration_subtype}.
	IntegrationSubtype *string `json:"integrationSubtype" yaml:"integrationSubtype"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#integration_uri Apigatewayv2Integration#integration_uri}.
	IntegrationUri *string `json:"integrationUri" yaml:"integrationUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#passthrough_behavior Apigatewayv2Integration#passthrough_behavior}.
	PassthroughBehavior *string `json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#payload_format_version Apigatewayv2Integration#payload_format_version}.
	PayloadFormatVersion *string `json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#request_parameters Apigatewayv2Integration#request_parameters}.
	RequestParameters *map[string]*string `json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#request_templates Apigatewayv2Integration#request_templates}.
	RequestTemplates *map[string]*string `json:"requestTemplates" yaml:"requestTemplates"`
	// response_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#response_parameters Apigatewayv2Integration#response_parameters}
	ResponseParameters interface{} `json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#template_selection_expression Apigatewayv2Integration#template_selection_expression}.
	TemplateSelectionExpression *string `json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#timeout_milliseconds Apigatewayv2Integration#timeout_milliseconds}.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#tls_config Apigatewayv2Integration#tls_config}
	TlsConfig *Apigatewayv2IntegrationTlsConfig `json:"tlsConfig" yaml:"tlsConfig"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response aws_apigatewayv2_integration_response}.
type Apigatewayv2IntegrationResponse interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	ContentHandlingStrategyInput() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IntegrationId() *string
	SetIntegrationId(val *string)
	IntegrationIdInput() *string
	IntegrationResponseKey() *string
	SetIntegrationResponseKey(val *string)
	IntegrationResponseKeyInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResponseTemplates() *map[string]*string
	SetResponseTemplates(val *map[string]*string)
	ResponseTemplatesInput() *map[string]*string
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
	TemplateSelectionExpressionInput() *string
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
	ResetContentHandlingStrategy()
	ResetOverrideLogicalId()
	ResetResponseTemplates()
	ResetTemplateSelectionExpression()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2IntegrationResponse
type jsiiProxy_Apigatewayv2IntegrationResponse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ContentHandlingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) IntegrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) IntegrationResponseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) IntegrationResponseKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ResponseTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) ResponseTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) TemplateSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response aws_apigatewayv2_integration_response} Resource.
func NewApigatewayv2IntegrationResponse(scope constructs.Construct, id *string, config *Apigatewayv2IntegrationResponseConfig) Apigatewayv2IntegrationResponse {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2IntegrationResponse{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationResponse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response aws_apigatewayv2_integration_response} Resource.
func NewApigatewayv2IntegrationResponse_Override(a Apigatewayv2IntegrationResponse, scope constructs.Construct, id *string, config *Apigatewayv2IntegrationResponseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationResponse",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetIntegrationId(val *string) {
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetIntegrationResponseKey(val *string) {
	_jsii_.Set(
		j,
		"integrationResponseKey",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetResponseTemplates(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponse) SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Apigatewayv2IntegrationResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2IntegrationResponse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationResponse",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ResetContentHandlingStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandlingStrategy",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ResetResponseTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ResetTemplateSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponse) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationResponse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2IntegrationResponseConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#api_id Apigatewayv2IntegrationResponse#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#integration_id Apigatewayv2IntegrationResponse#integration_id}.
	IntegrationId *string `json:"integrationId" yaml:"integrationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#integration_response_key Apigatewayv2IntegrationResponse#integration_response_key}.
	IntegrationResponseKey *string `json:"integrationResponseKey" yaml:"integrationResponseKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#content_handling_strategy Apigatewayv2IntegrationResponse#content_handling_strategy}.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#response_templates Apigatewayv2IntegrationResponse#response_templates}.
	ResponseTemplates *map[string]*string `json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration_response#template_selection_expression Apigatewayv2IntegrationResponse#template_selection_expression}.
	TemplateSelectionExpression *string `json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

type Apigatewayv2IntegrationResponseParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#mappings Apigatewayv2Integration#mappings}.
	Mappings *map[string]*string `json:"mappings" yaml:"mappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#status_code Apigatewayv2Integration#status_code}.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
}

type Apigatewayv2IntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_integration#server_name_to_verify Apigatewayv2Integration#server_name_to_verify}.
	ServerNameToVerify *string `json:"serverNameToVerify" yaml:"serverNameToVerify"`
}

type Apigatewayv2IntegrationTlsConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *Apigatewayv2IntegrationTlsConfig
	SetInternalValue(val *Apigatewayv2IntegrationTlsConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ServerNameToVerify() *string
	SetServerNameToVerify(val *string)
	ServerNameToVerifyInput() *string
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
	ResetServerNameToVerify()
}

// The jsii proxy struct for Apigatewayv2IntegrationTlsConfigOutputReference
type jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) InternalValue() *Apigatewayv2IntegrationTlsConfig {
	var returns *Apigatewayv2IntegrationTlsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) ServerNameToVerify() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameToVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) ServerNameToVerifyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameToVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApigatewayv2IntegrationTlsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2IntegrationTlsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2IntegrationTlsConfigOutputReference_Override(a Apigatewayv2IntegrationTlsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2IntegrationTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) SetInternalValue(val *Apigatewayv2IntegrationTlsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) SetServerNameToVerify(val *string) {
	_jsii_.Set(
		j,
		"serverNameToVerify",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference) ResetServerNameToVerify() {
	_jsii_.InvokeVoid(
		a,
		"resetServerNameToVerify",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model aws_apigatewayv2_model}.
type Apigatewayv2Model interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Model
type jsiiProxy_Apigatewayv2Model struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Model) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Model) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model aws_apigatewayv2_model} Resource.
func NewApigatewayv2Model(scope constructs.Construct, id *string, config *Apigatewayv2ModelConfig) Apigatewayv2Model {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Model{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Model",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model aws_apigatewayv2_model} Resource.
func NewApigatewayv2Model_Override(a Apigatewayv2Model, scope constructs.Construct, id *string, config *Apigatewayv2ModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Model",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Model) SetSchema(val *string) {
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
func Apigatewayv2Model_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Model",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Model_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Model",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Model) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Model) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Model) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Model) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Model) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Model) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Model) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Model) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Model) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Model) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Model) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Model) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Model) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Model) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Model) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Model) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Model) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Model) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2ModelConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model#api_id Apigatewayv2Model#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model#content_type Apigatewayv2Model#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model#name Apigatewayv2Model#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model#schema Apigatewayv2Model#schema}.
	Schema *string `json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_model#description Apigatewayv2Model#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route aws_apigatewayv2_route}.
type Apigatewayv2Route interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	ApiKeyRequiredInput() interface{}
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	AuthorizationScopesInput() *[]*string
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthorizationTypeInput() *string
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
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelSelectionExpression() *string
	SetModelSelectionExpression(val *string)
	ModelSelectionExpressionInput() *string
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
	RequestParameter() interface{}
	SetRequestParameter(val interface{})
	RequestParameterInput() interface{}
	RouteKey() *string
	SetRouteKey(val *string)
	RouteKeyInput() *string
	RouteResponseSelectionExpression() *string
	SetRouteResponseSelectionExpression(val *string)
	RouteResponseSelectionExpressionInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetAuthorizationType()
	ResetAuthorizerId()
	ResetModelSelectionExpression()
	ResetOperationName()
	ResetOverrideLogicalId()
	ResetRequestModels()
	ResetRequestParameter()
	ResetRouteResponseSelectionExpression()
	ResetTarget()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Route
type jsiiProxy_Apigatewayv2Route struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Route) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ApiKeyRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizationScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ModelSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) ModelSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RequestModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RequestModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RequestParameter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RequestParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RouteKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RouteResponseSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) RouteResponseSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Route) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route aws_apigatewayv2_route} Resource.
func NewApigatewayv2Route(scope constructs.Construct, id *string, config *Apigatewayv2RouteConfig) Apigatewayv2Route {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Route{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Route",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route aws_apigatewayv2_route} Resource.
func NewApigatewayv2Route_Override(a Apigatewayv2Route, scope constructs.Construct, id *string, config *Apigatewayv2RouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Route",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetApiKeyRequired(val interface{}) {
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetAuthorizationScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetAuthorizationType(val *string) {
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetAuthorizerId(val *string) {
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetModelSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"modelSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetOperationName(val *string) {
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetRequestModels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetRequestParameter(val interface{}) {
	_jsii_.Set(
		j,
		"requestParameter",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetRouteKey(val *string) {
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetRouteResponseSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"routeResponseSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Route) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Apigatewayv2Route_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Route",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Route_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Route",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Route) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Route) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Route) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Route) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Route) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Route) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Route) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Route) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Route) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Route) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Route) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Route) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetApiKeyRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeyRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetAuthorizationScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetAuthorizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetModelSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetModelSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetOperationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Route) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetRequestModels() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetRequestParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetRouteResponseSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteResponseSelectionExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) ResetTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Route) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Route) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Route) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Route) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2RouteConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#api_id Apigatewayv2Route#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#route_key Apigatewayv2Route#route_key}.
	RouteKey *string `json:"routeKey" yaml:"routeKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#api_key_required Apigatewayv2Route#api_key_required}.
	ApiKeyRequired interface{} `json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#authorization_scopes Apigatewayv2Route#authorization_scopes}.
	AuthorizationScopes *[]*string `json:"authorizationScopes" yaml:"authorizationScopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#authorization_type Apigatewayv2Route#authorization_type}.
	AuthorizationType *string `json:"authorizationType" yaml:"authorizationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#authorizer_id Apigatewayv2Route#authorizer_id}.
	AuthorizerId *string `json:"authorizerId" yaml:"authorizerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#model_selection_expression Apigatewayv2Route#model_selection_expression}.
	ModelSelectionExpression *string `json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#operation_name Apigatewayv2Route#operation_name}.
	OperationName *string `json:"operationName" yaml:"operationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#request_models Apigatewayv2Route#request_models}.
	RequestModels *map[string]*string `json:"requestModels" yaml:"requestModels"`
	// request_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#request_parameter Apigatewayv2Route#request_parameter}
	RequestParameter interface{} `json:"requestParameter" yaml:"requestParameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#route_response_selection_expression Apigatewayv2Route#route_response_selection_expression}.
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression" yaml:"routeResponseSelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#target Apigatewayv2Route#target}.
	Target *string `json:"target" yaml:"target"`
}

type Apigatewayv2RouteRequestParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#request_parameter_key Apigatewayv2Route#request_parameter_key}.
	RequestParameterKey *string `json:"requestParameterKey" yaml:"requestParameterKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route#required Apigatewayv2Route#required}.
	Required interface{} `json:"required" yaml:"required"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response aws_apigatewayv2_route_response}.
type Apigatewayv2RouteResponse interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
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
	ModelSelectionExpression() *string
	SetModelSelectionExpression(val *string)
	ModelSelectionExpressionInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResponseModels() *map[string]*string
	SetResponseModels(val *map[string]*string)
	ResponseModelsInput() *map[string]*string
	RouteId() *string
	SetRouteId(val *string)
	RouteIdInput() *string
	RouteResponseKey() *string
	SetRouteResponseKey(val *string)
	RouteResponseKeyInput() *string
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
	ResetModelSelectionExpression()
	ResetOverrideLogicalId()
	ResetResponseModels()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2RouteResponse
type jsiiProxy_Apigatewayv2RouteResponse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ModelSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ModelSelectionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ResponseModels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) ResponseModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) RouteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) RouteResponseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) RouteResponseKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response aws_apigatewayv2_route_response} Resource.
func NewApigatewayv2RouteResponse(scope constructs.Construct, id *string, config *Apigatewayv2RouteResponseConfig) Apigatewayv2RouteResponse {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2RouteResponse{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2RouteResponse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response aws_apigatewayv2_route_response} Resource.
func NewApigatewayv2RouteResponse_Override(a Apigatewayv2RouteResponse, scope constructs.Construct, id *string, config *Apigatewayv2RouteResponseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2RouteResponse",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetModelSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"modelSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetResponseModels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"responseModels",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetRouteId(val *string) {
	_jsii_.Set(
		j,
		"routeId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2RouteResponse) SetRouteResponseKey(val *string) {
	_jsii_.Set(
		j,
		"routeResponseKey",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Apigatewayv2RouteResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2RouteResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2RouteResponse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2RouteResponse",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2RouteResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2RouteResponse) ResetModelSelectionExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetModelSelectionExpression",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2RouteResponse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2RouteResponse) ResetResponseModels() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseModels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2RouteResponse) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2RouteResponse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2RouteResponseConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response#api_id Apigatewayv2RouteResponse#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response#route_id Apigatewayv2RouteResponse#route_id}.
	RouteId *string `json:"routeId" yaml:"routeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response#route_response_key Apigatewayv2RouteResponse#route_response_key}.
	RouteResponseKey *string `json:"routeResponseKey" yaml:"routeResponseKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response#model_selection_expression Apigatewayv2RouteResponse#model_selection_expression}.
	ModelSelectionExpression *string `json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_route_response#response_models Apigatewayv2RouteResponse#response_models}.
	ResponseModels *map[string]*string `json:"responseModels" yaml:"responseModels"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage aws_apigatewayv2_stage}.
type Apigatewayv2Stage interface {
	cdktf.TerraformResource
	AccessLogSettings() Apigatewayv2StageAccessLogSettingsOutputReference
	AccessLogSettingsInput() *Apigatewayv2StageAccessLogSettings
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
	AutoDeploy() interface{}
	SetAutoDeploy(val interface{})
	AutoDeployInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	ClientCertificateIdInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DefaultRouteSettings() Apigatewayv2StageDefaultRouteSettingsOutputReference
	DefaultRouteSettingsInput() *Apigatewayv2StageDefaultRouteSettings
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentId() *string
	SetDeploymentId(val *string)
	DeploymentIdInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RouteSettings() interface{}
	SetRouteSettings(val interface{})
	RouteSettingsInput() interface{}
	StageVariables() *map[string]*string
	SetStageVariables(val *map[string]*string)
	StageVariablesInput() *map[string]*string
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
	PutAccessLogSettings(value *Apigatewayv2StageAccessLogSettings)
	PutDefaultRouteSettings(value *Apigatewayv2StageDefaultRouteSettings)
	ResetAccessLogSettings()
	ResetAutoDeploy()
	ResetClientCertificateId()
	ResetDefaultRouteSettings()
	ResetDeploymentId()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetRouteSettings()
	ResetStageVariables()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2Stage
type jsiiProxy_Apigatewayv2Stage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2Stage) AccessLogSettings() Apigatewayv2StageAccessLogSettingsOutputReference {
	var returns Apigatewayv2StageAccessLogSettingsOutputReference
	_jsii_.Get(
		j,
		"accessLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) AccessLogSettingsInput() *Apigatewayv2StageAccessLogSettings {
	var returns *Apigatewayv2StageAccessLogSettings
	_jsii_.Get(
		j,
		"accessLogSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) AutoDeploy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) AutoDeployInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ClientCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DefaultRouteSettings() Apigatewayv2StageDefaultRouteSettingsOutputReference {
	var returns Apigatewayv2StageDefaultRouteSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DefaultRouteSettingsInput() *Apigatewayv2StageDefaultRouteSettings {
	var returns *Apigatewayv2StageDefaultRouteSettings
	_jsii_.Get(
		j,
		"defaultRouteSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) InvokeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) RouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) RouteSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) StageVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) StageVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2Stage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage aws_apigatewayv2_stage} Resource.
func NewApigatewayv2Stage(scope constructs.Construct, id *string, config *Apigatewayv2StageConfig) Apigatewayv2Stage {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2Stage{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Stage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage aws_apigatewayv2_stage} Resource.
func NewApigatewayv2Stage_Override(a Apigatewayv2Stage, scope constructs.Construct, id *string, config *Apigatewayv2StageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Stage",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetAutoDeploy(val interface{}) {
	_jsii_.Set(
		j,
		"autoDeploy",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetClientCertificateId(val *string) {
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetDeploymentId(val *string) {
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"routeSettings",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetStageVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"stageVariables",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2Stage) SetTagsAll(val *map[string]*string) {
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
func Apigatewayv2Stage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Stage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2Stage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2Stage",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Stage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2Stage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2Stage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2Stage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2Stage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) PutAccessLogSettings(value *Apigatewayv2StageAccessLogSettings) {
	_jsii_.InvokeVoid(
		a,
		"putAccessLogSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) PutDefaultRouteSettings(value *Apigatewayv2StageDefaultRouteSettings) {
	_jsii_.InvokeVoid(
		a,
		"putDefaultRouteSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetAccessLogSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetAutoDeploy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDeploy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetClientCertificateId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetDefaultRouteSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRouteSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetDeploymentId() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2Stage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetRouteSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetStageVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetStageVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2Stage) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2Stage) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2Stage) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2Stage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Apigatewayv2StageAccessLogSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#destination_arn Apigatewayv2Stage#destination_arn}.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#format Apigatewayv2Stage#format}.
	Format *string `json:"format" yaml:"format"`
}

type Apigatewayv2StageAccessLogSettingsOutputReference interface {
	cdktf.ComplexObject
	DestinationArn() *string
	SetDestinationArn(val *string)
	DestinationArnInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	InternalValue() *Apigatewayv2StageAccessLogSettings
	SetInternalValue(val *Apigatewayv2StageAccessLogSettings)
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

// The jsii proxy struct for Apigatewayv2StageAccessLogSettingsOutputReference
type jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) DestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) InternalValue() *Apigatewayv2StageAccessLogSettings {
	var returns *Apigatewayv2StageAccessLogSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApigatewayv2StageAccessLogSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2StageAccessLogSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2StageAccessLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2StageAccessLogSettingsOutputReference_Override(a Apigatewayv2StageAccessLogSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2StageAccessLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationArn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetInternalValue(val *Apigatewayv2StageAccessLogSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2StageAccessLogSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2StageConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#api_id Apigatewayv2Stage#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#name Apigatewayv2Stage#name}.
	Name *string `json:"name" yaml:"name"`
	// access_log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#access_log_settings Apigatewayv2Stage#access_log_settings}
	AccessLogSettings *Apigatewayv2StageAccessLogSettings `json:"accessLogSettings" yaml:"accessLogSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#auto_deploy Apigatewayv2Stage#auto_deploy}.
	AutoDeploy interface{} `json:"autoDeploy" yaml:"autoDeploy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#client_certificate_id Apigatewayv2Stage#client_certificate_id}.
	ClientCertificateId *string `json:"clientCertificateId" yaml:"clientCertificateId"`
	// default_route_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#default_route_settings Apigatewayv2Stage#default_route_settings}
	DefaultRouteSettings *Apigatewayv2StageDefaultRouteSettings `json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#deployment_id Apigatewayv2Stage#deployment_id}.
	DeploymentId *string `json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#description Apigatewayv2Stage#description}.
	Description *string `json:"description" yaml:"description"`
	// route_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#route_settings Apigatewayv2Stage#route_settings}
	RouteSettings interface{} `json:"routeSettings" yaml:"routeSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#stage_variables Apigatewayv2Stage#stage_variables}.
	StageVariables *map[string]*string `json:"stageVariables" yaml:"stageVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#tags Apigatewayv2Stage#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#tags_all Apigatewayv2Stage#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type Apigatewayv2StageDefaultRouteSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#data_trace_enabled Apigatewayv2Stage#data_trace_enabled}.
	DataTraceEnabled interface{} `json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#detailed_metrics_enabled Apigatewayv2Stage#detailed_metrics_enabled}.
	DetailedMetricsEnabled interface{} `json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#logging_level Apigatewayv2Stage#logging_level}.
	LoggingLevel *string `json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#throttling_burst_limit Apigatewayv2Stage#throttling_burst_limit}.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#throttling_rate_limit Apigatewayv2Stage#throttling_rate_limit}.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

type Apigatewayv2StageDefaultRouteSettingsOutputReference interface {
	cdktf.ComplexObject
	DataTraceEnabled() interface{}
	SetDataTraceEnabled(val interface{})
	DataTraceEnabledInput() interface{}
	DetailedMetricsEnabled() interface{}
	SetDetailedMetricsEnabled(val interface{})
	DetailedMetricsEnabledInput() interface{}
	InternalValue() *Apigatewayv2StageDefaultRouteSettings
	SetInternalValue(val *Apigatewayv2StageDefaultRouteSettings)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
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
	ResetDataTraceEnabled()
	ResetDetailedMetricsEnabled()
	ResetLoggingLevel()
	ResetThrottlingBurstLimit()
	ResetThrottlingRateLimit()
}

// The jsii proxy struct for Apigatewayv2StageDefaultRouteSettingsOutputReference
type jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) DetailedMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) DetailedMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) InternalValue() *Apigatewayv2StageDefaultRouteSettings {
	var returns *Apigatewayv2StageDefaultRouteSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}

func NewApigatewayv2StageDefaultRouteSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) Apigatewayv2StageDefaultRouteSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2StageDefaultRouteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApigatewayv2StageDefaultRouteSettingsOutputReference_Override(a Apigatewayv2StageDefaultRouteSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2StageDefaultRouteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetDataTraceEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetDetailedMetricsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"detailedMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetInternalValue(val *Apigatewayv2StageDefaultRouteSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetLoggingLevel(val *string) {
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetThrottlingBurstLimit(val *float64) {
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) SetThrottlingRateLimit(val *float64) {
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ResetDetailedMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2StageDefaultRouteSettingsOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

type Apigatewayv2StageRouteSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#route_key Apigatewayv2Stage#route_key}.
	RouteKey *string `json:"routeKey" yaml:"routeKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#data_trace_enabled Apigatewayv2Stage#data_trace_enabled}.
	DataTraceEnabled interface{} `json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#detailed_metrics_enabled Apigatewayv2Stage#detailed_metrics_enabled}.
	DetailedMetricsEnabled interface{} `json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#logging_level Apigatewayv2Stage#logging_level}.
	LoggingLevel *string `json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#throttling_burst_limit Apigatewayv2Stage#throttling_burst_limit}.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_stage#throttling_rate_limit Apigatewayv2Stage#throttling_rate_limit}.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link aws_apigatewayv2_vpc_link}.
type Apigatewayv2VpcLink interface {
	cdktf.TerraformResource
	Arn() *string
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Apigatewayv2VpcLink
type jsiiProxy_Apigatewayv2VpcLink struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Apigatewayv2VpcLink) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link aws_apigatewayv2_vpc_link} Resource.
func NewApigatewayv2VpcLink(scope constructs.Construct, id *string, config *Apigatewayv2VpcLinkConfig) Apigatewayv2VpcLink {
	_init_.Initialize()

	j := jsiiProxy_Apigatewayv2VpcLink{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2VpcLink",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link aws_apigatewayv2_vpc_link} Resource.
func NewApigatewayv2VpcLink_Override(a Apigatewayv2VpcLink, scope constructs.Construct, id *string, config *Apigatewayv2VpcLinkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.Apigatewayv2VpcLink",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Apigatewayv2VpcLink) SetTagsAll(val *map[string]*string) {
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
func Apigatewayv2VpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.Apigatewayv2VpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Apigatewayv2VpcLink_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.Apigatewayv2VpcLink",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2VpcLink) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_Apigatewayv2VpcLink) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_Apigatewayv2VpcLink) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2VpcLink) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2VpcLink) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Apigatewayv2VpcLink) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) ToMetadata() interface{} {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) ToString() *string {
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
func (a *jsiiProxy_Apigatewayv2VpcLink) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type Apigatewayv2VpcLinkConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link#name Apigatewayv2VpcLink#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link#security_group_ids Apigatewayv2VpcLink#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link#subnet_ids Apigatewayv2VpcLink#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link#tags Apigatewayv2VpcLink#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_vpc_link#tags_all Apigatewayv2VpcLink#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_api aws_apigatewayv2_api}.
type DataAwsApigatewayv2Api interface {
	cdktf.TerraformDataSource
	ApiEndpoint() *string
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	ApiKeySelectionExpression() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	DisableExecuteApiEndpoint() cdktf.IResolvable
	ExecutionArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	ProtocolType() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RouteSelectionExpression() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	AddOverride(path *string, value interface{})
	CorsConfiguration(index *string) DataAwsApigatewayv2ApiCorsConfiguration
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

// The jsii proxy struct for DataAwsApigatewayv2Api
type jsiiProxy_DataAwsApigatewayv2Api struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ApiKeySelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) DisableExecuteApiEndpoint() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ExecutionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) RouteSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_api aws_apigatewayv2_api} Data Source.
func NewDataAwsApigatewayv2Api(scope constructs.Construct, id *string, config *DataAwsApigatewayv2ApiConfig) DataAwsApigatewayv2Api {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApigatewayv2Api{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Api",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_api aws_apigatewayv2_api} Data Source.
func NewDataAwsApigatewayv2Api_Override(d DataAwsApigatewayv2Api, scope constructs.Construct, id *string, config *DataAwsApigatewayv2ApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Api",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Api) SetTags(val *map[string]*string) {
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
func DataAwsApigatewayv2Api_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Api",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApigatewayv2Api_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Api",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Api) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Api) CorsConfiguration(index *string) DataAwsApigatewayv2ApiCorsConfiguration {
	var returns DataAwsApigatewayv2ApiCorsConfiguration

	_jsii_.Invoke(
		d,
		"corsConfiguration",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Api) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Api) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Api) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) ToString() *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Api) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type DataAwsApigatewayv2ApiConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_api#api_id DataAwsApigatewayv2Api#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_api#tags DataAwsApigatewayv2Api#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsApigatewayv2ApiCorsConfiguration interface {
	cdktf.ComplexComputedList
	AllowCredentials() cdktf.IResolvable
	AllowHeaders() *[]*string
	AllowMethods() *[]*string
	AllowOrigins() *[]*string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ExposeHeaders() *[]*string
	MaxAge() *float64
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

// The jsii proxy struct for DataAwsApigatewayv2ApiCorsConfiguration
type jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) AllowCredentials() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) AllowHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) AllowMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) AllowOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) ExposeHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsApigatewayv2ApiCorsConfiguration(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsApigatewayv2ApiCorsConfiguration {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2ApiCorsConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsApigatewayv2ApiCorsConfiguration_Override(d DataAwsApigatewayv2ApiCorsConfiguration, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2ApiCorsConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2ApiCorsConfiguration) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis aws_apigatewayv2_apis}.
type DataAwsApigatewayv2Apis interface {
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
	Ids() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
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
	ResetName()
	ResetOverrideLogicalId()
	ResetProtocolType()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApigatewayv2Apis
type jsiiProxy_DataAwsApigatewayv2Apis struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis aws_apigatewayv2_apis} Data Source.
func NewDataAwsApigatewayv2Apis(scope constructs.Construct, id *string, config *DataAwsApigatewayv2ApisConfig) DataAwsApigatewayv2Apis {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApigatewayv2Apis{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Apis",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis aws_apigatewayv2_apis} Data Source.
func NewDataAwsApigatewayv2Apis_Override(d DataAwsApigatewayv2Apis, scope constructs.Construct, id *string, config *DataAwsApigatewayv2ApisConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Apis",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetProtocolType(val *string) {
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Apis) SetTags(val *map[string]*string) {
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
func DataAwsApigatewayv2Apis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Apis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApigatewayv2Apis_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Apis",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Apis) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Apis) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Apis) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Apis) ResetProtocolType() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocolType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Apis) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Apis) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) ToString() *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Apis) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type DataAwsApigatewayv2ApisConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis#name DataAwsApigatewayv2Apis#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis#protocol_type DataAwsApigatewayv2Apis#protocol_type}.
	ProtocolType *string `json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_apis#tags DataAwsApigatewayv2Apis#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export aws_apigatewayv2_export}.
type DataAwsApigatewayv2Export interface {
	cdktf.TerraformDataSource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Body() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExportVersion() *string
	SetExportVersion(val *string)
	ExportVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IncludeExtensions() interface{}
	SetIncludeExtensions(val interface{})
	IncludeExtensionsInput() interface{}
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Specification() *string
	SetSpecification(val *string)
	SpecificationInput() *string
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
	ResetExportVersion()
	ResetIncludeExtensions()
	ResetOverrideLogicalId()
	ResetStageName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsApigatewayv2Export
type jsiiProxy_DataAwsApigatewayv2Export struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ExportVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ExportVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) IncludeExtensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) IncludeExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Specification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export aws_apigatewayv2_export} Data Source.
func NewDataAwsApigatewayv2Export(scope constructs.Construct, id *string, config *DataAwsApigatewayv2ExportConfig) DataAwsApigatewayv2Export {
	_init_.Initialize()

	j := jsiiProxy_DataAwsApigatewayv2Export{}

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Export",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export aws_apigatewayv2_export} Data Source.
func NewDataAwsApigatewayv2Export_Override(d DataAwsApigatewayv2Export, scope constructs.Construct, id *string, config *DataAwsApigatewayv2ExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Export",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetExportVersion(val *string) {
	_jsii_.Set(
		j,
		"exportVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetIncludeExtensions(val interface{}) {
	_jsii_.Set(
		j,
		"includeExtensions",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetOutputType(val *string) {
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetSpecification(val *string) {
	_jsii_.Set(
		j,
		"specification",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SetStageName(val *string) {
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
func DataAwsApigatewayv2Export_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Export",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApigatewayv2Export_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.apigatewayv2.DataAwsApigatewayv2Export",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Export) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetExportVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetExportVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetIncludeExtensions() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeExtensions",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetStageName() {
	_jsii_.InvokeVoid(
		d,
		"resetStageName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) ToString() *string {
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
func (d *jsiiProxy_DataAwsApigatewayv2Export) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon API Gateway Websocket and HTTP APIs.
type DataAwsApigatewayv2ExportConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#api_id DataAwsApigatewayv2Export#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#output_type DataAwsApigatewayv2Export#output_type}.
	OutputType *string `json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#specification DataAwsApigatewayv2Export#specification}.
	Specification *string `json:"specification" yaml:"specification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#export_version DataAwsApigatewayv2Export#export_version}.
	ExportVersion *string `json:"exportVersion" yaml:"exportVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#include_extensions DataAwsApigatewayv2Export#include_extensions}.
	IncludeExtensions interface{} `json:"includeExtensions" yaml:"includeExtensions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/apigatewayv2_export#stage_name DataAwsApigatewayv2Export#stage_name}.
	StageName *string `json:"stageName" yaml:"stageName"`
}
