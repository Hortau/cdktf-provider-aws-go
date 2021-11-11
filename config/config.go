package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/config/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html aws_config_aggregate_authorization}.
type ConfigAggregateAuthorization interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigAggregateAuthorization
type jsiiProxy_ConfigAggregateAuthorization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigAggregateAuthorization) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigAggregateAuthorization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html aws_config_aggregate_authorization} Resource.
func NewConfigAggregateAuthorization(scope constructs.Construct, id *string, config *ConfigAggregateAuthorizationConfig) ConfigAggregateAuthorization {
	_init_.Initialize()

	j := jsiiProxy_ConfigAggregateAuthorization{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigAggregateAuthorization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html aws_config_aggregate_authorization} Resource.
func NewConfigAggregateAuthorization_Override(c ConfigAggregateAuthorization, scope constructs.Construct, id *string, config *ConfigAggregateAuthorizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigAggregateAuthorization",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ConfigAggregateAuthorization) SetTagsAll(val interface{}) {
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
func ConfigAggregateAuthorization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigAggregateAuthorization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigAggregateAuthorization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigAggregateAuthorization",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigAggregateAuthorization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigAggregateAuthorization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigAggregateAuthorization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigAggregateAuthorization) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigAggregateAuthorization) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigAggregateAuthorization) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) ToString() *string {
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
func (c *jsiiProxy_ConfigAggregateAuthorization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigAggregateAuthorizationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html#account_id ConfigAggregateAuthorization#account_id}.
	AccountId *string `json:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html#region ConfigAggregateAuthorization#region}.
	Region *string `json:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html#tags ConfigAggregateAuthorization#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_aggregate_authorization.html#tags_all ConfigAggregateAuthorization#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html aws_config_config_rule}.
type ConfigConfigRule interface {
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
	InputParameters() *string
	SetInputParameters(val *string)
	InputParametersInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	MaximumExecutionFrequencyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RuleId() *string
	Scope() ConfigConfigRuleScopeOutputReference
	ScopeInput() *ConfigConfigRuleScope
	Source() ConfigConfigRuleSourceOutputReference
	SourceInput() *ConfigConfigRuleSource
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
	PutScope(value *ConfigConfigRuleScope)
	PutSource(value *ConfigConfigRuleSource)
	ResetDescription()
	ResetInputParameters()
	ResetMaximumExecutionFrequency()
	ResetOverrideLogicalId()
	ResetScope()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigConfigRule
type jsiiProxy_ConfigConfigRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigConfigRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Scope() ConfigConfigRuleScopeOutputReference {
	var returns ConfigConfigRuleScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) ScopeInput() *ConfigConfigRuleScope {
	var returns *ConfigConfigRuleScope
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Source() ConfigConfigRuleSourceOutputReference {
	var returns ConfigConfigRuleSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) SourceInput() *ConfigConfigRuleSource {
	var returns *ConfigConfigRuleSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html aws_config_config_rule} Resource.
func NewConfigConfigRule(scope constructs.Construct, id *string, config *ConfigConfigRuleConfig) ConfigConfigRule {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigRule{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html aws_config_config_rule} Resource.
func NewConfigConfigRule_Override(c ConfigConfigRule, scope constructs.Construct, id *string, config *ConfigConfigRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetInputParameters(val *string) {
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetMaximumExecutionFrequency(val *string) {
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRule) SetTagsAll(val interface{}) {
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
func ConfigConfigRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigConfigRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigConfigRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigConfigRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigRule) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigRule) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigConfigRule) PutScope(value *ConfigConfigRuleScope) {
	_jsii_.InvokeVoid(
		c,
		"putScope",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigRule) PutSource(value *ConfigConfigRuleSource) {
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigConfigRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetScope() {
	_jsii_.InvokeVoid(
		c,
		"resetScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRule) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigConfigRule) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigConfigRule) ToString() *string {
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
func (c *jsiiProxy_ConfigConfigRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigConfigRuleConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#name ConfigConfigRule#name}.
	Name *string `json:"name"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#source ConfigConfigRule#source}
	Source *ConfigConfigRuleSource `json:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#description ConfigConfigRule#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#input_parameters ConfigConfigRule#input_parameters}.
	InputParameters *string `json:"inputParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#scope ConfigConfigRule#scope}
	Scope *ConfigConfigRuleScope `json:"scope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#tags ConfigConfigRule#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#tags_all ConfigConfigRule#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type ConfigConfigRuleScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#compliance_resource_id ConfigConfigRule#compliance_resource_id}.
	ComplianceResourceId *string `json:"complianceResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#compliance_resource_types ConfigConfigRule#compliance_resource_types}.
	ComplianceResourceTypes *[]*string `json:"complianceResourceTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#tag_key ConfigConfigRule#tag_key}.
	TagKey *string `json:"tagKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#tag_value ConfigConfigRule#tag_value}.
	TagValue *string `json:"tagValue"`
}

type ConfigConfigRuleScopeOutputReference interface {
	cdktf.ComplexObject
	ComplianceResourceId() *string
	SetComplianceResourceId(val *string)
	ComplianceResourceIdInput() *string
	ComplianceResourceTypes() *[]*string
	SetComplianceResourceTypes(val *[]*string)
	ComplianceResourceTypesInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TagKey() *string
	SetTagKey(val *string)
	TagKeyInput() *string
	TagValue() *string
	SetTagValue(val *string)
	TagValueInput() *string
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
	ResetComplianceResourceId()
	ResetComplianceResourceTypes()
	ResetTagKey()
	ResetTagValue()
}

// The jsii proxy struct for ConfigConfigRuleScopeOutputReference
type jsiiProxy_ConfigConfigRuleScopeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) ComplianceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) ComplianceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) ComplianceResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"complianceResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) ComplianceResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"complianceResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigConfigRuleScopeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigConfigRuleScopeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigRuleScopeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRuleScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigConfigRuleScopeOutputReference_Override(c ConfigConfigRuleScopeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRuleScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetComplianceResourceId(val *string) {
	_jsii_.Set(
		j,
		"complianceResourceId",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetComplianceResourceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"complianceResourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetTagKey(val *string) {
	_jsii_.Set(
		j,
		"tagKey",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetTagValue(val *string) {
	_jsii_.Set(
		j,
		"tagValue",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleScopeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) ResetComplianceResourceId() {
	_jsii_.InvokeVoid(
		c,
		"resetComplianceResourceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) ResetComplianceResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetComplianceResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) ResetTagKey() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRuleScopeOutputReference) ResetTagValue() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValue",
		nil, // no parameters
	)
}

type ConfigConfigRuleSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#owner ConfigConfigRule#owner}.
	Owner *string `json:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#source_identifier ConfigConfigRule#source_identifier}.
	SourceIdentifier *string `json:"sourceIdentifier"`
	// source_detail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#source_detail ConfigConfigRule#source_detail}
	SourceDetail *[]*ConfigConfigRuleSourceSourceDetail `json:"sourceDetail"`
}

type ConfigConfigRuleSourceOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	SourceDetail() *[]*ConfigConfigRuleSourceSourceDetail
	SetSourceDetail(val *[]*ConfigConfigRuleSourceSourceDetail)
	SourceDetailInput() *[]*ConfigConfigRuleSourceSourceDetail
	SourceIdentifier() *string
	SetSourceIdentifier(val *string)
	SourceIdentifierInput() *string
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
	ResetSourceDetail()
}

// The jsii proxy struct for ConfigConfigRuleSourceOutputReference
type jsiiProxy_ConfigConfigRuleSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SourceDetail() *[]*ConfigConfigRuleSourceSourceDetail {
	var returns *[]*ConfigConfigRuleSourceSourceDetail
	_jsii_.Get(
		j,
		"sourceDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SourceDetailInput() *[]*ConfigConfigRuleSourceSourceDetail {
	var returns *[]*ConfigConfigRuleSourceSourceDetail
	_jsii_.Get(
		j,
		"sourceDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigConfigRuleSourceOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigConfigRuleSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigRuleSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigConfigRuleSourceOutputReference_Override(c ConfigConfigRuleSourceOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetSourceDetail(val *[]*ConfigConfigRuleSourceSourceDetail) {
	_jsii_.Set(
		j,
		"sourceDetail",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetSourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceOutputReference) ResetSourceDetail() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDetail",
		nil, // no parameters
	)
}

type ConfigConfigRuleSourceSourceDetail struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#event_source ConfigConfigRule#event_source}.
	EventSource *string `json:"eventSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_config_rule.html#message_type ConfigConfigRule#message_type}.
	MessageType *string `json:"messageType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html aws_config_configuration_aggregator}.
type ConfigConfigurationAggregator interface {
	cdktf.TerraformResource
	AccountAggregationSource() ConfigConfigurationAggregatorAccountAggregationSourceOutputReference
	AccountAggregationSourceInput() *ConfigConfigurationAggregatorAccountAggregationSource
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
	OrganizationAggregationSource() ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference
	OrganizationAggregationSourceInput() *ConfigConfigurationAggregatorOrganizationAggregationSource
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
	PutAccountAggregationSource(value *ConfigConfigurationAggregatorAccountAggregationSource)
	PutOrganizationAggregationSource(value *ConfigConfigurationAggregatorOrganizationAggregationSource)
	ResetAccountAggregationSource()
	ResetOrganizationAggregationSource()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigConfigurationAggregator
type jsiiProxy_ConfigConfigurationAggregator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigConfigurationAggregator) AccountAggregationSource() ConfigConfigurationAggregatorAccountAggregationSourceOutputReference {
	var returns ConfigConfigurationAggregatorAccountAggregationSourceOutputReference
	_jsii_.Get(
		j,
		"accountAggregationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) AccountAggregationSourceInput() *ConfigConfigurationAggregatorAccountAggregationSource {
	var returns *ConfigConfigurationAggregatorAccountAggregationSource
	_jsii_.Get(
		j,
		"accountAggregationSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) OrganizationAggregationSource() ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference {
	var returns ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference
	_jsii_.Get(
		j,
		"organizationAggregationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) OrganizationAggregationSourceInput() *ConfigConfigurationAggregatorOrganizationAggregationSource {
	var returns *ConfigConfigurationAggregatorOrganizationAggregationSource
	_jsii_.Get(
		j,
		"organizationAggregationSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html aws_config_configuration_aggregator} Resource.
func NewConfigConfigurationAggregator(scope constructs.Construct, id *string, config *ConfigConfigurationAggregatorConfig) ConfigConfigurationAggregator {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationAggregator{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html aws_config_configuration_aggregator} Resource.
func NewConfigConfigurationAggregator_Override(c ConfigConfigurationAggregator, scope constructs.Construct, id *string, config *ConfigConfigurationAggregatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregator",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregator) SetTagsAll(val interface{}) {
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
func ConfigConfigurationAggregator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigConfigurationAggregator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigConfigurationAggregator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigConfigurationAggregator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) PutAccountAggregationSource(value *ConfigConfigurationAggregatorAccountAggregationSource) {
	_jsii_.InvokeVoid(
		c,
		"putAccountAggregationSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) PutOrganizationAggregationSource(value *ConfigConfigurationAggregatorOrganizationAggregationSource) {
	_jsii_.InvokeVoid(
		c,
		"putOrganizationAggregationSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) ResetAccountAggregationSource() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountAggregationSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) ResetOrganizationAggregationSource() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganizationAggregationSource",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregator) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) ToString() *string {
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
func (c *jsiiProxy_ConfigConfigurationAggregator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigConfigurationAggregatorAccountAggregationSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#account_ids ConfigConfigurationAggregator#account_ids}.
	AccountIds *[]*string `json:"accountIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#all_regions ConfigConfigurationAggregator#all_regions}.
	AllRegions interface{} `json:"allRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#regions ConfigConfigurationAggregator#regions}.
	Regions *[]*string `json:"regions"`
}

type ConfigConfigurationAggregatorAccountAggregationSourceOutputReference interface {
	cdktf.ComplexObject
	AccountIds() *[]*string
	SetAccountIds(val *[]*string)
	AccountIdsInput() *[]*string
	AllRegions() interface{}
	SetAllRegions(val interface{})
	AllRegionsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
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
	ResetAllRegions()
	ResetRegions()
}

// The jsii proxy struct for ConfigConfigurationAggregatorAccountAggregationSourceOutputReference
type jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) AccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) AccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) AllRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) AllRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigConfigurationAggregatorAccountAggregationSourceOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigConfigurationAggregatorAccountAggregationSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregatorAccountAggregationSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigConfigurationAggregatorAccountAggregationSourceOutputReference_Override(c ConfigConfigurationAggregatorAccountAggregationSourceOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregatorAccountAggregationSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"accountIds",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetAllRegions(val interface{}) {
	_jsii_.Set(
		j,
		"allRegions",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) ResetAllRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetAllRegions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregatorAccountAggregationSourceOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetRegions",
		nil, // no parameters
	)
}

type ConfigConfigurationAggregatorConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#name ConfigConfigurationAggregator#name}.
	Name *string `json:"name"`
	// account_aggregation_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#account_aggregation_source ConfigConfigurationAggregator#account_aggregation_source}
	AccountAggregationSource *ConfigConfigurationAggregatorAccountAggregationSource `json:"accountAggregationSource"`
	// organization_aggregation_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#organization_aggregation_source ConfigConfigurationAggregator#organization_aggregation_source}
	OrganizationAggregationSource *ConfigConfigurationAggregatorOrganizationAggregationSource `json:"organizationAggregationSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#tags ConfigConfigurationAggregator#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#tags_all ConfigConfigurationAggregator#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type ConfigConfigurationAggregatorOrganizationAggregationSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#role_arn ConfigConfigurationAggregator#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#all_regions ConfigConfigurationAggregator#all_regions}.
	AllRegions interface{} `json:"allRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_aggregator.html#regions ConfigConfigurationAggregator#regions}.
	Regions *[]*string `json:"regions"`
}

type ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference interface {
	cdktf.ComplexObject
	AllRegions() interface{}
	SetAllRegions(val interface{})
	AllRegionsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetAllRegions()
	ResetRegions()
}

// The jsii proxy struct for ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference
type jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) AllRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) AllRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference_Override(c ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetAllRegions(val interface{}) {
	_jsii_.Set(
		j,
		"allRegions",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) ResetAllRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetAllRegions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationAggregatorOrganizationAggregationSourceOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetRegions",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html aws_config_configuration_recorder}.
type ConfigConfigurationRecorder interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RecordingGroup() ConfigConfigurationRecorderRecordingGroupOutputReference
	RecordingGroupInput() *ConfigConfigurationRecorderRecordingGroup
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	PutRecordingGroup(value *ConfigConfigurationRecorderRecordingGroup)
	ResetName()
	ResetOverrideLogicalId()
	ResetRecordingGroup()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigConfigurationRecorder
type jsiiProxy_ConfigConfigurationRecorder struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigConfigurationRecorder) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) RecordingGroup() ConfigConfigurationRecorderRecordingGroupOutputReference {
	var returns ConfigConfigurationRecorderRecordingGroupOutputReference
	_jsii_.Get(
		j,
		"recordingGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) RecordingGroupInput() *ConfigConfigurationRecorderRecordingGroup {
	var returns *ConfigConfigurationRecorderRecordingGroup
	_jsii_.Get(
		j,
		"recordingGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorder) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html aws_config_configuration_recorder} Resource.
func NewConfigConfigurationRecorder(scope constructs.Construct, id *string, config *ConfigConfigurationRecorderConfig) ConfigConfigurationRecorder {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationRecorder{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorder",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html aws_config_configuration_recorder} Resource.
func NewConfigConfigurationRecorder_Override(c ConfigConfigurationRecorder, scope constructs.Construct, id *string, config *ConfigConfigurationRecorderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorder",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorder) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigConfigurationRecorder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigConfigurationRecorder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigConfigurationRecorder_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigConfigurationRecorder",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorder) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorder) PutRecordingGroup(value *ConfigConfigurationRecorderRecordingGroup) {
	_jsii_.InvokeVoid(
		c,
		"putRecordingGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorder) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorder) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorder) ResetRecordingGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetRecordingGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorder) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) ToString() *string {
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
func (c *jsiiProxy_ConfigConfigurationRecorder) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigConfigurationRecorderConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#role_arn ConfigConfigurationRecorder#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#name ConfigConfigurationRecorder#name}.
	Name *string `json:"name"`
	// recording_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#recording_group ConfigConfigurationRecorder#recording_group}
	RecordingGroup *ConfigConfigurationRecorderRecordingGroup `json:"recordingGroup"`
}

type ConfigConfigurationRecorderRecordingGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#all_supported ConfigConfigurationRecorder#all_supported}.
	AllSupported interface{} `json:"allSupported"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#include_global_resource_types ConfigConfigurationRecorder#include_global_resource_types}.
	IncludeGlobalResourceTypes interface{} `json:"includeGlobalResourceTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html#resource_types ConfigConfigurationRecorder#resource_types}.
	ResourceTypes *[]*string `json:"resourceTypes"`
}

type ConfigConfigurationRecorderRecordingGroupOutputReference interface {
	cdktf.ComplexObject
	AllSupported() interface{}
	SetAllSupported(val interface{})
	AllSupportedInput() interface{}
	IncludeGlobalResourceTypes() interface{}
	SetIncludeGlobalResourceTypes(val interface{})
	IncludeGlobalResourceTypesInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
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
	ResetAllSupported()
	ResetIncludeGlobalResourceTypes()
	ResetResourceTypes()
}

// The jsii proxy struct for ConfigConfigurationRecorderRecordingGroupOutputReference
type jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigConfigurationRecorderRecordingGroupOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigConfigurationRecorderRecordingGroupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigConfigurationRecorderRecordingGroupOutputReference_Override(c ConfigConfigurationRecorderRecordingGroupOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetAllSupported(val interface{}) {
	_jsii_.Set(
		j,
		"allSupported",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetIncludeGlobalResourceTypes(val interface{}) {
	_jsii_.Set(
		j,
		"includeGlobalResourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetResourceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetAllSupported() {
	_jsii_.InvokeVoid(
		c,
		"resetAllSupported",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetIncludeGlobalResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeGlobalResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypes",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html aws_config_configuration_recorder_status}.
type ConfigConfigurationRecorderStatus interface {
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
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
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

// The jsii proxy struct for ConfigConfigurationRecorderStatus
type jsiiProxy_ConfigConfigurationRecorderStatus struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html aws_config_configuration_recorder_status} Resource.
func NewConfigConfigurationRecorderStatus(scope constructs.Construct, id *string, config *ConfigConfigurationRecorderStatusConfig) ConfigConfigurationRecorderStatus {
	_init_.Initialize()

	j := jsiiProxy_ConfigConfigurationRecorderStatus{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorderStatus",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html aws_config_configuration_recorder_status} Resource.
func NewConfigConfigurationRecorderStatus_Override(c ConfigConfigurationRecorderStatus, scope constructs.Construct, id *string, config *ConfigConfigurationRecorderStatusConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConfigurationRecorderStatus",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderStatus) SetProvider(val cdktf.TerraformProvider) {
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
func ConfigConfigurationRecorderStatus_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigConfigurationRecorderStatus",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigConfigurationRecorderStatus_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigConfigurationRecorderStatus",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderStatus) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) ToString() *string {
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
func (c *jsiiProxy_ConfigConfigurationRecorderStatus) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigConfigurationRecorderStatusConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html#is_enabled ConfigConfigurationRecorderStatus#is_enabled}.
	IsEnabled interface{} `json:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html#name ConfigConfigurationRecorderStatus#name}.
	Name *string `json:"name"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html aws_config_conformance_pack}.
type ConfigConformancePack interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DeliveryS3Bucket() *string
	SetDeliveryS3Bucket(val *string)
	DeliveryS3BucketInput() *string
	DeliveryS3KeyPrefix() *string
	SetDeliveryS3KeyPrefix(val *string)
	DeliveryS3KeyPrefixInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InputParameter() *[]*ConfigConformancePackInputParameter
	SetInputParameter(val *[]*ConfigConformancePackInputParameter)
	InputParameterInput() *[]*ConfigConformancePackInputParameter
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TemplateBody() *string
	SetTemplateBody(val *string)
	TemplateBodyInput() *string
	TemplateS3Uri() *string
	SetTemplateS3Uri(val *string)
	TemplateS3UriInput() *string
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
	ResetDeliveryS3Bucket()
	ResetDeliveryS3KeyPrefix()
	ResetInputParameter()
	ResetOverrideLogicalId()
	ResetTemplateBody()
	ResetTemplateS3Uri()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigConformancePack
type jsiiProxy_ConfigConformancePack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigConformancePack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) DeliveryS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) DeliveryS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) DeliveryS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) DeliveryS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) InputParameter() *[]*ConfigConformancePackInputParameter {
	var returns *[]*ConfigConformancePackInputParameter
	_jsii_.Get(
		j,
		"inputParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) InputParameterInput() *[]*ConfigConformancePackInputParameter {
	var returns *[]*ConfigConformancePackInputParameter
	_jsii_.Get(
		j,
		"inputParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TemplateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TemplateS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TemplateS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConformancePack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html aws_config_conformance_pack} Resource.
func NewConfigConformancePack(scope constructs.Construct, id *string, config *ConfigConformancePackConfig) ConfigConformancePack {
	_init_.Initialize()

	j := jsiiProxy_ConfigConformancePack{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConformancePack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html aws_config_conformance_pack} Resource.
func NewConfigConformancePack_Override(c ConfigConformancePack, scope constructs.Construct, id *string, config *ConfigConformancePackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigConformancePack",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetDeliveryS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3Bucket",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetDeliveryS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetInputParameter(val *[]*ConfigConformancePackInputParameter) {
	_jsii_.Set(
		j,
		"inputParameter",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_ConfigConformancePack) SetTemplateS3Uri(val *string) {
	_jsii_.Set(
		j,
		"templateS3Uri",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigConformancePack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigConformancePack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigConformancePack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigConformancePack",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigConformancePack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigConformancePack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConformancePack) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigConformancePack) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigConformancePack) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigConformancePack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigConformancePack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigConformancePack) ResetDeliveryS3Bucket() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryS3Bucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConformancePack) ResetDeliveryS3KeyPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryS3KeyPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConformancePack) ResetInputParameter() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameter",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigConformancePack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConformancePack) ResetTemplateBody() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConformancePack) ResetTemplateS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConformancePack) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigConformancePack) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigConformancePack) ToString() *string {
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
func (c *jsiiProxy_ConfigConformancePack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigConformancePackConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#name ConfigConformancePack#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#delivery_s3_bucket ConfigConformancePack#delivery_s3_bucket}.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#delivery_s3_key_prefix ConfigConformancePack#delivery_s3_key_prefix}.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#input_parameter ConfigConformancePack#input_parameter}
	InputParameter *[]*ConfigConformancePackInputParameter `json:"inputParameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#template_body ConfigConformancePack#template_body}.
	TemplateBody *string `json:"templateBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#template_s3_uri ConfigConformancePack#template_s3_uri}.
	TemplateS3Uri *string `json:"templateS3Uri"`
}

type ConfigConformancePackInputParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#parameter_name ConfigConformancePack#parameter_name}.
	ParameterName *string `json:"parameterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack.html#parameter_value ConfigConformancePack#parameter_value}.
	ParameterValue *string `json:"parameterValue"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html aws_config_delivery_channel}.
type ConfigDeliveryChannel interface {
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
	S3KmsKeyArn() *string
	SetS3KmsKeyArn(val *string)
	S3KmsKeyArnInput() *string
	SnapshotDeliveryProperties() ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference
	SnapshotDeliveryPropertiesInput() *ConfigDeliveryChannelSnapshotDeliveryProperties
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	SnsTopicArnInput() *string
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
	PutSnapshotDeliveryProperties(value *ConfigDeliveryChannelSnapshotDeliveryProperties)
	ResetName()
	ResetOverrideLogicalId()
	ResetS3KeyPrefix()
	ResetS3KmsKeyArn()
	ResetSnapshotDeliveryProperties()
	ResetSnsTopicArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigDeliveryChannel
type jsiiProxy_ConfigDeliveryChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigDeliveryChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) S3KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) SnapshotDeliveryProperties() ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference {
	var returns ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference
	_jsii_.Get(
		j,
		"snapshotDeliveryProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) SnapshotDeliveryPropertiesInput() *ConfigDeliveryChannelSnapshotDeliveryProperties {
	var returns *ConfigDeliveryChannelSnapshotDeliveryProperties
	_jsii_.Get(
		j,
		"snapshotDeliveryPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) SnsTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html aws_config_delivery_channel} Resource.
func NewConfigDeliveryChannel(scope constructs.Construct, id *string, config *ConfigDeliveryChannelConfig) ConfigDeliveryChannel {
	_init_.Initialize()

	j := jsiiProxy_ConfigDeliveryChannel{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigDeliveryChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html aws_config_delivery_channel} Resource.
func NewConfigDeliveryChannel_Override(c ConfigDeliveryChannel, scope constructs.Construct, id *string, config *ConfigDeliveryChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigDeliveryChannel",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetS3KmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"s3KmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannel) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigDeliveryChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigDeliveryChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigDeliveryChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigDeliveryChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigDeliveryChannel) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigDeliveryChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigDeliveryChannel) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigDeliveryChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigDeliveryChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) PutSnapshotDeliveryProperties(value *ConfigDeliveryChannelSnapshotDeliveryProperties) {
	_jsii_.InvokeVoid(
		c,
		"putSnapshotDeliveryProperties",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) ResetS3KmsKeyArn() {
	_jsii_.InvokeVoid(
		c,
		"resetS3KmsKeyArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) ResetSnapshotDeliveryProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotDeliveryProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) ResetSnsTopicArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeliveryChannel) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigDeliveryChannel) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigDeliveryChannel) ToString() *string {
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
func (c *jsiiProxy_ConfigDeliveryChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigDeliveryChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#s3_bucket_name ConfigDeliveryChannel#s3_bucket_name}.
	S3BucketName *string `json:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#name ConfigDeliveryChannel#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#s3_key_prefix ConfigDeliveryChannel#s3_key_prefix}.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#s3_kms_key_arn ConfigDeliveryChannel#s3_kms_key_arn}.
	S3KmsKeyArn *string `json:"s3KmsKeyArn"`
	// snapshot_delivery_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#snapshot_delivery_properties ConfigDeliveryChannel#snapshot_delivery_properties}
	SnapshotDeliveryProperties *ConfigDeliveryChannelSnapshotDeliveryProperties `json:"snapshotDeliveryProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#sns_topic_arn ConfigDeliveryChannel#sns_topic_arn}.
	SnsTopicArn *string `json:"snsTopicArn"`
}

type ConfigDeliveryChannelSnapshotDeliveryProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html#delivery_frequency ConfigDeliveryChannel#delivery_frequency}.
	DeliveryFrequency *string `json:"deliveryFrequency"`
}

type ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference interface {
	cdktf.ComplexObject
	DeliveryFrequency() *string
	SetDeliveryFrequency(val *string)
	DeliveryFrequencyInput() *string
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
	ResetDeliveryFrequency()
}

// The jsii proxy struct for ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference
type jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) DeliveryFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) DeliveryFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference_Override(c ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) SetDeliveryFrequency(val *string) {
	_jsii_.Set(
		j,
		"deliveryFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigDeliveryChannelSnapshotDeliveryPropertiesOutputReference) ResetDeliveryFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryFrequency",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html aws_config_organization_conformance_pack}.
type ConfigOrganizationConformancePack interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DeliveryS3Bucket() *string
	SetDeliveryS3Bucket(val *string)
	DeliveryS3BucketInput() *string
	DeliveryS3KeyPrefix() *string
	SetDeliveryS3KeyPrefix(val *string)
	DeliveryS3KeyPrefixInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	ExcludedAccountsInput() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InputParameter() *[]*ConfigOrganizationConformancePackInputParameter
	SetInputParameter(val *[]*ConfigOrganizationConformancePackInputParameter)
	InputParameterInput() *[]*ConfigOrganizationConformancePackInputParameter
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TemplateBody() *string
	SetTemplateBody(val *string)
	TemplateBodyInput() *string
	TemplateS3Uri() *string
	SetTemplateS3Uri(val *string)
	TemplateS3UriInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ConfigOrganizationConformancePackTimeoutsOutputReference
	TimeoutsInput() *ConfigOrganizationConformancePackTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ConfigOrganizationConformancePackTimeouts)
	ResetDeliveryS3Bucket()
	ResetDeliveryS3KeyPrefix()
	ResetExcludedAccounts()
	ResetInputParameter()
	ResetOverrideLogicalId()
	ResetTemplateBody()
	ResetTemplateS3Uri()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigOrganizationConformancePack
type jsiiProxy_ConfigOrganizationConformancePack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) DeliveryS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) DeliveryS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) DeliveryS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) DeliveryS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) ExcludedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) InputParameter() *[]*ConfigOrganizationConformancePackInputParameter {
	var returns *[]*ConfigOrganizationConformancePackInputParameter
	_jsii_.Get(
		j,
		"inputParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) InputParameterInput() *[]*ConfigOrganizationConformancePackInputParameter {
	var returns *[]*ConfigOrganizationConformancePackInputParameter
	_jsii_.Get(
		j,
		"inputParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TemplateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TemplateS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TemplateS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) Timeouts() ConfigOrganizationConformancePackTimeoutsOutputReference {
	var returns ConfigOrganizationConformancePackTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) TimeoutsInput() *ConfigOrganizationConformancePackTimeouts {
	var returns *ConfigOrganizationConformancePackTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html aws_config_organization_conformance_pack} Resource.
func NewConfigOrganizationConformancePack(scope constructs.Construct, id *string, config *ConfigOrganizationConformancePackConfig) ConfigOrganizationConformancePack {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationConformancePack{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationConformancePack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html aws_config_organization_conformance_pack} Resource.
func NewConfigOrganizationConformancePack_Override(c ConfigOrganizationConformancePack, scope constructs.Construct, id *string, config *ConfigOrganizationConformancePackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationConformancePack",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetDeliveryS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3Bucket",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetDeliveryS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetInputParameter(val *[]*ConfigOrganizationConformancePackInputParameter) {
	_jsii_.Set(
		j,
		"inputParameter",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePack) SetTemplateS3Uri(val *string) {
	_jsii_.Set(
		j,
		"templateS3Uri",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigOrganizationConformancePack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigOrganizationConformancePack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigOrganizationConformancePack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigOrganizationConformancePack",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) PutTimeouts(value *ConfigOrganizationConformancePackTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetDeliveryS3Bucket() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryS3Bucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetDeliveryS3KeyPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryS3KeyPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetExcludedAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetInputParameter() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameter",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetTemplateBody() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetTemplateS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePack) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) ToString() *string {
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
func (c *jsiiProxy_ConfigOrganizationConformancePack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigOrganizationConformancePackConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#name ConfigOrganizationConformancePack#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#delivery_s3_bucket ConfigOrganizationConformancePack#delivery_s3_bucket}.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#delivery_s3_key_prefix ConfigOrganizationConformancePack#delivery_s3_key_prefix}.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#excluded_accounts ConfigOrganizationConformancePack#excluded_accounts}.
	ExcludedAccounts *[]*string `json:"excludedAccounts"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#input_parameter ConfigOrganizationConformancePack#input_parameter}
	InputParameter *[]*ConfigOrganizationConformancePackInputParameter `json:"inputParameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#template_body ConfigOrganizationConformancePack#template_body}.
	TemplateBody *string `json:"templateBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#template_s3_uri ConfigOrganizationConformancePack#template_s3_uri}.
	TemplateS3Uri *string `json:"templateS3Uri"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#timeouts ConfigOrganizationConformancePack#timeouts}
	Timeouts *ConfigOrganizationConformancePackTimeouts `json:"timeouts"`
}

type ConfigOrganizationConformancePackInputParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#parameter_name ConfigOrganizationConformancePack#parameter_name}.
	ParameterName *string `json:"parameterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#parameter_value ConfigOrganizationConformancePack#parameter_value}.
	ParameterValue *string `json:"parameterValue"`
}

type ConfigOrganizationConformancePackTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#create ConfigOrganizationConformancePack#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#delete ConfigOrganizationConformancePack#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_conformance_pack.html#update ConfigOrganizationConformancePack#update}.
	Update *string `json:"update"`
}

type ConfigOrganizationConformancePackTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
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
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for ConfigOrganizationConformancePackTimeoutsOutputReference
type jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewConfigOrganizationConformancePackTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigOrganizationConformancePackTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationConformancePackTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigOrganizationConformancePackTimeoutsOutputReference_Override(c ConfigOrganizationConformancePackTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationConformancePackTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationConformancePackTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html aws_config_organization_custom_rule}.
type ConfigOrganizationCustomRule interface {
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
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	ExcludedAccountsInput() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InputParameters() *string
	SetInputParameters(val *string)
	InputParametersInput() *string
	LambdaFunctionArn() *string
	SetLambdaFunctionArn(val *string)
	LambdaFunctionArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	MaximumExecutionFrequencyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceIdScope() *string
	SetResourceIdScope(val *string)
	ResourceIdScopeInput() *string
	ResourceTypesScope() *[]*string
	SetResourceTypesScope(val *[]*string)
	ResourceTypesScopeInput() *[]*string
	TagKeyScope() *string
	SetTagKeyScope(val *string)
	TagKeyScopeInput() *string
	TagValueScope() *string
	SetTagValueScope(val *string)
	TagValueScopeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ConfigOrganizationCustomRuleTimeoutsOutputReference
	TimeoutsInput() *ConfigOrganizationCustomRuleTimeouts
	TriggerTypes() *[]*string
	SetTriggerTypes(val *[]*string)
	TriggerTypesInput() *[]*string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ConfigOrganizationCustomRuleTimeouts)
	ResetDescription()
	ResetExcludedAccounts()
	ResetInputParameters()
	ResetMaximumExecutionFrequency()
	ResetOverrideLogicalId()
	ResetResourceIdScope()
	ResetResourceTypesScope()
	ResetTagKeyScope()
	ResetTagValueScope()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigOrganizationCustomRule
type jsiiProxy_ConfigOrganizationCustomRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ExcludedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) LambdaFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) LambdaFunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) Timeouts() ConfigOrganizationCustomRuleTimeoutsOutputReference {
	var returns ConfigOrganizationCustomRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TimeoutsInput() *ConfigOrganizationCustomRuleTimeouts {
	var returns *ConfigOrganizationCustomRuleTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TriggerTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) TriggerTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"triggerTypesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html aws_config_organization_custom_rule} Resource.
func NewConfigOrganizationCustomRule(scope constructs.Construct, id *string, config *ConfigOrganizationCustomRuleConfig) ConfigOrganizationCustomRule {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationCustomRule{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationCustomRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html aws_config_organization_custom_rule} Resource.
func NewConfigOrganizationCustomRule_Override(c ConfigOrganizationCustomRule, scope constructs.Construct, id *string, config *ConfigOrganizationCustomRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationCustomRule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetInputParameters(val *string) {
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetLambdaFunctionArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaFunctionArn",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetMaximumExecutionFrequency(val *string) {
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetResourceIdScope(val *string) {
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetResourceTypesScope(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetTagKeyScope(val *string) {
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetTagValueScope(val *string) {
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRule) SetTriggerTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"triggerTypes",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigOrganizationCustomRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigOrganizationCustomRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigOrganizationCustomRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigOrganizationCustomRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) PutTimeouts(value *ConfigOrganizationCustomRuleTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetExcludedAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRule) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) ToString() *string {
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
func (c *jsiiProxy_ConfigOrganizationCustomRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigOrganizationCustomRuleConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#lambda_function_arn ConfigOrganizationCustomRule#lambda_function_arn}.
	LambdaFunctionArn *string `json:"lambdaFunctionArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#name ConfigOrganizationCustomRule#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#trigger_types ConfigOrganizationCustomRule#trigger_types}.
	TriggerTypes *[]*string `json:"triggerTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#description ConfigOrganizationCustomRule#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#excluded_accounts ConfigOrganizationCustomRule#excluded_accounts}.
	ExcludedAccounts *[]*string `json:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#input_parameters ConfigOrganizationCustomRule#input_parameters}.
	InputParameters *string `json:"inputParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#maximum_execution_frequency ConfigOrganizationCustomRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#resource_id_scope ConfigOrganizationCustomRule#resource_id_scope}.
	ResourceIdScope *string `json:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#resource_types_scope ConfigOrganizationCustomRule#resource_types_scope}.
	ResourceTypesScope *[]*string `json:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#tag_key_scope ConfigOrganizationCustomRule#tag_key_scope}.
	TagKeyScope *string `json:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#tag_value_scope ConfigOrganizationCustomRule#tag_value_scope}.
	TagValueScope *string `json:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#timeouts ConfigOrganizationCustomRule#timeouts}
	Timeouts *ConfigOrganizationCustomRuleTimeouts `json:"timeouts"`
}

type ConfigOrganizationCustomRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#create ConfigOrganizationCustomRule#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#delete ConfigOrganizationCustomRule#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_custom_rule.html#update ConfigOrganizationCustomRule#update}.
	Update *string `json:"update"`
}

type ConfigOrganizationCustomRuleTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
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
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for ConfigOrganizationCustomRuleTimeoutsOutputReference
type jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewConfigOrganizationCustomRuleTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigOrganizationCustomRuleTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationCustomRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigOrganizationCustomRuleTimeoutsOutputReference_Override(c ConfigOrganizationCustomRuleTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationCustomRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationCustomRuleTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html aws_config_organization_managed_rule}.
type ConfigOrganizationManagedRule interface {
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
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	ExcludedAccountsInput() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InputParameters() *string
	SetInputParameters(val *string)
	InputParametersInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	MaximumExecutionFrequencyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceIdScope() *string
	SetResourceIdScope(val *string)
	ResourceIdScopeInput() *string
	ResourceTypesScope() *[]*string
	SetResourceTypesScope(val *[]*string)
	ResourceTypesScopeInput() *[]*string
	RuleIdentifier() *string
	SetRuleIdentifier(val *string)
	RuleIdentifierInput() *string
	TagKeyScope() *string
	SetTagKeyScope(val *string)
	TagKeyScopeInput() *string
	TagValueScope() *string
	SetTagValueScope(val *string)
	TagValueScopeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ConfigOrganizationManagedRuleTimeoutsOutputReference
	TimeoutsInput() *ConfigOrganizationManagedRuleTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ConfigOrganizationManagedRuleTimeouts)
	ResetDescription()
	ResetExcludedAccounts()
	ResetInputParameters()
	ResetMaximumExecutionFrequency()
	ResetOverrideLogicalId()
	ResetResourceIdScope()
	ResetResourceTypesScope()
	ResetTagKeyScope()
	ResetTagValueScope()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigOrganizationManagedRule
type jsiiProxy_ConfigOrganizationManagedRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ExcludedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) RuleIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) RuleIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) Timeouts() ConfigOrganizationManagedRuleTimeoutsOutputReference {
	var returns ConfigOrganizationManagedRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) TimeoutsInput() *ConfigOrganizationManagedRuleTimeouts {
	var returns *ConfigOrganizationManagedRuleTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html aws_config_organization_managed_rule} Resource.
func NewConfigOrganizationManagedRule(scope constructs.Construct, id *string, config *ConfigOrganizationManagedRuleConfig) ConfigOrganizationManagedRule {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationManagedRule{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationManagedRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html aws_config_organization_managed_rule} Resource.
func NewConfigOrganizationManagedRule_Override(c ConfigOrganizationManagedRule, scope constructs.Construct, id *string, config *ConfigOrganizationManagedRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationManagedRule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetInputParameters(val *string) {
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetMaximumExecutionFrequency(val *string) {
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetResourceIdScope(val *string) {
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetResourceTypesScope(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetRuleIdentifier(val *string) {
	_jsii_.Set(
		j,
		"ruleIdentifier",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetTagKeyScope(val *string) {
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRule) SetTagValueScope(val *string) {
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigOrganizationManagedRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigOrganizationManagedRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigOrganizationManagedRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigOrganizationManagedRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) PutTimeouts(value *ConfigOrganizationManagedRuleTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetExcludedAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetInputParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRule) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) ToString() *string {
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
func (c *jsiiProxy_ConfigOrganizationManagedRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigOrganizationManagedRuleConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#name ConfigOrganizationManagedRule#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#rule_identifier ConfigOrganizationManagedRule#rule_identifier}.
	RuleIdentifier *string `json:"ruleIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#description ConfigOrganizationManagedRule#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#excluded_accounts ConfigOrganizationManagedRule#excluded_accounts}.
	ExcludedAccounts *[]*string `json:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#input_parameters ConfigOrganizationManagedRule#input_parameters}.
	InputParameters *string `json:"inputParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#maximum_execution_frequency ConfigOrganizationManagedRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#resource_id_scope ConfigOrganizationManagedRule#resource_id_scope}.
	ResourceIdScope *string `json:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#resource_types_scope ConfigOrganizationManagedRule#resource_types_scope}.
	ResourceTypesScope *[]*string `json:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#tag_key_scope ConfigOrganizationManagedRule#tag_key_scope}.
	TagKeyScope *string `json:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#tag_value_scope ConfigOrganizationManagedRule#tag_value_scope}.
	TagValueScope *string `json:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#timeouts ConfigOrganizationManagedRule#timeouts}
	Timeouts *ConfigOrganizationManagedRuleTimeouts `json:"timeouts"`
}

type ConfigOrganizationManagedRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#create ConfigOrganizationManagedRule#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#delete ConfigOrganizationManagedRule#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html#update ConfigOrganizationManagedRule#update}.
	Update *string `json:"update"`
}

type ConfigOrganizationManagedRuleTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
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
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for ConfigOrganizationManagedRuleTimeoutsOutputReference
type jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewConfigOrganizationManagedRuleTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ConfigOrganizationManagedRuleTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationManagedRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewConfigOrganizationManagedRuleTimeoutsOutputReference_Override(c ConfigOrganizationManagedRuleTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigOrganizationManagedRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigOrganizationManagedRuleTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html aws_config_remediation_configuration}.
type ConfigRemediationConfiguration interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfigRuleName() *string
	SetConfigRuleName(val *string)
	ConfigRuleNameInput() *string
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
	Parameter() *[]*ConfigRemediationConfigurationParameter
	SetParameter(val *[]*ConfigRemediationConfigurationParameter)
	ParameterInput() *[]*ConfigRemediationConfigurationParameter
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	TargetId() *string
	SetTargetId(val *string)
	TargetIdInput() *string
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	TargetVersion() *string
	SetTargetVersion(val *string)
	TargetVersionInput() *string
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
	ResetParameter()
	ResetResourceType()
	ResetTargetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConfigRemediationConfiguration
type jsiiProxy_ConfigRemediationConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConfigRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Parameter() *[]*ConfigRemediationConfigurationParameter {
	var returns *[]*ConfigRemediationConfigurationParameter
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ParameterInput() *[]*ConfigRemediationConfigurationParameter {
	var returns *[]*ConfigRemediationConfigurationParameter
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TargetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html aws_config_remediation_configuration} Resource.
func NewConfigRemediationConfiguration(scope constructs.Construct, id *string, config *ConfigRemediationConfigurationConfig) ConfigRemediationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_ConfigRemediationConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigRemediationConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html aws_config_remediation_configuration} Resource.
func NewConfigRemediationConfiguration_Override(c ConfigRemediationConfiguration, scope constructs.Construct, id *string, config *ConfigRemediationConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Config.ConfigRemediationConfiguration",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetConfigRuleName(val *string) {
	_jsii_.Set(
		j,
		"configRuleName",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetParameter(val *[]*ConfigRemediationConfigurationParameter) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetTargetId(val *string) {
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfiguration) SetTargetVersion(val *string) {
	_jsii_.Set(
		j,
		"targetVersion",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ConfigRemediationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Config.ConfigRemediationConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigRemediationConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Config.ConfigRemediationConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConfigRemediationConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ConfigRemediationConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConfigRemediationConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetParameter() {
	_jsii_.InvokeVoid(
		c,
		"resetParameter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetResourceType() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) ResetTargetVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfiguration) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) ToMetadata() interface{} {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) ToString() *string {
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
func (c *jsiiProxy_ConfigRemediationConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ConfigRemediationConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#config_rule_name ConfigRemediationConfiguration#config_rule_name}.
	ConfigRuleName *string `json:"configRuleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#target_id ConfigRemediationConfiguration#target_id}.
	TargetId *string `json:"targetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#target_type ConfigRemediationConfiguration#target_type}.
	TargetType *string `json:"targetType"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#parameter ConfigRemediationConfiguration#parameter}
	Parameter *[]*ConfigRemediationConfigurationParameter `json:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#resource_type ConfigRemediationConfiguration#resource_type}.
	ResourceType *string `json:"resourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#target_version ConfigRemediationConfiguration#target_version}.
	TargetVersion *string `json:"targetVersion"`
}

type ConfigRemediationConfigurationParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#name ConfigRemediationConfiguration#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#resource_value ConfigRemediationConfiguration#resource_value}.
	ResourceValue *string `json:"resourceValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration.html#static_value ConfigRemediationConfiguration#static_value}.
	StaticValue *string `json:"staticValue"`
}
