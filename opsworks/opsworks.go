package opsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/opsworks/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application aws_opsworks_application}.
type OpsworksApplication interface {
	cdktf.TerraformResource
	AppSource() interface{}
	SetAppSource(val interface{})
	AppSourceInput() interface{}
	AutoBundleOnDeploy() *string
	SetAutoBundleOnDeploy(val *string)
	AutoBundleOnDeployInput() *string
	AwsFlowRubySettings() *string
	SetAwsFlowRubySettings(val *string)
	AwsFlowRubySettingsInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	Environment() interface{}
	SetEnvironment(val interface{})
	EnvironmentInput() interface{}
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
	SslConfiguration() interface{}
	SetSslConfiguration(val interface{})
	SslConfigurationInput() interface{}
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

func (j *jsiiProxy_OpsworksApplication) AppSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AppSourceInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksApplication) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksApplication) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnvironmentInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksApplication) SslConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) SslConfigurationInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application aws_opsworks_application} Resource.
func NewOpsworksApplication(scope constructs.Construct, id *string, config *OpsworksApplicationConfig) OpsworksApplication {
	_init_.Initialize()

	j := jsiiProxy_OpsworksApplication{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application aws_opsworks_application} Resource.
func NewOpsworksApplication_Override(o OpsworksApplication, scope constructs.Construct, id *string, config *OpsworksApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksApplication",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksApplication) SetAppSource(val interface{}) {
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

func (j *jsiiProxy_OpsworksApplication) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksApplication) SetEnvironment(val interface{}) {
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

func (j *jsiiProxy_OpsworksApplication) SetSslConfiguration(val interface{}) {
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
		"hashicorp_aws.opsworks.OpsworksApplication",
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
		"hashicorp_aws.opsworks.OpsworksApplication",
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
func (o *jsiiProxy_OpsworksApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#type OpsworksApplication#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#password OpsworksApplication#password}.
	Password *string `json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#revision OpsworksApplication#revision}.
	Revision *string `json:"revision" yaml:"revision"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#ssh_key OpsworksApplication#ssh_key}.
	SshKey *string `json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#url OpsworksApplication#url}.
	Url *string `json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#username OpsworksApplication#username}.
	Username *string `json:"username" yaml:"username"`
}

// AWS OpsWorks.
type OpsworksApplicationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#name OpsworksApplication#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#stack_id OpsworksApplication#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#type OpsworksApplication#type}.
	Type *string `json:"type" yaml:"type"`
	// app_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#app_source OpsworksApplication#app_source}
	AppSource interface{} `json:"appSource" yaml:"appSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#auto_bundle_on_deploy OpsworksApplication#auto_bundle_on_deploy}.
	AutoBundleOnDeploy *string `json:"autoBundleOnDeploy" yaml:"autoBundleOnDeploy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#aws_flow_ruby_settings OpsworksApplication#aws_flow_ruby_settings}.
	AwsFlowRubySettings *string `json:"awsFlowRubySettings" yaml:"awsFlowRubySettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#data_source_arn OpsworksApplication#data_source_arn}.
	DataSourceArn *string `json:"dataSourceArn" yaml:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#data_source_database_name OpsworksApplication#data_source_database_name}.
	DataSourceDatabaseName *string `json:"dataSourceDatabaseName" yaml:"dataSourceDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#data_source_type OpsworksApplication#data_source_type}.
	DataSourceType *string `json:"dataSourceType" yaml:"dataSourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#description OpsworksApplication#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#document_root OpsworksApplication#document_root}.
	DocumentRoot *string `json:"documentRoot" yaml:"documentRoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#domains OpsworksApplication#domains}.
	Domains *[]*string `json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#enable_ssl OpsworksApplication#enable_ssl}.
	EnableSsl interface{} `json:"enableSsl" yaml:"enableSsl"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#environment OpsworksApplication#environment}
	Environment interface{} `json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#rails_env OpsworksApplication#rails_env}.
	RailsEnv *string `json:"railsEnv" yaml:"railsEnv"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#short_name OpsworksApplication#short_name}.
	ShortName *string `json:"shortName" yaml:"shortName"`
	// ssl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#ssl_configuration OpsworksApplication#ssl_configuration}
	SslConfiguration interface{} `json:"sslConfiguration" yaml:"sslConfiguration"`
}

type OpsworksApplicationEnvironment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#key OpsworksApplication#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#value OpsworksApplication#value}.
	Value *string `json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#secure OpsworksApplication#secure}.
	Secure interface{} `json:"secure" yaml:"secure"`
}

type OpsworksApplicationSslConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#certificate OpsworksApplication#certificate}.
	Certificate *string `json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#private_key OpsworksApplication#private_key}.
	PrivateKey *string `json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_application#chain OpsworksApplication#chain}.
	Chain *string `json:"chain" yaml:"chain"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer aws_opsworks_custom_layer}.
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
	CloudwatchConfiguration() OpsworksCustomLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksCustomLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksCustomLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksCustomLayer) CloudwatchConfiguration() OpsworksCustomLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksCustomLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) CloudwatchConfigurationInput() *OpsworksCustomLayerCloudwatchConfiguration {
	var returns *OpsworksCustomLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksCustomLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksCustomLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksCustomLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer aws_opsworks_custom_layer} Resource.
func NewOpsworksCustomLayer(scope constructs.Construct, id *string, config *OpsworksCustomLayerConfig) OpsworksCustomLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksCustomLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksCustomLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer aws_opsworks_custom_layer} Resource.
func NewOpsworksCustomLayer_Override(o OpsworksCustomLayer, scope constructs.Construct, id *string, config *OpsworksCustomLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksCustomLayer",
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

func (j *jsiiProxy_OpsworksCustomLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksCustomLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksCustomLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksCustomLayer",
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
		"hashicorp_aws.opsworks.OpsworksCustomLayer",
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
func (o *jsiiProxy_OpsworksCustomLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksCustomLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksCustomLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksCustomLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksCustomLayer) PutCloudwatchConfiguration(value *OpsworksCustomLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksCustomLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksCustomLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#enabled OpsworksCustomLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#log_streams OpsworksCustomLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksCustomLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#file OpsworksCustomLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#log_group_name OpsworksCustomLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#batch_count OpsworksCustomLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#batch_size OpsworksCustomLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#buffer_duration OpsworksCustomLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#datetime_format OpsworksCustomLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#encoding OpsworksCustomLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#file_fingerprint_lines OpsworksCustomLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#initial_position OpsworksCustomLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#multiline_start_pattern OpsworksCustomLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#time_zone OpsworksCustomLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksCustomLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksCustomLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksCustomLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksCustomLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksCustomLayerCloudwatchConfiguration {
	var returns *OpsworksCustomLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksCustomLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksCustomLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksCustomLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksCustomLayerCloudwatchConfigurationOutputReference_Override(o OpsworksCustomLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksCustomLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksCustomLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksCustomLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#name OpsworksCustomLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#short_name OpsworksCustomLayer#short_name}.
	ShortName *string `json:"shortName" yaml:"shortName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#stack_id OpsworksCustomLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#auto_assign_elastic_ips OpsworksCustomLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#auto_assign_public_ips OpsworksCustomLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#auto_healing OpsworksCustomLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#cloudwatch_configuration OpsworksCustomLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksCustomLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_configure_recipes OpsworksCustomLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_deploy_recipes OpsworksCustomLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_instance_profile_arn OpsworksCustomLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_json OpsworksCustomLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_security_group_ids OpsworksCustomLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_setup_recipes OpsworksCustomLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_shutdown_recipes OpsworksCustomLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#custom_undeploy_recipes OpsworksCustomLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#drain_elb_on_shutdown OpsworksCustomLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#ebs_volume OpsworksCustomLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#elastic_load_balancer OpsworksCustomLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#install_updates_on_boot OpsworksCustomLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#instance_shutdown_timeout OpsworksCustomLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#system_packages OpsworksCustomLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#tags OpsworksCustomLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#tags_all OpsworksCustomLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#use_ebs_optimized_instances OpsworksCustomLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksCustomLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#mount_point OpsworksCustomLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#number_of_disks OpsworksCustomLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#size OpsworksCustomLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#encrypted OpsworksCustomLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#iops OpsworksCustomLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#raid_level OpsworksCustomLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_custom_layer#type OpsworksCustomLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer aws_opsworks_ganglia_layer}.
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
	CloudwatchConfiguration() OpsworksGangliaLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksGangliaLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	PutCloudwatchConfiguration(value *OpsworksGangliaLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksGangliaLayer) CloudwatchConfiguration() OpsworksGangliaLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksGangliaLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) CloudwatchConfigurationInput() *OpsworksGangliaLayerCloudwatchConfiguration {
	var returns *OpsworksGangliaLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksGangliaLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksGangliaLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksGangliaLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer aws_opsworks_ganglia_layer} Resource.
func NewOpsworksGangliaLayer(scope constructs.Construct, id *string, config *OpsworksGangliaLayerConfig) OpsworksGangliaLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksGangliaLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksGangliaLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer aws_opsworks_ganglia_layer} Resource.
func NewOpsworksGangliaLayer_Override(o OpsworksGangliaLayer, scope constructs.Construct, id *string, config *OpsworksGangliaLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksGangliaLayer",
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

func (j *jsiiProxy_OpsworksGangliaLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksGangliaLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksGangliaLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksGangliaLayer",
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
		"hashicorp_aws.opsworks.OpsworksGangliaLayer",
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
func (o *jsiiProxy_OpsworksGangliaLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksGangliaLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksGangliaLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksGangliaLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksGangliaLayer) PutCloudwatchConfiguration(value *OpsworksGangliaLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksGangliaLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksGangliaLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#enabled OpsworksGangliaLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#log_streams OpsworksGangliaLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksGangliaLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#file OpsworksGangliaLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#log_group_name OpsworksGangliaLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#batch_count OpsworksGangliaLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#batch_size OpsworksGangliaLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#buffer_duration OpsworksGangliaLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#datetime_format OpsworksGangliaLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#encoding OpsworksGangliaLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#file_fingerprint_lines OpsworksGangliaLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#initial_position OpsworksGangliaLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#multiline_start_pattern OpsworksGangliaLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#time_zone OpsworksGangliaLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksGangliaLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksGangliaLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksGangliaLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksGangliaLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksGangliaLayerCloudwatchConfiguration {
	var returns *OpsworksGangliaLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksGangliaLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksGangliaLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksGangliaLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksGangliaLayerCloudwatchConfigurationOutputReference_Override(o OpsworksGangliaLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksGangliaLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksGangliaLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksGangliaLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksGangliaLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#password OpsworksGangliaLayer#password}.
	Password *string `json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#stack_id OpsworksGangliaLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#auto_assign_elastic_ips OpsworksGangliaLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#auto_assign_public_ips OpsworksGangliaLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#auto_healing OpsworksGangliaLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#cloudwatch_configuration OpsworksGangliaLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksGangliaLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_configure_recipes OpsworksGangliaLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_deploy_recipes OpsworksGangliaLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_instance_profile_arn OpsworksGangliaLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_json OpsworksGangliaLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_security_group_ids OpsworksGangliaLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_setup_recipes OpsworksGangliaLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_shutdown_recipes OpsworksGangliaLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#custom_undeploy_recipes OpsworksGangliaLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#drain_elb_on_shutdown OpsworksGangliaLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#ebs_volume OpsworksGangliaLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#elastic_load_balancer OpsworksGangliaLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#install_updates_on_boot OpsworksGangliaLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#instance_shutdown_timeout OpsworksGangliaLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#name OpsworksGangliaLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#system_packages OpsworksGangliaLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#tags OpsworksGangliaLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#tags_all OpsworksGangliaLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#url OpsworksGangliaLayer#url}.
	Url *string `json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#use_ebs_optimized_instances OpsworksGangliaLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#username OpsworksGangliaLayer#username}.
	Username *string `json:"username" yaml:"username"`
}

type OpsworksGangliaLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#mount_point OpsworksGangliaLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#number_of_disks OpsworksGangliaLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#size OpsworksGangliaLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#encrypted OpsworksGangliaLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#iops OpsworksGangliaLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#raid_level OpsworksGangliaLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ganglia_layer#type OpsworksGangliaLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer aws_opsworks_haproxy_layer}.
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
	CloudwatchConfiguration() OpsworksHaproxyLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksHaproxyLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksHaproxyLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksHaproxyLayer) CloudwatchConfiguration() OpsworksHaproxyLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksHaproxyLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) CloudwatchConfigurationInput() *OpsworksHaproxyLayerCloudwatchConfiguration {
	var returns *OpsworksHaproxyLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksHaproxyLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksHaproxyLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksHaproxyLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer aws_opsworks_haproxy_layer} Resource.
func NewOpsworksHaproxyLayer(scope constructs.Construct, id *string, config *OpsworksHaproxyLayerConfig) OpsworksHaproxyLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksHaproxyLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksHaproxyLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer aws_opsworks_haproxy_layer} Resource.
func NewOpsworksHaproxyLayer_Override(o OpsworksHaproxyLayer, scope constructs.Construct, id *string, config *OpsworksHaproxyLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksHaproxyLayer",
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

func (j *jsiiProxy_OpsworksHaproxyLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksHaproxyLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksHaproxyLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksHaproxyLayer",
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
		"hashicorp_aws.opsworks.OpsworksHaproxyLayer",
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
func (o *jsiiProxy_OpsworksHaproxyLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksHaproxyLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksHaproxyLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksHaproxyLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksHaproxyLayer) PutCloudwatchConfiguration(value *OpsworksHaproxyLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksHaproxyLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksHaproxyLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#enabled OpsworksHaproxyLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#log_streams OpsworksHaproxyLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksHaproxyLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#file OpsworksHaproxyLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#log_group_name OpsworksHaproxyLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#batch_count OpsworksHaproxyLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#batch_size OpsworksHaproxyLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#buffer_duration OpsworksHaproxyLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#datetime_format OpsworksHaproxyLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#encoding OpsworksHaproxyLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#file_fingerprint_lines OpsworksHaproxyLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#initial_position OpsworksHaproxyLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#multiline_start_pattern OpsworksHaproxyLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#time_zone OpsworksHaproxyLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksHaproxyLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksHaproxyLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksHaproxyLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksHaproxyLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksHaproxyLayerCloudwatchConfiguration {
	var returns *OpsworksHaproxyLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksHaproxyLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksHaproxyLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksHaproxyLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksHaproxyLayerCloudwatchConfigurationOutputReference_Override(o OpsworksHaproxyLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksHaproxyLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksHaproxyLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksHaproxyLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksHaproxyLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#stack_id OpsworksHaproxyLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#stats_password OpsworksHaproxyLayer#stats_password}.
	StatsPassword *string `json:"statsPassword" yaml:"statsPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#auto_assign_elastic_ips OpsworksHaproxyLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#auto_assign_public_ips OpsworksHaproxyLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#auto_healing OpsworksHaproxyLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#cloudwatch_configuration OpsworksHaproxyLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksHaproxyLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_configure_recipes OpsworksHaproxyLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_deploy_recipes OpsworksHaproxyLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_instance_profile_arn OpsworksHaproxyLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_json OpsworksHaproxyLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_security_group_ids OpsworksHaproxyLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_setup_recipes OpsworksHaproxyLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_shutdown_recipes OpsworksHaproxyLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#custom_undeploy_recipes OpsworksHaproxyLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#drain_elb_on_shutdown OpsworksHaproxyLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#ebs_volume OpsworksHaproxyLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#elastic_load_balancer OpsworksHaproxyLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#healthcheck_method OpsworksHaproxyLayer#healthcheck_method}.
	HealthcheckMethod *string `json:"healthcheckMethod" yaml:"healthcheckMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#healthcheck_url OpsworksHaproxyLayer#healthcheck_url}.
	HealthcheckUrl *string `json:"healthcheckUrl" yaml:"healthcheckUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#install_updates_on_boot OpsworksHaproxyLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#instance_shutdown_timeout OpsworksHaproxyLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#name OpsworksHaproxyLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#stats_enabled OpsworksHaproxyLayer#stats_enabled}.
	StatsEnabled interface{} `json:"statsEnabled" yaml:"statsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#stats_url OpsworksHaproxyLayer#stats_url}.
	StatsUrl *string `json:"statsUrl" yaml:"statsUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#stats_user OpsworksHaproxyLayer#stats_user}.
	StatsUser *string `json:"statsUser" yaml:"statsUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#system_packages OpsworksHaproxyLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#tags OpsworksHaproxyLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#tags_all OpsworksHaproxyLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#use_ebs_optimized_instances OpsworksHaproxyLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksHaproxyLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#mount_point OpsworksHaproxyLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#number_of_disks OpsworksHaproxyLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#size OpsworksHaproxyLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#encrypted OpsworksHaproxyLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#iops OpsworksHaproxyLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#raid_level OpsworksHaproxyLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_haproxy_layer#type OpsworksHaproxyLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance aws_opsworks_instance}.
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
	Count() *float64
	SetCount(val *float64)
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
	EbsBlockDevice() interface{}
	SetEbsBlockDevice(val interface{})
	EbsBlockDeviceInput() interface{}
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
	EphemeralBlockDevice() interface{}
	SetEphemeralBlockDevice(val interface{})
	EphemeralBlockDeviceInput() interface{}
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
	RootBlockDevice() interface{}
	SetRootBlockDevice(val interface{})
	RootBlockDeviceInput() interface{}
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

func (j *jsiiProxy_OpsworksInstance) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksInstance) EbsBlockDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsBlockDeviceInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksInstance) RootBlockDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootBlockDeviceInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance aws_opsworks_instance} Resource.
func NewOpsworksInstance(scope constructs.Construct, id *string, config *OpsworksInstanceConfig) OpsworksInstance {
	_init_.Initialize()

	j := jsiiProxy_OpsworksInstance{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance aws_opsworks_instance} Resource.
func NewOpsworksInstance_Override(o OpsworksInstance, scope constructs.Construct, id *string, config *OpsworksInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksInstance",
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

func (j *jsiiProxy_OpsworksInstance) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksInstance) SetEbsBlockDevice(val interface{}) {
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

func (j *jsiiProxy_OpsworksInstance) SetEphemeralBlockDevice(val interface{}) {
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

func (j *jsiiProxy_OpsworksInstance) SetRootBlockDevice(val interface{}) {
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
		"hashicorp_aws.opsworks.OpsworksInstance",
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
		"hashicorp_aws.opsworks.OpsworksInstance",
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
func (o *jsiiProxy_OpsworksInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// AWS OpsWorks.
type OpsworksInstanceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#layer_ids OpsworksInstance#layer_ids}.
	LayerIds *[]*string `json:"layerIds" yaml:"layerIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#stack_id OpsworksInstance#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#agent_version OpsworksInstance#agent_version}.
	AgentVersion *string `json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ami_id OpsworksInstance#ami_id}.
	AmiId *string `json:"amiId" yaml:"amiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#architecture OpsworksInstance#architecture}.
	Architecture *string `json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#auto_scaling_type OpsworksInstance#auto_scaling_type}.
	AutoScalingType *string `json:"autoScalingType" yaml:"autoScalingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#availability_zone OpsworksInstance#availability_zone}.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#created_at OpsworksInstance#created_at}.
	CreatedAt *string `json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#delete_ebs OpsworksInstance#delete_ebs}.
	DeleteEbs interface{} `json:"deleteEbs" yaml:"deleteEbs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#delete_eip OpsworksInstance#delete_eip}.
	DeleteEip interface{} `json:"deleteEip" yaml:"deleteEip"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ebs_block_device OpsworksInstance#ebs_block_device}
	EbsBlockDevice interface{} `json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ebs_optimized OpsworksInstance#ebs_optimized}.
	EbsOptimized interface{} `json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ecs_cluster_arn OpsworksInstance#ecs_cluster_arn}.
	EcsClusterArn *string `json:"ecsClusterArn" yaml:"ecsClusterArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#elastic_ip OpsworksInstance#elastic_ip}.
	ElasticIp *string `json:"elasticIp" yaml:"elasticIp"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ephemeral_block_device OpsworksInstance#ephemeral_block_device}
	EphemeralBlockDevice interface{} `json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#hostname OpsworksInstance#hostname}.
	Hostname *string `json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#infrastructure_class OpsworksInstance#infrastructure_class}.
	InfrastructureClass *string `json:"infrastructureClass" yaml:"infrastructureClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#install_updates_on_boot OpsworksInstance#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#instance_profile_arn OpsworksInstance#instance_profile_arn}.
	InstanceProfileArn *string `json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#instance_type OpsworksInstance#instance_type}.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#last_service_error_id OpsworksInstance#last_service_error_id}.
	LastServiceErrorId *string `json:"lastServiceErrorId" yaml:"lastServiceErrorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#os OpsworksInstance#os}.
	Os *string `json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#platform OpsworksInstance#platform}.
	Platform *string `json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#private_dns OpsworksInstance#private_dns}.
	PrivateDns *string `json:"privateDns" yaml:"privateDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#private_ip OpsworksInstance#private_ip}.
	PrivateIp *string `json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#public_dns OpsworksInstance#public_dns}.
	PublicDns *string `json:"publicDns" yaml:"publicDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#public_ip OpsworksInstance#public_ip}.
	PublicIp *string `json:"publicIp" yaml:"publicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#registered_by OpsworksInstance#registered_by}.
	RegisteredBy *string `json:"registeredBy" yaml:"registeredBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#reported_agent_version OpsworksInstance#reported_agent_version}.
	ReportedAgentVersion *string `json:"reportedAgentVersion" yaml:"reportedAgentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#reported_os_family OpsworksInstance#reported_os_family}.
	ReportedOsFamily *string `json:"reportedOsFamily" yaml:"reportedOsFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#reported_os_name OpsworksInstance#reported_os_name}.
	ReportedOsName *string `json:"reportedOsName" yaml:"reportedOsName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#reported_os_version OpsworksInstance#reported_os_version}.
	ReportedOsVersion *string `json:"reportedOsVersion" yaml:"reportedOsVersion"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#root_block_device OpsworksInstance#root_block_device}
	RootBlockDevice interface{} `json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#root_device_type OpsworksInstance#root_device_type}.
	RootDeviceType *string `json:"rootDeviceType" yaml:"rootDeviceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#root_device_volume_id OpsworksInstance#root_device_volume_id}.
	RootDeviceVolumeId *string `json:"rootDeviceVolumeId" yaml:"rootDeviceVolumeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#security_group_ids OpsworksInstance#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ssh_host_dsa_key_fingerprint OpsworksInstance#ssh_host_dsa_key_fingerprint}.
	SshHostDsaKeyFingerprint *string `json:"sshHostDsaKeyFingerprint" yaml:"sshHostDsaKeyFingerprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ssh_host_rsa_key_fingerprint OpsworksInstance#ssh_host_rsa_key_fingerprint}.
	SshHostRsaKeyFingerprint *string `json:"sshHostRsaKeyFingerprint" yaml:"sshHostRsaKeyFingerprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#ssh_key_name OpsworksInstance#ssh_key_name}.
	SshKeyName *string `json:"sshKeyName" yaml:"sshKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#state OpsworksInstance#state}.
	State *string `json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#status OpsworksInstance#status}.
	Status *string `json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#subnet_id OpsworksInstance#subnet_id}.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#tenancy OpsworksInstance#tenancy}.
	Tenancy *string `json:"tenancy" yaml:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#timeouts OpsworksInstance#timeouts}
	Timeouts *OpsworksInstanceTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#virtualization_type OpsworksInstance#virtualization_type}.
	VirtualizationType *string `json:"virtualizationType" yaml:"virtualizationType"`
}

type OpsworksInstanceEbsBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#device_name OpsworksInstance#device_name}.
	DeviceName *string `json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#delete_on_termination OpsworksInstance#delete_on_termination}.
	DeleteOnTermination interface{} `json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#iops OpsworksInstance#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#snapshot_id OpsworksInstance#snapshot_id}.
	SnapshotId *string `json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#volume_size OpsworksInstance#volume_size}.
	VolumeSize *float64 `json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#volume_type OpsworksInstance#volume_type}.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

type OpsworksInstanceEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#device_name OpsworksInstance#device_name}.
	DeviceName *string `json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#virtual_name OpsworksInstance#virtual_name}.
	VirtualName *string `json:"virtualName" yaml:"virtualName"`
}

type OpsworksInstanceRootBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#delete_on_termination OpsworksInstance#delete_on_termination}.
	DeleteOnTermination interface{} `json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#iops OpsworksInstance#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#volume_size OpsworksInstance#volume_size}.
	VolumeSize *float64 `json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#volume_type OpsworksInstance#volume_type}.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

type OpsworksInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#create OpsworksInstance#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#delete OpsworksInstance#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_instance#update OpsworksInstance#update}.
	Update *string `json:"update" yaml:"update"`
}

type OpsworksInstanceTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *OpsworksInstanceTimeouts
	SetInternalValue(val *OpsworksInstanceTimeouts)
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

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) InternalValue() *OpsworksInstanceTimeouts {
	var returns *OpsworksInstanceTimeouts
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
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

func NewOpsworksInstanceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksInstanceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksInstanceTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksInstanceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksInstanceTimeoutsOutputReference_Override(o OpsworksInstanceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksInstanceTimeoutsOutputReference",
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

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetInternalValue(val *OpsworksInstanceTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksInstanceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer aws_opsworks_java_app_layer}.
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
	CloudwatchConfiguration() OpsworksJavaAppLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksJavaAppLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksJavaAppLayerCloudwatchConfiguration)
	ResetAppServer()
	ResetAppServerVersion()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksJavaAppLayer) CloudwatchConfiguration() OpsworksJavaAppLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksJavaAppLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) CloudwatchConfigurationInput() *OpsworksJavaAppLayerCloudwatchConfiguration {
	var returns *OpsworksJavaAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksJavaAppLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksJavaAppLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksJavaAppLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer aws_opsworks_java_app_layer} Resource.
func NewOpsworksJavaAppLayer(scope constructs.Construct, id *string, config *OpsworksJavaAppLayerConfig) OpsworksJavaAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksJavaAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksJavaAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer aws_opsworks_java_app_layer} Resource.
func NewOpsworksJavaAppLayer_Override(o OpsworksJavaAppLayer, scope constructs.Construct, id *string, config *OpsworksJavaAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksJavaAppLayer",
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

func (j *jsiiProxy_OpsworksJavaAppLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksJavaAppLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksJavaAppLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksJavaAppLayer",
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
		"hashicorp_aws.opsworks.OpsworksJavaAppLayer",
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
func (o *jsiiProxy_OpsworksJavaAppLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksJavaAppLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksJavaAppLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksJavaAppLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksJavaAppLayer) PutCloudwatchConfiguration(value *OpsworksJavaAppLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksJavaAppLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksJavaAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#enabled OpsworksJavaAppLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#log_streams OpsworksJavaAppLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksJavaAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#file OpsworksJavaAppLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#log_group_name OpsworksJavaAppLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#batch_count OpsworksJavaAppLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#batch_size OpsworksJavaAppLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#buffer_duration OpsworksJavaAppLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#datetime_format OpsworksJavaAppLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#encoding OpsworksJavaAppLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#file_fingerprint_lines OpsworksJavaAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#initial_position OpsworksJavaAppLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#multiline_start_pattern OpsworksJavaAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#time_zone OpsworksJavaAppLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksJavaAppLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksJavaAppLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksJavaAppLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksJavaAppLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksJavaAppLayerCloudwatchConfiguration {
	var returns *OpsworksJavaAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksJavaAppLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksJavaAppLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksJavaAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksJavaAppLayerCloudwatchConfigurationOutputReference_Override(o OpsworksJavaAppLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksJavaAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksJavaAppLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksJavaAppLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#stack_id OpsworksJavaAppLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#app_server OpsworksJavaAppLayer#app_server}.
	AppServer *string `json:"appServer" yaml:"appServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#app_server_version OpsworksJavaAppLayer#app_server_version}.
	AppServerVersion *string `json:"appServerVersion" yaml:"appServerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#auto_assign_elastic_ips OpsworksJavaAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#auto_assign_public_ips OpsworksJavaAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#auto_healing OpsworksJavaAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#cloudwatch_configuration OpsworksJavaAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksJavaAppLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_configure_recipes OpsworksJavaAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_deploy_recipes OpsworksJavaAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_instance_profile_arn OpsworksJavaAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_json OpsworksJavaAppLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_security_group_ids OpsworksJavaAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_setup_recipes OpsworksJavaAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_shutdown_recipes OpsworksJavaAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#custom_undeploy_recipes OpsworksJavaAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#drain_elb_on_shutdown OpsworksJavaAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#ebs_volume OpsworksJavaAppLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#elastic_load_balancer OpsworksJavaAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#install_updates_on_boot OpsworksJavaAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#instance_shutdown_timeout OpsworksJavaAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#jvm_options OpsworksJavaAppLayer#jvm_options}.
	JvmOptions *string `json:"jvmOptions" yaml:"jvmOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#jvm_type OpsworksJavaAppLayer#jvm_type}.
	JvmType *string `json:"jvmType" yaml:"jvmType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#jvm_version OpsworksJavaAppLayer#jvm_version}.
	JvmVersion *string `json:"jvmVersion" yaml:"jvmVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#name OpsworksJavaAppLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#system_packages OpsworksJavaAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#tags OpsworksJavaAppLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#tags_all OpsworksJavaAppLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#use_ebs_optimized_instances OpsworksJavaAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksJavaAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#mount_point OpsworksJavaAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#number_of_disks OpsworksJavaAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#size OpsworksJavaAppLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#encrypted OpsworksJavaAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#iops OpsworksJavaAppLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#raid_level OpsworksJavaAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_java_app_layer#type OpsworksJavaAppLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer aws_opsworks_memcached_layer}.
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
	CloudwatchConfiguration() OpsworksMemcachedLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksMemcachedLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksMemcachedLayerCloudwatchConfiguration)
	ResetAllocatedMemory()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksMemcachedLayer) CloudwatchConfiguration() OpsworksMemcachedLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksMemcachedLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) CloudwatchConfigurationInput() *OpsworksMemcachedLayerCloudwatchConfiguration {
	var returns *OpsworksMemcachedLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksMemcachedLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksMemcachedLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksMemcachedLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer aws_opsworks_memcached_layer} Resource.
func NewOpsworksMemcachedLayer(scope constructs.Construct, id *string, config *OpsworksMemcachedLayerConfig) OpsworksMemcachedLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMemcachedLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMemcachedLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer aws_opsworks_memcached_layer} Resource.
func NewOpsworksMemcachedLayer_Override(o OpsworksMemcachedLayer, scope constructs.Construct, id *string, config *OpsworksMemcachedLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMemcachedLayer",
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

func (j *jsiiProxy_OpsworksMemcachedLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksMemcachedLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksMemcachedLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksMemcachedLayer",
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
		"hashicorp_aws.opsworks.OpsworksMemcachedLayer",
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
func (o *jsiiProxy_OpsworksMemcachedLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksMemcachedLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksMemcachedLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksMemcachedLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksMemcachedLayer) PutCloudwatchConfiguration(value *OpsworksMemcachedLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksMemcachedLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksMemcachedLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#enabled OpsworksMemcachedLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#log_streams OpsworksMemcachedLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksMemcachedLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#file OpsworksMemcachedLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#log_group_name OpsworksMemcachedLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#batch_count OpsworksMemcachedLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#batch_size OpsworksMemcachedLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#buffer_duration OpsworksMemcachedLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#datetime_format OpsworksMemcachedLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#encoding OpsworksMemcachedLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#file_fingerprint_lines OpsworksMemcachedLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#initial_position OpsworksMemcachedLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#multiline_start_pattern OpsworksMemcachedLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#time_zone OpsworksMemcachedLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksMemcachedLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksMemcachedLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksMemcachedLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksMemcachedLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksMemcachedLayerCloudwatchConfiguration {
	var returns *OpsworksMemcachedLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksMemcachedLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksMemcachedLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMemcachedLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksMemcachedLayerCloudwatchConfigurationOutputReference_Override(o OpsworksMemcachedLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMemcachedLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksMemcachedLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksMemcachedLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#stack_id OpsworksMemcachedLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#allocated_memory OpsworksMemcachedLayer#allocated_memory}.
	AllocatedMemory *float64 `json:"allocatedMemory" yaml:"allocatedMemory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#auto_assign_elastic_ips OpsworksMemcachedLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#auto_assign_public_ips OpsworksMemcachedLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#auto_healing OpsworksMemcachedLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#cloudwatch_configuration OpsworksMemcachedLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksMemcachedLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_configure_recipes OpsworksMemcachedLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_deploy_recipes OpsworksMemcachedLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_instance_profile_arn OpsworksMemcachedLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_json OpsworksMemcachedLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_security_group_ids OpsworksMemcachedLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_setup_recipes OpsworksMemcachedLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_shutdown_recipes OpsworksMemcachedLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#custom_undeploy_recipes OpsworksMemcachedLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#drain_elb_on_shutdown OpsworksMemcachedLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#ebs_volume OpsworksMemcachedLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#elastic_load_balancer OpsworksMemcachedLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#install_updates_on_boot OpsworksMemcachedLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#instance_shutdown_timeout OpsworksMemcachedLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#name OpsworksMemcachedLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#system_packages OpsworksMemcachedLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#tags OpsworksMemcachedLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#tags_all OpsworksMemcachedLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#use_ebs_optimized_instances OpsworksMemcachedLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksMemcachedLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#mount_point OpsworksMemcachedLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#number_of_disks OpsworksMemcachedLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#size OpsworksMemcachedLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#encrypted OpsworksMemcachedLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#iops OpsworksMemcachedLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#raid_level OpsworksMemcachedLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_memcached_layer#type OpsworksMemcachedLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer aws_opsworks_mysql_layer}.
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
	CloudwatchConfiguration() OpsworksMysqlLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksMysqlLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksMysqlLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksMysqlLayer) CloudwatchConfiguration() OpsworksMysqlLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksMysqlLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) CloudwatchConfigurationInput() *OpsworksMysqlLayerCloudwatchConfiguration {
	var returns *OpsworksMysqlLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksMysqlLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksMysqlLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksMysqlLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer aws_opsworks_mysql_layer} Resource.
func NewOpsworksMysqlLayer(scope constructs.Construct, id *string, config *OpsworksMysqlLayerConfig) OpsworksMysqlLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMysqlLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMysqlLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer aws_opsworks_mysql_layer} Resource.
func NewOpsworksMysqlLayer_Override(o OpsworksMysqlLayer, scope constructs.Construct, id *string, config *OpsworksMysqlLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMysqlLayer",
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

func (j *jsiiProxy_OpsworksMysqlLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksMysqlLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksMysqlLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksMysqlLayer",
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
		"hashicorp_aws.opsworks.OpsworksMysqlLayer",
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
func (o *jsiiProxy_OpsworksMysqlLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksMysqlLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksMysqlLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksMysqlLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksMysqlLayer) PutCloudwatchConfiguration(value *OpsworksMysqlLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksMysqlLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksMysqlLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#enabled OpsworksMysqlLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#log_streams OpsworksMysqlLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksMysqlLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#file OpsworksMysqlLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#log_group_name OpsworksMysqlLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#batch_count OpsworksMysqlLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#batch_size OpsworksMysqlLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#buffer_duration OpsworksMysqlLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#datetime_format OpsworksMysqlLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#encoding OpsworksMysqlLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#file_fingerprint_lines OpsworksMysqlLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#initial_position OpsworksMysqlLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#multiline_start_pattern OpsworksMysqlLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#time_zone OpsworksMysqlLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksMysqlLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksMysqlLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksMysqlLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksMysqlLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksMysqlLayerCloudwatchConfiguration {
	var returns *OpsworksMysqlLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksMysqlLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksMysqlLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMysqlLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksMysqlLayerCloudwatchConfigurationOutputReference_Override(o OpsworksMysqlLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksMysqlLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksMysqlLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMysqlLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksMysqlLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#stack_id OpsworksMysqlLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#auto_assign_elastic_ips OpsworksMysqlLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#auto_assign_public_ips OpsworksMysqlLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#auto_healing OpsworksMysqlLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#cloudwatch_configuration OpsworksMysqlLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksMysqlLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_configure_recipes OpsworksMysqlLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_deploy_recipes OpsworksMysqlLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_instance_profile_arn OpsworksMysqlLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_json OpsworksMysqlLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_security_group_ids OpsworksMysqlLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_setup_recipes OpsworksMysqlLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_shutdown_recipes OpsworksMysqlLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#custom_undeploy_recipes OpsworksMysqlLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#drain_elb_on_shutdown OpsworksMysqlLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#ebs_volume OpsworksMysqlLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#elastic_load_balancer OpsworksMysqlLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#install_updates_on_boot OpsworksMysqlLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#instance_shutdown_timeout OpsworksMysqlLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#name OpsworksMysqlLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#root_password OpsworksMysqlLayer#root_password}.
	RootPassword *string `json:"rootPassword" yaml:"rootPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#root_password_on_all_instances OpsworksMysqlLayer#root_password_on_all_instances}.
	RootPasswordOnAllInstances interface{} `json:"rootPasswordOnAllInstances" yaml:"rootPasswordOnAllInstances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#system_packages OpsworksMysqlLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#tags OpsworksMysqlLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#tags_all OpsworksMysqlLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#use_ebs_optimized_instances OpsworksMysqlLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksMysqlLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#mount_point OpsworksMysqlLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#number_of_disks OpsworksMysqlLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#size OpsworksMysqlLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#encrypted OpsworksMysqlLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#iops OpsworksMysqlLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#raid_level OpsworksMysqlLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_mysql_layer#type OpsworksMysqlLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer}.
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
	CloudwatchConfiguration() OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksNodejsAppLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksNodejsAppLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) CloudwatchConfiguration() OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CloudwatchConfigurationInput() *OpsworksNodejsAppLayerCloudwatchConfiguration {
	var returns *OpsworksNodejsAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer(scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) OpsworksNodejsAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksNodejsAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer_Override(o OpsworksNodejsAppLayer, scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayer",
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayer",
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
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayer",
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
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksNodejsAppLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksNodejsAppLayer) PutCloudwatchConfiguration(value *OpsworksNodejsAppLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksNodejsAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#enabled OpsworksNodejsAppLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#log_streams OpsworksNodejsAppLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksNodejsAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#file OpsworksNodejsAppLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#log_group_name OpsworksNodejsAppLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#batch_count OpsworksNodejsAppLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#batch_size OpsworksNodejsAppLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#buffer_duration OpsworksNodejsAppLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#datetime_format OpsworksNodejsAppLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#encoding OpsworksNodejsAppLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#file_fingerprint_lines OpsworksNodejsAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#initial_position OpsworksNodejsAppLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#multiline_start_pattern OpsworksNodejsAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#time_zone OpsworksNodejsAppLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksNodejsAppLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksNodejsAppLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksNodejsAppLayerCloudwatchConfiguration {
	var returns *OpsworksNodejsAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksNodejsAppLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksNodejsAppLayerCloudwatchConfigurationOutputReference_Override(o OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksNodejsAppLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksNodejsAppLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#stack_id OpsworksNodejsAppLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#auto_assign_elastic_ips OpsworksNodejsAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#auto_assign_public_ips OpsworksNodejsAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#auto_healing OpsworksNodejsAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#cloudwatch_configuration OpsworksNodejsAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksNodejsAppLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_configure_recipes OpsworksNodejsAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_deploy_recipes OpsworksNodejsAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_instance_profile_arn OpsworksNodejsAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_json OpsworksNodejsAppLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_security_group_ids OpsworksNodejsAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_setup_recipes OpsworksNodejsAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_shutdown_recipes OpsworksNodejsAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#custom_undeploy_recipes OpsworksNodejsAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#drain_elb_on_shutdown OpsworksNodejsAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#ebs_volume OpsworksNodejsAppLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#elastic_load_balancer OpsworksNodejsAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#install_updates_on_boot OpsworksNodejsAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#instance_shutdown_timeout OpsworksNodejsAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#name OpsworksNodejsAppLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#nodejs_version OpsworksNodejsAppLayer#nodejs_version}.
	NodejsVersion *string `json:"nodejsVersion" yaml:"nodejsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#system_packages OpsworksNodejsAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#tags OpsworksNodejsAppLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#tags_all OpsworksNodejsAppLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#use_ebs_optimized_instances OpsworksNodejsAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksNodejsAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#mount_point OpsworksNodejsAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#number_of_disks OpsworksNodejsAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#size OpsworksNodejsAppLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#encrypted OpsworksNodejsAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#iops OpsworksNodejsAppLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#raid_level OpsworksNodejsAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_nodejs_app_layer#type OpsworksNodejsAppLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission aws_opsworks_permission}.
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
	Count() *float64
	SetCount(val *float64)
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

func (j *jsiiProxy_OpsworksPermission) Count() *float64 {
	var returns *float64
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission aws_opsworks_permission} Resource.
func NewOpsworksPermission(scope constructs.Construct, id *string, config *OpsworksPermissionConfig) OpsworksPermission {
	_init_.Initialize()

	j := jsiiProxy_OpsworksPermission{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission aws_opsworks_permission} Resource.
func NewOpsworksPermission_Override(o OpsworksPermission, scope constructs.Construct, id *string, config *OpsworksPermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPermission",
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

func (j *jsiiProxy_OpsworksPermission) SetCount(val *float64) {
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
		"hashicorp_aws.opsworks.OpsworksPermission",
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
		"hashicorp_aws.opsworks.OpsworksPermission",
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
func (o *jsiiProxy_OpsworksPermission) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksPermission) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksPermission) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPermission) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksPermission) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// AWS OpsWorks.
type OpsworksPermissionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission#user_arn OpsworksPermission#user_arn}.
	UserArn *string `json:"userArn" yaml:"userArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission#allow_ssh OpsworksPermission#allow_ssh}.
	AllowSsh interface{} `json:"allowSsh" yaml:"allowSsh"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission#allow_sudo OpsworksPermission#allow_sudo}.
	AllowSudo interface{} `json:"allowSudo" yaml:"allowSudo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission#level OpsworksPermission#level}.
	Level *string `json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_permission#stack_id OpsworksPermission#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer aws_opsworks_php_app_layer}.
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
	CloudwatchConfiguration() OpsworksPhpAppLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksPhpAppLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksPhpAppLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksPhpAppLayer) CloudwatchConfiguration() OpsworksPhpAppLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksPhpAppLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) CloudwatchConfigurationInput() *OpsworksPhpAppLayerCloudwatchConfiguration {
	var returns *OpsworksPhpAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksPhpAppLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksPhpAppLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksPhpAppLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer aws_opsworks_php_app_layer} Resource.
func NewOpsworksPhpAppLayer(scope constructs.Construct, id *string, config *OpsworksPhpAppLayerConfig) OpsworksPhpAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksPhpAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPhpAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer aws_opsworks_php_app_layer} Resource.
func NewOpsworksPhpAppLayer_Override(o OpsworksPhpAppLayer, scope constructs.Construct, id *string, config *OpsworksPhpAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPhpAppLayer",
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

func (j *jsiiProxy_OpsworksPhpAppLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksPhpAppLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksPhpAppLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksPhpAppLayer",
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
		"hashicorp_aws.opsworks.OpsworksPhpAppLayer",
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
func (o *jsiiProxy_OpsworksPhpAppLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksPhpAppLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksPhpAppLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksPhpAppLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksPhpAppLayer) PutCloudwatchConfiguration(value *OpsworksPhpAppLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksPhpAppLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksPhpAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#enabled OpsworksPhpAppLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#log_streams OpsworksPhpAppLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksPhpAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#file OpsworksPhpAppLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#log_group_name OpsworksPhpAppLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#batch_count OpsworksPhpAppLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#batch_size OpsworksPhpAppLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#buffer_duration OpsworksPhpAppLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#datetime_format OpsworksPhpAppLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#encoding OpsworksPhpAppLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#file_fingerprint_lines OpsworksPhpAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#initial_position OpsworksPhpAppLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#multiline_start_pattern OpsworksPhpAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#time_zone OpsworksPhpAppLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksPhpAppLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksPhpAppLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksPhpAppLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksPhpAppLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksPhpAppLayerCloudwatchConfiguration {
	var returns *OpsworksPhpAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksPhpAppLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksPhpAppLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPhpAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksPhpAppLayerCloudwatchConfigurationOutputReference_Override(o OpsworksPhpAppLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksPhpAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksPhpAppLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksPhpAppLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksPhpAppLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#stack_id OpsworksPhpAppLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#auto_assign_elastic_ips OpsworksPhpAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#auto_assign_public_ips OpsworksPhpAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#auto_healing OpsworksPhpAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#cloudwatch_configuration OpsworksPhpAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksPhpAppLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_configure_recipes OpsworksPhpAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_deploy_recipes OpsworksPhpAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_instance_profile_arn OpsworksPhpAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_json OpsworksPhpAppLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_security_group_ids OpsworksPhpAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_setup_recipes OpsworksPhpAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_shutdown_recipes OpsworksPhpAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#custom_undeploy_recipes OpsworksPhpAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#drain_elb_on_shutdown OpsworksPhpAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#ebs_volume OpsworksPhpAppLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#elastic_load_balancer OpsworksPhpAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#install_updates_on_boot OpsworksPhpAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#instance_shutdown_timeout OpsworksPhpAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#name OpsworksPhpAppLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#system_packages OpsworksPhpAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#tags OpsworksPhpAppLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#tags_all OpsworksPhpAppLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#use_ebs_optimized_instances OpsworksPhpAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksPhpAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#mount_point OpsworksPhpAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#number_of_disks OpsworksPhpAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#size OpsworksPhpAppLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#encrypted OpsworksPhpAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#iops OpsworksPhpAppLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#raid_level OpsworksPhpAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_php_app_layer#type OpsworksPhpAppLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer aws_opsworks_rails_app_layer}.
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
	CloudwatchConfiguration() OpsworksRailsAppLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksRailsAppLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksRailsAppLayerCloudwatchConfiguration)
	ResetAppServer()
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetBundlerVersion()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksRailsAppLayer) CloudwatchConfiguration() OpsworksRailsAppLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksRailsAppLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) CloudwatchConfigurationInput() *OpsworksRailsAppLayerCloudwatchConfiguration {
	var returns *OpsworksRailsAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksRailsAppLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksRailsAppLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksRailsAppLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer aws_opsworks_rails_app_layer} Resource.
func NewOpsworksRailsAppLayer(scope constructs.Construct, id *string, config *OpsworksRailsAppLayerConfig) OpsworksRailsAppLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksRailsAppLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRailsAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer aws_opsworks_rails_app_layer} Resource.
func NewOpsworksRailsAppLayer_Override(o OpsworksRailsAppLayer, scope constructs.Construct, id *string, config *OpsworksRailsAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRailsAppLayer",
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

func (j *jsiiProxy_OpsworksRailsAppLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksRailsAppLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksRailsAppLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksRailsAppLayer",
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
		"hashicorp_aws.opsworks.OpsworksRailsAppLayer",
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
func (o *jsiiProxy_OpsworksRailsAppLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksRailsAppLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksRailsAppLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksRailsAppLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksRailsAppLayer) PutCloudwatchConfiguration(value *OpsworksRailsAppLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksRailsAppLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksRailsAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#enabled OpsworksRailsAppLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#log_streams OpsworksRailsAppLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksRailsAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#file OpsworksRailsAppLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#log_group_name OpsworksRailsAppLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#batch_count OpsworksRailsAppLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#batch_size OpsworksRailsAppLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#buffer_duration OpsworksRailsAppLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#datetime_format OpsworksRailsAppLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#encoding OpsworksRailsAppLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#file_fingerprint_lines OpsworksRailsAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#initial_position OpsworksRailsAppLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#multiline_start_pattern OpsworksRailsAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#time_zone OpsworksRailsAppLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksRailsAppLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksRailsAppLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksRailsAppLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksRailsAppLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksRailsAppLayerCloudwatchConfiguration {
	var returns *OpsworksRailsAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksRailsAppLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksRailsAppLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRailsAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksRailsAppLayerCloudwatchConfigurationOutputReference_Override(o OpsworksRailsAppLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRailsAppLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksRailsAppLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksRailsAppLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#stack_id OpsworksRailsAppLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#app_server OpsworksRailsAppLayer#app_server}.
	AppServer *string `json:"appServer" yaml:"appServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_assign_elastic_ips OpsworksRailsAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_assign_public_ips OpsworksRailsAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_healing OpsworksRailsAppLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#bundler_version OpsworksRailsAppLayer#bundler_version}.
	BundlerVersion *string `json:"bundlerVersion" yaml:"bundlerVersion"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#cloudwatch_configuration OpsworksRailsAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksRailsAppLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_configure_recipes OpsworksRailsAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_deploy_recipes OpsworksRailsAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_instance_profile_arn OpsworksRailsAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_json OpsworksRailsAppLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_security_group_ids OpsworksRailsAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_setup_recipes OpsworksRailsAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_shutdown_recipes OpsworksRailsAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_undeploy_recipes OpsworksRailsAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#drain_elb_on_shutdown OpsworksRailsAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#ebs_volume OpsworksRailsAppLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#elastic_load_balancer OpsworksRailsAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#install_updates_on_boot OpsworksRailsAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#instance_shutdown_timeout OpsworksRailsAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#manage_bundler OpsworksRailsAppLayer#manage_bundler}.
	ManageBundler interface{} `json:"manageBundler" yaml:"manageBundler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#name OpsworksRailsAppLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#passenger_version OpsworksRailsAppLayer#passenger_version}.
	PassengerVersion *string `json:"passengerVersion" yaml:"passengerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#rubygems_version OpsworksRailsAppLayer#rubygems_version}.
	RubygemsVersion *string `json:"rubygemsVersion" yaml:"rubygemsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#ruby_version OpsworksRailsAppLayer#ruby_version}.
	RubyVersion *string `json:"rubyVersion" yaml:"rubyVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#system_packages OpsworksRailsAppLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#tags OpsworksRailsAppLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#tags_all OpsworksRailsAppLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#use_ebs_optimized_instances OpsworksRailsAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksRailsAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#mount_point OpsworksRailsAppLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#number_of_disks OpsworksRailsAppLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#size OpsworksRailsAppLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#encrypted OpsworksRailsAppLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#iops OpsworksRailsAppLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#raid_level OpsworksRailsAppLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#type OpsworksRailsAppLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance aws_opsworks_rds_db_instance}.
type OpsworksRdsDbInstance interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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

func (j *jsiiProxy_OpsworksRdsDbInstance) Count() *float64 {
	var returns *float64
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance aws_opsworks_rds_db_instance} Resource.
func NewOpsworksRdsDbInstance(scope constructs.Construct, id *string, config *OpsworksRdsDbInstanceConfig) OpsworksRdsDbInstance {
	_init_.Initialize()

	j := jsiiProxy_OpsworksRdsDbInstance{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRdsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance aws_opsworks_rds_db_instance} Resource.
func NewOpsworksRdsDbInstance_Override(o OpsworksRdsDbInstance, scope constructs.Construct, id *string, config *OpsworksRdsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksRdsDbInstance",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksRdsDbInstance) SetCount(val *float64) {
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
		"hashicorp_aws.opsworks.OpsworksRdsDbInstance",
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
		"hashicorp_aws.opsworks.OpsworksRdsDbInstance",
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
func (o *jsiiProxy_OpsworksRdsDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksRdsDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksRdsDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksRdsDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksRdsDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// AWS OpsWorks.
type OpsworksRdsDbInstanceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance#db_password OpsworksRdsDbInstance#db_password}.
	DbPassword *string `json:"dbPassword" yaml:"dbPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance#db_user OpsworksRdsDbInstance#db_user}.
	DbUser *string `json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance#rds_db_instance_arn OpsworksRdsDbInstance#rds_db_instance_arn}.
	RdsDbInstanceArn *string `json:"rdsDbInstanceArn" yaml:"rdsDbInstanceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rds_db_instance#stack_id OpsworksRdsDbInstance#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack aws_opsworks_stack}.
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
	Count() *float64
	SetCount(val *float64)
	CustomCookbooksSource() interface{}
	SetCustomCookbooksSource(val interface{})
	CustomCookbooksSourceInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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

func (j *jsiiProxy_OpsworksStack) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSourceInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksStack) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack aws_opsworks_stack} Resource.
func NewOpsworksStack(scope constructs.Construct, id *string, config *OpsworksStackConfig) OpsworksStack {
	_init_.Initialize()

	j := jsiiProxy_OpsworksStack{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack aws_opsworks_stack} Resource.
func NewOpsworksStack_Override(o OpsworksStack, scope constructs.Construct, id *string, config *OpsworksStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStack",
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

func (j *jsiiProxy_OpsworksStack) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetCustomCookbooksSource(val interface{}) {
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

func (j *jsiiProxy_OpsworksStack) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksStack",
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
		"hashicorp_aws.opsworks.OpsworksStack",
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
func (o *jsiiProxy_OpsworksStack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksStack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksStack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksStack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// AWS OpsWorks.
type OpsworksStackConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_instance_profile_arn OpsworksStack#default_instance_profile_arn}.
	DefaultInstanceProfileArn *string `json:"defaultInstanceProfileArn" yaml:"defaultInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#name OpsworksStack#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#region OpsworksStack#region}.
	Region *string `json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#service_role_arn OpsworksStack#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#agent_version OpsworksStack#agent_version}.
	AgentVersion *string `json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#berkshelf_version OpsworksStack#berkshelf_version}.
	BerkshelfVersion *string `json:"berkshelfVersion" yaml:"berkshelfVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#color OpsworksStack#color}.
	Color *string `json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#configuration_manager_name OpsworksStack#configuration_manager_name}.
	ConfigurationManagerName *string `json:"configurationManagerName" yaml:"configurationManagerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#configuration_manager_version OpsworksStack#configuration_manager_version}.
	ConfigurationManagerVersion *string `json:"configurationManagerVersion" yaml:"configurationManagerVersion"`
	// custom_cookbooks_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#custom_cookbooks_source OpsworksStack#custom_cookbooks_source}
	CustomCookbooksSource interface{} `json:"customCookbooksSource" yaml:"customCookbooksSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#custom_json OpsworksStack#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_availability_zone OpsworksStack#default_availability_zone}.
	DefaultAvailabilityZone *string `json:"defaultAvailabilityZone" yaml:"defaultAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_os OpsworksStack#default_os}.
	DefaultOs *string `json:"defaultOs" yaml:"defaultOs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_root_device_type OpsworksStack#default_root_device_type}.
	DefaultRootDeviceType *string `json:"defaultRootDeviceType" yaml:"defaultRootDeviceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_ssh_key_name OpsworksStack#default_ssh_key_name}.
	DefaultSshKeyName *string `json:"defaultSshKeyName" yaml:"defaultSshKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#default_subnet_id OpsworksStack#default_subnet_id}.
	DefaultSubnetId *string `json:"defaultSubnetId" yaml:"defaultSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#hostname_theme OpsworksStack#hostname_theme}.
	HostnameTheme *string `json:"hostnameTheme" yaml:"hostnameTheme"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#manage_berkshelf OpsworksStack#manage_berkshelf}.
	ManageBerkshelf interface{} `json:"manageBerkshelf" yaml:"manageBerkshelf"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#tags OpsworksStack#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#tags_all OpsworksStack#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#use_custom_cookbooks OpsworksStack#use_custom_cookbooks}.
	UseCustomCookbooks interface{} `json:"useCustomCookbooks" yaml:"useCustomCookbooks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#use_opsworks_security_groups OpsworksStack#use_opsworks_security_groups}.
	UseOpsworksSecurityGroups interface{} `json:"useOpsworksSecurityGroups" yaml:"useOpsworksSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#vpc_id OpsworksStack#vpc_id}.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

type OpsworksStackCustomCookbooksSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#type OpsworksStack#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#url OpsworksStack#url}.
	Url *string `json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#password OpsworksStack#password}.
	Password *string `json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#revision OpsworksStack#revision}.
	Revision *string `json:"revision" yaml:"revision"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#ssh_key OpsworksStack#ssh_key}.
	SshKey *string `json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#username OpsworksStack#username}.
	Username *string `json:"username" yaml:"username"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer aws_opsworks_static_web_layer}.
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
	CloudwatchConfiguration() OpsworksStaticWebLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksStaticWebLayerCloudwatchConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	EbsVolume() interface{}
	SetEbsVolume(val interface{})
	EbsVolumeInput() interface{}
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksStaticWebLayerCloudwatchConfiguration)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
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

func (j *jsiiProxy_OpsworksStaticWebLayer) CloudwatchConfiguration() OpsworksStaticWebLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksStaticWebLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) CloudwatchConfigurationInput() *OpsworksStaticWebLayerCloudwatchConfiguration {
	var returns *OpsworksStaticWebLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
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

func (j *jsiiProxy_OpsworksStaticWebLayer) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_OpsworksStaticWebLayer) EbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) EbsVolumeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OpsworksStaticWebLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer aws_opsworks_static_web_layer} Resource.
func NewOpsworksStaticWebLayer(scope constructs.Construct, id *string, config *OpsworksStaticWebLayerConfig) OpsworksStaticWebLayer {
	_init_.Initialize()

	j := jsiiProxy_OpsworksStaticWebLayer{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStaticWebLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer aws_opsworks_static_web_layer} Resource.
func NewOpsworksStaticWebLayer_Override(o OpsworksStaticWebLayer, scope constructs.Construct, id *string, config *OpsworksStaticWebLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStaticWebLayer",
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

func (j *jsiiProxy_OpsworksStaticWebLayer) SetCount(val *float64) {
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

func (j *jsiiProxy_OpsworksStaticWebLayer) SetEbsVolume(val interface{}) {
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

func (j *jsiiProxy_OpsworksStaticWebLayer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayer) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.opsworks.OpsworksStaticWebLayer",
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
		"hashicorp_aws.opsworks.OpsworksStaticWebLayer",
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
func (o *jsiiProxy_OpsworksStaticWebLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksStaticWebLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksStaticWebLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksStaticWebLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

func (o *jsiiProxy_OpsworksStaticWebLayer) PutCloudwatchConfiguration(value *OpsworksStaticWebLayerCloudwatchConfiguration) {
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
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

func (o *jsiiProxy_OpsworksStaticWebLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
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

type OpsworksStaticWebLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#enabled OpsworksStaticWebLayer#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#log_streams OpsworksStaticWebLayer#log_streams}
	LogStreams interface{} `json:"logStreams" yaml:"logStreams"`
}

type OpsworksStaticWebLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#file OpsworksStaticWebLayer#file}.
	File *string `json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#log_group_name OpsworksStaticWebLayer#log_group_name}.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#batch_count OpsworksStaticWebLayer#batch_count}.
	BatchCount *float64 `json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#batch_size OpsworksStaticWebLayer#batch_size}.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#buffer_duration OpsworksStaticWebLayer#buffer_duration}.
	BufferDuration *float64 `json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#datetime_format OpsworksStaticWebLayer#datetime_format}.
	DatetimeFormat *string `json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#encoding OpsworksStaticWebLayer#encoding}.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#file_fingerprint_lines OpsworksStaticWebLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#initial_position OpsworksStaticWebLayer#initial_position}.
	InitialPosition *string `json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#multiline_start_pattern OpsworksStaticWebLayer#multiline_start_pattern}.
	MultilineStartPattern *string `json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#time_zone OpsworksStaticWebLayer#time_zone}.
	TimeZone *string `json:"timeZone" yaml:"timeZone"`
}

type OpsworksStaticWebLayerCloudwatchConfigurationOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *OpsworksStaticWebLayerCloudwatchConfiguration
	SetInternalValue(val *OpsworksStaticWebLayerCloudwatchConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogStreams() interface{}
	SetLogStreams(val interface{})
	LogStreamsInput() interface{}
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
	ResetEnabled()
	ResetLogStreams()
}

// The jsii proxy struct for OpsworksStaticWebLayerCloudwatchConfigurationOutputReference
type jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) InternalValue() *OpsworksStaticWebLayerCloudwatchConfiguration {
	var returns *OpsworksStaticWebLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) LogStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) LogStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpsworksStaticWebLayerCloudwatchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) OpsworksStaticWebLayerCloudwatchConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStaticWebLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewOpsworksStaticWebLayerCloudwatchConfigurationOutputReference_Override(o OpsworksStaticWebLayerCloudwatchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksStaticWebLayerCloudwatchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		o,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetInternalValue(val *OpsworksStaticWebLayerCloudwatchConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetLogStreams(val interface{}) {
	_jsii_.Set(
		j,
		"logStreams",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStaticWebLayerCloudwatchConfigurationOutputReference) ResetLogStreams() {
	_jsii_.InvokeVoid(
		o,
		"resetLogStreams",
		nil, // no parameters
	)
}

// AWS OpsWorks.
type OpsworksStaticWebLayerConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#stack_id OpsworksStaticWebLayer#stack_id}.
	StackId *string `json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#auto_assign_elastic_ips OpsworksStaticWebLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#auto_assign_public_ips OpsworksStaticWebLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#auto_healing OpsworksStaticWebLayer#auto_healing}.
	AutoHealing interface{} `json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#cloudwatch_configuration OpsworksStaticWebLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksStaticWebLayerCloudwatchConfiguration `json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_configure_recipes OpsworksStaticWebLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_deploy_recipes OpsworksStaticWebLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_instance_profile_arn OpsworksStaticWebLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_json OpsworksStaticWebLayer#custom_json}.
	CustomJson *string `json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_security_group_ids OpsworksStaticWebLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_setup_recipes OpsworksStaticWebLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_shutdown_recipes OpsworksStaticWebLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#custom_undeploy_recipes OpsworksStaticWebLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#drain_elb_on_shutdown OpsworksStaticWebLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#ebs_volume OpsworksStaticWebLayer#ebs_volume}
	EbsVolume interface{} `json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#elastic_load_balancer OpsworksStaticWebLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#install_updates_on_boot OpsworksStaticWebLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#instance_shutdown_timeout OpsworksStaticWebLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#name OpsworksStaticWebLayer#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#system_packages OpsworksStaticWebLayer#system_packages}.
	SystemPackages *[]*string `json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#tags OpsworksStaticWebLayer#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#tags_all OpsworksStaticWebLayer#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#use_ebs_optimized_instances OpsworksStaticWebLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

type OpsworksStaticWebLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#mount_point OpsworksStaticWebLayer#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#number_of_disks OpsworksStaticWebLayer#number_of_disks}.
	NumberOfDisks *float64 `json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#size OpsworksStaticWebLayer#size}.
	Size *float64 `json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#encrypted OpsworksStaticWebLayer#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#iops OpsworksStaticWebLayer#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#raid_level OpsworksStaticWebLayer#raid_level}.
	RaidLevel *string `json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_static_web_layer#type OpsworksStaticWebLayer#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile aws_opsworks_user_profile}.
type OpsworksUserProfile interface {
	cdktf.TerraformResource
	AllowSelfManagement() interface{}
	SetAllowSelfManagement(val interface{})
	AllowSelfManagementInput() interface{}
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

func (j *jsiiProxy_OpsworksUserProfile) Count() *float64 {
	var returns *float64
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile aws_opsworks_user_profile} Resource.
func NewOpsworksUserProfile(scope constructs.Construct, id *string, config *OpsworksUserProfileConfig) OpsworksUserProfile {
	_init_.Initialize()

	j := jsiiProxy_OpsworksUserProfile{}

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksUserProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile aws_opsworks_user_profile} Resource.
func NewOpsworksUserProfile_Override(o OpsworksUserProfile, scope constructs.Construct, id *string, config *OpsworksUserProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.opsworks.OpsworksUserProfile",
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

func (j *jsiiProxy_OpsworksUserProfile) SetCount(val *float64) {
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
		"hashicorp_aws.opsworks.OpsworksUserProfile",
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
		"hashicorp_aws.opsworks.OpsworksUserProfile",
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
func (o *jsiiProxy_OpsworksUserProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (o *jsiiProxy_OpsworksUserProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
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
func (o *jsiiProxy_OpsworksUserProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OpsworksUserProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
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
func (o *jsiiProxy_OpsworksUserProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
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

// AWS OpsWorks.
type OpsworksUserProfileConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile#ssh_username OpsworksUserProfile#ssh_username}.
	SshUsername *string `json:"sshUsername" yaml:"sshUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile#user_arn OpsworksUserProfile#user_arn}.
	UserArn *string `json:"userArn" yaml:"userArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile#allow_self_management OpsworksUserProfile#allow_self_management}.
	AllowSelfManagement interface{} `json:"allowSelfManagement" yaml:"allowSelfManagement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_user_profile#ssh_public_key OpsworksUserProfile#ssh_public_key}.
	SshPublicKey *string `json:"sshPublicKey" yaml:"sshPublicKey"`
}
