package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/securityhub/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account aws_securityhub_account}.
type SecurityhubAccount interface {
	cdktf.TerraformResource
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

// The jsii proxy struct for SecurityhubAccount
type jsiiProxy_SecurityhubAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account aws_securityhub_account} Resource.
func NewSecurityhubAccount(scope constructs.Construct, id *string, config *SecurityhubAccountConfig) SecurityhubAccount {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubAccount{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account aws_securityhub_account} Resource.
func NewSecurityhubAccount_Override(s SecurityhubAccount, scope constructs.Construct, id *string, config *SecurityhubAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubAccount",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubAccount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAccount) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubAccount) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAccount) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubAccount) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubAccount) ToString() *string {
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
func (s *jsiiProxy_SecurityhubAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubAccountConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target aws_securityhub_action_target}.
type SecurityhubActionTarget interface {
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
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
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

// The jsii proxy struct for SecurityhubActionTarget
type jsiiProxy_SecurityhubActionTarget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubActionTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubActionTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target aws_securityhub_action_target} Resource.
func NewSecurityhubActionTarget(scope constructs.Construct, id *string, config *SecurityhubActionTargetConfig) SecurityhubActionTarget {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubActionTarget{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubActionTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target aws_securityhub_action_target} Resource.
func NewSecurityhubActionTarget_Override(s SecurityhubActionTarget, scope constructs.Construct, id *string, config *SecurityhubActionTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubActionTarget",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubActionTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubActionTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubActionTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubActionTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubActionTarget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubActionTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubActionTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubActionTarget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubActionTarget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubActionTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubActionTarget) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubActionTarget) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubActionTarget) ToString() *string {
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
func (s *jsiiProxy_SecurityhubActionTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubActionTargetConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target#description SecurityhubActionTarget#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target#identifier SecurityhubActionTarget#identifier}.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target#name SecurityhubActionTarget#name}.
	Name *string `json:"name" yaml:"name"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_finding_aggregator aws_securityhub_finding_aggregator}.
type SecurityhubFindingAggregator interface {
	cdktf.TerraformResource
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
	LinkingMode() *string
	SetLinkingMode(val *string)
	LinkingModeInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SpecifiedRegions() *[]*string
	SetSpecifiedRegions(val *[]*string)
	SpecifiedRegionsInput() *[]*string
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
	ResetSpecifiedRegions()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurityhubFindingAggregator
type jsiiProxy_SecurityhubFindingAggregator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubFindingAggregator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) LinkingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) LinkingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SpecifiedRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"specifiedRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SpecifiedRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"specifiedRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubFindingAggregator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_finding_aggregator aws_securityhub_finding_aggregator} Resource.
func NewSecurityhubFindingAggregator(scope constructs.Construct, id *string, config *SecurityhubFindingAggregatorConfig) SecurityhubFindingAggregator {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubFindingAggregator{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_finding_aggregator aws_securityhub_finding_aggregator} Resource.
func NewSecurityhubFindingAggregator_Override(s SecurityhubFindingAggregator, scope constructs.Construct, id *string, config *SecurityhubFindingAggregatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregator",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetLinkingMode(val *string) {
	_jsii_.Set(
		j,
		"linkingMode",
		val,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurityhubFindingAggregator) SetSpecifiedRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"specifiedRegions",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SecurityhubFindingAggregator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubFindingAggregator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubFindingAggregator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubFindingAggregator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubFindingAggregator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubFindingAggregator) ResetSpecifiedRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetSpecifiedRegions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubFindingAggregator) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) ToString() *string {
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
func (s *jsiiProxy_SecurityhubFindingAggregator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubFindingAggregatorConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_finding_aggregator#linking_mode SecurityhubFindingAggregator#linking_mode}.
	LinkingMode *string `json:"linkingMode" yaml:"linkingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_finding_aggregator#specified_regions SecurityhubFindingAggregator#specified_regions}.
	SpecifiedRegions *[]*string `json:"specifiedRegions" yaml:"specifiedRegions"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight aws_securityhub_insight}.
type SecurityhubInsight interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Filters() SecurityhubInsightFiltersOutputReference
	FiltersInput() *SecurityhubInsightFilters
	Fqn() *string
	FriendlyUniqueId() *string
	GroupByAttribute() *string
	SetGroupByAttribute(val *string)
	GroupByAttributeInput() *string
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
	PutFilters(value *SecurityhubInsightFilters)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurityhubInsight
type jsiiProxy_SecurityhubInsight struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubInsight) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Filters() SecurityhubInsightFiltersOutputReference {
	var returns SecurityhubInsightFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) FiltersInput() *SecurityhubInsightFilters {
	var returns *SecurityhubInsightFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) GroupByAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) GroupByAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsight) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight aws_securityhub_insight} Resource.
func NewSecurityhubInsight(scope constructs.Construct, id *string, config *SecurityhubInsightConfig) SecurityhubInsight {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsight{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsight",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight aws_securityhub_insight} Resource.
func NewSecurityhubInsight_Override(s SecurityhubInsight, scope constructs.Construct, id *string, config *SecurityhubInsightConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsight",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetGroupByAttribute(val *string) {
	_jsii_.Set(
		j,
		"groupByAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubInsight_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubInsight",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsight) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsight) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsight) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsight) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsight) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsight) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsight) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsight) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsight) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsight) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsight) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsight) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurityhubInsight) PutFilters(value *SecurityhubInsightFilters) {
	_jsii_.InvokeVoid(
		s,
		"putFilters",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubInsight) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsight) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsight) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubInsight) ToString() *string {
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
func (s *jsiiProxy_SecurityhubInsight) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubInsightConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#filters SecurityhubInsight#filters}
	Filters *SecurityhubInsightFilters `json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#group_by_attribute SecurityhubInsight#group_by_attribute}.
	GroupByAttribute *string `json:"groupByAttribute" yaml:"groupByAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#name SecurityhubInsight#name}.
	Name *string `json:"name" yaml:"name"`
}

type SecurityhubInsightFilters struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#aws_account_id SecurityhubInsight#aws_account_id}
	AwsAccountId interface{} `json:"awsAccountId" yaml:"awsAccountId"`
	// company_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#company_name SecurityhubInsight#company_name}
	CompanyName interface{} `json:"companyName" yaml:"companyName"`
	// compliance_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#compliance_status SecurityhubInsight#compliance_status}
	ComplianceStatus interface{} `json:"complianceStatus" yaml:"complianceStatus"`
	// confidence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#confidence SecurityhubInsight#confidence}
	Confidence interface{} `json:"confidence" yaml:"confidence"`
	// created_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#created_at SecurityhubInsight#created_at}
	CreatedAt interface{} `json:"createdAt" yaml:"createdAt"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#criticality SecurityhubInsight#criticality}
	Criticality interface{} `json:"criticality" yaml:"criticality"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#description SecurityhubInsight#description}
	Description interface{} `json:"description" yaml:"description"`
	// finding_provider_fields_confidence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_confidence SecurityhubInsight#finding_provider_fields_confidence}
	FindingProviderFieldsConfidence interface{} `json:"findingProviderFieldsConfidence" yaml:"findingProviderFieldsConfidence"`
	// finding_provider_fields_criticality block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_criticality SecurityhubInsight#finding_provider_fields_criticality}
	FindingProviderFieldsCriticality interface{} `json:"findingProviderFieldsCriticality" yaml:"findingProviderFieldsCriticality"`
	// finding_provider_fields_related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_related_findings_id SecurityhubInsight#finding_provider_fields_related_findings_id}
	FindingProviderFieldsRelatedFindingsId interface{} `json:"findingProviderFieldsRelatedFindingsId" yaml:"findingProviderFieldsRelatedFindingsId"`
	// finding_provider_fields_related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_related_findings_product_arn SecurityhubInsight#finding_provider_fields_related_findings_product_arn}
	FindingProviderFieldsRelatedFindingsProductArn interface{} `json:"findingProviderFieldsRelatedFindingsProductArn" yaml:"findingProviderFieldsRelatedFindingsProductArn"`
	// finding_provider_fields_severity_label block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_severity_label SecurityhubInsight#finding_provider_fields_severity_label}
	FindingProviderFieldsSeverityLabel interface{} `json:"findingProviderFieldsSeverityLabel" yaml:"findingProviderFieldsSeverityLabel"`
	// finding_provider_fields_severity_original block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_severity_original SecurityhubInsight#finding_provider_fields_severity_original}
	FindingProviderFieldsSeverityOriginal interface{} `json:"findingProviderFieldsSeverityOriginal" yaml:"findingProviderFieldsSeverityOriginal"`
	// finding_provider_fields_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#finding_provider_fields_types SecurityhubInsight#finding_provider_fields_types}
	FindingProviderFieldsTypes interface{} `json:"findingProviderFieldsTypes" yaml:"findingProviderFieldsTypes"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#first_observed_at SecurityhubInsight#first_observed_at}
	FirstObservedAt interface{} `json:"firstObservedAt" yaml:"firstObservedAt"`
	// generator_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#generator_id SecurityhubInsight#generator_id}
	GeneratorId interface{} `json:"generatorId" yaml:"generatorId"`
	// id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#id SecurityhubInsight#id}
	Id interface{} `json:"id" yaml:"id"`
	// keyword block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#keyword SecurityhubInsight#keyword}
	Keyword interface{} `json:"keyword" yaml:"keyword"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#last_observed_at SecurityhubInsight#last_observed_at}
	LastObservedAt interface{} `json:"lastObservedAt" yaml:"lastObservedAt"`
	// malware_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#malware_name SecurityhubInsight#malware_name}
	MalwareName interface{} `json:"malwareName" yaml:"malwareName"`
	// malware_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#malware_path SecurityhubInsight#malware_path}
	MalwarePath interface{} `json:"malwarePath" yaml:"malwarePath"`
	// malware_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#malware_state SecurityhubInsight#malware_state}
	MalwareState interface{} `json:"malwareState" yaml:"malwareState"`
	// malware_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#malware_type SecurityhubInsight#malware_type}
	MalwareType interface{} `json:"malwareType" yaml:"malwareType"`
	// network_destination_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_destination_domain SecurityhubInsight#network_destination_domain}
	NetworkDestinationDomain interface{} `json:"networkDestinationDomain" yaml:"networkDestinationDomain"`
	// network_destination_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_destination_ipv4 SecurityhubInsight#network_destination_ipv4}
	NetworkDestinationIpv4 interface{} `json:"networkDestinationIpv4" yaml:"networkDestinationIpv4"`
	// network_destination_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_destination_ipv6 SecurityhubInsight#network_destination_ipv6}
	NetworkDestinationIpv6 interface{} `json:"networkDestinationIpv6" yaml:"networkDestinationIpv6"`
	// network_destination_port block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_destination_port SecurityhubInsight#network_destination_port}
	NetworkDestinationPort interface{} `json:"networkDestinationPort" yaml:"networkDestinationPort"`
	// network_direction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_direction SecurityhubInsight#network_direction}
	NetworkDirection interface{} `json:"networkDirection" yaml:"networkDirection"`
	// network_protocol block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_protocol SecurityhubInsight#network_protocol}
	NetworkProtocol interface{} `json:"networkProtocol" yaml:"networkProtocol"`
	// network_source_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_source_domain SecurityhubInsight#network_source_domain}
	NetworkSourceDomain interface{} `json:"networkSourceDomain" yaml:"networkSourceDomain"`
	// network_source_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_source_ipv4 SecurityhubInsight#network_source_ipv4}
	NetworkSourceIpv4 interface{} `json:"networkSourceIpv4" yaml:"networkSourceIpv4"`
	// network_source_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_source_ipv6 SecurityhubInsight#network_source_ipv6}
	NetworkSourceIpv6 interface{} `json:"networkSourceIpv6" yaml:"networkSourceIpv6"`
	// network_source_mac block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_source_mac SecurityhubInsight#network_source_mac}
	NetworkSourceMac interface{} `json:"networkSourceMac" yaml:"networkSourceMac"`
	// network_source_port block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#network_source_port SecurityhubInsight#network_source_port}
	NetworkSourcePort interface{} `json:"networkSourcePort" yaml:"networkSourcePort"`
	// note_text block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#note_text SecurityhubInsight#note_text}
	NoteText interface{} `json:"noteText" yaml:"noteText"`
	// note_updated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#note_updated_at SecurityhubInsight#note_updated_at}
	NoteUpdatedAt interface{} `json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// note_updated_by block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#note_updated_by SecurityhubInsight#note_updated_by}
	NoteUpdatedBy interface{} `json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// process_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_launched_at SecurityhubInsight#process_launched_at}
	ProcessLaunchedAt interface{} `json:"processLaunchedAt" yaml:"processLaunchedAt"`
	// process_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_name SecurityhubInsight#process_name}
	ProcessName interface{} `json:"processName" yaml:"processName"`
	// process_parent_pid block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_parent_pid SecurityhubInsight#process_parent_pid}
	ProcessParentPid interface{} `json:"processParentPid" yaml:"processParentPid"`
	// process_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_path SecurityhubInsight#process_path}
	ProcessPath interface{} `json:"processPath" yaml:"processPath"`
	// process_pid block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_pid SecurityhubInsight#process_pid}
	ProcessPid interface{} `json:"processPid" yaml:"processPid"`
	// process_terminated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#process_terminated_at SecurityhubInsight#process_terminated_at}
	ProcessTerminatedAt interface{} `json:"processTerminatedAt" yaml:"processTerminatedAt"`
	// product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#product_arn SecurityhubInsight#product_arn}
	ProductArn interface{} `json:"productArn" yaml:"productArn"`
	// product_fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#product_fields SecurityhubInsight#product_fields}
	ProductFields interface{} `json:"productFields" yaml:"productFields"`
	// product_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#product_name SecurityhubInsight#product_name}
	ProductName interface{} `json:"productName" yaml:"productName"`
	// recommendation_text block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#recommendation_text SecurityhubInsight#recommendation_text}
	RecommendationText interface{} `json:"recommendationText" yaml:"recommendationText"`
	// record_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#record_state SecurityhubInsight#record_state}
	RecordState interface{} `json:"recordState" yaml:"recordState"`
	// related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#related_findings_id SecurityhubInsight#related_findings_id}
	RelatedFindingsId interface{} `json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#related_findings_product_arn SecurityhubInsight#related_findings_product_arn}
	RelatedFindingsProductArn interface{} `json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// resource_aws_ec2_instance_iam_instance_profile_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_iam_instance_profile_arn SecurityhubInsight#resource_aws_ec2_instance_iam_instance_profile_arn}
	ResourceAwsEc2InstanceIamInstanceProfileArn interface{} `json:"resourceAwsEc2InstanceIamInstanceProfileArn" yaml:"resourceAwsEc2InstanceIamInstanceProfileArn"`
	// resource_aws_ec2_instance_image_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_image_id SecurityhubInsight#resource_aws_ec2_instance_image_id}
	ResourceAwsEc2InstanceImageId interface{} `json:"resourceAwsEc2InstanceImageId" yaml:"resourceAwsEc2InstanceImageId"`
	// resource_aws_ec2_instance_ipv4_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_ipv4_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv4_addresses}
	ResourceAwsEc2InstanceIpv4Addresses interface{} `json:"resourceAwsEc2InstanceIpv4Addresses" yaml:"resourceAwsEc2InstanceIpv4Addresses"`
	// resource_aws_ec2_instance_ipv6_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_ipv6_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv6_addresses}
	ResourceAwsEc2InstanceIpv6Addresses interface{} `json:"resourceAwsEc2InstanceIpv6Addresses" yaml:"resourceAwsEc2InstanceIpv6Addresses"`
	// resource_aws_ec2_instance_key_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_key_name SecurityhubInsight#resource_aws_ec2_instance_key_name}
	ResourceAwsEc2InstanceKeyName interface{} `json:"resourceAwsEc2InstanceKeyName" yaml:"resourceAwsEc2InstanceKeyName"`
	// resource_aws_ec2_instance_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_launched_at SecurityhubInsight#resource_aws_ec2_instance_launched_at}
	ResourceAwsEc2InstanceLaunchedAt interface{} `json:"resourceAwsEc2InstanceLaunchedAt" yaml:"resourceAwsEc2InstanceLaunchedAt"`
	// resource_aws_ec2_instance_subnet_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_subnet_id SecurityhubInsight#resource_aws_ec2_instance_subnet_id}
	ResourceAwsEc2InstanceSubnetId interface{} `json:"resourceAwsEc2InstanceSubnetId" yaml:"resourceAwsEc2InstanceSubnetId"`
	// resource_aws_ec2_instance_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_type SecurityhubInsight#resource_aws_ec2_instance_type}
	ResourceAwsEc2InstanceType interface{} `json:"resourceAwsEc2InstanceType" yaml:"resourceAwsEc2InstanceType"`
	// resource_aws_ec2_instance_vpc_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_ec2_instance_vpc_id SecurityhubInsight#resource_aws_ec2_instance_vpc_id}
	ResourceAwsEc2InstanceVpcId interface{} `json:"resourceAwsEc2InstanceVpcId" yaml:"resourceAwsEc2InstanceVpcId"`
	// resource_aws_iam_access_key_created_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_iam_access_key_created_at SecurityhubInsight#resource_aws_iam_access_key_created_at}
	ResourceAwsIamAccessKeyCreatedAt interface{} `json:"resourceAwsIamAccessKeyCreatedAt" yaml:"resourceAwsIamAccessKeyCreatedAt"`
	// resource_aws_iam_access_key_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_iam_access_key_status SecurityhubInsight#resource_aws_iam_access_key_status}
	ResourceAwsIamAccessKeyStatus interface{} `json:"resourceAwsIamAccessKeyStatus" yaml:"resourceAwsIamAccessKeyStatus"`
	// resource_aws_iam_access_key_user_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_iam_access_key_user_name SecurityhubInsight#resource_aws_iam_access_key_user_name}
	ResourceAwsIamAccessKeyUserName interface{} `json:"resourceAwsIamAccessKeyUserName" yaml:"resourceAwsIamAccessKeyUserName"`
	// resource_aws_s3_bucket_owner_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_s3_bucket_owner_id SecurityhubInsight#resource_aws_s3_bucket_owner_id}
	ResourceAwsS3BucketOwnerId interface{} `json:"resourceAwsS3BucketOwnerId" yaml:"resourceAwsS3BucketOwnerId"`
	// resource_aws_s3_bucket_owner_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_aws_s3_bucket_owner_name SecurityhubInsight#resource_aws_s3_bucket_owner_name}
	ResourceAwsS3BucketOwnerName interface{} `json:"resourceAwsS3BucketOwnerName" yaml:"resourceAwsS3BucketOwnerName"`
	// resource_container_image_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_container_image_id SecurityhubInsight#resource_container_image_id}
	ResourceContainerImageId interface{} `json:"resourceContainerImageId" yaml:"resourceContainerImageId"`
	// resource_container_image_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_container_image_name SecurityhubInsight#resource_container_image_name}
	ResourceContainerImageName interface{} `json:"resourceContainerImageName" yaml:"resourceContainerImageName"`
	// resource_container_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_container_launched_at SecurityhubInsight#resource_container_launched_at}
	ResourceContainerLaunchedAt interface{} `json:"resourceContainerLaunchedAt" yaml:"resourceContainerLaunchedAt"`
	// resource_container_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_container_name SecurityhubInsight#resource_container_name}
	ResourceContainerName interface{} `json:"resourceContainerName" yaml:"resourceContainerName"`
	// resource_details_other block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_details_other SecurityhubInsight#resource_details_other}
	ResourceDetailsOther interface{} `json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_id SecurityhubInsight#resource_id}
	ResourceId interface{} `json:"resourceId" yaml:"resourceId"`
	// resource_partition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_partition SecurityhubInsight#resource_partition}
	ResourcePartition interface{} `json:"resourcePartition" yaml:"resourcePartition"`
	// resource_region block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_region SecurityhubInsight#resource_region}
	ResourceRegion interface{} `json:"resourceRegion" yaml:"resourceRegion"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_tags SecurityhubInsight#resource_tags}
	ResourceTags interface{} `json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#resource_type SecurityhubInsight#resource_type}
	ResourceType interface{} `json:"resourceType" yaml:"resourceType"`
	// severity_label block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#severity_label SecurityhubInsight#severity_label}
	SeverityLabel interface{} `json:"severityLabel" yaml:"severityLabel"`
	// source_url block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#source_url SecurityhubInsight#source_url}
	SourceUrl interface{} `json:"sourceUrl" yaml:"sourceUrl"`
	// threat_intel_indicator_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_category SecurityhubInsight#threat_intel_indicator_category}
	ThreatIntelIndicatorCategory interface{} `json:"threatIntelIndicatorCategory" yaml:"threatIntelIndicatorCategory"`
	// threat_intel_indicator_last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_last_observed_at SecurityhubInsight#threat_intel_indicator_last_observed_at}
	ThreatIntelIndicatorLastObservedAt interface{} `json:"threatIntelIndicatorLastObservedAt" yaml:"threatIntelIndicatorLastObservedAt"`
	// threat_intel_indicator_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_source SecurityhubInsight#threat_intel_indicator_source}
	ThreatIntelIndicatorSource interface{} `json:"threatIntelIndicatorSource" yaml:"threatIntelIndicatorSource"`
	// threat_intel_indicator_source_url block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_source_url SecurityhubInsight#threat_intel_indicator_source_url}
	ThreatIntelIndicatorSourceUrl interface{} `json:"threatIntelIndicatorSourceUrl" yaml:"threatIntelIndicatorSourceUrl"`
	// threat_intel_indicator_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_type SecurityhubInsight#threat_intel_indicator_type}
	ThreatIntelIndicatorType interface{} `json:"threatIntelIndicatorType" yaml:"threatIntelIndicatorType"`
	// threat_intel_indicator_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#threat_intel_indicator_value SecurityhubInsight#threat_intel_indicator_value}
	ThreatIntelIndicatorValue interface{} `json:"threatIntelIndicatorValue" yaml:"threatIntelIndicatorValue"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#title SecurityhubInsight#title}
	Title interface{} `json:"title" yaml:"title"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#type SecurityhubInsight#type}
	Type interface{} `json:"type" yaml:"type"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#updated_at SecurityhubInsight#updated_at}
	UpdatedAt interface{} `json:"updatedAt" yaml:"updatedAt"`
	// user_defined_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#user_defined_values SecurityhubInsight#user_defined_values}
	UserDefinedValues interface{} `json:"userDefinedValues" yaml:"userDefinedValues"`
	// verification_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#verification_state SecurityhubInsight#verification_state}
	VerificationState interface{} `json:"verificationState" yaml:"verificationState"`
	// workflow_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#workflow_status SecurityhubInsight#workflow_status}
	WorkflowStatus interface{} `json:"workflowStatus" yaml:"workflowStatus"`
}

type SecurityhubInsightFiltersAwsAccountId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersCompanyName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersComplianceStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersConfidence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersCreatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersCreatedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersCreatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersCreatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersCreatedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersCreatedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersCreatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersCreatedAtDateRange {
	var returns *SecurityhubInsightFiltersCreatedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersCreatedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersCreatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersCreatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersCreatedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersCreatedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersCriticality struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersDescription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsConfidence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersFindingProviderFieldsCriticality struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsTypes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFirstObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersFirstObservedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersFirstObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersFirstObservedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersFirstObservedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersFirstObservedAtDateRange {
	var returns *SecurityhubInsightFiltersFirstObservedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersFirstObservedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersGeneratorId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersKeyword struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersLastObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersLastObservedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersLastObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersLastObservedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersLastObservedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersLastObservedAtDateRange {
	var returns *SecurityhubInsightFiltersLastObservedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersLastObservedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersLastObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersLastObservedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersMalwareName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersMalwarePath struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersMalwareState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersMalwareType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkDestinationDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkDestinationIpv4 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersNetworkDestinationIpv6 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersNetworkDestinationPort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersNetworkDirection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkProtocol struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkSourceDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkSourceIpv4 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersNetworkSourceIpv6 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersNetworkSourceMac struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNetworkSourcePort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersNoteText struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNoteUpdatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersNoteUpdatedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersNoteUpdatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersNoteUpdatedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersNoteUpdatedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersNoteUpdatedAtDateRange {
	var returns *SecurityhubInsightFiltersNoteUpdatedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersNoteUpdatedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersNoteUpdatedBy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() interface{}
	SetAwsAccountId(val interface{})
	AwsAccountIdInput() interface{}
	CompanyName() interface{}
	SetCompanyName(val interface{})
	CompanyNameInput() interface{}
	ComplianceStatus() interface{}
	SetComplianceStatus(val interface{})
	ComplianceStatusInput() interface{}
	Confidence() interface{}
	SetConfidence(val interface{})
	ConfidenceInput() interface{}
	CreatedAt() interface{}
	SetCreatedAt(val interface{})
	CreatedAtInput() interface{}
	Criticality() interface{}
	SetCriticality(val interface{})
	CriticalityInput() interface{}
	Description() interface{}
	SetDescription(val interface{})
	DescriptionInput() interface{}
	FindingProviderFieldsConfidence() interface{}
	SetFindingProviderFieldsConfidence(val interface{})
	FindingProviderFieldsConfidenceInput() interface{}
	FindingProviderFieldsCriticality() interface{}
	SetFindingProviderFieldsCriticality(val interface{})
	FindingProviderFieldsCriticalityInput() interface{}
	FindingProviderFieldsRelatedFindingsId() interface{}
	SetFindingProviderFieldsRelatedFindingsId(val interface{})
	FindingProviderFieldsRelatedFindingsIdInput() interface{}
	FindingProviderFieldsRelatedFindingsProductArn() interface{}
	SetFindingProviderFieldsRelatedFindingsProductArn(val interface{})
	FindingProviderFieldsRelatedFindingsProductArnInput() interface{}
	FindingProviderFieldsSeverityLabel() interface{}
	SetFindingProviderFieldsSeverityLabel(val interface{})
	FindingProviderFieldsSeverityLabelInput() interface{}
	FindingProviderFieldsSeverityOriginal() interface{}
	SetFindingProviderFieldsSeverityOriginal(val interface{})
	FindingProviderFieldsSeverityOriginalInput() interface{}
	FindingProviderFieldsTypes() interface{}
	SetFindingProviderFieldsTypes(val interface{})
	FindingProviderFieldsTypesInput() interface{}
	FirstObservedAt() interface{}
	SetFirstObservedAt(val interface{})
	FirstObservedAtInput() interface{}
	GeneratorId() interface{}
	SetGeneratorId(val interface{})
	GeneratorIdInput() interface{}
	Id() interface{}
	SetId(val interface{})
	IdInput() interface{}
	InternalValue() *SecurityhubInsightFilters
	SetInternalValue(val *SecurityhubInsightFilters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Keyword() interface{}
	SetKeyword(val interface{})
	KeywordInput() interface{}
	LastObservedAt() interface{}
	SetLastObservedAt(val interface{})
	LastObservedAtInput() interface{}
	MalwareName() interface{}
	SetMalwareName(val interface{})
	MalwareNameInput() interface{}
	MalwarePath() interface{}
	SetMalwarePath(val interface{})
	MalwarePathInput() interface{}
	MalwareState() interface{}
	SetMalwareState(val interface{})
	MalwareStateInput() interface{}
	MalwareType() interface{}
	SetMalwareType(val interface{})
	MalwareTypeInput() interface{}
	NetworkDestinationDomain() interface{}
	SetNetworkDestinationDomain(val interface{})
	NetworkDestinationDomainInput() interface{}
	NetworkDestinationIpv4() interface{}
	SetNetworkDestinationIpv4(val interface{})
	NetworkDestinationIpv4Input() interface{}
	NetworkDestinationIpv6() interface{}
	SetNetworkDestinationIpv6(val interface{})
	NetworkDestinationIpv6Input() interface{}
	NetworkDestinationPort() interface{}
	SetNetworkDestinationPort(val interface{})
	NetworkDestinationPortInput() interface{}
	NetworkDirection() interface{}
	SetNetworkDirection(val interface{})
	NetworkDirectionInput() interface{}
	NetworkProtocol() interface{}
	SetNetworkProtocol(val interface{})
	NetworkProtocolInput() interface{}
	NetworkSourceDomain() interface{}
	SetNetworkSourceDomain(val interface{})
	NetworkSourceDomainInput() interface{}
	NetworkSourceIpv4() interface{}
	SetNetworkSourceIpv4(val interface{})
	NetworkSourceIpv4Input() interface{}
	NetworkSourceIpv6() interface{}
	SetNetworkSourceIpv6(val interface{})
	NetworkSourceIpv6Input() interface{}
	NetworkSourceMac() interface{}
	SetNetworkSourceMac(val interface{})
	NetworkSourceMacInput() interface{}
	NetworkSourcePort() interface{}
	SetNetworkSourcePort(val interface{})
	NetworkSourcePortInput() interface{}
	NoteText() interface{}
	SetNoteText(val interface{})
	NoteTextInput() interface{}
	NoteUpdatedAt() interface{}
	SetNoteUpdatedAt(val interface{})
	NoteUpdatedAtInput() interface{}
	NoteUpdatedBy() interface{}
	SetNoteUpdatedBy(val interface{})
	NoteUpdatedByInput() interface{}
	ProcessLaunchedAt() interface{}
	SetProcessLaunchedAt(val interface{})
	ProcessLaunchedAtInput() interface{}
	ProcessName() interface{}
	SetProcessName(val interface{})
	ProcessNameInput() interface{}
	ProcessParentPid() interface{}
	SetProcessParentPid(val interface{})
	ProcessParentPidInput() interface{}
	ProcessPath() interface{}
	SetProcessPath(val interface{})
	ProcessPathInput() interface{}
	ProcessPid() interface{}
	SetProcessPid(val interface{})
	ProcessPidInput() interface{}
	ProcessTerminatedAt() interface{}
	SetProcessTerminatedAt(val interface{})
	ProcessTerminatedAtInput() interface{}
	ProductArn() interface{}
	SetProductArn(val interface{})
	ProductArnInput() interface{}
	ProductFields() interface{}
	SetProductFields(val interface{})
	ProductFieldsInput() interface{}
	ProductName() interface{}
	SetProductName(val interface{})
	ProductNameInput() interface{}
	RecommendationText() interface{}
	SetRecommendationText(val interface{})
	RecommendationTextInput() interface{}
	RecordState() interface{}
	SetRecordState(val interface{})
	RecordStateInput() interface{}
	RelatedFindingsId() interface{}
	SetRelatedFindingsId(val interface{})
	RelatedFindingsIdInput() interface{}
	RelatedFindingsProductArn() interface{}
	SetRelatedFindingsProductArn(val interface{})
	RelatedFindingsProductArnInput() interface{}
	ResourceAwsEc2InstanceIamInstanceProfileArn() interface{}
	SetResourceAwsEc2InstanceIamInstanceProfileArn(val interface{})
	ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{}
	ResourceAwsEc2InstanceImageId() interface{}
	SetResourceAwsEc2InstanceImageId(val interface{})
	ResourceAwsEc2InstanceImageIdInput() interface{}
	ResourceAwsEc2InstanceIpv4Addresses() interface{}
	SetResourceAwsEc2InstanceIpv4Addresses(val interface{})
	ResourceAwsEc2InstanceIpv4AddressesInput() interface{}
	ResourceAwsEc2InstanceIpv6Addresses() interface{}
	SetResourceAwsEc2InstanceIpv6Addresses(val interface{})
	ResourceAwsEc2InstanceIpv6AddressesInput() interface{}
	ResourceAwsEc2InstanceKeyName() interface{}
	SetResourceAwsEc2InstanceKeyName(val interface{})
	ResourceAwsEc2InstanceKeyNameInput() interface{}
	ResourceAwsEc2InstanceLaunchedAt() interface{}
	SetResourceAwsEc2InstanceLaunchedAt(val interface{})
	ResourceAwsEc2InstanceLaunchedAtInput() interface{}
	ResourceAwsEc2InstanceSubnetId() interface{}
	SetResourceAwsEc2InstanceSubnetId(val interface{})
	ResourceAwsEc2InstanceSubnetIdInput() interface{}
	ResourceAwsEc2InstanceType() interface{}
	SetResourceAwsEc2InstanceType(val interface{})
	ResourceAwsEc2InstanceTypeInput() interface{}
	ResourceAwsEc2InstanceVpcId() interface{}
	SetResourceAwsEc2InstanceVpcId(val interface{})
	ResourceAwsEc2InstanceVpcIdInput() interface{}
	ResourceAwsIamAccessKeyCreatedAt() interface{}
	SetResourceAwsIamAccessKeyCreatedAt(val interface{})
	ResourceAwsIamAccessKeyCreatedAtInput() interface{}
	ResourceAwsIamAccessKeyStatus() interface{}
	SetResourceAwsIamAccessKeyStatus(val interface{})
	ResourceAwsIamAccessKeyStatusInput() interface{}
	ResourceAwsIamAccessKeyUserName() interface{}
	SetResourceAwsIamAccessKeyUserName(val interface{})
	ResourceAwsIamAccessKeyUserNameInput() interface{}
	ResourceAwsS3BucketOwnerId() interface{}
	SetResourceAwsS3BucketOwnerId(val interface{})
	ResourceAwsS3BucketOwnerIdInput() interface{}
	ResourceAwsS3BucketOwnerName() interface{}
	SetResourceAwsS3BucketOwnerName(val interface{})
	ResourceAwsS3BucketOwnerNameInput() interface{}
	ResourceContainerImageId() interface{}
	SetResourceContainerImageId(val interface{})
	ResourceContainerImageIdInput() interface{}
	ResourceContainerImageName() interface{}
	SetResourceContainerImageName(val interface{})
	ResourceContainerImageNameInput() interface{}
	ResourceContainerLaunchedAt() interface{}
	SetResourceContainerLaunchedAt(val interface{})
	ResourceContainerLaunchedAtInput() interface{}
	ResourceContainerName() interface{}
	SetResourceContainerName(val interface{})
	ResourceContainerNameInput() interface{}
	ResourceDetailsOther() interface{}
	SetResourceDetailsOther(val interface{})
	ResourceDetailsOtherInput() interface{}
	ResourceId() interface{}
	SetResourceId(val interface{})
	ResourceIdInput() interface{}
	ResourcePartition() interface{}
	SetResourcePartition(val interface{})
	ResourcePartitionInput() interface{}
	ResourceRegion() interface{}
	SetResourceRegion(val interface{})
	ResourceRegionInput() interface{}
	ResourceTags() interface{}
	SetResourceTags(val interface{})
	ResourceTagsInput() interface{}
	ResourceType() interface{}
	SetResourceType(val interface{})
	ResourceTypeInput() interface{}
	SeverityLabel() interface{}
	SetSeverityLabel(val interface{})
	SeverityLabelInput() interface{}
	SourceUrl() interface{}
	SetSourceUrl(val interface{})
	SourceUrlInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreatIntelIndicatorCategory() interface{}
	SetThreatIntelIndicatorCategory(val interface{})
	ThreatIntelIndicatorCategoryInput() interface{}
	ThreatIntelIndicatorLastObservedAt() interface{}
	SetThreatIntelIndicatorLastObservedAt(val interface{})
	ThreatIntelIndicatorLastObservedAtInput() interface{}
	ThreatIntelIndicatorSource() interface{}
	SetThreatIntelIndicatorSource(val interface{})
	ThreatIntelIndicatorSourceInput() interface{}
	ThreatIntelIndicatorSourceUrl() interface{}
	SetThreatIntelIndicatorSourceUrl(val interface{})
	ThreatIntelIndicatorSourceUrlInput() interface{}
	ThreatIntelIndicatorType() interface{}
	SetThreatIntelIndicatorType(val interface{})
	ThreatIntelIndicatorTypeInput() interface{}
	ThreatIntelIndicatorValue() interface{}
	SetThreatIntelIndicatorValue(val interface{})
	ThreatIntelIndicatorValueInput() interface{}
	Title() interface{}
	SetTitle(val interface{})
	TitleInput() interface{}
	Type() interface{}
	SetType(val interface{})
	TypeInput() interface{}
	UpdatedAt() interface{}
	SetUpdatedAt(val interface{})
	UpdatedAtInput() interface{}
	UserDefinedValues() interface{}
	SetUserDefinedValues(val interface{})
	UserDefinedValuesInput() interface{}
	VerificationState() interface{}
	SetVerificationState(val interface{})
	VerificationStateInput() interface{}
	WorkflowStatus() interface{}
	SetWorkflowStatus(val interface{})
	WorkflowStatusInput() interface{}
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
	ResetAwsAccountId()
	ResetCompanyName()
	ResetComplianceStatus()
	ResetConfidence()
	ResetCreatedAt()
	ResetCriticality()
	ResetDescription()
	ResetFindingProviderFieldsConfidence()
	ResetFindingProviderFieldsCriticality()
	ResetFindingProviderFieldsRelatedFindingsId()
	ResetFindingProviderFieldsRelatedFindingsProductArn()
	ResetFindingProviderFieldsSeverityLabel()
	ResetFindingProviderFieldsSeverityOriginal()
	ResetFindingProviderFieldsTypes()
	ResetFirstObservedAt()
	ResetGeneratorId()
	ResetId()
	ResetKeyword()
	ResetLastObservedAt()
	ResetMalwareName()
	ResetMalwarePath()
	ResetMalwareState()
	ResetMalwareType()
	ResetNetworkDestinationDomain()
	ResetNetworkDestinationIpv4()
	ResetNetworkDestinationIpv6()
	ResetNetworkDestinationPort()
	ResetNetworkDirection()
	ResetNetworkProtocol()
	ResetNetworkSourceDomain()
	ResetNetworkSourceIpv4()
	ResetNetworkSourceIpv6()
	ResetNetworkSourceMac()
	ResetNetworkSourcePort()
	ResetNoteText()
	ResetNoteUpdatedAt()
	ResetNoteUpdatedBy()
	ResetProcessLaunchedAt()
	ResetProcessName()
	ResetProcessParentPid()
	ResetProcessPath()
	ResetProcessPid()
	ResetProcessTerminatedAt()
	ResetProductArn()
	ResetProductFields()
	ResetProductName()
	ResetRecommendationText()
	ResetRecordState()
	ResetRelatedFindingsId()
	ResetRelatedFindingsProductArn()
	ResetResourceAwsEc2InstanceIamInstanceProfileArn()
	ResetResourceAwsEc2InstanceImageId()
	ResetResourceAwsEc2InstanceIpv4Addresses()
	ResetResourceAwsEc2InstanceIpv6Addresses()
	ResetResourceAwsEc2InstanceKeyName()
	ResetResourceAwsEc2InstanceLaunchedAt()
	ResetResourceAwsEc2InstanceSubnetId()
	ResetResourceAwsEc2InstanceType()
	ResetResourceAwsEc2InstanceVpcId()
	ResetResourceAwsIamAccessKeyCreatedAt()
	ResetResourceAwsIamAccessKeyStatus()
	ResetResourceAwsIamAccessKeyUserName()
	ResetResourceAwsS3BucketOwnerId()
	ResetResourceAwsS3BucketOwnerName()
	ResetResourceContainerImageId()
	ResetResourceContainerImageName()
	ResetResourceContainerLaunchedAt()
	ResetResourceContainerName()
	ResetResourceDetailsOther()
	ResetResourceId()
	ResetResourcePartition()
	ResetResourceRegion()
	ResetResourceTags()
	ResetResourceType()
	ResetSeverityLabel()
	ResetSourceUrl()
	ResetThreatIntelIndicatorCategory()
	ResetThreatIntelIndicatorLastObservedAt()
	ResetThreatIntelIndicatorSource()
	ResetThreatIntelIndicatorSourceUrl()
	ResetThreatIntelIndicatorType()
	ResetThreatIntelIndicatorValue()
	ResetTitle()
	ResetType()
	ResetUpdatedAt()
	ResetUserDefinedValues()
	ResetVerificationState()
	ResetWorkflowStatus()
}

// The jsii proxy struct for SecurityhubInsightFiltersOutputReference
type jsiiProxy_SecurityhubInsightFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Confidence() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Criticality() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Description() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidence() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticality() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Id() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) InternalValue() *SecurityhubInsightFilters {
	var returns *SecurityhubInsightFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Keyword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) KeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwarePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedBy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processParentPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processTerminatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4Addresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6Addresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOther() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Title() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Type() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersOutputReference_Override(s SecurityhubInsightFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetAwsAccountId(val interface{}) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCompanyName(val interface{}) {
	_jsii_.Set(
		j,
		"companyName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetComplianceStatus(val interface{}) {
	_jsii_.Set(
		j,
		"complianceStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetConfidence(val interface{}) {
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCreatedAt(val interface{}) {
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCriticality(val interface{}) {
	_jsii_.Set(
		j,
		"criticality",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetDescription(val interface{}) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsConfidence(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsConfidence",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsCriticality(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsCriticality",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsRelatedFindingsId(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsRelatedFindingsId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsRelatedFindingsProductArn(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsSeverityLabel(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsSeverityLabel",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsSeverityOriginal(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsSeverityOriginal",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsTypes(val interface{}) {
	_jsii_.Set(
		j,
		"findingProviderFieldsTypes",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFirstObservedAt(val interface{}) {
	_jsii_.Set(
		j,
		"firstObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetGeneratorId(val interface{}) {
	_jsii_.Set(
		j,
		"generatorId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetId(val interface{}) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetInternalValue(val *SecurityhubInsightFilters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetKeyword(val interface{}) {
	_jsii_.Set(
		j,
		"keyword",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetLastObservedAt(val interface{}) {
	_jsii_.Set(
		j,
		"lastObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareName(val interface{}) {
	_jsii_.Set(
		j,
		"malwareName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwarePath(val interface{}) {
	_jsii_.Set(
		j,
		"malwarePath",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareState(val interface{}) {
	_jsii_.Set(
		j,
		"malwareState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareType(val interface{}) {
	_jsii_.Set(
		j,
		"malwareType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationDomain(val interface{}) {
	_jsii_.Set(
		j,
		"networkDestinationDomain",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationIpv4(val interface{}) {
	_jsii_.Set(
		j,
		"networkDestinationIpv4",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationIpv6(val interface{}) {
	_jsii_.Set(
		j,
		"networkDestinationIpv6",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationPort(val interface{}) {
	_jsii_.Set(
		j,
		"networkDestinationPort",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDirection(val interface{}) {
	_jsii_.Set(
		j,
		"networkDirection",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkProtocol(val interface{}) {
	_jsii_.Set(
		j,
		"networkProtocol",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceDomain(val interface{}) {
	_jsii_.Set(
		j,
		"networkSourceDomain",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceIpv4(val interface{}) {
	_jsii_.Set(
		j,
		"networkSourceIpv4",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceIpv6(val interface{}) {
	_jsii_.Set(
		j,
		"networkSourceIpv6",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceMac(val interface{}) {
	_jsii_.Set(
		j,
		"networkSourceMac",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourcePort(val interface{}) {
	_jsii_.Set(
		j,
		"networkSourcePort",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteText(val interface{}) {
	_jsii_.Set(
		j,
		"noteText",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteUpdatedAt(val interface{}) {
	_jsii_.Set(
		j,
		"noteUpdatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteUpdatedBy(val interface{}) {
	_jsii_.Set(
		j,
		"noteUpdatedBy",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessLaunchedAt(val interface{}) {
	_jsii_.Set(
		j,
		"processLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessName(val interface{}) {
	_jsii_.Set(
		j,
		"processName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessParentPid(val interface{}) {
	_jsii_.Set(
		j,
		"processParentPid",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessPath(val interface{}) {
	_jsii_.Set(
		j,
		"processPath",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessPid(val interface{}) {
	_jsii_.Set(
		j,
		"processPid",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessTerminatedAt(val interface{}) {
	_jsii_.Set(
		j,
		"processTerminatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductArn(val interface{}) {
	_jsii_.Set(
		j,
		"productArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductFields(val interface{}) {
	_jsii_.Set(
		j,
		"productFields",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductName(val interface{}) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRecommendationText(val interface{}) {
	_jsii_.Set(
		j,
		"recommendationText",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRecordState(val interface{}) {
	_jsii_.Set(
		j,
		"recordState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRelatedFindingsId(val interface{}) {
	_jsii_.Set(
		j,
		"relatedFindingsId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRelatedFindingsProductArn(val interface{}) {
	_jsii_.Set(
		j,
		"relatedFindingsProductArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIamInstanceProfileArn(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceImageId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceImageId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIpv4Addresses(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIpv6Addresses(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceKeyName(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceKeyName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceLaunchedAt(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceSubnetId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceSubnetId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceType(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceVpcId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceVpcId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyCreatedAt(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyStatus(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyUserName(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyUserName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsS3BucketOwnerId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsS3BucketOwnerId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsS3BucketOwnerName(val interface{}) {
	_jsii_.Set(
		j,
		"resourceAwsS3BucketOwnerName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerImageId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceContainerImageId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerImageName(val interface{}) {
	_jsii_.Set(
		j,
		"resourceContainerImageName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerLaunchedAt(val interface{}) {
	_jsii_.Set(
		j,
		"resourceContainerLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerName(val interface{}) {
	_jsii_.Set(
		j,
		"resourceContainerName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceDetailsOther(val interface{}) {
	_jsii_.Set(
		j,
		"resourceDetailsOther",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceId(val interface{}) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourcePartition(val interface{}) {
	_jsii_.Set(
		j,
		"resourcePartition",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceRegion(val interface{}) {
	_jsii_.Set(
		j,
		"resourceRegion",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceType(val interface{}) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetSeverityLabel(val interface{}) {
	_jsii_.Set(
		j,
		"severityLabel",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetSourceUrl(val interface{}) {
	_jsii_.Set(
		j,
		"sourceUrl",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorCategory(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorCategory",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorLastObservedAt(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorLastObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorSource(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorSource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorSourceUrl(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorSourceUrl",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorType(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorValue(val interface{}) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetTitle(val interface{}) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetType(val interface{}) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetUpdatedAt(val interface{}) {
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetUserDefinedValues(val interface{}) {
	_jsii_.Set(
		j,
		"userDefinedValues",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetVerificationState(val interface{}) {
	_jsii_.Set(
		j,
		"verificationState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetWorkflowStatus(val interface{}) {
	_jsii_.Set(
		j,
		"workflowStatus",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		s,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsCriticality() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsCriticality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsRelatedFindingsId() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsRelatedFindingsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsSeverityLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsSeverityLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsSeverityOriginal() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsSeverityOriginal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFindingProviderFieldsTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetFindingProviderFieldsTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		s,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetKeyword() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareName() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwarePath() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwarePath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareState() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetMalwareType() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationIpv4() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationIpv4",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationIpv6() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationIpv6",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDestinationPort() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDestinationPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkDirection() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDirection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceIpv4() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceIpv4",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceIpv6() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceIpv6",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourceMac() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourceMac",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNetworkSourcePort() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSourcePort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessName() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessParentPid() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessParentPid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessPath() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessPid() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessPid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProcessTerminatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessTerminatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductFields() {
	_jsii_.InvokeVoid(
		s,
		"resetProductFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRecommendationText() {
	_jsii_.InvokeVoid(
		s,
		"resetRecommendationText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyCreatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyCreatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsIamAccessKeyUserName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsIamAccessKeyUserName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsS3BucketOwnerId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsS3BucketOwnerId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceAwsS3BucketOwnerName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceAwsS3BucketOwnerName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerImageName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerImageName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerLaunchedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerLaunchedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceContainerName() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceContainerName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		s,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorCategory() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorCategory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorLastObservedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorLastObservedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorSource() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorSourceUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorSourceUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorType() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetThreatIntelIndicatorValue() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatIntelIndicatorValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetUserDefinedValues() {
	_jsii_.InvokeVoid(
		s,
		"resetUserDefinedValues",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		s,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

type SecurityhubInsightFiltersProcessLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersProcessLaunchedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersProcessLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersProcessLaunchedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersProcessLaunchedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersProcessLaunchedAtDateRange {
	var returns *SecurityhubInsightFiltersProcessLaunchedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersProcessLaunchedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersProcessName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProcessParentPid struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersProcessPath struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProcessPid struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte" yaml:"lte"`
}

type SecurityhubInsightFiltersProcessTerminatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersProcessTerminatedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersProcessTerminatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersProcessTerminatedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersProcessTerminatedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersProcessTerminatedAtDateRange {
	var returns *SecurityhubInsightFiltersProcessTerminatedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersProcessTerminatedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersProductArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProductFields struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#key SecurityhubInsight#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersProductName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersRecommendationText struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersRecordState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersRelatedFindingsId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersRelatedFindingsProductArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceImageId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange {
	var returns *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange {
	var returns *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsS3BucketOwnerId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceAwsS3BucketOwnerName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceContainerImageId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceContainerImageName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange {
	var returns *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersResourceContainerName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceDetailsOther struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#key SecurityhubInsight#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourcePartition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceRegion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#key SecurityhubInsight#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersResourceType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersSeverityLabel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersSourceUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorCategory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange {
	var returns *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersThreatIntelIndicatorSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersTitle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersUpdatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersUpdatedAtDateRange `json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `json:"start" yaml:"start"`
}

type SecurityhubInsightFiltersUpdatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SecurityhubInsightFiltersUpdatedAtDateRange
	SetInternalValue(val *SecurityhubInsightFiltersUpdatedAtDateRange)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) InternalValue() *SecurityhubInsightFiltersUpdatedAtDateRange {
	var returns *SecurityhubInsightFiltersUpdatedAtDateRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSecurityhubInsightFiltersUpdatedAtDateRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersUpdatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetInternalValue(val *SecurityhubInsightFiltersUpdatedAtDateRange) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SecurityhubInsightFiltersUserDefinedValues struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#key SecurityhubInsight#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersVerificationState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

type SecurityhubInsightFiltersWorkflowStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `json:"value" yaml:"value"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter aws_securityhub_invite_accepter}.
type SecurityhubInviteAccepter interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InvitationId() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterId() *string
	SetMasterId(val *string)
	MasterIdInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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

// The jsii proxy struct for SecurityhubInviteAccepter
type jsiiProxy_SecurityhubInviteAccepter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubInviteAccepter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) InvitationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) MasterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) MasterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInviteAccepter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter aws_securityhub_invite_accepter} Resource.
func NewSecurityhubInviteAccepter(scope constructs.Construct, id *string, config *SecurityhubInviteAccepterConfig) SecurityhubInviteAccepter {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInviteAccepter{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter aws_securityhub_invite_accepter} Resource.
func NewSecurityhubInviteAccepter_Override(s SecurityhubInviteAccepter, scope constructs.Construct, id *string, config *SecurityhubInviteAccepterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepter",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetMasterId(val *string) {
	_jsii_.Set(
		j,
		"masterId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubInviteAccepter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubInviteAccepter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepter",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubInviteAccepter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInviteAccepter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubInviteAccepter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInviteAccepter) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) ToString() *string {
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
func (s *jsiiProxy_SecurityhubInviteAccepter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubInviteAccepterConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter#master_id SecurityhubInviteAccepter#master_id}.
	MasterId *string `json:"masterId" yaml:"masterId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member aws_securityhub_member}.
type SecurityhubMember interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Invite() interface{}
	SetInvite(val interface{})
	InviteInput() interface{}
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterId() *string
	MemberStatus() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetInvite()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurityhubMember
type jsiiProxy_SecurityhubMember struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubMember) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Invite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) InviteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inviteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) MasterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) MemberStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubMember) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member aws_securityhub_member} Resource.
func NewSecurityhubMember(scope constructs.Construct, id *string, config *SecurityhubMemberConfig) SecurityhubMember {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubMember{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubMember",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member aws_securityhub_member} Resource.
func NewSecurityhubMember_Override(s SecurityhubMember, scope constructs.Construct, id *string, config *SecurityhubMemberConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubMember",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetInvite(val interface{}) {
	_jsii_.Set(
		j,
		"invite",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubMember) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubMember_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubMember",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubMember_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubMember",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubMember) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubMember) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubMember) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubMember) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubMember) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubMember) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubMember) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubMember) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubMember) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubMember) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubMember) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubMember) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurityhubMember) ResetInvite() {
	_jsii_.InvokeVoid(
		s,
		"resetInvite",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubMember) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubMember) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubMember) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubMember) ToString() *string {
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
func (s *jsiiProxy_SecurityhubMember) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubMemberConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member#account_id SecurityhubMember#account_id}.
	AccountId *string `json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member#email SecurityhubMember#email}.
	Email *string `json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member#invite SecurityhubMember#invite}.
	Invite interface{} `json:"invite" yaml:"invite"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account aws_securityhub_organization_admin_account}.
type SecurityhubOrganizationAdminAccount interface {
	cdktf.TerraformResource
	AdminAccountId() *string
	SetAdminAccountId(val *string)
	AdminAccountIdInput() *string
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

// The jsii proxy struct for SecurityhubOrganizationAdminAccount
type jsiiProxy_SecurityhubOrganizationAdminAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) AdminAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) AdminAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account aws_securityhub_organization_admin_account} Resource.
func NewSecurityhubOrganizationAdminAccount(scope constructs.Construct, id *string, config *SecurityhubOrganizationAdminAccountConfig) SecurityhubOrganizationAdminAccount {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubOrganizationAdminAccount{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account aws_securityhub_organization_admin_account} Resource.
func NewSecurityhubOrganizationAdminAccount_Override(s SecurityhubOrganizationAdminAccount, scope constructs.Construct, id *string, config *SecurityhubOrganizationAdminAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccount",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetAdminAccountId(val *string) {
	_jsii_.Set(
		j,
		"adminAccountId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubOrganizationAdminAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubOrganizationAdminAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) ToString() *string {
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
func (s *jsiiProxy_SecurityhubOrganizationAdminAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubOrganizationAdminAccountConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account#admin_account_id SecurityhubOrganizationAdminAccount#admin_account_id}.
	AdminAccountId *string `json:"adminAccountId" yaml:"adminAccountId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration aws_securityhub_organization_configuration}.
type SecurityhubOrganizationConfiguration interface {
	cdktf.TerraformResource
	AutoEnable() interface{}
	SetAutoEnable(val interface{})
	AutoEnableInput() interface{}
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

// The jsii proxy struct for SecurityhubOrganizationConfiguration
type jsiiProxy_SecurityhubOrganizationConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) AutoEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) AutoEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration aws_securityhub_organization_configuration} Resource.
func NewSecurityhubOrganizationConfiguration(scope constructs.Construct, id *string, config *SecurityhubOrganizationConfigurationConfig) SecurityhubOrganizationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubOrganizationConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration aws_securityhub_organization_configuration} Resource.
func NewSecurityhubOrganizationConfiguration_Override(s SecurityhubOrganizationConfiguration, scope constructs.Construct, id *string, config *SecurityhubOrganizationConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfiguration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetAutoEnable(val interface{}) {
	_jsii_.Set(
		j,
		"autoEnable",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubOrganizationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubOrganizationConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubOrganizationConfiguration) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) ToString() *string {
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
func (s *jsiiProxy_SecurityhubOrganizationConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubOrganizationConfigurationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration#auto_enable SecurityhubOrganizationConfiguration#auto_enable}.
	AutoEnable interface{} `json:"autoEnable" yaml:"autoEnable"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription aws_securityhub_product_subscription}.
type SecurityhubProductSubscription interface {
	cdktf.TerraformResource
	Arn() *string
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
	ProductArn() *string
	SetProductArn(val *string)
	ProductArnInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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

// The jsii proxy struct for SecurityhubProductSubscription
type jsiiProxy_SecurityhubProductSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubProductSubscription) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) ProductArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) ProductArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubProductSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription aws_securityhub_product_subscription} Resource.
func NewSecurityhubProductSubscription(scope constructs.Construct, id *string, config *SecurityhubProductSubscriptionConfig) SecurityhubProductSubscription {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubProductSubscription{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubProductSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription aws_securityhub_product_subscription} Resource.
func NewSecurityhubProductSubscription_Override(s SecurityhubProductSubscription, scope constructs.Construct, id *string, config *SecurityhubProductSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubProductSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetProductArn(val *string) {
	_jsii_.Set(
		j,
		"productArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetProvider(val cdktf.TerraformProvider) {
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
func SecurityhubProductSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubProductSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubProductSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubProductSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubProductSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubProductSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubProductSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubProductSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubProductSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubProductSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubProductSubscription) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubProductSubscription) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubProductSubscription) ToString() *string {
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
func (s *jsiiProxy_SecurityhubProductSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubProductSubscriptionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription#product_arn SecurityhubProductSubscription#product_arn}.
	ProductArn *string `json:"productArn" yaml:"productArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control aws_securityhub_standards_control}.
type SecurityhubStandardsControl interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ControlId() *string
	ControlStatus() *string
	SetControlStatus(val *string)
	ControlStatusInput() *string
	ControlStatusUpdatedAt() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	DisabledReason() *string
	SetDisabledReason(val *string)
	DisabledReasonInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RelatedRequirements() *[]*string
	RemediationUrl() *string
	SeverityRating() *string
	StandardsControlArn() *string
	SetStandardsControlArn(val *string)
	StandardsControlArnInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Title() *string
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
	ResetDisabledReason()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurityhubStandardsControl
type jsiiProxy_SecurityhubStandardsControl struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubStandardsControl) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) ControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) ControlStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) ControlStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) ControlStatusUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlStatusUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) DisabledReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) DisabledReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) RelatedRequirements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relatedRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) RemediationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remediationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) SeverityRating() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityRating",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) StandardsControlArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardsControlArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) StandardsControlArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardsControlArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsControl) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control aws_securityhub_standards_control} Resource.
func NewSecurityhubStandardsControl(scope constructs.Construct, id *string, config *SecurityhubStandardsControlConfig) SecurityhubStandardsControl {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubStandardsControl{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubStandardsControl",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control aws_securityhub_standards_control} Resource.
func NewSecurityhubStandardsControl_Override(s SecurityhubStandardsControl, scope constructs.Construct, id *string, config *SecurityhubStandardsControlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubStandardsControl",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetControlStatus(val *string) {
	_jsii_.Set(
		j,
		"controlStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetDisabledReason(val *string) {
	_jsii_.Set(
		j,
		"disabledReason",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsControl) SetStandardsControlArn(val *string) {
	_jsii_.Set(
		j,
		"standardsControlArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SecurityhubStandardsControl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubStandardsControl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubStandardsControl_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubStandardsControl",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubStandardsControl) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubStandardsControl) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubStandardsControl) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubStandardsControl) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubStandardsControl) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurityhubStandardsControl) ResetDisabledReason() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabledReason",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubStandardsControl) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubStandardsControl) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsControl) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsControl) ToString() *string {
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
func (s *jsiiProxy_SecurityhubStandardsControl) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubStandardsControlConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control#control_status SecurityhubStandardsControl#control_status}.
	ControlStatus *string `json:"controlStatus" yaml:"controlStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control#standards_control_arn SecurityhubStandardsControl#standards_control_arn}.
	StandardsControlArn *string `json:"standardsControlArn" yaml:"standardsControlArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control#disabled_reason SecurityhubStandardsControl#disabled_reason}.
	DisabledReason *string `json:"disabledReason" yaml:"disabledReason"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription aws_securityhub_standards_subscription}.
type SecurityhubStandardsSubscription interface {
	cdktf.TerraformResource
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
	StandardsArn() *string
	SetStandardsArn(val *string)
	StandardsArnInput() *string
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

// The jsii proxy struct for SecurityhubStandardsSubscription
type jsiiProxy_SecurityhubStandardsSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) StandardsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) StandardsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription aws_securityhub_standards_subscription} Resource.
func NewSecurityhubStandardsSubscription(scope constructs.Construct, id *string, config *SecurityhubStandardsSubscriptionConfig) SecurityhubStandardsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubStandardsSubscription{}

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription aws_securityhub_standards_subscription} Resource.
func NewSecurityhubStandardsSubscription_Override(s SecurityhubStandardsSubscription, scope constructs.Construct, id *string, config *SecurityhubStandardsSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetStandardsArn(val *string) {
	_jsii_.Set(
		j,
		"standardsArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SecurityhubStandardsSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityhubStandardsSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SecurityhubStandardsSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SecurityhubStandardsSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubStandardsSubscription) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) ToMetadata() interface{} {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) ToString() *string {
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
func (s *jsiiProxy_SecurityhubStandardsSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Security Hub.
type SecurityhubStandardsSubscriptionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription#standards_arn SecurityhubStandardsSubscription#standards_arn}.
	StandardsArn *string `json:"standardsArn" yaml:"standardsArn"`
}
