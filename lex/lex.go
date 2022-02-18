package lex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/lex/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lex_bot aws_lex_bot}.
type DataAwsLexBot interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ChildDirected() cdktf.IResolvable
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	DetectSentiment() cdktf.IResolvable
	EnableModelImprovements() cdktf.IResolvable
	FailureReason() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdleSessionTtlInSeconds() *float64
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NluIntentConfidenceThreshold() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VoiceId() *string
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
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLexBot
type jsiiProxy_DataAwsLexBot struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLexBot) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) ChildDirected() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"childDirected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) DetectSentiment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detectSentiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) EnableModelImprovements() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableModelImprovements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) NluIntentConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBot) VoiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceId",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_bot aws_lex_bot} Data Source.
func NewDataAwsLexBot(scope constructs.Construct, id *string, config *DataAwsLexBotConfig) DataAwsLexBot {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLexBot{}

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_bot aws_lex_bot} Data Source.
func NewDataAwsLexBot_Override(d DataAwsLexBot, scope constructs.Construct, id *string, config *DataAwsLexBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexBot",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBot) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLexBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.DataAwsLexBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLexBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.DataAwsLexBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLexBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLexBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexBot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLexBot) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLexBot) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLexBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLexBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLexBot) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLexBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLexBot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexBot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLexBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexBot) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexBot) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexBot) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLexBot) ToString() *string {
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
func (d *jsiiProxy_DataAwsLexBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lex_bot_alias aws_lex_bot_alias}.
type DataAwsLexBotAlias interface {
	cdktf.TerraformDataSource
	Arn() *string
	BotName() *string
	SetBotName(val *string)
	BotNameInput() *string
	BotVersion() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
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

// The jsii proxy struct for DataAwsLexBotAlias
type jsiiProxy_DataAwsLexBotAlias struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLexBotAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) BotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) BotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexBotAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_bot_alias aws_lex_bot_alias} Data Source.
func NewDataAwsLexBotAlias(scope constructs.Construct, id *string, config *DataAwsLexBotAliasConfig) DataAwsLexBotAlias {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLexBotAlias{}

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexBotAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_bot_alias aws_lex_bot_alias} Data Source.
func NewDataAwsLexBotAlias_Override(d DataAwsLexBotAlias, scope constructs.Construct, id *string, config *DataAwsLexBotAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexBotAlias",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetBotName(val *string) {
	_jsii_.Set(
		j,
		"botName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexBotAlias) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsLexBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.DataAwsLexBotAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLexBotAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.DataAwsLexBotAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLexBotAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLexBotAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLexBotAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLexBotAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexBotAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLexBotAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexBotAlias) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexBotAlias) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLexBotAlias) ToString() *string {
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
func (d *jsiiProxy_DataAwsLexBotAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lex.
type DataAwsLexBotAliasConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_bot_alias#bot_name DataAwsLexBotAlias#bot_name}.
	BotName *string `json:"botName" yaml:"botName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_bot_alias#name DataAwsLexBotAlias#name}.
	Name *string `json:"name" yaml:"name"`
}

// AWS Lex.
type DataAwsLexBotConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_bot#name DataAwsLexBot#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_bot#version DataAwsLexBot#version}.
	Version *string `json:"version" yaml:"version"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lex_intent aws_lex_intent}.
type DataAwsLexIntent interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ParentIntentSignature() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLexIntent
type jsiiProxy_DataAwsLexIntent struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLexIntent) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexIntent) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_intent aws_lex_intent} Data Source.
func NewDataAwsLexIntent(scope constructs.Construct, id *string, config *DataAwsLexIntentConfig) DataAwsLexIntent {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLexIntent{}

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_intent aws_lex_intent} Data Source.
func NewDataAwsLexIntent_Override(d DataAwsLexIntent, scope constructs.Construct, id *string, config *DataAwsLexIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexIntent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexIntent) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLexIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.DataAwsLexIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLexIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.DataAwsLexIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLexIntent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLexIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexIntent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLexIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLexIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLexIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLexIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLexIntent) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLexIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLexIntent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexIntent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLexIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexIntent) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexIntent) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexIntent) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLexIntent) ToString() *string {
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
func (d *jsiiProxy_DataAwsLexIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lex.
type DataAwsLexIntentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_intent#name DataAwsLexIntent#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_intent#version DataAwsLexIntent#version}.
	Version *string `json:"version" yaml:"version"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lex_slot_type aws_lex_slot_type}.
type DataAwsLexSlotType interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
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
	ValueSelectionStrategy() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	AddOverride(path *string, value interface{})
	EnumerationValue(index *string) DataAwsLexSlotTypeEnumerationValue
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
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLexSlotType
type jsiiProxy_DataAwsLexSlotType struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLexSlotType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) ValueSelectionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueSelectionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotType) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_slot_type aws_lex_slot_type} Data Source.
func NewDataAwsLexSlotType(scope constructs.Construct, id *string, config *DataAwsLexSlotTypeConfig) DataAwsLexSlotType {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLexSlotType{}

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexSlotType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lex_slot_type aws_lex_slot_type} Data Source.
func NewDataAwsLexSlotType_Override(d DataAwsLexSlotType, scope constructs.Construct, id *string, config *DataAwsLexSlotTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexSlotType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotType) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsLexSlotType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.DataAwsLexSlotType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLexSlotType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.DataAwsLexSlotType",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLexSlotType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLexSlotType) EnumerationValue(index *string) DataAwsLexSlotTypeEnumerationValue {
	var returns DataAwsLexSlotTypeEnumerationValue

	_jsii_.Invoke(
		d,
		"enumerationValue",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLexSlotType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLexSlotType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLexSlotType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexSlotType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLexSlotType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexSlotType) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLexSlotType) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexSlotType) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLexSlotType) ToString() *string {
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
func (d *jsiiProxy_DataAwsLexSlotType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lex.
type DataAwsLexSlotTypeConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_slot_type#name DataAwsLexSlotType#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lex_slot_type#version DataAwsLexSlotType#version}.
	Version *string `json:"version" yaml:"version"`
}

type DataAwsLexSlotTypeEnumerationValue interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Synonyms() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
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

// The jsii proxy struct for DataAwsLexSlotTypeEnumerationValue
type jsiiProxy_DataAwsLexSlotTypeEnumerationValue struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) Synonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"synonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLexSlotTypeEnumerationValue(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLexSlotTypeEnumerationValue {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLexSlotTypeEnumerationValue{}

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexSlotTypeEnumerationValue",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLexSlotTypeEnumerationValue_Override(d DataAwsLexSlotTypeEnumerationValue, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.DataAwsLexSlotTypeEnumerationValue",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValue) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lex_bot aws_lex_bot}.
type LexBot interface {
	cdktf.TerraformResource
	AbortStatement() LexBotAbortStatementOutputReference
	AbortStatementInput() *LexBotAbortStatement
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ChildDirected() interface{}
	SetChildDirected(val interface{})
	ChildDirectedInput() interface{}
	ClarificationPrompt() LexBotClarificationPromptOutputReference
	ClarificationPromptInput() *LexBotClarificationPrompt
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	CreateVersion() interface{}
	SetCreateVersion(val interface{})
	CreateVersionInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetectSentiment() interface{}
	SetDetectSentiment(val interface{})
	DetectSentimentInput() interface{}
	EnableModelImprovements() interface{}
	SetEnableModelImprovements(val interface{})
	EnableModelImprovementsInput() interface{}
	FailureReason() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	IdleSessionTtlInSecondsInput() *float64
	Intent() interface{}
	SetIntent(val interface{})
	IntentInput() interface{}
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NluIntentConfidenceThreshold() *float64
	SetNluIntentConfidenceThreshold(val *float64)
	NluIntentConfidenceThresholdInput() *float64
	Node() constructs.Node
	ProcessBehavior() *string
	SetProcessBehavior(val *string)
	ProcessBehaviorInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() LexBotTimeoutsOutputReference
	TimeoutsInput() *LexBotTimeouts
	Version() *string
	VoiceId() *string
	SetVoiceId(val *string)
	VoiceIdInput() *string
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
	PutAbortStatement(value *LexBotAbortStatement)
	PutClarificationPrompt(value *LexBotClarificationPrompt)
	PutTimeouts(value *LexBotTimeouts)
	ResetClarificationPrompt()
	ResetCreateVersion()
	ResetDescription()
	ResetDetectSentiment()
	ResetEnableModelImprovements()
	ResetIdleSessionTtlInSeconds()
	ResetLocale()
	ResetNluIntentConfidenceThreshold()
	ResetOverrideLogicalId()
	ResetProcessBehavior()
	ResetTimeouts()
	ResetVoiceId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LexBot
type jsiiProxy_LexBot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexBot) AbortStatement() LexBotAbortStatementOutputReference {
	var returns LexBotAbortStatementOutputReference
	_jsii_.Get(
		j,
		"abortStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) AbortStatementInput() *LexBotAbortStatement {
	var returns *LexBotAbortStatement
	_jsii_.Get(
		j,
		"abortStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ChildDirected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ChildDirectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ClarificationPrompt() LexBotClarificationPromptOutputReference {
	var returns LexBotClarificationPromptOutputReference
	_jsii_.Get(
		j,
		"clarificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ClarificationPromptInput() *LexBotClarificationPrompt {
	var returns *LexBotClarificationPrompt
	_jsii_.Get(
		j,
		"clarificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DetectSentiment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DetectSentimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) EnableModelImprovements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) EnableModelImprovementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Intent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IntentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NluIntentConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NluIntentConfidenceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ProcessBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ProcessBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Timeouts() LexBotTimeoutsOutputReference {
	var returns LexBotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TimeoutsInput() *LexBotTimeouts {
	var returns *LexBotTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) VoiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) VoiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_bot aws_lex_bot} Resource.
func NewLexBot(scope constructs.Construct, id *string, config *LexBotConfig) LexBot {
	_init_.Initialize()

	j := jsiiProxy_LexBot{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_bot aws_lex_bot} Resource.
func NewLexBot_Override(l LexBot, scope constructs.Construct, id *string, config *LexBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBot",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexBot) SetChildDirected(val interface{}) {
	_jsii_.Set(
		j,
		"childDirected",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetCreateVersion(val interface{}) {
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetDetectSentiment(val interface{}) {
	_jsii_.Set(
		j,
		"detectSentiment",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetEnableModelImprovements(val interface{}) {
	_jsii_.Set(
		j,
		"enableModelImprovements",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetIdleSessionTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetIntent(val interface{}) {
	_jsii_.Set(
		j,
		"intent",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetLocale(val *string) {
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetNluIntentConfidenceThreshold(val *float64) {
	_jsii_.Set(
		j,
		"nluIntentConfidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetProcessBehavior(val *string) {
	_jsii_.Set(
		j,
		"processBehavior",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LexBot) SetVoiceId(val *string) {
	_jsii_.Set(
		j,
		"voiceId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LexBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.LexBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.LexBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LexBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LexBot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexBot) PutAbortStatement(value *LexBotAbortStatement) {
	_jsii_.InvokeVoid(
		l,
		"putAbortStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) PutClarificationPrompt(value *LexBotClarificationPrompt) {
	_jsii_.InvokeVoid(
		l,
		"putClarificationPrompt",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) PutTimeouts(value *LexBotTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) ResetClarificationPrompt() {
	_jsii_.InvokeVoid(
		l,
		"resetClarificationPrompt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetDetectSentiment() {
	_jsii_.InvokeVoid(
		l,
		"resetDetectSentiment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetEnableModelImprovements() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableModelImprovements",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetLocale() {
	_jsii_.InvokeVoid(
		l,
		"resetLocale",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetNluIntentConfidenceThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetNluIntentConfidenceThreshold",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LexBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetProcessBehavior() {
	_jsii_.InvokeVoid(
		l,
		"resetProcessBehavior",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetVoiceId() {
	_jsii_.InvokeVoid(
		l,
		"resetVoiceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LexBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LexBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LexBotAbortStatement struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#message LexBot#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#response_card LexBot#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexBotAbortStatementMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#content LexBot#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#content_type LexBot#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#group_number LexBot#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexBotAbortStatementOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexBotAbortStatement
	SetInternalValue(val *LexBotAbortStatement)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexBotAbortStatementOutputReference
type jsiiProxy_LexBotAbortStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) InternalValue() *LexBotAbortStatement {
	var returns *LexBotAbortStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexBotAbortStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexBotAbortStatementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexBotAbortStatementOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAbortStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexBotAbortStatementOutputReference_Override(l LexBotAbortStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAbortStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetInternalValue(val *LexBotAbortStatement) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotAbortStatementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAbortStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAbortStatementOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias aws_lex_bot_alias}.
type LexBotAlias interface {
	cdktf.TerraformResource
	Arn() *string
	BotName() *string
	SetBotName(val *string)
	BotNameInput() *string
	BotVersion() *string
	SetBotVersion(val *string)
	BotVersionInput() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConstructNodeMetadata() *map[string]interface{}
	ConversationLogs() LexBotAliasConversationLogsOutputReference
	ConversationLogsInput() *LexBotAliasConversationLogs
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
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
	Timeouts() LexBotAliasTimeoutsOutputReference
	TimeoutsInput() *LexBotAliasTimeouts
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
	PutConversationLogs(value *LexBotAliasConversationLogs)
	PutTimeouts(value *LexBotAliasTimeouts)
	ResetConversationLogs()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LexBotAlias
type jsiiProxy_LexBotAlias struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexBotAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConversationLogs() LexBotAliasConversationLogsOutputReference {
	var returns LexBotAliasConversationLogsOutputReference
	_jsii_.Get(
		j,
		"conversationLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) ConversationLogsInput() *LexBotAliasConversationLogs {
	var returns *LexBotAliasConversationLogs
	_jsii_.Get(
		j,
		"conversationLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) Timeouts() LexBotAliasTimeoutsOutputReference {
	var returns LexBotAliasTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAlias) TimeoutsInput() *LexBotAliasTimeouts {
	var returns *LexBotAliasTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias aws_lex_bot_alias} Resource.
func NewLexBotAlias(scope constructs.Construct, id *string, config *LexBotAliasConfig) LexBotAlias {
	_init_.Initialize()

	j := jsiiProxy_LexBotAlias{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias aws_lex_bot_alias} Resource.
func NewLexBotAlias_Override(l LexBotAlias, scope constructs.Construct, id *string, config *LexBotAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAlias",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexBotAlias) SetBotName(val *string) {
	_jsii_.Set(
		j,
		"botName",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetBotVersion(val *string) {
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexBotAlias) SetProvider(val cdktf.TerraformProvider) {
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
func LexBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.LexBotAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexBotAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.LexBotAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LexBotAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexBotAlias) PutConversationLogs(value *LexBotAliasConversationLogs) {
	_jsii_.InvokeVoid(
		l,
		"putConversationLogs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) PutTimeouts(value *LexBotAliasTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBotAlias) ResetConversationLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetConversationLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LexBotAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LexBotAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LexBotAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lex.
type LexBotAliasConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#bot_name LexBotAlias#bot_name}.
	BotName *string `json:"botName" yaml:"botName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#bot_version LexBotAlias#bot_version}.
	BotVersion *string `json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#name LexBotAlias#name}.
	Name *string `json:"name" yaml:"name"`
	// conversation_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#conversation_logs LexBotAlias#conversation_logs}
	ConversationLogs *LexBotAliasConversationLogs `json:"conversationLogs" yaml:"conversationLogs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#description LexBotAlias#description}.
	Description *string `json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#timeouts LexBotAlias#timeouts}
	Timeouts *LexBotAliasTimeouts `json:"timeouts" yaml:"timeouts"`
}

type LexBotAliasConversationLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#iam_role_arn LexBotAlias#iam_role_arn}.
	IamRoleArn *string `json:"iamRoleArn" yaml:"iamRoleArn"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#log_settings LexBotAlias#log_settings}
	LogSettings interface{} `json:"logSettings" yaml:"logSettings"`
}

type LexBotAliasConversationLogsLogSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#destination LexBotAlias#destination}.
	Destination *string `json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#log_type LexBotAlias#log_type}.
	LogType *string `json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#resource_arn LexBotAlias#resource_arn}.
	ResourceArn *string `json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#kms_key_arn LexBotAlias#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

type LexBotAliasConversationLogsOutputReference interface {
	cdktf.ComplexObject
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	InternalValue() *LexBotAliasConversationLogs
	SetInternalValue(val *LexBotAliasConversationLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LogSettings() interface{}
	SetLogSettings(val interface{})
	LogSettingsInput() interface{}
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
	ResetLogSettings()
}

// The jsii proxy struct for LexBotAliasConversationLogsOutputReference
type jsiiProxy_LexBotAliasConversationLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) InternalValue() *LexBotAliasConversationLogs {
	var returns *LexBotAliasConversationLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) LogSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) LogSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexBotAliasConversationLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexBotAliasConversationLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexBotAliasConversationLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAliasConversationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexBotAliasConversationLogsOutputReference_Override(l LexBotAliasConversationLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAliasConversationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetInternalValue(val *LexBotAliasConversationLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetLogSettings(val interface{}) {
	_jsii_.Set(
		j,
		"logSettings",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasConversationLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAliasConversationLogsOutputReference) ResetLogSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetLogSettings",
		nil, // no parameters
	)
}

type LexBotAliasTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#create LexBotAlias#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#delete LexBotAlias#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#update LexBotAlias#update}.
	Update *string `json:"update" yaml:"update"`
}

type LexBotAliasTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *LexBotAliasTimeouts
	SetInternalValue(val *LexBotAliasTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for LexBotAliasTimeoutsOutputReference
type jsiiProxy_LexBotAliasTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) InternalValue() *LexBotAliasTimeouts {
	var returns *LexBotAliasTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewLexBotAliasTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexBotAliasTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexBotAliasTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAliasTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexBotAliasTimeoutsOutputReference_Override(l LexBotAliasTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotAliasTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetInternalValue(val *LexBotAliasTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexBotAliasTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		l,
		"resetDelete",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotAliasTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdate",
		nil, // no parameters
	)
}

type LexBotClarificationPrompt struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#max_attempts LexBot#max_attempts}.
	MaxAttempts *float64 `json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#message LexBot#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#response_card LexBot#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexBotClarificationPromptMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#content LexBot#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#content_type LexBot#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#group_number LexBot#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexBotClarificationPromptOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexBotClarificationPrompt
	SetInternalValue(val *LexBotClarificationPrompt)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexBotClarificationPromptOutputReference
type jsiiProxy_LexBotClarificationPromptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) InternalValue() *LexBotClarificationPrompt {
	var returns *LexBotClarificationPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexBotClarificationPromptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexBotClarificationPromptOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexBotClarificationPromptOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotClarificationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexBotClarificationPromptOutputReference_Override(l LexBotClarificationPromptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotClarificationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetInternalValue(val *LexBotClarificationPrompt) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetMaxAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotClarificationPromptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotClarificationPromptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotClarificationPromptOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

// AWS Lex.
type LexBotConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// abort_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#abort_statement LexBot#abort_statement}
	AbortStatement *LexBotAbortStatement `json:"abortStatement" yaml:"abortStatement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#child_directed LexBot#child_directed}.
	ChildDirected interface{} `json:"childDirected" yaml:"childDirected"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#intent LexBot#intent}
	Intent interface{} `json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#name LexBot#name}.
	Name *string `json:"name" yaml:"name"`
	// clarification_prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#clarification_prompt LexBot#clarification_prompt}
	ClarificationPrompt *LexBotClarificationPrompt `json:"clarificationPrompt" yaml:"clarificationPrompt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#create_version LexBot#create_version}.
	CreateVersion interface{} `json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#description LexBot#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#detect_sentiment LexBot#detect_sentiment}.
	DetectSentiment interface{} `json:"detectSentiment" yaml:"detectSentiment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#enable_model_improvements LexBot#enable_model_improvements}.
	EnableModelImprovements interface{} `json:"enableModelImprovements" yaml:"enableModelImprovements"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#idle_session_ttl_in_seconds LexBot#idle_session_ttl_in_seconds}.
	IdleSessionTtlInSeconds *float64 `json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#locale LexBot#locale}.
	Locale *string `json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#nlu_intent_confidence_threshold LexBot#nlu_intent_confidence_threshold}.
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold" yaml:"nluIntentConfidenceThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#process_behavior LexBot#process_behavior}.
	ProcessBehavior *string `json:"processBehavior" yaml:"processBehavior"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#timeouts LexBot#timeouts}
	Timeouts *LexBotTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#voice_id LexBot#voice_id}.
	VoiceId *string `json:"voiceId" yaml:"voiceId"`
}

type LexBotIntent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#intent_name LexBot#intent_name}.
	IntentName *string `json:"intentName" yaml:"intentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#intent_version LexBot#intent_version}.
	IntentVersion *string `json:"intentVersion" yaml:"intentVersion"`
}

type LexBotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#create LexBot#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#delete LexBot#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot#update LexBot#update}.
	Update *string `json:"update" yaml:"update"`
}

type LexBotTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *LexBotTimeouts
	SetInternalValue(val *LexBotTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for LexBotTimeoutsOutputReference
type jsiiProxy_LexBotTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) InternalValue() *LexBotTimeouts {
	var returns *LexBotTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewLexBotTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexBotTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexBotTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexBotTimeoutsOutputReference_Override(l LexBotTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexBotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetInternalValue(val *LexBotTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexBotTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexBotTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBotTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		l,
		"resetDelete",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBotTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lex_intent aws_lex_intent}.
type LexIntent interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConclusionStatement() LexIntentConclusionStatementOutputReference
	ConclusionStatementInput() *LexIntentConclusionStatement
	ConfirmationPrompt() LexIntentConfirmationPromptOutputReference
	ConfirmationPromptInput() *LexIntentConfirmationPrompt
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	CreateVersion() interface{}
	SetCreateVersion(val interface{})
	CreateVersionInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DialogCodeHook() LexIntentDialogCodeHookOutputReference
	DialogCodeHookInput() *LexIntentDialogCodeHook
	FollowUpPrompt() LexIntentFollowUpPromptOutputReference
	FollowUpPromptInput() *LexIntentFollowUpPrompt
	Fqn() *string
	FriendlyUniqueId() *string
	FulfillmentActivity() LexIntentFulfillmentActivityOutputReference
	FulfillmentActivityInput() *LexIntentFulfillmentActivity
	Id() *string
	LastUpdatedDate() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ParentIntentSignature() *string
	SetParentIntentSignature(val *string)
	ParentIntentSignatureInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RejectionStatement() LexIntentRejectionStatementOutputReference
	RejectionStatementInput() *LexIntentRejectionStatement
	SampleUtterances() *[]*string
	SetSampleUtterances(val *[]*string)
	SampleUtterancesInput() *[]*string
	Slot() interface{}
	SetSlot(val interface{})
	SlotInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() LexIntentTimeoutsOutputReference
	TimeoutsInput() *LexIntentTimeouts
	Version() *string
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
	PutConclusionStatement(value *LexIntentConclusionStatement)
	PutConfirmationPrompt(value *LexIntentConfirmationPrompt)
	PutDialogCodeHook(value *LexIntentDialogCodeHook)
	PutFollowUpPrompt(value *LexIntentFollowUpPrompt)
	PutFulfillmentActivity(value *LexIntentFulfillmentActivity)
	PutRejectionStatement(value *LexIntentRejectionStatement)
	PutTimeouts(value *LexIntentTimeouts)
	ResetConclusionStatement()
	ResetConfirmationPrompt()
	ResetCreateVersion()
	ResetDescription()
	ResetDialogCodeHook()
	ResetFollowUpPrompt()
	ResetOverrideLogicalId()
	ResetParentIntentSignature()
	ResetRejectionStatement()
	ResetSampleUtterances()
	ResetSlot()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LexIntent
type jsiiProxy_LexIntent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexIntent) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ConclusionStatement() LexIntentConclusionStatementOutputReference {
	var returns LexIntentConclusionStatementOutputReference
	_jsii_.Get(
		j,
		"conclusionStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ConclusionStatementInput() *LexIntentConclusionStatement {
	var returns *LexIntentConclusionStatement
	_jsii_.Get(
		j,
		"conclusionStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ConfirmationPrompt() LexIntentConfirmationPromptOutputReference {
	var returns LexIntentConfirmationPromptOutputReference
	_jsii_.Get(
		j,
		"confirmationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ConfirmationPromptInput() *LexIntentConfirmationPrompt {
	var returns *LexIntentConfirmationPrompt
	_jsii_.Get(
		j,
		"confirmationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) DialogCodeHook() LexIntentDialogCodeHookOutputReference {
	var returns LexIntentDialogCodeHookOutputReference
	_jsii_.Get(
		j,
		"dialogCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) DialogCodeHookInput() *LexIntentDialogCodeHook {
	var returns *LexIntentDialogCodeHook
	_jsii_.Get(
		j,
		"dialogCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) FollowUpPrompt() LexIntentFollowUpPromptOutputReference {
	var returns LexIntentFollowUpPromptOutputReference
	_jsii_.Get(
		j,
		"followUpPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) FollowUpPromptInput() *LexIntentFollowUpPrompt {
	var returns *LexIntentFollowUpPrompt
	_jsii_.Get(
		j,
		"followUpPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) FulfillmentActivity() LexIntentFulfillmentActivityOutputReference {
	var returns LexIntentFulfillmentActivityOutputReference
	_jsii_.Get(
		j,
		"fulfillmentActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) FulfillmentActivityInput() *LexIntentFulfillmentActivity {
	var returns *LexIntentFulfillmentActivity
	_jsii_.Get(
		j,
		"fulfillmentActivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) ParentIntentSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) RejectionStatement() LexIntentRejectionStatementOutputReference {
	var returns LexIntentRejectionStatementOutputReference
	_jsii_.Get(
		j,
		"rejectionStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) RejectionStatementInput() *LexIntentRejectionStatement {
	var returns *LexIntentRejectionStatement
	_jsii_.Get(
		j,
		"rejectionStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) SampleUtterances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) SampleUtterancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Slot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) SlotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Timeouts() LexIntentTimeoutsOutputReference {
	var returns LexIntentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) TimeoutsInput() *LexIntentTimeouts {
	var returns *LexIntentTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntent) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_intent aws_lex_intent} Resource.
func NewLexIntent(scope constructs.Construct, id *string, config *LexIntentConfig) LexIntent {
	_init_.Initialize()

	j := jsiiProxy_LexIntent{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_intent aws_lex_intent} Resource.
func NewLexIntent_Override(l LexIntent, scope constructs.Construct, id *string, config *LexIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntent",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexIntent) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetCreateVersion(val interface{}) {
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetParentIntentSignature(val *string) {
	_jsii_.Set(
		j,
		"parentIntentSignature",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetSampleUtterances(val *[]*string) {
	_jsii_.Set(
		j,
		"sampleUtterances",
		val,
	)
}

func (j *jsiiProxy_LexIntent) SetSlot(val interface{}) {
	_jsii_.Set(
		j,
		"slot",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LexIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.LexIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.LexIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LexIntent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexIntent) PutConclusionStatement(value *LexIntentConclusionStatement) {
	_jsii_.InvokeVoid(
		l,
		"putConclusionStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutConfirmationPrompt(value *LexIntentConfirmationPrompt) {
	_jsii_.InvokeVoid(
		l,
		"putConfirmationPrompt",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutDialogCodeHook(value *LexIntentDialogCodeHook) {
	_jsii_.InvokeVoid(
		l,
		"putDialogCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutFollowUpPrompt(value *LexIntentFollowUpPrompt) {
	_jsii_.InvokeVoid(
		l,
		"putFollowUpPrompt",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutFulfillmentActivity(value *LexIntentFulfillmentActivity) {
	_jsii_.InvokeVoid(
		l,
		"putFulfillmentActivity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutRejectionStatement(value *LexIntentRejectionStatement) {
	_jsii_.InvokeVoid(
		l,
		"putRejectionStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) PutTimeouts(value *LexIntentTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntent) ResetConclusionStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetConclusionStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetConfirmationPrompt() {
	_jsii_.InvokeVoid(
		l,
		"resetConfirmationPrompt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetDialogCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetDialogCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetFollowUpPrompt() {
	_jsii_.InvokeVoid(
		l,
		"resetFollowUpPrompt",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LexIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetParentIntentSignature() {
	_jsii_.InvokeVoid(
		l,
		"resetParentIntentSignature",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetRejectionStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetRejectionStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetSampleUtterances() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleUtterances",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetSlot() {
	_jsii_.InvokeVoid(
		l,
		"resetSlot",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LexIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LexIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LexIntentConclusionStatement struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentConclusionStatementMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentConclusionStatementOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentConclusionStatement
	SetInternalValue(val *LexIntentConclusionStatement)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentConclusionStatementOutputReference
type jsiiProxy_LexIntentConclusionStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) InternalValue() *LexIntentConclusionStatement {
	var returns *LexIntentConclusionStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentConclusionStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentConclusionStatementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentConclusionStatementOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentConclusionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentConclusionStatementOutputReference_Override(l LexIntentConclusionStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentConclusionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetInternalValue(val *LexIntentConclusionStatement) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentConclusionStatementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentConclusionStatementOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

// AWS Lex.
type LexIntentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// fulfillment_activity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#fulfillment_activity LexIntent#fulfillment_activity}
	FulfillmentActivity *LexIntentFulfillmentActivity `json:"fulfillmentActivity" yaml:"fulfillmentActivity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#name LexIntent#name}.
	Name *string `json:"name" yaml:"name"`
	// conclusion_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#conclusion_statement LexIntent#conclusion_statement}
	ConclusionStatement *LexIntentConclusionStatement `json:"conclusionStatement" yaml:"conclusionStatement"`
	// confirmation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#confirmation_prompt LexIntent#confirmation_prompt}
	ConfirmationPrompt *LexIntentConfirmationPrompt `json:"confirmationPrompt" yaml:"confirmationPrompt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#create_version LexIntent#create_version}.
	CreateVersion interface{} `json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#description LexIntent#description}.
	Description *string `json:"description" yaml:"description"`
	// dialog_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#dialog_code_hook LexIntent#dialog_code_hook}
	DialogCodeHook *LexIntentDialogCodeHook `json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// follow_up_prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#follow_up_prompt LexIntent#follow_up_prompt}
	FollowUpPrompt *LexIntentFollowUpPrompt `json:"followUpPrompt" yaml:"followUpPrompt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#parent_intent_signature LexIntent#parent_intent_signature}.
	ParentIntentSignature *string `json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#rejection_statement LexIntent#rejection_statement}
	RejectionStatement *LexIntentRejectionStatement `json:"rejectionStatement" yaml:"rejectionStatement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#sample_utterances LexIntent#sample_utterances}.
	SampleUtterances *[]*string `json:"sampleUtterances" yaml:"sampleUtterances"`
	// slot block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#slot LexIntent#slot}
	Slot interface{} `json:"slot" yaml:"slot"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#timeouts LexIntent#timeouts}
	Timeouts *LexIntentTimeouts `json:"timeouts" yaml:"timeouts"`
}

type LexIntentConfirmationPrompt struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#max_attempts LexIntent#max_attempts}.
	MaxAttempts *float64 `json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentConfirmationPromptMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentConfirmationPromptOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentConfirmationPrompt
	SetInternalValue(val *LexIntentConfirmationPrompt)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentConfirmationPromptOutputReference
type jsiiProxy_LexIntentConfirmationPromptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) InternalValue() *LexIntentConfirmationPrompt {
	var returns *LexIntentConfirmationPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentConfirmationPromptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentConfirmationPromptOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentConfirmationPromptOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentConfirmationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentConfirmationPromptOutputReference_Override(l LexIntentConfirmationPromptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentConfirmationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetInternalValue(val *LexIntentConfirmationPrompt) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetMaxAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentConfirmationPromptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentConfirmationPromptOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

type LexIntentDialogCodeHook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message_version LexIntent#message_version}.
	MessageVersion *string `json:"messageVersion" yaml:"messageVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#uri LexIntent#uri}.
	Uri *string `json:"uri" yaml:"uri"`
}

type LexIntentDialogCodeHookOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentDialogCodeHook
	SetInternalValue(val *LexIntentDialogCodeHook)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageVersion() *string
	SetMessageVersion(val *string)
	MessageVersionInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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

// The jsii proxy struct for LexIntentDialogCodeHookOutputReference
type jsiiProxy_LexIntentDialogCodeHookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) InternalValue() *LexIntentDialogCodeHook {
	var returns *LexIntentDialogCodeHook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) MessageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) MessageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func NewLexIntentDialogCodeHookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentDialogCodeHookOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentDialogCodeHookOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentDialogCodeHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentDialogCodeHookOutputReference_Override(l LexIntentDialogCodeHookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentDialogCodeHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetInternalValue(val *LexIntentDialogCodeHook) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetMessageVersion(val *string) {
	_jsii_.Set(
		j,
		"messageVersion",
		val,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexIntentDialogCodeHookOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentDialogCodeHookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LexIntentFollowUpPrompt struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#prompt LexIntent#prompt}
	Prompt *LexIntentFollowUpPromptPrompt `json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#rejection_statement LexIntent#rejection_statement}
	RejectionStatement *LexIntentFollowUpPromptRejectionStatement `json:"rejectionStatement" yaml:"rejectionStatement"`
}

type LexIntentFollowUpPromptOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentFollowUpPrompt
	SetInternalValue(val *LexIntentFollowUpPrompt)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Prompt() LexIntentFollowUpPromptPromptOutputReference
	PromptInput() *LexIntentFollowUpPromptPrompt
	RejectionStatement() LexIntentFollowUpPromptRejectionStatementOutputReference
	RejectionStatementInput() *LexIntentFollowUpPromptRejectionStatement
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
	PutPrompt(value *LexIntentFollowUpPromptPrompt)
	PutRejectionStatement(value *LexIntentFollowUpPromptRejectionStatement)
}

// The jsii proxy struct for LexIntentFollowUpPromptOutputReference
type jsiiProxy_LexIntentFollowUpPromptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) InternalValue() *LexIntentFollowUpPrompt {
	var returns *LexIntentFollowUpPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) Prompt() LexIntentFollowUpPromptPromptOutputReference {
	var returns LexIntentFollowUpPromptPromptOutputReference
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) PromptInput() *LexIntentFollowUpPromptPrompt {
	var returns *LexIntentFollowUpPromptPrompt
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) RejectionStatement() LexIntentFollowUpPromptRejectionStatementOutputReference {
	var returns LexIntentFollowUpPromptRejectionStatementOutputReference
	_jsii_.Get(
		j,
		"rejectionStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) RejectionStatementInput() *LexIntentFollowUpPromptRejectionStatement {
	var returns *LexIntentFollowUpPromptRejectionStatement
	_jsii_.Get(
		j,
		"rejectionStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentFollowUpPromptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentFollowUpPromptOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentFollowUpPromptOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentFollowUpPromptOutputReference_Override(l LexIntentFollowUpPromptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) SetInternalValue(val *LexIntentFollowUpPrompt) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) PutPrompt(value *LexIntentFollowUpPromptPrompt) {
	_jsii_.InvokeVoid(
		l,
		"putPrompt",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntentFollowUpPromptOutputReference) PutRejectionStatement(value *LexIntentFollowUpPromptRejectionStatement) {
	_jsii_.InvokeVoid(
		l,
		"putRejectionStatement",
		[]interface{}{value},
	)
}

type LexIntentFollowUpPromptPrompt struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#max_attempts LexIntent#max_attempts}.
	MaxAttempts *float64 `json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentFollowUpPromptPromptMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentFollowUpPromptPromptOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentFollowUpPromptPrompt
	SetInternalValue(val *LexIntentFollowUpPromptPrompt)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentFollowUpPromptPromptOutputReference
type jsiiProxy_LexIntentFollowUpPromptPromptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) InternalValue() *LexIntentFollowUpPromptPrompt {
	var returns *LexIntentFollowUpPromptPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentFollowUpPromptPromptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentFollowUpPromptPromptOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentFollowUpPromptPromptOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentFollowUpPromptPromptOutputReference_Override(l LexIntentFollowUpPromptPromptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetInternalValue(val *LexIntentFollowUpPromptPrompt) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetMaxAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptPromptOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

type LexIntentFollowUpPromptRejectionStatement struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentFollowUpPromptRejectionStatementMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentFollowUpPromptRejectionStatementOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentFollowUpPromptRejectionStatement
	SetInternalValue(val *LexIntentFollowUpPromptRejectionStatement)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentFollowUpPromptRejectionStatementOutputReference
type jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) InternalValue() *LexIntentFollowUpPromptRejectionStatement {
	var returns *LexIntentFollowUpPromptRejectionStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentFollowUpPromptRejectionStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentFollowUpPromptRejectionStatementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptRejectionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentFollowUpPromptRejectionStatementOutputReference_Override(l LexIntentFollowUpPromptRejectionStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFollowUpPromptRejectionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetInternalValue(val *LexIntentFollowUpPromptRejectionStatement) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFollowUpPromptRejectionStatementOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

type LexIntentFulfillmentActivity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#type LexIntent#type}.
	Type *string `json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#code_hook LexIntent#code_hook}
	CodeHook *LexIntentFulfillmentActivityCodeHook `json:"codeHook" yaml:"codeHook"`
}

type LexIntentFulfillmentActivityCodeHook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message_version LexIntent#message_version}.
	MessageVersion *string `json:"messageVersion" yaml:"messageVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#uri LexIntent#uri}.
	Uri *string `json:"uri" yaml:"uri"`
}

type LexIntentFulfillmentActivityCodeHookOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentFulfillmentActivityCodeHook
	SetInternalValue(val *LexIntentFulfillmentActivityCodeHook)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageVersion() *string
	SetMessageVersion(val *string)
	MessageVersionInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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

// The jsii proxy struct for LexIntentFulfillmentActivityCodeHookOutputReference
type jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) InternalValue() *LexIntentFulfillmentActivityCodeHook {
	var returns *LexIntentFulfillmentActivityCodeHook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) MessageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) MessageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func NewLexIntentFulfillmentActivityCodeHookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentFulfillmentActivityCodeHookOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFulfillmentActivityCodeHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentFulfillmentActivityCodeHookOutputReference_Override(l LexIntentFulfillmentActivityCodeHookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFulfillmentActivityCodeHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetInternalValue(val *LexIntentFulfillmentActivityCodeHook) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetMessageVersion(val *string) {
	_jsii_.Set(
		j,
		"messageVersion",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityCodeHookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type LexIntentFulfillmentActivityOutputReference interface {
	cdktf.ComplexObject
	CodeHook() LexIntentFulfillmentActivityCodeHookOutputReference
	CodeHookInput() *LexIntentFulfillmentActivityCodeHook
	InternalValue() *LexIntentFulfillmentActivity
	SetInternalValue(val *LexIntentFulfillmentActivity)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutCodeHook(value *LexIntentFulfillmentActivityCodeHook)
	ResetCodeHook()
}

// The jsii proxy struct for LexIntentFulfillmentActivityOutputReference
type jsiiProxy_LexIntentFulfillmentActivityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) CodeHook() LexIntentFulfillmentActivityCodeHookOutputReference {
	var returns LexIntentFulfillmentActivityCodeHookOutputReference
	_jsii_.Get(
		j,
		"codeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) CodeHookInput() *LexIntentFulfillmentActivityCodeHook {
	var returns *LexIntentFulfillmentActivityCodeHook
	_jsii_.Get(
		j,
		"codeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) InternalValue() *LexIntentFulfillmentActivity {
	var returns *LexIntentFulfillmentActivity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewLexIntentFulfillmentActivityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentFulfillmentActivityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentFulfillmentActivityOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFulfillmentActivityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentFulfillmentActivityOutputReference_Override(l LexIntentFulfillmentActivityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentFulfillmentActivityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) SetInternalValue(val *LexIntentFulfillmentActivity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexIntentFulfillmentActivityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) PutCodeHook(value *LexIntentFulfillmentActivityCodeHook) {
	_jsii_.InvokeVoid(
		l,
		"putCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexIntentFulfillmentActivityOutputReference) ResetCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeHook",
		nil, // no parameters
	)
}

type LexIntentRejectionStatement struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentRejectionStatementMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentRejectionStatementOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentRejectionStatement
	SetInternalValue(val *LexIntentRejectionStatement)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentRejectionStatementOutputReference
type jsiiProxy_LexIntentRejectionStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) InternalValue() *LexIntentRejectionStatement {
	var returns *LexIntentRejectionStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentRejectionStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentRejectionStatementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentRejectionStatementOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentRejectionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentRejectionStatementOutputReference_Override(l LexIntentRejectionStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentRejectionStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetInternalValue(val *LexIntentRejectionStatement) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentRejectionStatementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentRejectionStatementOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

type LexIntentSlot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#name LexIntent#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#slot_constraint LexIntent#slot_constraint}.
	SlotConstraint *string `json:"slotConstraint" yaml:"slotConstraint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#slot_type LexIntent#slot_type}.
	SlotType *string `json:"slotType" yaml:"slotType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#description LexIntent#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#priority LexIntent#priority}.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#sample_utterances LexIntent#sample_utterances}.
	SampleUtterances *[]*string `json:"sampleUtterances" yaml:"sampleUtterances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#slot_type_version LexIntent#slot_type_version}.
	SlotTypeVersion *string `json:"slotTypeVersion" yaml:"slotTypeVersion"`
	// value_elicitation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#value_elicitation_prompt LexIntent#value_elicitation_prompt}
	ValueElicitationPrompt *LexIntentSlotValueElicitationPrompt `json:"valueElicitationPrompt" yaml:"valueElicitationPrompt"`
}

type LexIntentSlotValueElicitationPrompt struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#max_attempts LexIntent#max_attempts}.
	MaxAttempts *float64 `json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message LexIntent#message}
	Message interface{} `json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `json:"responseCard" yaml:"responseCard"`
}

type LexIntentSlotValueElicitationPromptMessage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content LexIntent#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `json:"groupNumber" yaml:"groupNumber"`
}

type LexIntentSlotValueElicitationPromptOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *LexIntentSlotValueElicitationPrompt
	SetInternalValue(val *LexIntentSlotValueElicitationPrompt)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Message() interface{}
	SetMessage(val interface{})
	MessageInput() interface{}
	ResponseCard() *string
	SetResponseCard(val *string)
	ResponseCardInput() *string
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
	ResetResponseCard()
}

// The jsii proxy struct for LexIntentSlotValueElicitationPromptOutputReference
type jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) InternalValue() *LexIntentSlotValueElicitationPrompt {
	var returns *LexIntentSlotValueElicitationPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) Message() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLexIntentSlotValueElicitationPromptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentSlotValueElicitationPromptOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentSlotValueElicitationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentSlotValueElicitationPromptOutputReference_Override(l LexIntentSlotValueElicitationPromptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentSlotValueElicitationPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetInternalValue(val *LexIntentSlotValueElicitationPrompt) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetMaxAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetMessage(val interface{}) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetResponseCard(val *string) {
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentSlotValueElicitationPromptOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetResponseCard",
		nil, // no parameters
	)
}

type LexIntentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#create LexIntent#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#delete LexIntent#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#update LexIntent#update}.
	Update *string `json:"update" yaml:"update"`
}

type LexIntentTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *LexIntentTimeouts
	SetInternalValue(val *LexIntentTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for LexIntentTimeoutsOutputReference
type jsiiProxy_LexIntentTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) InternalValue() *LexIntentTimeouts {
	var returns *LexIntentTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewLexIntentTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexIntentTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexIntentTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexIntentTimeoutsOutputReference_Override(l LexIntentTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexIntentTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetInternalValue(val *LexIntentTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexIntentTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexIntentTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexIntentTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntentTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		l,
		"resetDelete",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexIntentTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type aws_lex_slot_type}.
type LexSlotType interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	CreateVersion() interface{}
	SetCreateVersion(val interface{})
	CreateVersionInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnumerationValue() interface{}
	SetEnumerationValue(val interface{})
	EnumerationValueInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedDate() *string
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
	Timeouts() LexSlotTypeTimeoutsOutputReference
	TimeoutsInput() *LexSlotTypeTimeouts
	ValueSelectionStrategy() *string
	SetValueSelectionStrategy(val *string)
	ValueSelectionStrategyInput() *string
	Version() *string
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
	PutTimeouts(value *LexSlotTypeTimeouts)
	ResetCreateVersion()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetValueSelectionStrategy()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LexSlotType
type jsiiProxy_LexSlotType struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexSlotType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) EnumerationValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enumerationValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) EnumerationValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enumerationValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Timeouts() LexSlotTypeTimeoutsOutputReference {
	var returns LexSlotTypeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) TimeoutsInput() *LexSlotTypeTimeouts {
	var returns *LexSlotTypeTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) ValueSelectionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueSelectionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) ValueSelectionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueSelectionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotType) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type aws_lex_slot_type} Resource.
func NewLexSlotType(scope constructs.Construct, id *string, config *LexSlotTypeConfig) LexSlotType {
	_init_.Initialize()

	j := jsiiProxy_LexSlotType{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexSlotType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type aws_lex_slot_type} Resource.
func NewLexSlotType_Override(l LexSlotType, scope constructs.Construct, id *string, config *LexSlotTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexSlotType",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexSlotType) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetCreateVersion(val interface{}) {
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetEnumerationValue(val interface{}) {
	_jsii_.Set(
		j,
		"enumerationValue",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LexSlotType) SetValueSelectionStrategy(val *string) {
	_jsii_.Set(
		j,
		"valueSelectionStrategy",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LexSlotType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lex.LexSlotType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexSlotType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lex.LexSlotType",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LexSlotType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexSlotType) PutTimeouts(value *LexSlotTypeTimeouts) {
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexSlotType) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotType) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LexSlotType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotType) ResetValueSelectionStrategy() {
	_jsii_.InvokeVoid(
		l,
		"resetValueSelectionStrategy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LexSlotType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LexSlotType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lex.
type LexSlotTypeConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// enumeration_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#enumeration_value LexSlotType#enumeration_value}
	EnumerationValue interface{} `json:"enumerationValue" yaml:"enumerationValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#name LexSlotType#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#create_version LexSlotType#create_version}.
	CreateVersion interface{} `json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#description LexSlotType#description}.
	Description *string `json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#timeouts LexSlotType#timeouts}
	Timeouts *LexSlotTypeTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#value_selection_strategy LexSlotType#value_selection_strategy}.
	ValueSelectionStrategy *string `json:"valueSelectionStrategy" yaml:"valueSelectionStrategy"`
}

type LexSlotTypeEnumerationValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#value LexSlotType#value}.
	Value *string `json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#synonyms LexSlotType#synonyms}.
	Synonyms *[]*string `json:"synonyms" yaml:"synonyms"`
}

type LexSlotTypeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#create LexSlotType#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#delete LexSlotType#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#update LexSlotType#update}.
	Update *string `json:"update" yaml:"update"`
}

type LexSlotTypeTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *LexSlotTypeTimeouts
	SetInternalValue(val *LexSlotTypeTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for LexSlotTypeTimeoutsOutputReference
type jsiiProxy_LexSlotTypeTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) InternalValue() *LexSlotTypeTimeouts {
	var returns *LexSlotTypeTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewLexSlotTypeTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LexSlotTypeTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LexSlotTypeTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lex.LexSlotTypeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLexSlotTypeTimeoutsOutputReference_Override(l LexSlotTypeTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lex.LexSlotTypeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetInternalValue(val *LexSlotTypeTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LexSlotTypeTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		l,
		"resetCreate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		l,
		"resetDelete",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexSlotTypeTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetUpdate",
		nil, // no parameters
	)
}
