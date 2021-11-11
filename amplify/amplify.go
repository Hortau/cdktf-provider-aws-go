package amplify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/amplify/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html aws_amplify_app}.
type AmplifyApp interface {
	cdktf.TerraformResource
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Arn() *string
	AutoBranchCreationConfig() AmplifyAppAutoBranchCreationConfigOutputReference
	AutoBranchCreationConfigInput() *AmplifyAppAutoBranchCreationConfig
	AutoBranchCreationPatterns() *[]*string
	SetAutoBranchCreationPatterns(val *[]*string)
	AutoBranchCreationPatternsInput() *[]*string
	BasicAuthCredentials() *string
	SetBasicAuthCredentials(val *string)
	BasicAuthCredentialsInput() *string
	BuildSpec() *string
	SetBuildSpec(val *string)
	BuildSpecInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomRule() *[]*AmplifyAppCustomRule
	SetCustomRule(val *[]*AmplifyAppCustomRule)
	CustomRuleInput() *[]*AmplifyAppCustomRule
	DefaultDomain() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableAutoBranchCreation() interface{}
	SetEnableAutoBranchCreation(val interface{})
	EnableAutoBranchCreationInput() interface{}
	EnableBasicAuth() interface{}
	SetEnableBasicAuth(val interface{})
	EnableBasicAuthInput() interface{}
	EnableBranchAutoBuild() interface{}
	SetEnableBranchAutoBuild(val interface{})
	EnableBranchAutoBuildInput() interface{}
	EnableBranchAutoDeletion() interface{}
	SetEnableBranchAutoDeletion(val interface{})
	EnableBranchAutoDeletionInput() interface{}
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	EnvironmentVariablesInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	IamServiceRoleArn() *string
	SetIamServiceRoleArn(val *string)
	IamServiceRoleArnInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OauthToken() *string
	SetOauthToken(val *string)
	OauthTokenInput() *string
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
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
	ProductionBranch(index *string) AmplifyAppProductionBranch
	PutAutoBranchCreationConfig(value *AmplifyAppAutoBranchCreationConfig)
	ResetAccessToken()
	ResetAutoBranchCreationConfig()
	ResetAutoBranchCreationPatterns()
	ResetBasicAuthCredentials()
	ResetBuildSpec()
	ResetCustomRule()
	ResetDescription()
	ResetEnableAutoBranchCreation()
	ResetEnableBasicAuth()
	ResetEnableBranchAutoBuild()
	ResetEnableBranchAutoDeletion()
	ResetEnvironmentVariables()
	ResetIamServiceRoleArn()
	ResetOauthToken()
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRepository()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AmplifyApp
type jsiiProxy_AmplifyApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyApp) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationConfig() AmplifyAppAutoBranchCreationConfigOutputReference {
	var returns AmplifyAppAutoBranchCreationConfigOutputReference
	_jsii_.Get(
		j,
		"autoBranchCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationConfigInput() *AmplifyAppAutoBranchCreationConfig {
	var returns *AmplifyAppAutoBranchCreationConfig
	_jsii_.Get(
		j,
		"autoBranchCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomRule() *[]*AmplifyAppCustomRule {
	var returns *[]*AmplifyAppCustomRule
	_jsii_.Get(
		j,
		"customRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomRuleInput() *[]*AmplifyAppCustomRule {
	var returns *[]*AmplifyAppCustomRule
	_jsii_.Get(
		j,
		"customRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableAutoBranchCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableAutoBranchCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) IamServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) IamServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) OauthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html aws_amplify_app} Resource.
func NewAmplifyApp(scope constructs.Construct, id *string, config *AmplifyAppConfig) AmplifyApp {
	_init_.Initialize()

	j := jsiiProxy_AmplifyApp{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html aws_amplify_app} Resource.
func NewAmplifyApp_Override(a AmplifyApp, scope constructs.Construct, id *string, config *AmplifyAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyApp",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyApp) SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetAutoBranchCreationPatterns(val *[]*string) {
	_jsii_.Set(
		j,
		"autoBranchCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetBasicAuthCredentials(val *string) {
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetBuildSpec(val *string) {
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetCustomRule(val *[]*AmplifyAppCustomRule) {
	_jsii_.Set(
		j,
		"customRule",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetEnableAutoBranchCreation(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoBranchCreation",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetEnableBasicAuth(val interface{}) {
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetEnableBranchAutoBuild(val interface{}) {
	_jsii_.Set(
		j,
		"enableBranchAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetEnableBranchAutoDeletion(val interface{}) {
	_jsii_.Set(
		j,
		"enableBranchAutoDeletion",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetIamServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamServiceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetOauthToken(val *string) {
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetPlatform(val *string) {
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp) SetTagsAll(val interface{}) {
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
func AmplifyApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Amplify.AmplifyApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Amplify.AmplifyApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyApp) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyApp) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyApp) ProductionBranch(index *string) AmplifyAppProductionBranch {
	var returns AmplifyAppProductionBranch

	_jsii_.Invoke(
		a,
		"productionBranch",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) PutAutoBranchCreationConfig(value *AmplifyAppAutoBranchCreationConfig) {
	_jsii_.InvokeVoid(
		a,
		"putAutoBranchCreationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAutoBranchCreationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAutoBranchCreationPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetCustomRule() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableAutoBranchCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBranchCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBranchAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBranchAutoDeletion() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoDeletion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetIamServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetOauthToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthToken",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AmplifyApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetPlatform() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatform",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AmplifyApp) ToMetadata() interface{} {
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
func (a *jsiiProxy_AmplifyApp) ToString() *string {
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
func (a *jsiiProxy_AmplifyApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AmplifyAppAutoBranchCreationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#basic_auth_credentials AmplifyApp#basic_auth_credentials}.
	BasicAuthCredentials *string `json:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#build_spec AmplifyApp#build_spec}.
	BuildSpec *string `json:"buildSpec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_auto_build AmplifyApp#enable_auto_build}.
	EnableAutoBuild interface{} `json:"enableAutoBuild"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_basic_auth AmplifyApp#enable_basic_auth}.
	EnableBasicAuth interface{} `json:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_performance_mode AmplifyApp#enable_performance_mode}.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_pull_request_preview AmplifyApp#enable_pull_request_preview}.
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#environment_variables AmplifyApp#environment_variables}.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#framework AmplifyApp#framework}.
	Framework *string `json:"framework"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#pull_request_environment_name AmplifyApp#pull_request_environment_name}.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#stage AmplifyApp#stage}.
	Stage *string `json:"stage"`
}

type AmplifyAppAutoBranchCreationConfigOutputReference interface {
	cdktf.ComplexObject
	BasicAuthCredentials() *string
	SetBasicAuthCredentials(val *string)
	BasicAuthCredentialsInput() *string
	BuildSpec() *string
	SetBuildSpec(val *string)
	BuildSpecInput() *string
	EnableAutoBuild() interface{}
	SetEnableAutoBuild(val interface{})
	EnableAutoBuildInput() interface{}
	EnableBasicAuth() interface{}
	SetEnableBasicAuth(val interface{})
	EnableBasicAuthInput() interface{}
	EnablePerformanceMode() interface{}
	SetEnablePerformanceMode(val interface{})
	EnablePerformanceModeInput() interface{}
	EnablePullRequestPreview() interface{}
	SetEnablePullRequestPreview(val interface{})
	EnablePullRequestPreviewInput() interface{}
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	EnvironmentVariablesInput() interface{}
	Framework() *string
	SetFramework(val *string)
	FrameworkInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PullRequestEnvironmentName() *string
	SetPullRequestEnvironmentName(val *string)
	PullRequestEnvironmentNameInput() *string
	Stage() *string
	SetStage(val *string)
	StageInput() *string
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
	ResetBasicAuthCredentials()
	ResetBuildSpec()
	ResetEnableAutoBuild()
	ResetEnableBasicAuth()
	ResetEnablePerformanceMode()
	ResetEnablePullRequestPreview()
	ResetEnvironmentVariables()
	ResetFramework()
	ResetPullRequestEnvironmentName()
	ResetStage()
}

// The jsii proxy struct for AmplifyAppAutoBranchCreationConfigOutputReference
type jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnableAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnableAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnablePerformanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnablePerformanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnablePullRequestPreview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnablePullRequestPreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) EnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) PullRequestEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAmplifyAppAutoBranchCreationConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AmplifyAppAutoBranchCreationConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyAppAutoBranchCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAmplifyAppAutoBranchCreationConfigOutputReference_Override(a AmplifyAppAutoBranchCreationConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyAppAutoBranchCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetBasicAuthCredentials(val *string) {
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetBuildSpec(val *string) {
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetEnableAutoBuild(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetEnableBasicAuth(val interface{}) {
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetEnablePerformanceMode(val interface{}) {
	_jsii_.Set(
		j,
		"enablePerformanceMode",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetEnablePullRequestPreview(val interface{}) {
	_jsii_.Set(
		j,
		"enablePullRequestPreview",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetFramework(val *string) {
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetPullRequestEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"pullRequestEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetEnableAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetEnablePerformanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePerformanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetEnablePullRequestPreview() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePullRequestPreview",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetFramework() {
	_jsii_.InvokeVoid(
		a,
		"resetFramework",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetPullRequestEnvironmentName() {
	_jsii_.InvokeVoid(
		a,
		"resetPullRequestEnvironmentName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyAppAutoBranchCreationConfigOutputReference) ResetStage() {
	_jsii_.InvokeVoid(
		a,
		"resetStage",
		nil, // no parameters
	)
}

type AmplifyAppConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#name AmplifyApp#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#access_token AmplifyApp#access_token}.
	AccessToken *string `json:"accessToken"`
	// auto_branch_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#auto_branch_creation_config AmplifyApp#auto_branch_creation_config}
	AutoBranchCreationConfig *AmplifyAppAutoBranchCreationConfig `json:"autoBranchCreationConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#auto_branch_creation_patterns AmplifyApp#auto_branch_creation_patterns}.
	AutoBranchCreationPatterns *[]*string `json:"autoBranchCreationPatterns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#basic_auth_credentials AmplifyApp#basic_auth_credentials}.
	BasicAuthCredentials *string `json:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#build_spec AmplifyApp#build_spec}.
	BuildSpec *string `json:"buildSpec"`
	// custom_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#custom_rule AmplifyApp#custom_rule}
	CustomRule *[]*AmplifyAppCustomRule `json:"customRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#description AmplifyApp#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_auto_branch_creation AmplifyApp#enable_auto_branch_creation}.
	EnableAutoBranchCreation interface{} `json:"enableAutoBranchCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_basic_auth AmplifyApp#enable_basic_auth}.
	EnableBasicAuth interface{} `json:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_branch_auto_build AmplifyApp#enable_branch_auto_build}.
	EnableBranchAutoBuild interface{} `json:"enableBranchAutoBuild"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#enable_branch_auto_deletion AmplifyApp#enable_branch_auto_deletion}.
	EnableBranchAutoDeletion interface{} `json:"enableBranchAutoDeletion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#environment_variables AmplifyApp#environment_variables}.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#iam_service_role_arn AmplifyApp#iam_service_role_arn}.
	IamServiceRoleArn *string `json:"iamServiceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#oauth_token AmplifyApp#oauth_token}.
	OauthToken *string `json:"oauthToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#platform AmplifyApp#platform}.
	Platform *string `json:"platform"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#repository AmplifyApp#repository}.
	Repository *string `json:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#tags AmplifyApp#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#tags_all AmplifyApp#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type AmplifyAppCustomRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#source AmplifyApp#source}.
	Source *string `json:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#target AmplifyApp#target}.
	Target *string `json:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#condition AmplifyApp#condition}.
	Condition *string `json:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app.html#status AmplifyApp#status}.
	Status *string `json:"status"`
}

type AmplifyAppProductionBranch interface {
	cdktf.ComplexComputedList
	BranchName() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	LastDeployTime() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	ThumbnailUrl() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for AmplifyAppProductionBranch
type jsiiProxy_AmplifyAppProductionBranch struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_AmplifyAppProductionBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) LastDeployTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDeployTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyAppProductionBranch) ThumbnailUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailUrl",
		&returns,
	)
	return returns
}

// Experimental.
func NewAmplifyAppProductionBranch(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) AmplifyAppProductionBranch {
	_init_.Initialize()

	j := jsiiProxy_AmplifyAppProductionBranch{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyAppProductionBranch",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewAmplifyAppProductionBranch_Override(a AmplifyAppProductionBranch, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyAppProductionBranch",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		a,
	)
}

func (j *jsiiProxy_AmplifyAppProductionBranch) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppProductionBranch) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AmplifyAppProductionBranch) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyAppProductionBranch) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyAppProductionBranch) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyAppProductionBranch) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyAppProductionBranch) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyAppProductionBranch) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html aws_amplify_backend_environment}.
type AmplifyBackendEnvironment interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentArtifacts() *string
	SetDeploymentArtifacts(val *string)
	DeploymentArtifactsInput() *string
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	EnvironmentNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackName() *string
	SetStackName(val *string)
	StackNameInput() *string
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
	ResetDeploymentArtifacts()
	ResetOverrideLogicalId()
	ResetStackName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AmplifyBackendEnvironment
type jsiiProxy_AmplifyBackendEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyBackendEnvironment) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) DeploymentArtifacts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) DeploymentArtifactsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) EnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) StackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBackendEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html aws_amplify_backend_environment} Resource.
func NewAmplifyBackendEnvironment(scope constructs.Construct, id *string, config *AmplifyBackendEnvironmentConfig) AmplifyBackendEnvironment {
	_init_.Initialize()

	j := jsiiProxy_AmplifyBackendEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyBackendEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html aws_amplify_backend_environment} Resource.
func NewAmplifyBackendEnvironment_Override(a AmplifyBackendEnvironment, scope constructs.Construct, id *string, config *AmplifyBackendEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyBackendEnvironment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetDeploymentArtifacts(val *string) {
	_jsii_.Set(
		j,
		"deploymentArtifacts",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyBackendEnvironment) SetStackName(val *string) {
	_jsii_.Set(
		j,
		"stackName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AmplifyBackendEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Amplify.AmplifyBackendEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyBackendEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Amplify.AmplifyBackendEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyBackendEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyBackendEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyBackendEnvironment) ResetDeploymentArtifacts() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentArtifacts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AmplifyBackendEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBackendEnvironment) ResetStackName() {
	_jsii_.InvokeVoid(
		a,
		"resetStackName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBackendEnvironment) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) ToMetadata() interface{} {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) ToString() *string {
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
func (a *jsiiProxy_AmplifyBackendEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AmplifyBackendEnvironmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html#app_id AmplifyBackendEnvironment#app_id}.
	AppId *string `json:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html#environment_name AmplifyBackendEnvironment#environment_name}.
	EnvironmentName *string `json:"environmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html#deployment_artifacts AmplifyBackendEnvironment#deployment_artifacts}.
	DeploymentArtifacts *string `json:"deploymentArtifacts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_backend_environment.html#stack_name AmplifyBackendEnvironment#stack_name}.
	StackName *string `json:"stackName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html aws_amplify_branch}.
type AmplifyBranch interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	AssociatedResources() *[]*string
	BackendEnvironmentArn() *string
	SetBackendEnvironmentArn(val *string)
	BackendEnvironmentArnInput() *string
	BasicAuthCredentials() *string
	SetBasicAuthCredentials(val *string)
	BasicAuthCredentialsInput() *string
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomDomains() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationBranch() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnableAutoBuild() interface{}
	SetEnableAutoBuild(val interface{})
	EnableAutoBuildInput() interface{}
	EnableBasicAuth() interface{}
	SetEnableBasicAuth(val interface{})
	EnableBasicAuthInput() interface{}
	EnableNotification() interface{}
	SetEnableNotification(val interface{})
	EnableNotificationInput() interface{}
	EnablePerformanceMode() interface{}
	SetEnablePerformanceMode(val interface{})
	EnablePerformanceModeInput() interface{}
	EnablePullRequestPreview() interface{}
	SetEnablePullRequestPreview(val interface{})
	EnablePullRequestPreviewInput() interface{}
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	EnvironmentVariablesInput() interface{}
	Fqn() *string
	Framework() *string
	SetFramework(val *string)
	FrameworkInput() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PullRequestEnvironmentName() *string
	SetPullRequestEnvironmentName(val *string)
	PullRequestEnvironmentNameInput() *string
	RawOverrides() interface{}
	SourceBranch() *string
	Stage() *string
	SetStage(val *string)
	StageInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetBackendEnvironmentArn()
	ResetBasicAuthCredentials()
	ResetDescription()
	ResetDisplayName()
	ResetEnableAutoBuild()
	ResetEnableBasicAuth()
	ResetEnableNotification()
	ResetEnablePerformanceMode()
	ResetEnablePullRequestPreview()
	ResetEnvironmentVariables()
	ResetFramework()
	ResetOverrideLogicalId()
	ResetPullRequestEnvironmentName()
	ResetStage()
	ResetTags()
	ResetTagsAll()
	ResetTtl()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AmplifyBranch
type jsiiProxy_AmplifyBranch struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyBranch) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) AssociatedResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BackendEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BackendEnvironmentArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendEnvironmentArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) CustomDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DestinationBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableNotification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnableNotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePerformanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePerformanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePullRequestPreview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnablePullRequestPreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePullRequestPreviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) EnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) PullRequestEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) SourceBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyBranch) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html aws_amplify_branch} Resource.
func NewAmplifyBranch(scope constructs.Construct, id *string, config *AmplifyBranchConfig) AmplifyBranch {
	_init_.Initialize()

	j := jsiiProxy_AmplifyBranch{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyBranch",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html aws_amplify_branch} Resource.
func NewAmplifyBranch_Override(a AmplifyBranch, scope constructs.Construct, id *string, config *AmplifyBranchConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyBranch",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetBackendEnvironmentArn(val *string) {
	_jsii_.Set(
		j,
		"backendEnvironmentArn",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetBasicAuthCredentials(val *string) {
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetBranchName(val *string) {
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnableAutoBuild(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnableBasicAuth(val interface{}) {
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnableNotification(val interface{}) {
	_jsii_.Set(
		j,
		"enableNotification",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnablePerformanceMode(val interface{}) {
	_jsii_.Set(
		j,
		"enablePerformanceMode",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnablePullRequestPreview(val interface{}) {
	_jsii_.Set(
		j,
		"enablePullRequestPreview",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetFramework(val *string) {
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetPullRequestEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"pullRequestEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AmplifyBranch) SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AmplifyBranch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Amplify.AmplifyBranch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyBranch_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Amplify.AmplifyBranch",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyBranch) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyBranch) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyBranch) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyBranch) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyBranch) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyBranch) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyBranch) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetBackendEnvironmentArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendEnvironmentArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnableAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnableNotification() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableNotification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnablePerformanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePerformanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnablePullRequestPreview() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePullRequestPreview",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetFramework() {
	_jsii_.InvokeVoid(
		a,
		"resetFramework",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AmplifyBranch) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetPullRequestEnvironmentName() {
	_jsii_.InvokeVoid(
		a,
		"resetPullRequestEnvironmentName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetStage() {
	_jsii_.InvokeVoid(
		a,
		"resetStage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyBranch) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AmplifyBranch) ToMetadata() interface{} {
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
func (a *jsiiProxy_AmplifyBranch) ToString() *string {
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
func (a *jsiiProxy_AmplifyBranch) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AmplifyBranchConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#app_id AmplifyBranch#app_id}.
	AppId *string `json:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#branch_name AmplifyBranch#branch_name}.
	BranchName *string `json:"branchName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#backend_environment_arn AmplifyBranch#backend_environment_arn}.
	BackendEnvironmentArn *string `json:"backendEnvironmentArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#basic_auth_credentials AmplifyBranch#basic_auth_credentials}.
	BasicAuthCredentials *string `json:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#description AmplifyBranch#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#display_name AmplifyBranch#display_name}.
	DisplayName *string `json:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#enable_auto_build AmplifyBranch#enable_auto_build}.
	EnableAutoBuild interface{} `json:"enableAutoBuild"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#enable_basic_auth AmplifyBranch#enable_basic_auth}.
	EnableBasicAuth interface{} `json:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#enable_notification AmplifyBranch#enable_notification}.
	EnableNotification interface{} `json:"enableNotification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#enable_performance_mode AmplifyBranch#enable_performance_mode}.
	EnablePerformanceMode interface{} `json:"enablePerformanceMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#enable_pull_request_preview AmplifyBranch#enable_pull_request_preview}.
	EnablePullRequestPreview interface{} `json:"enablePullRequestPreview"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#environment_variables AmplifyBranch#environment_variables}.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#framework AmplifyBranch#framework}.
	Framework *string `json:"framework"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#pull_request_environment_name AmplifyBranch#pull_request_environment_name}.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#stage AmplifyBranch#stage}.
	Stage *string `json:"stage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#tags AmplifyBranch#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#tags_all AmplifyBranch#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_branch.html#ttl AmplifyBranch#ttl}.
	Ttl *string `json:"ttl"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html aws_amplify_domain_association}.
type AmplifyDomainAssociation interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CertificateVerificationDnsRecord() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	SubDomain() *[]*AmplifyDomainAssociationSubDomain
	SetSubDomain(val *[]*AmplifyDomainAssociationSubDomain)
	SubDomainInput() *[]*AmplifyDomainAssociationSubDomain
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WaitForVerification() interface{}
	SetWaitForVerification(val interface{})
	WaitForVerificationInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetWaitForVerification()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AmplifyDomainAssociation
type jsiiProxy_AmplifyDomainAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyDomainAssociation) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) CertificateVerificationDnsRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateVerificationDnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) SubDomain() *[]*AmplifyDomainAssociationSubDomain {
	var returns *[]*AmplifyDomainAssociationSubDomain
	_jsii_.Get(
		j,
		"subDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) SubDomainInput() *[]*AmplifyDomainAssociationSubDomain {
	var returns *[]*AmplifyDomainAssociationSubDomain
	_jsii_.Get(
		j,
		"subDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) WaitForVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomainAssociation) WaitForVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForVerificationInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html aws_amplify_domain_association} Resource.
func NewAmplifyDomainAssociation(scope constructs.Construct, id *string, config *AmplifyDomainAssociationConfig) AmplifyDomainAssociation {
	_init_.Initialize()

	j := jsiiProxy_AmplifyDomainAssociation{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyDomainAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html aws_amplify_domain_association} Resource.
func NewAmplifyDomainAssociation_Override(a AmplifyDomainAssociation, scope constructs.Construct, id *string, config *AmplifyDomainAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyDomainAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetSubDomain(val *[]*AmplifyDomainAssociationSubDomain) {
	_jsii_.Set(
		j,
		"subDomain",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomainAssociation) SetWaitForVerification(val interface{}) {
	_jsii_.Set(
		j,
		"waitForVerification",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AmplifyDomainAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Amplify.AmplifyDomainAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyDomainAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Amplify.AmplifyDomainAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyDomainAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyDomainAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyDomainAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyDomainAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyDomainAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyDomainAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyDomainAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AmplifyDomainAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomainAssociation) ResetWaitForVerification() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForVerification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomainAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AmplifyDomainAssociation) ToMetadata() interface{} {
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
func (a *jsiiProxy_AmplifyDomainAssociation) ToString() *string {
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
func (a *jsiiProxy_AmplifyDomainAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AmplifyDomainAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#app_id AmplifyDomainAssociation#app_id}.
	AppId *string `json:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#domain_name AmplifyDomainAssociation#domain_name}.
	DomainName *string `json:"domainName"`
	// sub_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#sub_domain AmplifyDomainAssociation#sub_domain}
	SubDomain *[]*AmplifyDomainAssociationSubDomain `json:"subDomain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#wait_for_verification AmplifyDomainAssociation#wait_for_verification}.
	WaitForVerification interface{} `json:"waitForVerification"`
}

type AmplifyDomainAssociationSubDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#branch_name AmplifyDomainAssociation#branch_name}.
	BranchName *string `json:"branchName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_domain_association.html#prefix AmplifyDomainAssociation#prefix}.
	Prefix *string `json:"prefix"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html aws_amplify_webhook}.
type AmplifyWebhook interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
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
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Url() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDescription()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AmplifyWebhook
type jsiiProxy_AmplifyWebhook struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyWebhook) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyWebhook) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html aws_amplify_webhook} Resource.
func NewAmplifyWebhook(scope constructs.Construct, id *string, config *AmplifyWebhookConfig) AmplifyWebhook {
	_init_.Initialize()

	j := jsiiProxy_AmplifyWebhook{}

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyWebhook",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html aws_amplify_webhook} Resource.
func NewAmplifyWebhook_Override(a AmplifyWebhook, scope constructs.Construct, id *string, config *AmplifyWebhookConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Amplify.AmplifyWebhook",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetBranchName(val *string) {
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyWebhook) SetProvider(val cdktf.TerraformProvider) {
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
func AmplifyWebhook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Amplify.AmplifyWebhook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyWebhook_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Amplify.AmplifyWebhook",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AmplifyWebhook) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AmplifyWebhook) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyWebhook) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AmplifyWebhook) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AmplifyWebhook) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AmplifyWebhook) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AmplifyWebhook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyWebhook) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AmplifyWebhook) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyWebhook) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AmplifyWebhook) ToMetadata() interface{} {
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
func (a *jsiiProxy_AmplifyWebhook) ToString() *string {
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
func (a *jsiiProxy_AmplifyWebhook) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AmplifyWebhookConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html#app_id AmplifyWebhook#app_id}.
	AppId *string `json:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html#branch_name AmplifyWebhook#branch_name}.
	BranchName *string `json:"branchName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_webhook.html#description AmplifyWebhook#description}.
	Description *string `json:"description"`
}
