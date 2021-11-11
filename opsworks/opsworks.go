package opsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/opsworks/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html aws_opsworks_application}.
type OpsworksApplication interface {
	cdktf.TerraformResource
	AppSource() *[]*OpsworksApplicationAppSource
	SetAppSource(val *[]*OpsworksApplicationAppSource)
	AppSourceInput() *[]*OpsworksApplicationAppSource
	AutoBundleOnDeploy() *string
	SetAutoBundleOnDeploy(val *string)
	AutoBundleOnDeployInput() *string
	AwsFlowRubySettings() *string
	SetAwsFlowRubySettings(val *string)
	AwsFlowRubySettingsInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DataSourceArn() *string
	SetDataSourceArn(val *string)
	DataSourceArnInput() *string
	DataSourceDatabaseName() *string
	SetDataSourceDatabaseName(val *string)
	DataSourceDatabaseNameInput() *string
	DataSourceType() *string
	SetDataSourceType(val *string)
	DataSourceTypeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentRoot() *string
	SetDocumentRoot(val *string)
	DocumentRootInput() *string
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	EnableSsl() interface{}
	SetEnableSsl(val interface{})
	EnableSslInput() interface{}
	Environment() *[]*OpsworksApplicationEnvironment
	SetEnvironment(val *[]*OpsworksApplicationEnvironment)
	EnvironmentInput() *[]*OpsworksApplicationEnvironment
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
	RailsEnv() *string
	SetRailsEnv(val *string)
	RailsEnvInput() *string
	RawOverrides() interface{}
	ShortName() *string
	SetShortName(val *string)
	ShortNameInput() *string
	SslConfiguration() *[]*OpsworksApplicationSslConfiguration
	SetSslConfiguration(val *[]*OpsworksApplicationSslConfiguration)
	SslConfigurationInput() *[]*OpsworksApplicationSslConfiguration
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAppSource()
	ResetAutoBundleOnDeploy()
	ResetAwsFlowRubySettings()
	ResetDataSourceArn()
	ResetDataSourceDatabaseName()
	ResetDataSourceType()
	ResetDescription()
	ResetDocumentRoot()
	ResetDomains()
	ResetEnableSsl()
	ResetEnvironment()
	ResetOverrideLogicalId()
	ResetRailsEnv()
	ResetShortName()
	ResetSslConfiguration()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksApplication
type jsiiProxy_OpsworksApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksApplication) AppSource() *[]*OpsworksApplicationAppSource {
	var returns *[]*OpsworksApplicationAppSource
	_jsii_.Get(
		j,
		"appSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AppSourceInput() *[]*OpsworksApplicationAppSource {
	var returns *[]*OpsworksApplicationAppSource
	_jsii_.Get(
		j,
		"appSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AutoBundleOnDeploy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoBundleOnDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AutoBundleOnDeployInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoBundleOnDeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AwsFlowRubySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsFlowRubySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AwsFlowRubySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsFlowRubySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DocumentRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DocumentRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnableSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnableSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Environment() *[]*OpsworksApplicationEnvironment {
	var returns *[]*OpsworksApplicationEnvironment
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnvironmentInput() *[]*OpsworksApplicationEnvironment {
	var returns *[]*OpsworksApplicationEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RailsEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"railsEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RailsEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"railsEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ShortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) SslConfiguration() *[]*OpsworksApplicationSslConfiguration {
	var returns *[]*OpsworksApplicationSslConfiguration
	_jsii_.Get(
		j,
		"sslConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) SslConfigurationInput() *[]*OpsworksApplicationSslConfiguration {
	var returns *[]*OpsworksApplicationSslConfiguration
	_jsii_.Get(
		j,
		"sslConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html aws_opsworks_application} Resource.
func NewOpsworksApplication(scope constructs.Construct, id *string, config *OpsworksApplicationConfig) OpsworksApplication {
	_init_.Initialize()

	j := jsiiProxy_OpsworksApplication{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html aws_opsworks_application} Resource.
func NewOpsworksApplication_Override(o OpsworksApplication, scope constructs.Construct, id *string, config *OpsworksApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksApplication",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetAppSource(val *[]*OpsworksApplicationAppSource) {
	_jsii_.Set(
		j,
		"appSource",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetAutoBundleOnDeploy(val *string) {
	_jsii_.Set(
		j,
		"autoBundleOnDeploy",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetAwsFlowRubySettings(val *string) {
	_jsii_.Set(
		j,
		"awsFlowRubySettings",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDataSourceArn(val *string) {
	_jsii_.Set(
		j,
		"dataSourceArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDataSourceDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"dataSourceDatabaseName",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDataSourceType(val *string) {
	_jsii_.Set(
		j,
		"dataSourceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDocumentRoot(val *string) {
	_jsii_.Set(
		j,
		"documentRoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetEnableSsl(val interface{}) {
	_jsii_.Set(
		j,
		"enableSsl",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetEnvironment(val *[]*OpsworksApplicationEnvironment) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetRailsEnv(val *string) {
	_jsii_.Set(
		j,
		"railsEnv",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetShortName(val *string) {
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetSslConfiguration(val *[]*OpsworksApplicationSslConfiguration) {
	_jsii_.Set(
		j,
		"sslConfiguration",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetType(val *string) {
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
func OpsworksApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAppSource() {
	_jsii_.InvokeVoid(
		o,
		"resetAppSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAutoBundleOnDeploy() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoBundleOnDeploy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAwsFlowRubySettings() {
	_jsii_.InvokeVoid(
		o,
		"resetAwsFlowRubySettings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceArn() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceDatabaseName() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceDatabaseName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceType() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDocumentRoot() {
	_jsii_.InvokeVoid(
		o,
		"resetDocumentRoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDomains() {
	_jsii_.InvokeVoid(
		o,
		"resetDomains",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetEnableSsl() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableSsl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetEnvironment() {
	_jsii_.InvokeVoid(
		o,
		"resetEnvironment",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetRailsEnv() {
	_jsii_.InvokeVoid(
		o,
		"resetRailsEnv",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetShortName() {
	_jsii_.InvokeVoid(
		o,
		"resetShortName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetSslConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetSslConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksApplicationAppSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#type OpsworksApplication#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#password OpsworksApplication#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#revision OpsworksApplication#revision}.
	Revision *string `json:"revision"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#ssh_key OpsworksApplication#ssh_key}.
	SshKey *string `json:"sshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#url OpsworksApplication#url}.
	Url *string `json:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#username OpsworksApplication#username}.
	Username *string `json:"username"`
}

type OpsworksApplicationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#name OpsworksApplication#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#stack_id OpsworksApplication#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#type OpsworksApplication#type}.
	Type *string `json:"type"`
	// app_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#app_source OpsworksApplication#app_source}
	AppSource *[]*OpsworksApplicationAppSource `json:"appSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#auto_bundle_on_deploy OpsworksApplication#auto_bundle_on_deploy}.
	AutoBundleOnDeploy *string `json:"autoBundleOnDeploy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#aws_flow_ruby_settings OpsworksApplication#aws_flow_ruby_settings}.
	AwsFlowRubySettings *string `json:"awsFlowRubySettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#data_source_arn OpsworksApplication#data_source_arn}.
	DataSourceArn *string `json:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#data_source_database_name OpsworksApplication#data_source_database_name}.
	DataSourceDatabaseName *string `json:"dataSourceDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#data_source_type OpsworksApplication#data_source_type}.
	DataSourceType *string `json:"dataSourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#description OpsworksApplication#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#document_root OpsworksApplication#document_root}.
	DocumentRoot *string `json:"documentRoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#domains OpsworksApplication#domains}.
	Domains *[]*string `json:"domains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#enable_ssl OpsworksApplication#enable_ssl}.
	EnableSsl interface{} `json:"enableSsl"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#environment OpsworksApplication#environment}
	Environment *[]*OpsworksApplicationEnvironment `json:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#rails_env OpsworksApplication#rails_env}.
	RailsEnv *string `json:"railsEnv"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#short_name OpsworksApplication#short_name}.
	ShortName *string `json:"shortName"`
	// ssl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#ssl_configuration OpsworksApplication#ssl_configuration}
	SslConfiguration *[]*OpsworksApplicationSslConfiguration `json:"sslConfiguration"`
}

type OpsworksApplicationEnvironment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#key OpsworksApplication#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#value OpsworksApplication#value}.
	Value *string `json:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#secure OpsworksApplication#secure}.
	Secure interface{} `json:"secure"`
}

type OpsworksApplicationSslConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#certificate OpsworksApplication#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#private_key OpsworksApplication#private_key}.
	PrivateKey *string `json:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application.html#chain OpsworksApplication#chain}.
	Chain *string `json:"chain"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html aws_opsworks_custom_layer}.
type OpsworksCustomLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksCustomLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksCustomLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksCustomLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ShortName() *string
	SetShortName(val *string)
	ShortNameInput() *string
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksCustomLayer
type jsiiProxy_OpsworksCustomLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksCustomLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) EbsVolume() *[]*OpsworksCustomLayerEbsVolume {
	var returns *[]*OpsworksCustomLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) EbsVolumeInput() *[]*OpsworksCustomLayerEbsVolume {
	var returns *[]*OpsworksCustomLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) ShortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html aws_opsworks_custom_layer} Resource.
func NewOpsworksCustomLayer(scope constructs.Construct, id *string, config *OpsworksCustomLayerConfig) OpsworksCustomLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksCustomLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksCustomLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html aws_opsworks_custom_layer} Resource.
func NewOpsworksCustomLayer_Override(o OpsworksCustomLayer, scope constructs.Construct, id *string, config *OpsworksCustomLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksCustomLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetEbsVolume(val *[]*OpsworksCustomLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetShortName(val *string) {
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksCustomLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksCustomLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksCustomLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksCustomLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksCustomLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksCustomLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#name OpsworksCustomLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#short_name OpsworksCustomLayer#short_name}.
	ShortName *string `json:"shortName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#stack_id OpsworksCustomLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#auto_assign_elastic_ips OpsworksCustomLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#auto_assign_public_ips OpsworksCustomLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#auto_healing OpsworksCustomLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_configure_recipes OpsworksCustomLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_deploy_recipes OpsworksCustomLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_instance_profile_arn OpsworksCustomLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_json OpsworksCustomLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_security_group_ids OpsworksCustomLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_setup_recipes OpsworksCustomLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_shutdown_recipes OpsworksCustomLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#custom_undeploy_recipes OpsworksCustomLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#drain_elb_on_shutdown OpsworksCustomLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#ebs_volume OpsworksCustomLayer#ebs_volume}
	EbsVolume *[]*OpsworksCustomLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#elastic_load_balancer OpsworksCustomLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#install_updates_on_boot OpsworksCustomLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#instance_shutdown_timeout OpsworksCustomLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#system_packages OpsworksCustomLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#tags OpsworksCustomLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#tags_all OpsworksCustomLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#use_ebs_optimized_instances OpsworksCustomLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksCustomLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#mount_point OpsworksCustomLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#number_of_disks OpsworksCustomLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#size OpsworksCustomLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#encrypted OpsworksCustomLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#iops OpsworksCustomLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#raid_level OpsworksCustomLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer.html#type OpsworksCustomLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html aws_opsworks_ganglia_layer}.
type OpsworksGangliaLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksGangliaLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksGangliaLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksGangliaLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUrl()
	ResetUseEbsOptimizedInstances()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksGangliaLayer
type jsiiProxy_OpsworksGangliaLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksGangliaLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) EbsVolume() *[]*OpsworksGangliaLayerEbsVolume {
	var returns *[]*OpsworksGangliaLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) EbsVolumeInput() *[]*OpsworksGangliaLayerEbsVolume {
	var returns *[]*OpsworksGangliaLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html aws_opsworks_ganglia_layer} Resource.
func NewOpsworksGangliaLayer(scope constructs.Construct, id *string, config *OpsworksGangliaLayerConfig) OpsworksGangliaLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksGangliaLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksGangliaLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html aws_opsworks_ganglia_layer} Resource.
func NewOpsworksGangliaLayer_Override(o OpsworksGangliaLayer, scope constructs.Construct, id *string, config *OpsworksGangliaLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksGangliaLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetEbsVolume(val *[]*OpsworksGangliaLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksGangliaLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksGangliaLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksGangliaLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksGangliaLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) ResetUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksGangliaLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksGangliaLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#password OpsworksGangliaLayer#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#stack_id OpsworksGangliaLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#auto_assign_elastic_ips OpsworksGangliaLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#auto_assign_public_ips OpsworksGangliaLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#auto_healing OpsworksGangliaLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_configure_recipes OpsworksGangliaLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_deploy_recipes OpsworksGangliaLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_instance_profile_arn OpsworksGangliaLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_json OpsworksGangliaLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_security_group_ids OpsworksGangliaLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_setup_recipes OpsworksGangliaLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_shutdown_recipes OpsworksGangliaLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#custom_undeploy_recipes OpsworksGangliaLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#drain_elb_on_shutdown OpsworksGangliaLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#ebs_volume OpsworksGangliaLayer#ebs_volume}
	EbsVolume *[]*OpsworksGangliaLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#elastic_load_balancer OpsworksGangliaLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#install_updates_on_boot OpsworksGangliaLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#instance_shutdown_timeout OpsworksGangliaLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#name OpsworksGangliaLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#system_packages OpsworksGangliaLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#tags OpsworksGangliaLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#tags_all OpsworksGangliaLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#url OpsworksGangliaLayer#url}.
	Url *string `json:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#use_ebs_optimized_instances OpsworksGangliaLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#username OpsworksGangliaLayer#username}.
	Username *string `json:"username"`
}

type OpsworksGangliaLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#mount_point OpsworksGangliaLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#number_of_disks OpsworksGangliaLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#size OpsworksGangliaLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#encrypted OpsworksGangliaLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#iops OpsworksGangliaLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#raid_level OpsworksGangliaLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer.html#type OpsworksGangliaLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html aws_opsworks_haproxy_layer}.
type OpsworksHaproxyLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksHaproxyLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksHaproxyLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksHaproxyLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HealthcheckMethod() *string
	SetHealthcheckMethod(val *string)
	HealthcheckMethodInput() *string
	HealthcheckUrl() *string
	SetHealthcheckUrl(val *string)
	HealthcheckUrlInput() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	StatsEnabled() interface{}
	SetStatsEnabled(val interface{})
	StatsEnabledInput() interface{}
	StatsPassword() *string
	SetStatsPassword(val *string)
	StatsPasswordInput() *string
	StatsUrl() *string
	SetStatsUrl(val *string)
	StatsUrlInput() *string
	StatsUser() *string
	SetStatsUser(val *string)
	StatsUserInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetHealthcheckMethod()
	ResetHealthcheckUrl()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetStatsEnabled()
	ResetStatsUrl()
	ResetStatsUser()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksHaproxyLayer
type jsiiProxy_OpsworksHaproxyLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) EbsVolume() *[]*OpsworksHaproxyLayerEbsVolume {
	var returns *[]*OpsworksHaproxyLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) EbsVolumeInput() *[]*OpsworksHaproxyLayerEbsVolume {
	var returns *[]*OpsworksHaproxyLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) HealthcheckMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcheckMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) HealthcheckMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcheckMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) HealthcheckUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcheckUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) HealthcheckUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcheckUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) StatsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html aws_opsworks_haproxy_layer} Resource.
func NewOpsworksHaproxyLayer(scope constructs.Construct, id *string, config *OpsworksHaproxyLayerConfig) OpsworksHaproxyLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksHaproxyLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksHaproxyLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html aws_opsworks_haproxy_layer} Resource.
func NewOpsworksHaproxyLayer_Override(o OpsworksHaproxyLayer, scope constructs.Construct, id *string, config *OpsworksHaproxyLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksHaproxyLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetEbsVolume(val *[]*OpsworksHaproxyLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetHealthcheckMethod(val *string) {
	_jsii_.Set(
		j,
		"healthcheckMethod",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetHealthcheckUrl(val *string) {
	_jsii_.Set(
		j,
		"healthcheckUrl",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetStatsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"statsEnabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetStatsPassword(val *string) {
	_jsii_.Set(
		j,
		"statsPassword",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetStatsUrl(val *string) {
	_jsii_.Set(
		j,
		"statsUrl",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetStatsUser(val *string) {
	_jsii_.Set(
		j,
		"statsUser",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksHaproxyLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksHaproxyLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksHaproxyLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksHaproxyLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetHealthcheckMethod() {
	_jsii_.InvokeVoid(
		o,
		"resetHealthcheckMethod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetHealthcheckUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetHealthcheckUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetStatsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetStatsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetStatsUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetStatsUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetStatsUser() {
	_jsii_.InvokeVoid(
		o,
		"resetStatsUser",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksHaproxyLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksHaproxyLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#stack_id OpsworksHaproxyLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#stats_password OpsworksHaproxyLayer#stats_password}.
	StatsPassword *string `json:"statsPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#auto_assign_elastic_ips OpsworksHaproxyLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#auto_assign_public_ips OpsworksHaproxyLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#auto_healing OpsworksHaproxyLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_configure_recipes OpsworksHaproxyLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_deploy_recipes OpsworksHaproxyLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_instance_profile_arn OpsworksHaproxyLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_json OpsworksHaproxyLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_security_group_ids OpsworksHaproxyLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_setup_recipes OpsworksHaproxyLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_shutdown_recipes OpsworksHaproxyLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#custom_undeploy_recipes OpsworksHaproxyLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#drain_elb_on_shutdown OpsworksHaproxyLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#ebs_volume OpsworksHaproxyLayer#ebs_volume}
	EbsVolume *[]*OpsworksHaproxyLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#elastic_load_balancer OpsworksHaproxyLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#healthcheck_method OpsworksHaproxyLayer#healthcheck_method}.
	HealthcheckMethod *string `json:"healthcheckMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#healthcheck_url OpsworksHaproxyLayer#healthcheck_url}.
	HealthcheckUrl *string `json:"healthcheckUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#install_updates_on_boot OpsworksHaproxyLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#instance_shutdown_timeout OpsworksHaproxyLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#name OpsworksHaproxyLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#stats_enabled OpsworksHaproxyLayer#stats_enabled}.
	StatsEnabled interface{} `json:"statsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#stats_url OpsworksHaproxyLayer#stats_url}.
	StatsUrl *string `json:"statsUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#stats_user OpsworksHaproxyLayer#stats_user}.
	StatsUser *string `json:"statsUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#system_packages OpsworksHaproxyLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#tags OpsworksHaproxyLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#tags_all OpsworksHaproxyLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#use_ebs_optimized_instances OpsworksHaproxyLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksHaproxyLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#mount_point OpsworksHaproxyLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#number_of_disks OpsworksHaproxyLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#size OpsworksHaproxyLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#encrypted OpsworksHaproxyLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#iops OpsworksHaproxyLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#raid_level OpsworksHaproxyLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer.html#type OpsworksHaproxyLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html aws_opsworks_instance}.
type OpsworksInstance interface {
	cdktf.TerraformResource
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	AmiId() *string
	SetAmiId(val *string)
	AmiIdInput() *string
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	AutoScalingType() *string
	SetAutoScalingType(val *string)
	AutoScalingTypeInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedAt() *string
	SetCreatedAt(val *string)
	CreatedAtInput() *string
	DeleteEbs() interface{}
	SetDeleteEbs(val interface{})
	DeleteEbsInput() interface{}
	DeleteEip() interface{}
	SetDeleteEip(val interface{})
	DeleteEipInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EbsBlockDevice() *[]*OpsworksInstanceEbsBlockDevice
	SetEbsBlockDevice(val *[]*OpsworksInstanceEbsBlockDevice)
	EbsBlockDeviceInput() *[]*OpsworksInstanceEbsBlockDevice
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	Ec2InstanceId() *string
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	EcsClusterArnInput() *string
	ElasticIp() *string
	SetElasticIp(val *string)
	ElasticIpInput() *string
	EphemeralBlockDevice() *[]*OpsworksInstanceEphemeralBlockDevice
	SetEphemeralBlockDevice(val *[]*OpsworksInstanceEphemeralBlockDevice)
	EphemeralBlockDeviceInput() *[]*OpsworksInstanceEphemeralBlockDevice
	Fqn() *string
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	InfrastructureClass() *string
	SetInfrastructureClass(val *string)
	InfrastructureClassInput() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	LastServiceErrorId() *string
	SetLastServiceErrorId(val *string)
	LastServiceErrorIdInput() *string
	LayerIds() *[]*string
	SetLayerIds(val *[]*string)
	LayerIdsInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Os() *string
	SetOs(val *string)
	OsInput() *string
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
	PrivateDns() *string
	SetPrivateDns(val *string)
	PrivateDnsInput() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PublicDns() *string
	SetPublicDns(val *string)
	PublicDnsInput() *string
	PublicIp() *string
	SetPublicIp(val *string)
	PublicIpInput() *string
	RawOverrides() interface{}
	RegisteredBy() *string
	SetRegisteredBy(val *string)
	RegisteredByInput() *string
	ReportedAgentVersion() *string
	SetReportedAgentVersion(val *string)
	ReportedAgentVersionInput() *string
	ReportedOsFamily() *string
	SetReportedOsFamily(val *string)
	ReportedOsFamilyInput() *string
	ReportedOsName() *string
	SetReportedOsName(val *string)
	ReportedOsNameInput() *string
	ReportedOsVersion() *string
	SetReportedOsVersion(val *string)
	ReportedOsVersionInput() *string
	RootBlockDevice() *[]*OpsworksInstanceRootBlockDevice
	SetRootBlockDevice(val *[]*OpsworksInstanceRootBlockDevice)
	RootBlockDeviceInput() *[]*OpsworksInstanceRootBlockDevice
	RootDeviceType() *string
	SetRootDeviceType(val *string)
	RootDeviceTypeInput() *string
	RootDeviceVolumeId() *string
	SetRootDeviceVolumeId(val *string)
	RootDeviceVolumeIdInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshHostDsaKeyFingerprint() *string
	SetSshHostDsaKeyFingerprint(val *string)
	SshHostDsaKeyFingerprintInput() *string
	SshHostRsaKeyFingerprint() *string
	SetSshHostRsaKeyFingerprint(val *string)
	SshHostRsaKeyFingerprintInput() *string
	SshKeyName() *string
	SetSshKeyName(val *string)
	SshKeyNameInput() *string
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() OpsworksInstanceTimeoutsOutputReference
	TimeoutsInput() *OpsworksInstanceTimeouts
	VirtualizationType() *string
	SetVirtualizationType(val *string)
	VirtualizationTypeInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *OpsworksInstanceTimeouts)
	ResetAgentVersion()
	ResetAmiId()
	ResetArchitecture()
	ResetAutoScalingType()
	ResetAvailabilityZone()
	ResetCreatedAt()
	ResetDeleteEbs()
	ResetDeleteEip()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetEcsClusterArn()
	ResetElasticIp()
	ResetEphemeralBlockDevice()
	ResetHostname()
	ResetInfrastructureClass()
	ResetInstallUpdatesOnBoot()
	ResetInstanceProfileArn()
	ResetInstanceType()
	ResetLastServiceErrorId()
	ResetOs()
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetPrivateDns()
	ResetPrivateIp()
	ResetPublicDns()
	ResetPublicIp()
	ResetRegisteredBy()
	ResetReportedAgentVersion()
	ResetReportedOsFamily()
	ResetReportedOsName()
	ResetReportedOsVersion()
	ResetRootBlockDevice()
	ResetRootDeviceType()
	ResetRootDeviceVolumeId()
	ResetSecurityGroupIds()
	ResetSshHostDsaKeyFingerprint()
	ResetSshHostRsaKeyFingerprint()
	ResetSshKeyName()
	ResetState()
	ResetStatus()
	ResetSubnetId()
	ResetTenancy()
	ResetTimeouts()
	ResetVirtualizationType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksInstance
type jsiiProxy_OpsworksInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksInstance) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AutoScalingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AutoScalingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CreatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEbs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsBlockDevice() *[]*OpsworksInstanceEbsBlockDevice {
	var returns *[]*OpsworksInstanceEbsBlockDevice
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsBlockDeviceInput() *[]*OpsworksInstanceEbsBlockDevice {
	var returns *[]*OpsworksInstanceEbsBlockDevice
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Ec2InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EcsClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDevice() *[]*OpsworksInstanceEphemeralBlockDevice {
	var returns *[]*OpsworksInstanceEphemeralBlockDevice
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDeviceInput() *[]*OpsworksInstanceEphemeralBlockDevice {
	var returns *[]*OpsworksInstanceEphemeralBlockDevice
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InfrastructureClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InfrastructureClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LastServiceErrorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastServiceErrorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LastServiceErrorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastServiceErrorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LayerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LayerIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateDnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicDnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RegisteredBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RegisteredByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedAgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedAgentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootBlockDevice() *[]*OpsworksInstanceRootBlockDevice {
	var returns *[]*OpsworksInstanceRootBlockDevice
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootBlockDeviceInput() *[]*OpsworksInstanceRootBlockDevice {
	var returns *[]*OpsworksInstanceRootBlockDevice
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceVolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceVolumeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostDsaKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostDsaKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostDsaKeyFingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostDsaKeyFingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostRsaKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostRsaKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostRsaKeyFingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostRsaKeyFingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Timeouts() OpsworksInstanceTimeoutsOutputReference {
	var returns OpsworksInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TimeoutsInput() *OpsworksInstanceTimeouts {
	var returns *OpsworksInstanceTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) VirtualizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationTypeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html aws_opsworks_instance} Resource.
func NewOpsworksInstance(scope constructs.Construct, id *string, config *OpsworksInstanceConfig) OpsworksInstance {
	_init_.Initialize()

	j := jsiiProxy_OpsworksInstance{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html aws_opsworks_instance} Resource.
func NewOpsworksInstance_Override(o OpsworksInstance, scope constructs.Construct, id *string, config *OpsworksInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksInstance",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetAmiId(val *string) {
	_jsii_.Set(
		j,
		"amiId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetAutoScalingType(val *string) {
	_jsii_.Set(
		j,
		"autoScalingType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetCreatedAt(val *string) {
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetDeleteEbs(val interface{}) {
	_jsii_.Set(
		j,
		"deleteEbs",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetDeleteEip(val interface{}) {
	_jsii_.Set(
		j,
		"deleteEip",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetEbsBlockDevice(val *[]*OpsworksInstanceEbsBlockDevice) {
	_jsii_.Set(
		j,
		"ebsBlockDevice",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetEbsOptimized(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetEcsClusterArn(val *string) {
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetElasticIp(val *string) {
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetEphemeralBlockDevice(val *[]*OpsworksInstanceEphemeralBlockDevice) {
	_jsii_.Set(
		j,
		"ephemeralBlockDevice",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetInfrastructureClass(val *string) {
	_jsii_.Set(
		j,
		"infrastructureClass",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetLastServiceErrorId(val *string) {
	_jsii_.Set(
		j,
		"lastServiceErrorId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetLayerIds(val *[]*string) {
	_jsii_.Set(
		j,
		"layerIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetOs(val *string) {
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetPlatform(val *string) {
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetPrivateDns(val *string) {
	_jsii_.Set(
		j,
		"privateDns",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetPrivateIp(val *string) {
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetPublicDns(val *string) {
	_jsii_.Set(
		j,
		"publicDns",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetPublicIp(val *string) {
	_jsii_.Set(
		j,
		"publicIp",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetRegisteredBy(val *string) {
	_jsii_.Set(
		j,
		"registeredBy",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetReportedAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"reportedAgentVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetReportedOsFamily(val *string) {
	_jsii_.Set(
		j,
		"reportedOsFamily",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetReportedOsName(val *string) {
	_jsii_.Set(
		j,
		"reportedOsName",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetReportedOsVersion(val *string) {
	_jsii_.Set(
		j,
		"reportedOsVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetRootBlockDevice(val *[]*OpsworksInstanceRootBlockDevice) {
	_jsii_.Set(
		j,
		"rootBlockDevice",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"rootDeviceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetRootDeviceVolumeId(val *string) {
	_jsii_.Set(
		j,
		"rootDeviceVolumeId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetSshHostDsaKeyFingerprint(val *string) {
	_jsii_.Set(
		j,
		"sshHostDsaKeyFingerprint",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetSshHostRsaKeyFingerprint(val *string) {
	_jsii_.Set(
		j,
		"sshHostRsaKeyFingerprint",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"sshKeyName",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetTenancy(val *string) {
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance) SetVirtualizationType(val *string) {
	_jsii_.Set(
		j,
		"virtualizationType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksInstance) PutTimeouts(value *OpsworksInstanceTimeouts) {
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAgentVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAgentVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAmiId() {
	_jsii_.InvokeVoid(
		o,
		"resetAmiId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetArchitecture() {
	_jsii_.InvokeVoid(
		o,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAutoScalingType() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoScalingType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		o,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		o,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetDeleteEbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDeleteEbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetDeleteEip() {
	_jsii_.InvokeVoid(
		o,
		"resetDeleteEip",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEcsClusterArn() {
	_jsii_.InvokeVoid(
		o,
		"resetEcsClusterArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetElasticIp() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetHostname() {
	_jsii_.InvokeVoid(
		o,
		"resetHostname",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInfrastructureClass() {
	_jsii_.InvokeVoid(
		o,
		"resetInfrastructureClass",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetLastServiceErrorId() {
	_jsii_.InvokeVoid(
		o,
		"resetLastServiceErrorId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetOs() {
	_jsii_.InvokeVoid(
		o,
		"resetOs",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetPlatform() {
	_jsii_.InvokeVoid(
		o,
		"resetPlatform",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetPrivateDns() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateDns",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetPublicDns() {
	_jsii_.InvokeVoid(
		o,
		"resetPublicDns",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetPublicIp() {
	_jsii_.InvokeVoid(
		o,
		"resetPublicIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRegisteredBy() {
	_jsii_.InvokeVoid(
		o,
		"resetRegisteredBy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetReportedAgentVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetReportedAgentVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetReportedOsFamily() {
	_jsii_.InvokeVoid(
		o,
		"resetReportedOsFamily",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetReportedOsName() {
	_jsii_.InvokeVoid(
		o,
		"resetReportedOsName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetReportedOsVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetReportedOsVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRootDeviceType() {
	_jsii_.InvokeVoid(
		o,
		"resetRootDeviceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRootDeviceVolumeId() {
	_jsii_.InvokeVoid(
		o,
		"resetRootDeviceVolumeId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSshHostDsaKeyFingerprint() {
	_jsii_.InvokeVoid(
		o,
		"resetSshHostDsaKeyFingerprint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSshHostRsaKeyFingerprint() {
	_jsii_.InvokeVoid(
		o,
		"resetSshHostRsaKeyFingerprint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSshKeyName() {
	_jsii_.InvokeVoid(
		o,
		"resetSshKeyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetState() {
	_jsii_.InvokeVoid(
		o,
		"resetState",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetStatus() {
	_jsii_.InvokeVoid(
		o,
		"resetStatus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSubnetId() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetTenancy() {
	_jsii_.InvokeVoid(
		o,
		"resetTenancy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetVirtualizationType() {
	_jsii_.InvokeVoid(
		o,
		"resetVirtualizationType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksInstanceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#layer_ids OpsworksInstance#layer_ids}.
	LayerIds *[]*string `json:"layerIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#stack_id OpsworksInstance#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#agent_version OpsworksInstance#agent_version}.
	AgentVersion *string `json:"agentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ami_id OpsworksInstance#ami_id}.
	AmiId *string `json:"amiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#architecture OpsworksInstance#architecture}.
	Architecture *string `json:"architecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#auto_scaling_type OpsworksInstance#auto_scaling_type}.
	AutoScalingType *string `json:"autoScalingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#availability_zone OpsworksInstance#availability_zone}.
	AvailabilityZone *string `json:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#created_at OpsworksInstance#created_at}.
	CreatedAt *string `json:"createdAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#delete_ebs OpsworksInstance#delete_ebs}.
	DeleteEbs interface{} `json:"deleteEbs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#delete_eip OpsworksInstance#delete_eip}.
	DeleteEip interface{} `json:"deleteEip"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ebs_block_device OpsworksInstance#ebs_block_device}
	EbsBlockDevice *[]*OpsworksInstanceEbsBlockDevice `json:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ebs_optimized OpsworksInstance#ebs_optimized}.
	EbsOptimized interface{} `json:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ecs_cluster_arn OpsworksInstance#ecs_cluster_arn}.
	EcsClusterArn *string `json:"ecsClusterArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#elastic_ip OpsworksInstance#elastic_ip}.
	ElasticIp *string `json:"elasticIp"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ephemeral_block_device OpsworksInstance#ephemeral_block_device}
	EphemeralBlockDevice *[]*OpsworksInstanceEphemeralBlockDevice `json:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#hostname OpsworksInstance#hostname}.
	Hostname *string `json:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#infrastructure_class OpsworksInstance#infrastructure_class}.
	InfrastructureClass *string `json:"infrastructureClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#install_updates_on_boot OpsworksInstance#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#instance_profile_arn OpsworksInstance#instance_profile_arn}.
	InstanceProfileArn *string `json:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#instance_type OpsworksInstance#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#last_service_error_id OpsworksInstance#last_service_error_id}.
	LastServiceErrorId *string `json:"lastServiceErrorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#os OpsworksInstance#os}.
	Os *string `json:"os"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#platform OpsworksInstance#platform}.
	Platform *string `json:"platform"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#private_dns OpsworksInstance#private_dns}.
	PrivateDns *string `json:"privateDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#private_ip OpsworksInstance#private_ip}.
	PrivateIp *string `json:"privateIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#public_dns OpsworksInstance#public_dns}.
	PublicDns *string `json:"publicDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#public_ip OpsworksInstance#public_ip}.
	PublicIp *string `json:"publicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#registered_by OpsworksInstance#registered_by}.
	RegisteredBy *string `json:"registeredBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#reported_agent_version OpsworksInstance#reported_agent_version}.
	ReportedAgentVersion *string `json:"reportedAgentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#reported_os_family OpsworksInstance#reported_os_family}.
	ReportedOsFamily *string `json:"reportedOsFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#reported_os_name OpsworksInstance#reported_os_name}.
	ReportedOsName *string `json:"reportedOsName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#reported_os_version OpsworksInstance#reported_os_version}.
	ReportedOsVersion *string `json:"reportedOsVersion"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#root_block_device OpsworksInstance#root_block_device}
	RootBlockDevice *[]*OpsworksInstanceRootBlockDevice `json:"rootBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#root_device_type OpsworksInstance#root_device_type}.
	RootDeviceType *string `json:"rootDeviceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#root_device_volume_id OpsworksInstance#root_device_volume_id}.
	RootDeviceVolumeId *string `json:"rootDeviceVolumeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#security_group_ids OpsworksInstance#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ssh_host_dsa_key_fingerprint OpsworksInstance#ssh_host_dsa_key_fingerprint}.
	SshHostDsaKeyFingerprint *string `json:"sshHostDsaKeyFingerprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ssh_host_rsa_key_fingerprint OpsworksInstance#ssh_host_rsa_key_fingerprint}.
	SshHostRsaKeyFingerprint *string `json:"sshHostRsaKeyFingerprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#ssh_key_name OpsworksInstance#ssh_key_name}.
	SshKeyName *string `json:"sshKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#state OpsworksInstance#state}.
	State *string `json:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#status OpsworksInstance#status}.
	Status *string `json:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#subnet_id OpsworksInstance#subnet_id}.
	SubnetId *string `json:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#tenancy OpsworksInstance#tenancy}.
	Tenancy *string `json:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#timeouts OpsworksInstance#timeouts}
	Timeouts *OpsworksInstanceTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#virtualization_type OpsworksInstance#virtualization_type}.
	VirtualizationType *string `json:"virtualizationType"`
}

type OpsworksInstanceEbsBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#device_name OpsworksInstance#device_name}.
	DeviceName *string `json:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#delete_on_termination OpsworksInstance#delete_on_termination}.
	DeleteOnTermination interface{} `json:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#iops OpsworksInstance#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#snapshot_id OpsworksInstance#snapshot_id}.
	SnapshotId *string `json:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#volume_size OpsworksInstance#volume_size}.
	VolumeSize *float64 `json:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#volume_type OpsworksInstance#volume_type}.
	VolumeType *string `json:"volumeType"`
}

type OpsworksInstanceEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#device_name OpsworksInstance#device_name}.
	DeviceName *string `json:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#virtual_name OpsworksInstance#virtual_name}.
	VirtualName *string `json:"virtualName"`
}

type OpsworksInstanceRootBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#delete_on_termination OpsworksInstance#delete_on_termination}.
	DeleteOnTermination interface{} `json:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#iops OpsworksInstance#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#volume_size OpsworksInstance#volume_size}.
	VolumeSize *float64 `json:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#volume_type OpsworksInstance#volume_type}.
	VolumeType *string `json:"volumeType"`
}

type OpsworksInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#create OpsworksInstance#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#delete OpsworksInstance#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance.html#update OpsworksInstance#update}.
	Update *string `json:"update"`
}

type OpsworksInstanceTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for OpsworksInstanceTimeoutsOutputReference
type jsiiProxy_OpsworksInstanceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewOpsworksInstanceTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) OpsworksInstanceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksInstanceTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksInstanceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksInstanceTimeoutsOutputReference_Override(o OpsworksInstanceTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksInstanceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		o,
		"resetCreate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		o,
		"resetDelete",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html aws_opsworks_java_app_layer}.
type OpsworksJavaAppLayer interface {
	cdktf.TerraformResource
	AppServer() *string
	SetAppServer(val *string)
	AppServerInput() *string
	AppServerVersion() *string
	SetAppServerVersion(val *string)
	AppServerVersionInput() *string
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksJavaAppLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksJavaAppLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksJavaAppLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	JvmOptions() *string
	SetJvmOptions(val *string)
	JvmOptionsInput() *string
	JvmType() *string
	SetJvmType(val *string)
	JvmTypeInput() *string
	JvmVersion() *string
	SetJvmVersion(val *string)
	JvmVersionInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAppServer()
	ResetAppServerVersion()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetJvmOptions()
	ResetJvmType()
	ResetJvmVersion()
	ResetName()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksJavaAppLayer
type jsiiProxy_OpsworksJavaAppLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AppServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AppServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AppServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AppServerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) EbsVolume() *[]*OpsworksJavaAppLayerEbsVolume {
	var returns *[]*OpsworksJavaAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) EbsVolumeInput() *[]*OpsworksJavaAppLayerEbsVolume {
	var returns *[]*OpsworksJavaAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) JvmVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html aws_opsworks_java_app_layer} Resource.
func NewOpsworksJavaAppLayer(scope constructs.Construct, id *string, config *OpsworksJavaAppLayerConfig) OpsworksJavaAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksJavaAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksJavaAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html aws_opsworks_java_app_layer} Resource.
func NewOpsworksJavaAppLayer_Override(o OpsworksJavaAppLayer, scope constructs.Construct, id *string, config *OpsworksJavaAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksJavaAppLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetAppServer(val *string) {
	_jsii_.Set(
		j,
		"appServer",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetAppServerVersion(val *string) {
	_jsii_.Set(
		j,
		"appServerVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetEbsVolume(val *[]*OpsworksJavaAppLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetJvmOptions(val *string) {
	_jsii_.Set(
		j,
		"jvmOptions",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetJvmType(val *string) {
	_jsii_.Set(
		j,
		"jvmType",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetJvmVersion(val *string) {
	_jsii_.Set(
		j,
		"jvmVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksJavaAppLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksJavaAppLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksJavaAppLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksJavaAppLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetAppServer() {
	_jsii_.InvokeVoid(
		o,
		"resetAppServer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetAppServerVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAppServerVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetJvmOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetJvmOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetJvmType() {
	_jsii_.InvokeVoid(
		o,
		"resetJvmType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetJvmVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetJvmVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksJavaAppLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksJavaAppLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#stack_id OpsworksJavaAppLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#app_server OpsworksJavaAppLayer#app_server}.
	AppServer *string `json:"appServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#app_server_version OpsworksJavaAppLayer#app_server_version}.
	AppServerVersion *string `json:"appServerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#auto_assign_elastic_ips OpsworksJavaAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#auto_assign_public_ips OpsworksJavaAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#auto_healing OpsworksJavaAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_configure_recipes OpsworksJavaAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_deploy_recipes OpsworksJavaAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_instance_profile_arn OpsworksJavaAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_json OpsworksJavaAppLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_security_group_ids OpsworksJavaAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_setup_recipes OpsworksJavaAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_shutdown_recipes OpsworksJavaAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#custom_undeploy_recipes OpsworksJavaAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#drain_elb_on_shutdown OpsworksJavaAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#ebs_volume OpsworksJavaAppLayer#ebs_volume}
	EbsVolume *[]*OpsworksJavaAppLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#elastic_load_balancer OpsworksJavaAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#install_updates_on_boot OpsworksJavaAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#instance_shutdown_timeout OpsworksJavaAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#jvm_options OpsworksJavaAppLayer#jvm_options}.
	JvmOptions *string `json:"jvmOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#jvm_type OpsworksJavaAppLayer#jvm_type}.
	JvmType *string `json:"jvmType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#jvm_version OpsworksJavaAppLayer#jvm_version}.
	JvmVersion *string `json:"jvmVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#name OpsworksJavaAppLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#system_packages OpsworksJavaAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#tags OpsworksJavaAppLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#tags_all OpsworksJavaAppLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#use_ebs_optimized_instances OpsworksJavaAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksJavaAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#mount_point OpsworksJavaAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#number_of_disks OpsworksJavaAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#size OpsworksJavaAppLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#encrypted OpsworksJavaAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#iops OpsworksJavaAppLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#raid_level OpsworksJavaAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer.html#type OpsworksJavaAppLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html aws_opsworks_memcached_layer}.
type OpsworksMemcachedLayer interface {
	cdktf.TerraformResource
	AllocatedMemory() *float64
	SetAllocatedMemory(val *float64)
	AllocatedMemoryInput() *float64
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksMemcachedLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksMemcachedLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksMemcachedLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAllocatedMemory()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksMemcachedLayer
type jsiiProxy_OpsworksMemcachedLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AllocatedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AllocatedMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) EbsVolume() *[]*OpsworksMemcachedLayerEbsVolume {
	var returns *[]*OpsworksMemcachedLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) EbsVolumeInput() *[]*OpsworksMemcachedLayerEbsVolume {
	var returns *[]*OpsworksMemcachedLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html aws_opsworks_memcached_layer} Resource.
func NewOpsworksMemcachedLayer(scope constructs.Construct, id *string, config *OpsworksMemcachedLayerConfig) OpsworksMemcachedLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMemcachedLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksMemcachedLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html aws_opsworks_memcached_layer} Resource.
func NewOpsworksMemcachedLayer_Override(o OpsworksMemcachedLayer, scope constructs.Construct, id *string, config *OpsworksMemcachedLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksMemcachedLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetAllocatedMemory(val *float64) {
	_jsii_.Set(
		j,
		"allocatedMemory",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetEbsVolume(val *[]*OpsworksMemcachedLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksMemcachedLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksMemcachedLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksMemcachedLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksMemcachedLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetAllocatedMemory() {
	_jsii_.InvokeVoid(
		o,
		"resetAllocatedMemory",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksMemcachedLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksMemcachedLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#stack_id OpsworksMemcachedLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#allocated_memory OpsworksMemcachedLayer#allocated_memory}.
	AllocatedMemory *float64 `json:"allocatedMemory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#auto_assign_elastic_ips OpsworksMemcachedLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#auto_assign_public_ips OpsworksMemcachedLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#auto_healing OpsworksMemcachedLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_configure_recipes OpsworksMemcachedLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_deploy_recipes OpsworksMemcachedLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_instance_profile_arn OpsworksMemcachedLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_json OpsworksMemcachedLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_security_group_ids OpsworksMemcachedLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_setup_recipes OpsworksMemcachedLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_shutdown_recipes OpsworksMemcachedLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#custom_undeploy_recipes OpsworksMemcachedLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#drain_elb_on_shutdown OpsworksMemcachedLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#ebs_volume OpsworksMemcachedLayer#ebs_volume}
	EbsVolume *[]*OpsworksMemcachedLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#elastic_load_balancer OpsworksMemcachedLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#install_updates_on_boot OpsworksMemcachedLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#instance_shutdown_timeout OpsworksMemcachedLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#name OpsworksMemcachedLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#system_packages OpsworksMemcachedLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#tags OpsworksMemcachedLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#tags_all OpsworksMemcachedLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#use_ebs_optimized_instances OpsworksMemcachedLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksMemcachedLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#mount_point OpsworksMemcachedLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#number_of_disks OpsworksMemcachedLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#size OpsworksMemcachedLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#encrypted OpsworksMemcachedLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#iops OpsworksMemcachedLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#raid_level OpsworksMemcachedLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer.html#type OpsworksMemcachedLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html aws_opsworks_mysql_layer}.
type OpsworksMysqlLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksMysqlLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksMysqlLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksMysqlLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootPassword() *string
	SetRootPassword(val *string)
	RootPasswordInput() *string
	RootPasswordOnAllInstances() interface{}
	SetRootPasswordOnAllInstances(val interface{})
	RootPasswordOnAllInstancesInput() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetRootPassword()
	ResetRootPasswordOnAllInstances()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksMysqlLayer
type jsiiProxy_OpsworksMysqlLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksMysqlLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) EbsVolume() *[]*OpsworksMysqlLayerEbsVolume {
	var returns *[]*OpsworksMysqlLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) EbsVolumeInput() *[]*OpsworksMysqlLayerEbsVolume {
	var returns *[]*OpsworksMysqlLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) RootPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) RootPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) RootPasswordOnAllInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootPasswordOnAllInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) RootPasswordOnAllInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootPasswordOnAllInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html aws_opsworks_mysql_layer} Resource.
func NewOpsworksMysqlLayer(scope constructs.Construct, id *string, config *OpsworksMysqlLayerConfig) OpsworksMysqlLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMysqlLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksMysqlLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html aws_opsworks_mysql_layer} Resource.
func NewOpsworksMysqlLayer_Override(o OpsworksMysqlLayer, scope constructs.Construct, id *string, config *OpsworksMysqlLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksMysqlLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetEbsVolume(val *[]*OpsworksMysqlLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetRootPassword(val *string) {
	_jsii_.Set(
		j,
		"rootPassword",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetRootPasswordOnAllInstances(val interface{}) {
	_jsii_.Set(
		j,
		"rootPasswordOnAllInstances",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksMysqlLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksMysqlLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksMysqlLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksMysqlLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetRootPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetRootPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetRootPasswordOnAllInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetRootPasswordOnAllInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksMysqlLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksMysqlLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#stack_id OpsworksMysqlLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#auto_assign_elastic_ips OpsworksMysqlLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#auto_assign_public_ips OpsworksMysqlLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#auto_healing OpsworksMysqlLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_configure_recipes OpsworksMysqlLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_deploy_recipes OpsworksMysqlLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_instance_profile_arn OpsworksMysqlLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_json OpsworksMysqlLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_security_group_ids OpsworksMysqlLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_setup_recipes OpsworksMysqlLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_shutdown_recipes OpsworksMysqlLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#custom_undeploy_recipes OpsworksMysqlLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#drain_elb_on_shutdown OpsworksMysqlLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#ebs_volume OpsworksMysqlLayer#ebs_volume}
	EbsVolume *[]*OpsworksMysqlLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#elastic_load_balancer OpsworksMysqlLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#install_updates_on_boot OpsworksMysqlLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#instance_shutdown_timeout OpsworksMysqlLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#name OpsworksMysqlLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#root_password OpsworksMysqlLayer#root_password}.
	RootPassword *string `json:"rootPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#root_password_on_all_instances OpsworksMysqlLayer#root_password_on_all_instances}.
	RootPasswordOnAllInstances interface{} `json:"rootPasswordOnAllInstances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#system_packages OpsworksMysqlLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#tags OpsworksMysqlLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#tags_all OpsworksMysqlLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#use_ebs_optimized_instances OpsworksMysqlLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksMysqlLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#mount_point OpsworksMysqlLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#number_of_disks OpsworksMysqlLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#size OpsworksMysqlLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#encrypted OpsworksMysqlLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#iops OpsworksMysqlLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#raid_level OpsworksMysqlLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer.html#type OpsworksMysqlLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html aws_opsworks_nodejs_app_layer}.
type OpsworksNodejsAppLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksNodejsAppLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksNodejsAppLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksNodejsAppLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	NodejsVersion() *string
	SetNodejsVersion(val *string)
	NodejsVersionInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetNodejsVersion()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksNodejsAppLayer
type jsiiProxy_OpsworksNodejsAppLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolume() *[]*OpsworksNodejsAppLayerEbsVolume {
	var returns *[]*OpsworksNodejsAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolumeInput() *[]*OpsworksNodejsAppLayerEbsVolume {
	var returns *[]*OpsworksNodejsAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NodejsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodejsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NodejsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodejsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer(scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) OpsworksNodejsAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksNodejsAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksNodejsAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer_Override(o OpsworksNodejsAppLayer, scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksNodejsAppLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetEbsVolume(val *[]*OpsworksNodejsAppLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetNodejsVersion(val *string) {
	_jsii_.Set(
		j,
		"nodejsVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksNodejsAppLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksNodejsAppLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksNodejsAppLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksNodejsAppLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetNodejsVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetNodejsVersion",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksNodejsAppLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksNodejsAppLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#stack_id OpsworksNodejsAppLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#auto_assign_elastic_ips OpsworksNodejsAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#auto_assign_public_ips OpsworksNodejsAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#auto_healing OpsworksNodejsAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_configure_recipes OpsworksNodejsAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_deploy_recipes OpsworksNodejsAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_instance_profile_arn OpsworksNodejsAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_json OpsworksNodejsAppLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_security_group_ids OpsworksNodejsAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_setup_recipes OpsworksNodejsAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_shutdown_recipes OpsworksNodejsAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#custom_undeploy_recipes OpsworksNodejsAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#drain_elb_on_shutdown OpsworksNodejsAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#ebs_volume OpsworksNodejsAppLayer#ebs_volume}
	EbsVolume *[]*OpsworksNodejsAppLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#elastic_load_balancer OpsworksNodejsAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#install_updates_on_boot OpsworksNodejsAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#instance_shutdown_timeout OpsworksNodejsAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#name OpsworksNodejsAppLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#nodejs_version OpsworksNodejsAppLayer#nodejs_version}.
	NodejsVersion *string `json:"nodejsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#system_packages OpsworksNodejsAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#tags OpsworksNodejsAppLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#tags_all OpsworksNodejsAppLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#use_ebs_optimized_instances OpsworksNodejsAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksNodejsAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#mount_point OpsworksNodejsAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#number_of_disks OpsworksNodejsAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#size OpsworksNodejsAppLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#encrypted OpsworksNodejsAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#iops OpsworksNodejsAppLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#raid_level OpsworksNodejsAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer.html#type OpsworksNodejsAppLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html aws_opsworks_permission}.
type OpsworksPermission interface {
	cdktf.TerraformResource
	AllowSsh() interface{}
	SetAllowSsh(val interface{})
	AllowSshInput() interface{}
	AllowSudo() interface{}
	SetAllowSudo(val interface{})
	AllowSudoInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserArn() *string
	SetUserArn(val *string)
	UserArnInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAllowSsh()
	ResetAllowSudo()
	ResetLevel()
	ResetOverrideLogicalId()
	ResetStackId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksPermission
type jsiiProxy_OpsworksPermission struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksPermission) AllowSsh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) AllowSshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) AllowSudo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSudo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) AllowSudoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSudoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) UserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPermission) UserArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArnInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html aws_opsworks_permission} Resource.
func NewOpsworksPermission(scope constructs.Construct, id *string, config *OpsworksPermissionConfig) OpsworksPermission {
	_init_.Initialize()

	j := jsiiProxy_OpsworksPermission{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksPermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html aws_opsworks_permission} Resource.
func NewOpsworksPermission_Override(o OpsworksPermission, scope constructs.Construct, id *string, config *OpsworksPermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksPermission",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetAllowSsh(val interface{}) {
	_jsii_.Set(
		j,
		"allowSsh",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetAllowSudo(val interface{}) {
	_jsii_.Set(
		j,
		"allowSudo",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetLevel(val *string) {
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksPermission) SetUserArn(val *string) {
	_jsii_.Set(
		j,
		"userArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksPermission_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksPermission",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksPermission) ResetAllowSsh() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowSsh",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPermission) ResetAllowSudo() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowSudo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPermission) ResetLevel() {
	_jsii_.InvokeVoid(
		o,
		"resetLevel",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksPermission) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPermission) ResetStackId() {
	_jsii_.InvokeVoid(
		o,
		"resetStackId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPermission) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksPermission) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksPermissionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html#user_arn OpsworksPermission#user_arn}.
	UserArn *string `json:"userArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html#allow_ssh OpsworksPermission#allow_ssh}.
	AllowSsh interface{} `json:"allowSsh"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html#allow_sudo OpsworksPermission#allow_sudo}.
	AllowSudo interface{} `json:"allowSudo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html#level OpsworksPermission#level}.
	Level *string `json:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission.html#stack_id OpsworksPermission#stack_id}.
	StackId *string `json:"stackId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html aws_opsworks_php_app_layer}.
type OpsworksPhpAppLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksPhpAppLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksPhpAppLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksPhpAppLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksPhpAppLayer
type jsiiProxy_OpsworksPhpAppLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) EbsVolume() *[]*OpsworksPhpAppLayerEbsVolume {
	var returns *[]*OpsworksPhpAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) EbsVolumeInput() *[]*OpsworksPhpAppLayerEbsVolume {
	var returns *[]*OpsworksPhpAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html aws_opsworks_php_app_layer} Resource.
func NewOpsworksPhpAppLayer(scope constructs.Construct, id *string, config *OpsworksPhpAppLayerConfig) OpsworksPhpAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksPhpAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksPhpAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html aws_opsworks_php_app_layer} Resource.
func NewOpsworksPhpAppLayer_Override(o OpsworksPhpAppLayer, scope constructs.Construct, id *string, config *OpsworksPhpAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksPhpAppLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetEbsVolume(val *[]*OpsworksPhpAppLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksPhpAppLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksPhpAppLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksPhpAppLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksPhpAppLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksPhpAppLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksPhpAppLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#stack_id OpsworksPhpAppLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#auto_assign_elastic_ips OpsworksPhpAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#auto_assign_public_ips OpsworksPhpAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#auto_healing OpsworksPhpAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_configure_recipes OpsworksPhpAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_deploy_recipes OpsworksPhpAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_instance_profile_arn OpsworksPhpAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_json OpsworksPhpAppLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_security_group_ids OpsworksPhpAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_setup_recipes OpsworksPhpAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_shutdown_recipes OpsworksPhpAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#custom_undeploy_recipes OpsworksPhpAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#drain_elb_on_shutdown OpsworksPhpAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#ebs_volume OpsworksPhpAppLayer#ebs_volume}
	EbsVolume *[]*OpsworksPhpAppLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#elastic_load_balancer OpsworksPhpAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#install_updates_on_boot OpsworksPhpAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#instance_shutdown_timeout OpsworksPhpAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#name OpsworksPhpAppLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#system_packages OpsworksPhpAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#tags OpsworksPhpAppLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#tags_all OpsworksPhpAppLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#use_ebs_optimized_instances OpsworksPhpAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksPhpAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#mount_point OpsworksPhpAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#number_of_disks OpsworksPhpAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#size OpsworksPhpAppLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#encrypted OpsworksPhpAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#iops OpsworksPhpAppLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#raid_level OpsworksPhpAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer.html#type OpsworksPhpAppLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html aws_opsworks_rails_app_layer}.
type OpsworksRailsAppLayer interface {
	cdktf.TerraformResource
	AppServer() *string
	SetAppServer(val *string)
	AppServerInput() *string
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	BundlerVersion() *string
	SetBundlerVersion(val *string)
	BundlerVersionInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksRailsAppLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksRailsAppLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksRailsAppLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageBundler() interface{}
	SetManageBundler(val interface{})
	ManageBundlerInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PassengerVersion() *string
	SetPassengerVersion(val *string)
	PassengerVersionInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RubygemsVersion() *string
	SetRubygemsVersion(val *string)
	RubygemsVersionInput() *string
	RubyVersion() *string
	SetRubyVersion(val *string)
	RubyVersionInput() *string
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAppServer()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetBundlerVersion()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetManageBundler()
	ResetName()
	ResetOverrideLogicalId()
	ResetPassengerVersion()
	ResetRubygemsVersion()
	ResetRubyVersion()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksRailsAppLayer
type jsiiProxy_OpsworksRailsAppLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AppServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AppServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) BundlerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundlerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) BundlerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundlerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) EbsVolume() *[]*OpsworksRailsAppLayerEbsVolume {
	var returns *[]*OpsworksRailsAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) EbsVolumeInput() *[]*OpsworksRailsAppLayerEbsVolume {
	var returns *[]*OpsworksRailsAppLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) ManageBundler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) ManageBundlerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBundlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) PassengerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passengerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) PassengerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passengerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) RubygemsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubygemsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) RubygemsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubygemsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) RubyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) RubyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html aws_opsworks_rails_app_layer} Resource.
func NewOpsworksRailsAppLayer(scope constructs.Construct, id *string, config *OpsworksRailsAppLayerConfig) OpsworksRailsAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksRailsAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksRailsAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html aws_opsworks_rails_app_layer} Resource.
func NewOpsworksRailsAppLayer_Override(o OpsworksRailsAppLayer, scope constructs.Construct, id *string, config *OpsworksRailsAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksRailsAppLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetAppServer(val *string) {
	_jsii_.Set(
		j,
		"appServer",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetBundlerVersion(val *string) {
	_jsii_.Set(
		j,
		"bundlerVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetEbsVolume(val *[]*OpsworksRailsAppLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetManageBundler(val interface{}) {
	_jsii_.Set(
		j,
		"manageBundler",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetPassengerVersion(val *string) {
	_jsii_.Set(
		j,
		"passengerVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetRubygemsVersion(val *string) {
	_jsii_.Set(
		j,
		"rubygemsVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetRubyVersion(val *string) {
	_jsii_.Set(
		j,
		"rubyVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksRailsAppLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksRailsAppLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksRailsAppLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksRailsAppLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetAppServer() {
	_jsii_.InvokeVoid(
		o,
		"resetAppServer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetBundlerVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetBundlerVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetManageBundler() {
	_jsii_.InvokeVoid(
		o,
		"resetManageBundler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetPassengerVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetPassengerVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetRubygemsVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetRubygemsVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetRubyVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetRubyVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksRailsAppLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksRailsAppLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#stack_id OpsworksRailsAppLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#app_server OpsworksRailsAppLayer#app_server}.
	AppServer *string `json:"appServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#auto_assign_elastic_ips OpsworksRailsAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#auto_assign_public_ips OpsworksRailsAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#auto_healing OpsworksRailsAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#bundler_version OpsworksRailsAppLayer#bundler_version}.
	BundlerVersion *string `json:"bundlerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_configure_recipes OpsworksRailsAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_deploy_recipes OpsworksRailsAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_instance_profile_arn OpsworksRailsAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_json OpsworksRailsAppLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_security_group_ids OpsworksRailsAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_setup_recipes OpsworksRailsAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_shutdown_recipes OpsworksRailsAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#custom_undeploy_recipes OpsworksRailsAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#drain_elb_on_shutdown OpsworksRailsAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#ebs_volume OpsworksRailsAppLayer#ebs_volume}
	EbsVolume *[]*OpsworksRailsAppLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#elastic_load_balancer OpsworksRailsAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#install_updates_on_boot OpsworksRailsAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#instance_shutdown_timeout OpsworksRailsAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#manage_bundler OpsworksRailsAppLayer#manage_bundler}.
	ManageBundler interface{} `json:"manageBundler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#name OpsworksRailsAppLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#passenger_version OpsworksRailsAppLayer#passenger_version}.
	PassengerVersion *string `json:"passengerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#rubygems_version OpsworksRailsAppLayer#rubygems_version}.
	RubygemsVersion *string `json:"rubygemsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#ruby_version OpsworksRailsAppLayer#ruby_version}.
	RubyVersion *string `json:"rubyVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#system_packages OpsworksRailsAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#tags OpsworksRailsAppLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#tags_all OpsworksRailsAppLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#use_ebs_optimized_instances OpsworksRailsAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksRailsAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#mount_point OpsworksRailsAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#number_of_disks OpsworksRailsAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#size OpsworksRailsAppLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#encrypted OpsworksRailsAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#iops OpsworksRailsAppLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#raid_level OpsworksRailsAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer.html#type OpsworksRailsAppLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html aws_opsworks_rds_db_instance}.
type OpsworksRdsDbInstance interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DbPassword() *string
	SetDbPassword(val *string)
	DbPasswordInput() *string
	DbUser() *string
	SetDbUser(val *string)
	DbUserInput() *string
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
	RdsDbInstanceArn() *string
	SetRdsDbInstanceArn(val *string)
	RdsDbInstanceArnInput() *string
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
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

// The jsii proxy struct for OpsworksRdsDbInstance
type jsiiProxy_OpsworksRdsDbInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksRdsDbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) DbPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) DbPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) DbUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) DbUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) RdsDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) RdsDbInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRdsDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html aws_opsworks_rds_db_instance} Resource.
func NewOpsworksRdsDbInstance(scope constructs.Construct, id *string, config *OpsworksRdsDbInstanceConfig) OpsworksRdsDbInstance {
	_init_.Initialize()

	j := jsiiProxy_OpsworksRdsDbInstance{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksRdsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html aws_opsworks_rds_db_instance} Resource.
func NewOpsworksRdsDbInstance_Override(o OpsworksRdsDbInstance, scope constructs.Construct, id *string, config *OpsworksRdsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksRdsDbInstance",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetDbPassword(val *string) {
	_jsii_.Set(
		j,
		"dbPassword",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetDbUser(val *string) {
	_jsii_.Set(
		j,
		"dbUser",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetRdsDbInstanceArn(val *string) {
	_jsii_.Set(
		j,
		"rdsDbInstanceArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksRdsDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksRdsDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksRdsDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksRdsDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRdsDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksRdsDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksRdsDbInstanceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html#db_password OpsworksRdsDbInstance#db_password}.
	DbPassword *string `json:"dbPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html#db_user OpsworksRdsDbInstance#db_user}.
	DbUser *string `json:"dbUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html#rds_db_instance_arn OpsworksRdsDbInstance#rds_db_instance_arn}.
	RdsDbInstanceArn *string `json:"rdsDbInstanceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance.html#stack_id OpsworksRdsDbInstance#stack_id}.
	StackId *string `json:"stackId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html aws_opsworks_stack}.
type OpsworksStack interface {
	cdktf.TerraformResource
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	Arn() *string
	BerkshelfVersion() *string
	SetBerkshelfVersion(val *string)
	BerkshelfVersionInput() *string
	CdktfStack() cdktf.TerraformStack
	Color() *string
	SetColor(val *string)
	ColorInput() *string
	ConfigurationManagerName() *string
	SetConfigurationManagerName(val *string)
	ConfigurationManagerNameInput() *string
	ConfigurationManagerVersion() *string
	SetConfigurationManagerVersion(val *string)
	ConfigurationManagerVersionInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomCookbooksSource() *[]*OpsworksStackCustomCookbooksSource
	SetCustomCookbooksSource(val *[]*OpsworksStackCustomCookbooksSource)
	CustomCookbooksSourceInput() *[]*OpsworksStackCustomCookbooksSource
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	DefaultAvailabilityZone() *string
	SetDefaultAvailabilityZone(val *string)
	DefaultAvailabilityZoneInput() *string
	DefaultInstanceProfileArn() *string
	SetDefaultInstanceProfileArn(val *string)
	DefaultInstanceProfileArnInput() *string
	DefaultOs() *string
	SetDefaultOs(val *string)
	DefaultOsInput() *string
	DefaultRootDeviceType() *string
	SetDefaultRootDeviceType(val *string)
	DefaultRootDeviceTypeInput() *string
	DefaultSshKeyName() *string
	SetDefaultSshKeyName(val *string)
	DefaultSshKeyNameInput() *string
	DefaultSubnetId() *string
	SetDefaultSubnetId(val *string)
	DefaultSubnetIdInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	HostnameTheme() *string
	SetHostnameTheme(val *string)
	HostnameThemeInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageBerkshelf() interface{}
	SetManageBerkshelf(val interface{})
	ManageBerkshelfInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	StackEndpoint() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseCustomCookbooks() interface{}
	SetUseCustomCookbooks(val interface{})
	UseCustomCookbooksInput() interface{}
	UseOpsworksSecurityGroups() interface{}
	SetUseOpsworksSecurityGroups(val interface{})
	UseOpsworksSecurityGroupsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAgentVersion()
	ResetBerkshelfVersion()
	ResetColor()
	ResetConfigurationManagerName()
	ResetConfigurationManagerVersion()
	ResetCustomCookbooksSource()
	ResetCustomJson()
	ResetDefaultAvailabilityZone()
	ResetDefaultOs()
	ResetDefaultRootDeviceType()
	ResetDefaultSshKeyName()
	ResetDefaultSubnetId()
	ResetHostnameTheme()
	ResetManageBerkshelf()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetUseCustomCookbooks()
	ResetUseOpsworksSecurityGroups()
	ResetVpcId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksStack
type jsiiProxy_OpsworksStack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksStack) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) BerkshelfVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"berkshelfVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) BerkshelfVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"berkshelfVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSource() *[]*OpsworksStackCustomCookbooksSource {
	var returns *[]*OpsworksStackCustomCookbooksSource
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSourceInput() *[]*OpsworksStackCustomCookbooksSource {
	var returns *[]*OpsworksStackCustomCookbooksSource
	_jsii_.Get(
		j,
		"customCookbooksSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultOs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultOsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultRootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultRootDeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSshKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) HostnameTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) HostnameThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameThemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ManageBerkshelf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBerkshelf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ManageBerkshelfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBerkshelfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) StackEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseCustomCookbooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseCustomCookbooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseOpsworksSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseOpsworksSecurityGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html aws_opsworks_stack} Resource.
func NewOpsworksStack(scope constructs.Construct, id *string, config *OpsworksStackConfig) OpsworksStack {
	_init_.Initialize()

	j := jsiiProxy_OpsworksStack{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html aws_opsworks_stack} Resource.
func NewOpsworksStack_Override(o OpsworksStack, scope constructs.Construct, id *string, config *OpsworksStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksStack",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksStack) SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetBerkshelfVersion(val *string) {
	_jsii_.Set(
		j,
		"berkshelfVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetColor(val *string) {
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetConfigurationManagerName(val *string) {
	_jsii_.Set(
		j,
		"configurationManagerName",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetConfigurationManagerVersion(val *string) {
	_jsii_.Set(
		j,
		"configurationManagerVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetCustomCookbooksSource(val *[]*OpsworksStackCustomCookbooksSource) {
	_jsii_.Set(
		j,
		"customCookbooksSource",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"defaultAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"defaultInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultOs(val *string) {
	_jsii_.Set(
		j,
		"defaultOs",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"defaultRootDeviceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"defaultSshKeyName",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDefaultSubnetId(val *string) {
	_jsii_.Set(
		j,
		"defaultSubnetId",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetHostnameTheme(val *string) {
	_jsii_.Set(
		j,
		"hostnameTheme",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetManageBerkshelf(val interface{}) {
	_jsii_.Set(
		j,
		"manageBerkshelf",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetUseCustomCookbooks(val interface{}) {
	_jsii_.Set(
		j,
		"useCustomCookbooks",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetUseOpsworksSecurityGroups(val interface{}) {
	_jsii_.Set(
		j,
		"useOpsworksSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksStack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksStack",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksStack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksStack) ResetAgentVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAgentVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetBerkshelfVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetBerkshelfVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetColor() {
	_jsii_.InvokeVoid(
		o,
		"resetColor",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetConfigurationManagerName() {
	_jsii_.InvokeVoid(
		o,
		"resetConfigurationManagerName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetConfigurationManagerVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetConfigurationManagerVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetCustomCookbooksSource() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomCookbooksSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultAvailabilityZone() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultAvailabilityZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultOs() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultOs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultRootDeviceType() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultRootDeviceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultSshKeyName() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultSshKeyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultSubnetId() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultSubnetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetHostnameTheme() {
	_jsii_.InvokeVoid(
		o,
		"resetHostnameTheme",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetManageBerkshelf() {
	_jsii_.InvokeVoid(
		o,
		"resetManageBerkshelf",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksStack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetUseCustomCookbooks() {
	_jsii_.InvokeVoid(
		o,
		"resetUseCustomCookbooks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetUseOpsworksSecurityGroups() {
	_jsii_.InvokeVoid(
		o,
		"resetUseOpsworksSecurityGroups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetVpcId() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksStackConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_instance_profile_arn OpsworksStack#default_instance_profile_arn}.
	DefaultInstanceProfileArn *string `json:"defaultInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#name OpsworksStack#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#region OpsworksStack#region}.
	Region *string `json:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#service_role_arn OpsworksStack#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#agent_version OpsworksStack#agent_version}.
	AgentVersion *string `json:"agentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#berkshelf_version OpsworksStack#berkshelf_version}.
	BerkshelfVersion *string `json:"berkshelfVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#color OpsworksStack#color}.
	Color *string `json:"color"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#configuration_manager_name OpsworksStack#configuration_manager_name}.
	ConfigurationManagerName *string `json:"configurationManagerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#configuration_manager_version OpsworksStack#configuration_manager_version}.
	ConfigurationManagerVersion *string `json:"configurationManagerVersion"`
	// custom_cookbooks_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#custom_cookbooks_source OpsworksStack#custom_cookbooks_source}
	CustomCookbooksSource *[]*OpsworksStackCustomCookbooksSource `json:"customCookbooksSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#custom_json OpsworksStack#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_availability_zone OpsworksStack#default_availability_zone}.
	DefaultAvailabilityZone *string `json:"defaultAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_os OpsworksStack#default_os}.
	DefaultOs *string `json:"defaultOs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_root_device_type OpsworksStack#default_root_device_type}.
	DefaultRootDeviceType *string `json:"defaultRootDeviceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_ssh_key_name OpsworksStack#default_ssh_key_name}.
	DefaultSshKeyName *string `json:"defaultSshKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#default_subnet_id OpsworksStack#default_subnet_id}.
	DefaultSubnetId *string `json:"defaultSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#hostname_theme OpsworksStack#hostname_theme}.
	HostnameTheme *string `json:"hostnameTheme"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#manage_berkshelf OpsworksStack#manage_berkshelf}.
	ManageBerkshelf interface{} `json:"manageBerkshelf"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#tags OpsworksStack#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#tags_all OpsworksStack#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#use_custom_cookbooks OpsworksStack#use_custom_cookbooks}.
	UseCustomCookbooks interface{} `json:"useCustomCookbooks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#use_opsworks_security_groups OpsworksStack#use_opsworks_security_groups}.
	UseOpsworksSecurityGroups interface{} `json:"useOpsworksSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#vpc_id OpsworksStack#vpc_id}.
	VpcId *string `json:"vpcId"`
}

type OpsworksStackCustomCookbooksSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#type OpsworksStack#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#url OpsworksStack#url}.
	Url *string `json:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#password OpsworksStack#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#revision OpsworksStack#revision}.
	Revision *string `json:"revision"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#ssh_key OpsworksStack#ssh_key}.
	SshKey *string `json:"sshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack.html#username OpsworksStack#username}.
	Username *string `json:"username"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html aws_opsworks_static_web_layer}.
type OpsworksStaticWebLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() *[]*OpsworksStaticWebLayerEbsVolume
	SetEbsVolume(val *[]*OpsworksStaticWebLayerEbsVolume)
	EbsVolumeInput() *[]*OpsworksStaticWebLayerEbsVolume
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetName()
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksStaticWebLayer
type jsiiProxy_OpsworksStaticWebLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) EbsVolume() *[]*OpsworksStaticWebLayerEbsVolume {
	var returns *[]*OpsworksStaticWebLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) EbsVolumeInput() *[]*OpsworksStaticWebLayerEbsVolume {
	var returns *[]*OpsworksStaticWebLayerEbsVolume
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html aws_opsworks_static_web_layer} Resource.
func NewOpsworksStaticWebLayer(scope constructs.Construct, id *string, config *OpsworksStaticWebLayerConfig) OpsworksStaticWebLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksStaticWebLayer{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksStaticWebLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html aws_opsworks_static_web_layer} Resource.
func NewOpsworksStaticWebLayer_Override(o OpsworksStaticWebLayer, scope constructs.Construct, id *string, config *OpsworksStaticWebLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksStaticWebLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetAutoAssignElasticIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetAutoAssignPublicIps(val interface{}) {
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetAutoHealing(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomConfigureRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomDeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomJson(val *string) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomSetupRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomShutdownRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCustomUndeployRecipes(val *[]*string) {
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetDrainElbOnShutdown(val interface{}) {
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetEbsVolume(val *[]*OpsworksStaticWebLayerEbsVolume) {
	_jsii_.Set(
		j,
		"ebsVolume",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetElasticLoadBalancer(val *string) {
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetInstallUpdatesOnBoot(val interface{}) {
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetInstanceShutdownTimeout(val *float64) {
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetStackId(val *string) {
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetSystemPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetUseEbsOptimizedInstances(val interface{}) {
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksStaticWebLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksStaticWebLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksStaticWebLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksStaticWebLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksStaticWebLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksStaticWebLayerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#stack_id OpsworksStaticWebLayer#stack_id}.
	StackId *string `json:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#auto_assign_elastic_ips OpsworksStaticWebLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#auto_assign_public_ips OpsworksStaticWebLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#auto_healing OpsworksStaticWebLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_configure_recipes OpsworksStaticWebLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_deploy_recipes OpsworksStaticWebLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_instance_profile_arn OpsworksStaticWebLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_json OpsworksStaticWebLayer#custom_json}.
	CustomJson *string `json:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_security_group_ids OpsworksStaticWebLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_setup_recipes OpsworksStaticWebLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_shutdown_recipes OpsworksStaticWebLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#custom_undeploy_recipes OpsworksStaticWebLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#drain_elb_on_shutdown OpsworksStaticWebLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#ebs_volume OpsworksStaticWebLayer#ebs_volume}
	EbsVolume *[]*OpsworksStaticWebLayerEbsVolume `json:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#elastic_load_balancer OpsworksStaticWebLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#install_updates_on_boot OpsworksStaticWebLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#instance_shutdown_timeout OpsworksStaticWebLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#name OpsworksStaticWebLayer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#system_packages OpsworksStaticWebLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#tags OpsworksStaticWebLayer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#tags_all OpsworksStaticWebLayer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#use_ebs_optimized_instances OpsworksStaticWebLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances"`
}

type OpsworksStaticWebLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#mount_point OpsworksStaticWebLayer#mount_point}.
	MountPoint *string `json:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#number_of_disks OpsworksStaticWebLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#size OpsworksStaticWebLayer#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#encrypted OpsworksStaticWebLayer#encrypted}.
	Encrypted interface{} `json:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#iops OpsworksStaticWebLayer#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#raid_level OpsworksStaticWebLayer#raid_level}.
	RaidLevel *string `json:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer.html#type OpsworksStaticWebLayer#type}.
	Type *string `json:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html aws_opsworks_user_profile}.
type OpsworksUserProfile interface {
	cdktf.TerraformResource
	AllowSelfManagement() interface{}
	SetAllowSelfManagement(val interface{})
	AllowSelfManagementInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshPublicKeyInput() *string
	SshUsername() *string
	SetSshUsername(val *string)
	SshUsernameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserArn() *string
	SetUserArn(val *string)
	UserArnInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAllowSelfManagement()
	ResetOverrideLogicalId()
	ResetSshPublicKey()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksUserProfile
type jsiiProxy_OpsworksUserProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksUserProfile) AllowSelfManagement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelfManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) AllowSelfManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelfManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) SshPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) SshUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) SshUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) UserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksUserProfile) UserArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArnInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html aws_opsworks_user_profile} Resource.
func NewOpsworksUserProfile(scope constructs.Construct, id *string, config *OpsworksUserProfileConfig) OpsworksUserProfile {
	_init_.Initialize()

	j := jsiiProxy_OpsworksUserProfile{}

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksUserProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html aws_opsworks_user_profile} Resource.
func NewOpsworksUserProfile_Override(o OpsworksUserProfile, scope constructs.Construct, id *string, config *OpsworksUserProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.OpsWorks.OpsworksUserProfile",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetAllowSelfManagement(val interface{}) {
	_jsii_.Set(
		j,
		"allowSelfManagement",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetSshPublicKey(val *string) {
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetSshUsername(val *string) {
	_jsii_.Set(
		j,
		"sshUsername",
		val,
	)
}

func (j *jsiiProxy_OpsworksUserProfile) SetUserArn(val *string) {
	_jsii_.Set(
		j,
		"userArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OpsworksUserProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.OpsWorks.OpsworksUserProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksUserProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.OpsWorks.OpsworksUserProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksUserProfile) ResetAllowSelfManagement() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowSelfManagement",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksUserProfile) ResetSshPublicKey() {
	_jsii_.InvokeVoid(
		o,
		"resetSshPublicKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksUserProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OpsworksUserProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OpsworksUserProfileConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html#ssh_username OpsworksUserProfile#ssh_username}.
	SshUsername *string `json:"sshUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html#user_arn OpsworksUserProfile#user_arn}.
	UserArn *string `json:"userArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html#allow_self_management OpsworksUserProfile#allow_self_management}.
	AllowSelfManagement interface{} `json:"allowSelfManagement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile.html#ssh_public_key OpsworksUserProfile#ssh_public_key}.
	SshPublicKey *string `json:"sshPublicKey"`
}
