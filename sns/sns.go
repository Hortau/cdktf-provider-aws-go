package sns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/sns/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/sns_topic.html aws_sns_topic}.
type DataAwsSnsTopic interface {
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

// The jsii proxy struct for DataAwsSnsTopic
type jsiiProxy_DataAwsSnsTopic struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSnsTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSnsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/sns_topic.html aws_sns_topic} Data Source.
func NewDataAwsSnsTopic(scope constructs.Construct, id *string, config *DataAwsSnsTopicConfig) DataAwsSnsTopic {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSnsTopic{}

	_jsii_.Create(
		"hashicorp_aws.SNS.DataAwsSnsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/sns_topic.html aws_sns_topic} Data Source.
func NewDataAwsSnsTopic_Override(d DataAwsSnsTopic, scope constructs.Construct, id *string, config *DataAwsSnsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.DataAwsSnsTopic",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSnsTopic) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSnsTopic) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSnsTopic) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSnsTopic) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsSnsTopic) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsSnsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.DataAwsSnsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSnsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.DataAwsSnsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSnsTopic) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSnsTopic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSnsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSnsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSnsTopic) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSnsTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSnsTopic) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSnsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSnsTopic) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSnsTopic) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSnsTopic) ToString() *string {
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
func (d *jsiiProxy_DataAwsSnsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsSnsTopicConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/sns_topic.html#name DataAwsSnsTopic#name}.
	Name *string `json:"name"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html aws_sns_platform_application}.
type SnsPlatformApplication interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EventDeliveryFailureTopicArn() *string
	SetEventDeliveryFailureTopicArn(val *string)
	EventDeliveryFailureTopicArnInput() *string
	EventEndpointCreatedTopicArn() *string
	SetEventEndpointCreatedTopicArn(val *string)
	EventEndpointCreatedTopicArnInput() *string
	EventEndpointDeletedTopicArn() *string
	SetEventEndpointDeletedTopicArn(val *string)
	EventEndpointDeletedTopicArnInput() *string
	EventEndpointUpdatedTopicArn() *string
	SetEventEndpointUpdatedTopicArn(val *string)
	EventEndpointUpdatedTopicArnInput() *string
	FailureFeedbackRoleArn() *string
	SetFailureFeedbackRoleArn(val *string)
	FailureFeedbackRoleArnInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Platform() *string
	SetPlatform(val *string)
	PlatformCredential() *string
	SetPlatformCredential(val *string)
	PlatformCredentialInput() *string
	PlatformInput() *string
	PlatformPrincipal() *string
	SetPlatformPrincipal(val *string)
	PlatformPrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SuccessFeedbackRoleArn() *string
	SetSuccessFeedbackRoleArn(val *string)
	SuccessFeedbackRoleArnInput() *string
	SuccessFeedbackSampleRate() *string
	SetSuccessFeedbackSampleRate(val *string)
	SuccessFeedbackSampleRateInput() *string
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
	ResetEventDeliveryFailureTopicArn()
	ResetEventEndpointCreatedTopicArn()
	ResetEventEndpointDeletedTopicArn()
	ResetEventEndpointUpdatedTopicArn()
	ResetFailureFeedbackRoleArn()
	ResetOverrideLogicalId()
	ResetPlatformPrincipal()
	ResetSuccessFeedbackRoleArn()
	ResetSuccessFeedbackSampleRate()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsPlatformApplication
type jsiiProxy_SnsPlatformApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsPlatformApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventDeliveryFailureTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventDeliveryFailureTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointCreatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointCreatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointDeletedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointDeletedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointUpdatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointUpdatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformCredential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformCredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackSampleRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackSampleRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html aws_sns_platform_application} Resource.
func NewSnsPlatformApplication(scope constructs.Construct, id *string, config *SnsPlatformApplicationConfig) SnsPlatformApplication {
	_init_.Initialize()

	j := jsiiProxy_SnsPlatformApplication{}

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsPlatformApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html aws_sns_platform_application} Resource.
func NewSnsPlatformApplication_Override(s SnsPlatformApplication, scope constructs.Construct, id *string, config *SnsPlatformApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsPlatformApplication",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetEventDeliveryFailureTopicArn(val *string) {
	_jsii_.Set(
		j,
		"eventDeliveryFailureTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetEventEndpointCreatedTopicArn(val *string) {
	_jsii_.Set(
		j,
		"eventEndpointCreatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetEventEndpointDeletedTopicArn(val *string) {
	_jsii_.Set(
		j,
		"eventEndpointDeletedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetEventEndpointUpdatedTopicArn(val *string) {
	_jsii_.Set(
		j,
		"eventEndpointUpdatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"failureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetPlatform(val *string) {
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetPlatformCredential(val *string) {
	_jsii_.Set(
		j,
		"platformCredential",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetPlatformPrincipal(val *string) {
	_jsii_.Set(
		j,
		"platformPrincipal",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"successFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication) SetSuccessFeedbackSampleRate(val *string) {
	_jsii_.Set(
		j,
		"successFeedbackSampleRate",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SnsPlatformApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.SnsPlatformApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsPlatformApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.SnsPlatformApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SnsPlatformApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SnsPlatformApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsPlatformApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SnsPlatformApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SnsPlatformApplication) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SnsPlatformApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsPlatformApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventDeliveryFailureTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventDeliveryFailureTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointCreatedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointCreatedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointDeletedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointDeletedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointUpdatedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointUpdatedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SnsPlatformApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetPlatformPrincipal() {
	_jsii_.InvokeVoid(
		s,
		"resetPlatformPrincipal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SnsPlatformApplication) ToMetadata() interface{} {
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
func (s *jsiiProxy_SnsPlatformApplication) ToString() *string {
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
func (s *jsiiProxy_SnsPlatformApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsPlatformApplicationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#name SnsPlatformApplication#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#platform SnsPlatformApplication#platform}.
	Platform *string `json:"platform"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#platform_credential SnsPlatformApplication#platform_credential}.
	PlatformCredential *string `json:"platformCredential"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#event_delivery_failure_topic_arn SnsPlatformApplication#event_delivery_failure_topic_arn}.
	EventDeliveryFailureTopicArn *string `json:"eventDeliveryFailureTopicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#event_endpoint_created_topic_arn SnsPlatformApplication#event_endpoint_created_topic_arn}.
	EventEndpointCreatedTopicArn *string `json:"eventEndpointCreatedTopicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#event_endpoint_deleted_topic_arn SnsPlatformApplication#event_endpoint_deleted_topic_arn}.
	EventEndpointDeletedTopicArn *string `json:"eventEndpointDeletedTopicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#event_endpoint_updated_topic_arn SnsPlatformApplication#event_endpoint_updated_topic_arn}.
	EventEndpointUpdatedTopicArn *string `json:"eventEndpointUpdatedTopicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#failure_feedback_role_arn SnsPlatformApplication#failure_feedback_role_arn}.
	FailureFeedbackRoleArn *string `json:"failureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#platform_principal SnsPlatformApplication#platform_principal}.
	PlatformPrincipal *string `json:"platformPrincipal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#success_feedback_role_arn SnsPlatformApplication#success_feedback_role_arn}.
	SuccessFeedbackRoleArn *string `json:"successFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_platform_application.html#success_feedback_sample_rate SnsPlatformApplication#success_feedback_sample_rate}.
	SuccessFeedbackSampleRate *string `json:"successFeedbackSampleRate"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html aws_sns_sms_preferences}.
type SnsSmsPreferences interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultSenderId() *string
	SetDefaultSenderId(val *string)
	DefaultSenderIdInput() *string
	DefaultSmsType() *string
	SetDefaultSmsType(val *string)
	DefaultSmsTypeInput() *string
	DeliveryStatusIamRoleArn() *string
	SetDeliveryStatusIamRoleArn(val *string)
	DeliveryStatusIamRoleArnInput() *string
	DeliveryStatusSuccessSamplingRate() *string
	SetDeliveryStatusSuccessSamplingRate(val *string)
	DeliveryStatusSuccessSamplingRateInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MonthlySpendLimit() *string
	SetMonthlySpendLimit(val *string)
	MonthlySpendLimitInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UsageReportS3Bucket() *string
	SetUsageReportS3Bucket(val *string)
	UsageReportS3BucketInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDefaultSenderId()
	ResetDefaultSmsType()
	ResetDeliveryStatusIamRoleArn()
	ResetDeliveryStatusSuccessSamplingRate()
	ResetMonthlySpendLimit()
	ResetOverrideLogicalId()
	ResetUsageReportS3Bucket()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsSmsPreferences
type jsiiProxy_SnsSmsPreferences struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsSmsPreferences) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSenderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSenderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSmsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSmsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSmsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSmsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusSuccessSamplingRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusSuccessSamplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusSuccessSamplingRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusSuccessSamplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) MonthlySpendLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlySpendLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) MonthlySpendLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlySpendLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) UsageReportS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageReportS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) UsageReportS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageReportS3BucketInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html aws_sns_sms_preferences} Resource.
func NewSnsSmsPreferences(scope constructs.Construct, id *string, config *SnsSmsPreferencesConfig) SnsSmsPreferences {
	_init_.Initialize()

	j := jsiiProxy_SnsSmsPreferences{}

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsSmsPreferences",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html aws_sns_sms_preferences} Resource.
func NewSnsSmsPreferences_Override(s SnsSmsPreferences, scope constructs.Construct, id *string, config *SnsSmsPreferencesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsSmsPreferences",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetDefaultSenderId(val *string) {
	_jsii_.Set(
		j,
		"defaultSenderId",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetDefaultSmsType(val *string) {
	_jsii_.Set(
		j,
		"defaultSmsType",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetDeliveryStatusIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"deliveryStatusIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetDeliveryStatusSuccessSamplingRate(val *string) {
	_jsii_.Set(
		j,
		"deliveryStatusSuccessSamplingRate",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetMonthlySpendLimit(val *string) {
	_jsii_.Set(
		j,
		"monthlySpendLimit",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences) SetUsageReportS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"usageReportS3Bucket",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SnsSmsPreferences_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.SnsSmsPreferences",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsSmsPreferences_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.SnsSmsPreferences",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SnsSmsPreferences) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SnsSmsPreferences) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsSmsPreferences) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SnsSmsPreferences) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SnsSmsPreferences) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SnsSmsPreferences) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsSmsPreferences) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDefaultSenderId() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultSenderId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDefaultSmsType() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultSmsType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDeliveryStatusIamRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryStatusIamRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDeliveryStatusSuccessSamplingRate() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryStatusSuccessSamplingRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetMonthlySpendLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetMonthlySpendLimit",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SnsSmsPreferences) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetUsageReportS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetUsageReportS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SnsSmsPreferences) ToMetadata() interface{} {
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
func (s *jsiiProxy_SnsSmsPreferences) ToString() *string {
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
func (s *jsiiProxy_SnsSmsPreferences) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsSmsPreferencesConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#default_sender_id SnsSmsPreferences#default_sender_id}.
	DefaultSenderId *string `json:"defaultSenderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#default_sms_type SnsSmsPreferences#default_sms_type}.
	DefaultSmsType *string `json:"defaultSmsType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#delivery_status_iam_role_arn SnsSmsPreferences#delivery_status_iam_role_arn}.
	DeliveryStatusIamRoleArn *string `json:"deliveryStatusIamRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#delivery_status_success_sampling_rate SnsSmsPreferences#delivery_status_success_sampling_rate}.
	DeliveryStatusSuccessSamplingRate *string `json:"deliveryStatusSuccessSamplingRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#monthly_spend_limit SnsSmsPreferences#monthly_spend_limit}.
	MonthlySpendLimit *string `json:"monthlySpendLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_sms_preferences.html#usage_report_s3_bucket SnsSmsPreferences#usage_report_s3_bucket}.
	UsageReportS3Bucket *string `json:"usageReportS3Bucket"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html aws_sns_topic}.
type SnsTopic interface {
	cdktf.TerraformResource
	ApplicationFailureFeedbackRoleArn() *string
	SetApplicationFailureFeedbackRoleArn(val *string)
	ApplicationFailureFeedbackRoleArnInput() *string
	ApplicationSuccessFeedbackRoleArn() *string
	SetApplicationSuccessFeedbackRoleArn(val *string)
	ApplicationSuccessFeedbackRoleArnInput() *string
	ApplicationSuccessFeedbackSampleRate() *float64
	SetApplicationSuccessFeedbackSampleRate(val *float64)
	ApplicationSuccessFeedbackSampleRateInput() *float64
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentBasedDeduplication() interface{}
	SetContentBasedDeduplication(val interface{})
	ContentBasedDeduplicationInput() interface{}
	Count() interface{}
	SetCount(val interface{})
	DeliveryPolicy() *string
	SetDeliveryPolicy(val *string)
	DeliveryPolicyInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FifoTopic() interface{}
	SetFifoTopic(val interface{})
	FifoTopicInput() interface{}
	FirehoseFailureFeedbackRoleArn() *string
	SetFirehoseFailureFeedbackRoleArn(val *string)
	FirehoseFailureFeedbackRoleArnInput() *string
	FirehoseSuccessFeedbackRoleArn() *string
	SetFirehoseSuccessFeedbackRoleArn(val *string)
	FirehoseSuccessFeedbackRoleArnInput() *string
	FirehoseSuccessFeedbackSampleRate() *float64
	SetFirehoseSuccessFeedbackSampleRate(val *float64)
	FirehoseSuccessFeedbackSampleRateInput() *float64
	Fqn() *string
	FriendlyUniqueId() *string
	HttpFailureFeedbackRoleArn() *string
	SetHttpFailureFeedbackRoleArn(val *string)
	HttpFailureFeedbackRoleArnInput() *string
	HttpSuccessFeedbackRoleArn() *string
	SetHttpSuccessFeedbackRoleArn(val *string)
	HttpSuccessFeedbackRoleArnInput() *string
	HttpSuccessFeedbackSampleRate() *float64
	SetHttpSuccessFeedbackSampleRate(val *float64)
	HttpSuccessFeedbackSampleRateInput() *float64
	Id() *string
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
	KmsMasterKeyIdInput() *string
	LambdaFailureFeedbackRoleArn() *string
	SetLambdaFailureFeedbackRoleArn(val *string)
	LambdaFailureFeedbackRoleArnInput() *string
	LambdaSuccessFeedbackRoleArn() *string
	SetLambdaSuccessFeedbackRoleArn(val *string)
	LambdaSuccessFeedbackRoleArnInput() *string
	LambdaSuccessFeedbackSampleRate() *float64
	SetLambdaSuccessFeedbackSampleRate(val *float64)
	LambdaSuccessFeedbackSampleRateInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Owner() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SqsFailureFeedbackRoleArn() *string
	SetSqsFailureFeedbackRoleArn(val *string)
	SqsFailureFeedbackRoleArnInput() *string
	SqsSuccessFeedbackRoleArn() *string
	SetSqsSuccessFeedbackRoleArn(val *string)
	SqsSuccessFeedbackRoleArnInput() *string
	SqsSuccessFeedbackSampleRate() *float64
	SetSqsSuccessFeedbackSampleRate(val *float64)
	SqsSuccessFeedbackSampleRateInput() *float64
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
	ResetApplicationFailureFeedbackRoleArn()
	ResetApplicationSuccessFeedbackRoleArn()
	ResetApplicationSuccessFeedbackSampleRate()
	ResetContentBasedDeduplication()
	ResetDeliveryPolicy()
	ResetDisplayName()
	ResetFifoTopic()
	ResetFirehoseFailureFeedbackRoleArn()
	ResetFirehoseSuccessFeedbackRoleArn()
	ResetFirehoseSuccessFeedbackSampleRate()
	ResetHttpFailureFeedbackRoleArn()
	ResetHttpSuccessFeedbackRoleArn()
	ResetHttpSuccessFeedbackSampleRate()
	ResetKmsMasterKeyId()
	ResetLambdaFailureFeedbackRoleArn()
	ResetLambdaSuccessFeedbackRoleArn()
	ResetLambdaSuccessFeedbackSampleRate()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetSqsFailureFeedbackRoleArn()
	ResetSqsSuccessFeedbackRoleArn()
	ResetSqsSuccessFeedbackSampleRate()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsTopic
type jsiiProxy_SnsTopic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopic) ApplicationFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html aws_sns_topic} Resource.
func NewSnsTopic(scope constructs.Construct, id *string, config *SnsTopicConfig) SnsTopic {
	_init_.Initialize()

	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html aws_sns_topic} Resource.
func NewSnsTopic_Override(s SnsTopic, scope constructs.Construct, id *string, config *SnsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopic",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopic) SetApplicationFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"applicationFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetApplicationSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetApplicationSuccessFeedbackSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetContentBasedDeduplication(val interface{}) {
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetDeliveryPolicy(val *string) {
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetFifoTopic(val interface{}) {
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetFirehoseFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"firehoseFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetFirehoseSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetFirehoseSuccessFeedbackSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetHttpFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"httpFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetHttpSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"httpSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetHttpSuccessFeedbackSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"httpSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetKmsMasterKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetLambdaFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetLambdaSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetLambdaSuccessFeedbackSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetSqsFailureFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"sqsFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetSqsSuccessFeedbackRoleArn(val *string) {
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetSqsSuccessFeedbackSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SnsTopic) SetTagsAll(val interface{}) {
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
func SnsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.SnsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.SnsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopic) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SnsTopic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SnsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SnsTopic) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SnsTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopic) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		s,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SnsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopic) ToMetadata() interface{} {
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
func (s *jsiiProxy_SnsTopic) ToString() *string {
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
func (s *jsiiProxy_SnsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsTopicConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#application_failure_feedback_role_arn SnsTopic#application_failure_feedback_role_arn}.
	ApplicationFailureFeedbackRoleArn *string `json:"applicationFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#application_success_feedback_role_arn SnsTopic#application_success_feedback_role_arn}.
	ApplicationSuccessFeedbackRoleArn *string `json:"applicationSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#application_success_feedback_sample_rate SnsTopic#application_success_feedback_sample_rate}.
	ApplicationSuccessFeedbackSampleRate *float64 `json:"applicationSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#content_based_deduplication SnsTopic#content_based_deduplication}.
	ContentBasedDeduplication interface{} `json:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#delivery_policy SnsTopic#delivery_policy}.
	DeliveryPolicy *string `json:"deliveryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#display_name SnsTopic#display_name}.
	DisplayName *string `json:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#fifo_topic SnsTopic#fifo_topic}.
	FifoTopic interface{} `json:"fifoTopic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#firehose_failure_feedback_role_arn SnsTopic#firehose_failure_feedback_role_arn}.
	FirehoseFailureFeedbackRoleArn *string `json:"firehoseFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#firehose_success_feedback_role_arn SnsTopic#firehose_success_feedback_role_arn}.
	FirehoseSuccessFeedbackRoleArn *string `json:"firehoseSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#firehose_success_feedback_sample_rate SnsTopic#firehose_success_feedback_sample_rate}.
	FirehoseSuccessFeedbackSampleRate *float64 `json:"firehoseSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#http_failure_feedback_role_arn SnsTopic#http_failure_feedback_role_arn}.
	HttpFailureFeedbackRoleArn *string `json:"httpFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#http_success_feedback_role_arn SnsTopic#http_success_feedback_role_arn}.
	HttpSuccessFeedbackRoleArn *string `json:"httpSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#http_success_feedback_sample_rate SnsTopic#http_success_feedback_sample_rate}.
	HttpSuccessFeedbackSampleRate *float64 `json:"httpSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#kms_master_key_id SnsTopic#kms_master_key_id}.
	KmsMasterKeyId *string `json:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#lambda_failure_feedback_role_arn SnsTopic#lambda_failure_feedback_role_arn}.
	LambdaFailureFeedbackRoleArn *string `json:"lambdaFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#lambda_success_feedback_role_arn SnsTopic#lambda_success_feedback_role_arn}.
	LambdaSuccessFeedbackRoleArn *string `json:"lambdaSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#lambda_success_feedback_sample_rate SnsTopic#lambda_success_feedback_sample_rate}.
	LambdaSuccessFeedbackSampleRate *float64 `json:"lambdaSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#name SnsTopic#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#name_prefix SnsTopic#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#policy SnsTopic#policy}.
	Policy *string `json:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#sqs_failure_feedback_role_arn SnsTopic#sqs_failure_feedback_role_arn}.
	SqsFailureFeedbackRoleArn *string `json:"sqsFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#sqs_success_feedback_role_arn SnsTopic#sqs_success_feedback_role_arn}.
	SqsSuccessFeedbackRoleArn *string `json:"sqsSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#sqs_success_feedback_sample_rate SnsTopic#sqs_success_feedback_sample_rate}.
	SqsSuccessFeedbackSampleRate *float64 `json:"sqsSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#tags SnsTopic#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic.html#tags_all SnsTopic#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_policy.html aws_sns_topic_policy}.
type SnsTopicPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	Owner() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
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

// The jsii proxy struct for SnsTopicPolicy
type jsiiProxy_SnsTopicPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopicPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_policy.html aws_sns_topic_policy} Resource.
func NewSnsTopicPolicy(scope constructs.Construct, id *string, config *SnsTopicPolicyConfig) SnsTopicPolicy {
	_init_.Initialize()

	j := jsiiProxy_SnsTopicPolicy{}

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopicPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_policy.html aws_sns_topic_policy} Resource.
func NewSnsTopicPolicy_Override(s SnsTopicPolicy, scope constructs.Construct, id *string, config *SnsTopicPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopicPolicy",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func SnsTopicPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.SnsTopicPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopicPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.SnsTopicPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopicPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SnsTopicPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopicPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SnsTopicPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SnsTopicPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SnsTopicPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopicPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SnsTopicPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopicPolicy) ToMetadata() interface{} {
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
func (s *jsiiProxy_SnsTopicPolicy) ToString() *string {
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
func (s *jsiiProxy_SnsTopicPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsTopicPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_policy.html#arn SnsTopicPolicy#arn}.
	Arn *string `json:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_policy.html#policy SnsTopicPolicy#policy}.
	Policy *string `json:"policy"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html aws_sns_topic_subscription}.
type SnsTopicSubscription interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfirmationTimeoutInMinutes() *float64
	SetConfirmationTimeoutInMinutes(val *float64)
	ConfirmationTimeoutInMinutesInput() *float64
	ConfirmationWasAuthenticated() interface{}
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DeliveryPolicy() *string
	SetDeliveryPolicy(val *string)
	DeliveryPolicyInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointAutoConfirms() interface{}
	SetEndpointAutoConfirms(val interface{})
	EndpointAutoConfirmsInput() interface{}
	EndpointInput() *string
	FilterPolicy() *string
	SetFilterPolicy(val *string)
	FilterPolicyInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OwnerId() *string
	PendingConfirmation() interface{}
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawMessageDelivery() interface{}
	SetRawMessageDelivery(val interface{})
	RawMessageDeliveryInput() interface{}
	RawOverrides() interface{}
	RedrivePolicy() *string
	SetRedrivePolicy(val *string)
	RedrivePolicyInput() *string
	SubscriptionRoleArn() *string
	SetSubscriptionRoleArn(val *string)
	SubscriptionRoleArnInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TopicArn() *string
	SetTopicArn(val *string)
	TopicArnInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetConfirmationTimeoutInMinutes()
	ResetDeliveryPolicy()
	ResetEndpointAutoConfirms()
	ResetFilterPolicy()
	ResetOverrideLogicalId()
	ResetRawMessageDelivery()
	ResetRedrivePolicy()
	ResetSubscriptionRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsTopicSubscription
type jsiiProxy_SnsTopicSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopicSubscription) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confirmationTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confirmationTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationWasAuthenticated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationWasAuthenticated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointAutoConfirms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointAutoConfirms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointAutoConfirmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointAutoConfirmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) PendingConfirmation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pendingConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawMessageDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawMessageDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawMessageDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawMessageDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RedrivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RedrivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) SubscriptionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) SubscriptionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArnInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html aws_sns_topic_subscription} Resource.
func NewSnsTopicSubscription(scope constructs.Construct, id *string, config *SnsTopicSubscriptionConfig) SnsTopicSubscription {
	_init_.Initialize()

	j := jsiiProxy_SnsTopicSubscription{}

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopicSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html aws_sns_topic_subscription} Resource.
func NewSnsTopicSubscription_Override(s SnsTopicSubscription, scope constructs.Construct, id *string, config *SnsTopicSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SNS.SnsTopicSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetConfirmationTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"confirmationTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetDeliveryPolicy(val *string) {
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetEndpointAutoConfirms(val interface{}) {
	_jsii_.Set(
		j,
		"endpointAutoConfirms",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetFilterPolicy(val *string) {
	_jsii_.Set(
		j,
		"filterPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetRawMessageDelivery(val interface{}) {
	_jsii_.Set(
		j,
		"rawMessageDelivery",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetRedrivePolicy(val *string) {
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetSubscriptionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"subscriptionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription) SetTopicArn(val *string) {
	_jsii_.Set(
		j,
		"topicArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SnsTopicSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.SNS.SnsTopicSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopicSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.SNS.SnsTopicSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopicSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SnsTopicSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopicSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SnsTopicSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SnsTopicSubscription) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SnsTopicSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SnsTopicSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetConfirmationTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetConfirmationTimeoutInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetEndpointAutoConfirms() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointAutoConfirms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetFilterPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetFilterPolicy",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SnsTopicSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetRawMessageDelivery() {
	_jsii_.InvokeVoid(
		s,
		"resetRawMessageDelivery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetRedrivePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRedrivePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetSubscriptionRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSubscriptionRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SnsTopicSubscription) ToMetadata() interface{} {
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
func (s *jsiiProxy_SnsTopicSubscription) ToString() *string {
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
func (s *jsiiProxy_SnsTopicSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsTopicSubscriptionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#endpoint SnsTopicSubscription#endpoint}.
	Endpoint *string `json:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#protocol SnsTopicSubscription#protocol}.
	Protocol *string `json:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#topic_arn SnsTopicSubscription#topic_arn}.
	TopicArn *string `json:"topicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#confirmation_timeout_in_minutes SnsTopicSubscription#confirmation_timeout_in_minutes}.
	ConfirmationTimeoutInMinutes *float64 `json:"confirmationTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#delivery_policy SnsTopicSubscription#delivery_policy}.
	DeliveryPolicy *string `json:"deliveryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#endpoint_auto_confirms SnsTopicSubscription#endpoint_auto_confirms}.
	EndpointAutoConfirms interface{} `json:"endpointAutoConfirms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#filter_policy SnsTopicSubscription#filter_policy}.
	FilterPolicy *string `json:"filterPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#raw_message_delivery SnsTopicSubscription#raw_message_delivery}.
	RawMessageDelivery interface{} `json:"rawMessageDelivery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#redrive_policy SnsTopicSubscription#redrive_policy}.
	RedrivePolicy *string `json:"redrivePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sns_topic_subscription.html#subscription_role_arn SnsTopicSubscription#subscription_role_arn}.
	SubscriptionRoleArn *string `json:"subscriptionRoleArn"`
}
