package worklink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/worklink/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html aws_worklink_fleet}.
type WorklinkFleet interface {
	cdktf.TerraformResource
	Arn() *string
	AuditStreamArn() *string
	SetAuditStreamArn(val *string)
	AuditStreamArnInput() *string
	CdktfStack() cdktf.TerraformStack
	CompanyCode() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeviceCaCertificate() *string
	SetDeviceCaCertificate(val *string)
	DeviceCaCertificateInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentityProvider() WorklinkFleetIdentityProviderOutputReference
	IdentityProviderInput() *WorklinkFleetIdentityProvider
	LastUpdatedTime() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() WorklinkFleetNetworkOutputReference
	NetworkInput() *WorklinkFleetNetwork
	Node() constructs.Node
	OptimizeForEndUserLocation() interface{}
	SetOptimizeForEndUserLocation(val interface{})
	OptimizeForEndUserLocationInput() interface{}
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
	PutIdentityProvider(value *WorklinkFleetIdentityProvider)
	PutNetwork(value *WorklinkFleetNetwork)
	ResetAuditStreamArn()
	ResetDeviceCaCertificate()
	ResetDisplayName()
	ResetIdentityProvider()
	ResetNetwork()
	ResetOptimizeForEndUserLocation()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for WorklinkFleet
type jsiiProxy_WorklinkFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorklinkFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) AuditStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) AuditStreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditStreamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) CompanyCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) DeviceCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) DeviceCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) IdentityProvider() WorklinkFleetIdentityProviderOutputReference {
	var returns WorklinkFleetIdentityProviderOutputReference
	_jsii_.Get(
		j,
		"identityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) IdentityProviderInput() *WorklinkFleetIdentityProvider {
	var returns *WorklinkFleetIdentityProvider
	_jsii_.Get(
		j,
		"identityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Network() WorklinkFleetNetworkOutputReference {
	var returns WorklinkFleetNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) NetworkInput() *WorklinkFleetNetwork {
	var returns *WorklinkFleetNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) OptimizeForEndUserLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeForEndUserLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) OptimizeForEndUserLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeForEndUserLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html aws_worklink_fleet} Resource.
func NewWorklinkFleet(scope constructs.Construct, id *string, config *WorklinkFleetConfig) WorklinkFleet {
	_init_.Initialize()

	j := jsiiProxy_WorklinkFleet{}

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html aws_worklink_fleet} Resource.
func NewWorklinkFleet_Override(w WorklinkFleet, scope constructs.Construct, id *string, config *WorklinkFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleet",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetAuditStreamArn(val *string) {
	_jsii_.Set(
		j,
		"auditStreamArn",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetDeviceCaCertificate(val *string) {
	_jsii_.Set(
		j,
		"deviceCaCertificate",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetOptimizeForEndUserLocation(val interface{}) {
	_jsii_.Set(
		j,
		"optimizeForEndUserLocation",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleet) SetProvider(val cdktf.TerraformProvider) {
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
func WorklinkFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.WorkLink.WorklinkFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorklinkFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.WorkLink.WorklinkFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (w *jsiiProxy_WorklinkFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorklinkFleet) PutIdentityProvider(value *WorklinkFleetIdentityProvider) {
	_jsii_.InvokeVoid(
		w,
		"putIdentityProvider",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorklinkFleet) PutNetwork(value *WorklinkFleetNetwork) {
	_jsii_.InvokeVoid(
		w,
		"putNetwork",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetAuditStreamArn() {
	_jsii_.InvokeVoid(
		w,
		"resetAuditStreamArn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetDeviceCaCertificate() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceCaCertificate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetDisplayName() {
	_jsii_.InvokeVoid(
		w,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetIdentityProvider() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentityProvider",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetNetwork() {
	_jsii_.InvokeVoid(
		w,
		"resetNetwork",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) ResetOptimizeForEndUserLocation() {
	_jsii_.InvokeVoid(
		w,
		"resetOptimizeForEndUserLocation",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (w *jsiiProxy_WorklinkFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorklinkFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (w *jsiiProxy_WorklinkFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WorklinkFleetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#name WorklinkFleet#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#audit_stream_arn WorklinkFleet#audit_stream_arn}.
	AuditStreamArn *string `json:"auditStreamArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#device_ca_certificate WorklinkFleet#device_ca_certificate}.
	DeviceCaCertificate *string `json:"deviceCaCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#display_name WorklinkFleet#display_name}.
	DisplayName *string `json:"displayName"`
	// identity_provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#identity_provider WorklinkFleet#identity_provider}
	IdentityProvider *WorklinkFleetIdentityProvider `json:"identityProvider"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#network WorklinkFleet#network}
	Network *WorklinkFleetNetwork `json:"network"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#optimize_for_end_user_location WorklinkFleet#optimize_for_end_user_location}.
	OptimizeForEndUserLocation interface{} `json:"optimizeForEndUserLocation"`
}

type WorklinkFleetIdentityProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#saml_metadata WorklinkFleet#saml_metadata}.
	SamlMetadata *string `json:"samlMetadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#type WorklinkFleet#type}.
	Type *string `json:"type"`
}

type WorklinkFleetIdentityProviderOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SamlMetadata() *string
	SetSamlMetadata(val *string)
	SamlMetadataInput() *string
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
}

// The jsii proxy struct for WorklinkFleetIdentityProviderOutputReference
type jsiiProxy_WorklinkFleetIdentityProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SamlMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SamlMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewWorklinkFleetIdentityProviderOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) WorklinkFleetIdentityProviderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorklinkFleetIdentityProviderOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleetIdentityProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorklinkFleetIdentityProviderOutputReference_Override(w WorklinkFleetIdentityProviderOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleetIdentityProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SetSamlMetadata(val *string) {
	_jsii_.Set(
		j,
		"samlMetadata",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetIdentityProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type WorklinkFleetNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#security_group_ids WorklinkFleet#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#subnet_ids WorklinkFleet#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet.html#vpc_id WorklinkFleet#vpc_id}.
	VpcId *string `json:"vpcId"`
}

type WorklinkFleetNetworkOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for WorklinkFleetNetworkOutputReference
type jsiiProxy_WorklinkFleetNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func NewWorklinkFleetNetworkOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) WorklinkFleetNetworkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorklinkFleetNetworkOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleetNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorklinkFleetNetworkOutputReference_Override(w WorklinkFleetNetworkOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkFleetNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorklinkFleetNetworkOutputReference) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkFleetNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html aws_worklink_website_certificate_authority_association}.
type WorklinkWebsiteCertificateAuthorityAssociation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FleetArn() *string
	SetFleetArn(val *string)
	FleetArnInput() *string
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
	WebsiteCaId() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDisplayName()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for WorklinkWebsiteCertificateAuthorityAssociation
type jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) FleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) FleetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) WebsiteCaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteCaId",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html aws_worklink_website_certificate_authority_association} Resource.
func NewWorklinkWebsiteCertificateAuthorityAssociation(scope constructs.Construct, id *string, config *WorklinkWebsiteCertificateAuthorityAssociationConfig) WorklinkWebsiteCertificateAuthorityAssociation {
	_init_.Initialize()

	j := jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation{}

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkWebsiteCertificateAuthorityAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html aws_worklink_website_certificate_authority_association} Resource.
func NewWorklinkWebsiteCertificateAuthorityAssociation_Override(w WorklinkWebsiteCertificateAuthorityAssociation, scope constructs.Construct, id *string, config *WorklinkWebsiteCertificateAuthorityAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.WorkLink.WorklinkWebsiteCertificateAuthorityAssociation",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetFleetArn(val *string) {
	_jsii_.Set(
		j,
		"fleetArn",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SetProvider(val cdktf.TerraformProvider) {
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
func WorklinkWebsiteCertificateAuthorityAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.WorkLink.WorklinkWebsiteCertificateAuthorityAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorklinkWebsiteCertificateAuthorityAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.WorkLink.WorklinkWebsiteCertificateAuthorityAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ResetDisplayName() {
	_jsii_.InvokeVoid(
		w,
		"resetDisplayName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (w *jsiiProxy_WorklinkWebsiteCertificateAuthorityAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WorklinkWebsiteCertificateAuthorityAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html#certificate WorklinkWebsiteCertificateAuthorityAssociation#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html#fleet_arn WorklinkWebsiteCertificateAuthorityAssociation#fleet_arn}.
	FleetArn *string `json:"fleetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_website_certificate_authority_association.html#display_name WorklinkWebsiteCertificateAuthorityAssociation#display_name}.
	DisplayName *string `json:"displayName"`
}
