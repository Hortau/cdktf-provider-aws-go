package macie2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/macie2/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_account.html aws_macie2_account}.
type Macie2Account interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedAt() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FindingPublishingFrequency() *string
	SetFindingPublishingFrequency(val *string)
	FindingPublishingFrequencyInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRole() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UpdatedAt() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetFindingPublishingFrequency()
	ResetOverrideLogicalId()
	ResetStatus()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Macie2Account
type jsiiProxy_Macie2Account struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2Account) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) FindingPublishingFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"findingPublishingFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) FindingPublishingFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"findingPublishingFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Account) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_account.html aws_macie2_account} Resource.
func NewMacie2Account(scope constructs.Construct, id *string, config *Macie2AccountConfig) Macie2Account {
	_init_.Initialize()

	j := jsiiProxy_Macie2Account{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2Account",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_account.html aws_macie2_account} Resource.
func NewMacie2Account_Override(m Macie2Account, scope constructs.Construct, id *string, config *Macie2AccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2Account",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2Account) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2Account) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2Account) SetFindingPublishingFrequency(val *string) {
	_jsii_.Set(
		j,
		"findingPublishingFrequency",
		val,
	)
}

func (j *jsiiProxy_Macie2Account) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2Account) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2Account) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Macie2Account_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2Account",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2Account_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2Account",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2Account) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2Account) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2Account) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2Account) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2Account) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2Account) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2Account) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2Account) ResetFindingPublishingFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetFindingPublishingFrequency",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2Account) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Account) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Account) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2Account) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2Account) ToString() *string {
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
func (m *jsiiProxy_Macie2Account) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2AccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_account.html#finding_publishing_frequency Macie2Account#finding_publishing_frequency}.
	FindingPublishingFrequency *string `json:"findingPublishingFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_account.html#status Macie2Account#status}.
	Status *string `json:"status"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html aws_macie2_classification_job}.
type Macie2ClassificationJob interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedAt() *string
	CustomDataIdentifierIds() *[]*string
	SetCustomDataIdentifierIds(val *[]*string)
	CustomDataIdentifierIdsInput() *[]*string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InitialRun() interface{}
	SetInitialRun(val interface{})
	InitialRunInput() interface{}
	JobArn() *string
	JobId() *string
	JobStatus() *string
	SetJobStatus(val *string)
	JobStatusInput() *string
	JobType() *string
	SetJobType(val *string)
	JobTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3JobDefinition() Macie2ClassificationJobS3JobDefinitionOutputReference
	S3JobDefinitionInput() *Macie2ClassificationJobS3JobDefinition
	SamplingPercentage() *float64
	SetSamplingPercentage(val *float64)
	SamplingPercentageInput() *float64
	ScheduleFrequency() Macie2ClassificationJobScheduleFrequencyOutputReference
	ScheduleFrequencyInput() *Macie2ClassificationJobScheduleFrequency
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
	PutS3JobDefinition(value *Macie2ClassificationJobS3JobDefinition)
	PutScheduleFrequency(value *Macie2ClassificationJobScheduleFrequency)
	ResetCustomDataIdentifierIds()
	ResetDescription()
	ResetInitialRun()
	ResetJobStatus()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetSamplingPercentage()
	ResetScheduleFrequency()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	UserPausedDetails(index *string) Macie2ClassificationJobUserPausedDetails
}

// The jsii proxy struct for Macie2ClassificationJob
type jsiiProxy_Macie2ClassificationJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2ClassificationJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CustomDataIdentifierIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDataIdentifierIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CustomDataIdentifierIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDataIdentifierIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) InitialRun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) InitialRunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) S3JobDefinition() Macie2ClassificationJobS3JobDefinitionOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionOutputReference
	_jsii_.Get(
		j,
		"s3JobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) S3JobDefinitionInput() *Macie2ClassificationJobS3JobDefinition {
	var returns *Macie2ClassificationJobS3JobDefinition
	_jsii_.Get(
		j,
		"s3JobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) SamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) SamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ScheduleFrequency() Macie2ClassificationJobScheduleFrequencyOutputReference {
	var returns Macie2ClassificationJobScheduleFrequencyOutputReference
	_jsii_.Get(
		j,
		"scheduleFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ScheduleFrequencyInput() *Macie2ClassificationJobScheduleFrequency {
	var returns *Macie2ClassificationJobScheduleFrequency
	_jsii_.Get(
		j,
		"scheduleFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html aws_macie2_classification_job} Resource.
func NewMacie2ClassificationJob(scope constructs.Construct, id *string, config *Macie2ClassificationJobConfig) Macie2ClassificationJob {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJob{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html aws_macie2_classification_job} Resource.
func NewMacie2ClassificationJob_Override(m Macie2ClassificationJob, scope constructs.Construct, id *string, config *Macie2ClassificationJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJob",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetCustomDataIdentifierIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customDataIdentifierIds",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetInitialRun(val interface{}) {
	_jsii_.Set(
		j,
		"initialRun",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetJobStatus(val *string) {
	_jsii_.Set(
		j,
		"jobStatus",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetJobType(val *string) {
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetSamplingPercentage(val *float64) {
	_jsii_.Set(
		j,
		"samplingPercentage",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob) SetTagsAll(val interface{}) {
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
func Macie2ClassificationJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2ClassificationJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2ClassificationJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2ClassificationJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJob) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJob) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2ClassificationJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) PutS3JobDefinition(value *Macie2ClassificationJobS3JobDefinition) {
	_jsii_.InvokeVoid(
		m,
		"putS3JobDefinition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) PutScheduleFrequency(value *Macie2ClassificationJobScheduleFrequency) {
	_jsii_.InvokeVoid(
		m,
		"putScheduleFrequency",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetCustomDataIdentifierIds() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomDataIdentifierIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetInitialRun() {
	_jsii_.InvokeVoid(
		m,
		"resetInitialRun",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetJobStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetJobStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2ClassificationJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetSamplingPercentage() {
	_jsii_.InvokeVoid(
		m,
		"resetSamplingPercentage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetScheduleFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduleFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2ClassificationJob) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2ClassificationJob) ToString() *string {
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
func (m *jsiiProxy_Macie2ClassificationJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) UserPausedDetails(index *string) Macie2ClassificationJobUserPausedDetails {
	var returns Macie2ClassificationJobUserPausedDetails

	_jsii_.Invoke(
		m,
		"userPausedDetails",
		[]interface{}{index},
		&returns,
	)

	return returns
}

type Macie2ClassificationJobConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#job_type Macie2ClassificationJob#job_type}.
	JobType *string `json:"jobType"`
	// s3_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#s3_job_definition Macie2ClassificationJob#s3_job_definition}
	S3JobDefinition *Macie2ClassificationJobS3JobDefinition `json:"s3JobDefinition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#custom_data_identifier_ids Macie2ClassificationJob#custom_data_identifier_ids}.
	CustomDataIdentifierIds *[]*string `json:"customDataIdentifierIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#description Macie2ClassificationJob#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#initial_run Macie2ClassificationJob#initial_run}.
	InitialRun interface{} `json:"initialRun"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#job_status Macie2ClassificationJob#job_status}.
	JobStatus *string `json:"jobStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#name Macie2ClassificationJob#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#name_prefix Macie2ClassificationJob#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#sampling_percentage Macie2ClassificationJob#sampling_percentage}.
	SamplingPercentage *float64 `json:"samplingPercentage"`
	// schedule_frequency block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#schedule_frequency Macie2ClassificationJob#schedule_frequency}
	ScheduleFrequency *Macie2ClassificationJobScheduleFrequency `json:"scheduleFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tags Macie2ClassificationJob#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tags_all Macie2ClassificationJob#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type Macie2ClassificationJobS3JobDefinition struct {
	// bucket_definitions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#bucket_definitions Macie2ClassificationJob#bucket_definitions}
	BucketDefinitions *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions `json:"bucketDefinitions"`
	// scoping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#scoping Macie2ClassificationJob#scoping}
	Scoping *Macie2ClassificationJobS3JobDefinitionScoping `json:"scoping"`
}

type Macie2ClassificationJobS3JobDefinitionBucketDefinitions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#account_id Macie2ClassificationJob#account_id}.
	AccountId *string `json:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#buckets Macie2ClassificationJob#buckets}.
	Buckets *[]*string `json:"buckets"`
}

type Macie2ClassificationJobS3JobDefinitionOutputReference interface {
	cdktf.ComplexObject
	BucketDefinitions() *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions
	SetBucketDefinitions(val *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions)
	BucketDefinitionsInput() *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Scoping() Macie2ClassificationJobS3JobDefinitionScopingOutputReference
	ScopingInput() *Macie2ClassificationJobS3JobDefinitionScoping
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
	PutScoping(value *Macie2ClassificationJobS3JobDefinitionScoping)
	ResetBucketDefinitions()
	ResetScoping()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketDefinitions() *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions
	_jsii_.Get(
		j,
		"bucketDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketDefinitionsInput() *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions
	_jsii_.Get(
		j,
		"bucketDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) Scoping() Macie2ClassificationJobS3JobDefinitionScopingOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingOutputReference
	_jsii_.Get(
		j,
		"scoping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ScopingInput() *Macie2ClassificationJobS3JobDefinitionScoping {
	var returns *Macie2ClassificationJobS3JobDefinitionScoping
	_jsii_.Get(
		j,
		"scopingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) SetBucketDefinitions(val *[]*Macie2ClassificationJobS3JobDefinitionBucketDefinitions) {
	_jsii_.Set(
		j,
		"bucketDefinitions",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) PutScoping(value *Macie2ClassificationJobS3JobDefinitionScoping) {
	_jsii_.InvokeVoid(
		m,
		"putScoping",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ResetBucketDefinitions() {
	_jsii_.InvokeVoid(
		m,
		"resetBucketDefinitions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ResetScoping() {
	_jsii_.InvokeVoid(
		m,
		"resetScoping",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScoping struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#excludes Macie2ClassificationJob#excludes}
	Excludes *Macie2ClassificationJobS3JobDefinitionScopingExcludes `json:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#includes Macie2ClassificationJob#includes}
	Includes *Macie2ClassificationJobS3JobDefinitionScopingIncludes `json:"includes"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludes struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#and Macie2ClassificationJob#and}
	And *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd `json:"and"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#simple_scope_term Macie2ClassificationJob#simple_scope_term}
	SimpleScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm `json:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tag_scope_term Macie2ClassificationJob#tag_scope_term}
	TagScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm `json:"tagScopeTerm"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `json:"comparator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#values Macie2ClassificationJob#values}.
	Values *[]*string `json:"values"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetComparator()
	ResetKey()
	ResetValues()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetComparator(val *string) {
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		m,
		"resetComparator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		m,
		"resetValues",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `json:"comparator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tag_values Macie2ClassificationJob#tag_values}
	TagValues *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues `json:"tagValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#target Macie2ClassificationJob#target}.
	Target *string `json:"target"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TagValues() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues
	SetTagValues(val *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues)
	TagValuesInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetComparator()
	ResetKey()
	ResetTagValues()
	ResetTarget()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) TagValues() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) TagValuesInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues
	_jsii_.Get(
		j,
		"tagValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetComparator(val *string) {
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetTagValues(val *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues) {
	_jsii_.Set(
		j,
		"tagValues",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		m,
		"resetComparator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) ResetTagValues() {
	_jsii_.InvokeVoid(
		m,
		"resetTagValues",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetTarget",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValues struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#value Macie2ClassificationJob#value}.
	Value *string `json:"value"`
}

type Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference interface {
	cdktf.ComplexObject
	And() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd
	SetAnd(val *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd)
	AndInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd
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
	ResetAnd()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) And() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) AndInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) SetAnd(val *[]*Macie2ClassificationJobS3JobDefinitionScopingExcludesAnd) {
	_jsii_.Set(
		j,
		"and",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		m,
		"resetAnd",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludes struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#and Macie2ClassificationJob#and}
	And *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd `json:"and"`
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#simple_scope_term Macie2ClassificationJob#simple_scope_term}
	SimpleScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTerm `json:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tag_scope_term Macie2ClassificationJob#tag_scope_term}
	TagScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm `json:"tagScopeTerm"`
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTerm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `json:"comparator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#values Macie2ClassificationJob#values}.
	Values *[]*string `json:"values"`
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetComparator()
	ResetKey()
	ResetValues()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetComparator(val *string) {
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		m,
		"resetComparator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTermOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		m,
		"resetValues",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `json:"comparator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#tag_values Macie2ClassificationJob#tag_values}
	TagValues *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues `json:"tagValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#target Macie2ClassificationJob#target}.
	Target *string `json:"target"`
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TagValues() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues
	SetTagValues(val *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues)
	TagValuesInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetComparator()
	ResetKey()
	ResetTagValues()
	ResetTarget()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TagValues() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TagValuesInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues
	_jsii_.Get(
		j,
		"tagValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetComparator(val *string) {
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetTagValues(val *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues) {
	_jsii_.Set(
		j,
		"tagValues",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		m,
		"resetComparator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetTagValues() {
	_jsii_.InvokeVoid(
		m,
		"resetTagValues",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetTarget",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValues struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#key Macie2ClassificationJob#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#value Macie2ClassificationJob#value}.
	Value *string `json:"value"`
}

type Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference interface {
	cdktf.ComplexObject
	And() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd
	SetAnd(val *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd)
	AndInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd
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
	ResetAnd()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) And() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) AndInput() *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd {
	var returns *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) SetAnd(val *[]*Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd) {
	_jsii_.Set(
		j,
		"and",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		m,
		"resetAnd",
		nil, // no parameters
	)
}

type Macie2ClassificationJobS3JobDefinitionScopingOutputReference interface {
	cdktf.ComplexObject
	Excludes() Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference
	ExcludesInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludes
	Includes() Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference
	IncludesInput() *Macie2ClassificationJobS3JobDefinitionScopingIncludes
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
	PutExcludes(value *Macie2ClassificationJobS3JobDefinitionScopingExcludes)
	PutIncludes(value *Macie2ClassificationJobS3JobDefinitionScopingIncludes)
	ResetExcludes()
	ResetIncludes()
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) Excludes() Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingExcludesOutputReference
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) ExcludesInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludes {
	var returns *Macie2ClassificationJobS3JobDefinitionScopingExcludes
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) Includes() Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingIncludesOutputReference
	_jsii_.Get(
		j,
		"includes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) IncludesInput() *Macie2ClassificationJobS3JobDefinitionScopingIncludes {
	var returns *Macie2ClassificationJobS3JobDefinitionScopingIncludes
	_jsii_.Get(
		j,
		"includesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobS3JobDefinitionScopingOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobS3JobDefinitionScopingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobS3JobDefinitionScopingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) PutExcludes(value *Macie2ClassificationJobS3JobDefinitionScopingExcludes) {
	_jsii_.InvokeVoid(
		m,
		"putExcludes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) PutIncludes(value *Macie2ClassificationJobS3JobDefinitionScopingIncludes) {
	_jsii_.InvokeVoid(
		m,
		"putIncludes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) ResetExcludes() {
	_jsii_.InvokeVoid(
		m,
		"resetExcludes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingOutputReference) ResetIncludes() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludes",
		nil, // no parameters
	)
}

type Macie2ClassificationJobScheduleFrequency struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#daily_schedule Macie2ClassificationJob#daily_schedule}.
	DailySchedule interface{} `json:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#monthly_schedule Macie2ClassificationJob#monthly_schedule}.
	MonthlySchedule *float64 `json:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job.html#weekly_schedule Macie2ClassificationJob#weekly_schedule}.
	WeeklySchedule *string `json:"weeklySchedule"`
}

type Macie2ClassificationJobScheduleFrequencyOutputReference interface {
	cdktf.ComplexObject
	DailySchedule() interface{}
	SetDailySchedule(val interface{})
	DailyScheduleInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MonthlySchedule() *float64
	SetMonthlySchedule(val *float64)
	MonthlyScheduleInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WeeklySchedule() *string
	SetWeeklySchedule(val *string)
	WeeklyScheduleInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDailySchedule()
	ResetMonthlySchedule()
	ResetWeeklySchedule()
}

// The jsii proxy struct for Macie2ClassificationJobScheduleFrequencyOutputReference
type jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) DailySchedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) DailyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) MonthlySchedule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) MonthlyScheduleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) WeeklySchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) WeeklyScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}

func NewMacie2ClassificationJobScheduleFrequencyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2ClassificationJobScheduleFrequencyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobScheduleFrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobScheduleFrequencyOutputReference_Override(m Macie2ClassificationJobScheduleFrequencyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobScheduleFrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetDailySchedule(val interface{}) {
	_jsii_.Set(
		j,
		"dailySchedule",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetMonthlySchedule(val *float64) {
	_jsii_.Set(
		j,
		"monthlySchedule",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) SetWeeklySchedule(val *string) {
	_jsii_.Set(
		j,
		"weeklySchedule",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobScheduleFrequencyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

type Macie2ClassificationJobUserPausedDetails interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	JobExpiresAt() *string
	JobImminentExpirationHealthEventArn() *string
	JobPausedAt() *string
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

// The jsii proxy struct for Macie2ClassificationJobUserPausedDetails
type jsiiProxy_Macie2ClassificationJobUserPausedDetails struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) JobExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobExpiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) JobImminentExpirationHealthEventArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobImminentExpirationHealthEventArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) JobPausedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobPausedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewMacie2ClassificationJobUserPausedDetails(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) Macie2ClassificationJobUserPausedDetails {
	_init_.Initialize()

	j := jsiiProxy_Macie2ClassificationJobUserPausedDetails{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobUserPausedDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewMacie2ClassificationJobUserPausedDetails_Override(m Macie2ClassificationJobUserPausedDetails, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2ClassificationJobUserPausedDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobUserPausedDetails) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobUserPausedDetails) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2ClassificationJobUserPausedDetails) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2ClassificationJobUserPausedDetails) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2ClassificationJobUserPausedDetails) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2ClassificationJobUserPausedDetails) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html aws_macie2_custom_data_identifier}.
type Macie2CustomDataIdentifier interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedAt() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IgnoreWords() *[]*string
	SetIgnoreWords(val *[]*string)
	IgnoreWordsInput() *[]*string
	Keywords() *[]*string
	SetKeywords(val *[]*string)
	KeywordsInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumMatchDistance() *float64
	SetMaximumMatchDistance(val *float64)
	MaximumMatchDistanceInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
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
	ResetDescription()
	ResetIgnoreWords()
	ResetKeywords()
	ResetMaximumMatchDistance()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetRegex()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Macie2CustomDataIdentifier
type jsiiProxy_Macie2CustomDataIdentifier struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) IgnoreWords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreWords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) IgnoreWordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreWordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Keywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) KeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) MaximumMatchDistance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumMatchDistance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) MaximumMatchDistanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumMatchDistanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html aws_macie2_custom_data_identifier} Resource.
func NewMacie2CustomDataIdentifier(scope constructs.Construct, id *string, config *Macie2CustomDataIdentifierConfig) Macie2CustomDataIdentifier {
	_init_.Initialize()

	j := jsiiProxy_Macie2CustomDataIdentifier{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2CustomDataIdentifier",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html aws_macie2_custom_data_identifier} Resource.
func NewMacie2CustomDataIdentifier_Override(m Macie2CustomDataIdentifier, scope constructs.Construct, id *string, config *Macie2CustomDataIdentifierConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2CustomDataIdentifier",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetIgnoreWords(val *[]*string) {
	_jsii_.Set(
		j,
		"ignoreWords",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetKeywords(val *[]*string) {
	_jsii_.Set(
		j,
		"keywords",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetMaximumMatchDistance(val *float64) {
	_jsii_.Set(
		j,
		"maximumMatchDistance",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetRegex(val *string) {
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Macie2CustomDataIdentifier) SetTagsAll(val interface{}) {
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
func Macie2CustomDataIdentifier_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2CustomDataIdentifier",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2CustomDataIdentifier_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2CustomDataIdentifier",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2CustomDataIdentifier) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2CustomDataIdentifier) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2CustomDataIdentifier) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetIgnoreWords() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreWords",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetKeywords() {
	_jsii_.InvokeVoid(
		m,
		"resetKeywords",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetMaximumMatchDistance() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximumMatchDistance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetRegex() {
	_jsii_.InvokeVoid(
		m,
		"resetRegex",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2CustomDataIdentifier) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) ToString() *string {
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
func (m *jsiiProxy_Macie2CustomDataIdentifier) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2CustomDataIdentifierConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#description Macie2CustomDataIdentifier#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#ignore_words Macie2CustomDataIdentifier#ignore_words}.
	IgnoreWords *[]*string `json:"ignoreWords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#keywords Macie2CustomDataIdentifier#keywords}.
	Keywords *[]*string `json:"keywords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#maximum_match_distance Macie2CustomDataIdentifier#maximum_match_distance}.
	MaximumMatchDistance *float64 `json:"maximumMatchDistance"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#name Macie2CustomDataIdentifier#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#name_prefix Macie2CustomDataIdentifier#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#regex Macie2CustomDataIdentifier#regex}.
	Regex *string `json:"regex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#tags Macie2CustomDataIdentifier#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_custom_data_identifier.html#tags_all Macie2CustomDataIdentifier#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html aws_macie2_findings_filter}.
type Macie2FindingsFilter interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FindingCriteria() Macie2FindingsFilterFindingCriteriaOutputReference
	FindingCriteriaInput() *Macie2FindingsFilterFindingCriteria
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Position() *float64
	SetPosition(val *float64)
	PositionInput() *float64
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
	PutFindingCriteria(value *Macie2FindingsFilterFindingCriteria)
	ResetDescription()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetPosition()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Macie2FindingsFilter
type jsiiProxy_Macie2FindingsFilter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2FindingsFilter) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) FindingCriteria() Macie2FindingsFilterFindingCriteriaOutputReference {
	var returns Macie2FindingsFilterFindingCriteriaOutputReference
	_jsii_.Get(
		j,
		"findingCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) FindingCriteriaInput() *Macie2FindingsFilterFindingCriteria {
	var returns *Macie2FindingsFilterFindingCriteria
	_jsii_.Get(
		j,
		"findingCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) PositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html aws_macie2_findings_filter} Resource.
func NewMacie2FindingsFilter(scope constructs.Construct, id *string, config *Macie2FindingsFilterConfig) Macie2FindingsFilter {
	_init_.Initialize()

	j := jsiiProxy_Macie2FindingsFilter{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2FindingsFilter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html aws_macie2_findings_filter} Resource.
func NewMacie2FindingsFilter_Override(m Macie2FindingsFilter, scope constructs.Construct, id *string, config *Macie2FindingsFilterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2FindingsFilter",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetPosition(val *float64) {
	_jsii_.Set(
		j,
		"position",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilter) SetTagsAll(val interface{}) {
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
func Macie2FindingsFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2FindingsFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2FindingsFilter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2FindingsFilter",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2FindingsFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2FindingsFilter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2FindingsFilter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2FindingsFilter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2FindingsFilter) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2FindingsFilter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2FindingsFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) PutFindingCriteria(value *Macie2FindingsFilterFindingCriteria) {
	_jsii_.InvokeVoid(
		m,
		"putFindingCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2FindingsFilter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetPosition() {
	_jsii_.InvokeVoid(
		m,
		"resetPosition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilter) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2FindingsFilter) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2FindingsFilter) ToString() *string {
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
func (m *jsiiProxy_Macie2FindingsFilter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2FindingsFilterConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#action Macie2FindingsFilter#action}.
	Action *string `json:"action"`
	// finding_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#finding_criteria Macie2FindingsFilter#finding_criteria}
	FindingCriteria *Macie2FindingsFilterFindingCriteria `json:"findingCriteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#description Macie2FindingsFilter#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#name Macie2FindingsFilter#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#name_prefix Macie2FindingsFilter#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#position Macie2FindingsFilter#position}.
	Position *float64 `json:"position"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#tags Macie2FindingsFilter#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#tags_all Macie2FindingsFilter#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type Macie2FindingsFilterFindingCriteria struct {
	// criterion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#criterion Macie2FindingsFilter#criterion}
	Criterion *[]*Macie2FindingsFilterFindingCriteriaCriterion `json:"criterion"`
}

type Macie2FindingsFilterFindingCriteriaCriterion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#field Macie2FindingsFilter#field}.
	Field *string `json:"field"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#eq Macie2FindingsFilter#eq}.
	Eq *[]*string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#eq_exact_match Macie2FindingsFilter#eq_exact_match}.
	EqExactMatch *[]*string `json:"eqExactMatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#gt Macie2FindingsFilter#gt}.
	Gt *string `json:"gt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#gte Macie2FindingsFilter#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#lt Macie2FindingsFilter#lt}.
	Lt *string `json:"lt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#lte Macie2FindingsFilter#lte}.
	Lte *string `json:"lte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_findings_filter.html#neq Macie2FindingsFilter#neq}.
	Neq *[]*string `json:"neq"`
}

type Macie2FindingsFilterFindingCriteriaOutputReference interface {
	cdktf.ComplexObject
	Criterion() *[]*Macie2FindingsFilterFindingCriteriaCriterion
	SetCriterion(val *[]*Macie2FindingsFilterFindingCriteriaCriterion)
	CriterionInput() *[]*Macie2FindingsFilterFindingCriteriaCriterion
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
	ResetCriterion()
}

// The jsii proxy struct for Macie2FindingsFilterFindingCriteriaOutputReference
type jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) Criterion() *[]*Macie2FindingsFilterFindingCriteriaCriterion {
	var returns *[]*Macie2FindingsFilterFindingCriteriaCriterion
	_jsii_.Get(
		j,
		"criterion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) CriterionInput() *[]*Macie2FindingsFilterFindingCriteriaCriterion {
	var returns *[]*Macie2FindingsFilterFindingCriteriaCriterion
	_jsii_.Get(
		j,
		"criterionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2FindingsFilterFindingCriteriaOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2FindingsFilterFindingCriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2FindingsFilterFindingCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2FindingsFilterFindingCriteriaOutputReference_Override(m Macie2FindingsFilterFindingCriteriaOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2FindingsFilterFindingCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) SetCriterion(val *[]*Macie2FindingsFilterFindingCriteriaCriterion) {
	_jsii_.Set(
		j,
		"criterion",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaOutputReference) ResetCriterion() {
	_jsii_.InvokeVoid(
		m,
		"resetCriterion",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html aws_macie2_invitation_accepter}.
type Macie2InvitationAccepter interface {
	cdktf.TerraformResource
	AdministratorAccountId() *string
	SetAdministratorAccountId(val *string)
	AdministratorAccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InvitationId() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() Macie2InvitationAccepterTimeoutsOutputReference
	TimeoutsInput() *Macie2InvitationAccepterTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *Macie2InvitationAccepterTimeouts)
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Macie2InvitationAccepter
type jsiiProxy_Macie2InvitationAccepter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2InvitationAccepter) AdministratorAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) AdministratorAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) InvitationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) Timeouts() Macie2InvitationAccepterTimeoutsOutputReference {
	var returns Macie2InvitationAccepterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepter) TimeoutsInput() *Macie2InvitationAccepterTimeouts {
	var returns *Macie2InvitationAccepterTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html aws_macie2_invitation_accepter} Resource.
func NewMacie2InvitationAccepter(scope constructs.Construct, id *string, config *Macie2InvitationAccepterConfig) Macie2InvitationAccepter {
	_init_.Initialize()

	j := jsiiProxy_Macie2InvitationAccepter{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2InvitationAccepter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html aws_macie2_invitation_accepter} Resource.
func NewMacie2InvitationAccepter_Override(m Macie2InvitationAccepter, scope constructs.Construct, id *string, config *Macie2InvitationAccepterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2InvitationAccepter",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepter) SetAdministratorAccountId(val *string) {
	_jsii_.Set(
		j,
		"administratorAccountId",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepter) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepter) SetProvider(val cdktf.TerraformProvider) {
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
func Macie2InvitationAccepter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2InvitationAccepter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2InvitationAccepter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2InvitationAccepter",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2InvitationAccepter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2InvitationAccepter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2InvitationAccepter) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2InvitationAccepter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2InvitationAccepter) PutTimeouts(value *Macie2InvitationAccepterTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2InvitationAccepter) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2InvitationAccepter) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2InvitationAccepter) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2InvitationAccepter) ToString() *string {
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
func (m *jsiiProxy_Macie2InvitationAccepter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2InvitationAccepterConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html#administrator_account_id Macie2InvitationAccepter#administrator_account_id}.
	AdministratorAccountId *string `json:"administratorAccountId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html#timeouts Macie2InvitationAccepter#timeouts}
	Timeouts *Macie2InvitationAccepterTimeouts `json:"timeouts"`
}

type Macie2InvitationAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_invitation_accepter.html#create Macie2InvitationAccepter#create}.
	Create *string `json:"create"`
}

type Macie2InvitationAccepterTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
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
	ResetCreate()
}

// The jsii proxy struct for Macie2InvitationAccepterTimeoutsOutputReference
type jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacie2InvitationAccepterTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2InvitationAccepterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2InvitationAccepterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2InvitationAccepterTimeoutsOutputReference_Override(m Macie2InvitationAccepterTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2InvitationAccepterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2InvitationAccepterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html aws_macie2_member}.
type Macie2Member interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AdministratorAccountId() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InvitationDisableEmailNotification() *string
	SetInvitationDisableEmailNotification(val *string)
	InvitationDisableEmailNotificationInput() *string
	InvitationMessage() *string
	SetInvitationMessage(val *string)
	InvitationMessageInput() *string
	Invite() interface{}
	SetInvite(val interface{})
	InvitedAt() *string
	InviteInput() interface{}
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterAccountId() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RelationshipStatus() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() Macie2MemberTimeoutsOutputReference
	TimeoutsInput() *Macie2MemberTimeouts
	UpdatedAt() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *Macie2MemberTimeouts)
	ResetInvitationDisableEmailNotification()
	ResetInvitationMessage()
	ResetInvite()
	ResetOverrideLogicalId()
	ResetStatus()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for Macie2Member
type jsiiProxy_Macie2Member struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2Member) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) AdministratorAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InvitationDisableEmailNotification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationDisableEmailNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InvitationDisableEmailNotificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationDisableEmailNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InvitationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InvitationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Invite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InvitedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) InviteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inviteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) MasterAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) RelationshipStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationshipStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) Timeouts() Macie2MemberTimeoutsOutputReference {
	var returns Macie2MemberTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) TimeoutsInput() *Macie2MemberTimeouts {
	var returns *Macie2MemberTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2Member) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html aws_macie2_member} Resource.
func NewMacie2Member(scope constructs.Construct, id *string, config *Macie2MemberConfig) Macie2Member {
	_init_.Initialize()

	j := jsiiProxy_Macie2Member{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2Member",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html aws_macie2_member} Resource.
func NewMacie2Member_Override(m Macie2Member, scope constructs.Construct, id *string, config *Macie2MemberConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2Member",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2Member) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetInvitationDisableEmailNotification(val *string) {
	_jsii_.Set(
		j,
		"invitationDisableEmailNotification",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetInvitationMessage(val *string) {
	_jsii_.Set(
		j,
		"invitationMessage",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetInvite(val interface{}) {
	_jsii_.Set(
		j,
		"invite",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Macie2Member) SetTagsAll(val interface{}) {
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
func Macie2Member_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2Member",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2Member_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2Member",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2Member) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2Member) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2Member) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2Member) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2Member) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2Member) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2Member) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2Member) PutTimeouts(value *Macie2MemberTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2Member) ResetInvitationDisableEmailNotification() {
	_jsii_.InvokeVoid(
		m,
		"resetInvitationDisableEmailNotification",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetInvitationMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetInvitationMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetInvite() {
	_jsii_.InvokeVoid(
		m,
		"resetInvite",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2Member) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2Member) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2Member) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2Member) ToString() *string {
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
func (m *jsiiProxy_Macie2Member) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2MemberConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#account_id Macie2Member#account_id}.
	AccountId *string `json:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#email Macie2Member#email}.
	Email *string `json:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#invitation_disable_email_notification Macie2Member#invitation_disable_email_notification}.
	InvitationDisableEmailNotification *string `json:"invitationDisableEmailNotification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#invitation_message Macie2Member#invitation_message}.
	InvitationMessage *string `json:"invitationMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#invite Macie2Member#invite}.
	Invite interface{} `json:"invite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#status Macie2Member#status}.
	Status *string `json:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#tags Macie2Member#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#tags_all Macie2Member#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#timeouts Macie2Member#timeouts}
	Timeouts *Macie2MemberTimeouts `json:"timeouts"`
}

type Macie2MemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#create Macie2Member#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_member.html#update Macie2Member#update}.
	Update *string `json:"update"`
}

type Macie2MemberTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
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
	ResetUpdate()
}

// The jsii proxy struct for Macie2MemberTimeoutsOutputReference
type jsiiProxy_Macie2MemberTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewMacie2MemberTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) Macie2MemberTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_Macie2MemberTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2MemberTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacie2MemberTimeoutsOutputReference_Override(m Macie2MemberTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2MemberTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Macie2MemberTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2MemberTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie2_organization_admin_account.html aws_macie2_organization_admin_account}.
type Macie2OrganizationAdminAccount interface {
	cdktf.TerraformResource
	AdminAccountId() *string
	SetAdminAccountId(val *string)
	AdminAccountIdInput() *string
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

// The jsii proxy struct for Macie2OrganizationAdminAccount
type jsiiProxy_Macie2OrganizationAdminAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) AdminAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) AdminAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_organization_admin_account.html aws_macie2_organization_admin_account} Resource.
func NewMacie2OrganizationAdminAccount(scope constructs.Construct, id *string, config *Macie2OrganizationAdminAccountConfig) Macie2OrganizationAdminAccount {
	_init_.Initialize()

	j := jsiiProxy_Macie2OrganizationAdminAccount{}

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2OrganizationAdminAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie2_organization_admin_account.html aws_macie2_organization_admin_account} Resource.
func NewMacie2OrganizationAdminAccount_Override(m Macie2OrganizationAdminAccount, scope constructs.Construct, id *string, config *Macie2OrganizationAdminAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie2.Macie2OrganizationAdminAccount",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) SetAdminAccountId(val *string) {
	_jsii_.Set(
		j,
		"adminAccountId",
		val,
	)
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2OrganizationAdminAccount) SetProvider(val cdktf.TerraformProvider) {
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
func Macie2OrganizationAdminAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie2.Macie2OrganizationAdminAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2OrganizationAdminAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie2.Macie2OrganizationAdminAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Macie2OrganizationAdminAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_Macie2OrganizationAdminAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_Macie2OrganizationAdminAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_Macie2OrganizationAdminAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2OrganizationAdminAccount) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) ToMetadata() interface{} {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) ToString() *string {
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
func (m *jsiiProxy_Macie2OrganizationAdminAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type Macie2OrganizationAdminAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_organization_admin_account.html#admin_account_id Macie2OrganizationAdminAccount#admin_account_id}.
	AdminAccountId *string `json:"adminAccountId"`
}
