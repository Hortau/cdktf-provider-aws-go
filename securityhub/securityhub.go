package securityhub

import (
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hortau/cdktf-provider-aws-go/securityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account.html aws_securityhub_account}.
type SecurityhubAccount interface {
	cdktf.TerraformResource
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

func (j *jsiiProxy_SecurityhubAccount) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account.html aws_securityhub_account} Resource.
func NewSecurityhubAccount(scope constructs.Construct, id *string, config *SecurityhubAccountConfig) SecurityhubAccount {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubAccount{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_account.html aws_securityhub_account} Resource.
func NewSecurityhubAccount_Override(s SecurityhubAccount, scope constructs.Construct, id *string, config *SecurityhubAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubAccount",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubAccount) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubAccount",
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
		"hashicorp_aws.SecurityHub.SecurityhubAccount",
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

type SecurityhubAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html aws_securityhub_action_target}.
type SecurityhubActionTarget interface {
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

func (j *jsiiProxy_SecurityhubActionTarget) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html aws_securityhub_action_target} Resource.
func NewSecurityhubActionTarget(scope constructs.Construct, id *string, config *SecurityhubActionTargetConfig) SecurityhubActionTarget {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubActionTarget{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubActionTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html aws_securityhub_action_target} Resource.
func NewSecurityhubActionTarget_Override(s SecurityhubActionTarget, scope constructs.Construct, id *string, config *SecurityhubActionTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubActionTarget",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubActionTarget) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubActionTarget",
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
		"hashicorp_aws.SecurityHub.SecurityhubActionTarget",
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

type SecurityhubActionTargetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html#description SecurityhubActionTarget#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html#identifier SecurityhubActionTarget#identifier}.
	Identifier *string `json:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_action_target.html#name SecurityhubActionTarget#name}.
	Name *string `json:"name"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html aws_securityhub_insight}.
type SecurityhubInsight interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_SecurityhubInsight) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html aws_securityhub_insight} Resource.
func NewSecurityhubInsight(scope constructs.Construct, id *string, config *SecurityhubInsightConfig) SecurityhubInsight {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsight{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsight",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html aws_securityhub_insight} Resource.
func NewSecurityhubInsight_Override(s SecurityhubInsight, scope constructs.Construct, id *string, config *SecurityhubInsightConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsight",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsight) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubInsight",
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
		"hashicorp_aws.SecurityHub.SecurityhubInsight",
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

type SecurityhubInsightConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#filters SecurityhubInsight#filters}
	Filters *SecurityhubInsightFilters `json:"filters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#group_by_attribute SecurityhubInsight#group_by_attribute}.
	GroupByAttribute *string `json:"groupByAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#name SecurityhubInsight#name}.
	Name *string `json:"name"`
}

type SecurityhubInsightFilters struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#aws_account_id SecurityhubInsight#aws_account_id}
	AwsAccountId *[]*SecurityhubInsightFiltersAwsAccountId `json:"awsAccountId"`
	// company_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#company_name SecurityhubInsight#company_name}
	CompanyName *[]*SecurityhubInsightFiltersCompanyName `json:"companyName"`
	// compliance_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#compliance_status SecurityhubInsight#compliance_status}
	ComplianceStatus *[]*SecurityhubInsightFiltersComplianceStatus `json:"complianceStatus"`
	// confidence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#confidence SecurityhubInsight#confidence}
	Confidence *[]*SecurityhubInsightFiltersConfidence `json:"confidence"`
	// created_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#created_at SecurityhubInsight#created_at}
	CreatedAt *[]*SecurityhubInsightFiltersCreatedAt `json:"createdAt"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#criticality SecurityhubInsight#criticality}
	Criticality *[]*SecurityhubInsightFiltersCriticality `json:"criticality"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#description SecurityhubInsight#description}
	Description *[]*SecurityhubInsightFiltersDescription `json:"description"`
	// finding_provider_fields_confidence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_confidence SecurityhubInsight#finding_provider_fields_confidence}
	FindingProviderFieldsConfidence *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence `json:"findingProviderFieldsConfidence"`
	// finding_provider_fields_criticality block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_criticality SecurityhubInsight#finding_provider_fields_criticality}
	FindingProviderFieldsCriticality *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality `json:"findingProviderFieldsCriticality"`
	// finding_provider_fields_related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_related_findings_id SecurityhubInsight#finding_provider_fields_related_findings_id}
	FindingProviderFieldsRelatedFindingsId *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId `json:"findingProviderFieldsRelatedFindingsId"`
	// finding_provider_fields_related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_related_findings_product_arn SecurityhubInsight#finding_provider_fields_related_findings_product_arn}
	FindingProviderFieldsRelatedFindingsProductArn *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn `json:"findingProviderFieldsRelatedFindingsProductArn"`
	// finding_provider_fields_severity_label block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_severity_label SecurityhubInsight#finding_provider_fields_severity_label}
	FindingProviderFieldsSeverityLabel *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel `json:"findingProviderFieldsSeverityLabel"`
	// finding_provider_fields_severity_original block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_severity_original SecurityhubInsight#finding_provider_fields_severity_original}
	FindingProviderFieldsSeverityOriginal *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal `json:"findingProviderFieldsSeverityOriginal"`
	// finding_provider_fields_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#finding_provider_fields_types SecurityhubInsight#finding_provider_fields_types}
	FindingProviderFieldsTypes *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes `json:"findingProviderFieldsTypes"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#first_observed_at SecurityhubInsight#first_observed_at}
	FirstObservedAt *[]*SecurityhubInsightFiltersFirstObservedAt `json:"firstObservedAt"`
	// generator_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#generator_id SecurityhubInsight#generator_id}
	GeneratorId *[]*SecurityhubInsightFiltersGeneratorId `json:"generatorId"`
	// id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#id SecurityhubInsight#id}
	Id *[]*SecurityhubInsightFiltersId `json:"id"`
	// keyword block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#keyword SecurityhubInsight#keyword}
	Keyword *[]*SecurityhubInsightFiltersKeyword `json:"keyword"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#last_observed_at SecurityhubInsight#last_observed_at}
	LastObservedAt *[]*SecurityhubInsightFiltersLastObservedAt `json:"lastObservedAt"`
	// malware_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#malware_name SecurityhubInsight#malware_name}
	MalwareName *[]*SecurityhubInsightFiltersMalwareName `json:"malwareName"`
	// malware_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#malware_path SecurityhubInsight#malware_path}
	MalwarePath *[]*SecurityhubInsightFiltersMalwarePath `json:"malwarePath"`
	// malware_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#malware_state SecurityhubInsight#malware_state}
	MalwareState *[]*SecurityhubInsightFiltersMalwareState `json:"malwareState"`
	// malware_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#malware_type SecurityhubInsight#malware_type}
	MalwareType *[]*SecurityhubInsightFiltersMalwareType `json:"malwareType"`
	// network_destination_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_destination_domain SecurityhubInsight#network_destination_domain}
	NetworkDestinationDomain *[]*SecurityhubInsightFiltersNetworkDestinationDomain `json:"networkDestinationDomain"`
	// network_destination_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_destination_ipv4 SecurityhubInsight#network_destination_ipv4}
	NetworkDestinationIpv4 *[]*SecurityhubInsightFiltersNetworkDestinationIpv4 `json:"networkDestinationIpv4"`
	// network_destination_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_destination_ipv6 SecurityhubInsight#network_destination_ipv6}
	NetworkDestinationIpv6 *[]*SecurityhubInsightFiltersNetworkDestinationIpv6 `json:"networkDestinationIpv6"`
	// network_destination_port block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_destination_port SecurityhubInsight#network_destination_port}
	NetworkDestinationPort *[]*SecurityhubInsightFiltersNetworkDestinationPort `json:"networkDestinationPort"`
	// network_direction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_direction SecurityhubInsight#network_direction}
	NetworkDirection *[]*SecurityhubInsightFiltersNetworkDirection `json:"networkDirection"`
	// network_protocol block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_protocol SecurityhubInsight#network_protocol}
	NetworkProtocol *[]*SecurityhubInsightFiltersNetworkProtocol `json:"networkProtocol"`
	// network_source_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_source_domain SecurityhubInsight#network_source_domain}
	NetworkSourceDomain *[]*SecurityhubInsightFiltersNetworkSourceDomain `json:"networkSourceDomain"`
	// network_source_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_source_ipv4 SecurityhubInsight#network_source_ipv4}
	NetworkSourceIpv4 *[]*SecurityhubInsightFiltersNetworkSourceIpv4 `json:"networkSourceIpv4"`
	// network_source_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_source_ipv6 SecurityhubInsight#network_source_ipv6}
	NetworkSourceIpv6 *[]*SecurityhubInsightFiltersNetworkSourceIpv6 `json:"networkSourceIpv6"`
	// network_source_mac block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_source_mac SecurityhubInsight#network_source_mac}
	NetworkSourceMac *[]*SecurityhubInsightFiltersNetworkSourceMac `json:"networkSourceMac"`
	// network_source_port block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#network_source_port SecurityhubInsight#network_source_port}
	NetworkSourcePort *[]*SecurityhubInsightFiltersNetworkSourcePort `json:"networkSourcePort"`
	// note_text block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#note_text SecurityhubInsight#note_text}
	NoteText *[]*SecurityhubInsightFiltersNoteText `json:"noteText"`
	// note_updated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#note_updated_at SecurityhubInsight#note_updated_at}
	NoteUpdatedAt *[]*SecurityhubInsightFiltersNoteUpdatedAt `json:"noteUpdatedAt"`
	// note_updated_by block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#note_updated_by SecurityhubInsight#note_updated_by}
	NoteUpdatedBy *[]*SecurityhubInsightFiltersNoteUpdatedBy `json:"noteUpdatedBy"`
	// process_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_launched_at SecurityhubInsight#process_launched_at}
	ProcessLaunchedAt *[]*SecurityhubInsightFiltersProcessLaunchedAt `json:"processLaunchedAt"`
	// process_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_name SecurityhubInsight#process_name}
	ProcessName *[]*SecurityhubInsightFiltersProcessName `json:"processName"`
	// process_parent_pid block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_parent_pid SecurityhubInsight#process_parent_pid}
	ProcessParentPid *[]*SecurityhubInsightFiltersProcessParentPid `json:"processParentPid"`
	// process_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_path SecurityhubInsight#process_path}
	ProcessPath *[]*SecurityhubInsightFiltersProcessPath `json:"processPath"`
	// process_pid block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_pid SecurityhubInsight#process_pid}
	ProcessPid *[]*SecurityhubInsightFiltersProcessPid `json:"processPid"`
	// process_terminated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#process_terminated_at SecurityhubInsight#process_terminated_at}
	ProcessTerminatedAt *[]*SecurityhubInsightFiltersProcessTerminatedAt `json:"processTerminatedAt"`
	// product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#product_arn SecurityhubInsight#product_arn}
	ProductArn *[]*SecurityhubInsightFiltersProductArn `json:"productArn"`
	// product_fields block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#product_fields SecurityhubInsight#product_fields}
	ProductFields *[]*SecurityhubInsightFiltersProductFields `json:"productFields"`
	// product_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#product_name SecurityhubInsight#product_name}
	ProductName *[]*SecurityhubInsightFiltersProductName `json:"productName"`
	// recommendation_text block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#recommendation_text SecurityhubInsight#recommendation_text}
	RecommendationText *[]*SecurityhubInsightFiltersRecommendationText `json:"recommendationText"`
	// record_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#record_state SecurityhubInsight#record_state}
	RecordState *[]*SecurityhubInsightFiltersRecordState `json:"recordState"`
	// related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#related_findings_id SecurityhubInsight#related_findings_id}
	RelatedFindingsId *[]*SecurityhubInsightFiltersRelatedFindingsId `json:"relatedFindingsId"`
	// related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#related_findings_product_arn SecurityhubInsight#related_findings_product_arn}
	RelatedFindingsProductArn *[]*SecurityhubInsightFiltersRelatedFindingsProductArn `json:"relatedFindingsProductArn"`
	// resource_aws_ec2_instance_iam_instance_profile_arn block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_iam_instance_profile_arn SecurityhubInsight#resource_aws_ec2_instance_iam_instance_profile_arn}
	ResourceAwsEc2InstanceIamInstanceProfileArn *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn `json:"resourceAwsEc2InstanceIamInstanceProfileArn"`
	// resource_aws_ec2_instance_image_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_image_id SecurityhubInsight#resource_aws_ec2_instance_image_id}
	ResourceAwsEc2InstanceImageId *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId `json:"resourceAwsEc2InstanceImageId"`
	// resource_aws_ec2_instance_ipv4_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_ipv4_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv4_addresses}
	ResourceAwsEc2InstanceIpv4Addresses *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses `json:"resourceAwsEc2InstanceIpv4Addresses"`
	// resource_aws_ec2_instance_ipv6_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_ipv6_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv6_addresses}
	ResourceAwsEc2InstanceIpv6Addresses *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses `json:"resourceAwsEc2InstanceIpv6Addresses"`
	// resource_aws_ec2_instance_key_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_key_name SecurityhubInsight#resource_aws_ec2_instance_key_name}
	ResourceAwsEc2InstanceKeyName *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName `json:"resourceAwsEc2InstanceKeyName"`
	// resource_aws_ec2_instance_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_launched_at SecurityhubInsight#resource_aws_ec2_instance_launched_at}
	ResourceAwsEc2InstanceLaunchedAt *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt `json:"resourceAwsEc2InstanceLaunchedAt"`
	// resource_aws_ec2_instance_subnet_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_subnet_id SecurityhubInsight#resource_aws_ec2_instance_subnet_id}
	ResourceAwsEc2InstanceSubnetId *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId `json:"resourceAwsEc2InstanceSubnetId"`
	// resource_aws_ec2_instance_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_type SecurityhubInsight#resource_aws_ec2_instance_type}
	ResourceAwsEc2InstanceType *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType `json:"resourceAwsEc2InstanceType"`
	// resource_aws_ec2_instance_vpc_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_ec2_instance_vpc_id SecurityhubInsight#resource_aws_ec2_instance_vpc_id}
	ResourceAwsEc2InstanceVpcId *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId `json:"resourceAwsEc2InstanceVpcId"`
	// resource_aws_iam_access_key_created_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_iam_access_key_created_at SecurityhubInsight#resource_aws_iam_access_key_created_at}
	ResourceAwsIamAccessKeyCreatedAt *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt `json:"resourceAwsIamAccessKeyCreatedAt"`
	// resource_aws_iam_access_key_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_iam_access_key_status SecurityhubInsight#resource_aws_iam_access_key_status}
	ResourceAwsIamAccessKeyStatus *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus `json:"resourceAwsIamAccessKeyStatus"`
	// resource_aws_iam_access_key_user_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_iam_access_key_user_name SecurityhubInsight#resource_aws_iam_access_key_user_name}
	ResourceAwsIamAccessKeyUserName *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName `json:"resourceAwsIamAccessKeyUserName"`
	// resource_aws_s3_bucket_owner_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_s3_bucket_owner_id SecurityhubInsight#resource_aws_s3_bucket_owner_id}
	ResourceAwsS3BucketOwnerId *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId `json:"resourceAwsS3BucketOwnerId"`
	// resource_aws_s3_bucket_owner_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_aws_s3_bucket_owner_name SecurityhubInsight#resource_aws_s3_bucket_owner_name}
	ResourceAwsS3BucketOwnerName *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName `json:"resourceAwsS3BucketOwnerName"`
	// resource_container_image_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_container_image_id SecurityhubInsight#resource_container_image_id}
	ResourceContainerImageId *[]*SecurityhubInsightFiltersResourceContainerImageId `json:"resourceContainerImageId"`
	// resource_container_image_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_container_image_name SecurityhubInsight#resource_container_image_name}
	ResourceContainerImageName *[]*SecurityhubInsightFiltersResourceContainerImageName `json:"resourceContainerImageName"`
	// resource_container_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_container_launched_at SecurityhubInsight#resource_container_launched_at}
	ResourceContainerLaunchedAt *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt `json:"resourceContainerLaunchedAt"`
	// resource_container_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_container_name SecurityhubInsight#resource_container_name}
	ResourceContainerName *[]*SecurityhubInsightFiltersResourceContainerName `json:"resourceContainerName"`
	// resource_details_other block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_details_other SecurityhubInsight#resource_details_other}
	ResourceDetailsOther *[]*SecurityhubInsightFiltersResourceDetailsOther `json:"resourceDetailsOther"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_id SecurityhubInsight#resource_id}
	ResourceId *[]*SecurityhubInsightFiltersResourceId `json:"resourceId"`
	// resource_partition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_partition SecurityhubInsight#resource_partition}
	ResourcePartition *[]*SecurityhubInsightFiltersResourcePartition `json:"resourcePartition"`
	// resource_region block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_region SecurityhubInsight#resource_region}
	ResourceRegion *[]*SecurityhubInsightFiltersResourceRegion `json:"resourceRegion"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_tags SecurityhubInsight#resource_tags}
	ResourceTags *[]*SecurityhubInsightFiltersResourceTags `json:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#resource_type SecurityhubInsight#resource_type}
	ResourceType *[]*SecurityhubInsightFiltersResourceType `json:"resourceType"`
	// severity_label block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#severity_label SecurityhubInsight#severity_label}
	SeverityLabel *[]*SecurityhubInsightFiltersSeverityLabel `json:"severityLabel"`
	// source_url block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#source_url SecurityhubInsight#source_url}
	SourceUrl *[]*SecurityhubInsightFiltersSourceUrl `json:"sourceUrl"`
	// threat_intel_indicator_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_category SecurityhubInsight#threat_intel_indicator_category}
	ThreatIntelIndicatorCategory *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory `json:"threatIntelIndicatorCategory"`
	// threat_intel_indicator_last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_last_observed_at SecurityhubInsight#threat_intel_indicator_last_observed_at}
	ThreatIntelIndicatorLastObservedAt *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt `json:"threatIntelIndicatorLastObservedAt"`
	// threat_intel_indicator_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_source SecurityhubInsight#threat_intel_indicator_source}
	ThreatIntelIndicatorSource *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource `json:"threatIntelIndicatorSource"`
	// threat_intel_indicator_source_url block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_source_url SecurityhubInsight#threat_intel_indicator_source_url}
	ThreatIntelIndicatorSourceUrl *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl `json:"threatIntelIndicatorSourceUrl"`
	// threat_intel_indicator_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_type SecurityhubInsight#threat_intel_indicator_type}
	ThreatIntelIndicatorType *[]*SecurityhubInsightFiltersThreatIntelIndicatorType `json:"threatIntelIndicatorType"`
	// threat_intel_indicator_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#threat_intel_indicator_value SecurityhubInsight#threat_intel_indicator_value}
	ThreatIntelIndicatorValue *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue `json:"threatIntelIndicatorValue"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#title SecurityhubInsight#title}
	Title *[]*SecurityhubInsightFiltersTitle `json:"title"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#type SecurityhubInsight#type}
	Type *[]*SecurityhubInsightFiltersType `json:"type"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#updated_at SecurityhubInsight#updated_at}
	UpdatedAt *[]*SecurityhubInsightFiltersUpdatedAt `json:"updatedAt"`
	// user_defined_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#user_defined_values SecurityhubInsight#user_defined_values}
	UserDefinedValues *[]*SecurityhubInsightFiltersUserDefinedValues `json:"userDefinedValues"`
	// verification_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#verification_state SecurityhubInsight#verification_state}
	VerificationState *[]*SecurityhubInsightFiltersVerificationState `json:"verificationState"`
	// workflow_status block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#workflow_status SecurityhubInsight#workflow_status}
	WorkflowStatus *[]*SecurityhubInsightFiltersWorkflowStatus `json:"workflowStatus"`
}

type SecurityhubInsightFiltersAwsAccountId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersCompanyName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersComplianceStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersConfidence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersCreatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersCreatedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersCreatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersCreatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersCreatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersCreatedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersCreatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersCreatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersCreatedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersDescription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsConfidence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersFindingProviderFieldsCriticality struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFindingProviderFieldsTypes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersFirstObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersFirstObservedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersFirstObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersKeyword struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersLastObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersLastObservedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersLastObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersLastObservedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersLastObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersMalwarePath struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersMalwareState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersMalwareType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkDestinationDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkDestinationIpv4 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersNetworkDestinationIpv6 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersNetworkDestinationPort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersNetworkDirection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkProtocol struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkSourceDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkSourceIpv4 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersNetworkSourceIpv6 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersNetworkSourceMac struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNetworkSourcePort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersNoteText struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersNoteUpdatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersNoteUpdatedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersNoteUpdatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() *[]*SecurityhubInsightFiltersAwsAccountId
	SetAwsAccountId(val *[]*SecurityhubInsightFiltersAwsAccountId)
	AwsAccountIdInput() *[]*SecurityhubInsightFiltersAwsAccountId
	CompanyName() *[]*SecurityhubInsightFiltersCompanyName
	SetCompanyName(val *[]*SecurityhubInsightFiltersCompanyName)
	CompanyNameInput() *[]*SecurityhubInsightFiltersCompanyName
	ComplianceStatus() *[]*SecurityhubInsightFiltersComplianceStatus
	SetComplianceStatus(val *[]*SecurityhubInsightFiltersComplianceStatus)
	ComplianceStatusInput() *[]*SecurityhubInsightFiltersComplianceStatus
	Confidence() *[]*SecurityhubInsightFiltersConfidence
	SetConfidence(val *[]*SecurityhubInsightFiltersConfidence)
	ConfidenceInput() *[]*SecurityhubInsightFiltersConfidence
	CreatedAt() *[]*SecurityhubInsightFiltersCreatedAt
	SetCreatedAt(val *[]*SecurityhubInsightFiltersCreatedAt)
	CreatedAtInput() *[]*SecurityhubInsightFiltersCreatedAt
	Criticality() *[]*SecurityhubInsightFiltersCriticality
	SetCriticality(val *[]*SecurityhubInsightFiltersCriticality)
	CriticalityInput() *[]*SecurityhubInsightFiltersCriticality
	Description() *[]*SecurityhubInsightFiltersDescription
	SetDescription(val *[]*SecurityhubInsightFiltersDescription)
	DescriptionInput() *[]*SecurityhubInsightFiltersDescription
	FindingProviderFieldsConfidence() *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence
	SetFindingProviderFieldsConfidence(val *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence)
	FindingProviderFieldsConfidenceInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence
	FindingProviderFieldsCriticality() *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality
	SetFindingProviderFieldsCriticality(val *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality)
	FindingProviderFieldsCriticalityInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality
	FindingProviderFieldsRelatedFindingsId() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId
	SetFindingProviderFieldsRelatedFindingsId(val *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId)
	FindingProviderFieldsRelatedFindingsIdInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId
	FindingProviderFieldsRelatedFindingsProductArn() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn
	SetFindingProviderFieldsRelatedFindingsProductArn(val *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn)
	FindingProviderFieldsRelatedFindingsProductArnInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn
	FindingProviderFieldsSeverityLabel() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel
	SetFindingProviderFieldsSeverityLabel(val *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel)
	FindingProviderFieldsSeverityLabelInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel
	FindingProviderFieldsSeverityOriginal() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal
	SetFindingProviderFieldsSeverityOriginal(val *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal)
	FindingProviderFieldsSeverityOriginalInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal
	FindingProviderFieldsTypes() *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes
	SetFindingProviderFieldsTypes(val *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes)
	FindingProviderFieldsTypesInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes
	FirstObservedAt() *[]*SecurityhubInsightFiltersFirstObservedAt
	SetFirstObservedAt(val *[]*SecurityhubInsightFiltersFirstObservedAt)
	FirstObservedAtInput() *[]*SecurityhubInsightFiltersFirstObservedAt
	GeneratorId() *[]*SecurityhubInsightFiltersGeneratorId
	SetGeneratorId(val *[]*SecurityhubInsightFiltersGeneratorId)
	GeneratorIdInput() *[]*SecurityhubInsightFiltersGeneratorId
	Id() *[]*SecurityhubInsightFiltersId
	SetId(val *[]*SecurityhubInsightFiltersId)
	IdInput() *[]*SecurityhubInsightFiltersId
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Keyword() *[]*SecurityhubInsightFiltersKeyword
	SetKeyword(val *[]*SecurityhubInsightFiltersKeyword)
	KeywordInput() *[]*SecurityhubInsightFiltersKeyword
	LastObservedAt() *[]*SecurityhubInsightFiltersLastObservedAt
	SetLastObservedAt(val *[]*SecurityhubInsightFiltersLastObservedAt)
	LastObservedAtInput() *[]*SecurityhubInsightFiltersLastObservedAt
	MalwareName() *[]*SecurityhubInsightFiltersMalwareName
	SetMalwareName(val *[]*SecurityhubInsightFiltersMalwareName)
	MalwareNameInput() *[]*SecurityhubInsightFiltersMalwareName
	MalwarePath() *[]*SecurityhubInsightFiltersMalwarePath
	SetMalwarePath(val *[]*SecurityhubInsightFiltersMalwarePath)
	MalwarePathInput() *[]*SecurityhubInsightFiltersMalwarePath
	MalwareState() *[]*SecurityhubInsightFiltersMalwareState
	SetMalwareState(val *[]*SecurityhubInsightFiltersMalwareState)
	MalwareStateInput() *[]*SecurityhubInsightFiltersMalwareState
	MalwareType() *[]*SecurityhubInsightFiltersMalwareType
	SetMalwareType(val *[]*SecurityhubInsightFiltersMalwareType)
	MalwareTypeInput() *[]*SecurityhubInsightFiltersMalwareType
	NetworkDestinationDomain() *[]*SecurityhubInsightFiltersNetworkDestinationDomain
	SetNetworkDestinationDomain(val *[]*SecurityhubInsightFiltersNetworkDestinationDomain)
	NetworkDestinationDomainInput() *[]*SecurityhubInsightFiltersNetworkDestinationDomain
	NetworkDestinationIpv4() *[]*SecurityhubInsightFiltersNetworkDestinationIpv4
	SetNetworkDestinationIpv4(val *[]*SecurityhubInsightFiltersNetworkDestinationIpv4)
	NetworkDestinationIpv4Input() *[]*SecurityhubInsightFiltersNetworkDestinationIpv4
	NetworkDestinationIpv6() *[]*SecurityhubInsightFiltersNetworkDestinationIpv6
	SetNetworkDestinationIpv6(val *[]*SecurityhubInsightFiltersNetworkDestinationIpv6)
	NetworkDestinationIpv6Input() *[]*SecurityhubInsightFiltersNetworkDestinationIpv6
	NetworkDestinationPort() *[]*SecurityhubInsightFiltersNetworkDestinationPort
	SetNetworkDestinationPort(val *[]*SecurityhubInsightFiltersNetworkDestinationPort)
	NetworkDestinationPortInput() *[]*SecurityhubInsightFiltersNetworkDestinationPort
	NetworkDirection() *[]*SecurityhubInsightFiltersNetworkDirection
	SetNetworkDirection(val *[]*SecurityhubInsightFiltersNetworkDirection)
	NetworkDirectionInput() *[]*SecurityhubInsightFiltersNetworkDirection
	NetworkProtocol() *[]*SecurityhubInsightFiltersNetworkProtocol
	SetNetworkProtocol(val *[]*SecurityhubInsightFiltersNetworkProtocol)
	NetworkProtocolInput() *[]*SecurityhubInsightFiltersNetworkProtocol
	NetworkSourceDomain() *[]*SecurityhubInsightFiltersNetworkSourceDomain
	SetNetworkSourceDomain(val *[]*SecurityhubInsightFiltersNetworkSourceDomain)
	NetworkSourceDomainInput() *[]*SecurityhubInsightFiltersNetworkSourceDomain
	NetworkSourceIpv4() *[]*SecurityhubInsightFiltersNetworkSourceIpv4
	SetNetworkSourceIpv4(val *[]*SecurityhubInsightFiltersNetworkSourceIpv4)
	NetworkSourceIpv4Input() *[]*SecurityhubInsightFiltersNetworkSourceIpv4
	NetworkSourceIpv6() *[]*SecurityhubInsightFiltersNetworkSourceIpv6
	SetNetworkSourceIpv6(val *[]*SecurityhubInsightFiltersNetworkSourceIpv6)
	NetworkSourceIpv6Input() *[]*SecurityhubInsightFiltersNetworkSourceIpv6
	NetworkSourceMac() *[]*SecurityhubInsightFiltersNetworkSourceMac
	SetNetworkSourceMac(val *[]*SecurityhubInsightFiltersNetworkSourceMac)
	NetworkSourceMacInput() *[]*SecurityhubInsightFiltersNetworkSourceMac
	NetworkSourcePort() *[]*SecurityhubInsightFiltersNetworkSourcePort
	SetNetworkSourcePort(val *[]*SecurityhubInsightFiltersNetworkSourcePort)
	NetworkSourcePortInput() *[]*SecurityhubInsightFiltersNetworkSourcePort
	NoteText() *[]*SecurityhubInsightFiltersNoteText
	SetNoteText(val *[]*SecurityhubInsightFiltersNoteText)
	NoteTextInput() *[]*SecurityhubInsightFiltersNoteText
	NoteUpdatedAt() *[]*SecurityhubInsightFiltersNoteUpdatedAt
	SetNoteUpdatedAt(val *[]*SecurityhubInsightFiltersNoteUpdatedAt)
	NoteUpdatedAtInput() *[]*SecurityhubInsightFiltersNoteUpdatedAt
	NoteUpdatedBy() *[]*SecurityhubInsightFiltersNoteUpdatedBy
	SetNoteUpdatedBy(val *[]*SecurityhubInsightFiltersNoteUpdatedBy)
	NoteUpdatedByInput() *[]*SecurityhubInsightFiltersNoteUpdatedBy
	ProcessLaunchedAt() *[]*SecurityhubInsightFiltersProcessLaunchedAt
	SetProcessLaunchedAt(val *[]*SecurityhubInsightFiltersProcessLaunchedAt)
	ProcessLaunchedAtInput() *[]*SecurityhubInsightFiltersProcessLaunchedAt
	ProcessName() *[]*SecurityhubInsightFiltersProcessName
	SetProcessName(val *[]*SecurityhubInsightFiltersProcessName)
	ProcessNameInput() *[]*SecurityhubInsightFiltersProcessName
	ProcessParentPid() *[]*SecurityhubInsightFiltersProcessParentPid
	SetProcessParentPid(val *[]*SecurityhubInsightFiltersProcessParentPid)
	ProcessParentPidInput() *[]*SecurityhubInsightFiltersProcessParentPid
	ProcessPath() *[]*SecurityhubInsightFiltersProcessPath
	SetProcessPath(val *[]*SecurityhubInsightFiltersProcessPath)
	ProcessPathInput() *[]*SecurityhubInsightFiltersProcessPath
	ProcessPid() *[]*SecurityhubInsightFiltersProcessPid
	SetProcessPid(val *[]*SecurityhubInsightFiltersProcessPid)
	ProcessPidInput() *[]*SecurityhubInsightFiltersProcessPid
	ProcessTerminatedAt() *[]*SecurityhubInsightFiltersProcessTerminatedAt
	SetProcessTerminatedAt(val *[]*SecurityhubInsightFiltersProcessTerminatedAt)
	ProcessTerminatedAtInput() *[]*SecurityhubInsightFiltersProcessTerminatedAt
	ProductArn() *[]*SecurityhubInsightFiltersProductArn
	SetProductArn(val *[]*SecurityhubInsightFiltersProductArn)
	ProductArnInput() *[]*SecurityhubInsightFiltersProductArn
	ProductFields() *[]*SecurityhubInsightFiltersProductFields
	SetProductFields(val *[]*SecurityhubInsightFiltersProductFields)
	ProductFieldsInput() *[]*SecurityhubInsightFiltersProductFields
	ProductName() *[]*SecurityhubInsightFiltersProductName
	SetProductName(val *[]*SecurityhubInsightFiltersProductName)
	ProductNameInput() *[]*SecurityhubInsightFiltersProductName
	RecommendationText() *[]*SecurityhubInsightFiltersRecommendationText
	SetRecommendationText(val *[]*SecurityhubInsightFiltersRecommendationText)
	RecommendationTextInput() *[]*SecurityhubInsightFiltersRecommendationText
	RecordState() *[]*SecurityhubInsightFiltersRecordState
	SetRecordState(val *[]*SecurityhubInsightFiltersRecordState)
	RecordStateInput() *[]*SecurityhubInsightFiltersRecordState
	RelatedFindingsId() *[]*SecurityhubInsightFiltersRelatedFindingsId
	SetRelatedFindingsId(val *[]*SecurityhubInsightFiltersRelatedFindingsId)
	RelatedFindingsIdInput() *[]*SecurityhubInsightFiltersRelatedFindingsId
	RelatedFindingsProductArn() *[]*SecurityhubInsightFiltersRelatedFindingsProductArn
	SetRelatedFindingsProductArn(val *[]*SecurityhubInsightFiltersRelatedFindingsProductArn)
	RelatedFindingsProductArnInput() *[]*SecurityhubInsightFiltersRelatedFindingsProductArn
	ResourceAwsEc2InstanceIamInstanceProfileArn() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn
	SetResourceAwsEc2InstanceIamInstanceProfileArn(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn)
	ResourceAwsEc2InstanceIamInstanceProfileArnInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn
	ResourceAwsEc2InstanceImageId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId
	SetResourceAwsEc2InstanceImageId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId)
	ResourceAwsEc2InstanceImageIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId
	ResourceAwsEc2InstanceIpv4Addresses() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses
	SetResourceAwsEc2InstanceIpv4Addresses(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses)
	ResourceAwsEc2InstanceIpv4AddressesInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses
	ResourceAwsEc2InstanceIpv6Addresses() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses
	SetResourceAwsEc2InstanceIpv6Addresses(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses)
	ResourceAwsEc2InstanceIpv6AddressesInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses
	ResourceAwsEc2InstanceKeyName() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName
	SetResourceAwsEc2InstanceKeyName(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName)
	ResourceAwsEc2InstanceKeyNameInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName
	ResourceAwsEc2InstanceLaunchedAt() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt
	SetResourceAwsEc2InstanceLaunchedAt(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt)
	ResourceAwsEc2InstanceLaunchedAtInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt
	ResourceAwsEc2InstanceSubnetId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId
	SetResourceAwsEc2InstanceSubnetId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId)
	ResourceAwsEc2InstanceSubnetIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId
	ResourceAwsEc2InstanceType() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType
	SetResourceAwsEc2InstanceType(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType)
	ResourceAwsEc2InstanceTypeInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType
	ResourceAwsEc2InstanceVpcId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId
	SetResourceAwsEc2InstanceVpcId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId)
	ResourceAwsEc2InstanceVpcIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId
	ResourceAwsIamAccessKeyCreatedAt() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt
	SetResourceAwsIamAccessKeyCreatedAt(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt)
	ResourceAwsIamAccessKeyCreatedAtInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt
	ResourceAwsIamAccessKeyStatus() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus
	SetResourceAwsIamAccessKeyStatus(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus)
	ResourceAwsIamAccessKeyStatusInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus
	ResourceAwsIamAccessKeyUserName() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName
	SetResourceAwsIamAccessKeyUserName(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName)
	ResourceAwsIamAccessKeyUserNameInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName
	ResourceAwsS3BucketOwnerId() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId
	SetResourceAwsS3BucketOwnerId(val *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId)
	ResourceAwsS3BucketOwnerIdInput() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId
	ResourceAwsS3BucketOwnerName() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName
	SetResourceAwsS3BucketOwnerName(val *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName)
	ResourceAwsS3BucketOwnerNameInput() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName
	ResourceContainerImageId() *[]*SecurityhubInsightFiltersResourceContainerImageId
	SetResourceContainerImageId(val *[]*SecurityhubInsightFiltersResourceContainerImageId)
	ResourceContainerImageIdInput() *[]*SecurityhubInsightFiltersResourceContainerImageId
	ResourceContainerImageName() *[]*SecurityhubInsightFiltersResourceContainerImageName
	SetResourceContainerImageName(val *[]*SecurityhubInsightFiltersResourceContainerImageName)
	ResourceContainerImageNameInput() *[]*SecurityhubInsightFiltersResourceContainerImageName
	ResourceContainerLaunchedAt() *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt
	SetResourceContainerLaunchedAt(val *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt)
	ResourceContainerLaunchedAtInput() *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt
	ResourceContainerName() *[]*SecurityhubInsightFiltersResourceContainerName
	SetResourceContainerName(val *[]*SecurityhubInsightFiltersResourceContainerName)
	ResourceContainerNameInput() *[]*SecurityhubInsightFiltersResourceContainerName
	ResourceDetailsOther() *[]*SecurityhubInsightFiltersResourceDetailsOther
	SetResourceDetailsOther(val *[]*SecurityhubInsightFiltersResourceDetailsOther)
	ResourceDetailsOtherInput() *[]*SecurityhubInsightFiltersResourceDetailsOther
	ResourceId() *[]*SecurityhubInsightFiltersResourceId
	SetResourceId(val *[]*SecurityhubInsightFiltersResourceId)
	ResourceIdInput() *[]*SecurityhubInsightFiltersResourceId
	ResourcePartition() *[]*SecurityhubInsightFiltersResourcePartition
	SetResourcePartition(val *[]*SecurityhubInsightFiltersResourcePartition)
	ResourcePartitionInput() *[]*SecurityhubInsightFiltersResourcePartition
	ResourceRegion() *[]*SecurityhubInsightFiltersResourceRegion
	SetResourceRegion(val *[]*SecurityhubInsightFiltersResourceRegion)
	ResourceRegionInput() *[]*SecurityhubInsightFiltersResourceRegion
	ResourceTags() *[]*SecurityhubInsightFiltersResourceTags
	SetResourceTags(val *[]*SecurityhubInsightFiltersResourceTags)
	ResourceTagsInput() *[]*SecurityhubInsightFiltersResourceTags
	ResourceType() *[]*SecurityhubInsightFiltersResourceType
	SetResourceType(val *[]*SecurityhubInsightFiltersResourceType)
	ResourceTypeInput() *[]*SecurityhubInsightFiltersResourceType
	SeverityLabel() *[]*SecurityhubInsightFiltersSeverityLabel
	SetSeverityLabel(val *[]*SecurityhubInsightFiltersSeverityLabel)
	SeverityLabelInput() *[]*SecurityhubInsightFiltersSeverityLabel
	SourceUrl() *[]*SecurityhubInsightFiltersSourceUrl
	SetSourceUrl(val *[]*SecurityhubInsightFiltersSourceUrl)
	SourceUrlInput() *[]*SecurityhubInsightFiltersSourceUrl
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	ThreatIntelIndicatorCategory() *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory
	SetThreatIntelIndicatorCategory(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory)
	ThreatIntelIndicatorCategoryInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory
	ThreatIntelIndicatorLastObservedAt() *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt
	SetThreatIntelIndicatorLastObservedAt(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt)
	ThreatIntelIndicatorLastObservedAtInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt
	ThreatIntelIndicatorSource() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource
	SetThreatIntelIndicatorSource(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource)
	ThreatIntelIndicatorSourceInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource
	ThreatIntelIndicatorSourceUrl() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl
	SetThreatIntelIndicatorSourceUrl(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl)
	ThreatIntelIndicatorSourceUrlInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl
	ThreatIntelIndicatorType() *[]*SecurityhubInsightFiltersThreatIntelIndicatorType
	SetThreatIntelIndicatorType(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorType)
	ThreatIntelIndicatorTypeInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorType
	ThreatIntelIndicatorValue() *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue
	SetThreatIntelIndicatorValue(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue)
	ThreatIntelIndicatorValueInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue
	Title() *[]*SecurityhubInsightFiltersTitle
	SetTitle(val *[]*SecurityhubInsightFiltersTitle)
	TitleInput() *[]*SecurityhubInsightFiltersTitle
	Type() *[]*SecurityhubInsightFiltersType
	SetType(val *[]*SecurityhubInsightFiltersType)
	TypeInput() *[]*SecurityhubInsightFiltersType
	UpdatedAt() *[]*SecurityhubInsightFiltersUpdatedAt
	SetUpdatedAt(val *[]*SecurityhubInsightFiltersUpdatedAt)
	UpdatedAtInput() *[]*SecurityhubInsightFiltersUpdatedAt
	UserDefinedValues() *[]*SecurityhubInsightFiltersUserDefinedValues
	SetUserDefinedValues(val *[]*SecurityhubInsightFiltersUserDefinedValues)
	UserDefinedValuesInput() *[]*SecurityhubInsightFiltersUserDefinedValues
	VerificationState() *[]*SecurityhubInsightFiltersVerificationState
	SetVerificationState(val *[]*SecurityhubInsightFiltersVerificationState)
	VerificationStateInput() *[]*SecurityhubInsightFiltersVerificationState
	WorkflowStatus() *[]*SecurityhubInsightFiltersWorkflowStatus
	SetWorkflowStatus(val *[]*SecurityhubInsightFiltersWorkflowStatus)
	WorkflowStatusInput() *[]*SecurityhubInsightFiltersWorkflowStatus
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountId() *[]*SecurityhubInsightFiltersAwsAccountId {
	var returns *[]*SecurityhubInsightFiltersAwsAccountId
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) AwsAccountIdInput() *[]*SecurityhubInsightFiltersAwsAccountId {
	var returns *[]*SecurityhubInsightFiltersAwsAccountId
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyName() *[]*SecurityhubInsightFiltersCompanyName {
	var returns *[]*SecurityhubInsightFiltersCompanyName
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CompanyNameInput() *[]*SecurityhubInsightFiltersCompanyName {
	var returns *[]*SecurityhubInsightFiltersCompanyName
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatus() *[]*SecurityhubInsightFiltersComplianceStatus {
	var returns *[]*SecurityhubInsightFiltersComplianceStatus
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ComplianceStatusInput() *[]*SecurityhubInsightFiltersComplianceStatus {
	var returns *[]*SecurityhubInsightFiltersComplianceStatus
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Confidence() *[]*SecurityhubInsightFiltersConfidence {
	var returns *[]*SecurityhubInsightFiltersConfidence
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ConfidenceInput() *[]*SecurityhubInsightFiltersConfidence {
	var returns *[]*SecurityhubInsightFiltersConfidence
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAt() *[]*SecurityhubInsightFiltersCreatedAt {
	var returns *[]*SecurityhubInsightFiltersCreatedAt
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CreatedAtInput() *[]*SecurityhubInsightFiltersCreatedAt {
	var returns *[]*SecurityhubInsightFiltersCreatedAt
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Criticality() *[]*SecurityhubInsightFiltersCriticality {
	var returns *[]*SecurityhubInsightFiltersCriticality
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) CriticalityInput() *[]*SecurityhubInsightFiltersCriticality {
	var returns *[]*SecurityhubInsightFiltersCriticality
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Description() *[]*SecurityhubInsightFiltersDescription {
	var returns *[]*SecurityhubInsightFiltersDescription
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) DescriptionInput() *[]*SecurityhubInsightFiltersDescription {
	var returns *[]*SecurityhubInsightFiltersDescription
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidence() *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidenceInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticality() *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticalityInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsId() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsIdInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArn() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArnInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabel() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabelInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginal() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginalInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypes() *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypesInput() *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes {
	var returns *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes
	_jsii_.Get(
		j,
		"findingProviderFieldsTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAt() *[]*SecurityhubInsightFiltersFirstObservedAt {
	var returns *[]*SecurityhubInsightFiltersFirstObservedAt
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) FirstObservedAtInput() *[]*SecurityhubInsightFiltersFirstObservedAt {
	var returns *[]*SecurityhubInsightFiltersFirstObservedAt
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorId() *[]*SecurityhubInsightFiltersGeneratorId {
	var returns *[]*SecurityhubInsightFiltersGeneratorId
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) GeneratorIdInput() *[]*SecurityhubInsightFiltersGeneratorId {
	var returns *[]*SecurityhubInsightFiltersGeneratorId
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Id() *[]*SecurityhubInsightFiltersId {
	var returns *[]*SecurityhubInsightFiltersId
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) IdInput() *[]*SecurityhubInsightFiltersId {
	var returns *[]*SecurityhubInsightFiltersId
	_jsii_.Get(
		j,
		"idInput",
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

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Keyword() *[]*SecurityhubInsightFiltersKeyword {
	var returns *[]*SecurityhubInsightFiltersKeyword
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) KeywordInput() *[]*SecurityhubInsightFiltersKeyword {
	var returns *[]*SecurityhubInsightFiltersKeyword
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAt() *[]*SecurityhubInsightFiltersLastObservedAt {
	var returns *[]*SecurityhubInsightFiltersLastObservedAt
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) LastObservedAtInput() *[]*SecurityhubInsightFiltersLastObservedAt {
	var returns *[]*SecurityhubInsightFiltersLastObservedAt
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareName() *[]*SecurityhubInsightFiltersMalwareName {
	var returns *[]*SecurityhubInsightFiltersMalwareName
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareNameInput() *[]*SecurityhubInsightFiltersMalwareName {
	var returns *[]*SecurityhubInsightFiltersMalwareName
	_jsii_.Get(
		j,
		"malwareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePath() *[]*SecurityhubInsightFiltersMalwarePath {
	var returns *[]*SecurityhubInsightFiltersMalwarePath
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwarePathInput() *[]*SecurityhubInsightFiltersMalwarePath {
	var returns *[]*SecurityhubInsightFiltersMalwarePath
	_jsii_.Get(
		j,
		"malwarePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareState() *[]*SecurityhubInsightFiltersMalwareState {
	var returns *[]*SecurityhubInsightFiltersMalwareState
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareStateInput() *[]*SecurityhubInsightFiltersMalwareState {
	var returns *[]*SecurityhubInsightFiltersMalwareState
	_jsii_.Get(
		j,
		"malwareStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareType() *[]*SecurityhubInsightFiltersMalwareType {
	var returns *[]*SecurityhubInsightFiltersMalwareType
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) MalwareTypeInput() *[]*SecurityhubInsightFiltersMalwareType {
	var returns *[]*SecurityhubInsightFiltersMalwareType
	_jsii_.Get(
		j,
		"malwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomain() *[]*SecurityhubInsightFiltersNetworkDestinationDomain {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationDomain
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationDomainInput() *[]*SecurityhubInsightFiltersNetworkDestinationDomain {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationDomain
	_jsii_.Get(
		j,
		"networkDestinationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4() *[]*SecurityhubInsightFiltersNetworkDestinationIpv4 {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationIpv4
	_jsii_.Get(
		j,
		"networkDestinationIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv4Input() *[]*SecurityhubInsightFiltersNetworkDestinationIpv4 {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationIpv4
	_jsii_.Get(
		j,
		"networkDestinationIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6() *[]*SecurityhubInsightFiltersNetworkDestinationIpv6 {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationIpv6
	_jsii_.Get(
		j,
		"networkDestinationIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationIpv6Input() *[]*SecurityhubInsightFiltersNetworkDestinationIpv6 {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationIpv6
	_jsii_.Get(
		j,
		"networkDestinationIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPort() *[]*SecurityhubInsightFiltersNetworkDestinationPort {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationPort
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDestinationPortInput() *[]*SecurityhubInsightFiltersNetworkDestinationPort {
	var returns *[]*SecurityhubInsightFiltersNetworkDestinationPort
	_jsii_.Get(
		j,
		"networkDestinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirection() *[]*SecurityhubInsightFiltersNetworkDirection {
	var returns *[]*SecurityhubInsightFiltersNetworkDirection
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkDirectionInput() *[]*SecurityhubInsightFiltersNetworkDirection {
	var returns *[]*SecurityhubInsightFiltersNetworkDirection
	_jsii_.Get(
		j,
		"networkDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocol() *[]*SecurityhubInsightFiltersNetworkProtocol {
	var returns *[]*SecurityhubInsightFiltersNetworkProtocol
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkProtocolInput() *[]*SecurityhubInsightFiltersNetworkProtocol {
	var returns *[]*SecurityhubInsightFiltersNetworkProtocol
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomain() *[]*SecurityhubInsightFiltersNetworkSourceDomain {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceDomain
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceDomainInput() *[]*SecurityhubInsightFiltersNetworkSourceDomain {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceDomain
	_jsii_.Get(
		j,
		"networkSourceDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4() *[]*SecurityhubInsightFiltersNetworkSourceIpv4 {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceIpv4
	_jsii_.Get(
		j,
		"networkSourceIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv4Input() *[]*SecurityhubInsightFiltersNetworkSourceIpv4 {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceIpv4
	_jsii_.Get(
		j,
		"networkSourceIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6() *[]*SecurityhubInsightFiltersNetworkSourceIpv6 {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceIpv6
	_jsii_.Get(
		j,
		"networkSourceIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceIpv6Input() *[]*SecurityhubInsightFiltersNetworkSourceIpv6 {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceIpv6
	_jsii_.Get(
		j,
		"networkSourceIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMac() *[]*SecurityhubInsightFiltersNetworkSourceMac {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceMac
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourceMacInput() *[]*SecurityhubInsightFiltersNetworkSourceMac {
	var returns *[]*SecurityhubInsightFiltersNetworkSourceMac
	_jsii_.Get(
		j,
		"networkSourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePort() *[]*SecurityhubInsightFiltersNetworkSourcePort {
	var returns *[]*SecurityhubInsightFiltersNetworkSourcePort
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NetworkSourcePortInput() *[]*SecurityhubInsightFiltersNetworkSourcePort {
	var returns *[]*SecurityhubInsightFiltersNetworkSourcePort
	_jsii_.Get(
		j,
		"networkSourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteText() *[]*SecurityhubInsightFiltersNoteText {
	var returns *[]*SecurityhubInsightFiltersNoteText
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteTextInput() *[]*SecurityhubInsightFiltersNoteText {
	var returns *[]*SecurityhubInsightFiltersNoteText
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAt() *[]*SecurityhubInsightFiltersNoteUpdatedAt {
	var returns *[]*SecurityhubInsightFiltersNoteUpdatedAt
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedAtInput() *[]*SecurityhubInsightFiltersNoteUpdatedAt {
	var returns *[]*SecurityhubInsightFiltersNoteUpdatedAt
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedBy() *[]*SecurityhubInsightFiltersNoteUpdatedBy {
	var returns *[]*SecurityhubInsightFiltersNoteUpdatedBy
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) NoteUpdatedByInput() *[]*SecurityhubInsightFiltersNoteUpdatedBy {
	var returns *[]*SecurityhubInsightFiltersNoteUpdatedBy
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAt() *[]*SecurityhubInsightFiltersProcessLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersProcessLaunchedAt
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessLaunchedAtInput() *[]*SecurityhubInsightFiltersProcessLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersProcessLaunchedAt
	_jsii_.Get(
		j,
		"processLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessName() *[]*SecurityhubInsightFiltersProcessName {
	var returns *[]*SecurityhubInsightFiltersProcessName
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessNameInput() *[]*SecurityhubInsightFiltersProcessName {
	var returns *[]*SecurityhubInsightFiltersProcessName
	_jsii_.Get(
		j,
		"processNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPid() *[]*SecurityhubInsightFiltersProcessParentPid {
	var returns *[]*SecurityhubInsightFiltersProcessParentPid
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessParentPidInput() *[]*SecurityhubInsightFiltersProcessParentPid {
	var returns *[]*SecurityhubInsightFiltersProcessParentPid
	_jsii_.Get(
		j,
		"processParentPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPath() *[]*SecurityhubInsightFiltersProcessPath {
	var returns *[]*SecurityhubInsightFiltersProcessPath
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPathInput() *[]*SecurityhubInsightFiltersProcessPath {
	var returns *[]*SecurityhubInsightFiltersProcessPath
	_jsii_.Get(
		j,
		"processPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPid() *[]*SecurityhubInsightFiltersProcessPid {
	var returns *[]*SecurityhubInsightFiltersProcessPid
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessPidInput() *[]*SecurityhubInsightFiltersProcessPid {
	var returns *[]*SecurityhubInsightFiltersProcessPid
	_jsii_.Get(
		j,
		"processPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAt() *[]*SecurityhubInsightFiltersProcessTerminatedAt {
	var returns *[]*SecurityhubInsightFiltersProcessTerminatedAt
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProcessTerminatedAtInput() *[]*SecurityhubInsightFiltersProcessTerminatedAt {
	var returns *[]*SecurityhubInsightFiltersProcessTerminatedAt
	_jsii_.Get(
		j,
		"processTerminatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArn() *[]*SecurityhubInsightFiltersProductArn {
	var returns *[]*SecurityhubInsightFiltersProductArn
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductArnInput() *[]*SecurityhubInsightFiltersProductArn {
	var returns *[]*SecurityhubInsightFiltersProductArn
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFields() *[]*SecurityhubInsightFiltersProductFields {
	var returns *[]*SecurityhubInsightFiltersProductFields
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductFieldsInput() *[]*SecurityhubInsightFiltersProductFields {
	var returns *[]*SecurityhubInsightFiltersProductFields
	_jsii_.Get(
		j,
		"productFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductName() *[]*SecurityhubInsightFiltersProductName {
	var returns *[]*SecurityhubInsightFiltersProductName
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ProductNameInput() *[]*SecurityhubInsightFiltersProductName {
	var returns *[]*SecurityhubInsightFiltersProductName
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationText() *[]*SecurityhubInsightFiltersRecommendationText {
	var returns *[]*SecurityhubInsightFiltersRecommendationText
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecommendationTextInput() *[]*SecurityhubInsightFiltersRecommendationText {
	var returns *[]*SecurityhubInsightFiltersRecommendationText
	_jsii_.Get(
		j,
		"recommendationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordState() *[]*SecurityhubInsightFiltersRecordState {
	var returns *[]*SecurityhubInsightFiltersRecordState
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RecordStateInput() *[]*SecurityhubInsightFiltersRecordState {
	var returns *[]*SecurityhubInsightFiltersRecordState
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsId() *[]*SecurityhubInsightFiltersRelatedFindingsId {
	var returns *[]*SecurityhubInsightFiltersRelatedFindingsId
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsIdInput() *[]*SecurityhubInsightFiltersRelatedFindingsId {
	var returns *[]*SecurityhubInsightFiltersRelatedFindingsId
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArn() *[]*SecurityhubInsightFiltersRelatedFindingsProductArn {
	var returns *[]*SecurityhubInsightFiltersRelatedFindingsProductArn
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) RelatedFindingsProductArnInput() *[]*SecurityhubInsightFiltersRelatedFindingsProductArn {
	var returns *[]*SecurityhubInsightFiltersRelatedFindingsProductArn
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArnInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4Addresses() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv4AddressesInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6Addresses() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpv6AddressesInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyName() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyNameInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAt() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAtInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceType() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceTypeInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcId() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcIdInput() *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAt() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAtInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatus() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatusInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserName() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserNameInput() *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerId() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerIdInput() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId {
	var returns *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerName() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerNameInput() *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName {
	var returns *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageId() *[]*SecurityhubInsightFiltersResourceContainerImageId {
	var returns *[]*SecurityhubInsightFiltersResourceContainerImageId
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageIdInput() *[]*SecurityhubInsightFiltersResourceContainerImageId {
	var returns *[]*SecurityhubInsightFiltersResourceContainerImageId
	_jsii_.Get(
		j,
		"resourceContainerImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageName() *[]*SecurityhubInsightFiltersResourceContainerImageName {
	var returns *[]*SecurityhubInsightFiltersResourceContainerImageName
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerImageNameInput() *[]*SecurityhubInsightFiltersResourceContainerImageName {
	var returns *[]*SecurityhubInsightFiltersResourceContainerImageName
	_jsii_.Get(
		j,
		"resourceContainerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAt() *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAtInput() *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt {
	var returns *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerName() *[]*SecurityhubInsightFiltersResourceContainerName {
	var returns *[]*SecurityhubInsightFiltersResourceContainerName
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceContainerNameInput() *[]*SecurityhubInsightFiltersResourceContainerName {
	var returns *[]*SecurityhubInsightFiltersResourceContainerName
	_jsii_.Get(
		j,
		"resourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOther() *[]*SecurityhubInsightFiltersResourceDetailsOther {
	var returns *[]*SecurityhubInsightFiltersResourceDetailsOther
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceDetailsOtherInput() *[]*SecurityhubInsightFiltersResourceDetailsOther {
	var returns *[]*SecurityhubInsightFiltersResourceDetailsOther
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceId() *[]*SecurityhubInsightFiltersResourceId {
	var returns *[]*SecurityhubInsightFiltersResourceId
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceIdInput() *[]*SecurityhubInsightFiltersResourceId {
	var returns *[]*SecurityhubInsightFiltersResourceId
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartition() *[]*SecurityhubInsightFiltersResourcePartition {
	var returns *[]*SecurityhubInsightFiltersResourcePartition
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourcePartitionInput() *[]*SecurityhubInsightFiltersResourcePartition {
	var returns *[]*SecurityhubInsightFiltersResourcePartition
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegion() *[]*SecurityhubInsightFiltersResourceRegion {
	var returns *[]*SecurityhubInsightFiltersResourceRegion
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceRegionInput() *[]*SecurityhubInsightFiltersResourceRegion {
	var returns *[]*SecurityhubInsightFiltersResourceRegion
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTags() *[]*SecurityhubInsightFiltersResourceTags {
	var returns *[]*SecurityhubInsightFiltersResourceTags
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTagsInput() *[]*SecurityhubInsightFiltersResourceTags {
	var returns *[]*SecurityhubInsightFiltersResourceTags
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceType() *[]*SecurityhubInsightFiltersResourceType {
	var returns *[]*SecurityhubInsightFiltersResourceType
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ResourceTypeInput() *[]*SecurityhubInsightFiltersResourceType {
	var returns *[]*SecurityhubInsightFiltersResourceType
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabel() *[]*SecurityhubInsightFiltersSeverityLabel {
	var returns *[]*SecurityhubInsightFiltersSeverityLabel
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SeverityLabelInput() *[]*SecurityhubInsightFiltersSeverityLabel {
	var returns *[]*SecurityhubInsightFiltersSeverityLabel
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrl() *[]*SecurityhubInsightFiltersSourceUrl {
	var returns *[]*SecurityhubInsightFiltersSourceUrl
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SourceUrlInput() *[]*SecurityhubInsightFiltersSourceUrl {
	var returns *[]*SecurityhubInsightFiltersSourceUrl
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

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategory() *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategoryInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAt() *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAtInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSource() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrl() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrlInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorType() *[]*SecurityhubInsightFiltersThreatIntelIndicatorType {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorType
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorTypeInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorType {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorType
	_jsii_.Get(
		j,
		"threatIntelIndicatorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValue() *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValueInput() *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue {
	var returns *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue
	_jsii_.Get(
		j,
		"threatIntelIndicatorValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Title() *[]*SecurityhubInsightFiltersTitle {
	var returns *[]*SecurityhubInsightFiltersTitle
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TitleInput() *[]*SecurityhubInsightFiltersTitle {
	var returns *[]*SecurityhubInsightFiltersTitle
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) Type() *[]*SecurityhubInsightFiltersType {
	var returns *[]*SecurityhubInsightFiltersType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) TypeInput() *[]*SecurityhubInsightFiltersType {
	var returns *[]*SecurityhubInsightFiltersType
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAt() *[]*SecurityhubInsightFiltersUpdatedAt {
	var returns *[]*SecurityhubInsightFiltersUpdatedAt
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UpdatedAtInput() *[]*SecurityhubInsightFiltersUpdatedAt {
	var returns *[]*SecurityhubInsightFiltersUpdatedAt
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValues() *[]*SecurityhubInsightFiltersUserDefinedValues {
	var returns *[]*SecurityhubInsightFiltersUserDefinedValues
	_jsii_.Get(
		j,
		"userDefinedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) UserDefinedValuesInput() *[]*SecurityhubInsightFiltersUserDefinedValues {
	var returns *[]*SecurityhubInsightFiltersUserDefinedValues
	_jsii_.Get(
		j,
		"userDefinedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationState() *[]*SecurityhubInsightFiltersVerificationState {
	var returns *[]*SecurityhubInsightFiltersVerificationState
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) VerificationStateInput() *[]*SecurityhubInsightFiltersVerificationState {
	var returns *[]*SecurityhubInsightFiltersVerificationState
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatus() *[]*SecurityhubInsightFiltersWorkflowStatus {
	var returns *[]*SecurityhubInsightFiltersWorkflowStatus
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) WorkflowStatusInput() *[]*SecurityhubInsightFiltersWorkflowStatus {
	var returns *[]*SecurityhubInsightFiltersWorkflowStatus
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


func NewSecurityhubInsightFiltersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersOutputReference_Override(s SecurityhubInsightFiltersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetAwsAccountId(val *[]*SecurityhubInsightFiltersAwsAccountId) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCompanyName(val *[]*SecurityhubInsightFiltersCompanyName) {
	_jsii_.Set(
		j,
		"companyName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetComplianceStatus(val *[]*SecurityhubInsightFiltersComplianceStatus) {
	_jsii_.Set(
		j,
		"complianceStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetConfidence(val *[]*SecurityhubInsightFiltersConfidence) {
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCreatedAt(val *[]*SecurityhubInsightFiltersCreatedAt) {
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetCriticality(val *[]*SecurityhubInsightFiltersCriticality) {
	_jsii_.Set(
		j,
		"criticality",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetDescription(val *[]*SecurityhubInsightFiltersDescription) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsConfidence(val *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence) {
	_jsii_.Set(
		j,
		"findingProviderFieldsConfidence",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsCriticality(val *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality) {
	_jsii_.Set(
		j,
		"findingProviderFieldsCriticality",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsRelatedFindingsId(val *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId) {
	_jsii_.Set(
		j,
		"findingProviderFieldsRelatedFindingsId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsRelatedFindingsProductArn(val *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn) {
	_jsii_.Set(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsSeverityLabel(val *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel) {
	_jsii_.Set(
		j,
		"findingProviderFieldsSeverityLabel",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsSeverityOriginal(val *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal) {
	_jsii_.Set(
		j,
		"findingProviderFieldsSeverityOriginal",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFindingProviderFieldsTypes(val *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes) {
	_jsii_.Set(
		j,
		"findingProviderFieldsTypes",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetFirstObservedAt(val *[]*SecurityhubInsightFiltersFirstObservedAt) {
	_jsii_.Set(
		j,
		"firstObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetGeneratorId(val *[]*SecurityhubInsightFiltersGeneratorId) {
	_jsii_.Set(
		j,
		"generatorId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetId(val *[]*SecurityhubInsightFiltersId) {
	_jsii_.Set(
		j,
		"id",
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

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetKeyword(val *[]*SecurityhubInsightFiltersKeyword) {
	_jsii_.Set(
		j,
		"keyword",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetLastObservedAt(val *[]*SecurityhubInsightFiltersLastObservedAt) {
	_jsii_.Set(
		j,
		"lastObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareName(val *[]*SecurityhubInsightFiltersMalwareName) {
	_jsii_.Set(
		j,
		"malwareName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwarePath(val *[]*SecurityhubInsightFiltersMalwarePath) {
	_jsii_.Set(
		j,
		"malwarePath",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareState(val *[]*SecurityhubInsightFiltersMalwareState) {
	_jsii_.Set(
		j,
		"malwareState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetMalwareType(val *[]*SecurityhubInsightFiltersMalwareType) {
	_jsii_.Set(
		j,
		"malwareType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationDomain(val *[]*SecurityhubInsightFiltersNetworkDestinationDomain) {
	_jsii_.Set(
		j,
		"networkDestinationDomain",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationIpv4(val *[]*SecurityhubInsightFiltersNetworkDestinationIpv4) {
	_jsii_.Set(
		j,
		"networkDestinationIpv4",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationIpv6(val *[]*SecurityhubInsightFiltersNetworkDestinationIpv6) {
	_jsii_.Set(
		j,
		"networkDestinationIpv6",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDestinationPort(val *[]*SecurityhubInsightFiltersNetworkDestinationPort) {
	_jsii_.Set(
		j,
		"networkDestinationPort",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkDirection(val *[]*SecurityhubInsightFiltersNetworkDirection) {
	_jsii_.Set(
		j,
		"networkDirection",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkProtocol(val *[]*SecurityhubInsightFiltersNetworkProtocol) {
	_jsii_.Set(
		j,
		"networkProtocol",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceDomain(val *[]*SecurityhubInsightFiltersNetworkSourceDomain) {
	_jsii_.Set(
		j,
		"networkSourceDomain",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceIpv4(val *[]*SecurityhubInsightFiltersNetworkSourceIpv4) {
	_jsii_.Set(
		j,
		"networkSourceIpv4",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceIpv6(val *[]*SecurityhubInsightFiltersNetworkSourceIpv6) {
	_jsii_.Set(
		j,
		"networkSourceIpv6",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourceMac(val *[]*SecurityhubInsightFiltersNetworkSourceMac) {
	_jsii_.Set(
		j,
		"networkSourceMac",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNetworkSourcePort(val *[]*SecurityhubInsightFiltersNetworkSourcePort) {
	_jsii_.Set(
		j,
		"networkSourcePort",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteText(val *[]*SecurityhubInsightFiltersNoteText) {
	_jsii_.Set(
		j,
		"noteText",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteUpdatedAt(val *[]*SecurityhubInsightFiltersNoteUpdatedAt) {
	_jsii_.Set(
		j,
		"noteUpdatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetNoteUpdatedBy(val *[]*SecurityhubInsightFiltersNoteUpdatedBy) {
	_jsii_.Set(
		j,
		"noteUpdatedBy",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessLaunchedAt(val *[]*SecurityhubInsightFiltersProcessLaunchedAt) {
	_jsii_.Set(
		j,
		"processLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessName(val *[]*SecurityhubInsightFiltersProcessName) {
	_jsii_.Set(
		j,
		"processName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessParentPid(val *[]*SecurityhubInsightFiltersProcessParentPid) {
	_jsii_.Set(
		j,
		"processParentPid",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessPath(val *[]*SecurityhubInsightFiltersProcessPath) {
	_jsii_.Set(
		j,
		"processPath",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessPid(val *[]*SecurityhubInsightFiltersProcessPid) {
	_jsii_.Set(
		j,
		"processPid",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProcessTerminatedAt(val *[]*SecurityhubInsightFiltersProcessTerminatedAt) {
	_jsii_.Set(
		j,
		"processTerminatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductArn(val *[]*SecurityhubInsightFiltersProductArn) {
	_jsii_.Set(
		j,
		"productArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductFields(val *[]*SecurityhubInsightFiltersProductFields) {
	_jsii_.Set(
		j,
		"productFields",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetProductName(val *[]*SecurityhubInsightFiltersProductName) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRecommendationText(val *[]*SecurityhubInsightFiltersRecommendationText) {
	_jsii_.Set(
		j,
		"recommendationText",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRecordState(val *[]*SecurityhubInsightFiltersRecordState) {
	_jsii_.Set(
		j,
		"recordState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRelatedFindingsId(val *[]*SecurityhubInsightFiltersRelatedFindingsId) {
	_jsii_.Set(
		j,
		"relatedFindingsId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetRelatedFindingsProductArn(val *[]*SecurityhubInsightFiltersRelatedFindingsProductArn) {
	_jsii_.Set(
		j,
		"relatedFindingsProductArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIamInstanceProfileArn(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceImageId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceImageId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIpv4Addresses(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceIpv6Addresses(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceKeyName(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceKeyName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceLaunchedAt(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceSubnetId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceSubnetId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceType(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsEc2InstanceVpcId(val *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId) {
	_jsii_.Set(
		j,
		"resourceAwsEc2InstanceVpcId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyCreatedAt(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyStatus(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsIamAccessKeyUserName(val *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName) {
	_jsii_.Set(
		j,
		"resourceAwsIamAccessKeyUserName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsS3BucketOwnerId(val *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId) {
	_jsii_.Set(
		j,
		"resourceAwsS3BucketOwnerId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceAwsS3BucketOwnerName(val *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName) {
	_jsii_.Set(
		j,
		"resourceAwsS3BucketOwnerName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerImageId(val *[]*SecurityhubInsightFiltersResourceContainerImageId) {
	_jsii_.Set(
		j,
		"resourceContainerImageId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerImageName(val *[]*SecurityhubInsightFiltersResourceContainerImageName) {
	_jsii_.Set(
		j,
		"resourceContainerImageName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerLaunchedAt(val *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt) {
	_jsii_.Set(
		j,
		"resourceContainerLaunchedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceContainerName(val *[]*SecurityhubInsightFiltersResourceContainerName) {
	_jsii_.Set(
		j,
		"resourceContainerName",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceDetailsOther(val *[]*SecurityhubInsightFiltersResourceDetailsOther) {
	_jsii_.Set(
		j,
		"resourceDetailsOther",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceId(val *[]*SecurityhubInsightFiltersResourceId) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourcePartition(val *[]*SecurityhubInsightFiltersResourcePartition) {
	_jsii_.Set(
		j,
		"resourcePartition",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceRegion(val *[]*SecurityhubInsightFiltersResourceRegion) {
	_jsii_.Set(
		j,
		"resourceRegion",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceTags(val *[]*SecurityhubInsightFiltersResourceTags) {
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetResourceType(val *[]*SecurityhubInsightFiltersResourceType) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetSeverityLabel(val *[]*SecurityhubInsightFiltersSeverityLabel) {
	_jsii_.Set(
		j,
		"severityLabel",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetSourceUrl(val *[]*SecurityhubInsightFiltersSourceUrl) {
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

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorCategory(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorCategory",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorLastObservedAt(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorLastObservedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorSource(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorSource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorSourceUrl(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorSourceUrl",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorType(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorType) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorType",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetThreatIntelIndicatorValue(val *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue) {
	_jsii_.Set(
		j,
		"threatIntelIndicatorValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetTitle(val *[]*SecurityhubInsightFiltersTitle) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetType(val *[]*SecurityhubInsightFiltersType) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetUpdatedAt(val *[]*SecurityhubInsightFiltersUpdatedAt) {
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetUserDefinedValues(val *[]*SecurityhubInsightFiltersUserDefinedValues) {
	_jsii_.Set(
		j,
		"userDefinedValues",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetVerificationState(val *[]*SecurityhubInsightFiltersVerificationState) {
	_jsii_.Set(
		j,
		"verificationState",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) SetWorkflowStatus(val *[]*SecurityhubInsightFiltersWorkflowStatus) {
	_jsii_.Set(
		j,
		"workflowStatus",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersProcessLaunchedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersProcessLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersProcessParentPid struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersProcessPath struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersProcessPid struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#eq SecurityhubInsight#eq}.
	Eq *string `json:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#gte SecurityhubInsight#gte}.
	Gte *string `json:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#lte SecurityhubInsight#lte}.
	Lte *string `json:"lte"`
}

type SecurityhubInsightFiltersProcessTerminatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersProcessTerminatedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersProcessTerminatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersProductFields struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#key SecurityhubInsight#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersProductName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersRecommendationText struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersRecordState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersRelatedFindingsId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersRelatedFindingsProductArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceImageId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#cidr SecurityhubInsight#cidr}.
	Cidr *string `json:"cidr"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsS3BucketOwnerId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceAwsS3BucketOwnerName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceContainerImageId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceContainerImageName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceDetailsOther struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#key SecurityhubInsight#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourcePartition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceRegion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#key SecurityhubInsight#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersResourceType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersSeverityLabel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersSourceUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorCategory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersThreatIntelIndicatorValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersTitle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersUpdatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersUpdatedAtDateRange `json:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#end SecurityhubInsight#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#start SecurityhubInsight#start}.
	Start *string `json:"start"`
}

type SecurityhubInsightFiltersUpdatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#unit SecurityhubInsight#unit}.
	Unit *string `json:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *float64 `json:"value"`
}

type SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference
type jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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


func NewSecurityhubInsightFiltersUpdatedAtDateRangeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersUpdatedAtDateRangeOutputReference_Override(s SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
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

func (j *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (s *jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#key SecurityhubInsight#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersVerificationState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

type SecurityhubInsightFiltersWorkflowStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#comparison SecurityhubInsight#comparison}.
	Comparison *string `json:"comparison"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight.html#value SecurityhubInsight#value}.
	Value *string `json:"value"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter.html aws_securityhub_invite_accepter}.
type SecurityhubInviteAccepter interface {
	cdktf.TerraformResource
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

func (j *jsiiProxy_SecurityhubInviteAccepter) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter.html aws_securityhub_invite_accepter} Resource.
func NewSecurityhubInviteAccepter(scope constructs.Construct, id *string, config *SecurityhubInviteAccepterConfig) SecurityhubInviteAccepter {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubInviteAccepter{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInviteAccepter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter.html aws_securityhub_invite_accepter} Resource.
func NewSecurityhubInviteAccepter_Override(s SecurityhubInviteAccepter, scope constructs.Construct, id *string, config *SecurityhubInviteAccepterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubInviteAccepter",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInviteAccepter) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubInviteAccepter",
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
		"hashicorp_aws.SecurityHub.SecurityhubInviteAccepter",
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

type SecurityhubInviteAccepterConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_invite_accepter.html#master_id SecurityhubInviteAccepter#master_id}.
	MasterId *string `json:"masterId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html aws_securityhub_member}.
type SecurityhubMember interface {
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_SecurityhubMember) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html aws_securityhub_member} Resource.
func NewSecurityhubMember(scope constructs.Construct, id *string, config *SecurityhubMemberConfig) SecurityhubMember {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubMember{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubMember",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html aws_securityhub_member} Resource.
func NewSecurityhubMember_Override(s SecurityhubMember, scope constructs.Construct, id *string, config *SecurityhubMemberConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubMember",
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

func (j *jsiiProxy_SecurityhubMember) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubMember",
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
		"hashicorp_aws.SecurityHub.SecurityhubMember",
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

type SecurityhubMemberConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html#account_id SecurityhubMember#account_id}.
	AccountId *string `json:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html#email SecurityhubMember#email}.
	Email *string `json:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_member.html#invite SecurityhubMember#invite}.
	Invite interface{} `json:"invite"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account.html aws_securityhub_organization_admin_account}.
type SecurityhubOrganizationAdminAccount interface {
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

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account.html aws_securityhub_organization_admin_account} Resource.
func NewSecurityhubOrganizationAdminAccount(scope constructs.Construct, id *string, config *SecurityhubOrganizationAdminAccountConfig) SecurityhubOrganizationAdminAccount {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubOrganizationAdminAccount{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationAdminAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account.html aws_securityhub_organization_admin_account} Resource.
func NewSecurityhubOrganizationAdminAccount_Override(s SecurityhubOrganizationAdminAccount, scope constructs.Construct, id *string, config *SecurityhubOrganizationAdminAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationAdminAccount",
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

func (j *jsiiProxy_SecurityhubOrganizationAdminAccount) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationAdminAccount",
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
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationAdminAccount",
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

type SecurityhubOrganizationAdminAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_admin_account.html#admin_account_id SecurityhubOrganizationAdminAccount#admin_account_id}.
	AdminAccountId *string `json:"adminAccountId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration.html aws_securityhub_organization_configuration}.
type SecurityhubOrganizationConfiguration interface {
	cdktf.TerraformResource
	AutoEnable() interface{}
	SetAutoEnable(val interface{})
	AutoEnableInput() interface{}
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

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration.html aws_securityhub_organization_configuration} Resource.
func NewSecurityhubOrganizationConfiguration(scope constructs.Construct, id *string, config *SecurityhubOrganizationConfigurationConfig) SecurityhubOrganizationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubOrganizationConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration.html aws_securityhub_organization_configuration} Resource.
func NewSecurityhubOrganizationConfiguration_Override(s SecurityhubOrganizationConfiguration, scope constructs.Construct, id *string, config *SecurityhubOrganizationConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationConfiguration",
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

func (j *jsiiProxy_SecurityhubOrganizationConfiguration) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationConfiguration",
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
		"hashicorp_aws.SecurityHub.SecurityhubOrganizationConfiguration",
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

type SecurityhubOrganizationConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_organization_configuration.html#auto_enable SecurityhubOrganizationConfiguration#auto_enable}.
	AutoEnable interface{} `json:"autoEnable"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription.html aws_securityhub_product_subscription}.
type SecurityhubProductSubscription interface {
	cdktf.TerraformResource
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

func (j *jsiiProxy_SecurityhubProductSubscription) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription.html aws_securityhub_product_subscription} Resource.
func NewSecurityhubProductSubscription(scope constructs.Construct, id *string, config *SecurityhubProductSubscriptionConfig) SecurityhubProductSubscription {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubProductSubscription{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubProductSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription.html aws_securityhub_product_subscription} Resource.
func NewSecurityhubProductSubscription_Override(s SecurityhubProductSubscription, scope constructs.Construct, id *string, config *SecurityhubProductSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubProductSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubProductSubscription) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubProductSubscription",
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
		"hashicorp_aws.SecurityHub.SecurityhubProductSubscription",
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

type SecurityhubProductSubscriptionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_product_subscription.html#product_arn SecurityhubProductSubscription#product_arn}.
	ProductArn *string `json:"productArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html aws_securityhub_standards_control}.
type SecurityhubStandardsControl interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ControlId() *string
	ControlStatus() *string
	SetControlStatus(val *string)
	ControlStatusInput() *string
	ControlStatusUpdatedAt() *string
	Count() interface{}
	SetCount(val interface{})
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_SecurityhubStandardsControl) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html aws_securityhub_standards_control} Resource.
func NewSecurityhubStandardsControl(scope constructs.Construct, id *string, config *SecurityhubStandardsControlConfig) SecurityhubStandardsControl {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubStandardsControl{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubStandardsControl",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html aws_securityhub_standards_control} Resource.
func NewSecurityhubStandardsControl_Override(s SecurityhubStandardsControl, scope constructs.Construct, id *string, config *SecurityhubStandardsControlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubStandardsControl",
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

func (j *jsiiProxy_SecurityhubStandardsControl) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubStandardsControl",
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
		"hashicorp_aws.SecurityHub.SecurityhubStandardsControl",
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

type SecurityhubStandardsControlConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html#control_status SecurityhubStandardsControl#control_status}.
	ControlStatus *string `json:"controlStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html#standards_control_arn SecurityhubStandardsControl#standards_control_arn}.
	StandardsControlArn *string `json:"standardsControlArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_control.html#disabled_reason SecurityhubStandardsControl#disabled_reason}.
	DisabledReason *string `json:"disabledReason"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription.html aws_securityhub_standards_subscription}.
type SecurityhubStandardsSubscription interface {
	cdktf.TerraformResource
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
	StandardsArn() *string
	SetStandardsArn(val *string)
	StandardsArnInput() *string
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

func (j *jsiiProxy_SecurityhubStandardsSubscription) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription.html aws_securityhub_standards_subscription} Resource.
func NewSecurityhubStandardsSubscription(scope constructs.Construct, id *string, config *SecurityhubStandardsSubscriptionConfig) SecurityhubStandardsSubscription {
	_init_.Initialize()

	j := jsiiProxy_SecurityhubStandardsSubscription{}

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubStandardsSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription.html aws_securityhub_standards_subscription} Resource.
func NewSecurityhubStandardsSubscription_Override(s SecurityhubStandardsSubscription, scope constructs.Construct, id *string, config *SecurityhubStandardsSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.SecurityHub.SecurityhubStandardsSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityhubStandardsSubscription) SetCount(val interface{}) {
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
		"hashicorp_aws.SecurityHub.SecurityhubStandardsSubscription",
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
		"hashicorp_aws.SecurityHub.SecurityhubStandardsSubscription",
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

type SecurityhubStandardsSubscriptionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_standards_subscription.html#standards_arn SecurityhubStandardsSubscription#standards_arn}.
	StandardsArn *string `json:"standardsArn"`
}

