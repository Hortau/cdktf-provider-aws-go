package mwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/mwaa/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment aws_mwaa_environment}.
type MwaaEnvironment interface {
	cdktf.TerraformResource
	AirflowConfigurationOptions() *map[string]*string
	SetAirflowConfigurationOptions(val *map[string]*string)
	AirflowConfigurationOptionsInput() *map[string]*string
	AirflowVersion() *string
	SetAirflowVersion(val *string)
	AirflowVersionInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedAt() *string
	DagS3Path() *string
	SetDagS3Path(val *string)
	DagS3PathInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnvironmentClass() *string
	SetEnvironmentClass(val *string)
	EnvironmentClassInput() *string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfiguration() MwaaEnvironmentLoggingConfigurationOutputReference
	LoggingConfigurationInput() *MwaaEnvironmentLoggingConfiguration
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MaxWorkersInput() *float64
	MinWorkers() *float64
	SetMinWorkers(val *float64)
	MinWorkersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfiguration() MwaaEnvironmentNetworkConfigurationOutputReference
	NetworkConfigurationInput() *MwaaEnvironmentNetworkConfiguration
	Node() constructs.Node
	PluginsS3ObjectVersion() *string
	SetPluginsS3ObjectVersion(val *string)
	PluginsS3ObjectVersionInput() *string
	PluginsS3Path() *string
	SetPluginsS3Path(val *string)
	PluginsS3PathInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequirementsS3ObjectVersion() *string
	SetRequirementsS3ObjectVersion(val *string)
	RequirementsS3ObjectVersionInput() *string
	RequirementsS3Path() *string
	SetRequirementsS3Path(val *string)
	RequirementsS3PathInput() *string
	ServiceRoleArn() *string
	SourceBucketArn() *string
	SetSourceBucketArn(val *string)
	SourceBucketArnInput() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WebserverAccessMode() *string
	SetWebserverAccessMode(val *string)
	WebserverAccessModeInput() *string
	WebserverUrl() *string
	WeeklyMaintenanceWindowStart() *string
	SetWeeklyMaintenanceWindowStart(val *string)
	WeeklyMaintenanceWindowStartInput() *string
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
	LastUpdated(index *string) MwaaEnvironmentLastUpdated
	OverrideLogicalId(newLogicalId *string)
	PutLoggingConfiguration(value *MwaaEnvironmentLoggingConfiguration)
	PutNetworkConfiguration(value *MwaaEnvironmentNetworkConfiguration)
	ResetAirflowConfigurationOptions()
	ResetAirflowVersion()
	ResetEnvironmentClass()
	ResetKmsKey()
	ResetLoggingConfiguration()
	ResetMaxWorkers()
	ResetMinWorkers()
	ResetOverrideLogicalId()
	ResetPluginsS3ObjectVersion()
	ResetPluginsS3Path()
	ResetRequirementsS3ObjectVersion()
	ResetRequirementsS3Path()
	ResetTags()
	ResetTagsAll()
	ResetWebserverAccessMode()
	ResetWeeklyMaintenanceWindowStart()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MwaaEnvironment
type jsiiProxy_MwaaEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MwaaEnvironment) AirflowConfigurationOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowConfigurationOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DagS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EnvironmentClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EnvironmentClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) LoggingConfiguration() MwaaEnvironmentLoggingConfigurationOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) LoggingConfigurationInput() *MwaaEnvironmentLoggingConfiguration {
	var returns *MwaaEnvironmentLoggingConfiguration
	_jsii_.Get(
		j,
		"loggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NetworkConfiguration() MwaaEnvironmentNetworkConfigurationOutputReference {
	var returns MwaaEnvironmentNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NetworkConfigurationInput() *MwaaEnvironmentNetworkConfiguration {
	var returns *MwaaEnvironmentNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) SourceBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) SourceBucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WeeklyMaintenanceWindowStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStartInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment aws_mwaa_environment} Resource.
func NewMwaaEnvironment(scope constructs.Construct, id *string, config *MwaaEnvironmentConfig) MwaaEnvironment {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment aws_mwaa_environment} Resource.
func NewMwaaEnvironment_Override(m MwaaEnvironment, scope constructs.Construct, id *string, config *MwaaEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironment",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetAirflowConfigurationOptions(val *map[string]*string) {
	_jsii_.Set(
		j,
		"airflowConfigurationOptions",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetAirflowVersion(val *string) {
	_jsii_.Set(
		j,
		"airflowVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetDagS3Path(val *string) {
	_jsii_.Set(
		j,
		"dagS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetEnvironmentClass(val *string) {
	_jsii_.Set(
		j,
		"environmentClass",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetMaxWorkers(val *float64) {
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetMinWorkers(val *float64) {
	_jsii_.Set(
		j,
		"minWorkers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetPluginsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetPluginsS3Path(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetRequirementsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetRequirementsS3Path(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetSourceBucketArn(val *string) {
	_jsii_.Set(
		j,
		"sourceBucketArn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetWebserverAccessMode(val *string) {
	_jsii_.Set(
		j,
		"webserverAccessMode",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment) SetWeeklyMaintenanceWindowStart(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceWindowStart",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func MwaaEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.mwaa.MwaaEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwaaEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.mwaa.MwaaEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) LastUpdated(index *string) MwaaEnvironmentLastUpdated {
	var returns MwaaEnvironmentLastUpdated

	_jsii_.Invoke(
		m,
		"lastUpdated",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MwaaEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwaaEnvironment) PutLoggingConfiguration(value *MwaaEnvironmentLoggingConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putLoggingConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironment) PutNetworkConfiguration(value *MwaaEnvironmentNetworkConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetAirflowConfigurationOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetAirflowConfigurationOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetAirflowVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetAirflowVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetEnvironmentClass() {
	_jsii_.InvokeVoid(
		m,
		"resetEnvironmentClass",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetKmsKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetLoggingConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetLoggingConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMinWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMinWorkers",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MwaaEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetPluginsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginsS3ObjectVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetPluginsS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginsS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetRequirementsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetRequirementsS3ObjectVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetRequirementsS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetRequirementsS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetWebserverAccessMode() {
	_jsii_.InvokeVoid(
		m,
		"resetWebserverAccessMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetWeeklyMaintenanceWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklyMaintenanceWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MwaaEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MwaaEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Managed Workloads for Apache Airflow.
type MwaaEnvironmentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#dag_s3_path MwaaEnvironment#dag_s3_path}.
	DagS3Path *string `json:"dagS3Path" yaml:"dagS3Path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#execution_role_arn MwaaEnvironment#execution_role_arn}.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#name MwaaEnvironment#name}.
	Name *string `json:"name" yaml:"name"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#network_configuration MwaaEnvironment#network_configuration}
	NetworkConfiguration *MwaaEnvironmentNetworkConfiguration `json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#source_bucket_arn MwaaEnvironment#source_bucket_arn}.
	SourceBucketArn *string `json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#airflow_configuration_options MwaaEnvironment#airflow_configuration_options}.
	AirflowConfigurationOptions *map[string]*string `json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#airflow_version MwaaEnvironment#airflow_version}.
	AirflowVersion *string `json:"airflowVersion" yaml:"airflowVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#environment_class MwaaEnvironment#environment_class}.
	EnvironmentClass *string `json:"environmentClass" yaml:"environmentClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#kms_key MwaaEnvironment#kms_key}.
	KmsKey *string `json:"kmsKey" yaml:"kmsKey"`
	// logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#logging_configuration MwaaEnvironment#logging_configuration}
	LoggingConfiguration *MwaaEnvironmentLoggingConfiguration `json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#max_workers MwaaEnvironment#max_workers}.
	MaxWorkers *float64 `json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#min_workers MwaaEnvironment#min_workers}.
	MinWorkers *float64 `json:"minWorkers" yaml:"minWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#plugins_s3_object_version MwaaEnvironment#plugins_s3_object_version}.
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#plugins_s3_path MwaaEnvironment#plugins_s3_path}.
	PluginsS3Path *string `json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#requirements_s3_object_version MwaaEnvironment#requirements_s3_object_version}.
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#requirements_s3_path MwaaEnvironment#requirements_s3_path}.
	RequirementsS3Path *string `json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#tags MwaaEnvironment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#tags_all MwaaEnvironment#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#webserver_access_mode MwaaEnvironment#webserver_access_mode}.
	WebserverAccessMode *string `json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#weekly_maintenance_window_start MwaaEnvironment#weekly_maintenance_window_start}.
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

type MwaaEnvironmentLastUpdated interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CreatedAt() *string
	Error() cdktf.IResolvable
	Status() *string
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

// The jsii proxy struct for MwaaEnvironmentLastUpdated
type jsiiProxy_MwaaEnvironmentLastUpdated struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) Error() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewMwaaEnvironmentLastUpdated(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) MwaaEnvironmentLastUpdated {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLastUpdated{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdated",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewMwaaEnvironmentLastUpdated_Override(m MwaaEnvironmentLastUpdated, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdated",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdated) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdated) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type MwaaEnvironmentLastUpdatedError interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ErrorCode() *string
	ErrorMessage() *string
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

// The jsii proxy struct for MwaaEnvironmentLastUpdatedError
type jsiiProxy_MwaaEnvironmentLastUpdatedError struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) ErrorCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewMwaaEnvironmentLastUpdatedError(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) MwaaEnvironmentLastUpdatedError {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLastUpdatedError{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdatedError",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewMwaaEnvironmentLastUpdatedError_Override(m MwaaEnvironmentLastUpdatedError, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdatedError",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLastUpdatedError) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLastUpdatedError) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type MwaaEnvironmentLoggingConfiguration struct {
	// dag_processing_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#dag_processing_logs MwaaEnvironment#dag_processing_logs}
	DagProcessingLogs *MwaaEnvironmentLoggingConfigurationDagProcessingLogs `json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// scheduler_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#scheduler_logs MwaaEnvironment#scheduler_logs}
	SchedulerLogs *MwaaEnvironmentLoggingConfigurationSchedulerLogs `json:"schedulerLogs" yaml:"schedulerLogs"`
	// task_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#task_logs MwaaEnvironment#task_logs}
	TaskLogs *MwaaEnvironmentLoggingConfigurationTaskLogs `json:"taskLogs" yaml:"taskLogs"`
	// webserver_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#webserver_logs MwaaEnvironment#webserver_logs}
	WebserverLogs *MwaaEnvironmentLoggingConfigurationWebserverLogs `json:"webserverLogs" yaml:"webserverLogs"`
	// worker_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#worker_logs MwaaEnvironment#worker_logs}
	WorkerLogs *MwaaEnvironmentLoggingConfigurationWorkerLogs `json:"workerLogs" yaml:"workerLogs"`
}

type MwaaEnvironmentLoggingConfigurationDagProcessingLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

type MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogGroupArn() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	SetInternalValue(val *MwaaEnvironmentLoggingConfigurationDagProcessingLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
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
	ResetLogLevel()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) InternalValue() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs {
	var returns *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference_Override(m MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfigurationDagProcessingLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

type MwaaEnvironmentLoggingConfigurationOutputReference interface {
	cdktf.ComplexObject
	DagProcessingLogs() MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference
	DagProcessingLogsInput() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	InternalValue() *MwaaEnvironmentLoggingConfiguration
	SetInternalValue(val *MwaaEnvironmentLoggingConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SchedulerLogs() MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference
	SchedulerLogsInput() *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	TaskLogs() MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference
	TaskLogsInput() *MwaaEnvironmentLoggingConfigurationTaskLogs
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebserverLogs() MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference
	WebserverLogsInput() *MwaaEnvironmentLoggingConfigurationWebserverLogs
	WorkerLogs() MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference
	WorkerLogsInput() *MwaaEnvironmentLoggingConfigurationWorkerLogs
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
	PutDagProcessingLogs(value *MwaaEnvironmentLoggingConfigurationDagProcessingLogs)
	PutSchedulerLogs(value *MwaaEnvironmentLoggingConfigurationSchedulerLogs)
	PutTaskLogs(value *MwaaEnvironmentLoggingConfigurationTaskLogs)
	PutWebserverLogs(value *MwaaEnvironmentLoggingConfigurationWebserverLogs)
	PutWorkerLogs(value *MwaaEnvironmentLoggingConfigurationWorkerLogs)
	ResetDagProcessingLogs()
	ResetSchedulerLogs()
	ResetTaskLogs()
	ResetWebserverLogs()
	ResetWorkerLogs()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) DagProcessingLogs() MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference
	_jsii_.Get(
		j,
		"dagProcessingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) DagProcessingLogsInput() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs {
	var returns *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	_jsii_.Get(
		j,
		"dagProcessingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InternalValue() *MwaaEnvironmentLoggingConfiguration {
	var returns *MwaaEnvironmentLoggingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SchedulerLogs() MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference
	_jsii_.Get(
		j,
		"schedulerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SchedulerLogsInput() *MwaaEnvironmentLoggingConfigurationSchedulerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	_jsii_.Get(
		j,
		"schedulerLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TaskLogs() MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference
	_jsii_.Get(
		j,
		"taskLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TaskLogsInput() *MwaaEnvironmentLoggingConfigurationTaskLogs {
	var returns *MwaaEnvironmentLoggingConfigurationTaskLogs
	_jsii_.Get(
		j,
		"taskLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WebserverLogs() MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference
	_jsii_.Get(
		j,
		"webserverLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WebserverLogsInput() *MwaaEnvironmentLoggingConfigurationWebserverLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWebserverLogs
	_jsii_.Get(
		j,
		"webserverLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WorkerLogs() MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference
	_jsii_.Get(
		j,
		"workerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WorkerLogsInput() *MwaaEnvironmentLoggingConfigurationWorkerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWorkerLogs
	_jsii_.Get(
		j,
		"workerLogsInput",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationOutputReference_Override(m MwaaEnvironmentLoggingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutDagProcessingLogs(value *MwaaEnvironmentLoggingConfigurationDagProcessingLogs) {
	_jsii_.InvokeVoid(
		m,
		"putDagProcessingLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutSchedulerLogs(value *MwaaEnvironmentLoggingConfigurationSchedulerLogs) {
	_jsii_.InvokeVoid(
		m,
		"putSchedulerLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutTaskLogs(value *MwaaEnvironmentLoggingConfigurationTaskLogs) {
	_jsii_.InvokeVoid(
		m,
		"putTaskLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutWebserverLogs(value *MwaaEnvironmentLoggingConfigurationWebserverLogs) {
	_jsii_.InvokeVoid(
		m,
		"putWebserverLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutWorkerLogs(value *MwaaEnvironmentLoggingConfigurationWorkerLogs) {
	_jsii_.InvokeVoid(
		m,
		"putWorkerLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetDagProcessingLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetDagProcessingLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetSchedulerLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetSchedulerLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetTaskLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetWebserverLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetWebserverLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetWorkerLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkerLogs",
		nil, // no parameters
	)
}

type MwaaEnvironmentLoggingConfigurationSchedulerLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

type MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogGroupArn() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	SetInternalValue(val *MwaaEnvironmentLoggingConfigurationSchedulerLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
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
	ResetLogLevel()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) InternalValue() *MwaaEnvironmentLoggingConfigurationSchedulerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference_Override(m MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfigurationSchedulerLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

type MwaaEnvironmentLoggingConfigurationTaskLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

type MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogGroupArn() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *MwaaEnvironmentLoggingConfigurationTaskLogs
	SetInternalValue(val *MwaaEnvironmentLoggingConfigurationTaskLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
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
	ResetLogLevel()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) InternalValue() *MwaaEnvironmentLoggingConfigurationTaskLogs {
	var returns *MwaaEnvironmentLoggingConfigurationTaskLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationTaskLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationTaskLogsOutputReference_Override(m MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfigurationTaskLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

type MwaaEnvironmentLoggingConfigurationWebserverLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

type MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogGroupArn() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *MwaaEnvironmentLoggingConfigurationWebserverLogs
	SetInternalValue(val *MwaaEnvironmentLoggingConfigurationWebserverLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
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
	ResetLogLevel()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) InternalValue() *MwaaEnvironmentLoggingConfigurationWebserverLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWebserverLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference_Override(m MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfigurationWebserverLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

type MwaaEnvironmentLoggingConfigurationWorkerLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

type MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogGroupArn() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *MwaaEnvironmentLoggingConfigurationWorkerLogs
	SetInternalValue(val *MwaaEnvironmentLoggingConfigurationWorkerLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
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
	ResetLogLevel()
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) InternalValue() *MwaaEnvironmentLoggingConfigurationWorkerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWorkerLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference_Override(m MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetInternalValue(val *MwaaEnvironmentLoggingConfigurationWorkerLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

type MwaaEnvironmentNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#security_group_ids MwaaEnvironment#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#subnet_ids MwaaEnvironment#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type MwaaEnvironmentNetworkConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *MwaaEnvironmentNetworkConfiguration
	SetInternalValue(val *MwaaEnvironmentNetworkConfiguration)
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

// The jsii proxy struct for MwaaEnvironmentNetworkConfigurationOutputReference
type jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) InternalValue() *MwaaEnvironmentNetworkConfiguration {
	var returns *MwaaEnvironmentNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMwaaEnvironmentNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) MwaaEnvironmentNetworkConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentNetworkConfigurationOutputReference_Override(m MwaaEnvironmentNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.mwaa.MwaaEnvironmentNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetInternalValue(val *MwaaEnvironmentNetworkConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
