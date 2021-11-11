package cognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/cognito/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html aws_cognito_identity_pool}.
type CognitoIdentityPool interface {
	cdktf.TerraformResource
	AllowClassicFlow() interface{}
	SetAllowClassicFlow(val interface{})
	AllowClassicFlowInput() interface{}
	AllowUnauthenticatedIdentities() interface{}
	SetAllowUnauthenticatedIdentities(val interface{})
	AllowUnauthenticatedIdentitiesInput() interface{}
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CognitoIdentityProviders() *[]*CognitoIdentityPoolCognitoIdentityProviders
	SetCognitoIdentityProviders(val *[]*CognitoIdentityPoolCognitoIdentityProviders)
	CognitoIdentityProvidersInput() *[]*CognitoIdentityPoolCognitoIdentityProviders
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeveloperProviderName() *string
	SetDeveloperProviderName(val *string)
	DeveloperProviderNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentityPoolName() *string
	SetIdentityPoolName(val *string)
	IdentityPoolNameInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OpenidConnectProviderArns() *[]*string
	SetOpenidConnectProviderArns(val *[]*string)
	OpenidConnectProviderArnsInput() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SamlProviderArns() *[]*string
	SetSamlProviderArns(val *[]*string)
	SamlProviderArnsInput() *[]*string
	SupportedLoginProviders() interface{}
	SetSupportedLoginProviders(val interface{})
	SupportedLoginProvidersInput() interface{}
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAllowClassicFlow()
	ResetAllowUnauthenticatedIdentities()
	ResetCognitoIdentityProviders()
	ResetDeveloperProviderName()
	ResetOpenidConnectProviderArns()
	ResetOverrideLogicalId()
	ResetSamlProviderArns()
	ResetSupportedLoginProviders()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoIdentityPool
type jsiiProxy_CognitoIdentityPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProviders() *[]*CognitoIdentityPoolCognitoIdentityProviders {
	var returns *[]*CognitoIdentityPoolCognitoIdentityProviders
	_jsii_.Get(
		j,
		"cognitoIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProvidersInput() *[]*CognitoIdentityPoolCognitoIdentityProviders {
	var returns *[]*CognitoIdentityPoolCognitoIdentityProviders
	_jsii_.Get(
		j,
		"cognitoIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenidConnectProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openidConnectProviderArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenidConnectProviderArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openidConnectProviderArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedLoginProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedLoginProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html aws_cognito_identity_pool} Resource.
func NewCognitoIdentityPool(scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) CognitoIdentityPool {
	_init_.Initialize()

	j := jsiiProxy_CognitoIdentityPool{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html aws_cognito_identity_pool} Resource.
func NewCognitoIdentityPool_Override(c CognitoIdentityPool, scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetAllowClassicFlow(val interface{}) {
	_jsii_.Set(
		j,
		"allowClassicFlow",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetAllowUnauthenticatedIdentities(val interface{}) {
	_jsii_.Set(
		j,
		"allowUnauthenticatedIdentities",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetCognitoIdentityProviders(val *[]*CognitoIdentityPoolCognitoIdentityProviders) {
	_jsii_.Set(
		j,
		"cognitoIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetDeveloperProviderName(val *string) {
	_jsii_.Set(
		j,
		"developerProviderName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetIdentityPoolName(val *string) {
	_jsii_.Set(
		j,
		"identityPoolName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetOpenidConnectProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"openidConnectProviderArns",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetSamlProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"samlProviderArns",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetSupportedLoginProviders(val interface{}) {
	_jsii_.Set(
		j,
		"supportedLoginProviders",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool) SetTagsAll(val interface{}) {
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
func CognitoIdentityPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoIdentityPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoIdentityPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoIdentityPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetAllowClassicFlow() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowClassicFlow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetAllowUnauthenticatedIdentities() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowUnauthenticatedIdentities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetDeveloperProviderName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeveloperProviderName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetOpenidConnectProviderArns() {
	_jsii_.InvokeVoid(
		c,
		"resetOpenidConnectProviderArns",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSamlProviderArns() {
	_jsii_.InvokeVoid(
		c,
		"resetSamlProviderArns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSupportedLoginProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedLoginProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoIdentityPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoIdentityPoolCognitoIdentityProviders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#client_id CognitoIdentityPool#client_id}.
	ClientId *string `json:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#provider_name CognitoIdentityPool#provider_name}.
	ProviderName *string `json:"providerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#server_side_token_check CognitoIdentityPool#server_side_token_check}.
	ServerSideTokenCheck interface{} `json:"serverSideTokenCheck"`
}

type CognitoIdentityPoolConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#identity_pool_name CognitoIdentityPool#identity_pool_name}.
	IdentityPoolName *string `json:"identityPoolName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#allow_classic_flow CognitoIdentityPool#allow_classic_flow}.
	AllowClassicFlow interface{} `json:"allowClassicFlow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#allow_unauthenticated_identities CognitoIdentityPool#allow_unauthenticated_identities}.
	AllowUnauthenticatedIdentities interface{} `json:"allowUnauthenticatedIdentities"`
	// cognito_identity_providers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#cognito_identity_providers CognitoIdentityPool#cognito_identity_providers}
	CognitoIdentityProviders *[]*CognitoIdentityPoolCognitoIdentityProviders `json:"cognitoIdentityProviders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#developer_provider_name CognitoIdentityPool#developer_provider_name}.
	DeveloperProviderName *string `json:"developerProviderName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#openid_connect_provider_arns CognitoIdentityPool#openid_connect_provider_arns}.
	OpenidConnectProviderArns *[]*string `json:"openidConnectProviderArns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#saml_provider_arns CognitoIdentityPool#saml_provider_arns}.
	SamlProviderArns *[]*string `json:"samlProviderArns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#supported_login_providers CognitoIdentityPool#supported_login_providers}.
	SupportedLoginProviders interface{} `json:"supportedLoginProviders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#tags CognitoIdentityPool#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool.html#tags_all CognitoIdentityPool#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html aws_cognito_identity_pool_roles_attachment}.
type CognitoIdentityPoolRolesAttachment interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleMapping() *[]*CognitoIdentityPoolRolesAttachmentRoleMapping
	SetRoleMapping(val *[]*CognitoIdentityPoolRolesAttachmentRoleMapping)
	RoleMappingInput() *[]*CognitoIdentityPoolRolesAttachmentRoleMapping
	Roles() interface{}
	SetRoles(val interface{})
	RolesInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetRoleMapping()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoIdentityPoolRolesAttachment
type jsiiProxy_CognitoIdentityPoolRolesAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) RoleMapping() *[]*CognitoIdentityPoolRolesAttachmentRoleMapping {
	var returns *[]*CognitoIdentityPoolRolesAttachmentRoleMapping
	_jsii_.Get(
		j,
		"roleMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) RoleMappingInput() *[]*CognitoIdentityPoolRolesAttachmentRoleMapping {
	var returns *[]*CognitoIdentityPoolRolesAttachmentRoleMapping
	_jsii_.Get(
		j,
		"roleMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) Roles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) RolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html aws_cognito_identity_pool_roles_attachment} Resource.
func NewCognitoIdentityPoolRolesAttachment(scope constructs.Construct, id *string, config *CognitoIdentityPoolRolesAttachmentConfig) CognitoIdentityPoolRolesAttachment {
	_init_.Initialize()

	j := jsiiProxy_CognitoIdentityPoolRolesAttachment{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html aws_cognito_identity_pool_roles_attachment} Resource.
func NewCognitoIdentityPoolRolesAttachment_Override(c CognitoIdentityPoolRolesAttachment, scope constructs.Construct, id *string, config *CognitoIdentityPoolRolesAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachment",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetIdentityPoolId(val *string) {
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetRoleMapping(val *[]*CognitoIdentityPoolRolesAttachmentRoleMapping) {
	_jsii_.Set(
		j,
		"roleMapping",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPoolRolesAttachment) SetRoles(val interface{}) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoIdentityPoolRolesAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoIdentityPoolRolesAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) ResetRoleMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoIdentityPoolRolesAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoIdentityPoolRolesAttachmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#identity_pool_id CognitoIdentityPoolRolesAttachment#identity_pool_id}.
	IdentityPoolId *string `json:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#roles CognitoIdentityPoolRolesAttachment#roles}.
	Roles interface{} `json:"roles"`
	// role_mapping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#role_mapping CognitoIdentityPoolRolesAttachment#role_mapping}
	RoleMapping *[]*CognitoIdentityPoolRolesAttachmentRoleMapping `json:"roleMapping"`
}

type CognitoIdentityPoolRolesAttachmentRoleMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#identity_provider CognitoIdentityPoolRolesAttachment#identity_provider}.
	IdentityProvider *string `json:"identityProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#type CognitoIdentityPoolRolesAttachment#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#ambiguous_role_resolution CognitoIdentityPoolRolesAttachment#ambiguous_role_resolution}.
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution"`
	// mapping_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#mapping_rule CognitoIdentityPoolRolesAttachment#mapping_rule}
	MappingRule *[]*CognitoIdentityPoolRolesAttachmentRoleMappingMappingRule `json:"mappingRule"`
}

type CognitoIdentityPoolRolesAttachmentRoleMappingMappingRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#claim CognitoIdentityPoolRolesAttachment#claim}.
	Claim *string `json:"claim"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#match_type CognitoIdentityPoolRolesAttachment#match_type}.
	MatchType *string `json:"matchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#role_arn CognitoIdentityPoolRolesAttachment#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool_roles_attachment.html#value CognitoIdentityPoolRolesAttachment#value}.
	Value *string `json:"value"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html aws_cognito_identity_provider}.
type CognitoIdentityProvider interface {
	cdktf.TerraformResource
	AttributeMapping() interface{}
	SetAttributeMapping(val interface{})
	AttributeMappingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdpIdentifiers() *[]*string
	SetIdpIdentifiers(val *[]*string)
	IdpIdentifiersInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProviderDetails() interface{}
	SetProviderDetails(val interface{})
	ProviderDetailsInput() interface{}
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAttributeMapping()
	ResetIdpIdentifiers()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoIdentityProvider
type jsiiProxy_CognitoIdentityProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoIdentityProvider) AttributeMapping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) AttributeMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) IdpIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) IdpIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityProvider) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html aws_cognito_identity_provider} Resource.
func NewCognitoIdentityProvider(scope constructs.Construct, id *string, config *CognitoIdentityProviderConfig) CognitoIdentityProvider {
	_init_.Initialize()

	j := jsiiProxy_CognitoIdentityProvider{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html aws_cognito_identity_provider} Resource.
func NewCognitoIdentityProvider_Override(c CognitoIdentityProvider, scope constructs.Construct, id *string, config *CognitoIdentityProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoIdentityProvider",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetAttributeMapping(val interface{}) {
	_jsii_.Set(
		j,
		"attributeMapping",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetIdpIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"idpIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetProviderDetails(val interface{}) {
	_jsii_.Set(
		j,
		"providerDetails",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetProviderName(val *string) {
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetProviderType(val *string) {
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityProvider) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoIdentityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoIdentityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoIdentityProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoIdentityProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoIdentityProvider) ResetAttributeMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributeMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityProvider) ResetIdpIdentifiers() {
	_jsii_.InvokeVoid(
		c,
		"resetIdpIdentifiers",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoIdentityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoIdentityProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoIdentityProviderConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#provider_details CognitoIdentityProvider#provider_details}.
	ProviderDetails interface{} `json:"providerDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#provider_name CognitoIdentityProvider#provider_name}.
	ProviderName *string `json:"providerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#provider_type CognitoIdentityProvider#provider_type}.
	ProviderType *string `json:"providerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#user_pool_id CognitoIdentityProvider#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#attribute_mapping CognitoIdentityProvider#attribute_mapping}.
	AttributeMapping interface{} `json:"attributeMapping"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_provider.html#idp_identifiers CognitoIdentityProvider#idp_identifiers}.
	IdpIdentifiers *[]*string `json:"idpIdentifiers"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html aws_cognito_resource_server}.
type CognitoResourceServer interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Scope() *[]*CognitoResourceServerScope
	SetScope(val *[]*CognitoResourceServerScope)
	ScopeIdentifiers() *[]*string
	ScopeInput() *[]*CognitoResourceServerScope
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetScope()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoResourceServer
type jsiiProxy_CognitoResourceServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoResourceServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) Scope() *[]*CognitoResourceServerScope {
	var returns *[]*CognitoResourceServerScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) ScopeIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) ScopeInput() *[]*CognitoResourceServerScope {
	var returns *[]*CognitoResourceServerScope
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoResourceServer) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html aws_cognito_resource_server} Resource.
func NewCognitoResourceServer(scope constructs.Construct, id *string, config *CognitoResourceServerConfig) CognitoResourceServer {
	_init_.Initialize()

	j := jsiiProxy_CognitoResourceServer{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoResourceServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html aws_cognito_resource_server} Resource.
func NewCognitoResourceServer_Override(c CognitoResourceServer, scope constructs.Construct, id *string, config *CognitoResourceServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoResourceServer",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetScope(val *[]*CognitoResourceServerScope) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CognitoResourceServer) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoResourceServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoResourceServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoResourceServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoResourceServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoResourceServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoResourceServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoResourceServer) ResetScope() {
	_jsii_.InvokeVoid(
		c,
		"resetScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoResourceServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoResourceServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoResourceServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoResourceServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoResourceServerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#identifier CognitoResourceServer#identifier}.
	Identifier *string `json:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#name CognitoResourceServer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#user_pool_id CognitoResourceServer#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#scope CognitoResourceServer#scope}
	Scope *[]*CognitoResourceServerScope `json:"scope"`
}

type CognitoResourceServerScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#scope_description CognitoResourceServer#scope_description}.
	ScopeDescription *string `json:"scopeDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server.html#scope_name CognitoResourceServer#scope_name}.
	ScopeName *string `json:"scopeName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html aws_cognito_user_group}.
type CognitoUserGroup interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetPrecedence()
	ResetRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserGroup
type jsiiProxy_CognitoUserGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserGroup) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html aws_cognito_user_group} Resource.
func NewCognitoUserGroup(scope constructs.Construct, id *string, config *CognitoUserGroupConfig) CognitoUserGroup {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserGroup{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html aws_cognito_user_group} Resource.
func NewCognitoUserGroup_Override(c CognitoUserGroup, scope constructs.Construct, id *string, config *CognitoUserGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetPrecedence(val *float64) {
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserGroup) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoUserGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoUserGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoUserGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoUserGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoUserGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserGroup) ResetPrecedence() {
	_jsii_.InvokeVoid(
		c,
		"resetPrecedence",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserGroup) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoUserGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoUserGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoUserGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html#name CognitoUserGroup#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html#user_pool_id CognitoUserGroup#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html#description CognitoUserGroup#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html#precedence CognitoUserGroup#precedence}.
	Precedence *float64 `json:"precedence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_group.html#role_arn CognitoUserGroup#role_arn}.
	RoleArn *string `json:"roleArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html aws_cognito_user_pool}.
type CognitoUserPool interface {
	cdktf.TerraformResource
	AccountRecoverySetting() CognitoUserPoolAccountRecoverySettingOutputReference
	AccountRecoverySettingInput() *CognitoUserPoolAccountRecoverySetting
	AdminCreateUserConfig() CognitoUserPoolAdminCreateUserConfigOutputReference
	AdminCreateUserConfigInput() *CognitoUserPoolAdminCreateUserConfig
	AliasAttributes() *[]*string
	SetAliasAttributes(val *[]*string)
	AliasAttributesInput() *[]*string
	Arn() *string
	AutoVerifiedAttributes() *[]*string
	SetAutoVerifiedAttributes(val *[]*string)
	AutoVerifiedAttributesInput() *[]*string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreationDate() *string
	CustomDomain() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeviceConfiguration() CognitoUserPoolDeviceConfigurationOutputReference
	DeviceConfigurationInput() *CognitoUserPoolDeviceConfiguration
	Domain() *string
	EmailConfiguration() CognitoUserPoolEmailConfigurationOutputReference
	EmailConfigurationInput() *CognitoUserPoolEmailConfiguration
	EmailVerificationMessage() *string
	SetEmailVerificationMessage(val *string)
	EmailVerificationMessageInput() *string
	EmailVerificationSubject() *string
	SetEmailVerificationSubject(val *string)
	EmailVerificationSubjectInput() *string
	Endpoint() *string
	EstimatedNumberOfUsers() *float64
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LambdaConfig() CognitoUserPoolLambdaConfigOutputReference
	LambdaConfigInput() *CognitoUserPoolLambdaConfig
	LastModifiedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MfaConfiguration() *string
	SetMfaConfiguration(val *string)
	MfaConfigurationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PasswordPolicy() CognitoUserPoolPasswordPolicyOutputReference
	PasswordPolicyInput() *CognitoUserPoolPasswordPolicy
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Schema() *[]*CognitoUserPoolSchema
	SetSchema(val *[]*CognitoUserPoolSchema)
	SchemaInput() *[]*CognitoUserPoolSchema
	SmsAuthenticationMessage() *string
	SetSmsAuthenticationMessage(val *string)
	SmsAuthenticationMessageInput() *string
	SmsConfiguration() CognitoUserPoolSmsConfigurationOutputReference
	SmsConfigurationInput() *CognitoUserPoolSmsConfiguration
	SmsVerificationMessage() *string
	SetSmsVerificationMessage(val *string)
	SmsVerificationMessageInput() *string
	SoftwareTokenMfaConfiguration() CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference
	SoftwareTokenMfaConfigurationInput() *CognitoUserPoolSoftwareTokenMfaConfiguration
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UsernameAttributes() *[]*string
	SetUsernameAttributes(val *[]*string)
	UsernameAttributesInput() *[]*string
	UsernameConfiguration() CognitoUserPoolUsernameConfigurationOutputReference
	UsernameConfigurationInput() *CognitoUserPoolUsernameConfiguration
	UserPoolAddOns() CognitoUserPoolUserPoolAddOnsOutputReference
	UserPoolAddOnsInput() *CognitoUserPoolUserPoolAddOns
	VerificationMessageTemplate() CognitoUserPoolVerificationMessageTemplateOutputReference
	VerificationMessageTemplateInput() *CognitoUserPoolVerificationMessageTemplate
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAccountRecoverySetting(value *CognitoUserPoolAccountRecoverySetting)
	PutAdminCreateUserConfig(value *CognitoUserPoolAdminCreateUserConfig)
	PutDeviceConfiguration(value *CognitoUserPoolDeviceConfiguration)
	PutEmailConfiguration(value *CognitoUserPoolEmailConfiguration)
	PutLambdaConfig(value *CognitoUserPoolLambdaConfig)
	PutPasswordPolicy(value *CognitoUserPoolPasswordPolicy)
	PutSmsConfiguration(value *CognitoUserPoolSmsConfiguration)
	PutSoftwareTokenMfaConfiguration(value *CognitoUserPoolSoftwareTokenMfaConfiguration)
	PutUsernameConfiguration(value *CognitoUserPoolUsernameConfiguration)
	PutUserPoolAddOns(value *CognitoUserPoolUserPoolAddOns)
	PutVerificationMessageTemplate(value *CognitoUserPoolVerificationMessageTemplate)
	ResetAccountRecoverySetting()
	ResetAdminCreateUserConfig()
	ResetAliasAttributes()
	ResetAutoVerifiedAttributes()
	ResetDeviceConfiguration()
	ResetEmailConfiguration()
	ResetEmailVerificationMessage()
	ResetEmailVerificationSubject()
	ResetLambdaConfig()
	ResetMfaConfiguration()
	ResetOverrideLogicalId()
	ResetPasswordPolicy()
	ResetSchema()
	ResetSmsAuthenticationMessage()
	ResetSmsConfiguration()
	ResetSmsVerificationMessage()
	ResetSoftwareTokenMfaConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetUsernameAttributes()
	ResetUsernameConfiguration()
	ResetUserPoolAddOns()
	ResetVerificationMessageTemplate()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPool
type jsiiProxy_CognitoUserPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPool) AccountRecoverySetting() CognitoUserPoolAccountRecoverySettingOutputReference {
	var returns CognitoUserPoolAccountRecoverySettingOutputReference
	_jsii_.Get(
		j,
		"accountRecoverySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AccountRecoverySettingInput() *CognitoUserPoolAccountRecoverySetting {
	var returns *CognitoUserPoolAccountRecoverySetting
	_jsii_.Get(
		j,
		"accountRecoverySettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AdminCreateUserConfig() CognitoUserPoolAdminCreateUserConfigOutputReference {
	var returns CognitoUserPoolAdminCreateUserConfigOutputReference
	_jsii_.Get(
		j,
		"adminCreateUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AdminCreateUserConfigInput() *CognitoUserPoolAdminCreateUserConfig {
	var returns *CognitoUserPoolAdminCreateUserConfig
	_jsii_.Get(
		j,
		"adminCreateUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AliasAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AliasAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AutoVerifiedAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) AutoVerifiedAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoVerifiedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeviceConfiguration() CognitoUserPoolDeviceConfigurationOutputReference {
	var returns CognitoUserPoolDeviceConfigurationOutputReference
	_jsii_.Get(
		j,
		"deviceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) DeviceConfigurationInput() *CognitoUserPoolDeviceConfiguration {
	var returns *CognitoUserPoolDeviceConfiguration
	_jsii_.Get(
		j,
		"deviceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailConfiguration() CognitoUserPoolEmailConfigurationOutputReference {
	var returns CognitoUserPoolEmailConfigurationOutputReference
	_jsii_.Get(
		j,
		"emailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailConfigurationInput() *CognitoUserPoolEmailConfiguration {
	var returns *CognitoUserPoolEmailConfiguration
	_jsii_.Get(
		j,
		"emailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EmailVerificationSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailVerificationSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) EstimatedNumberOfUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedNumberOfUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LambdaConfig() CognitoUserPoolLambdaConfigOutputReference {
	var returns CognitoUserPoolLambdaConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LambdaConfigInput() *CognitoUserPoolLambdaConfig {
	var returns *CognitoUserPoolLambdaConfig
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) MfaConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) MfaConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) PasswordPolicy() CognitoUserPoolPasswordPolicyOutputReference {
	var returns CognitoUserPoolPasswordPolicyOutputReference
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) PasswordPolicyInput() *CognitoUserPoolPasswordPolicy {
	var returns *CognitoUserPoolPasswordPolicy
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Schema() *[]*CognitoUserPoolSchema {
	var returns *[]*CognitoUserPoolSchema
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SchemaInput() *[]*CognitoUserPoolSchema {
	var returns *[]*CognitoUserPoolSchema
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsAuthenticationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsAuthenticationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsAuthenticationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsConfiguration() CognitoUserPoolSmsConfigurationOutputReference {
	var returns CognitoUserPoolSmsConfigurationOutputReference
	_jsii_.Get(
		j,
		"smsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsConfigurationInput() *CognitoUserPoolSmsConfiguration {
	var returns *CognitoUserPoolSmsConfiguration
	_jsii_.Get(
		j,
		"smsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsVerificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SmsVerificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsVerificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SoftwareTokenMfaConfiguration() CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference {
	var returns CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference
	_jsii_.Get(
		j,
		"softwareTokenMfaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) SoftwareTokenMfaConfigurationInput() *CognitoUserPoolSoftwareTokenMfaConfiguration {
	var returns *CognitoUserPoolSoftwareTokenMfaConfiguration
	_jsii_.Get(
		j,
		"softwareTokenMfaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usernameAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameConfiguration() CognitoUserPoolUsernameConfigurationOutputReference {
	var returns CognitoUserPoolUsernameConfigurationOutputReference
	_jsii_.Get(
		j,
		"usernameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UsernameConfigurationInput() *CognitoUserPoolUsernameConfiguration {
	var returns *CognitoUserPoolUsernameConfiguration
	_jsii_.Get(
		j,
		"usernameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserPoolAddOns() CognitoUserPoolUserPoolAddOnsOutputReference {
	var returns CognitoUserPoolUserPoolAddOnsOutputReference
	_jsii_.Get(
		j,
		"userPoolAddOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) UserPoolAddOnsInput() *CognitoUserPoolUserPoolAddOns {
	var returns *CognitoUserPoolUserPoolAddOns
	_jsii_.Get(
		j,
		"userPoolAddOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) VerificationMessageTemplate() CognitoUserPoolVerificationMessageTemplateOutputReference {
	var returns CognitoUserPoolVerificationMessageTemplateOutputReference
	_jsii_.Get(
		j,
		"verificationMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPool) VerificationMessageTemplateInput() *CognitoUserPoolVerificationMessageTemplate {
	var returns *CognitoUserPoolVerificationMessageTemplate
	_jsii_.Get(
		j,
		"verificationMessageTemplateInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html aws_cognito_user_pool} Resource.
func NewCognitoUserPool(scope constructs.Construct, id *string, config *CognitoUserPoolConfig) CognitoUserPool {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPool{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html aws_cognito_user_pool} Resource.
func NewCognitoUserPool_Override(c CognitoUserPool, scope constructs.Construct, id *string, config *CognitoUserPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPool",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetAliasAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"aliasAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetAutoVerifiedAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"autoVerifiedAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetEmailVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetEmailVerificationSubject(val *string) {
	_jsii_.Set(
		j,
		"emailVerificationSubject",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetMfaConfiguration(val *string) {
	_jsii_.Set(
		j,
		"mfaConfiguration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetSchema(val *[]*CognitoUserPoolSchema) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetSmsAuthenticationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsAuthenticationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetSmsVerificationMessage(val *string) {
	_jsii_.Set(
		j,
		"smsVerificationMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPool) SetUsernameAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"usernameAttributes",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoUserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoUserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoUserPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoUserPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutAccountRecoverySetting(value *CognitoUserPoolAccountRecoverySetting) {
	_jsii_.InvokeVoid(
		c,
		"putAccountRecoverySetting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutAdminCreateUserConfig(value *CognitoUserPoolAdminCreateUserConfig) {
	_jsii_.InvokeVoid(
		c,
		"putAdminCreateUserConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutDeviceConfiguration(value *CognitoUserPoolDeviceConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putDeviceConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutEmailConfiguration(value *CognitoUserPoolEmailConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putEmailConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutLambdaConfig(value *CognitoUserPoolLambdaConfig) {
	_jsii_.InvokeVoid(
		c,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutPasswordPolicy(value *CognitoUserPoolPasswordPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putPasswordPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutSmsConfiguration(value *CognitoUserPoolSmsConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putSmsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutSoftwareTokenMfaConfiguration(value *CognitoUserPoolSoftwareTokenMfaConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putSoftwareTokenMfaConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutUsernameConfiguration(value *CognitoUserPoolUsernameConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putUsernameConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutUserPoolAddOns(value *CognitoUserPoolUserPoolAddOns) {
	_jsii_.InvokeVoid(
		c,
		"putUserPoolAddOns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) PutVerificationMessageTemplate(value *CognitoUserPoolVerificationMessageTemplate) {
	_jsii_.InvokeVoid(
		c,
		"putVerificationMessageTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAccountRecoverySetting() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountRecoverySetting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAdminCreateUserConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminCreateUserConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAliasAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAliasAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetAutoVerifiedAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoVerifiedAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetDeviceConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailVerificationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailVerificationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetEmailVerificationSubject() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailVerificationSubject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetMfaConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetMfaConfiguration",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoUserPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSchema() {
	_jsii_.InvokeVoid(
		c,
		"resetSchema",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsAuthenticationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsAuthenticationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSmsVerificationMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsVerificationMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetSoftwareTokenMfaConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSoftwareTokenMfaConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUsernameAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUsernameConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetUserPoolAddOns() {
	_jsii_.InvokeVoid(
		c,
		"resetUserPoolAddOns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) ResetVerificationMessageTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetVerificationMessageTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoUserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoUserPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoUserPoolAccountRecoverySetting struct {
	// recovery_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#recovery_mechanism CognitoUserPool#recovery_mechanism}
	RecoveryMechanism *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism `json:"recoveryMechanism"`
}

type CognitoUserPoolAccountRecoverySettingOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RecoveryMechanism() *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism
	SetRecoveryMechanism(val *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism)
	RecoveryMechanismInput() *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolAccountRecoverySettingOutputReference
type jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) RecoveryMechanism() *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism {
	var returns *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism
	_jsii_.Get(
		j,
		"recoveryMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) RecoveryMechanismInput() *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism {
	var returns *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism
	_jsii_.Get(
		j,
		"recoveryMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolAccountRecoverySettingOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolAccountRecoverySettingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAccountRecoverySettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolAccountRecoverySettingOutputReference_Override(c CognitoUserPoolAccountRecoverySettingOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAccountRecoverySettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) SetRecoveryMechanism(val *[]*CognitoUserPoolAccountRecoverySettingRecoveryMechanism) {
	_jsii_.Set(
		j,
		"recoveryMechanism",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolAccountRecoverySettingRecoveryMechanism struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#name CognitoUserPool#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#priority CognitoUserPool#priority}.
	Priority *float64 `json:"priority"`
}

type CognitoUserPoolAdminCreateUserConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#allow_admin_create_user_only CognitoUserPool#allow_admin_create_user_only}.
	AllowAdminCreateUserOnly interface{} `json:"allowAdminCreateUserOnly"`
	// invite_message_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#invite_message_template CognitoUserPool#invite_message_template}
	InviteMessageTemplate *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate `json:"inviteMessageTemplate"`
}

type CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_message CognitoUserPool#email_message}.
	EmailMessage *string `json:"emailMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_subject CognitoUserPool#email_subject}.
	EmailSubject *string `json:"emailSubject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sms_message CognitoUserPool#sms_message}.
	SmsMessage *string `json:"smsMessage"`
}

type CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference interface {
	cdktf.ComplexObject
	EmailMessage() *string
	SetEmailMessage(val *string)
	EmailMessageInput() *string
	EmailSubject() *string
	SetEmailSubject(val *string)
	EmailSubjectInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SmsMessage() *string
	SetSmsMessage(val *string)
	SmsMessageInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEmailMessage()
	ResetEmailSubject()
	ResetSmsMessage()
}

// The jsii proxy struct for CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference
type jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) EmailMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) EmailMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SmsMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SmsMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference_Override(c CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetEmailMessage(val *string) {
	_jsii_.Set(
		j,
		"emailMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetEmailSubject(val *string) {
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetSmsMessage(val *string) {
	_jsii_.Set(
		j,
		"smsMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) ResetEmailMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference) ResetSmsMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsMessage",
		nil, // no parameters
	)
}

type CognitoUserPoolAdminCreateUserConfigOutputReference interface {
	cdktf.ComplexObject
	AllowAdminCreateUserOnly() interface{}
	SetAllowAdminCreateUserOnly(val interface{})
	AllowAdminCreateUserOnlyInput() interface{}
	InviteMessageTemplate() CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference
	InviteMessageTemplateInput() *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutInviteMessageTemplate(value *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate)
	ResetAllowAdminCreateUserOnly()
	ResetInviteMessageTemplate()
}

// The jsii proxy struct for CognitoUserPoolAdminCreateUserConfigOutputReference
type jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) AllowAdminCreateUserOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAdminCreateUserOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) AllowAdminCreateUserOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAdminCreateUserOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) InviteMessageTemplate() CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference {
	var returns CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference
	_jsii_.Get(
		j,
		"inviteMessageTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) InviteMessageTemplateInput() *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate {
	var returns *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate
	_jsii_.Get(
		j,
		"inviteMessageTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolAdminCreateUserConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolAdminCreateUserConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolAdminCreateUserConfigOutputReference_Override(c CognitoUserPoolAdminCreateUserConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) SetAllowAdminCreateUserOnly(val interface{}) {
	_jsii_.Set(
		j,
		"allowAdminCreateUserOnly",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) PutInviteMessageTemplate(value *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate) {
	_jsii_.InvokeVoid(
		c,
		"putInviteMessageTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) ResetAllowAdminCreateUserOnly() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowAdminCreateUserOnly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference) ResetInviteMessageTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetInviteMessageTemplate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html aws_cognito_user_pool_client}.
type CognitoUserPoolClient interface {
	cdktf.TerraformResource
	AccessTokenValidity() *float64
	SetAccessTokenValidity(val *float64)
	AccessTokenValidityInput() *float64
	AllowedOauthFlows() *[]*string
	SetAllowedOauthFlows(val *[]*string)
	AllowedOauthFlowsInput() *[]*string
	AllowedOauthFlowsUserPoolClient() interface{}
	SetAllowedOauthFlowsUserPoolClient(val interface{})
	AllowedOauthFlowsUserPoolClientInput() interface{}
	AllowedOauthScopes() *[]*string
	SetAllowedOauthScopes(val *[]*string)
	AllowedOauthScopesInput() *[]*string
	AnalyticsConfiguration() CognitoUserPoolClientAnalyticsConfigurationOutputReference
	AnalyticsConfigurationInput() *CognitoUserPoolClientAnalyticsConfiguration
	CallbackUrls() *[]*string
	SetCallbackUrls(val *[]*string)
	CallbackUrlsInput() *[]*string
	CdktfStack() cdktf.TerraformStack
	ClientSecret() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultRedirectUri() *string
	SetDefaultRedirectUri(val *string)
	DefaultRedirectUriInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnableTokenRevocation() interface{}
	SetEnableTokenRevocation(val interface{})
	EnableTokenRevocationInput() interface{}
	ExplicitAuthFlows() *[]*string
	SetExplicitAuthFlows(val *[]*string)
	ExplicitAuthFlowsInput() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	GenerateSecret() interface{}
	SetGenerateSecret(val interface{})
	GenerateSecretInput() interface{}
	Id() *string
	IdTokenValidity() *float64
	SetIdTokenValidity(val *float64)
	IdTokenValidityInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoutUrls() *[]*string
	SetLogoutUrls(val *[]*string)
	LogoutUrlsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PreventUserExistenceErrors() *string
	SetPreventUserExistenceErrors(val *string)
	PreventUserExistenceErrorsInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReadAttributes() *[]*string
	SetReadAttributes(val *[]*string)
	ReadAttributesInput() *[]*string
	RefreshTokenValidity() *float64
	SetRefreshTokenValidity(val *float64)
	RefreshTokenValidityInput() *float64
	SupportedIdentityProviders() *[]*string
	SetSupportedIdentityProviders(val *[]*string)
	SupportedIdentityProvidersInput() *[]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenValidityUnits() CognitoUserPoolClientTokenValidityUnitsOutputReference
	TokenValidityUnitsInput() *CognitoUserPoolClientTokenValidityUnits
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	WriteAttributes() *[]*string
	SetWriteAttributes(val *[]*string)
	WriteAttributesInput() *[]*string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAnalyticsConfiguration(value *CognitoUserPoolClientAnalyticsConfiguration)
	PutTokenValidityUnits(value *CognitoUserPoolClientTokenValidityUnits)
	ResetAccessTokenValidity()
	ResetAllowedOauthFlows()
	ResetAllowedOauthFlowsUserPoolClient()
	ResetAllowedOauthScopes()
	ResetAnalyticsConfiguration()
	ResetCallbackUrls()
	ResetDefaultRedirectUri()
	ResetEnableTokenRevocation()
	ResetExplicitAuthFlows()
	ResetGenerateSecret()
	ResetIdTokenValidity()
	ResetLogoutUrls()
	ResetOverrideLogicalId()
	ResetPreventUserExistenceErrors()
	ResetReadAttributes()
	ResetRefreshTokenValidity()
	ResetSupportedIdentityProviders()
	ResetTokenValidityUnits()
	ResetWriteAttributes()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPoolClient
type jsiiProxy_CognitoUserPoolClient struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthFlowsUserPoolClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOauthFlowsUserPoolClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AllowedOauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AnalyticsConfiguration() CognitoUserPoolClientAnalyticsConfigurationOutputReference {
	var returns CognitoUserPoolClientAnalyticsConfigurationOutputReference
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) AnalyticsConfigurationInput() *CognitoUserPoolClientAnalyticsConfiguration {
	var returns *CognitoUserPoolClientAnalyticsConfiguration
	_jsii_.Get(
		j,
		"analyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CallbackUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CallbackUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DefaultRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) EnableTokenRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ExplicitAuthFlowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) GenerateSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) GenerateSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) IdTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) LogoutUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) LogoutUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) PreventUserExistenceErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) ReadAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) RefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) SupportedIdentityProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TokenValidityUnits() CognitoUserPoolClientTokenValidityUnitsOutputReference {
	var returns CognitoUserPoolClientTokenValidityUnitsOutputReference
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) TokenValidityUnitsInput() *CognitoUserPoolClientTokenValidityUnits {
	var returns *CognitoUserPoolClientTokenValidityUnits
	_jsii_.Get(
		j,
		"tokenValidityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClient) WriteAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html aws_cognito_user_pool_client} Resource.
func NewCognitoUserPoolClient(scope constructs.Construct, id *string, config *CognitoUserPoolClientConfig) CognitoUserPoolClient {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolClient{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html aws_cognito_user_pool_client} Resource.
func NewCognitoUserPoolClient_Override(c CognitoUserPoolClient, scope constructs.Construct, id *string, config *CognitoUserPoolClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClient",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetAccessTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetAllowedOauthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOauthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetAllowedOauthFlowsUserPoolClient(val interface{}) {
	_jsii_.Set(
		j,
		"allowedOauthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetAllowedOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOauthScopes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetCallbackUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"callbackUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetDefaultRedirectUri(val *string) {
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetEnableTokenRevocation(val interface{}) {
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetExplicitAuthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetGenerateSecret(val interface{}) {
	_jsii_.Set(
		j,
		"generateSecret",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetIdTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetLogoutUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"logoutUrls",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetPreventUserExistenceErrors(val *string) {
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetReadAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetRefreshTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetSupportedIdentityProviders(val *[]*string) {
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClient) SetWriteAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"writeAttributes",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPoolClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoUserPoolClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) PutAnalyticsConfiguration(value *CognitoUserPoolClientAnalyticsConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) PutTokenValidityUnits(value *CognitoUserPoolClientTokenValidityUnits) {
	_jsii_.InvokeVoid(
		c,
		"putTokenValidityUnits",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAccessTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthFlowsUserPoolClient() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthFlowsUserPoolClient",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAllowedOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetCallbackUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetCallbackUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetDefaultRedirectUri() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRedirectUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetEnableTokenRevocation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTokenRevocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetExplicitAuthFlows() {
	_jsii_.InvokeVoid(
		c,
		"resetExplicitAuthFlows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetGenerateSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetGenerateSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetIdTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetLogoutUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetLogoutUrls",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetPreventUserExistenceErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUserExistenceErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetReadAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetReadAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshTokenValidity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetSupportedIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetTokenValidityUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenValidityUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) ResetWriteAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetWriteAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoUserPoolClientAnalyticsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#application_arn CognitoUserPoolClient#application_arn}.
	ApplicationArn *string `json:"applicationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#application_id CognitoUserPoolClient#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#external_id CognitoUserPoolClient#external_id}.
	ExternalId *string `json:"externalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#role_arn CognitoUserPoolClient#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#user_data_shared CognitoUserPoolClient#user_data_shared}.
	UserDataShared interface{} `json:"userDataShared"`
}

type CognitoUserPoolClientAnalyticsConfigurationOutputReference interface {
	cdktf.ComplexObject
	ApplicationArn() *string
	SetApplicationArn(val *string)
	ApplicationArnInput() *string
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserDataShared() interface{}
	SetUserDataShared(val interface{})
	UserDataSharedInput() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetApplicationArn()
	ResetApplicationId()
	ResetExternalId()
	ResetRoleArn()
	ResetUserDataShared()
}

// The jsii proxy struct for CognitoUserPoolClientAnalyticsConfigurationOutputReference
type jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ApplicationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) UserDataShared() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataShared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) UserDataSharedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataSharedInput",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolClientAnalyticsConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolClientAnalyticsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClientAnalyticsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolClientAnalyticsConfigurationOutputReference_Override(c CognitoUserPoolClientAnalyticsConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClientAnalyticsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetApplicationArn(val *string) {
	_jsii_.Set(
		j,
		"applicationArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetExternalId(val *string) {
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) SetUserDataShared(val interface{}) {
	_jsii_.Set(
		j,
		"userDataShared",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ResetApplicationArn() {
	_jsii_.InvokeVoid(
		c,
		"resetApplicationArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ResetApplicationId() {
	_jsii_.InvokeVoid(
		c,
		"resetApplicationId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference) ResetUserDataShared() {
	_jsii_.InvokeVoid(
		c,
		"resetUserDataShared",
		nil, // no parameters
	)
}

type CognitoUserPoolClientConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#name CognitoUserPoolClient#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#user_pool_id CognitoUserPoolClient#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#access_token_validity CognitoUserPoolClient#access_token_validity}.
	AccessTokenValidity *float64 `json:"accessTokenValidity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#allowed_oauth_flows CognitoUserPoolClient#allowed_oauth_flows}.
	AllowedOauthFlows *[]*string `json:"allowedOauthFlows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#allowed_oauth_flows_user_pool_client CognitoUserPoolClient#allowed_oauth_flows_user_pool_client}.
	AllowedOauthFlowsUserPoolClient interface{} `json:"allowedOauthFlowsUserPoolClient"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#allowed_oauth_scopes CognitoUserPoolClient#allowed_oauth_scopes}.
	AllowedOauthScopes *[]*string `json:"allowedOauthScopes"`
	// analytics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#analytics_configuration CognitoUserPoolClient#analytics_configuration}
	AnalyticsConfiguration *CognitoUserPoolClientAnalyticsConfiguration `json:"analyticsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#callback_urls CognitoUserPoolClient#callback_urls}.
	CallbackUrls *[]*string `json:"callbackUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#default_redirect_uri CognitoUserPoolClient#default_redirect_uri}.
	DefaultRedirectUri *string `json:"defaultRedirectUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#enable_token_revocation CognitoUserPoolClient#enable_token_revocation}.
	EnableTokenRevocation interface{} `json:"enableTokenRevocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#explicit_auth_flows CognitoUserPoolClient#explicit_auth_flows}.
	ExplicitAuthFlows *[]*string `json:"explicitAuthFlows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#generate_secret CognitoUserPoolClient#generate_secret}.
	GenerateSecret interface{} `json:"generateSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#id_token_validity CognitoUserPoolClient#id_token_validity}.
	IdTokenValidity *float64 `json:"idTokenValidity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#logout_urls CognitoUserPoolClient#logout_urls}.
	LogoutUrls *[]*string `json:"logoutUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#prevent_user_existence_errors CognitoUserPoolClient#prevent_user_existence_errors}.
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#read_attributes CognitoUserPoolClient#read_attributes}.
	ReadAttributes *[]*string `json:"readAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#refresh_token_validity CognitoUserPoolClient#refresh_token_validity}.
	RefreshTokenValidity *float64 `json:"refreshTokenValidity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#supported_identity_providers CognitoUserPoolClient#supported_identity_providers}.
	SupportedIdentityProviders *[]*string `json:"supportedIdentityProviders"`
	// token_validity_units block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#token_validity_units CognitoUserPoolClient#token_validity_units}
	TokenValidityUnits *CognitoUserPoolClientTokenValidityUnits `json:"tokenValidityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#write_attributes CognitoUserPoolClient#write_attributes}.
	WriteAttributes *[]*string `json:"writeAttributes"`
}

type CognitoUserPoolClientTokenValidityUnits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#access_token CognitoUserPoolClient#access_token}.
	AccessToken *string `json:"accessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#id_token CognitoUserPoolClient#id_token}.
	IdToken *string `json:"idToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_client.html#refresh_token CognitoUserPoolClient#refresh_token}.
	RefreshToken *string `json:"refreshToken"`
}

type CognitoUserPoolClientTokenValidityUnitsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	IdToken() *string
	SetIdToken(val *string)
	IdTokenInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RefreshToken() *string
	SetRefreshToken(val *string)
	RefreshTokenInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAccessToken()
	ResetIdToken()
	ResetRefreshToken()
}

// The jsii proxy struct for CognitoUserPoolClientTokenValidityUnitsOutputReference
type jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) IdToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) IdTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) RefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolClientTokenValidityUnitsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolClientTokenValidityUnitsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClientTokenValidityUnitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolClientTokenValidityUnitsOutputReference_Override(c CognitoUserPoolClientTokenValidityUnitsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolClientTokenValidityUnitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetIdToken(val *string) {
	_jsii_.Set(
		j,
		"idToken",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetRefreshToken(val *string) {
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) ResetIdToken() {
	_jsii_.InvokeVoid(
		c,
		"resetIdToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshToken",
		nil, // no parameters
	)
}

type CognitoUserPoolConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#name CognitoUserPool#name}.
	Name *string `json:"name"`
	// account_recovery_setting block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#account_recovery_setting CognitoUserPool#account_recovery_setting}
	AccountRecoverySetting *CognitoUserPoolAccountRecoverySetting `json:"accountRecoverySetting"`
	// admin_create_user_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#admin_create_user_config CognitoUserPool#admin_create_user_config}
	AdminCreateUserConfig *CognitoUserPoolAdminCreateUserConfig `json:"adminCreateUserConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#alias_attributes CognitoUserPool#alias_attributes}.
	AliasAttributes *[]*string `json:"aliasAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#auto_verified_attributes CognitoUserPool#auto_verified_attributes}.
	AutoVerifiedAttributes *[]*string `json:"autoVerifiedAttributes"`
	// device_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#device_configuration CognitoUserPool#device_configuration}
	DeviceConfiguration *CognitoUserPoolDeviceConfiguration `json:"deviceConfiguration"`
	// email_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_configuration CognitoUserPool#email_configuration}
	EmailConfiguration *CognitoUserPoolEmailConfiguration `json:"emailConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_verification_message CognitoUserPool#email_verification_message}.
	EmailVerificationMessage *string `json:"emailVerificationMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_verification_subject CognitoUserPool#email_verification_subject}.
	EmailVerificationSubject *string `json:"emailVerificationSubject"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#lambda_config CognitoUserPool#lambda_config}
	LambdaConfig *CognitoUserPoolLambdaConfig `json:"lambdaConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#mfa_configuration CognitoUserPool#mfa_configuration}.
	MfaConfiguration *string `json:"mfaConfiguration"`
	// password_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#password_policy CognitoUserPool#password_policy}
	PasswordPolicy *CognitoUserPoolPasswordPolicy `json:"passwordPolicy"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#schema CognitoUserPool#schema}
	Schema *[]*CognitoUserPoolSchema `json:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sms_authentication_message CognitoUserPool#sms_authentication_message}.
	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage"`
	// sms_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sms_configuration CognitoUserPool#sms_configuration}
	SmsConfiguration *CognitoUserPoolSmsConfiguration `json:"smsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sms_verification_message CognitoUserPool#sms_verification_message}.
	SmsVerificationMessage *string `json:"smsVerificationMessage"`
	// software_token_mfa_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#software_token_mfa_configuration CognitoUserPool#software_token_mfa_configuration}
	SoftwareTokenMfaConfiguration *CognitoUserPoolSoftwareTokenMfaConfiguration `json:"softwareTokenMfaConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#tags CognitoUserPool#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#tags_all CognitoUserPool#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#username_attributes CognitoUserPool#username_attributes}.
	UsernameAttributes *[]*string `json:"usernameAttributes"`
	// username_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#username_configuration CognitoUserPool#username_configuration}
	UsernameConfiguration *CognitoUserPoolUsernameConfiguration `json:"usernameConfiguration"`
	// user_pool_add_ons block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#user_pool_add_ons CognitoUserPool#user_pool_add_ons}
	UserPoolAddOns *CognitoUserPoolUserPoolAddOns `json:"userPoolAddOns"`
	// verification_message_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#verification_message_template CognitoUserPool#verification_message_template}
	VerificationMessageTemplate *CognitoUserPoolVerificationMessageTemplate `json:"verificationMessageTemplate"`
}

type CognitoUserPoolDeviceConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#challenge_required_on_new_device CognitoUserPool#challenge_required_on_new_device}.
	ChallengeRequiredOnNewDevice interface{} `json:"challengeRequiredOnNewDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#device_only_remembered_on_user_prompt CognitoUserPool#device_only_remembered_on_user_prompt}.
	DeviceOnlyRememberedOnUserPrompt interface{} `json:"deviceOnlyRememberedOnUserPrompt"`
}

type CognitoUserPoolDeviceConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChallengeRequiredOnNewDevice() interface{}
	SetChallengeRequiredOnNewDevice(val interface{})
	ChallengeRequiredOnNewDeviceInput() interface{}
	DeviceOnlyRememberedOnUserPrompt() interface{}
	SetDeviceOnlyRememberedOnUserPrompt(val interface{})
	DeviceOnlyRememberedOnUserPromptInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetChallengeRequiredOnNewDevice()
	ResetDeviceOnlyRememberedOnUserPrompt()
}

// The jsii proxy struct for CognitoUserPoolDeviceConfigurationOutputReference
type jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ChallengeRequiredOnNewDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"challengeRequiredOnNewDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ChallengeRequiredOnNewDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"challengeRequiredOnNewDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) DeviceOnlyRememberedOnUserPrompt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceOnlyRememberedOnUserPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) DeviceOnlyRememberedOnUserPromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceOnlyRememberedOnUserPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolDeviceConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolDeviceConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolDeviceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolDeviceConfigurationOutputReference_Override(c CognitoUserPoolDeviceConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolDeviceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetChallengeRequiredOnNewDevice(val interface{}) {
	_jsii_.Set(
		j,
		"challengeRequiredOnNewDevice",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetDeviceOnlyRememberedOnUserPrompt(val interface{}) {
	_jsii_.Set(
		j,
		"deviceOnlyRememberedOnUserPrompt",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ResetChallengeRequiredOnNewDevice() {
	_jsii_.InvokeVoid(
		c,
		"resetChallengeRequiredOnNewDevice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ResetDeviceOnlyRememberedOnUserPrompt() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceOnlyRememberedOnUserPrompt",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html aws_cognito_user_pool_domain}.
type CognitoUserPoolDomain interface {
	cdktf.TerraformResource
	AwsAccountId() *string
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	CloudfrontDistributionArn() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3Bucket() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	Version() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetCertificateArn()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPoolDomain
type jsiiProxy_CognitoUserPoolDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPoolDomain) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) CloudfrontDistributionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontDistributionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDomain) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html aws_cognito_user_pool_domain} Resource.
func NewCognitoUserPoolDomain(scope constructs.Construct, id *string, config *CognitoUserPoolDomainConfig) CognitoUserPoolDomain {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolDomain{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html aws_cognito_user_pool_domain} Resource.
func NewCognitoUserPoolDomain_Override(c CognitoUserPoolDomain, scope constructs.Construct, id *string, config *CognitoUserPoolDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolDomain",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDomain) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoUserPoolDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoUserPoolDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPoolDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoUserPoolDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPoolDomain) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificateArn",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoUserPoolDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoUserPoolDomainConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html#domain CognitoUserPoolDomain#domain}.
	Domain *string `json:"domain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html#user_pool_id CognitoUserPoolDomain#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_domain.html#certificate_arn CognitoUserPoolDomain#certificate_arn}.
	CertificateArn *string `json:"certificateArn"`
}

type CognitoUserPoolEmailConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#configuration_set CognitoUserPool#configuration_set}.
	ConfigurationSet *string `json:"configurationSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_sending_account CognitoUserPool#email_sending_account}.
	EmailSendingAccount *string `json:"emailSendingAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#from_email_address CognitoUserPool#from_email_address}.
	FromEmailAddress *string `json:"fromEmailAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#reply_to_email_address CognitoUserPool#reply_to_email_address}.
	ReplyToEmailAddress *string `json:"replyToEmailAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#source_arn CognitoUserPool#source_arn}.
	SourceArn *string `json:"sourceArn"`
}

type CognitoUserPoolEmailConfigurationOutputReference interface {
	cdktf.ComplexObject
	ConfigurationSet() *string
	SetConfigurationSet(val *string)
	ConfigurationSetInput() *string
	EmailSendingAccount() *string
	SetEmailSendingAccount(val *string)
	EmailSendingAccountInput() *string
	FromEmailAddress() *string
	SetFromEmailAddress(val *string)
	FromEmailAddressInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ReplyToEmailAddress() *string
	SetReplyToEmailAddress(val *string)
	ReplyToEmailAddressInput() *string
	SourceArn() *string
	SetSourceArn(val *string)
	SourceArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetConfigurationSet()
	ResetEmailSendingAccount()
	ResetFromEmailAddress()
	ResetReplyToEmailAddress()
	ResetSourceArn()
}

// The jsii proxy struct for CognitoUserPoolEmailConfigurationOutputReference
type jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ConfigurationSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) EmailSendingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) EmailSendingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) FromEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) FromEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ReplyToEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ReplyToEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolEmailConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolEmailConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolEmailConfigurationOutputReference_Override(c CognitoUserPoolEmailConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetConfigurationSet(val *string) {
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetEmailSendingAccount(val *string) {
	_jsii_.Set(
		j,
		"emailSendingAccount",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetFromEmailAddress(val *string) {
	_jsii_.Set(
		j,
		"fromEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetReplyToEmailAddress(val *string) {
	_jsii_.Set(
		j,
		"replyToEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetSourceArn(val *string) {
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetConfigurationSet() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigurationSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetEmailSendingAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSendingAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetFromEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetFromEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetReplyToEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetReplyToEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference) ResetSourceArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArn",
		nil, // no parameters
	)
}

type CognitoUserPoolLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#create_auth_challenge CognitoUserPool#create_auth_challenge}.
	CreateAuthChallenge *string `json:"createAuthChallenge"`
	// custom_email_sender block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#custom_email_sender CognitoUserPool#custom_email_sender}
	CustomEmailSender *CognitoUserPoolLambdaConfigCustomEmailSender `json:"customEmailSender"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#custom_message CognitoUserPool#custom_message}.
	CustomMessage *string `json:"customMessage"`
	// custom_sms_sender block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#custom_sms_sender CognitoUserPool#custom_sms_sender}
	CustomSmsSender *CognitoUserPoolLambdaConfigCustomSmsSender `json:"customSmsSender"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#define_auth_challenge CognitoUserPool#define_auth_challenge}.
	DefineAuthChallenge *string `json:"defineAuthChallenge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#kms_key_id CognitoUserPool#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#post_authentication CognitoUserPool#post_authentication}.
	PostAuthentication *string `json:"postAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#post_confirmation CognitoUserPool#post_confirmation}.
	PostConfirmation *string `json:"postConfirmation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#pre_authentication CognitoUserPool#pre_authentication}.
	PreAuthentication *string `json:"preAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#pre_sign_up CognitoUserPool#pre_sign_up}.
	PreSignUp *string `json:"preSignUp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#pre_token_generation CognitoUserPool#pre_token_generation}.
	PreTokenGeneration *string `json:"preTokenGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#user_migration CognitoUserPool#user_migration}.
	UserMigration *string `json:"userMigration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#verify_auth_challenge_response CognitoUserPool#verify_auth_challenge_response}.
	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse"`
}

type CognitoUserPoolLambdaConfigCustomEmailSender struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#lambda_arn CognitoUserPool#lambda_arn}.
	LambdaArn *string `json:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#lambda_version CognitoUserPool#lambda_version}.
	LambdaVersion *string `json:"lambdaVersion"`
}

type CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaArn() *string
	SetLambdaArn(val *string)
	LambdaArnInput() *string
	LambdaVersion() *string
	SetLambdaVersion(val *string)
	LambdaVersionInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference
type jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) LambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) LambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) LambdaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) LambdaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolLambdaConfigCustomEmailSenderOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolLambdaConfigCustomEmailSenderOutputReference_Override(c CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) SetLambdaArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) SetLambdaVersion(val *string) {
	_jsii_.Set(
		j,
		"lambdaVersion",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolLambdaConfigCustomSmsSender struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#lambda_arn CognitoUserPool#lambda_arn}.
	LambdaArn *string `json:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#lambda_version CognitoUserPool#lambda_version}.
	LambdaVersion *string `json:"lambdaVersion"`
}

type CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaArn() *string
	SetLambdaArn(val *string)
	LambdaArnInput() *string
	LambdaVersion() *string
	SetLambdaVersion(val *string)
	LambdaVersionInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference
type jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) LambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) LambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) LambdaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) LambdaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolLambdaConfigCustomSmsSenderOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolLambdaConfigCustomSmsSenderOutputReference_Override(c CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) SetLambdaArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) SetLambdaVersion(val *string) {
	_jsii_.Set(
		j,
		"lambdaVersion",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolLambdaConfigOutputReference interface {
	cdktf.ComplexObject
	CreateAuthChallenge() *string
	SetCreateAuthChallenge(val *string)
	CreateAuthChallengeInput() *string
	CustomEmailSender() CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference
	CustomEmailSenderInput() *CognitoUserPoolLambdaConfigCustomEmailSender
	CustomMessage() *string
	SetCustomMessage(val *string)
	CustomMessageInput() *string
	CustomSmsSender() CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference
	CustomSmsSenderInput() *CognitoUserPoolLambdaConfigCustomSmsSender
	DefineAuthChallenge() *string
	SetDefineAuthChallenge(val *string)
	DefineAuthChallengeInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	PostAuthentication() *string
	SetPostAuthentication(val *string)
	PostAuthenticationInput() *string
	PostConfirmation() *string
	SetPostConfirmation(val *string)
	PostConfirmationInput() *string
	PreAuthentication() *string
	SetPreAuthentication(val *string)
	PreAuthenticationInput() *string
	PreSignUp() *string
	SetPreSignUp(val *string)
	PreSignUpInput() *string
	PreTokenGeneration() *string
	SetPreTokenGeneration(val *string)
	PreTokenGenerationInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserMigration() *string
	SetUserMigration(val *string)
	UserMigrationInput() *string
	VerifyAuthChallengeResponse() *string
	SetVerifyAuthChallengeResponse(val *string)
	VerifyAuthChallengeResponseInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomEmailSender(value *CognitoUserPoolLambdaConfigCustomEmailSender)
	PutCustomSmsSender(value *CognitoUserPoolLambdaConfigCustomSmsSender)
	ResetCreateAuthChallenge()
	ResetCustomEmailSender()
	ResetCustomMessage()
	ResetCustomSmsSender()
	ResetDefineAuthChallenge()
	ResetKmsKeyId()
	ResetPostAuthentication()
	ResetPostConfirmation()
	ResetPreAuthentication()
	ResetPreSignUp()
	ResetPreTokenGeneration()
	ResetUserMigration()
	ResetVerifyAuthChallengeResponse()
}

// The jsii proxy struct for CognitoUserPoolLambdaConfigOutputReference
type jsiiProxy_CognitoUserPoolLambdaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CreateAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomEmailSender() CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference {
	var returns CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomEmailSenderInput() *CognitoUserPoolLambdaConfigCustomEmailSender {
	var returns *CognitoUserPoolLambdaConfigCustomEmailSender
	_jsii_.Get(
		j,
		"customEmailSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomSmsSender() CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference {
	var returns CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) CustomSmsSenderInput() *CognitoUserPoolLambdaConfigCustomSmsSender {
	var returns *CognitoUserPoolLambdaConfigCustomSmsSender
	_jsii_.Get(
		j,
		"customSmsSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) DefineAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PostConfirmationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreSignUpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PreTokenGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) UserMigrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) VerifyAuthChallengeResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponseInput",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolLambdaConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolLambdaConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolLambdaConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolLambdaConfigOutputReference_Override(c CognitoUserPoolLambdaConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetCreateAuthChallenge(val *string) {
	_jsii_.Set(
		j,
		"createAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetCustomMessage(val *string) {
	_jsii_.Set(
		j,
		"customMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetDefineAuthChallenge(val *string) {
	_jsii_.Set(
		j,
		"defineAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetPostAuthentication(val *string) {
	_jsii_.Set(
		j,
		"postAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetPostConfirmation(val *string) {
	_jsii_.Set(
		j,
		"postConfirmation",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetPreAuthentication(val *string) {
	_jsii_.Set(
		j,
		"preAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetPreSignUp(val *string) {
	_jsii_.Set(
		j,
		"preSignUp",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetPreTokenGeneration(val *string) {
	_jsii_.Set(
		j,
		"preTokenGeneration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetUserMigration(val *string) {
	_jsii_.Set(
		j,
		"userMigration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) SetVerifyAuthChallengeResponse(val *string) {
	_jsii_.Set(
		j,
		"verifyAuthChallengeResponse",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PutCustomEmailSender(value *CognitoUserPoolLambdaConfigCustomEmailSender) {
	_jsii_.InvokeVoid(
		c,
		"putCustomEmailSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) PutCustomSmsSender(value *CognitoUserPoolLambdaConfigCustomSmsSender) {
	_jsii_.InvokeVoid(
		c,
		"putCustomSmsSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCreateAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetCreateAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomEmailSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomEmailSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetCustomSmsSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomSmsSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetDefineAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetDefineAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPostAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPostAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPostConfirmation() {
	_jsii_.InvokeVoid(
		c,
		"resetPostConfirmation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPreAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreSignUp() {
	_jsii_.InvokeVoid(
		c,
		"resetPreSignUp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetPreTokenGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetPreTokenGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetUserMigration() {
	_jsii_.InvokeVoid(
		c,
		"resetUserMigration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolLambdaConfigOutputReference) ResetVerifyAuthChallengeResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetVerifyAuthChallengeResponse",
		nil, // no parameters
	)
}

type CognitoUserPoolPasswordPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#minimum_length CognitoUserPool#minimum_length}.
	MinimumLength *float64 `json:"minimumLength"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#require_lowercase CognitoUserPool#require_lowercase}.
	RequireLowercase interface{} `json:"requireLowercase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#require_numbers CognitoUserPool#require_numbers}.
	RequireNumbers interface{} `json:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#require_symbols CognitoUserPool#require_symbols}.
	RequireSymbols interface{} `json:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#require_uppercase CognitoUserPool#require_uppercase}.
	RequireUppercase interface{} `json:"requireUppercase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#temporary_password_validity_days CognitoUserPool#temporary_password_validity_days}.
	TemporaryPasswordValidityDays *float64 `json:"temporaryPasswordValidityDays"`
}

type CognitoUserPoolPasswordPolicyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MinimumLength() *float64
	SetMinimumLength(val *float64)
	MinimumLengthInput() *float64
	RequireLowercase() interface{}
	SetRequireLowercase(val interface{})
	RequireLowercaseInput() interface{}
	RequireNumbers() interface{}
	SetRequireNumbers(val interface{})
	RequireNumbersInput() interface{}
	RequireSymbols() interface{}
	SetRequireSymbols(val interface{})
	RequireSymbolsInput() interface{}
	RequireUppercase() interface{}
	SetRequireUppercase(val interface{})
	RequireUppercaseInput() interface{}
	TemporaryPasswordValidityDays() *float64
	SetTemporaryPasswordValidityDays(val *float64)
	TemporaryPasswordValidityDaysInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMinimumLength()
	ResetRequireLowercase()
	ResetRequireNumbers()
	ResetRequireSymbols()
	ResetRequireUppercase()
	ResetTemporaryPasswordValidityDays()
}

// The jsii proxy struct for CognitoUserPoolPasswordPolicyOutputReference
type jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) MinimumLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) MinimumLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) RequireUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TemporaryPasswordValidityDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TemporaryPasswordValidityDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolPasswordPolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolPasswordPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolPasswordPolicyOutputReference_Override(c CognitoUserPoolPasswordPolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetMinimumLength(val *float64) {
	_jsii_.Set(
		j,
		"minimumLength",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetRequireLowercase(val interface{}) {
	_jsii_.Set(
		j,
		"requireLowercase",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetRequireNumbers(val interface{}) {
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetRequireSymbols(val interface{}) {
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetRequireUppercase(val interface{}) {
	_jsii_.Set(
		j,
		"requireUppercase",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetTemporaryPasswordValidityDays(val *float64) {
	_jsii_.Set(
		j,
		"temporaryPasswordValidityDays",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetMinimumLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireLowercase() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireLowercase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetRequireUppercase() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireUppercase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference) ResetTemporaryPasswordValidityDays() {
	_jsii_.InvokeVoid(
		c,
		"resetTemporaryPasswordValidityDays",
		nil, // no parameters
	)
}

type CognitoUserPoolSchema struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#attribute_data_type CognitoUserPool#attribute_data_type}.
	AttributeDataType *string `json:"attributeDataType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#name CognitoUserPool#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#developer_only_attribute CognitoUserPool#developer_only_attribute}.
	DeveloperOnlyAttribute interface{} `json:"developerOnlyAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#mutable CognitoUserPool#mutable}.
	Mutable interface{} `json:"mutable"`
	// number_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#number_attribute_constraints CognitoUserPool#number_attribute_constraints}
	NumberAttributeConstraints *CognitoUserPoolSchemaNumberAttributeConstraints `json:"numberAttributeConstraints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#required CognitoUserPool#required}.
	Required interface{} `json:"required"`
	// string_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#string_attribute_constraints CognitoUserPool#string_attribute_constraints}
	StringAttributeConstraints *CognitoUserPoolSchemaStringAttributeConstraints `json:"stringAttributeConstraints"`
}

type CognitoUserPoolSchemaNumberAttributeConstraints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#max_value CognitoUserPool#max_value}.
	MaxValue *string `json:"maxValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#min_value CognitoUserPool#min_value}.
	MinValue *string `json:"minValue"`
}

type CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxValue() *string
	SetMaxValue(val *string)
	MaxValueInput() *string
	MinValue() *string
	SetMinValue(val *string)
	MinValueInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxValue()
	ResetMinValue()
}

// The jsii proxy struct for CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference
type jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) MaxValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) MaxValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) MinValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) MinValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolSchemaNumberAttributeConstraintsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolSchemaNumberAttributeConstraintsOutputReference_Override(c CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) SetMaxValue(val *string) {
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) SetMinValue(val *string) {
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) ResetMaxValue() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference) ResetMinValue() {
	_jsii_.InvokeVoid(
		c,
		"resetMinValue",
		nil, // no parameters
	)
}

type CognitoUserPoolSchemaStringAttributeConstraints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#max_length CognitoUserPool#max_length}.
	MaxLength *string `json:"maxLength"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#min_length CognitoUserPool#min_length}.
	MinLength *string `json:"minLength"`
}

type CognitoUserPoolSchemaStringAttributeConstraintsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxLength() *string
	SetMaxLength(val *string)
	MaxLengthInput() *string
	MinLength() *string
	SetMinLength(val *string)
	MinLengthInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxLength()
	ResetMinLength()
}

// The jsii proxy struct for CognitoUserPoolSchemaStringAttributeConstraintsOutputReference
type jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) MaxLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) MaxLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) MinLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) MinLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolSchemaStringAttributeConstraintsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolSchemaStringAttributeConstraintsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaStringAttributeConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolSchemaStringAttributeConstraintsOutputReference_Override(c CognitoUserPoolSchemaStringAttributeConstraintsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaStringAttributeConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) SetMaxLength(val *string) {
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) SetMinLength(val *string) {
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) ResetMaxLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference) ResetMinLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMinLength",
		nil, // no parameters
	)
}

type CognitoUserPoolSmsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#external_id CognitoUserPool#external_id}.
	ExternalId *string `json:"externalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sns_caller_arn CognitoUserPool#sns_caller_arn}.
	SnsCallerArn *string `json:"snsCallerArn"`
}

type CognitoUserPoolSmsConfigurationOutputReference interface {
	cdktf.ComplexObject
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SnsCallerArn() *string
	SetSnsCallerArn(val *string)
	SnsCallerArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolSmsConfigurationOutputReference
type jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SnsCallerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsCallerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SnsCallerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsCallerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolSmsConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolSmsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSmsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolSmsConfigurationOutputReference_Override(c CognitoUserPoolSmsConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSmsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SetExternalId(val *string) {
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SetSnsCallerArn(val *string) {
	_jsii_.Set(
		j,
		"snsCallerArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolSoftwareTokenMfaConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#enabled CognitoUserPool#enabled}.
	Enabled interface{} `json:"enabled"`
}

type CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference
type jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolSoftwareTokenMfaConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolSoftwareTokenMfaConfigurationOutputReference_Override(c CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html aws_cognito_user_pool_ui_customization}.
type CognitoUserPoolUiCustomization interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreationDate() *string
	Css() *string
	SetCss(val *string)
	CssInput() *string
	CssVersion() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	ImageFile() *string
	SetImageFile(val *string)
	ImageFileInput() *string
	ImageUrl() *string
	LastModifiedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetClientId()
	ResetCss()
	ResetImageFile()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CognitoUserPoolUiCustomization
type jsiiProxy_CognitoUserPoolUiCustomization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Css() *string {
	var returns *string
	_jsii_.Get(
		j,
		"css",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) CssInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) CssVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cssVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ImageFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ImageFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html aws_cognito_user_pool_ui_customization} Resource.
func NewCognitoUserPoolUiCustomization(scope constructs.Construct, id *string, config *CognitoUserPoolUiCustomizationConfig) CognitoUserPoolUiCustomization {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolUiCustomization{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html aws_cognito_user_pool_ui_customization} Resource.
func NewCognitoUserPoolUiCustomization_Override(c CognitoUserPoolUiCustomization, scope constructs.Construct, id *string, config *CognitoUserPoolUiCustomizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomization",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetCss(val *string) {
	_jsii_.Set(
		j,
		"css",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetImageFile(val *string) {
	_jsii_.Set(
		j,
		"imageFile",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUiCustomization) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoUserPoolUiCustomization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUserPoolUiCustomization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomization",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUserPoolUiCustomization) ResetClientId() {
	_jsii_.InvokeVoid(
		c,
		"resetClientId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolUiCustomization) ResetCss() {
	_jsii_.InvokeVoid(
		c,
		"resetCss",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolUiCustomization) ResetImageFile() {
	_jsii_.InvokeVoid(
		c,
		"resetImageFile",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolUiCustomization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CognitoUserPoolUiCustomization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoUserPoolUiCustomizationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html#user_pool_id CognitoUserPoolUiCustomization#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html#client_id CognitoUserPoolUiCustomization#client_id}.
	ClientId *string `json:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html#css CognitoUserPoolUiCustomization#css}.
	Css *string `json:"css"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool_ui_customization.html#image_file CognitoUserPoolUiCustomization#image_file}.
	ImageFile *string `json:"imageFile"`
}

type CognitoUserPoolUserPoolAddOns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#advanced_security_mode CognitoUserPool#advanced_security_mode}.
	AdvancedSecurityMode *string `json:"advancedSecurityMode"`
}

type CognitoUserPoolUserPoolAddOnsOutputReference interface {
	cdktf.ComplexObject
	AdvancedSecurityMode() *string
	SetAdvancedSecurityMode(val *string)
	AdvancedSecurityModeInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolUserPoolAddOnsOutputReference
type jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) AdvancedSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advancedSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) AdvancedSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advancedSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolUserPoolAddOnsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolUserPoolAddOnsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUserPoolAddOnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolUserPoolAddOnsOutputReference_Override(c CognitoUserPoolUserPoolAddOnsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUserPoolAddOnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) SetAdvancedSecurityMode(val *string) {
	_jsii_.Set(
		j,
		"advancedSecurityMode",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolUsernameConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#case_sensitive CognitoUserPool#case_sensitive}.
	CaseSensitive interface{} `json:"caseSensitive"`
}

type CognitoUserPoolUsernameConfigurationOutputReference interface {
	cdktf.ComplexObject
	CaseSensitive() interface{}
	SetCaseSensitive(val interface{})
	CaseSensitiveInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for CognitoUserPoolUsernameConfigurationOutputReference
type jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) CaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) CaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolUsernameConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolUsernameConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUsernameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolUsernameConfigurationOutputReference_Override(c CognitoUserPoolUsernameConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolUsernameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) SetCaseSensitive(val interface{}) {
	_jsii_.Set(
		j,
		"caseSensitive",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CognitoUserPoolVerificationMessageTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#default_email_option CognitoUserPool#default_email_option}.
	DefaultEmailOption *string `json:"defaultEmailOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_message CognitoUserPool#email_message}.
	EmailMessage *string `json:"emailMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_message_by_link CognitoUserPool#email_message_by_link}.
	EmailMessageByLink *string `json:"emailMessageByLink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_subject CognitoUserPool#email_subject}.
	EmailSubject *string `json:"emailSubject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#email_subject_by_link CognitoUserPool#email_subject_by_link}.
	EmailSubjectByLink *string `json:"emailSubjectByLink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html#sms_message CognitoUserPool#sms_message}.
	SmsMessage *string `json:"smsMessage"`
}

type CognitoUserPoolVerificationMessageTemplateOutputReference interface {
	cdktf.ComplexObject
	DefaultEmailOption() *string
	SetDefaultEmailOption(val *string)
	DefaultEmailOptionInput() *string
	EmailMessage() *string
	SetEmailMessage(val *string)
	EmailMessageByLink() *string
	SetEmailMessageByLink(val *string)
	EmailMessageByLinkInput() *string
	EmailMessageInput() *string
	EmailSubject() *string
	SetEmailSubject(val *string)
	EmailSubjectByLink() *string
	SetEmailSubjectByLink(val *string)
	EmailSubjectByLinkInput() *string
	EmailSubjectInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SmsMessage() *string
	SetSmsMessage(val *string)
	SmsMessageInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDefaultEmailOption()
	ResetEmailMessage()
	ResetEmailMessageByLink()
	ResetEmailSubject()
	ResetEmailSubjectByLink()
	ResetSmsMessage()
}

// The jsii proxy struct for CognitoUserPoolVerificationMessageTemplateOutputReference
type jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) DefaultEmailOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) DefaultEmailOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SmsMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SmsMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolVerificationMessageTemplateOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CognitoUserPoolVerificationMessageTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolVerificationMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCognitoUserPoolVerificationMessageTemplateOutputReference_Override(c CognitoUserPoolVerificationMessageTemplateOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.CognitoUserPoolVerificationMessageTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetDefaultEmailOption(val *string) {
	_jsii_.Set(
		j,
		"defaultEmailOption",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetEmailMessage(val *string) {
	_jsii_.Set(
		j,
		"emailMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetEmailMessageByLink(val *string) {
	_jsii_.Set(
		j,
		"emailMessageByLink",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetEmailSubject(val *string) {
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetEmailSubjectByLink(val *string) {
	_jsii_.Set(
		j,
		"emailSubjectByLink",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetSmsMessage(val *string) {
	_jsii_.Set(
		j,
		"smsMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetDefaultEmailOption() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultEmailOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailMessageByLink() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailMessageByLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetEmailSubjectByLink() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSubjectByLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference) ResetSmsMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetSmsMessage",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cognito_user_pools.html aws_cognito_user_pools}.
type DataAwsCognitoUserPools interface {
	cdktf.TerraformDataSource
	Arns() *[]*string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCognitoUserPools
type jsiiProxy_DataAwsCognitoUserPools struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Arns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPools) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cognito_user_pools.html aws_cognito_user_pools} Data Source.
func NewDataAwsCognitoUserPools(scope constructs.Construct, id *string, config *DataAwsCognitoUserPoolsConfig) DataAwsCognitoUserPools {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCognitoUserPools{}

	_jsii_.Create(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPools",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cognito_user_pools.html aws_cognito_user_pools} Data Source.
func NewDataAwsCognitoUserPools_Override(d DataAwsCognitoUserPools, scope constructs.Construct, id *string, config *DataAwsCognitoUserPoolsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPools",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPools) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPools) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPools) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPools) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPools) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsCognitoUserPools_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPools",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCognitoUserPools_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPools",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCognitoUserPools) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCognitoUserPools) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCognitoUserPools) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCognitoUserPools) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) ToString() *string {
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
func (d *jsiiProxy_DataAwsCognitoUserPools) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCognitoUserPoolsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cognito_user_pools.html#name DataAwsCognitoUserPools#name}.
	Name *string `json:"name"`
}
