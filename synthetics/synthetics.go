package synthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/synthetics/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary aws_synthetics_canary}.
type SyntheticsCanary interface {
	cdktf.TerraformResource
	Arn() *string
	ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference
	ArtifactConfigInput() *SyntheticsCanaryArtifactConfig
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	ArtifactS3LocationInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() SyntheticsCanaryVpcConfigOutputReference
	VpcConfigInput() *SyntheticsCanaryVpcConfig
	ZipFile() *string
	SetZipFile(val *string)
	ZipFileInput() *string
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
	PutArtifactConfig(value *SyntheticsCanaryArtifactConfig)
	PutRunConfig(value *SyntheticsCanaryRunConfig)
	PutSchedule(value *SyntheticsCanarySchedule)
	PutVpcConfig(value *SyntheticsCanaryVpcConfig)
	ResetArtifactConfig()
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

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference {
	var returns SyntheticsCanaryArtifactConfigOutputReference
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfigInput() *SyntheticsCanaryArtifactConfig {
	var returns *SyntheticsCanaryArtifactConfig
	_jsii_.Get(
		j,
		"artifactConfigInput",
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

func (j *jsiiProxy_SyntheticsCanary) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_SyntheticsCanary) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary aws_synthetics_canary} Resource.
func NewSyntheticsCanary(scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) SyntheticsCanary {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanary{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanary",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary aws_synthetics_canary} Resource.
func NewSyntheticsCanary_Override(s SyntheticsCanary, scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanary",
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

func (j *jsiiProxy_SyntheticsCanary) SetCount(val *float64) {
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

func (j *jsiiProxy_SyntheticsCanary) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.synthetics.SyntheticsCanary",
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
		"hashicorp_aws.synthetics.SyntheticsCanary",
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
func (s *jsiiProxy_SyntheticsCanary) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (s *jsiiProxy_SyntheticsCanary) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanary) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanary) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanary) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
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

func (s *jsiiProxy_SyntheticsCanary) PutArtifactConfig(value *SyntheticsCanaryArtifactConfig) {
	_jsii_.InvokeVoid(
		s,
		"putArtifactConfig",
		[]interface{}{value},
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

func (s *jsiiProxy_SyntheticsCanary) ResetArtifactConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetArtifactConfig",
		nil, // no parameters
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

type SyntheticsCanaryArtifactConfig struct {
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#s3_encryption SyntheticsCanary#s3_encryption}
	S3Encryption *SyntheticsCanaryArtifactConfigS3Encryption `json:"s3Encryption" yaml:"s3Encryption"`
}

type SyntheticsCanaryArtifactConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SyntheticsCanaryArtifactConfig
	SetInternalValue(val *SyntheticsCanaryArtifactConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3Encryption() SyntheticsCanaryArtifactConfigS3EncryptionOutputReference
	S3EncryptionInput() *SyntheticsCanaryArtifactConfigS3Encryption
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
	PutS3Encryption(value *SyntheticsCanaryArtifactConfigS3Encryption)
	ResetS3Encryption()
}

// The jsii proxy struct for SyntheticsCanaryArtifactConfigOutputReference
type jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) InternalValue() *SyntheticsCanaryArtifactConfig {
	var returns *SyntheticsCanaryArtifactConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) S3Encryption() SyntheticsCanaryArtifactConfigS3EncryptionOutputReference {
	var returns SyntheticsCanaryArtifactConfigS3EncryptionOutputReference
	_jsii_.Get(
		j,
		"s3Encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) S3EncryptionInput() *SyntheticsCanaryArtifactConfigS3Encryption {
	var returns *SyntheticsCanaryArtifactConfigS3Encryption
	_jsii_.Get(
		j,
		"s3EncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryArtifactConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryArtifactConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryArtifactConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryArtifactConfigOutputReference_Override(s SyntheticsCanaryArtifactConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryArtifactConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) SetInternalValue(val *SyntheticsCanaryArtifactConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) PutS3Encryption(value *SyntheticsCanaryArtifactConfigS3Encryption) {
	_jsii_.InvokeVoid(
		s,
		"putS3Encryption",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanaryArtifactConfigOutputReference) ResetS3Encryption() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Encryption",
		nil, // no parameters
	)
}

type SyntheticsCanaryArtifactConfigS3Encryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#encryption_mode SyntheticsCanary#encryption_mode}.
	EncryptionMode *string `json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#kms_key_arn SyntheticsCanary#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

type SyntheticsCanaryArtifactConfigS3EncryptionOutputReference interface {
	cdktf.ComplexObject
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	EncryptionModeInput() *string
	InternalValue() *SyntheticsCanaryArtifactConfigS3Encryption
	SetInternalValue(val *SyntheticsCanaryArtifactConfigS3Encryption)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	ResetEncryptionMode()
	ResetKmsKeyArn()
}

// The jsii proxy struct for SyntheticsCanaryArtifactConfigS3EncryptionOutputReference
type jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) InternalValue() *SyntheticsCanaryArtifactConfigS3Encryption {
	var returns *SyntheticsCanaryArtifactConfigS3Encryption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryArtifactConfigS3EncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryArtifactConfigS3EncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryArtifactConfigS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryArtifactConfigS3EncryptionOutputReference_Override(s SyntheticsCanaryArtifactConfigS3EncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryArtifactConfigS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetInternalValue(val *SyntheticsCanaryArtifactConfigS3Encryption) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		s,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanaryArtifactConfigS3EncryptionOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

// AWS Synthetics.
type SyntheticsCanaryConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#artifact_s3_location SyntheticsCanary#artifact_s3_location}.
	ArtifactS3Location *string `json:"artifactS3Location" yaml:"artifactS3Location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#execution_role_arn SyntheticsCanary#execution_role_arn}.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#handler SyntheticsCanary#handler}.
	Handler *string `json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#name SyntheticsCanary#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#runtime_version SyntheticsCanary#runtime_version}.
	RuntimeVersion *string `json:"runtimeVersion" yaml:"runtimeVersion"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#schedule SyntheticsCanary#schedule}
	Schedule *SyntheticsCanarySchedule `json:"schedule" yaml:"schedule"`
	// artifact_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#artifact_config SyntheticsCanary#artifact_config}
	ArtifactConfig *SyntheticsCanaryArtifactConfig `json:"artifactConfig" yaml:"artifactConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#failure_retention_period SyntheticsCanary#failure_retention_period}.
	FailureRetentionPeriod *float64 `json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// run_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#run_config SyntheticsCanary#run_config}
	RunConfig *SyntheticsCanaryRunConfig `json:"runConfig" yaml:"runConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#s3_bucket SyntheticsCanary#s3_bucket}.
	S3Bucket *string `json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#s3_key SyntheticsCanary#s3_key}.
	S3Key *string `json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#s3_version SyntheticsCanary#s3_version}.
	S3Version *string `json:"s3Version" yaml:"s3Version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#start_canary SyntheticsCanary#start_canary}.
	StartCanary interface{} `json:"startCanary" yaml:"startCanary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#success_retention_period SyntheticsCanary#success_retention_period}.
	SuccessRetentionPeriod *float64 `json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#tags SyntheticsCanary#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#tags_all SyntheticsCanary#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#vpc_config SyntheticsCanary#vpc_config}
	VpcConfig *SyntheticsCanaryVpcConfig `json:"vpcConfig" yaml:"vpcConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#zip_file SyntheticsCanary#zip_file}.
	ZipFile *string `json:"zipFile" yaml:"zipFile"`
}

type SyntheticsCanaryRunConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#active_tracing SyntheticsCanary#active_tracing}.
	ActiveTracing interface{} `json:"activeTracing" yaml:"activeTracing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#memory_in_mb SyntheticsCanary#memory_in_mb}.
	MemoryInMb *float64 `json:"memoryInMb" yaml:"memoryInMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#timeout_in_seconds SyntheticsCanary#timeout_in_seconds}.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

type SyntheticsCanaryRunConfigOutputReference interface {
	cdktf.ComplexObject
	ActiveTracing() interface{}
	SetActiveTracing(val interface{})
	ActiveTracingInput() interface{}
	InternalValue() *SyntheticsCanaryRunConfig
	SetInternalValue(val *SyntheticsCanaryRunConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MemoryInMb() *float64
	SetMemoryInMb(val *float64)
	MemoryInMbInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
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

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) InternalValue() *SyntheticsCanaryRunConfig {
	var returns *SyntheticsCanaryRunConfig
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
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

func NewSyntheticsCanaryRunConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryRunConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryRunConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryRunConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryRunConfigOutputReference_Override(s SyntheticsCanaryRunConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryRunConfigOutputReference",
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

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetInternalValue(val *SyntheticsCanaryRunConfig) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryRunConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#expression SyntheticsCanary#expression}.
	Expression *string `json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#duration_in_seconds SyntheticsCanary#duration_in_seconds}.
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

type SyntheticsCanaryScheduleOutputReference interface {
	cdktf.ComplexObject
	DurationInSeconds() *float64
	SetDurationInSeconds(val *float64)
	DurationInSecondsInput() *float64
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	InternalValue() *SyntheticsCanarySchedule
	SetInternalValue(val *SyntheticsCanarySchedule)
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

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) InternalValue() *SyntheticsCanarySchedule {
	var returns *SyntheticsCanarySchedule
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryScheduleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryScheduleOutputReference_Override(s SyntheticsCanaryScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryScheduleOutputReference",
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

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetInternalValue(val *SyntheticsCanarySchedule) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_SyntheticsCanaryScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
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

func (j *jsiiProxy_SyntheticsCanaryTimeline) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSyntheticsCanaryTimeline(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SyntheticsCanaryTimeline {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryTimeline{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryTimeline",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryTimeline_Override(s SyntheticsCanaryTimeline, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryTimeline",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
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

func (j *jsiiProxy_SyntheticsCanaryTimeline) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanaryTimeline) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryTimeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#security_group_ids SyntheticsCanary#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#subnet_ids SyntheticsCanary#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type SyntheticsCanaryVpcConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SyntheticsCanaryVpcConfig
	SetInternalValue(val *SyntheticsCanaryVpcConfig)
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
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for SyntheticsCanaryVpcConfigOutputReference
type jsiiProxy_SyntheticsCanaryVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) InternalValue() *SyntheticsCanaryVpcConfig {
	var returns *SyntheticsCanaryVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func NewSyntheticsCanaryVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SyntheticsCanaryVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSyntheticsCanaryVpcConfigOutputReference_Override(s SyntheticsCanaryVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.synthetics.SyntheticsCanaryVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetInternalValue(val *SyntheticsCanaryVpcConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
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

func (j *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
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
func (s *jsiiProxy_SyntheticsCanaryVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
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
