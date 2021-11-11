package transfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/transfer/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/transfer_server.html aws_transfer_server}.
type DataAwsTransferServer interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Domain() *string
	Endpoint() *string
	EndpointType() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentityProviderType() *string
	InvocationRole() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingRole() *string
	Node() constructs.Node
	Protocols() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SecurityPolicyName() *string
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Url() *string
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

// The jsii proxy struct for DataAwsTransferServer
type jsiiProxy_DataAwsTransferServer struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsTransferServer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) IdentityProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) InvocationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) LoggingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) SecurityPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsTransferServer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/transfer_server.html aws_transfer_server} Data Source.
func NewDataAwsTransferServer(scope constructs.Construct, id *string, config *DataAwsTransferServerConfig) DataAwsTransferServer {
	_init_.Initialize()

	j := jsiiProxy_DataAwsTransferServer{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.DataAwsTransferServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/transfer_server.html aws_transfer_server} Data Source.
func NewDataAwsTransferServer_Override(d DataAwsTransferServer, scope constructs.Construct, id *string, config *DataAwsTransferServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.DataAwsTransferServer",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsTransferServer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsTransferServer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsTransferServer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsTransferServer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsTransferServer) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsTransferServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Transfer.DataAwsTransferServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsTransferServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Transfer.DataAwsTransferServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsTransferServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsTransferServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsTransferServer) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsTransferServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsTransferServer) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsTransferServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsTransferServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsTransferServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsTransferServer) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsTransferServer) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsTransferServer) ToString() *string {
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
func (d *jsiiProxy_DataAwsTransferServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsTransferServerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/transfer_server.html#server_id DataAwsTransferServer#server_id}.
	ServerId *string `json:"serverId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html aws_transfer_access}.
type TransferAccess interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HomeDirectory() *string
	SetHomeDirectory(val *string)
	HomeDirectoryInput() *string
	HomeDirectoryMappings() *[]*TransferAccessHomeDirectoryMappings
	SetHomeDirectoryMappings(val *[]*TransferAccessHomeDirectoryMappings)
	HomeDirectoryMappingsInput() *[]*TransferAccessHomeDirectoryMappings
	HomeDirectoryType() *string
	SetHomeDirectoryType(val *string)
	HomeDirectoryTypeInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	PosixProfile() TransferAccessPosixProfileOutputReference
	PosixProfileInput() *TransferAccessPosixProfile
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
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
	PutPosixProfile(value *TransferAccessPosixProfile)
	ResetHomeDirectory()
	ResetHomeDirectoryMappings()
	ResetHomeDirectoryType()
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetPosixProfile()
	ResetRole()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TransferAccess
type jsiiProxy_TransferAccess struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TransferAccess) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectoryMappings() *[]*TransferAccessHomeDirectoryMappings {
	var returns *[]*TransferAccessHomeDirectoryMappings
	_jsii_.Get(
		j,
		"homeDirectoryMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectoryMappingsInput() *[]*TransferAccessHomeDirectoryMappings {
	var returns *[]*TransferAccessHomeDirectoryMappings
	_jsii_.Get(
		j,
		"homeDirectoryMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) HomeDirectoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) PosixProfile() TransferAccessPosixProfileOutputReference {
	var returns TransferAccessPosixProfileOutputReference
	_jsii_.Get(
		j,
		"posixProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) PosixProfileInput() *TransferAccessPosixProfile {
	var returns *TransferAccessPosixProfile
	_jsii_.Get(
		j,
		"posixProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccess) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html aws_transfer_access} Resource.
func NewTransferAccess(scope constructs.Construct, id *string, config *TransferAccessConfig) TransferAccess {
	_init_.Initialize()

	j := jsiiProxy_TransferAccess{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferAccess",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html aws_transfer_access} Resource.
func NewTransferAccess_Override(t TransferAccess, scope constructs.Construct, id *string, config *TransferAccessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferAccess",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransferAccess) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetExternalId(val *string) {
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetHomeDirectory(val *string) {
	_jsii_.Set(
		j,
		"homeDirectory",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetHomeDirectoryMappings(val *[]*TransferAccessHomeDirectoryMappings) {
	_jsii_.Set(
		j,
		"homeDirectoryMappings",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetHomeDirectoryType(val *string) {
	_jsii_.Set(
		j,
		"homeDirectoryType",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_TransferAccess) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TransferAccess_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Transfer.TransferAccess",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransferAccess_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Transfer.TransferAccess",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TransferAccess) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TransferAccess) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransferAccess) PutPosixProfile(value *TransferAccessPosixProfile) {
	_jsii_.InvokeVoid(
		t,
		"putPosixProfile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferAccess) ResetHomeDirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) ResetHomeDirectoryMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectoryMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) ResetHomeDirectoryType() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectoryType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TransferAccess) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) ResetPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) ResetPosixProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetPosixProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) ResetRole() {
	_jsii_.InvokeVoid(
		t,
		"resetRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferAccess) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccess) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TransferAccess) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TransferAccess) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TransferAccessConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#external_id TransferAccess#external_id}.
	ExternalId *string `json:"externalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#server_id TransferAccess#server_id}.
	ServerId *string `json:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#home_directory TransferAccess#home_directory}.
	HomeDirectory *string `json:"homeDirectory"`
	// home_directory_mappings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#home_directory_mappings TransferAccess#home_directory_mappings}
	HomeDirectoryMappings *[]*TransferAccessHomeDirectoryMappings `json:"homeDirectoryMappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#home_directory_type TransferAccess#home_directory_type}.
	HomeDirectoryType *string `json:"homeDirectoryType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#policy TransferAccess#policy}.
	Policy *string `json:"policy"`
	// posix_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#posix_profile TransferAccess#posix_profile}
	PosixProfile *TransferAccessPosixProfile `json:"posixProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#role TransferAccess#role}.
	Role *string `json:"role"`
}

type TransferAccessHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#entry TransferAccess#entry}.
	Entry *string `json:"entry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#target TransferAccess#target}.
	Target *string `json:"target"`
}

type TransferAccessPosixProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#gid TransferAccess#gid}.
	Gid *float64 `json:"gid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#uid TransferAccess#uid}.
	Uid *float64 `json:"uid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access.html#secondary_gids TransferAccess#secondary_gids}.
	SecondaryGids *[]*float64 `json:"secondaryGids"`
}

type TransferAccessPosixProfileOutputReference interface {
	cdktf.ComplexObject
	Gid() *float64
	SetGid(val *float64)
	GidInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecondaryGids() *[]*float64
	SetSecondaryGids(val *[]*float64)
	SecondaryGidsInput() *[]*float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Uid() *float64
	SetUid(val *float64)
	UidInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecondaryGids()
}

// The jsii proxy struct for TransferAccessPosixProfileOutputReference
type jsiiProxy_TransferAccessPosixProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) Gid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) GidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SecondaryGids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SecondaryGidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) Uid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) UidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func NewTransferAccessPosixProfileOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) TransferAccessPosixProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_TransferAccessPosixProfileOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferAccessPosixProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewTransferAccessPosixProfileOutputReference_Override(t TransferAccessPosixProfileOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferAccessPosixProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		t,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetGid(val *float64) {
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetSecondaryGids(val *[]*float64) {
	_jsii_.Set(
		j,
		"secondaryGids",
		val,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferAccessPosixProfileOutputReference) SetUid(val *float64) {
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferAccessPosixProfileOutputReference) ResetSecondaryGids() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryGids",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html aws_transfer_server}.
type TransferServer interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	Endpoint() *string
	EndpointDetails() TransferServerEndpointDetailsOutputReference
	EndpointDetailsInput() *TransferServerEndpointDetails
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	HostKey() *string
	SetHostKey(val *string)
	HostKeyFingerprint() *string
	HostKeyInput() *string
	Id() *string
	IdentityProviderType() *string
	SetIdentityProviderType(val *string)
	IdentityProviderTypeInput() *string
	InvocationRole() *string
	SetInvocationRole(val *string)
	InvocationRoleInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingRole() *string
	SetLoggingRole(val *string)
	LoggingRoleInput() *string
	Node() constructs.Node
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SecurityPolicyName() *string
	SetSecurityPolicyName(val *string)
	SecurityPolicyNameInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutEndpointDetails(value *TransferServerEndpointDetails)
	ResetCertificate()
	ResetDirectoryId()
	ResetDomain()
	ResetEndpointDetails()
	ResetEndpointType()
	ResetForceDestroy()
	ResetHostKey()
	ResetIdentityProviderType()
	ResetInvocationRole()
	ResetLoggingRole()
	ResetOverrideLogicalId()
	ResetProtocols()
	ResetSecurityPolicyName()
	ResetTags()
	ResetTagsAll()
	ResetUrl()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TransferServer
type jsiiProxy_TransferServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TransferServer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) EndpointDetails() TransferServerEndpointDetailsOutputReference {
	var returns TransferServerEndpointDetailsOutputReference
	_jsii_.Get(
		j,
		"endpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) EndpointDetailsInput() *TransferServerEndpointDetails {
	var returns *TransferServerEndpointDetails
	_jsii_.Get(
		j,
		"endpointDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) HostKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) IdentityProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) IdentityProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) InvocationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) InvocationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) LoggingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) LoggingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) SecurityPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) SecurityPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServer) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html aws_transfer_server} Resource.
func NewTransferServer(scope constructs.Construct, id *string, config *TransferServerConfig) TransferServer {
	_init_.Initialize()

	j := jsiiProxy_TransferServer{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html aws_transfer_server} Resource.
func NewTransferServer_Override(t TransferServer, scope constructs.Construct, id *string, config *TransferServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferServer",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransferServer) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetForceDestroy(val interface{}) {
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetHostKey(val *string) {
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetIdentityProviderType(val *string) {
	_jsii_.Set(
		j,
		"identityProviderType",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetInvocationRole(val *string) {
	_jsii_.Set(
		j,
		"invocationRole",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetLoggingRole(val *string) {
	_jsii_.Set(
		j,
		"loggingRole",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetSecurityPolicyName(val *string) {
	_jsii_.Set(
		j,
		"securityPolicyName",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TransferServer) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TransferServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Transfer.TransferServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransferServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Transfer.TransferServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TransferServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TransferServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransferServer) PutEndpointDetails(value *TransferServerEndpointDetails) {
	_jsii_.InvokeVoid(
		t,
		"putEndpointDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferServer) ResetCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetEndpointDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetEndpointType() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetHostKey() {
	_jsii_.InvokeVoid(
		t,
		"resetHostKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetIdentityProviderType() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityProviderType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetInvocationRole() {
	_jsii_.InvokeVoid(
		t,
		"resetInvocationRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetLoggingRole() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingRole",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TransferServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetProtocols() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetSecurityPolicyName() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityPolicyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) ResetUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TransferServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TransferServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TransferServerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#certificate TransferServer#certificate}.
	Certificate *string `json:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#directory_id TransferServer#directory_id}.
	DirectoryId *string `json:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#domain TransferServer#domain}.
	Domain *string `json:"domain"`
	// endpoint_details block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#endpoint_details TransferServer#endpoint_details}
	EndpointDetails *TransferServerEndpointDetails `json:"endpointDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#endpoint_type TransferServer#endpoint_type}.
	EndpointType *string `json:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#force_destroy TransferServer#force_destroy}.
	ForceDestroy interface{} `json:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#host_key TransferServer#host_key}.
	HostKey *string `json:"hostKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#identity_provider_type TransferServer#identity_provider_type}.
	IdentityProviderType *string `json:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#invocation_role TransferServer#invocation_role}.
	InvocationRole *string `json:"invocationRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#logging_role TransferServer#logging_role}.
	LoggingRole *string `json:"loggingRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#protocols TransferServer#protocols}.
	Protocols *[]*string `json:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#security_policy_name TransferServer#security_policy_name}.
	SecurityPolicyName *string `json:"securityPolicyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#tags TransferServer#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#tags_all TransferServer#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#url TransferServer#url}.
	Url *string `json:"url"`
}

type TransferServerEndpointDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#address_allocation_ids TransferServer#address_allocation_ids}.
	AddressAllocationIds *[]*string `json:"addressAllocationIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#security_group_ids TransferServer#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#subnet_ids TransferServer#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#vpc_endpoint_id TransferServer#vpc_endpoint_id}.
	VpcEndpointId *string `json:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server.html#vpc_id TransferServer#vpc_id}.
	VpcId *string `json:"vpcId"`
}

type TransferServerEndpointDetailsOutputReference interface {
	cdktf.ComplexObject
	AddressAllocationIds() *[]*string
	SetAddressAllocationIds(val *[]*string)
	AddressAllocationIdsInput() *[]*string
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
	VpcEndpointId() *string
	SetVpcEndpointId(val *string)
	VpcEndpointIdInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAddressAllocationIds()
	ResetSecurityGroupIds()
	ResetSubnetIds()
	ResetVpcEndpointId()
	ResetVpcId()
}

// The jsii proxy struct for TransferServerEndpointDetailsOutputReference
type jsiiProxy_TransferServerEndpointDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) AddressAllocationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressAllocationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) AddressAllocationIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressAllocationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func NewTransferServerEndpointDetailsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) TransferServerEndpointDetailsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_TransferServerEndpointDetailsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferServerEndpointDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewTransferServerEndpointDetailsOutputReference_Override(t TransferServerEndpointDetailsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferServerEndpointDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		t,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetAddressAllocationIds(val *[]*string) {
	_jsii_.Set(
		j,
		"addressAllocationIds",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetVpcEndpointId(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_TransferServerEndpointDetailsOutputReference) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) ResetAddressAllocationIds() {
	_jsii_.InvokeVoid(
		t,
		"resetAddressAllocationIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerEndpointDetailsOutputReference) ResetVpcId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcId",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html aws_transfer_ssh_key}.
type TransferSshKey interface {
	cdktf.TerraformResource
	Body() *string
	SetBody(val *string)
	BodyInput() *string
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
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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

// The jsii proxy struct for TransferSshKey
type jsiiProxy_TransferSshKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TransferSshKey) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferSshKey) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html aws_transfer_ssh_key} Resource.
func NewTransferSshKey(scope constructs.Construct, id *string, config *TransferSshKeyConfig) TransferSshKey {
	_init_.Initialize()

	j := jsiiProxy_TransferSshKey{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferSshKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html aws_transfer_ssh_key} Resource.
func NewTransferSshKey_Override(t TransferSshKey, scope constructs.Construct, id *string, config *TransferSshKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferSshKey",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransferSshKey) SetBody(val *string) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

func (j *jsiiProxy_TransferSshKey) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TransferSshKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Transfer.TransferSshKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransferSshKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Transfer.TransferSshKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TransferSshKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TransferSshKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferSshKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferSshKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TransferSshKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TransferSshKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TransferSshKeyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html#body TransferSshKey#body}.
	Body *string `json:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html#server_id TransferSshKey#server_id}.
	ServerId *string `json:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html#user_name TransferSshKey#user_name}.
	UserName *string `json:"userName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html aws_transfer_user}.
type TransferUser interface {
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
	HomeDirectory() *string
	SetHomeDirectory(val *string)
	HomeDirectoryInput() *string
	HomeDirectoryMappings() *[]*TransferUserHomeDirectoryMappings
	SetHomeDirectoryMappings(val *[]*TransferUserHomeDirectoryMappings)
	HomeDirectoryMappingsInput() *[]*TransferUserHomeDirectoryMappings
	HomeDirectoryType() *string
	SetHomeDirectoryType(val *string)
	HomeDirectoryTypeInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	PosixProfile() TransferUserPosixProfileOutputReference
	PosixProfileInput() *TransferUserPosixProfile
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutPosixProfile(value *TransferUserPosixProfile)
	ResetHomeDirectory()
	ResetHomeDirectoryMappings()
	ResetHomeDirectoryType()
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetPosixProfile()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TransferUser
type jsiiProxy_TransferUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TransferUser) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectoryMappings() *[]*TransferUserHomeDirectoryMappings {
	var returns *[]*TransferUserHomeDirectoryMappings
	_jsii_.Get(
		j,
		"homeDirectoryMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectoryMappingsInput() *[]*TransferUserHomeDirectoryMappings {
	var returns *[]*TransferUserHomeDirectoryMappings
	_jsii_.Get(
		j,
		"homeDirectoryMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) HomeDirectoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) PosixProfile() TransferUserPosixProfileOutputReference {
	var returns TransferUserPosixProfileOutputReference
	_jsii_.Get(
		j,
		"posixProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) PosixProfileInput() *TransferUserPosixProfile {
	var returns *TransferUserPosixProfile
	_jsii_.Get(
		j,
		"posixProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUser) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html aws_transfer_user} Resource.
func NewTransferUser(scope constructs.Construct, id *string, config *TransferUserConfig) TransferUser {
	_init_.Initialize()

	j := jsiiProxy_TransferUser{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html aws_transfer_user} Resource.
func NewTransferUser_Override(t TransferUser, scope constructs.Construct, id *string, config *TransferUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferUser",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransferUser) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetHomeDirectory(val *string) {
	_jsii_.Set(
		j,
		"homeDirectory",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetHomeDirectoryMappings(val *[]*TransferUserHomeDirectoryMappings) {
	_jsii_.Set(
		j,
		"homeDirectoryMappings",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetHomeDirectoryType(val *string) {
	_jsii_.Set(
		j,
		"homeDirectoryType",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TransferUser) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TransferUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Transfer.TransferUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransferUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Transfer.TransferUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TransferUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TransferUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransferUser) PutPosixProfile(value *TransferUserPosixProfile) {
	_jsii_.InvokeVoid(
		t,
		"putPosixProfile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferUser) ResetHomeDirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetHomeDirectoryMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectoryMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetHomeDirectoryType() {
	_jsii_.InvokeVoid(
		t,
		"resetHomeDirectoryType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TransferUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetPosixProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetPosixProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TransferUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TransferUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TransferUserConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#role TransferUser#role}.
	Role *string `json:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#server_id TransferUser#server_id}.
	ServerId *string `json:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#user_name TransferUser#user_name}.
	UserName *string `json:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#home_directory TransferUser#home_directory}.
	HomeDirectory *string `json:"homeDirectory"`
	// home_directory_mappings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#home_directory_mappings TransferUser#home_directory_mappings}
	HomeDirectoryMappings *[]*TransferUserHomeDirectoryMappings `json:"homeDirectoryMappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#home_directory_type TransferUser#home_directory_type}.
	HomeDirectoryType *string `json:"homeDirectoryType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#policy TransferUser#policy}.
	Policy *string `json:"policy"`
	// posix_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#posix_profile TransferUser#posix_profile}
	PosixProfile *TransferUserPosixProfile `json:"posixProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#tags TransferUser#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#tags_all TransferUser#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type TransferUserHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#entry TransferUser#entry}.
	Entry *string `json:"entry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#target TransferUser#target}.
	Target *string `json:"target"`
}

type TransferUserPosixProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#gid TransferUser#gid}.
	Gid *float64 `json:"gid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#uid TransferUser#uid}.
	Uid *float64 `json:"uid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_user.html#secondary_gids TransferUser#secondary_gids}.
	SecondaryGids *[]*float64 `json:"secondaryGids"`
}

type TransferUserPosixProfileOutputReference interface {
	cdktf.ComplexObject
	Gid() *float64
	SetGid(val *float64)
	GidInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecondaryGids() *[]*float64
	SetSecondaryGids(val *[]*float64)
	SecondaryGidsInput() *[]*float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Uid() *float64
	SetUid(val *float64)
	UidInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecondaryGids()
}

// The jsii proxy struct for TransferUserPosixProfileOutputReference
type jsiiProxy_TransferUserPosixProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) Gid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) GidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SecondaryGids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SecondaryGidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) Uid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) UidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func NewTransferUserPosixProfileOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) TransferUserPosixProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_TransferUserPosixProfileOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferUserPosixProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewTransferUserPosixProfileOutputReference_Override(t TransferUserPosixProfileOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Transfer.TransferUserPosixProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		t,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetGid(val *float64) {
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetSecondaryGids(val *[]*float64) {
	_jsii_.Set(
		j,
		"secondaryGids",
		val,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferUserPosixProfileOutputReference) SetUid(val *float64) {
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TransferUserPosixProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferUserPosixProfileOutputReference) ResetSecondaryGids() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryGids",
		nil, // no parameters
	)
}
