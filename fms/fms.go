package fms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/fms/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fms_admin_account.html aws_fms_admin_account}.
type FmsAdminAccount interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	ResetAccountId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FmsAdminAccount
type jsiiProxy_FmsAdminAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FmsAdminAccount) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsAdminAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fms_admin_account.html aws_fms_admin_account} Resource.
func NewFmsAdminAccount(scope constructs.Construct, id *string, config *FmsAdminAccountConfig) FmsAdminAccount {
	_init_.Initialize()

	j := jsiiProxy_FmsAdminAccount{}

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsAdminAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fms_admin_account.html aws_fms_admin_account} Resource.
func NewFmsAdminAccount_Override(f FmsAdminAccount, scope constructs.Construct, id *string, config *FmsAdminAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsAdminAccount",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FmsAdminAccount) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_FmsAdminAccount) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FmsAdminAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FmsAdminAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FmsAdminAccount) SetProvider(val cdktf.TerraformProvider) {
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
func FmsAdminAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.FMS.FmsAdminAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FmsAdminAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.FMS.FmsAdminAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FmsAdminAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FmsAdminAccount) ResetAccountId() {
	_jsii_.InvokeVoid(
		f,
		"resetAccountId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FmsAdminAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsAdminAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsAdminAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FmsAdminAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FmsAdminAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FmsAdminAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_admin_account.html#account_id FmsAdminAccount#account_id}.
	AccountId *string `json:"accountId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html aws_fms_policy}.
type FmsPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DeleteAllPolicyResources() interface{}
	SetDeleteAllPolicyResources(val interface{})
	DeleteAllPolicyResourcesInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExcludeMap() FmsPolicyExcludeMapOutputReference
	ExcludeMapInput() *FmsPolicyExcludeMap
	ExcludeResourceTags() interface{}
	SetExcludeResourceTags(val interface{})
	ExcludeResourceTagsInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IncludeMap() FmsPolicyIncludeMapOutputReference
	IncludeMapInput() *FmsPolicyIncludeMap
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PolicyUpdateToken() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RemediationEnabled() interface{}
	SetRemediationEnabled(val interface{})
	RemediationEnabledInput() interface{}
	ResourceTags() interface{}
	SetResourceTags(val interface{})
	ResourceTagsInput() interface{}
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	ResourceTypeList() *[]*string
	SetResourceTypeList(val *[]*string)
	ResourceTypeListInput() *[]*string
	SecurityServicePolicyData() FmsPolicySecurityServicePolicyDataOutputReference
	SecurityServicePolicyDataInput() *FmsPolicySecurityServicePolicyData
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
	PutExcludeMap(value *FmsPolicyExcludeMap)
	PutIncludeMap(value *FmsPolicyIncludeMap)
	PutSecurityServicePolicyData(value *FmsPolicySecurityServicePolicyData)
	ResetDeleteAllPolicyResources()
	ResetExcludeMap()
	ResetIncludeMap()
	ResetOverrideLogicalId()
	ResetRemediationEnabled()
	ResetResourceTags()
	ResetResourceType()
	ResetResourceTypeList()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FmsPolicy
type jsiiProxy_FmsPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FmsPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DeleteAllPolicyResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DeleteAllPolicyResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeMap() FmsPolicyExcludeMapOutputReference {
	var returns FmsPolicyExcludeMapOutputReference
	_jsii_.Get(
		j,
		"excludeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeMapInput() *FmsPolicyExcludeMap {
	var returns *FmsPolicyExcludeMap
	_jsii_.Get(
		j,
		"excludeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ExcludeResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) IncludeMap() FmsPolicyIncludeMapOutputReference {
	var returns FmsPolicyIncludeMapOutputReference
	_jsii_.Get(
		j,
		"includeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) IncludeMapInput() *FmsPolicyIncludeMap {
	var returns *FmsPolicyIncludeMap
	_jsii_.Get(
		j,
		"includeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) PolicyUpdateToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUpdateToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RemediationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) RemediationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) ResourceTypeListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) SecurityServicePolicyData() FmsPolicySecurityServicePolicyDataOutputReference {
	var returns FmsPolicySecurityServicePolicyDataOutputReference
	_jsii_.Get(
		j,
		"securityServicePolicyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) SecurityServicePolicyDataInput() *FmsPolicySecurityServicePolicyData {
	var returns *FmsPolicySecurityServicePolicyData
	_jsii_.Get(
		j,
		"securityServicePolicyDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html aws_fms_policy} Resource.
func NewFmsPolicy(scope constructs.Construct, id *string, config *FmsPolicyConfig) FmsPolicy {
	_init_.Initialize()

	j := jsiiProxy_FmsPolicy{}

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html aws_fms_policy} Resource.
func NewFmsPolicy_Override(f FmsPolicy, scope constructs.Construct, id *string, config *FmsPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicy",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FmsPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetDeleteAllPolicyResources(val interface{}) {
	_jsii_.Set(
		j,
		"deleteAllPolicyResources",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetExcludeResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"excludeResourceTags",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetRemediationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"remediationEnabled",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_FmsPolicy) SetResourceTypeList(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypeList",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FmsPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.FMS.FmsPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FmsPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.FMS.FmsPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FmsPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FmsPolicy) PutExcludeMap(value *FmsPolicyExcludeMap) {
	_jsii_.InvokeVoid(
		f,
		"putExcludeMap",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutIncludeMap(value *FmsPolicyIncludeMap) {
	_jsii_.InvokeVoid(
		f,
		"putIncludeMap",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) PutSecurityServicePolicyData(value *FmsPolicySecurityServicePolicyData) {
	_jsii_.InvokeVoid(
		f,
		"putSecurityServicePolicyData",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FmsPolicy) ResetDeleteAllPolicyResources() {
	_jsii_.InvokeVoid(
		f,
		"resetDeleteAllPolicyResources",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetExcludeMap() {
	_jsii_.InvokeVoid(
		f,
		"resetExcludeMap",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetIncludeMap() {
	_jsii_.InvokeVoid(
		f,
		"resetIncludeMap",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FmsPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetRemediationEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetRemediationEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceTags() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceType() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) ResetResourceTypeList() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceTypeList",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FmsPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FmsPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FmsPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#exclude_resource_tags FmsPolicy#exclude_resource_tags}.
	ExcludeResourceTags interface{} `json:"excludeResourceTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#name FmsPolicy#name}.
	Name *string `json:"name"`
	// security_service_policy_data block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#security_service_policy_data FmsPolicy#security_service_policy_data}
	SecurityServicePolicyData *FmsPolicySecurityServicePolicyData `json:"securityServicePolicyData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#delete_all_policy_resources FmsPolicy#delete_all_policy_resources}.
	DeleteAllPolicyResources interface{} `json:"deleteAllPolicyResources"`
	// exclude_map block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#exclude_map FmsPolicy#exclude_map}
	ExcludeMap *FmsPolicyExcludeMap `json:"excludeMap"`
	// include_map block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#include_map FmsPolicy#include_map}
	IncludeMap *FmsPolicyIncludeMap `json:"includeMap"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#remediation_enabled FmsPolicy#remediation_enabled}.
	RemediationEnabled interface{} `json:"remediationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#resource_tags FmsPolicy#resource_tags}.
	ResourceTags interface{} `json:"resourceTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#resource_type FmsPolicy#resource_type}.
	ResourceType *string `json:"resourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#resource_type_list FmsPolicy#resource_type_list}.
	ResourceTypeList *[]*string `json:"resourceTypeList"`
}

type FmsPolicyExcludeMap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#account FmsPolicy#account}.
	Account *[]*string `json:"account"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#orgunit FmsPolicy#orgunit}.
	Orgunit *[]*string `json:"orgunit"`
}

type FmsPolicyExcludeMapOutputReference interface {
	cdktf.ComplexObject
	Account() *[]*string
	SetAccount(val *[]*string)
	AccountInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Orgunit() *[]*string
	SetOrgunit(val *[]*string)
	OrgunitInput() *[]*string
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
	ResetAccount()
	ResetOrgunit()
}

// The jsii proxy struct for FmsPolicyExcludeMapOutputReference
type jsiiProxy_FmsPolicyExcludeMapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) Account() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) AccountInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) Orgunit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgunit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) OrgunitInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgunitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFmsPolicyExcludeMapOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FmsPolicyExcludeMapOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FmsPolicyExcludeMapOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicyExcludeMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFmsPolicyExcludeMapOutputReference_Override(f FmsPolicyExcludeMapOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicyExcludeMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) SetAccount(val *[]*string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) SetOrgunit(val *[]*string) {
	_jsii_.Set(
		j,
		"orgunit",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyExcludeMapOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		f,
		"resetAccount",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicyExcludeMapOutputReference) ResetOrgunit() {
	_jsii_.InvokeVoid(
		f,
		"resetOrgunit",
		nil, // no parameters
	)
}

type FmsPolicyIncludeMap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#account FmsPolicy#account}.
	Account *[]*string `json:"account"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#orgunit FmsPolicy#orgunit}.
	Orgunit *[]*string `json:"orgunit"`
}

type FmsPolicyIncludeMapOutputReference interface {
	cdktf.ComplexObject
	Account() *[]*string
	SetAccount(val *[]*string)
	AccountInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Orgunit() *[]*string
	SetOrgunit(val *[]*string)
	OrgunitInput() *[]*string
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
	ResetAccount()
	ResetOrgunit()
}

// The jsii proxy struct for FmsPolicyIncludeMapOutputReference
type jsiiProxy_FmsPolicyIncludeMapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) Account() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) AccountInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) Orgunit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgunit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) OrgunitInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgunitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFmsPolicyIncludeMapOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FmsPolicyIncludeMapOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FmsPolicyIncludeMapOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicyIncludeMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFmsPolicyIncludeMapOutputReference_Override(f FmsPolicyIncludeMapOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicyIncludeMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) SetAccount(val *[]*string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) SetOrgunit(val *[]*string) {
	_jsii_.Set(
		j,
		"orgunit",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FmsPolicyIncludeMapOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		f,
		"resetAccount",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FmsPolicyIncludeMapOutputReference) ResetOrgunit() {
	_jsii_.InvokeVoid(
		f,
		"resetOrgunit",
		nil, // no parameters
	)
}

type FmsPolicySecurityServicePolicyData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#type FmsPolicy#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy.html#managed_service_data FmsPolicy#managed_service_data}.
	ManagedServiceData *string `json:"managedServiceData"`
}

type FmsPolicySecurityServicePolicyDataOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ManagedServiceData() *string
	SetManagedServiceData(val *string)
	ManagedServiceDataInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetManagedServiceData()
}

// The jsii proxy struct for FmsPolicySecurityServicePolicyDataOutputReference
type jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) ManagedServiceData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServiceData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) ManagedServiceDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServiceDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewFmsPolicySecurityServicePolicyDataOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FmsPolicySecurityServicePolicyDataOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicySecurityServicePolicyDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFmsPolicySecurityServicePolicyDataOutputReference_Override(f FmsPolicySecurityServicePolicyDataOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FMS.FmsPolicySecurityServicePolicyDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) SetManagedServiceData(val *string) {
	_jsii_.Set(
		j,
		"managedServiceData",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FmsPolicySecurityServicePolicyDataOutputReference) ResetManagedServiceData() {
	_jsii_.InvokeVoid(
		f,
		"resetManagedServiceData",
		nil, // no parameters
	)
}
