package dlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/dlm/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy aws_dlm_lifecycle_policy}.
type DlmLifecyclePolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PolicyDetails() DlmLifecyclePolicyPolicyDetailsOutputReference
	PolicyDetailsInput() *DlmLifecyclePolicyPolicyDetails
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	PutPolicyDetails(value *DlmLifecyclePolicyPolicyDetails)
	ResetOverrideLogicalId()
	ResetState()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DlmLifecyclePolicy
type jsiiProxy_DlmLifecyclePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DlmLifecyclePolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) PolicyDetails() DlmLifecyclePolicyPolicyDetailsOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsOutputReference
	_jsii_.Get(
		j,
		"policyDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) PolicyDetailsInput() *DlmLifecyclePolicyPolicyDetails {
	var returns *DlmLifecyclePolicyPolicyDetails
	_jsii_.Get(
		j,
		"policyDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy aws_dlm_lifecycle_policy} Resource.
func NewDlmLifecyclePolicy(scope constructs.Construct, id *string, config *DlmLifecyclePolicyConfig) DlmLifecyclePolicy {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicy{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy aws_dlm_lifecycle_policy} Resource.
func NewDlmLifecyclePolicy_Override(d DlmLifecyclePolicy, scope constructs.Construct, id *string, config *DlmLifecyclePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicy) SetTagsAll(val *map[string]*string) {
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
func DlmLifecyclePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.dlm.DlmLifecyclePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DlmLifecyclePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.dlm.DlmLifecyclePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicy) PutPolicyDetails(value *DlmLifecyclePolicyPolicyDetails) {
	_jsii_.InvokeVoid(
		d,
		"putPolicyDetails",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicy) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicy) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DlmLifecyclePolicy) ToMetadata() interface{} {
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
func (d *jsiiProxy_DlmLifecyclePolicy) ToString() *string {
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
func (d *jsiiProxy_DlmLifecyclePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Data Lifecycle Manager.
type DlmLifecyclePolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#description DlmLifecyclePolicy#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#execution_role_arn DlmLifecyclePolicy#execution_role_arn}.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// policy_details block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#policy_details DlmLifecyclePolicy#policy_details}
	PolicyDetails *DlmLifecyclePolicyPolicyDetails `json:"policyDetails" yaml:"policyDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#state DlmLifecyclePolicy#state}.
	State *string `json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#tags DlmLifecyclePolicy#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#tags_all DlmLifecyclePolicy#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type DlmLifecyclePolicyPolicyDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#resource_types DlmLifecyclePolicy#resource_types}.
	ResourceTypes *[]*string `json:"resourceTypes" yaml:"resourceTypes"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#schedule DlmLifecyclePolicy#schedule}
	Schedule interface{} `json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#target_tags DlmLifecyclePolicy#target_tags}.
	TargetTags *map[string]*string `json:"targetTags" yaml:"targetTags"`
}

type DlmLifecyclePolicyPolicyDetailsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *DlmLifecyclePolicyPolicyDetails
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetails)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
	Schedule() interface{}
	SetSchedule(val interface{})
	ScheduleInput() interface{}
	TargetTags() *map[string]*string
	SetTargetTags(val *map[string]*string)
	TargetTagsInput() *map[string]*string
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

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetails {
	var returns *DlmLifecyclePolicyPolicyDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TargetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TargetTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDlmLifecyclePolicyPolicyDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DlmLifecyclePolicyPolicyDetailsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetInternalValue(val *DlmLifecyclePolicyPolicyDetails) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetResourceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetTargetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DlmLifecyclePolicyPolicyDetailsSchedule struct {
	// create_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#create_rule DlmLifecyclePolicy#create_rule}
	CreateRule *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule `json:"createRule" yaml:"createRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}.
	Name *string `json:"name" yaml:"name"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule `json:"retainRule" yaml:"retainRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}.
	CopyTags interface{} `json:"copyTags" yaml:"copyTags"`
	// cross_region_copy_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#cross_region_copy_rule DlmLifecyclePolicy#cross_region_copy_rule}
	CrossRegionCopyRule interface{} `json:"crossRegionCopyRule" yaml:"crossRegionCopyRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#tags_to_add DlmLifecyclePolicy#tags_to_add}.
	TagsToAdd *map[string]*string `json:"tagsToAdd" yaml:"tagsToAdd"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleCreateRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#times DlmLifecyclePolicy#times}.
	Times *[]*string `json:"times" yaml:"times"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IntervalUnit() *string
	SetIntervalUnit(val *string)
	IntervalUnitInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Times() *[]*string
	SetTimes(val *[]*string)
	TimesInput() *[]*string
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
	ResetIntervalUnit()
	ResetTimes()
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) Times() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"times",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) TimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timesInput",
		&returns,
	)
	return returns
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetIntervalUnit(val *string) {
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) SetTimes(val *[]*string) {
	_jsii_.Set(
		j,
		"times",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) ResetIntervalUnit() {
	_jsii_.InvokeVoid(
		d,
		"resetIntervalUnit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference) ResetTimes() {
	_jsii_.InvokeVoid(
		d,
		"resetTimes",
		nil, // no parameters
	)
}

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#encrypted DlmLifecyclePolicy#encrypted}.
	Encrypted interface{} `json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#target DlmLifecyclePolicy#target}.
	Target *string `json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#cmk_arn DlmLifecyclePolicy#cmk_arn}.
	CmkArn *string `json:"cmkArn" yaml:"cmkArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}.
	CopyTags interface{} `json:"copyTags" yaml:"copyTags"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#deprecate_rule DlmLifecyclePolicy#deprecate_rule}
	DeprecateRule *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule `json:"deprecateRule" yaml:"deprecateRule"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule `json:"retainRule" yaml:"retainRule"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `json:"intervalUnit" yaml:"intervalUnit"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IntervalUnit() *string
	SetIntervalUnit(val *string)
	IntervalUnitInput() *string
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
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetIntervalUnit(val *string) {
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `json:"intervalUnit" yaml:"intervalUnit"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IntervalUnit() *string
	SetIntervalUnit(val *string)
	IntervalUnitInput() *string
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
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetIntervalUnit(val *string) {
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DlmLifecyclePolicyPolicyDetailsScheduleRetainRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#count DlmLifecyclePolicy#count}.
	Count *float64 `json:"count" yaml:"count"`
}

type DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference interface {
	cdktf.ComplexObject
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule)
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
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.dlm.DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
