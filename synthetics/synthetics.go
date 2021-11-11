package synthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/synthetics/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html aws_synthetics_canary}.
type SyntheticsCanary interface {
	cdktf.TerraformResource
	Arn() *string
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	ArtifactS3LocationInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EngineArn() *string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	FailureRetentionPeriodInput() *float64
	Fqn() *string
	FriendlyUniqueId() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
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
	RunConfig() SyntheticsCanaryRunConfigOutputReference
	RunConfigInput() *SyntheticsCanaryRunConfig
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3Version() *string
	SetS3Version(val *string)
	S3VersionInput() *string
	Schedule() SyntheticsCanaryScheduleOutputReference
	ScheduleInput() *SyntheticsCanarySchedule
	SourceLocationArn() *string
	StartCanary() interface{}
	SetStartCanary(val interface{})
	StartCanaryInput() interface{}
	Status() *string
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	SuccessRetentionPeriodInput() *float64
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() SyntheticsCanaryVpcConfigOutputReference
	VpcConfigInput() *SyntheticsCanaryVpcConfig
	ZipFile() *string
	SetZipFile(val *string)
	ZipFileInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutRunConfig(value *SyntheticsCanaryRunConfig)
	PutSchedule(value *SyntheticsCanarySchedule)
	PutVpcConfig(value *SyntheticsCanaryVpcConfig)
	ResetFailureRetentionPeriod()
	ResetOverrideLogicalId()
	ResetRunConfig()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3Version()
	ResetStartCanary()
	ResetSuccessRetentionPeriod()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	ResetZipFile()
	SynthesizeAttributes() *map[string]interface{}
	Timeline(index *string) SyntheticsCanaryTimeline
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SyntheticsCanary
type jsiiProxy_SyntheticsCanary struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsCanary) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) EngineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfig() SyntheticsCanaryRunConfigOutputReference {
	var returns SyntheticsCanaryRunConfigOutputReference
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfigInput() *SyntheticsCanaryRunConfig {
	var returns *SyntheticsCanaryRunConfig
	_jsii_.Get(
		j,
		"runConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3VersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Schedule() SyntheticsCanaryScheduleOutputReference {
	var returns SyntheticsCanaryScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ScheduleInput() *SyntheticsCanarySchedule {
	var returns *SyntheticsCanarySchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SourceLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfig() SyntheticsCanaryVpcConfigOutputReference {
	var returns SyntheticsCanaryVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfigInput() *SyntheticsCanaryVpcConfig {
	var returns *SyntheticsCanaryVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ZipFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ZipFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFileInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html aws_synthetics_canary} Resource.
func NewSyntheticsCanary(scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) SyntheticsCanary {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanary{}

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanary",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html aws_synthetics_canary} Resource.
func NewSyntheticsCanary_Override(s SyntheticsCanary, scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanary",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetArtifactS3Location(val *string) {
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetFailureRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetHandler(val *string) {
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetS3Key(val *string) {
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetS3Version(val *string) {
	_jsii_.Set(
		j,
		"s3Version",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetStartCanary(val interface{}) {
	_jsii_.Set(
		j,
		"startCanary",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetSuccessRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetZipFile(val *string) {
	_jsii_.Set(
		j,
		"zipFile",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SyntheticsCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Synthetics.SyntheticsCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsCanary_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Synthetics.SyntheticsCanary",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SyntheticsCanary) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutRunConfig(value *SyntheticsCanaryRunConfig) {
	_jsii_.InvokeVoid(
		s,
		"putRunConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutSchedule(value *SyntheticsCanarySchedule) {
	_jsii_.InvokeVoid(
		s,
		"putSchedule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutVpcConfig(value *SyntheticsCanaryVpcConfig) {
	_jsii_.InvokeVoid(
		s,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetFailureRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureRetentionPeriod",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SyntheticsCanary) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetRunConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRunConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Key() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Key",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Version() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Version",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetStartCanary() {
	_jsii_.InvokeVoid(
		s,
		"resetStartCanary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetSuccessRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessRetentionPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetZipFile() {
	_jsii_.InvokeVoid(
		s,
		"resetZipFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) Timeline(index *string) SyntheticsCanaryTimeline {
	var returns SyntheticsCanaryTimeline

	_jsii_.Invoke(
		s,
		"timeline",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SyntheticsCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_SyntheticsCanary) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SyntheticsCanaryConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#artifact_s3_location SyntheticsCanary#artifact_s3_location}.
	ArtifactS3Location *string `json:"artifactS3Location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#execution_role_arn SyntheticsCanary#execution_role_arn}.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#handler SyntheticsCanary#handler}.
	Handler *string `json:"handler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#name SyntheticsCanary#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#runtime_version SyntheticsCanary#runtime_version}.
	RuntimeVersion *string `json:"runtimeVersion"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#schedule SyntheticsCanary#schedule}
	Schedule *SyntheticsCanarySchedule `json:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#failure_retention_period SyntheticsCanary#failure_retention_period}.
	FailureRetentionPeriod *float64 `json:"failureRetentionPeriod"`
	// run_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#run_config SyntheticsCanary#run_config}
	RunConfig *SyntheticsCanaryRunConfig `json:"runConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#s3_bucket SyntheticsCanary#s3_bucket}.
	S3Bucket *string `json:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#s3_key SyntheticsCanary#s3_key}.
	S3Key *string `json:"s3Key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#s3_version SyntheticsCanary#s3_version}.
	S3Version *string `json:"s3Version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#start_canary SyntheticsCanary#start_canary}.
	StartCanary interface{} `json:"startCanary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#success_retention_period SyntheticsCanary#success_retention_period}.
	SuccessRetentionPeriod *float64 `json:"successRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#tags SyntheticsCanary#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#tags_all SyntheticsCanary#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#vpc_config SyntheticsCanary#vpc_config}
	VpcConfig *SyntheticsCanaryVpcConfig `json:"vpcConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#zip_file SyntheticsCanary#zip_file}.
	ZipFile *string `json:"zipFile"`
}

type SyntheticsCanaryRunConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#active_tracing SyntheticsCanary#active_tracing}.
	ActiveTracing interface{} `json:"activeTracing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#memory_in_mb SyntheticsCanary#memory_in_mb}.
	MemoryInMb *float64 `json:"memoryInMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#timeout_in_seconds SyntheticsCanary#timeout_in_seconds}.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

type SyntheticsCanaryRunConfigOutputReference interface {
	cdktf.ComplexObject
	ActiveTracing() interface{}
	SetActiveTracing(val interface{})
	ActiveTracingInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MemoryInMb() *float64
	SetMemoryInMb(val *float64)
	MemoryInMbInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetActiveTracing()
	ResetMemoryInMb()
	ResetTimeoutInSeconds()
}

// The jsii proxy struct for SyntheticsCanaryRunConfigOutputReference
type jsiiProxy_SyntheticsCanaryRunConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) ActiveTracing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) ActiveTracingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTracingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) MemoryInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) MemoryInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryRunConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryRunConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryRunConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryRunConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryRunConfigOutputReference_Override(s SyntheticsCanaryRunConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryRunConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetActiveTracing(val interface{}) {
	_jsii_.Set(
		j,
		"activeTracing",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetMemoryInMb(val *float64) {
	_jsii_.Set(
		j,
		"memoryInMb",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) ResetActiveTracing() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveTracing",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) ResetMemoryInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMemoryInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

type SyntheticsCanarySchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#expression SyntheticsCanary#expression}.
	Expression *string `json:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#duration_in_seconds SyntheticsCanary#duration_in_seconds}.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

type SyntheticsCanaryScheduleOutputReference interface {
	cdktf.ComplexObject
	DurationInSeconds() *float64
	SetDurationInSeconds(val *float64)
	DurationInSecondsInput() *float64
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
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
	ResetDurationInSeconds()
}

// The jsii proxy struct for SyntheticsCanaryScheduleOutputReference
type jsiiProxy_SyntheticsCanaryScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) DurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) DurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryScheduleOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryScheduleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryScheduleOutputReference_Override(s SyntheticsCanaryScheduleOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetDurationInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"durationInSeconds",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetExpression(val *string) {
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) ResetDurationInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetDurationInSeconds",
		nil, // no parameters
	)
}

type SyntheticsCanaryTimeline interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Created() *string
	LastModified() *string
	LastStarted() *string
	LastStopped() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SyntheticsCanaryTimeline
type jsiiProxy_SyntheticsCanaryTimeline struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) LastStarted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastStarted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) LastStopped() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastStopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewSyntheticsCanaryTimeline(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) SyntheticsCanaryTimeline {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryTimeline{}

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryTimeline",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryTimeline_Override(s SyntheticsCanaryTimeline, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryTimeline",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SyntheticsCanaryVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#security_group_ids SyntheticsCanary#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary.html#subnet_ids SyntheticsCanary#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
}

type SyntheticsCanaryVpcConfigOutputReference interface {
	cdktf.ComplexObject
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for SyntheticsCanaryVpcConfigOutputReference
type jsiiProxy_SyntheticsCanaryVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryVpcConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryVpcConfigOutputReference_Override(s SyntheticsCanaryVpcConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Synthetics.SyntheticsCanaryVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSubnetIds",
		nil, // no parameters
	)
}
