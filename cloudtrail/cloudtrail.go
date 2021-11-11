package cloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/cloudtrail/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html aws_cloudtrail}.
type Cloudtrail interface {
	cdktf.TerraformResource
	AdvancedEventSelector() *[]*CloudtrailAdvancedEventSelector
	SetAdvancedEventSelector(val *[]*CloudtrailAdvancedEventSelector)
	AdvancedEventSelectorInput() *[]*CloudtrailAdvancedEventSelector
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CloudWatchLogsGroupArn() *string
	SetCloudWatchLogsGroupArn(val *string)
	CloudWatchLogsGroupArnInput() *string
	CloudWatchLogsRoleArn() *string
	SetCloudWatchLogsRoleArn(val *string)
	CloudWatchLogsRoleArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	EnableLogFileValidationInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
	EventSelector() *[]*CloudtrailEventSelector
	SetEventSelector(val *[]*CloudtrailEventSelector)
	EventSelectorInput() *[]*CloudtrailEventSelector
	Fqn() *string
	FriendlyUniqueId() *string
	HomeRegion() *string
	Id() *string
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	IncludeGlobalServiceEventsInput() interface{}
	InsightSelector() *[]*CloudtrailInsightSelector
	SetInsightSelector(val *[]*CloudtrailInsightSelector)
	InsightSelectorInput() *[]*CloudtrailInsightSelector
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	IsMultiRegionTrailInput() interface{}
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	IsOrganizationTrailInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KeyPrefixInput() *string
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	SnsTopicNameInput() *string
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
	ResetAdvancedEventSelector()
	ResetCloudWatchLogsGroupArn()
	ResetCloudWatchLogsRoleArn()
	ResetEnableLogFileValidation()
	ResetEnableLogging()
	ResetEventSelector()
	ResetIncludeGlobalServiceEvents()
	ResetInsightSelector()
	ResetIsMultiRegionTrail()
	ResetIsOrganizationTrail()
	ResetKmsKeyId()
	ResetOverrideLogicalId()
	ResetS3KeyPrefix()
	ResetSnsTopicName()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Cloudtrail
type jsiiProxy_Cloudtrail struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Cloudtrail) AdvancedEventSelector() *[]*CloudtrailAdvancedEventSelector {
	var returns *[]*CloudtrailAdvancedEventSelector
	_jsii_.Get(
		j,
		"advancedEventSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) AdvancedEventSelectorInput() *[]*CloudtrailAdvancedEventSelector {
	var returns *[]*CloudtrailAdvancedEventSelector
	_jsii_.Get(
		j,
		"advancedEventSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) CloudWatchLogsGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) CloudWatchLogsGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) CloudWatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) CloudWatchLogsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EnableLogFileValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EventSelector() *[]*CloudtrailEventSelector {
	var returns *[]*CloudtrailEventSelector
	_jsii_.Get(
		j,
		"eventSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) EventSelectorInput() *[]*CloudtrailEventSelector {
	var returns *[]*CloudtrailEventSelector
	_jsii_.Get(
		j,
		"eventSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IncludeGlobalServiceEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) InsightSelector() *[]*CloudtrailInsightSelector {
	var returns *[]*CloudtrailInsightSelector
	_jsii_.Get(
		j,
		"insightSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) InsightSelectorInput() *[]*CloudtrailInsightSelector {
	var returns *[]*CloudtrailInsightSelector
	_jsii_.Get(
		j,
		"insightSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IsMultiRegionTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) IsOrganizationTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) SnsTopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudtrail) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html aws_cloudtrail} Resource.
func NewCloudtrail(scope constructs.Construct, id *string, config *CloudtrailConfig) Cloudtrail {
	_init_.Initialize()

	j := jsiiProxy_Cloudtrail{}

	_jsii_.Create(
		"hashicorp_aws.CloudTrail.Cloudtrail",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html aws_cloudtrail} Resource.
func NewCloudtrail_Override(c Cloudtrail, scope constructs.Construct, id *string, config *CloudtrailConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudTrail.Cloudtrail",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Cloudtrail) SetAdvancedEventSelector(val *[]*CloudtrailAdvancedEventSelector) {
	_jsii_.Set(
		j,
		"advancedEventSelector",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetCloudWatchLogsGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsGroupArn",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetCloudWatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetEnableLogFileValidation(val interface{}) {
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetEnableLogging(val interface{}) {
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetEventSelector(val *[]*CloudtrailEventSelector) {
	_jsii_.Set(
		j,
		"eventSelector",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetIncludeGlobalServiceEvents(val interface{}) {
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetInsightSelector(val *[]*CloudtrailInsightSelector) {
	_jsii_.Set(
		j,
		"insightSelector",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetIsMultiRegionTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetIsOrganizationTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetSnsTopicName(val *string) {
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Cloudtrail) SetTagsAll(val interface{}) {
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
func Cloudtrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudTrail.Cloudtrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Cloudtrail_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudTrail.Cloudtrail",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_Cloudtrail) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_Cloudtrail) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_Cloudtrail) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_Cloudtrail) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_Cloudtrail) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_Cloudtrail) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_Cloudtrail) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Cloudtrail) ResetAdvancedEventSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvancedEventSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetCloudWatchLogsGroupArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudWatchLogsGroupArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetCloudWatchLogsRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudWatchLogsRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetEnableLogFileValidation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLogFileValidation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetEventSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetEventSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetIncludeGlobalServiceEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeGlobalServiceEvents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetInsightSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetInsightSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetIsMultiRegionTrail() {
	_jsii_.InvokeVoid(
		c,
		"resetIsMultiRegionTrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetIsOrganizationTrail() {
	_jsii_.InvokeVoid(
		c,
		"resetIsOrganizationTrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_Cloudtrail) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetSnsTopicName() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudtrail) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_Cloudtrail) ToMetadata() interface{} {
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
func (c *jsiiProxy_Cloudtrail) ToString() *string {
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
func (c *jsiiProxy_Cloudtrail) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudtrailAdvancedEventSelector struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#field_selector Cloudtrail#field_selector}
	FieldSelector *[]*CloudtrailAdvancedEventSelectorFieldSelector `json:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#name Cloudtrail#name}.
	Name *string `json:"name"`
}

type CloudtrailAdvancedEventSelectorFieldSelector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#field Cloudtrail#field}.
	Field *string `json:"field"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#ends_with Cloudtrail#ends_with}.
	EndsWith *[]*string `json:"endsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#equals Cloudtrail#equals}.
	EqualTo *[]*string `json:"equalTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#not_ends_with Cloudtrail#not_ends_with}.
	NotEndsWith *[]*string `json:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#not_equals Cloudtrail#not_equals}.
	NotEquals *[]*string `json:"notEquals"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#not_starts_with Cloudtrail#not_starts_with}.
	NotStartsWith *[]*string `json:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#starts_with Cloudtrail#starts_with}.
	StartsWith *[]*string `json:"startsWith"`
}

type CloudtrailConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#name Cloudtrail#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#s3_bucket_name Cloudtrail#s3_bucket_name}.
	S3BucketName *string `json:"s3BucketName"`
	// advanced_event_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#advanced_event_selector Cloudtrail#advanced_event_selector}
	AdvancedEventSelector *[]*CloudtrailAdvancedEventSelector `json:"advancedEventSelector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#cloud_watch_logs_group_arn Cloudtrail#cloud_watch_logs_group_arn}.
	CloudWatchLogsGroupArn *string `json:"cloudWatchLogsGroupArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#cloud_watch_logs_role_arn Cloudtrail#cloud_watch_logs_role_arn}.
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#enable_log_file_validation Cloudtrail#enable_log_file_validation}.
	EnableLogFileValidation interface{} `json:"enableLogFileValidation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#enable_logging Cloudtrail#enable_logging}.
	EnableLogging interface{} `json:"enableLogging"`
	// event_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#event_selector Cloudtrail#event_selector}
	EventSelector *[]*CloudtrailEventSelector `json:"eventSelector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#include_global_service_events Cloudtrail#include_global_service_events}.
	IncludeGlobalServiceEvents interface{} `json:"includeGlobalServiceEvents"`
	// insight_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#insight_selector Cloudtrail#insight_selector}
	InsightSelector *[]*CloudtrailInsightSelector `json:"insightSelector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#is_multi_region_trail Cloudtrail#is_multi_region_trail}.
	IsMultiRegionTrail interface{} `json:"isMultiRegionTrail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#is_organization_trail Cloudtrail#is_organization_trail}.
	IsOrganizationTrail interface{} `json:"isOrganizationTrail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#kms_key_id Cloudtrail#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#s3_key_prefix Cloudtrail#s3_key_prefix}.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#sns_topic_name Cloudtrail#sns_topic_name}.
	SnsTopicName *string `json:"snsTopicName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#tags Cloudtrail#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#tags_all Cloudtrail#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type CloudtrailEventSelector struct {
	// data_resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#data_resource Cloudtrail#data_resource}
	DataResource *[]*CloudtrailEventSelectorDataResource `json:"dataResource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#include_management_events Cloudtrail#include_management_events}.
	IncludeManagementEvents interface{} `json:"includeManagementEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#read_write_type Cloudtrail#read_write_type}.
	ReadWriteType *string `json:"readWriteType"`
}

type CloudtrailEventSelectorDataResource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#type Cloudtrail#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#values Cloudtrail#values}.
	Values *[]*string `json:"values"`
}

type CloudtrailInsightSelector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail.html#insight_type Cloudtrail#insight_type}.
	InsightType *string `json:"insightType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudtrail_service_account.html aws_cloudtrail_service_account}.
type DataAwsCloudtrailServiceAccount interface {
	cdktf.TerraformDataSource
	Arn() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetRegion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudtrailServiceAccount
type jsiiProxy_DataAwsCloudtrailServiceAccount struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudtrail_service_account.html aws_cloudtrail_service_account} Data Source.
func NewDataAwsCloudtrailServiceAccount(scope constructs.Construct, id *string, config *DataAwsCloudtrailServiceAccountConfig) DataAwsCloudtrailServiceAccount {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudtrailServiceAccount{}

	_jsii_.Create(
		"hashicorp_aws.CloudTrail.DataAwsCloudtrailServiceAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudtrail_service_account.html aws_cloudtrail_service_account} Data Source.
func NewDataAwsCloudtrailServiceAccount_Override(d DataAwsCloudtrailServiceAccount, scope constructs.Construct, id *string, config *DataAwsCloudtrailServiceAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudTrail.DataAwsCloudtrailServiceAccount",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudtrailServiceAccount) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsCloudtrailServiceAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudTrail.DataAwsCloudtrailServiceAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudtrailServiceAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudTrail.DataAwsCloudtrailServiceAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudtrailServiceAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudtrailServiceAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudtrail_service_account.html#region DataAwsCloudtrailServiceAccount#region}.
	Region *string `json:"region"`
}
