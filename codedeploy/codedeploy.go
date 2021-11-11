package codedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/codedeploy/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html aws_codedeploy_app}.
type CodedeployApp interface {
	cdktf.TerraformResource
	ApplicationId() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputePlatform() *string
	SetComputePlatform(val *string)
	ComputePlatformInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	GithubAccountName() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedToGithub() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetComputePlatform()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodedeployApp
type jsiiProxy_CodedeployApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodedeployApp) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) ComputePlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) GithubAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) LinkedToGithub() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedToGithub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html aws_codedeploy_app} Resource.
func NewCodedeployApp(scope constructs.Construct, id *string, config *CodedeployAppConfig) CodedeployApp {
	_init_.Initialize()

	j := jsiiProxy_CodedeployApp{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html aws_codedeploy_app} Resource.
func NewCodedeployApp_Override(c CodedeployApp, scope constructs.Construct, id *string, config *CodedeployAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployApp",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployApp) SetComputePlatform(val *string) {
	_jsii_.Set(
		j,
		"computePlatform",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodedeployApp) SetTagsAll(val interface{}) {
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
func CodedeployApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CodeDeploy.CodedeployApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodedeployApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CodeDeploy.CodedeployApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodedeployApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployApp) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployApp) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployApp) ResetComputePlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetComputePlatform",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodedeployApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployApp) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployApp) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployApp) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CodedeployApp) ToMetadata() interface{} {
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
func (c *jsiiProxy_CodedeployApp) ToString() *string {
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
func (c *jsiiProxy_CodedeployApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CodedeployAppConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html#name CodedeployApp#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html#compute_platform CodedeployApp#compute_platform}.
	ComputePlatform *string `json:"computePlatform"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html#tags CodedeployApp#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_app.html#tags_all CodedeployApp#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html aws_codedeploy_deployment_config}.
type CodedeployDeploymentConfig interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ComputePlatform() *string
	SetComputePlatform(val *string)
	ComputePlatformInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentConfigId() *string
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentConfigNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinimumHealthyHosts() CodedeployDeploymentConfigMinimumHealthyHostsOutputReference
	MinimumHealthyHostsInput() *CodedeployDeploymentConfigMinimumHealthyHosts
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TrafficRoutingConfig() CodedeployDeploymentConfigTrafficRoutingConfigOutputReference
	TrafficRoutingConfigInput() *CodedeployDeploymentConfigTrafficRoutingConfig
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutMinimumHealthyHosts(value *CodedeployDeploymentConfigMinimumHealthyHosts)
	PutTrafficRoutingConfig(value *CodedeployDeploymentConfigTrafficRoutingConfig)
	ResetComputePlatform()
	ResetMinimumHealthyHosts()
	ResetOverrideLogicalId()
	ResetTrafficRoutingConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodedeployDeploymentConfig
type jsiiProxy_CodedeployDeploymentConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodedeployDeploymentConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ComputePlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DeploymentConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) MinimumHealthyHosts() CodedeployDeploymentConfigMinimumHealthyHostsOutputReference {
	var returns CodedeployDeploymentConfigMinimumHealthyHostsOutputReference
	_jsii_.Get(
		j,
		"minimumHealthyHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) MinimumHealthyHostsInput() *CodedeployDeploymentConfigMinimumHealthyHosts {
	var returns *CodedeployDeploymentConfigMinimumHealthyHosts
	_jsii_.Get(
		j,
		"minimumHealthyHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TrafficRoutingConfig() CodedeployDeploymentConfigTrafficRoutingConfigOutputReference {
	var returns CodedeployDeploymentConfigTrafficRoutingConfigOutputReference
	_jsii_.Get(
		j,
		"trafficRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfig) TrafficRoutingConfigInput() *CodedeployDeploymentConfigTrafficRoutingConfig {
	var returns *CodedeployDeploymentConfigTrafficRoutingConfig
	_jsii_.Get(
		j,
		"trafficRoutingConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html aws_codedeploy_deployment_config} Resource.
func NewCodedeployDeploymentConfig(scope constructs.Construct, id *string, config *CodedeployDeploymentConfigConfig) CodedeployDeploymentConfig {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentConfig{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html aws_codedeploy_deployment_config} Resource.
func NewCodedeployDeploymentConfig_Override(c CodedeployDeploymentConfig, scope constructs.Construct, id *string, config *CodedeployDeploymentConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfig",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetComputePlatform(val *string) {
	_jsii_.Set(
		j,
		"computePlatform",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetDeploymentConfigName(val *string) {
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfig) SetProvider(val cdktf.TerraformProvider) {
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
func CodedeployDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodedeployDeploymentConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) PutMinimumHealthyHosts(value *CodedeployDeploymentConfigMinimumHealthyHosts) {
	_jsii_.InvokeVoid(
		c,
		"putMinimumHealthyHosts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) PutTrafficRoutingConfig(value *CodedeployDeploymentConfigTrafficRoutingConfig) {
	_jsii_.InvokeVoid(
		c,
		"putTrafficRoutingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetComputePlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetComputePlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetMinimumHealthyHosts() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumHealthyHosts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) ResetTrafficRoutingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTrafficRoutingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfig) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) ToMetadata() interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) ToString() *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CodedeployDeploymentConfigConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#deployment_config_name CodedeployDeploymentConfig#deployment_config_name}.
	DeploymentConfigName *string `json:"deploymentConfigName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#compute_platform CodedeployDeploymentConfig#compute_platform}.
	ComputePlatform *string `json:"computePlatform"`
	// minimum_healthy_hosts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#minimum_healthy_hosts CodedeployDeploymentConfig#minimum_healthy_hosts}
	MinimumHealthyHosts *CodedeployDeploymentConfigMinimumHealthyHosts `json:"minimumHealthyHosts"`
	// traffic_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#traffic_routing_config CodedeployDeploymentConfig#traffic_routing_config}
	TrafficRoutingConfig *CodedeployDeploymentConfigTrafficRoutingConfig `json:"trafficRoutingConfig"`
}

type CodedeployDeploymentConfigMinimumHealthyHosts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#type CodedeployDeploymentConfig#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#value CodedeployDeploymentConfig#value}.
	Value *float64 `json:"value"`
}

type CodedeployDeploymentConfigMinimumHealthyHostsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
	ResetValue()
}

// The jsii proxy struct for CodedeployDeploymentConfigMinimumHealthyHostsOutputReference
type jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentConfigMinimumHealthyHostsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentConfigMinimumHealthyHostsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigMinimumHealthyHostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentConfigMinimumHealthyHostsOutputReference_Override(c CodedeployDeploymentConfigMinimumHealthyHostsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigMinimumHealthyHostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

type CodedeployDeploymentConfigTrafficRoutingConfig struct {
	// time_based_canary block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#time_based_canary CodedeployDeploymentConfig#time_based_canary}
	TimeBasedCanary *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary `json:"timeBasedCanary"`
	// time_based_linear block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#time_based_linear CodedeployDeploymentConfig#time_based_linear}
	TimeBasedLinear *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear `json:"timeBasedLinear"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#type CodedeployDeploymentConfig#type}.
	Type *string `json:"type"`
}

type CodedeployDeploymentConfigTrafficRoutingConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TimeBasedCanary() CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference
	TimeBasedCanaryInput() *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary
	TimeBasedLinear() CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference
	TimeBasedLinearInput() *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear
	Type() *string
	SetType(val *string)
	TypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutTimeBasedCanary(value *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary)
	PutTimeBasedLinear(value *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear)
	ResetTimeBasedCanary()
	ResetTimeBasedLinear()
	ResetType()
}

// The jsii proxy struct for CodedeployDeploymentConfigTrafficRoutingConfigOutputReference
type jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TimeBasedCanary() CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference {
	var returns CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference
	_jsii_.Get(
		j,
		"timeBasedCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TimeBasedCanaryInput() *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary {
	var returns *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary
	_jsii_.Get(
		j,
		"timeBasedCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TimeBasedLinear() CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference {
	var returns CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference
	_jsii_.Get(
		j,
		"timeBasedLinear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TimeBasedLinearInput() *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear {
	var returns *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear
	_jsii_.Get(
		j,
		"timeBasedLinearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentConfigTrafficRoutingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigOutputReference_Override(c CodedeployDeploymentConfigTrafficRoutingConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) PutTimeBasedCanary(value *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary) {
	_jsii_.InvokeVoid(
		c,
		"putTimeBasedCanary",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) PutTimeBasedLinear(value *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear) {
	_jsii_.InvokeVoid(
		c,
		"putTimeBasedLinear",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) ResetTimeBasedCanary() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeBasedCanary",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) ResetTimeBasedLinear() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeBasedLinear",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#interval CodedeployDeploymentConfig#interval}.
	Interval *float64 `json:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#percentage CodedeployDeploymentConfig#percentage}.
	Percentage *float64 `json:"percentage"`
}

type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference interface {
	cdktf.ComplexObject
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
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
	ResetInterval()
	ResetPercentage()
}

// The jsii proxy struct for CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference
type jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference_Override(c CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) SetPercentage(val *float64) {
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetPercentage",
		nil, // no parameters
	)
}

type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#interval CodedeployDeploymentConfig#interval}.
	Interval *float64 `json:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config.html#percentage CodedeployDeploymentConfig#percentage}.
	Percentage *float64 `json:"percentage"`
}

type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference interface {
	cdktf.ComplexObject
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
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
	ResetInterval()
	ResetPercentage()
}

// The jsii proxy struct for CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference
type jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference_Override(c CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) SetPercentage(val *float64) {
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetPercentage",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html aws_codedeploy_deployment_group}.
type CodedeployDeploymentGroup interface {
	cdktf.TerraformResource
	AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference
	AlarmConfigurationInput() *CodedeployDeploymentGroupAlarmConfiguration
	AppName() *string
	SetAppName(val *string)
	AppNameInput() *string
	Arn() *string
	AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	AutoRollbackConfigurationInput() *CodedeployDeploymentGroupAutoRollbackConfiguration
	AutoscalingGroups() *[]*string
	SetAutoscalingGroups(val *[]*string)
	AutoscalingGroupsInput() *[]*string
	BlueGreenDeploymentConfig() CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference
	BlueGreenDeploymentConfigInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfig
	CdktfStack() cdktf.TerraformStack
	ComputePlatform() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentConfigNameInput() *string
	DeploymentGroupId() *string
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	DeploymentGroupNameInput() *string
	DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference
	DeploymentStyleInput() *CodedeployDeploymentGroupDeploymentStyle
	Ec2TagFilter() *[]*CodedeployDeploymentGroupEc2TagFilter
	SetEc2TagFilter(val *[]*CodedeployDeploymentGroupEc2TagFilter)
	Ec2TagFilterInput() *[]*CodedeployDeploymentGroupEc2TagFilter
	Ec2TagSet() *[]*CodedeployDeploymentGroupEc2TagSet
	SetEc2TagSet(val *[]*CodedeployDeploymentGroupEc2TagSet)
	Ec2TagSetInput() *[]*CodedeployDeploymentGroupEc2TagSet
	EcsService() CodedeployDeploymentGroupEcsServiceOutputReference
	EcsServiceInput() *CodedeployDeploymentGroupEcsService
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	LoadBalancerInfoInput() *CodedeployDeploymentGroupLoadBalancerInfo
	Node() constructs.Node
	OnPremisesInstanceTagFilter() *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter
	SetOnPremisesInstanceTagFilter(val *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter)
	OnPremisesInstanceTagFilterInput() *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TriggerConfiguration() *[]*CodedeployDeploymentGroupTriggerConfiguration
	SetTriggerConfiguration(val *[]*CodedeployDeploymentGroupTriggerConfiguration)
	TriggerConfigurationInput() *[]*CodedeployDeploymentGroupTriggerConfiguration
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration)
	PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration)
	PutBlueGreenDeploymentConfig(value *CodedeployDeploymentGroupBlueGreenDeploymentConfig)
	PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle)
	PutEcsService(value *CodedeployDeploymentGroupEcsService)
	PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo)
	ResetAlarmConfiguration()
	ResetAutoRollbackConfiguration()
	ResetAutoscalingGroups()
	ResetBlueGreenDeploymentConfig()
	ResetDeploymentConfigName()
	ResetDeploymentStyle()
	ResetEc2TagFilter()
	ResetEc2TagSet()
	ResetEcsService()
	ResetLoadBalancerInfo()
	ResetOnPremisesInstanceTagFilter()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTriggerConfiguration()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodedeployDeploymentGroup
type jsiiProxy_CodedeployDeploymentGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAlarmConfigurationOutputReference
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfigurationInput() *CodedeployDeploymentGroupAlarmConfiguration {
	var returns *CodedeployDeploymentGroupAlarmConfiguration
	_jsii_.Get(
		j,
		"alarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfigurationInput() *CodedeployDeploymentGroupAutoRollbackConfiguration {
	var returns *CodedeployDeploymentGroupAutoRollbackConfiguration
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoscalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoscalingGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfig() CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfigInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfig {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfig
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference {
	var returns CodedeployDeploymentGroupDeploymentStyleOutputReference
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyleInput() *CodedeployDeploymentGroupDeploymentStyle {
	var returns *CodedeployDeploymentGroupDeploymentStyle
	_jsii_.Get(
		j,
		"deploymentStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFilter() *[]*CodedeployDeploymentGroupEc2TagFilter {
	var returns *[]*CodedeployDeploymentGroupEc2TagFilter
	_jsii_.Get(
		j,
		"ec2TagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFilterInput() *[]*CodedeployDeploymentGroupEc2TagFilter {
	var returns *[]*CodedeployDeploymentGroupEc2TagFilter
	_jsii_.Get(
		j,
		"ec2TagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSet() *[]*CodedeployDeploymentGroupEc2TagSet {
	var returns *[]*CodedeployDeploymentGroupEc2TagSet
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSetInput() *[]*CodedeployDeploymentGroupEc2TagSet {
	var returns *[]*CodedeployDeploymentGroupEc2TagSet
	_jsii_.Get(
		j,
		"ec2TagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsService() CodedeployDeploymentGroupEcsServiceOutputReference {
	var returns CodedeployDeploymentGroupEcsServiceOutputReference
	_jsii_.Get(
		j,
		"ecsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsServiceInput() *CodedeployDeploymentGroupEcsService {
	var returns *CodedeployDeploymentGroupEcsService
	_jsii_.Get(
		j,
		"ecsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfoInput() *CodedeployDeploymentGroupLoadBalancerInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfo
	_jsii_.Get(
		j,
		"loadBalancerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFilter() *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter {
	var returns *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFilterInput() *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter {
	var returns *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfiguration() *[]*CodedeployDeploymentGroupTriggerConfiguration {
	var returns *[]*CodedeployDeploymentGroupTriggerConfiguration
	_jsii_.Get(
		j,
		"triggerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfigurationInput() *[]*CodedeployDeploymentGroupTriggerConfiguration {
	var returns *[]*CodedeployDeploymentGroupTriggerConfiguration
	_jsii_.Get(
		j,
		"triggerConfigurationInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html aws_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup(scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) CodedeployDeploymentGroup {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroup{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html aws_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup_Override(c CodedeployDeploymentGroup, scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetAppName(val *string) {
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetAutoscalingGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"autoscalingGroups",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDeploymentConfigName(val *string) {
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDeploymentGroupName(val *string) {
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetEc2TagFilter(val *[]*CodedeployDeploymentGroupEc2TagFilter) {
	_jsii_.Set(
		j,
		"ec2TagFilter",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetEc2TagSet(val *[]*CodedeployDeploymentGroupEc2TagSet) {
	_jsii_.Set(
		j,
		"ec2TagSet",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetOnPremisesInstanceTagFilter(val *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter) {
	_jsii_.Set(
		j,
		"onPremisesInstanceTagFilter",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetTriggerConfiguration(val *[]*CodedeployDeploymentGroupTriggerConfiguration) {
	_jsii_.Set(
		j,
		"triggerConfiguration",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CodedeployDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodedeployDeploymentGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putAlarmConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration) {
	_jsii_.InvokeVoid(
		c,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutBlueGreenDeploymentConfig(value *CodedeployDeploymentGroupBlueGreenDeploymentConfig) {
	_jsii_.InvokeVoid(
		c,
		"putBlueGreenDeploymentConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle) {
	_jsii_.InvokeVoid(
		c,
		"putDeploymentStyle",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEcsService(value *CodedeployDeploymentGroupEcsService) {
	_jsii_.InvokeVoid(
		c,
		"putEcsService",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo) {
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancerInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAlarmConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoscalingGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscalingGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetBlueGreenDeploymentConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBlueGreenDeploymentConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentStyle() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentStyle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagSet() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEcsService() {
	_jsii_.InvokeVoid(
		c,
		"resetEcsService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetLoadBalancerInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancerInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOnPremisesInstanceTagFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetOnPremisesInstanceTagFilter",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTriggerConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) ToMetadata() interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) ToString() *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CodedeployDeploymentGroupAlarmConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#alarms CodedeployDeploymentGroup#alarms}.
	Alarms *[]*string `json:"alarms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#enabled CodedeployDeploymentGroup#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#ignore_poll_alarm_failure CodedeployDeploymentGroup#ignore_poll_alarm_failure}.
	IgnorePollAlarmFailure interface{} `json:"ignorePollAlarmFailure"`
}

type CodedeployDeploymentGroupAlarmConfigurationOutputReference interface {
	cdktf.ComplexObject
	Alarms() *[]*string
	SetAlarms(val *[]*string)
	AlarmsInput() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IgnorePollAlarmFailure() interface{}
	SetIgnorePollAlarmFailure(val interface{})
	IgnorePollAlarmFailureInput() interface{}
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
	ResetAlarms()
	ResetEnabled()
	ResetIgnorePollAlarmFailure()
}

// The jsii proxy struct for CodedeployDeploymentGroupAlarmConfigurationOutputReference
type jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) Alarms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) AlarmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) IgnorePollAlarmFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePollAlarmFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) IgnorePollAlarmFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePollAlarmFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupAlarmConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupAlarmConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAlarmConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupAlarmConfigurationOutputReference_Override(c CodedeployDeploymentGroupAlarmConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAlarmConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetAlarms(val *[]*string) {
	_jsii_.Set(
		j,
		"alarms",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetIgnorePollAlarmFailure(val interface{}) {
	_jsii_.Set(
		j,
		"ignorePollAlarmFailure",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference) ResetIgnorePollAlarmFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnorePollAlarmFailure",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupAutoRollbackConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#enabled CodedeployDeploymentGroup#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#events CodedeployDeploymentGroup#events}.
	Events *[]*string `json:"events"`
}

type CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
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
	ResetEnabled()
	ResetEvents()
}

// The jsii proxy struct for CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
type jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupAutoRollbackConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupAutoRollbackConfigurationOutputReference_Override(c CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) SetEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetEvents",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfig struct {
	// deployment_ready_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_ready_option CodedeployDeploymentGroup#deployment_ready_option}
	DeploymentReadyOption *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption `json:"deploymentReadyOption"`
	// green_fleet_provisioning_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#green_fleet_provisioning_option CodedeployDeploymentGroup#green_fleet_provisioning_option}
	GreenFleetProvisioningOption *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption `json:"greenFleetProvisioningOption"`
	// terminate_blue_instances_on_deployment_success block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#terminate_blue_instances_on_deployment_success CodedeployDeploymentGroup#terminate_blue_instances_on_deployment_success}
	TerminateBlueInstancesOnDeploymentSuccess *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `json:"terminateBlueInstancesOnDeploymentSuccess"`
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#action_on_timeout CodedeployDeploymentGroup#action_on_timeout}.
	ActionOnTimeout *string `json:"actionOnTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#wait_time_in_minutes CodedeployDeploymentGroup#wait_time_in_minutes}.
	WaitTimeInMinutes *float64 `json:"waitTimeInMinutes"`
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference interface {
	cdktf.ComplexObject
	ActionOnTimeout() *string
	SetActionOnTimeout(val *string)
	ActionOnTimeoutInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WaitTimeInMinutes() *float64
	SetWaitTimeInMinutes(val *float64)
	WaitTimeInMinutesInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetActionOnTimeout()
	ResetWaitTimeInMinutes()
}

// The jsii proxy struct for CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference
type jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ActionOnTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ActionOnTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) WaitTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) WaitTimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeInMinutesInput",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference_Override(c CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) SetActionOnTimeout(val *string) {
	_jsii_.Set(
		j,
		"actionOnTimeout",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) SetWaitTimeInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"waitTimeInMinutes",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ResetActionOnTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetActionOnTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ResetWaitTimeInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitTimeInMinutes",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#action CodedeployDeploymentGroup#action}.
	Action *string `json:"action"`
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	ResetAction()
}

// The jsii proxy struct for CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference
type jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference_Override(c CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		c,
		"resetAction",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference interface {
	cdktf.ComplexObject
	DeploymentReadyOption() CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference
	DeploymentReadyOptionInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption
	GreenFleetProvisioningOption() CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference
	GreenFleetProvisioningOptionInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerminateBlueInstancesOnDeploymentSuccess() CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference
	TerminateBlueInstancesOnDeploymentSuccessInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess
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
	PutDeploymentReadyOption(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption)
	PutGreenFleetProvisioningOption(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption)
	PutTerminateBlueInstancesOnDeploymentSuccess(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess)
	ResetDeploymentReadyOption()
	ResetGreenFleetProvisioningOption()
	ResetTerminateBlueInstancesOnDeploymentSuccess()
}

// The jsii proxy struct for CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference
type jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) DeploymentReadyOption() CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference
	_jsii_.Get(
		j,
		"deploymentReadyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) DeploymentReadyOptionInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption
	_jsii_.Get(
		j,
		"deploymentReadyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GreenFleetProvisioningOption() CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference
	_jsii_.Get(
		j,
		"greenFleetProvisioningOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GreenFleetProvisioningOptionInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption
	_jsii_.Get(
		j,
		"greenFleetProvisioningOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) TerminateBlueInstancesOnDeploymentSuccess() CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) TerminateBlueInstancesOnDeploymentSuccessInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference_Override(c CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) PutDeploymentReadyOption(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption) {
	_jsii_.InvokeVoid(
		c,
		"putDeploymentReadyOption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) PutGreenFleetProvisioningOption(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption) {
	_jsii_.InvokeVoid(
		c,
		"putGreenFleetProvisioningOption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) PutTerminateBlueInstancesOnDeploymentSuccess(value *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess) {
	_jsii_.InvokeVoid(
		c,
		"putTerminateBlueInstancesOnDeploymentSuccess",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) ResetDeploymentReadyOption() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentReadyOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) ResetGreenFleetProvisioningOption() {
	_jsii_.InvokeVoid(
		c,
		"resetGreenFleetProvisioningOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference) ResetTerminateBlueInstancesOnDeploymentSuccess() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminateBlueInstancesOnDeploymentSuccess",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#action CodedeployDeploymentGroup#action}.
	Action *string `json:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#termination_wait_time_in_minutes CodedeployDeploymentGroup#termination_wait_time_in_minutes}.
	TerminationWaitTimeInMinutes *float64 `json:"terminationWaitTimeInMinutes"`
}

type CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerminationWaitTimeInMinutes() *float64
	SetTerminationWaitTimeInMinutes(val *float64)
	TerminationWaitTimeInMinutesInput() *float64
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
	ResetAction()
	ResetTerminationWaitTimeInMinutes()
}

// The jsii proxy struct for CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference
type jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) TerminationWaitTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) TerminationWaitTimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitTimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference_Override(c CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) SetTerminationWaitTimeInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"terminationWaitTimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		c,
		"resetAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference) ResetTerminationWaitTimeInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationWaitTimeInMinutes",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#app_name CodedeployDeploymentGroup#app_name}.
	AppName *string `json:"appName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_group_name CodedeployDeploymentGroup#deployment_group_name}.
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#service_role_arn CodedeployDeploymentGroup#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn"`
	// alarm_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#alarm_configuration CodedeployDeploymentGroup#alarm_configuration}
	AlarmConfiguration *CodedeployDeploymentGroupAlarmConfiguration `json:"alarmConfiguration"`
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#auto_rollback_configuration CodedeployDeploymentGroup#auto_rollback_configuration}
	AutoRollbackConfiguration *CodedeployDeploymentGroupAutoRollbackConfiguration `json:"autoRollbackConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#autoscaling_groups CodedeployDeploymentGroup#autoscaling_groups}.
	AutoscalingGroups *[]*string `json:"autoscalingGroups"`
	// blue_green_deployment_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#blue_green_deployment_config CodedeployDeploymentGroup#blue_green_deployment_config}
	BlueGreenDeploymentConfig *CodedeployDeploymentGroupBlueGreenDeploymentConfig `json:"blueGreenDeploymentConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_config_name CodedeployDeploymentGroup#deployment_config_name}.
	DeploymentConfigName *string `json:"deploymentConfigName"`
	// deployment_style block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_style CodedeployDeploymentGroup#deployment_style}
	DeploymentStyle *CodedeployDeploymentGroupDeploymentStyle `json:"deploymentStyle"`
	// ec2_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#ec2_tag_filter CodedeployDeploymentGroup#ec2_tag_filter}
	Ec2TagFilter *[]*CodedeployDeploymentGroupEc2TagFilter `json:"ec2TagFilter"`
	// ec2_tag_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#ec2_tag_set CodedeployDeploymentGroup#ec2_tag_set}
	Ec2TagSet *[]*CodedeployDeploymentGroupEc2TagSet `json:"ec2TagSet"`
	// ecs_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#ecs_service CodedeployDeploymentGroup#ecs_service}
	EcsService *CodedeployDeploymentGroupEcsService `json:"ecsService"`
	// load_balancer_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#load_balancer_info CodedeployDeploymentGroup#load_balancer_info}
	LoadBalancerInfo *CodedeployDeploymentGroupLoadBalancerInfo `json:"loadBalancerInfo"`
	// on_premises_instance_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#on_premises_instance_tag_filter CodedeployDeploymentGroup#on_premises_instance_tag_filter}
	OnPremisesInstanceTagFilter *[]*CodedeployDeploymentGroupOnPremisesInstanceTagFilter `json:"onPremisesInstanceTagFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#tags CodedeployDeploymentGroup#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#tags_all CodedeployDeploymentGroup#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// trigger_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#trigger_configuration CodedeployDeploymentGroup#trigger_configuration}
	TriggerConfiguration *[]*CodedeployDeploymentGroupTriggerConfiguration `json:"triggerConfiguration"`
}

type CodedeployDeploymentGroupDeploymentStyle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_option CodedeployDeploymentGroup#deployment_option}.
	DeploymentOption *string `json:"deploymentOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#deployment_type CodedeployDeploymentGroup#deployment_type}.
	DeploymentType *string `json:"deploymentType"`
}

type CodedeployDeploymentGroupDeploymentStyleOutputReference interface {
	cdktf.ComplexObject
	DeploymentOption() *string
	SetDeploymentOption(val *string)
	DeploymentOptionInput() *string
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
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
	ResetDeploymentOption()
	ResetDeploymentType()
}

// The jsii proxy struct for CodedeployDeploymentGroupDeploymentStyleOutputReference
type jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) DeploymentOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) DeploymentOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupDeploymentStyleOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupDeploymentStyleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupDeploymentStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupDeploymentStyleOutputReference_Override(c CodedeployDeploymentGroupDeploymentStyleOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupDeploymentStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) SetDeploymentOption(val *string) {
	_jsii_.Set(
		j,
		"deploymentOption",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) SetDeploymentType(val *string) {
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) ResetDeploymentOption() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentType",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupEc2TagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#key CodedeployDeploymentGroup#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#type CodedeployDeploymentGroup#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#value CodedeployDeploymentGroup#value}.
	Value *string `json:"value"`
}

type CodedeployDeploymentGroupEc2TagSet struct {
	// ec2_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#ec2_tag_filter CodedeployDeploymentGroup#ec2_tag_filter}
	Ec2TagFilter *[]*CodedeployDeploymentGroupEc2TagSetEc2TagFilter `json:"ec2TagFilter"`
}

type CodedeployDeploymentGroupEc2TagSetEc2TagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#key CodedeployDeploymentGroup#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#type CodedeployDeploymentGroup#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#value CodedeployDeploymentGroup#value}.
	Value *string `json:"value"`
}

type CodedeployDeploymentGroupEcsService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#cluster_name CodedeployDeploymentGroup#cluster_name}.
	ClusterName *string `json:"clusterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#service_name CodedeployDeploymentGroup#service_name}.
	ServiceName *string `json:"serviceName"`
}

type CodedeployDeploymentGroupEcsServiceOutputReference interface {
	cdktf.ComplexObject
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
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

// The jsii proxy struct for CodedeployDeploymentGroupEcsServiceOutputReference
type jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupEcsServiceOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupEcsServiceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEcsServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupEcsServiceOutputReference_Override(c CodedeployDeploymentGroupEcsServiceOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEcsServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodedeployDeploymentGroupLoadBalancerInfo struct {
	// elb_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#elb_info CodedeployDeploymentGroup#elb_info}
	ElbInfo *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo `json:"elbInfo"`
	// target_group_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#target_group_info CodedeployDeploymentGroup#target_group_info}
	TargetGroupInfo *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo `json:"targetGroupInfo"`
	// target_group_pair_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#target_group_pair_info CodedeployDeploymentGroup#target_group_pair_info}
	TargetGroupPairInfo *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo `json:"targetGroupPairInfo"`
}

type CodedeployDeploymentGroupLoadBalancerInfoElbInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#name CodedeployDeploymentGroup#name}.
	Name *string `json:"name"`
}

type CodedeployDeploymentGroupLoadBalancerInfoOutputReference interface {
	cdktf.ComplexObject
	ElbInfo() *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo
	SetElbInfo(val *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo)
	ElbInfoInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TargetGroupInfo() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo
	SetTargetGroupInfo(val *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo)
	TargetGroupInfoInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo
	TargetGroupPairInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
	TargetGroupPairInfoInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
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
	PutTargetGroupPairInfo(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo)
	ResetElbInfo()
	ResetTargetGroupInfo()
	ResetTargetGroupPairInfo()
}

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ElbInfo() *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo
	_jsii_.Get(
		j,
		"elbInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ElbInfoInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo
	_jsii_.Get(
		j,
		"elbInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupInfo() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo
	_jsii_.Get(
		j,
		"targetGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupInfoInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo
	_jsii_.Get(
		j,
		"targetGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupPairInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
	_jsii_.Get(
		j,
		"targetGroupPairInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupPairInfoInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
	_jsii_.Get(
		j,
		"targetGroupPairInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupLoadBalancerInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupLoadBalancerInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) SetElbInfo(val *[]*CodedeployDeploymentGroupLoadBalancerInfoElbInfo) {
	_jsii_.Set(
		j,
		"elbInfo",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) SetTargetGroupInfo(val *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo) {
	_jsii_.Set(
		j,
		"targetGroupInfo",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) PutTargetGroupPairInfo(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo) {
	_jsii_.InvokeVoid(
		c,
		"putTargetGroupPairInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetElbInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetElbInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetTargetGroupInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetTargetGroupPairInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupPairInfo",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#name CodedeployDeploymentGroup#name}.
	Name *string `json:"name"`
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo struct {
	// prod_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#prod_traffic_route CodedeployDeploymentGroup#prod_traffic_route}
	ProdTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute `json:"prodTrafficRoute"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#target_group CodedeployDeploymentGroup#target_group}
	TargetGroup *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup `json:"targetGroup"`
	// test_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#test_traffic_route CodedeployDeploymentGroup#test_traffic_route}
	TestTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute `json:"testTrafficRoute"`
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ProdTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference
	ProdTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute
	TargetGroup() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup
	SetTargetGroup(val *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup)
	TargetGroupInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TestTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference
	TestTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutProdTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)
	PutTestTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)
	ResetTestTrafficRoute()
}

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ProdTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference
	_jsii_.Get(
		j,
		"prodTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ProdTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute
	_jsii_.Get(
		j,
		"prodTrafficRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TargetGroup() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TargetGroupInput() *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup {
	var returns *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup
	_jsii_.Get(
		j,
		"targetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TestTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference
	_jsii_.Get(
		j,
		"testTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TestTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute
	_jsii_.Get(
		j,
		"testTrafficRouteInput",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) SetTargetGroup(val *[]*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroup",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) PutProdTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute) {
	_jsii_.InvokeVoid(
		c,
		"putProdTrafficRoute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) PutTestTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute) {
	_jsii_.InvokeVoid(
		c,
		"putTestTrafficRoute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ResetTestTrafficRoute() {
	_jsii_.InvokeVoid(
		c,
		"resetTestTrafficRoute",
		nil, // no parameters
	)
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#listener_arns CodedeployDeploymentGroup#listener_arns}.
	ListenerArns *[]*string `json:"listenerArns"`
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ListenerArns() *[]*string
	SetListenerArns(val *[]*string)
	ListenerArnsInput() *[]*string
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

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) ListenerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) ListenerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) SetListenerArns(val *[]*string) {
	_jsii_.Set(
		j,
		"listenerArns",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#name CodedeployDeploymentGroup#name}.
	Name *string `json:"name"`
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#listener_arns CodedeployDeploymentGroup#listener_arns}.
	ListenerArns *[]*string `json:"listenerArns"`
}

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ListenerArns() *[]*string
	SetListenerArns(val *[]*string)
	ListenerArnsInput() *[]*string
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

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) ListenerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) ListenerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) SetListenerArns(val *[]*string) {
	_jsii_.Set(
		j,
		"listenerArns",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodedeployDeploymentGroupOnPremisesInstanceTagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#key CodedeployDeploymentGroup#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#type CodedeployDeploymentGroup#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#value CodedeployDeploymentGroup#value}.
	Value *string `json:"value"`
}

type CodedeployDeploymentGroupTriggerConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#trigger_events CodedeployDeploymentGroup#trigger_events}.
	TriggerEvents *[]*string `json:"triggerEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#trigger_name CodedeployDeploymentGroup#trigger_name}.
	TriggerName *string `json:"triggerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group.html#trigger_target_arn CodedeployDeploymentGroup#trigger_target_arn}.
	TriggerTargetArn *string `json:"triggerTargetArn"`
}
