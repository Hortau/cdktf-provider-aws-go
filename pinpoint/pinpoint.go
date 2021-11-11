package pinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/pinpoint/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html aws_pinpoint_adm_channel}.
type PinpointAdmChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	ResetEnabled()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointAdmChannel
type jsiiProxy_PinpointAdmChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointAdmChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAdmChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html aws_pinpoint_adm_channel} Resource.
func NewPinpointAdmChannel(scope constructs.Construct, id *string, config *PinpointAdmChannelConfig) PinpointAdmChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointAdmChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAdmChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html aws_pinpoint_adm_channel} Resource.
func NewPinpointAdmChannel_Override(p PinpointAdmChannel, scope constructs.Construct, id *string, config *PinpointAdmChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAdmChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointAdmChannel) SetProvider(val cdktf.TerraformProvider) {
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
func PinpointAdmChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointAdmChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointAdmChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointAdmChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointAdmChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAdmChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointAdmChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointAdmChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointAdmChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html#application_id PinpointAdmChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html#client_id PinpointAdmChannel#client_id}.
	ClientId *string `json:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html#client_secret PinpointAdmChannel#client_secret}.
	ClientSecret *string `json:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_adm_channel.html#enabled PinpointAdmChannel#enabled}.
	Enabled interface{} `json:"enabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html aws_pinpoint_apns_channel}.
type PinpointApnsChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	DefaultAuthenticationMethodInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenKey() *string
	SetTokenKey(val *string)
	TokenKeyId() *string
	SetTokenKeyId(val *string)
	TokenKeyIdInput() *string
	TokenKeyInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetBundleId()
	ResetCertificate()
	ResetDefaultAuthenticationMethod()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetTeamId()
	ResetTokenKey()
	ResetTokenKeyId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApnsChannel
type jsiiProxy_PinpointApnsChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html aws_pinpoint_apns_channel} Resource.
func NewPinpointApnsChannel(scope constructs.Construct, id *string, config *PinpointApnsChannelConfig) PinpointApnsChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointApnsChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html aws_pinpoint_apns_channel} Resource.
func NewPinpointApnsChannel_Override(p PinpointApnsChannel, scope constructs.Construct, id *string, config *PinpointApnsChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointApnsChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointApnsChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointApnsChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointApnsChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointApnsChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointApnsChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#application_id PinpointApnsChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#bundle_id PinpointApnsChannel#bundle_id}.
	BundleId *string `json:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#certificate PinpointApnsChannel#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#default_authentication_method PinpointApnsChannel#default_authentication_method}.
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#enabled PinpointApnsChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#private_key PinpointApnsChannel#private_key}.
	PrivateKey *string `json:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#team_id PinpointApnsChannel#team_id}.
	TeamId *string `json:"teamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#token_key PinpointApnsChannel#token_key}.
	TokenKey *string `json:"tokenKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_channel.html#token_key_id PinpointApnsChannel#token_key_id}.
	TokenKeyId *string `json:"tokenKeyId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html aws_pinpoint_apns_sandbox_channel}.
type PinpointApnsSandboxChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	DefaultAuthenticationMethodInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenKey() *string
	SetTokenKey(val *string)
	TokenKeyId() *string
	SetTokenKeyId(val *string)
	TokenKeyIdInput() *string
	TokenKeyInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetBundleId()
	ResetCertificate()
	ResetDefaultAuthenticationMethod()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetTeamId()
	ResetTokenKey()
	ResetTokenKeyId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApnsSandboxChannel
type jsiiProxy_PinpointApnsSandboxChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html aws_pinpoint_apns_sandbox_channel} Resource.
func NewPinpointApnsSandboxChannel(scope constructs.Construct, id *string, config *PinpointApnsSandboxChannelConfig) PinpointApnsSandboxChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointApnsSandboxChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsSandboxChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html aws_pinpoint_apns_sandbox_channel} Resource.
func NewPinpointApnsSandboxChannel_Override(p PinpointApnsSandboxChannel, scope constructs.Construct, id *string, config *PinpointApnsSandboxChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsSandboxChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsSandboxChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointApnsSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointApnsSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsSandboxChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointApnsSandboxChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsSandboxChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointApnsSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointApnsSandboxChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointApnsSandboxChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#application_id PinpointApnsSandboxChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#bundle_id PinpointApnsSandboxChannel#bundle_id}.
	BundleId *string `json:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#certificate PinpointApnsSandboxChannel#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#default_authentication_method PinpointApnsSandboxChannel#default_authentication_method}.
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#enabled PinpointApnsSandboxChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#private_key PinpointApnsSandboxChannel#private_key}.
	PrivateKey *string `json:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#team_id PinpointApnsSandboxChannel#team_id}.
	TeamId *string `json:"teamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#token_key PinpointApnsSandboxChannel#token_key}.
	TokenKey *string `json:"tokenKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_sandbox_channel.html#token_key_id PinpointApnsSandboxChannel#token_key_id}.
	TokenKeyId *string `json:"tokenKeyId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html aws_pinpoint_apns_voip_channel}.
type PinpointApnsVoipChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	DefaultAuthenticationMethodInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenKey() *string
	SetTokenKey(val *string)
	TokenKeyId() *string
	SetTokenKeyId(val *string)
	TokenKeyIdInput() *string
	TokenKeyInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetBundleId()
	ResetCertificate()
	ResetDefaultAuthenticationMethod()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetTeamId()
	ResetTokenKey()
	ResetTokenKeyId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApnsVoipChannel
type jsiiProxy_PinpointApnsVoipChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsVoipChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html aws_pinpoint_apns_voip_channel} Resource.
func NewPinpointApnsVoipChannel(scope constructs.Construct, id *string, config *PinpointApnsVoipChannelConfig) PinpointApnsVoipChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointApnsVoipChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html aws_pinpoint_apns_voip_channel} Resource.
func NewPinpointApnsVoipChannel_Override(p PinpointApnsVoipChannel, scope constructs.Construct, id *string, config *PinpointApnsVoipChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointApnsVoipChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsVoipChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointApnsVoipChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointApnsVoipChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#application_id PinpointApnsVoipChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#bundle_id PinpointApnsVoipChannel#bundle_id}.
	BundleId *string `json:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#certificate PinpointApnsVoipChannel#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#default_authentication_method PinpointApnsVoipChannel#default_authentication_method}.
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#enabled PinpointApnsVoipChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#private_key PinpointApnsVoipChannel#private_key}.
	PrivateKey *string `json:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#team_id PinpointApnsVoipChannel#team_id}.
	TeamId *string `json:"teamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#token_key PinpointApnsVoipChannel#token_key}.
	TokenKey *string `json:"tokenKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_channel.html#token_key_id PinpointApnsVoipChannel#token_key_id}.
	TokenKeyId *string `json:"tokenKeyId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html aws_pinpoint_apns_voip_sandbox_channel}.
type PinpointApnsVoipSandboxChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	DefaultAuthenticationMethodInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenKey() *string
	SetTokenKey(val *string)
	TokenKeyId() *string
	SetTokenKeyId(val *string)
	TokenKeyIdInput() *string
	TokenKeyInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetBundleId()
	ResetCertificate()
	ResetDefaultAuthenticationMethod()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetTeamId()
	ResetTokenKey()
	ResetTokenKeyId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApnsVoipSandboxChannel
type jsiiProxy_PinpointApnsVoipSandboxChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DefaultAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) TokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html aws_pinpoint_apns_voip_sandbox_channel} Resource.
func NewPinpointApnsVoipSandboxChannel(scope constructs.Construct, id *string, config *PinpointApnsVoipSandboxChannelConfig) PinpointApnsVoipSandboxChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointApnsVoipSandboxChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipSandboxChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html aws_pinpoint_apns_voip_sandbox_channel} Resource.
func NewPinpointApnsVoipSandboxChannel_Override(p PinpointApnsVoipSandboxChannel, scope constructs.Construct, id *string, config *PinpointApnsVoipSandboxChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipSandboxChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_PinpointApnsVoipSandboxChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointApnsVoipSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApnsVoipSandboxChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointApnsVoipSandboxChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetBundleId() {
	_jsii_.InvokeVoid(
		p,
		"resetBundleId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetCertificate() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetDefaultAuthenticationMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAuthenticationMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTeamId() {
	_jsii_.InvokeVoid(
		p,
		"resetTeamId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTokenKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ResetTokenKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetTokenKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointApnsVoipSandboxChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointApnsVoipSandboxChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#application_id PinpointApnsVoipSandboxChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#bundle_id PinpointApnsVoipSandboxChannel#bundle_id}.
	BundleId *string `json:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#certificate PinpointApnsVoipSandboxChannel#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#default_authentication_method PinpointApnsVoipSandboxChannel#default_authentication_method}.
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#enabled PinpointApnsVoipSandboxChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#private_key PinpointApnsVoipSandboxChannel#private_key}.
	PrivateKey *string `json:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#team_id PinpointApnsVoipSandboxChannel#team_id}.
	TeamId *string `json:"teamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#token_key PinpointApnsVoipSandboxChannel#token_key}.
	TokenKey *string `json:"tokenKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_apns_voip_sandbox_channel.html#token_key_id PinpointApnsVoipSandboxChannel#token_key_id}.
	TokenKeyId *string `json:"tokenKeyId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html aws_pinpoint_app}.
type PinpointApp interface {
	cdktf.TerraformResource
	ApplicationId() *string
	Arn() *string
	CampaignHook() PinpointAppCampaignHookOutputReference
	CampaignHookInput() *PinpointAppCampaignHook
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
	Limits() PinpointAppLimitsOutputReference
	LimitsInput() *PinpointAppLimits
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	QuietTime() PinpointAppQuietTimeOutputReference
	QuietTimeInput() *PinpointAppQuietTime
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
	PutCampaignHook(value *PinpointAppCampaignHook)
	PutLimits(value *PinpointAppLimits)
	PutQuietTime(value *PinpointAppQuietTime)
	ResetCampaignHook()
	ResetLimits()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetQuietTime()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointApp
type jsiiProxy_PinpointApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointApp) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) CampaignHook() PinpointAppCampaignHookOutputReference {
	var returns PinpointAppCampaignHookOutputReference
	_jsii_.Get(
		j,
		"campaignHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) CampaignHookInput() *PinpointAppCampaignHook {
	var returns *PinpointAppCampaignHook
	_jsii_.Get(
		j,
		"campaignHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Limits() PinpointAppLimitsOutputReference {
	var returns PinpointAppLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) LimitsInput() *PinpointAppLimits {
	var returns *PinpointAppLimits
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) QuietTime() PinpointAppQuietTimeOutputReference {
	var returns PinpointAppQuietTimeOutputReference
	_jsii_.Get(
		j,
		"quietTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) QuietTimeInput() *PinpointAppQuietTime {
	var returns *PinpointAppQuietTime
	_jsii_.Get(
		j,
		"quietTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html aws_pinpoint_app} Resource.
func NewPinpointApp(scope constructs.Construct, id *string, config *PinpointAppConfig) PinpointApp {
	_init_.Initialize()

	j := jsiiProxy_PinpointApp{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html aws_pinpoint_app} Resource.
func NewPinpointApp_Override(p PinpointApp, scope constructs.Construct, id *string, config *PinpointAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointApp",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointApp) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PinpointApp) SetTagsAll(val interface{}) {
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
func PinpointApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointApp) PutCampaignHook(value *PinpointAppCampaignHook) {
	_jsii_.InvokeVoid(
		p,
		"putCampaignHook",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointApp) PutLimits(value *PinpointAppLimits) {
	_jsii_.InvokeVoid(
		p,
		"putLimits",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointApp) PutQuietTime(value *PinpointAppQuietTime) {
	_jsii_.InvokeVoid(
		p,
		"putQuietTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointApp) ResetCampaignHook() {
	_jsii_.InvokeVoid(
		p,
		"resetCampaignHook",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetLimits() {
	_jsii_.InvokeVoid(
		p,
		"resetLimits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		p,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetQuietTime() {
	_jsii_.InvokeVoid(
		p,
		"resetQuietTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) ResetTagsAll() {
	_jsii_.InvokeVoid(
		p,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointAppCampaignHook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#lambda_function_name PinpointApp#lambda_function_name}.
	LambdaFunctionName *string `json:"lambdaFunctionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#mode PinpointApp#mode}.
	Mode *string `json:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#web_url PinpointApp#web_url}.
	WebUrl *string `json:"webUrl"`
}

type PinpointAppCampaignHookOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaFunctionName() *string
	SetLambdaFunctionName(val *string)
	LambdaFunctionNameInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WebUrl() *string
	SetWebUrl(val *string)
	WebUrlInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetLambdaFunctionName()
	ResetMode()
	ResetWebUrl()
}

// The jsii proxy struct for PinpointAppCampaignHookOutputReference
type jsiiProxy_PinpointAppCampaignHookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) LambdaFunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) LambdaFunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) WebUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrlInput",
		&returns,
	)
	return returns
}

func NewPinpointAppCampaignHookOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) PinpointAppCampaignHookOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PinpointAppCampaignHookOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppCampaignHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewPinpointAppCampaignHookOutputReference_Override(p PinpointAppCampaignHookOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppCampaignHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		p,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetLambdaFunctionName(val *string) {
	_jsii_.Set(
		j,
		"lambdaFunctionName",
		val,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PinpointAppCampaignHookOutputReference) SetWebUrl(val *string) {
	_jsii_.Set(
		j,
		"webUrl",
		val,
	)
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) ResetLambdaFunctionName() {
	_jsii_.InvokeVoid(
		p,
		"resetLambdaFunctionName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		p,
		"resetMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppCampaignHookOutputReference) ResetWebUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetWebUrl",
		nil, // no parameters
	)
}

type PinpointAppConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// campaign_hook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#campaign_hook PinpointApp#campaign_hook}
	CampaignHook *PinpointAppCampaignHook `json:"campaignHook"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#limits PinpointApp#limits}
	Limits *PinpointAppLimits `json:"limits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#name PinpointApp#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#name_prefix PinpointApp#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// quiet_time block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#quiet_time PinpointApp#quiet_time}
	QuietTime *PinpointAppQuietTime `json:"quietTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#tags PinpointApp#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#tags_all PinpointApp#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type PinpointAppLimits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#daily PinpointApp#daily}.
	Daily *float64 `json:"daily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#maximum_duration PinpointApp#maximum_duration}.
	MaximumDuration *float64 `json:"maximumDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#messages_per_second PinpointApp#messages_per_second}.
	MessagesPerSecond *float64 `json:"messagesPerSecond"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#total PinpointApp#total}.
	Total *float64 `json:"total"`
}

type PinpointAppLimitsOutputReference interface {
	cdktf.ComplexObject
	Daily() *float64
	SetDaily(val *float64)
	DailyInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaximumDuration() *float64
	SetMaximumDuration(val *float64)
	MaximumDurationInput() *float64
	MessagesPerSecond() *float64
	SetMessagesPerSecond(val *float64)
	MessagesPerSecondInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Total() *float64
	SetTotal(val *float64)
	TotalInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDaily()
	ResetMaximumDuration()
	ResetMessagesPerSecond()
	ResetTotal()
}

// The jsii proxy struct for PinpointAppLimitsOutputReference
type jsiiProxy_PinpointAppLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) Daily() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) DailyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MaximumDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MaximumDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MessagesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) Total() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"total",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TotalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInput",
		&returns,
	)
	return returns
}

func NewPinpointAppLimitsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) PinpointAppLimitsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PinpointAppLimitsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewPinpointAppLimitsOutputReference_Override(p PinpointAppLimitsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		p,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetDaily(val *float64) {
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetMaximumDuration(val *float64) {
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetMessagesPerSecond(val *float64) {
	_jsii_.Set(
		j,
		"messagesPerSecond",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) SetTotal(val *float64) {
	_jsii_.Set(
		j,
		"total",
		val,
	)
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		p,
		"resetDaily",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetMessagesPerSecond() {
	_jsii_.InvokeVoid(
		p,
		"resetMessagesPerSecond",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetTotal() {
	_jsii_.InvokeVoid(
		p,
		"resetTotal",
		nil, // no parameters
	)
}

type PinpointAppQuietTime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#end PinpointApp#end}.
	End *string `json:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_app.html#start PinpointApp#start}.
	Start *string `json:"start"`
}

type PinpointAppQuietTimeOutputReference interface {
	cdktf.ComplexObject
	End() *string
	SetEnd(val *string)
	EndInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Start() *string
	SetStart(val *string)
	StartInput() *string
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
	ResetEnd()
	ResetStart()
}

// The jsii proxy struct for PinpointAppQuietTimeOutputReference
type jsiiProxy_PinpointAppQuietTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewPinpointAppQuietTimeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) PinpointAppQuietTimeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PinpointAppQuietTimeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppQuietTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewPinpointAppQuietTimeOutputReference_Override(p PinpointAppQuietTimeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointAppQuietTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		p,
	)
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) SetEnd(val *string) {
	_jsii_.Set(
		j,
		"end",
		val,
	)
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) SetStart(val *string) {
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointAppQuietTimeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		p,
		"resetEnd",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppQuietTimeOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		p,
		"resetStart",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html aws_pinpoint_baidu_channel}.
type PinpointBaiduChannel interface {
	cdktf.TerraformResource
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
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
	ResetEnabled()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointBaiduChannel
type jsiiProxy_PinpointBaiduChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointBaiduChannel) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointBaiduChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html aws_pinpoint_baidu_channel} Resource.
func NewPinpointBaiduChannel(scope constructs.Construct, id *string, config *PinpointBaiduChannelConfig) PinpointBaiduChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointBaiduChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointBaiduChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html aws_pinpoint_baidu_channel} Resource.
func NewPinpointBaiduChannel_Override(p PinpointBaiduChannel, scope constructs.Construct, id *string, config *PinpointBaiduChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointBaiduChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointBaiduChannel) SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointBaiduChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointBaiduChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointBaiduChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointBaiduChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointBaiduChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointBaiduChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointBaiduChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointBaiduChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointBaiduChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html#api_key PinpointBaiduChannel#api_key}.
	ApiKey *string `json:"apiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html#application_id PinpointBaiduChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html#secret_key PinpointBaiduChannel#secret_key}.
	SecretKey *string `json:"secretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_baidu_channel.html#enabled PinpointBaiduChannel#enabled}.
	Enabled interface{} `json:"enabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html aws_pinpoint_email_channel}.
type PinpointEmailChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConfigurationSet() *string
	SetConfigurationSet(val *string)
	ConfigurationSetInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	FromAddress() *string
	SetFromAddress(val *string)
	FromAddressInput() *string
	Id() *string
	Identity() *string
	SetIdentity(val *string)
	IdentityInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MessagesPerSecond() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetConfigurationSet()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointEmailChannel
type jsiiProxy_PinpointEmailChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointEmailChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) ConfigurationSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) FromAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) FromAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) IdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) MessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html aws_pinpoint_email_channel} Resource.
func NewPinpointEmailChannel(scope constructs.Construct, id *string, config *PinpointEmailChannelConfig) PinpointEmailChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointEmailChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointEmailChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html aws_pinpoint_email_channel} Resource.
func NewPinpointEmailChannel_Override(p PinpointEmailChannel, scope constructs.Construct, id *string, config *PinpointEmailChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointEmailChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetConfigurationSet(val *string) {
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetFromAddress(val *string) {
	_jsii_.Set(
		j,
		"fromAddress",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetIdentity(val *string) {
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailChannel) SetRoleArn(val *string) {
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
func PinpointEmailChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointEmailChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointEmailChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointEmailChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointEmailChannel) ResetConfigurationSet() {
	_jsii_.InvokeVoid(
		p,
		"resetConfigurationSet",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailChannel) ResetRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointEmailChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointEmailChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointEmailChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#application_id PinpointEmailChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#from_address PinpointEmailChannel#from_address}.
	FromAddress *string `json:"fromAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#identity PinpointEmailChannel#identity}.
	Identity *string `json:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#configuration_set PinpointEmailChannel#configuration_set}.
	ConfigurationSet *string `json:"configurationSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#enabled PinpointEmailChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_email_channel.html#role_arn PinpointEmailChannel#role_arn}.
	RoleArn *string `json:"roleArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html aws_pinpoint_event_stream}.
type PinpointEventStream interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DestinationStreamArn() *string
	SetDestinationStreamArn(val *string)
	DestinationStreamArnInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointEventStream
type jsiiProxy_PinpointEventStream struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointEventStream) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) DestinationStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) DestinationStreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationStreamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEventStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html aws_pinpoint_event_stream} Resource.
func NewPinpointEventStream(scope constructs.Construct, id *string, config *PinpointEventStreamConfig) PinpointEventStream {
	_init_.Initialize()

	j := jsiiProxy_PinpointEventStream{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointEventStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html aws_pinpoint_event_stream} Resource.
func NewPinpointEventStream_Override(p PinpointEventStream, scope constructs.Construct, id *string, config *PinpointEventStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointEventStream",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetDestinationStreamArn(val *string) {
	_jsii_.Set(
		j,
		"destinationStreamArn",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointEventStream) SetRoleArn(val *string) {
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
func PinpointEventStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointEventStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointEventStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointEventStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointEventStream) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointEventStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEventStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointEventStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointEventStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointEventStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointEventStreamConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html#application_id PinpointEventStream#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html#destination_stream_arn PinpointEventStream#destination_stream_arn}.
	DestinationStreamArn *string `json:"destinationStreamArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_event_stream.html#role_arn PinpointEventStream#role_arn}.
	RoleArn *string `json:"roleArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html aws_pinpoint_gcm_channel}.
type PinpointGcmChannel interface {
	cdktf.TerraformResource
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	ResetEnabled()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointGcmChannel
type jsiiProxy_PinpointGcmChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointGcmChannel) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointGcmChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html aws_pinpoint_gcm_channel} Resource.
func NewPinpointGcmChannel(scope constructs.Construct, id *string, config *PinpointGcmChannelConfig) PinpointGcmChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointGcmChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointGcmChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html aws_pinpoint_gcm_channel} Resource.
func NewPinpointGcmChannel_Override(p PinpointGcmChannel, scope constructs.Construct, id *string, config *PinpointGcmChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointGcmChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointGcmChannel) SetProvider(val cdktf.TerraformProvider) {
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
func PinpointGcmChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointGcmChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointGcmChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointGcmChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointGcmChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointGcmChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointGcmChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointGcmChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointGcmChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html#api_key PinpointGcmChannel#api_key}.
	ApiKey *string `json:"apiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html#application_id PinpointGcmChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_gcm_channel.html#enabled PinpointGcmChannel#enabled}.
	Enabled interface{} `json:"enabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html aws_pinpoint_sms_channel}.
type PinpointSmsChannel interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PromotionalMessagesPerSecond() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SenderId() *string
	SetSenderId(val *string)
	SenderIdInput() *string
	ShortCode() *string
	SetShortCode(val *string)
	ShortCodeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TransactionalMessagesPerSecond() *float64
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetSenderId()
	ResetShortCode()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PinpointSmsChannel
type jsiiProxy_PinpointSmsChannel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PinpointSmsChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) PromotionalMessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionalMessagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) SenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) SenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) ShortCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) ShortCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointSmsChannel) TransactionalMessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionalMessagesPerSecond",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html aws_pinpoint_sms_channel} Resource.
func NewPinpointSmsChannel(scope constructs.Construct, id *string, config *PinpointSmsChannelConfig) PinpointSmsChannel {
	_init_.Initialize()

	j := jsiiProxy_PinpointSmsChannel{}

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointSmsChannel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html aws_pinpoint_sms_channel} Resource.
func NewPinpointSmsChannel_Override(p PinpointSmsChannel, scope constructs.Construct, id *string, config *PinpointSmsChannelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Pinpoint.PinpointSmsChannel",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetSenderId(val *string) {
	_jsii_.Set(
		j,
		"senderId",
		val,
	)
}

func (j *jsiiProxy_PinpointSmsChannel) SetShortCode(val *string) {
	_jsii_.Set(
		j,
		"shortCode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PinpointSmsChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Pinpoint.PinpointSmsChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PinpointSmsChannel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Pinpoint.PinpointSmsChannel",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PinpointSmsChannel) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointSmsChannel) ResetSenderId() {
	_jsii_.InvokeVoid(
		p,
		"resetSenderId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointSmsChannel) ResetShortCode() {
	_jsii_.InvokeVoid(
		p,
		"resetShortCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointSmsChannel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PinpointSmsChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PinpointSmsChannel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PinpointSmsChannelConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html#application_id PinpointSmsChannel#application_id}.
	ApplicationId *string `json:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html#enabled PinpointSmsChannel#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html#sender_id PinpointSmsChannel#sender_id}.
	SenderId *string `json:"senderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pinpoint_sms_channel.html#short_code PinpointSmsChannel#short_code}.
	ShortCode *string `json:"shortCode"`
}
