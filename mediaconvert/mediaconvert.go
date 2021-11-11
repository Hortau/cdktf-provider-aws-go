package mediaconvert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/mediaconvert/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html aws_media_convert_queue}.
type MediaConvertQueue interface {
	cdktf.TerraformResource
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
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PricingPlan() *string
	SetPricingPlan(val *string)
	PricingPlanInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReservationPlanSettings() MediaConvertQueueReservationPlanSettingsOutputReference
	ReservationPlanSettingsInput() *MediaConvertQueueReservationPlanSettings
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
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutReservationPlanSettings(value *MediaConvertQueueReservationPlanSettings)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetPricingPlan()
	ResetReservationPlanSettings()
	ResetStatus()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MediaConvertQueue
type jsiiProxy_MediaConvertQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaConvertQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) PricingPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) PricingPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) ReservationPlanSettings() MediaConvertQueueReservationPlanSettingsOutputReference {
	var returns MediaConvertQueueReservationPlanSettingsOutputReference
	_jsii_.Get(
		j,
		"reservationPlanSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) ReservationPlanSettingsInput() *MediaConvertQueueReservationPlanSettings {
	var returns *MediaConvertQueueReservationPlanSettings
	_jsii_.Get(
		j,
		"reservationPlanSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html aws_media_convert_queue} Resource.
func NewMediaConvertQueue(scope constructs.Construct, id *string, config *MediaConvertQueueConfig) MediaConvertQueue {
	_init_.Initialize()

	j := jsiiProxy_MediaConvertQueue{}

	_jsii_.Create(
		"hashicorp_aws.MediaConvert.MediaConvertQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html aws_media_convert_queue} Resource.
func NewMediaConvertQueue_Override(m MediaConvertQueue, scope constructs.Construct, id *string, config *MediaConvertQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MediaConvert.MediaConvertQueue",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetPricingPlan(val *string) {
	_jsii_.Set(
		j,
		"pricingPlan",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueue) SetTagsAll(val interface{}) {
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
func MediaConvertQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.MediaConvert.MediaConvertQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaConvertQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.MediaConvert.MediaConvertQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MediaConvertQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MediaConvertQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_MediaConvertQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_MediaConvertQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_MediaConvertQueue) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_MediaConvertQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (m *jsiiProxy_MediaConvertQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaConvertQueue) PutReservationPlanSettings(value *MediaConvertQueueReservationPlanSettings) {
	_jsii_.InvokeVoid(
		m,
		"putReservationPlanSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MediaConvertQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetPricingPlan() {
	_jsii_.InvokeVoid(
		m,
		"resetPricingPlan",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetReservationPlanSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetReservationPlanSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaConvertQueue) SynthesizeAttributes() *map[string]interface{} {
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
func (m *jsiiProxy_MediaConvertQueue) ToMetadata() interface{} {
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
func (m *jsiiProxy_MediaConvertQueue) ToString() *string {
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
func (m *jsiiProxy_MediaConvertQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaConvertQueueConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#name MediaConvertQueue#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#description MediaConvertQueue#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#pricing_plan MediaConvertQueue#pricing_plan}.
	PricingPlan *string `json:"pricingPlan"`
	// reservation_plan_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#reservation_plan_settings MediaConvertQueue#reservation_plan_settings}
	ReservationPlanSettings *MediaConvertQueueReservationPlanSettings `json:"reservationPlanSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#status MediaConvertQueue#status}.
	Status *string `json:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#tags MediaConvertQueue#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#tags_all MediaConvertQueue#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type MediaConvertQueueReservationPlanSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#commitment MediaConvertQueue#commitment}.
	Commitment *string `json:"commitment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#renewal_type MediaConvertQueue#renewal_type}.
	RenewalType *string `json:"renewalType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/media_convert_queue.html#reserved_slots MediaConvertQueue#reserved_slots}.
	ReservedSlots *float64 `json:"reservedSlots"`
}

type MediaConvertQueueReservationPlanSettingsOutputReference interface {
	cdktf.ComplexObject
	Commitment() *string
	SetCommitment(val *string)
	CommitmentInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RenewalType() *string
	SetRenewalType(val *string)
	RenewalTypeInput() *string
	ReservedSlots() *float64
	SetReservedSlots(val *float64)
	ReservedSlotsInput() *float64
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

// The jsii proxy struct for MediaConvertQueueReservationPlanSettingsOutputReference
type jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) Commitment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) CommitmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) RenewalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) RenewalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) ReservedSlots() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedSlots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) ReservedSlotsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedSlotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMediaConvertQueueReservationPlanSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MediaConvertQueueReservationPlanSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MediaConvert.MediaConvertQueueReservationPlanSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMediaConvertQueueReservationPlanSettingsOutputReference_Override(m MediaConvertQueueReservationPlanSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MediaConvert.MediaConvertQueueReservationPlanSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetCommitment(val *string) {
	_jsii_.Set(
		j,
		"commitment",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetRenewalType(val *string) {
	_jsii_.Set(
		j,
		"renewalType",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetReservedSlots(val *float64) {
	_jsii_.Set(
		j,
		"reservedSlots",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (m *jsiiProxy_MediaConvertQueueReservationPlanSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
