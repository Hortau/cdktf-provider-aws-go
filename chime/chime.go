package chime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/chime/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html aws_chime_voice_connector}.
type ChimeVoiceConnector interface {
	cdktf.TerraformResource
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
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
	OutboundHostName() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequireEncryption() interface{}
	SetRequireEncryption(val interface{})
	RequireEncryptionInput() interface{}
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
	ResetAwsRegion()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnector
type jsiiProxy_ChimeVoiceConnector struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnector) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) OutboundHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) RequireEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) RequireEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html aws_chime_voice_connector} Resource.
func NewChimeVoiceConnector(scope constructs.Construct, id *string, config *ChimeVoiceConnectorConfig) ChimeVoiceConnector {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnector{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html aws_chime_voice_connector} Resource.
func NewChimeVoiceConnector_Override(c ChimeVoiceConnector, scope constructs.Construct, id *string, config *ChimeVoiceConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnector",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetAwsRegion(val *string) {
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnector) SetRequireEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"requireEncryption",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnector) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnector) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnector) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnector) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnector) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsRegion",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnector) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnector) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnector) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html#name ChimeVoiceConnector#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html#require_encryption ChimeVoiceConnector#require_encryption}.
	RequireEncryption interface{} `json:"requireEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector.html#aws_region ChimeVoiceConnector#aws_region}.
	AwsRegion *string `json:"awsRegion"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html aws_chime_voice_connector_group}.
type ChimeVoiceConnectorGroup interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Connector() *[]*ChimeVoiceConnectorGroupConnector
	SetConnector(val *[]*ChimeVoiceConnectorGroupConnector)
	ConnectorInput() *[]*ChimeVoiceConnectorGroupConnector
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
	ResetConnector()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorGroup
type jsiiProxy_ChimeVoiceConnectorGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Connector() *[]*ChimeVoiceConnectorGroupConnector {
	var returns *[]*ChimeVoiceConnectorGroupConnector
	_jsii_.Get(
		j,
		"connector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) ConnectorInput() *[]*ChimeVoiceConnectorGroupConnector {
	var returns *[]*ChimeVoiceConnectorGroupConnector
	_jsii_.Get(
		j,
		"connectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html aws_chime_voice_connector_group} Resource.
func NewChimeVoiceConnectorGroup(scope constructs.Construct, id *string, config *ChimeVoiceConnectorGroupConfig) ChimeVoiceConnectorGroup {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorGroup{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html aws_chime_voice_connector_group} Resource.
func NewChimeVoiceConnectorGroup_Override(c ChimeVoiceConnectorGroup, scope constructs.Construct, id *string, config *ChimeVoiceConnectorGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetConnector(val *[]*ChimeVoiceConnectorGroupConnector) {
	_jsii_.Set(
		j,
		"connector",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorGroup) SetProvider(val cdktf.TerraformProvider) {
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
func ChimeVoiceConnectorGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorGroup) ResetConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetConnector",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html#name ChimeVoiceConnectorGroup#name}.
	Name *string `json:"name"`
	// connector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html#connector ChimeVoiceConnectorGroup#connector}
	Connector *[]*ChimeVoiceConnectorGroupConnector `json:"connector"`
}

type ChimeVoiceConnectorGroupConnector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html#priority ChimeVoiceConnectorGroup#priority}.
	Priority *float64 `json:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_group.html#voice_connector_id ChimeVoiceConnectorGroup#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_logging.html aws_chime_voice_connector_logging}.
type ChimeVoiceConnectorLogging interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnableSipLogs() interface{}
	SetEnableSipLogs(val interface{})
	EnableSipLogsInput() interface{}
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
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetEnableSipLogs()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorLogging
type jsiiProxy_ChimeVoiceConnectorLogging struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) EnableSipLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSipLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) EnableSipLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSipLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_logging.html aws_chime_voice_connector_logging} Resource.
func NewChimeVoiceConnectorLogging(scope constructs.Construct, id *string, config *ChimeVoiceConnectorLoggingConfig) ChimeVoiceConnectorLogging {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorLogging{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorLogging",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_logging.html aws_chime_voice_connector_logging} Resource.
func NewChimeVoiceConnectorLogging_Override(c ChimeVoiceConnectorLogging, scope constructs.Construct, id *string, config *ChimeVoiceConnectorLoggingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorLogging",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetEnableSipLogs(val interface{}) {
	_jsii_.Set(
		j,
		"enableSipLogs",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorLogging) SetVoiceConnectorId(val *string) {
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnectorLogging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorLogging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorLogging_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorLogging",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorLogging) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorLogging) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorLogging) ResetEnableSipLogs() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableSipLogs",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorLogging) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorLogging) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorLogging) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorLoggingConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_logging.html#voice_connector_id ChimeVoiceConnectorLogging#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_logging.html#enable_sip_logs ChimeVoiceConnectorLogging#enable_sip_logs}.
	EnableSipLogs interface{} `json:"enableSipLogs"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html aws_chime_voice_connector_origination}.
type ChimeVoiceConnectorOrigination interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Route() *[]*ChimeVoiceConnectorOriginationRoute
	SetRoute(val *[]*ChimeVoiceConnectorOriginationRoute)
	RouteInput() *[]*ChimeVoiceConnectorOriginationRoute
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDisabled()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorOrigination
type jsiiProxy_ChimeVoiceConnectorOrigination struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) Route() *[]*ChimeVoiceConnectorOriginationRoute {
	var returns *[]*ChimeVoiceConnectorOriginationRoute
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) RouteInput() *[]*ChimeVoiceConnectorOriginationRoute {
	var returns *[]*ChimeVoiceConnectorOriginationRoute
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html aws_chime_voice_connector_origination} Resource.
func NewChimeVoiceConnectorOrigination(scope constructs.Construct, id *string, config *ChimeVoiceConnectorOriginationConfig) ChimeVoiceConnectorOrigination {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorOrigination{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorOrigination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html aws_chime_voice_connector_origination} Resource.
func NewChimeVoiceConnectorOrigination_Override(c ChimeVoiceConnectorOrigination, scope constructs.Construct, id *string, config *ChimeVoiceConnectorOriginationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorOrigination",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetRoute(val *[]*ChimeVoiceConnectorOriginationRoute) {
	_jsii_.Set(
		j,
		"route",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorOrigination) SetVoiceConnectorId(val *string) {
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnectorOrigination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorOrigination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorOrigination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorOrigination",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorOrigination) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorOrigination) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorOrigination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorOriginationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#route ChimeVoiceConnectorOrigination#route}
	Route *[]*ChimeVoiceConnectorOriginationRoute `json:"route"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#voice_connector_id ChimeVoiceConnectorOrigination#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#disabled ChimeVoiceConnectorOrigination#disabled}.
	Disabled interface{} `json:"disabled"`
}

type ChimeVoiceConnectorOriginationRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#host ChimeVoiceConnectorOrigination#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#priority ChimeVoiceConnectorOrigination#priority}.
	Priority *float64 `json:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#protocol ChimeVoiceConnectorOrigination#protocol}.
	Protocol *string `json:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#weight ChimeVoiceConnectorOrigination#weight}.
	Weight *float64 `json:"weight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_origination.html#port ChimeVoiceConnectorOrigination#port}.
	Port *float64 `json:"port"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html aws_chime_voice_connector_streaming}.
type ChimeVoiceConnectorStreaming interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DataRetention() *float64
	SetDataRetention(val *float64)
	DataRetentionInput() *float64
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StreamingNotificationTargets() *[]*string
	SetStreamingNotificationTargets(val *[]*string)
	StreamingNotificationTargetsInput() *[]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDisabled()
	ResetOverrideLogicalId()
	ResetStreamingNotificationTargets()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorStreaming
type jsiiProxy_ChimeVoiceConnectorStreaming struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DataRetention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DataRetentionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) StreamingNotificationTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingNotificationTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) StreamingNotificationTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingNotificationTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html aws_chime_voice_connector_streaming} Resource.
func NewChimeVoiceConnectorStreaming(scope constructs.Construct, id *string, config *ChimeVoiceConnectorStreamingConfig) ChimeVoiceConnectorStreaming {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorStreaming{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorStreaming",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html aws_chime_voice_connector_streaming} Resource.
func NewChimeVoiceConnectorStreaming_Override(c ChimeVoiceConnectorStreaming, scope constructs.Construct, id *string, config *ChimeVoiceConnectorStreamingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorStreaming",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetDataRetention(val *float64) {
	_jsii_.Set(
		j,
		"dataRetention",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetStreamingNotificationTargets(val *[]*string) {
	_jsii_.Set(
		j,
		"streamingNotificationTargets",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) SetVoiceConnectorId(val *string) {
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnectorStreaming_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorStreaming",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorStreaming_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorStreaming",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetStreamingNotificationTargets() {
	_jsii_.InvokeVoid(
		c,
		"resetStreamingNotificationTargets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorStreamingConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html#data_retention ChimeVoiceConnectorStreaming#data_retention}.
	DataRetention *float64 `json:"dataRetention"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html#voice_connector_id ChimeVoiceConnectorStreaming#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html#disabled ChimeVoiceConnectorStreaming#disabled}.
	Disabled interface{} `json:"disabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_streaming.html#streaming_notification_targets ChimeVoiceConnectorStreaming#streaming_notification_targets}.
	StreamingNotificationTargets *[]*string `json:"streamingNotificationTargets"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html aws_chime_voice_connector_termination}.
type ChimeVoiceConnectorTermination interface {
	cdktf.TerraformResource
	CallingRegions() *[]*string
	SetCallingRegions(val *[]*string)
	CallingRegionsInput() *[]*string
	CdktfStack() cdktf.TerraformStack
	CidrAllowList() *[]*string
	SetCidrAllowList(val *[]*string)
	CidrAllowListInput() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CpsLimit() *float64
	SetCpsLimit(val *float64)
	CpsLimitInput() *float64
	DefaultPhoneNumber() *string
	SetDefaultPhoneNumber(val *string)
	DefaultPhoneNumberInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
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
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetCpsLimit()
	ResetDefaultPhoneNumber()
	ResetDisabled()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorTermination
type jsiiProxy_ChimeVoiceConnectorTermination struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CallingRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callingRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CallingRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callingRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CidrAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CidrAllowListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrAllowListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CpsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) CpsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) DefaultPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) DefaultPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html aws_chime_voice_connector_termination} Resource.
func NewChimeVoiceConnectorTermination(scope constructs.Construct, id *string, config *ChimeVoiceConnectorTerminationConfig) ChimeVoiceConnectorTermination {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorTermination{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTermination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html aws_chime_voice_connector_termination} Resource.
func NewChimeVoiceConnectorTermination_Override(c ChimeVoiceConnectorTermination, scope constructs.Construct, id *string, config *ChimeVoiceConnectorTerminationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTermination",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetCallingRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"callingRegions",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetCidrAllowList(val *[]*string) {
	_jsii_.Set(
		j,
		"cidrAllowList",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetCpsLimit(val *float64) {
	_jsii_.Set(
		j,
		"cpsLimit",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetDefaultPhoneNumber(val *string) {
	_jsii_.Set(
		j,
		"defaultPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTermination) SetVoiceConnectorId(val *string) {
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnectorTermination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTermination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorTermination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTermination",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTermination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTermination) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorTermination) ResetCpsLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetCpsLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorTermination) ResetDefaultPhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultPhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorTermination) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTermination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorTermination) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTermination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorTerminationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#calling_regions ChimeVoiceConnectorTermination#calling_regions}.
	CallingRegions *[]*string `json:"callingRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#cidr_allow_list ChimeVoiceConnectorTermination#cidr_allow_list}.
	CidrAllowList *[]*string `json:"cidrAllowList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#voice_connector_id ChimeVoiceConnectorTermination#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#cps_limit ChimeVoiceConnectorTermination#cps_limit}.
	CpsLimit *float64 `json:"cpsLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#default_phone_number ChimeVoiceConnectorTermination#default_phone_number}.
	DefaultPhoneNumber *string `json:"defaultPhoneNumber"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html#disabled ChimeVoiceConnectorTermination#disabled}.
	Disabled interface{} `json:"disabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html aws_chime_voice_connector_termination_credentials}.
type ChimeVoiceConnectorTerminationCredentials interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Credentials() *[]*ChimeVoiceConnectorTerminationCredentialsCredentials
	SetCredentials(val *[]*ChimeVoiceConnectorTerminationCredentialsCredentials)
	CredentialsInput() *[]*ChimeVoiceConnectorTerminationCredentialsCredentials
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
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
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

// The jsii proxy struct for ChimeVoiceConnectorTerminationCredentials
type jsiiProxy_ChimeVoiceConnectorTerminationCredentials struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Credentials() *[]*ChimeVoiceConnectorTerminationCredentialsCredentials {
	var returns *[]*ChimeVoiceConnectorTerminationCredentialsCredentials
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) CredentialsInput() *[]*ChimeVoiceConnectorTerminationCredentialsCredentials {
	var returns *[]*ChimeVoiceConnectorTerminationCredentialsCredentials
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html aws_chime_voice_connector_termination_credentials} Resource.
func NewChimeVoiceConnectorTerminationCredentials(scope constructs.Construct, id *string, config *ChimeVoiceConnectorTerminationCredentialsConfig) ChimeVoiceConnectorTerminationCredentials {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorTerminationCredentials{}

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTerminationCredentials",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html aws_chime_voice_connector_termination_credentials} Resource.
func NewChimeVoiceConnectorTerminationCredentials_Override(c ChimeVoiceConnectorTerminationCredentials, scope constructs.Construct, id *string, config *ChimeVoiceConnectorTerminationCredentialsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTerminationCredentials",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetCredentials(val *[]*ChimeVoiceConnectorTerminationCredentialsCredentials) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SetVoiceConnectorId(val *string) {
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ChimeVoiceConnectorTerminationCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTerminationCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorTerminationCredentials_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Chime.ChimeVoiceConnectorTerminationCredentials",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) ToMetadata() interface{} {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) ToString() *string {
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
func (c *jsiiProxy_ChimeVoiceConnectorTerminationCredentials) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChimeVoiceConnectorTerminationCredentialsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html#credentials ChimeVoiceConnectorTerminationCredentials#credentials}
	Credentials *[]*ChimeVoiceConnectorTerminationCredentialsCredentials `json:"credentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html#voice_connector_id ChimeVoiceConnectorTerminationCredentials#voice_connector_id}.
	VoiceConnectorId *string `json:"voiceConnectorId"`
}

type ChimeVoiceConnectorTerminationCredentialsCredentials struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html#password ChimeVoiceConnectorTerminationCredentials#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination_credentials.html#username ChimeVoiceConnectorTerminationCredentials#username}.
	Username *string `json:"username"`
}
