package iot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/iot/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/iot_endpoint.html aws_iot_endpoint}.
type DataAwsIotEndpoint interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EndpointAddress() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
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
	ResetEndpointType()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsIotEndpoint
type jsiiProxy_DataAwsIotEndpoint struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsIotEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) EndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIotEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/iot_endpoint.html aws_iot_endpoint} Data Source.
func NewDataAwsIotEndpoint(scope constructs.Construct, id *string, config *DataAwsIotEndpointConfig) DataAwsIotEndpoint {
	_init_.Initialize()

	j := jsiiProxy_DataAwsIotEndpoint{}

	_jsii_.Create(
		"hashicorp_aws.IoT.DataAwsIotEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/iot_endpoint.html aws_iot_endpoint} Data Source.
func NewDataAwsIotEndpoint_Override(d DataAwsIotEndpoint, scope constructs.Construct, id *string, config *DataAwsIotEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.DataAwsIotEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsIotEndpoint) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsIotEndpoint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIotEndpoint) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_DataAwsIotEndpoint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsIotEndpoint) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsIotEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.DataAwsIotEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsIotEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.DataAwsIotEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsIotEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsIotEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIotEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsIotEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsIotEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsIotEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIotEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsIotEndpoint) ResetEndpointType() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsIotEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIotEndpoint) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsIotEndpoint) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsIotEndpoint) ToString() *string {
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
func (d *jsiiProxy_DataAwsIotEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsIotEndpointConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/iot_endpoint.html#endpoint_type DataAwsIotEndpoint#endpoint_type}.
	EndpointType *string `json:"endpointType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html aws_iot_authorizer}.
type IotAuthorizer interface {
	cdktf.TerraformResource
	Arn() *string
	AuthorizerFunctionArn() *string
	SetAuthorizerFunctionArn(val *string)
	AuthorizerFunctionArnInput() *string
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
	SigningDisabled() interface{}
	SetSigningDisabled(val interface{})
	SigningDisabledInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TokenKeyName() *string
	SetTokenKeyName(val *string)
	TokenKeyNameInput() *string
	TokenSigningPublicKeys() interface{}
	SetTokenSigningPublicKeys(val interface{})
	TokenSigningPublicKeysInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetSigningDisabled()
	ResetStatus()
	ResetTokenKeyName()
	ResetTokenSigningPublicKeys()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotAuthorizer
type jsiiProxy_IotAuthorizer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotAuthorizer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) AuthorizerFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) AuthorizerFunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerFunctionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) SigningDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signingDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) SigningDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signingDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TokenKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TokenKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TokenSigningPublicKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenSigningPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotAuthorizer) TokenSigningPublicKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenSigningPublicKeysInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html aws_iot_authorizer} Resource.
func NewIotAuthorizer(scope constructs.Construct, id *string, config *IotAuthorizerConfig) IotAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_IotAuthorizer{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotAuthorizer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html aws_iot_authorizer} Resource.
func NewIotAuthorizer_Override(i IotAuthorizer, scope constructs.Construct, id *string, config *IotAuthorizerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotAuthorizer",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetAuthorizerFunctionArn(val *string) {
	_jsii_.Set(
		j,
		"authorizerFunctionArn",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetSigningDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"signingDisabled",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetTokenKeyName(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyName",
		val,
	)
}

func (j *jsiiProxy_IotAuthorizer) SetTokenSigningPublicKeys(val interface{}) {
	_jsii_.Set(
		j,
		"tokenSigningPublicKeys",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotAuthorizer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotAuthorizer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotAuthorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotAuthorizer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAuthorizer) ResetSigningDisabled() {
	_jsii_.InvokeVoid(
		i,
		"resetSigningDisabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAuthorizer) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAuthorizer) ResetTokenKeyName() {
	_jsii_.InvokeVoid(
		i,
		"resetTokenKeyName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAuthorizer) ResetTokenSigningPublicKeys() {
	_jsii_.InvokeVoid(
		i,
		"resetTokenSigningPublicKeys",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotAuthorizer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotAuthorizer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotAuthorizer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotAuthorizerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#authorizer_function_arn IotAuthorizer#authorizer_function_arn}.
	AuthorizerFunctionArn *string `json:"authorizerFunctionArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#name IotAuthorizer#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#signing_disabled IotAuthorizer#signing_disabled}.
	SigningDisabled interface{} `json:"signingDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#status IotAuthorizer#status}.
	Status *string `json:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#token_key_name IotAuthorizer#token_key_name}.
	TokenKeyName *string `json:"tokenKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_authorizer.html#token_signing_public_keys IotAuthorizer#token_signing_public_keys}.
	TokenSigningPublicKeys interface{} `json:"tokenSigningPublicKeys"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_certificate.html aws_iot_certificate}.
type IotCertificate interface {
	cdktf.TerraformResource
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CertificatePem() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Csr() *string
	SetCsr(val *string)
	CsrInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PrivateKey() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PublicKey() *string
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
	ResetCsr()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotCertificate
type jsiiProxy_IotCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotCertificate) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) CertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Csr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) CsrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_certificate.html aws_iot_certificate} Resource.
func NewIotCertificate(scope constructs.Construct, id *string, config *IotCertificateConfig) IotCertificate {
	_init_.Initialize()

	j := jsiiProxy_IotCertificate{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_certificate.html aws_iot_certificate} Resource.
func NewIotCertificate_Override(i IotCertificate, scope constructs.Construct, id *string, config *IotCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotCertificate",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotCertificate) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_IotCertificate) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotCertificate) SetCsr(val *string) {
	_jsii_.Set(
		j,
		"csr",
		val,
	)
}

func (j *jsiiProxy_IotCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotCertificate) SetProvider(val cdktf.TerraformProvider) {
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
func IotCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotCertificate) ResetCsr() {
	_jsii_.InvokeVoid(
		i,
		"resetCsr",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCertificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotCertificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotCertificateConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_certificate.html#active IotCertificate#active}.
	Active interface{} `json:"active"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_certificate.html#csr IotCertificate#csr}.
	Csr *string `json:"csr"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_policy.html aws_iot_policy}.
type IotPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultVersionId() *string
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
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
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

// The jsii proxy struct for IotPolicy
type jsiiProxy_IotPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) DefaultVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_policy.html aws_iot_policy} Resource.
func NewIotPolicy(scope constructs.Construct, id *string, config *IotPolicyConfig) IotPolicy {
	_init_.Initialize()

	j := jsiiProxy_IotPolicy{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_policy.html aws_iot_policy} Resource.
func NewIotPolicy_Override(i IotPolicy, scope constructs.Construct, id *string, config *IotPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotPolicy",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotPolicy) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_IotPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func IotPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_policy_attachment.html aws_iot_policy_attachment}.
type IotPolicyAttachment interface {
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
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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

// The jsii proxy struct for IotPolicyAttachment
type jsiiProxy_IotPolicyAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotPolicyAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotPolicyAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_policy_attachment.html aws_iot_policy_attachment} Resource.
func NewIotPolicyAttachment(scope constructs.Construct, id *string, config *IotPolicyAttachmentConfig) IotPolicyAttachment {
	_init_.Initialize()

	j := jsiiProxy_IotPolicyAttachment{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotPolicyAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_policy_attachment.html aws_iot_policy_attachment} Resource.
func NewIotPolicyAttachment_Override(i IotPolicyAttachment, scope constructs.Construct, id *string, config *IotPolicyAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotPolicyAttachment",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotPolicyAttachment) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotPolicyAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotPolicyAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotPolicyAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotPolicyAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotPolicyAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotPolicyAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotPolicyAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotPolicyAttachmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_policy_attachment.html#policy IotPolicyAttachment#policy}.
	Policy *string `json:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_policy_attachment.html#target IotPolicyAttachment#target}.
	Target *string `json:"target"`
}

type IotPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_policy.html#name IotPolicy#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_policy.html#policy IotPolicy#policy}.
	Policy *string `json:"policy"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html aws_iot_role_alias}.
type IotRoleAlias interface {
	cdktf.TerraformResource
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CredentialDuration() *float64
	SetCredentialDuration(val *float64)
	CredentialDurationInput() *float64
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
	ResetCredentialDuration()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotRoleAlias
type jsiiProxy_IotRoleAlias struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotRoleAlias) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) CredentialDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"credentialDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) CredentialDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"credentialDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotRoleAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html aws_iot_role_alias} Resource.
func NewIotRoleAlias(scope constructs.Construct, id *string, config *IotRoleAliasConfig) IotRoleAlias {
	_init_.Initialize()

	j := jsiiProxy_IotRoleAlias{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotRoleAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html aws_iot_role_alias} Resource.
func NewIotRoleAlias_Override(i IotRoleAlias, scope constructs.Construct, id *string, config *IotRoleAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotRoleAlias",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetCredentialDuration(val *float64) {
	_jsii_.Set(
		j,
		"credentialDuration",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotRoleAlias) SetRoleArn(val *string) {
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
func IotRoleAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotRoleAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotRoleAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotRoleAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotRoleAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotRoleAlias) ResetCredentialDuration() {
	_jsii_.InvokeVoid(
		i,
		"resetCredentialDuration",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotRoleAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotRoleAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotRoleAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotRoleAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotRoleAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotRoleAliasConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html#alias IotRoleAlias#alias}.
	Alias *string `json:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html#role_arn IotRoleAlias#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_role_alias.html#credential_duration IotRoleAlias#credential_duration}.
	CredentialDuration *float64 `json:"credentialDuration"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html aws_iot_thing}.
type IotThing interface {
	cdktf.TerraformResource
	Arn() *string
	Attributes() interface{}
	SetAttributes(val interface{})
	AttributesInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultClientId() *string
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
	ThingTypeName() *string
	SetThingTypeName(val *string)
	ThingTypeNameInput() *string
	Version() *float64
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAttributes()
	ResetOverrideLogicalId()
	ResetThingTypeName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotThing
type jsiiProxy_IotThing struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotThing) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) DefaultClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) ThingTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) ThingTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThing) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html aws_iot_thing} Resource.
func NewIotThing(scope constructs.Construct, id *string, config *IotThingConfig) IotThing {
	_init_.Initialize()

	j := jsiiProxy_IotThing{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThing",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html aws_iot_thing} Resource.
func NewIotThing_Override(i IotThing, scope constructs.Construct, id *string, config *IotThingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThing",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotThing) SetAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotThing) SetThingTypeName(val *string) {
	_jsii_.Set(
		j,
		"thingTypeName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotThing_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotThing",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotThing_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotThing",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotThing) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotThing) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotThing) ResetAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetAttributes",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotThing) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThing) ResetThingTypeName() {
	_jsii_.InvokeVoid(
		i,
		"resetThingTypeName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThing) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThing) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotThing) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotThing) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotThingConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html#name IotThing#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html#attributes IotThing#attributes}.
	Attributes interface{} `json:"attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing.html#thing_type_name IotThing#thing_type_name}.
	ThingTypeName *string `json:"thingTypeName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_principal_attachment.html aws_iot_thing_principal_attachment}.
type IotThingPrincipalAttachment interface {
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
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Thing() *string
	SetThing(val *string)
	ThingInput() *string
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

// The jsii proxy struct for IotThingPrincipalAttachment
type jsiiProxy_IotThingPrincipalAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotThingPrincipalAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) Thing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingPrincipalAttachment) ThingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_principal_attachment.html aws_iot_thing_principal_attachment} Resource.
func NewIotThingPrincipalAttachment(scope constructs.Construct, id *string, config *IotThingPrincipalAttachmentConfig) IotThingPrincipalAttachment {
	_init_.Initialize()

	j := jsiiProxy_IotThingPrincipalAttachment{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingPrincipalAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_principal_attachment.html aws_iot_thing_principal_attachment} Resource.
func NewIotThingPrincipalAttachment_Override(i IotThingPrincipalAttachment, scope constructs.Construct, id *string, config *IotThingPrincipalAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingPrincipalAttachment",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotThingPrincipalAttachment) SetThing(val *string) {
	_jsii_.Set(
		j,
		"thing",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotThingPrincipalAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotThingPrincipalAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotThingPrincipalAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotThingPrincipalAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingPrincipalAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotThingPrincipalAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotThingPrincipalAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotThingPrincipalAttachmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_principal_attachment.html#principal IotThingPrincipalAttachment#principal}.
	Principal *string `json:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_principal_attachment.html#thing IotThingPrincipalAttachment#thing}.
	Thing *string `json:"thing"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html aws_iot_thing_type}.
type IotThingType interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Deprecated() interface{}
	SetDeprecated(val interface{})
	DeprecatedInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Properties() IotThingTypePropertiesOutputReference
	PropertiesInput() *IotThingTypeProperties
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
	PutProperties(value *IotThingTypeProperties)
	ResetDeprecated()
	ResetOverrideLogicalId()
	ResetProperties()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotThingType
type jsiiProxy_IotThingType struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotThingType) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Deprecated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) DeprecatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Properties() IotThingTypePropertiesOutputReference {
	var returns IotThingTypePropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) PropertiesInput() *IotThingTypeProperties {
	var returns *IotThingTypeProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html aws_iot_thing_type} Resource.
func NewIotThingType(scope constructs.Construct, id *string, config *IotThingTypeConfig) IotThingType {
	_init_.Initialize()

	j := jsiiProxy_IotThingType{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html aws_iot_thing_type} Resource.
func NewIotThingType_Override(i IotThingType, scope constructs.Construct, id *string, config *IotThingTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingType",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotThingType) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotThingType) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotThingType) SetDeprecated(val interface{}) {
	_jsii_.Set(
		j,
		"deprecated",
		val,
	)
}

func (j *jsiiProxy_IotThingType) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotThingType) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotThingType) SetProvider(val cdktf.TerraformProvider) {
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
func IotThingType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotThingType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotThingType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotThingType",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotThingType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotThingType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotThingType) PutProperties(value *IotThingTypeProperties) {
	_jsii_.InvokeVoid(
		i,
		"putProperties",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotThingType) ResetDeprecated() {
	_jsii_.InvokeVoid(
		i,
		"resetDeprecated",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotThingType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingType) ResetProperties() {
	_jsii_.InvokeVoid(
		i,
		"resetProperties",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotThingType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotThingType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotThingTypeConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html#name IotThingType#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html#deprecated IotThingType#deprecated}.
	Deprecated interface{} `json:"deprecated"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html#properties IotThingType#properties}
	Properties *IotThingTypeProperties `json:"properties"`
}

type IotThingTypeProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html#description IotThingType#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_type.html#searchable_attributes IotThingType#searchable_attributes}.
	SearchableAttributes *[]*string `json:"searchableAttributes"`
}

type IotThingTypePropertiesOutputReference interface {
	cdktf.ComplexObject
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SearchableAttributes() *[]*string
	SetSearchableAttributes(val *[]*string)
	SearchableAttributesInput() *[]*string
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
	ResetDescription()
	ResetSearchableAttributes()
}

// The jsii proxy struct for IotThingTypePropertiesOutputReference
type jsiiProxy_IotThingTypePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SearchableAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchableAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SearchableAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchableAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotThingTypePropertiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotThingTypePropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotThingTypePropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingTypePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotThingTypePropertiesOutputReference_Override(i IotThingTypePropertiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotThingTypePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SetSearchableAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"searchableAttributes",
		val,
	)
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotThingTypePropertiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotThingTypePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingTypePropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotThingTypePropertiesOutputReference) ResetSearchableAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetSearchableAttributes",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html aws_iot_topic_rule}.
type IotTopicRule interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CloudwatchAlarm() *[]*IotTopicRuleCloudwatchAlarm
	SetCloudwatchAlarm(val *[]*IotTopicRuleCloudwatchAlarm)
	CloudwatchAlarmInput() *[]*IotTopicRuleCloudwatchAlarm
	CloudwatchMetric() *[]*IotTopicRuleCloudwatchMetric
	SetCloudwatchMetric(val *[]*IotTopicRuleCloudwatchMetric)
	CloudwatchMetricInput() *[]*IotTopicRuleCloudwatchMetric
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dynamodb() *[]*IotTopicRuleDynamodb
	SetDynamodb(val *[]*IotTopicRuleDynamodb)
	DynamodbInput() *[]*IotTopicRuleDynamodb
	Dynamodbv2() *[]*IotTopicRuleDynamodbv2
	SetDynamodbv2(val *[]*IotTopicRuleDynamodbv2)
	Dynamodbv2Input() *[]*IotTopicRuleDynamodbv2
	Elasticsearch() *[]*IotTopicRuleElasticsearch
	SetElasticsearch(val *[]*IotTopicRuleElasticsearch)
	ElasticsearchInput() *[]*IotTopicRuleElasticsearch
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ErrorAction() IotTopicRuleErrorActionOutputReference
	ErrorActionInput() *IotTopicRuleErrorAction
	Firehose() *[]*IotTopicRuleFirehose
	SetFirehose(val *[]*IotTopicRuleFirehose)
	FirehoseInput() *[]*IotTopicRuleFirehose
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IotAnalytics() *[]*IotTopicRuleIotAnalytics
	SetIotAnalytics(val *[]*IotTopicRuleIotAnalytics)
	IotAnalyticsInput() *[]*IotTopicRuleIotAnalytics
	IotEvents() *[]*IotTopicRuleIotEvents
	SetIotEvents(val *[]*IotTopicRuleIotEvents)
	IotEventsInput() *[]*IotTopicRuleIotEvents
	Kinesis() *[]*IotTopicRuleKinesis
	SetKinesis(val *[]*IotTopicRuleKinesis)
	KinesisInput() *[]*IotTopicRuleKinesis
	Lambda() *[]*IotTopicRuleLambda
	SetLambda(val *[]*IotTopicRuleLambda)
	LambdaInput() *[]*IotTopicRuleLambda
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Republish() *[]*IotTopicRuleRepublish
	SetRepublish(val *[]*IotTopicRuleRepublish)
	RepublishInput() *[]*IotTopicRuleRepublish
	S3() *[]*IotTopicRuleS3
	SetS3(val *[]*IotTopicRuleS3)
	S3Input() *[]*IotTopicRuleS3
	Sns() *[]*IotTopicRuleSns
	SetSns(val *[]*IotTopicRuleSns)
	SnsInput() *[]*IotTopicRuleSns
	Sql() *string
	SetSql(val *string)
	SqlInput() *string
	SqlVersion() *string
	SetSqlVersion(val *string)
	SqlVersionInput() *string
	Sqs() *[]*IotTopicRuleSqs
	SetSqs(val *[]*IotTopicRuleSqs)
	SqsInput() *[]*IotTopicRuleSqs
	StepFunctions() *[]*IotTopicRuleStepFunctions
	SetStepFunctions(val *[]*IotTopicRuleStepFunctions)
	StepFunctionsInput() *[]*IotTopicRuleStepFunctions
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
	PutErrorAction(value *IotTopicRuleErrorAction)
	ResetCloudwatchAlarm()
	ResetCloudwatchMetric()
	ResetDescription()
	ResetDynamodb()
	ResetDynamodbv2()
	ResetElasticsearch()
	ResetErrorAction()
	ResetFirehose()
	ResetIotAnalytics()
	ResetIotEvents()
	ResetKinesis()
	ResetLambda()
	ResetOverrideLogicalId()
	ResetRepublish()
	ResetS3()
	ResetSns()
	ResetSqs()
	ResetStepFunctions()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for IotTopicRule
type jsiiProxy_IotTopicRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotTopicRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchAlarm() *[]*IotTopicRuleCloudwatchAlarm {
	var returns *[]*IotTopicRuleCloudwatchAlarm
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchAlarmInput() *[]*IotTopicRuleCloudwatchAlarm {
	var returns *[]*IotTopicRuleCloudwatchAlarm
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchMetric() *[]*IotTopicRuleCloudwatchMetric {
	var returns *[]*IotTopicRuleCloudwatchMetric
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchMetricInput() *[]*IotTopicRuleCloudwatchMetric {
	var returns *[]*IotTopicRuleCloudwatchMetric
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodb() *[]*IotTopicRuleDynamodb {
	var returns *[]*IotTopicRuleDynamodb
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DynamodbInput() *[]*IotTopicRuleDynamodb {
	var returns *[]*IotTopicRuleDynamodb
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodbv2() *[]*IotTopicRuleDynamodbv2 {
	var returns *[]*IotTopicRuleDynamodbv2
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodbv2Input() *[]*IotTopicRuleDynamodbv2 {
	var returns *[]*IotTopicRuleDynamodbv2
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Elasticsearch() *[]*IotTopicRuleElasticsearch {
	var returns *[]*IotTopicRuleElasticsearch
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ElasticsearchInput() *[]*IotTopicRuleElasticsearch {
	var returns *[]*IotTopicRuleElasticsearch
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ErrorAction() IotTopicRuleErrorActionOutputReference {
	var returns IotTopicRuleErrorActionOutputReference
	_jsii_.Get(
		j,
		"errorAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ErrorActionInput() *IotTopicRuleErrorAction {
	var returns *IotTopicRuleErrorAction
	_jsii_.Get(
		j,
		"errorActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Firehose() *[]*IotTopicRuleFirehose {
	var returns *[]*IotTopicRuleFirehose
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) FirehoseInput() *[]*IotTopicRuleFirehose {
	var returns *[]*IotTopicRuleFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotAnalytics() *[]*IotTopicRuleIotAnalytics {
	var returns *[]*IotTopicRuleIotAnalytics
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotAnalyticsInput() *[]*IotTopicRuleIotAnalytics {
	var returns *[]*IotTopicRuleIotAnalytics
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotEvents() *[]*IotTopicRuleIotEvents {
	var returns *[]*IotTopicRuleIotEvents
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotEventsInput() *[]*IotTopicRuleIotEvents {
	var returns *[]*IotTopicRuleIotEvents
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Kinesis() *[]*IotTopicRuleKinesis {
	var returns *[]*IotTopicRuleKinesis
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) KinesisInput() *[]*IotTopicRuleKinesis {
	var returns *[]*IotTopicRuleKinesis
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Lambda() *[]*IotTopicRuleLambda {
	var returns *[]*IotTopicRuleLambda
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) LambdaInput() *[]*IotTopicRuleLambda {
	var returns *[]*IotTopicRuleLambda
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Republish() *[]*IotTopicRuleRepublish {
	var returns *[]*IotTopicRuleRepublish
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) RepublishInput() *[]*IotTopicRuleRepublish {
	var returns *[]*IotTopicRuleRepublish
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) S3() *[]*IotTopicRuleS3 {
	var returns *[]*IotTopicRuleS3
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) S3Input() *[]*IotTopicRuleS3 {
	var returns *[]*IotTopicRuleS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sns() *[]*IotTopicRuleSns {
	var returns *[]*IotTopicRuleSns
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SnsInput() *[]*IotTopicRuleSns {
	var returns *[]*IotTopicRuleSns
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sqs() *[]*IotTopicRuleSqs {
	var returns *[]*IotTopicRuleSqs
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqsInput() *[]*IotTopicRuleSqs {
	var returns *[]*IotTopicRuleSqs
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) StepFunctions() *[]*IotTopicRuleStepFunctions {
	var returns *[]*IotTopicRuleStepFunctions
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) StepFunctionsInput() *[]*IotTopicRuleStepFunctions {
	var returns *[]*IotTopicRuleStepFunctions
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html aws_iot_topic_rule} Resource.
func NewIotTopicRule(scope constructs.Construct, id *string, config *IotTopicRuleConfig) IotTopicRule {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRule{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html aws_iot_topic_rule} Resource.
func NewIotTopicRule_Override(i IotTopicRule, scope constructs.Construct, id *string, config *IotTopicRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRule",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotTopicRule) SetCloudwatchAlarm(val *[]*IotTopicRuleCloudwatchAlarm) {
	_jsii_.Set(
		j,
		"cloudwatchAlarm",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetCloudwatchMetric(val *[]*IotTopicRuleCloudwatchMetric) {
	_jsii_.Set(
		j,
		"cloudwatchMetric",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetDynamodb(val *[]*IotTopicRuleDynamodb) {
	_jsii_.Set(
		j,
		"dynamodb",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetDynamodbv2(val *[]*IotTopicRuleDynamodbv2) {
	_jsii_.Set(
		j,
		"dynamodbv2",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetElasticsearch(val *[]*IotTopicRuleElasticsearch) {
	_jsii_.Set(
		j,
		"elasticsearch",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetFirehose(val *[]*IotTopicRuleFirehose) {
	_jsii_.Set(
		j,
		"firehose",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetIotAnalytics(val *[]*IotTopicRuleIotAnalytics) {
	_jsii_.Set(
		j,
		"iotAnalytics",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetIotEvents(val *[]*IotTopicRuleIotEvents) {
	_jsii_.Set(
		j,
		"iotEvents",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetKinesis(val *[]*IotTopicRuleKinesis) {
	_jsii_.Set(
		j,
		"kinesis",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetLambda(val *[]*IotTopicRuleLambda) {
	_jsii_.Set(
		j,
		"lambda",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetRepublish(val *[]*IotTopicRuleRepublish) {
	_jsii_.Set(
		j,
		"republish",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetS3(val *[]*IotTopicRuleS3) {
	_jsii_.Set(
		j,
		"s3",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetSns(val *[]*IotTopicRuleSns) {
	_jsii_.Set(
		j,
		"sns",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetSql(val *string) {
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetSqlVersion(val *string) {
	_jsii_.Set(
		j,
		"sqlVersion",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetSqs(val *[]*IotTopicRuleSqs) {
	_jsii_.Set(
		j,
		"sqs",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetStepFunctions(val *[]*IotTopicRuleStepFunctions) {
	_jsii_.Set(
		j,
		"stepFunctions",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule) SetTagsAll(val interface{}) {
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
func IotTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.IoT.IotTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotTopicRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.IoT.IotTopicRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (i *jsiiProxy_IotTopicRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotTopicRule) PutErrorAction(value *IotTopicRuleErrorAction) {
	_jsii_.InvokeVoid(
		i,
		"putErrorAction",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDynamodb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetErrorAction() {
	_jsii_.InvokeVoid(
		i,
		"resetErrorAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (i *jsiiProxy_IotTopicRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetTagsAll() {
	_jsii_.InvokeVoid(
		i,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (i *jsiiProxy_IotTopicRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotTopicRuleCloudwatchAlarm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#alarm_name IotTopicRule#alarm_name}.
	AlarmName *string `json:"alarmName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_reason IotTopicRule#state_reason}.
	StateReason *string `json:"stateReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_value IotTopicRule#state_value}.
	StateValue *string `json:"stateValue"`
}

type IotTopicRuleCloudwatchMetric struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_name IotTopicRule#metric_name}.
	MetricName *string `json:"metricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_namespace IotTopicRule#metric_namespace}.
	MetricNamespace *string `json:"metricNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_unit IotTopicRule#metric_unit}.
	MetricUnit *string `json:"metricUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_value IotTopicRule#metric_value}.
	MetricValue *string `json:"metricValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_timestamp IotTopicRule#metric_timestamp}.
	MetricTimestamp *string `json:"metricTimestamp"`
}

type IotTopicRuleConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#enabled IotTopicRule#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#name IotTopicRule#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sql IotTopicRule#sql}.
	Sql *string `json:"sql"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sql_version IotTopicRule#sql_version}.
	SqlVersion *string `json:"sqlVersion"`
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}
	CloudwatchAlarm *[]*IotTopicRuleCloudwatchAlarm `json:"cloudwatchAlarm"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#cloudwatch_metric IotTopicRule#cloudwatch_metric}
	CloudwatchMetric *[]*IotTopicRuleCloudwatchMetric `json:"cloudwatchMetric"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#description IotTopicRule#description}.
	Description *string `json:"description"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#dynamodb IotTopicRule#dynamodb}
	Dynamodb *[]*IotTopicRuleDynamodb `json:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#dynamodbv2 IotTopicRule#dynamodbv2}
	Dynamodbv2 *[]*IotTopicRuleDynamodbv2 `json:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#elasticsearch IotTopicRule#elasticsearch}
	Elasticsearch *[]*IotTopicRuleElasticsearch `json:"elasticsearch"`
	// error_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#error_action IotTopicRule#error_action}
	ErrorAction *IotTopicRuleErrorAction `json:"errorAction"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#firehose IotTopicRule#firehose}
	Firehose *[]*IotTopicRuleFirehose `json:"firehose"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#iot_analytics IotTopicRule#iot_analytics}
	IotAnalytics *[]*IotTopicRuleIotAnalytics `json:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#iot_events IotTopicRule#iot_events}
	IotEvents *[]*IotTopicRuleIotEvents `json:"iotEvents"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#kinesis IotTopicRule#kinesis}
	Kinesis *[]*IotTopicRuleKinesis `json:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#lambda IotTopicRule#lambda}
	Lambda *[]*IotTopicRuleLambda `json:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#republish IotTopicRule#republish}
	Republish *[]*IotTopicRuleRepublish `json:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#s3 IotTopicRule#s3}
	S3 *[]*IotTopicRuleS3 `json:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sns IotTopicRule#sns}
	Sns *[]*IotTopicRuleSns `json:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sqs IotTopicRule#sqs}
	Sqs *[]*IotTopicRuleSqs `json:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#step_functions IotTopicRule#step_functions}
	StepFunctions *[]*IotTopicRuleStepFunctions `json:"stepFunctions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#tags IotTopicRule#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#tags_all IotTopicRule#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type IotTopicRuleDynamodb struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_field IotTopicRule#hash_key_field}.
	HashKeyField *string `json:"hashKeyField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_value IotTopicRule#hash_key_value}.
	HashKeyValue *string `json:"hashKeyValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#table_name IotTopicRule#table_name}.
	TableName *string `json:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_type IotTopicRule#hash_key_type}.
	HashKeyType *string `json:"hashKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#operation IotTopicRule#operation}.
	Operation *string `json:"operation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#payload_field IotTopicRule#payload_field}.
	PayloadField *string `json:"payloadField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_field IotTopicRule#range_key_field}.
	RangeKeyField *string `json:"rangeKeyField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_type IotTopicRule#range_key_type}.
	RangeKeyType *string `json:"rangeKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_value IotTopicRule#range_key_value}.
	RangeKeyValue *string `json:"rangeKeyValue"`
}

type IotTopicRuleDynamodbv2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#put_item IotTopicRule#put_item}
	PutItem *IotTopicRuleDynamodbv2PutItem `json:"putItem"`
}

type IotTopicRuleDynamodbv2PutItem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#table_name IotTopicRule#table_name}.
	TableName *string `json:"tableName"`
}

type IotTopicRuleDynamodbv2PutItemOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
}

// The jsii proxy struct for IotTopicRuleDynamodbv2PutItemOutputReference
type jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleDynamodbv2PutItemOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleDynamodbv2PutItemOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleDynamodbv2PutItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleDynamodbv2PutItemOutputReference_Override(i IotTopicRuleDynamodbv2PutItemOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleDynamodbv2PutItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleDynamodbv2PutItemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleElasticsearch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#endpoint IotTopicRule#endpoint}.
	Endpoint *string `json:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#id IotTopicRule#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#index IotTopicRule#index}.
	Index *string `json:"index"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#type IotTopicRule#type}.
	Type *string `json:"type"`
}

type IotTopicRuleErrorAction struct {
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}
	CloudwatchAlarm *IotTopicRuleErrorActionCloudwatchAlarm `json:"cloudwatchAlarm"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#cloudwatch_metric IotTopicRule#cloudwatch_metric}
	CloudwatchMetric *IotTopicRuleErrorActionCloudwatchMetric `json:"cloudwatchMetric"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#dynamodb IotTopicRule#dynamodb}
	Dynamodb *IotTopicRuleErrorActionDynamodb `json:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#dynamodbv2 IotTopicRule#dynamodbv2}
	Dynamodbv2 *IotTopicRuleErrorActionDynamodbv2 `json:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#elasticsearch IotTopicRule#elasticsearch}
	Elasticsearch *IotTopicRuleErrorActionElasticsearch `json:"elasticsearch"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#firehose IotTopicRule#firehose}
	Firehose *IotTopicRuleErrorActionFirehose `json:"firehose"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#iot_analytics IotTopicRule#iot_analytics}
	IotAnalytics *IotTopicRuleErrorActionIotAnalytics `json:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#iot_events IotTopicRule#iot_events}
	IotEvents *IotTopicRuleErrorActionIotEvents `json:"iotEvents"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#kinesis IotTopicRule#kinesis}
	Kinesis *IotTopicRuleErrorActionKinesis `json:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#lambda IotTopicRule#lambda}
	Lambda *IotTopicRuleErrorActionLambda `json:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#republish IotTopicRule#republish}
	Republish *IotTopicRuleErrorActionRepublish `json:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#s3 IotTopicRule#s3}
	S3 *IotTopicRuleErrorActionS3 `json:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sns IotTopicRule#sns}
	Sns *IotTopicRuleErrorActionSns `json:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#sqs IotTopicRule#sqs}
	Sqs *IotTopicRuleErrorActionSqs `json:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#step_functions IotTopicRule#step_functions}
	StepFunctions *IotTopicRuleErrorActionStepFunctions `json:"stepFunctions"`
}

type IotTopicRuleErrorActionCloudwatchAlarm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#alarm_name IotTopicRule#alarm_name}.
	AlarmName *string `json:"alarmName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_reason IotTopicRule#state_reason}.
	StateReason *string `json:"stateReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_value IotTopicRule#state_value}.
	StateValue *string `json:"stateValue"`
}

type IotTopicRuleErrorActionCloudwatchAlarmOutputReference interface {
	cdktf.ComplexObject
	AlarmName() *string
	SetAlarmName(val *string)
	AlarmNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StateReason() *string
	SetStateReason(val *string)
	StateReasonInput() *string
	StateValue() *string
	SetStateValue(val *string)
	StateValueInput() *string
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
}

// The jsii proxy struct for IotTopicRuleErrorActionCloudwatchAlarmOutputReference
type jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) StateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) StateReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) StateValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) StateValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionCloudwatchAlarmOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionCloudwatchAlarmOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionCloudwatchAlarmOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionCloudwatchAlarmOutputReference_Override(i IotTopicRuleErrorActionCloudwatchAlarmOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionCloudwatchAlarmOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetAlarmName(val *string) {
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetStateReason(val *string) {
	_jsii_.Set(
		j,
		"stateReason",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetStateValue(val *string) {
	_jsii_.Set(
		j,
		"stateValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchAlarmOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionCloudwatchMetric struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_name IotTopicRule#metric_name}.
	MetricName *string `json:"metricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_namespace IotTopicRule#metric_namespace}.
	MetricNamespace *string `json:"metricNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_unit IotTopicRule#metric_unit}.
	MetricUnit *string `json:"metricUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_value IotTopicRule#metric_value}.
	MetricValue *string `json:"metricValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#metric_timestamp IotTopicRule#metric_timestamp}.
	MetricTimestamp *string `json:"metricTimestamp"`
}

type IotTopicRuleErrorActionCloudwatchMetricOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricNamespace() *string
	SetMetricNamespace(val *string)
	MetricNamespaceInput() *string
	MetricTimestamp() *string
	SetMetricTimestamp(val *string)
	MetricTimestampInput() *string
	MetricUnit() *string
	SetMetricUnit(val *string)
	MetricUnitInput() *string
	MetricValue() *string
	SetMetricValue(val *string)
	MetricValueInput() *string
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
	ResetMetricTimestamp()
}

// The jsii proxy struct for IotTopicRuleErrorActionCloudwatchMetricOutputReference
type jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) MetricValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionCloudwatchMetricOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionCloudwatchMetricOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionCloudwatchMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionCloudwatchMetricOutputReference_Override(i IotTopicRuleErrorActionCloudwatchMetricOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionCloudwatchMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetMetricNamespace(val *string) {
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetMetricTimestamp(val *string) {
	_jsii_.Set(
		j,
		"metricTimestamp",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetMetricUnit(val *string) {
	_jsii_.Set(
		j,
		"metricUnit",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetMetricValue(val *string) {
	_jsii_.Set(
		j,
		"metricValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionCloudwatchMetricOutputReference) ResetMetricTimestamp() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricTimestamp",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionDynamodb struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_field IotTopicRule#hash_key_field}.
	HashKeyField *string `json:"hashKeyField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_value IotTopicRule#hash_key_value}.
	HashKeyValue *string `json:"hashKeyValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#table_name IotTopicRule#table_name}.
	TableName *string `json:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#hash_key_type IotTopicRule#hash_key_type}.
	HashKeyType *string `json:"hashKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#operation IotTopicRule#operation}.
	Operation *string `json:"operation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#payload_field IotTopicRule#payload_field}.
	PayloadField *string `json:"payloadField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_field IotTopicRule#range_key_field}.
	RangeKeyField *string `json:"rangeKeyField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_type IotTopicRule#range_key_type}.
	RangeKeyType *string `json:"rangeKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#range_key_value IotTopicRule#range_key_value}.
	RangeKeyValue *string `json:"rangeKeyValue"`
}

type IotTopicRuleErrorActionDynamodbOutputReference interface {
	cdktf.ComplexObject
	HashKeyField() *string
	SetHashKeyField(val *string)
	HashKeyFieldInput() *string
	HashKeyType() *string
	SetHashKeyType(val *string)
	HashKeyTypeInput() *string
	HashKeyValue() *string
	SetHashKeyValue(val *string)
	HashKeyValueInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Operation() *string
	SetOperation(val *string)
	OperationInput() *string
	PayloadField() *string
	SetPayloadField(val *string)
	PayloadFieldInput() *string
	RangeKeyField() *string
	SetRangeKeyField(val *string)
	RangeKeyFieldInput() *string
	RangeKeyType() *string
	SetRangeKeyType(val *string)
	RangeKeyTypeInput() *string
	RangeKeyValue() *string
	SetRangeKeyValue(val *string)
	RangeKeyValueInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	ResetHashKeyType()
	ResetOperation()
	ResetPayloadField()
	ResetRangeKeyField()
	ResetRangeKeyType()
	ResetRangeKeyValue()
}

// The jsii proxy struct for IotTopicRuleErrorActionDynamodbOutputReference
type jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) HashKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) PayloadField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) PayloadFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RangeKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionDynamodbOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionDynamodbOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionDynamodbOutputReference_Override(i IotTopicRuleErrorActionDynamodbOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetHashKeyField(val *string) {
	_jsii_.Set(
		j,
		"hashKeyField",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetHashKeyType(val *string) {
	_jsii_.Set(
		j,
		"hashKeyType",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetHashKeyValue(val *string) {
	_jsii_.Set(
		j,
		"hashKeyValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetOperation(val *string) {
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetPayloadField(val *string) {
	_jsii_.Set(
		j,
		"payloadField",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetRangeKeyField(val *string) {
	_jsii_.Set(
		j,
		"rangeKeyField",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetRangeKeyType(val *string) {
	_jsii_.Set(
		j,
		"rangeKeyType",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetRangeKeyValue(val *string) {
	_jsii_.Set(
		j,
		"rangeKeyValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetHashKeyType() {
	_jsii_.InvokeVoid(
		i,
		"resetHashKeyType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetOperation() {
	_jsii_.InvokeVoid(
		i,
		"resetOperation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetPayloadField() {
	_jsii_.InvokeVoid(
		i,
		"resetPayloadField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetRangeKeyField() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetRangeKeyType() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbOutputReference) ResetRangeKeyValue() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyValue",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionDynamodbv2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#put_item IotTopicRule#put_item}
	PutItem *IotTopicRuleErrorActionDynamodbv2PutItem `json:"putItem"`
}

type IotTopicRuleErrorActionDynamodbv2OutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PutItem() IotTopicRuleErrorActionDynamodbv2PutItemOutputReference
	PutItemInput() *IotTopicRuleErrorActionDynamodbv2PutItem
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
	PutPutItem(value *IotTopicRuleErrorActionDynamodbv2PutItem)
	ResetPutItem()
}

// The jsii proxy struct for IotTopicRuleErrorActionDynamodbv2OutputReference
type jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) PutItem() IotTopicRuleErrorActionDynamodbv2PutItemOutputReference {
	var returns IotTopicRuleErrorActionDynamodbv2PutItemOutputReference
	_jsii_.Get(
		j,
		"putItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) PutItemInput() *IotTopicRuleErrorActionDynamodbv2PutItem {
	var returns *IotTopicRuleErrorActionDynamodbv2PutItem
	_jsii_.Get(
		j,
		"putItemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionDynamodbv2OutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionDynamodbv2OutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbv2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionDynamodbv2OutputReference_Override(i IotTopicRuleErrorActionDynamodbv2OutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbv2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) PutPutItem(value *IotTopicRuleErrorActionDynamodbv2PutItem) {
	_jsii_.InvokeVoid(
		i,
		"putPutItem",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2OutputReference) ResetPutItem() {
	_jsii_.InvokeVoid(
		i,
		"resetPutItem",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionDynamodbv2PutItem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#table_name IotTopicRule#table_name}.
	TableName *string `json:"tableName"`
}

type IotTopicRuleErrorActionDynamodbv2PutItemOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
}

// The jsii proxy struct for IotTopicRuleErrorActionDynamodbv2PutItemOutputReference
type jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionDynamodbv2PutItemOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionDynamodbv2PutItemOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbv2PutItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionDynamodbv2PutItemOutputReference_Override(i IotTopicRuleErrorActionDynamodbv2PutItemOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionDynamodbv2PutItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionDynamodbv2PutItemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionElasticsearch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#endpoint IotTopicRule#endpoint}.
	Endpoint *string `json:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#id IotTopicRule#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#index IotTopicRule#index}.
	Index *string `json:"index"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#type IotTopicRule#type}.
	Type *string `json:"type"`
}

type IotTopicRuleErrorActionElasticsearchOutputReference interface {
	cdktf.ComplexObject
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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

// The jsii proxy struct for IotTopicRuleErrorActionElasticsearchOutputReference
type jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionElasticsearchOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionElasticsearchOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionElasticsearchOutputReference_Override(i IotTopicRuleErrorActionElasticsearchOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetIndex(val *string) {
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionElasticsearchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionFirehose struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#delivery_stream_name IotTopicRule#delivery_stream_name}.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#separator IotTopicRule#separator}.
	Separator *string `json:"separator"`
}

type IotTopicRuleErrorActionFirehoseOutputReference interface {
	cdktf.ComplexObject
	DeliveryStreamName() *string
	SetDeliveryStreamName(val *string)
	DeliveryStreamNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Separator() *string
	SetSeparator(val *string)
	SeparatorInput() *string
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
	ResetSeparator()
}

// The jsii proxy struct for IotTopicRuleErrorActionFirehoseOutputReference
type jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) DeliveryStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) Separator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"separator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SeparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"separatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionFirehoseOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionFirehoseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionFirehoseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionFirehoseOutputReference_Override(i IotTopicRuleErrorActionFirehoseOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionFirehoseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetDeliveryStreamName(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetSeparator(val *string) {
	_jsii_.Set(
		j,
		"separator",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionFirehoseOutputReference) ResetSeparator() {
	_jsii_.InvokeVoid(
		i,
		"resetSeparator",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionIotAnalytics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#channel_name IotTopicRule#channel_name}.
	ChannelName *string `json:"channelName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
}

type IotTopicRuleErrorActionIotAnalyticsOutputReference interface {
	cdktf.ComplexObject
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
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
}

// The jsii proxy struct for IotTopicRuleErrorActionIotAnalyticsOutputReference
type jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionIotAnalyticsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionIotAnalyticsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionIotAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionIotAnalyticsOutputReference_Override(i IotTopicRuleErrorActionIotAnalyticsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionIotAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) SetChannelName(val *string) {
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotAnalyticsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionIotEvents struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#input_name IotTopicRule#input_name}.
	InputName *string `json:"inputName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#message_id IotTopicRule#message_id}.
	MessageId *string `json:"messageId"`
}

type IotTopicRuleErrorActionIotEventsOutputReference interface {
	cdktf.ComplexObject
	InputName() *string
	SetInputName(val *string)
	InputNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageId() *string
	SetMessageId(val *string)
	MessageIdInput() *string
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
	ResetMessageId()
}

// The jsii proxy struct for IotTopicRuleErrorActionIotEventsOutputReference
type jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) MessageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) MessageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionIotEventsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionIotEventsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionIotEventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionIotEventsOutputReference_Override(i IotTopicRuleErrorActionIotEventsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionIotEventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetInputName(val *string) {
	_jsii_.Set(
		j,
		"inputName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetMessageId(val *string) {
	_jsii_.Set(
		j,
		"messageId",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ResetMessageId() {
	_jsii_.InvokeVoid(
		i,
		"resetMessageId",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionKinesis struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#stream_name IotTopicRule#stream_name}.
	StreamName *string `json:"streamName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#partition_key IotTopicRule#partition_key}.
	PartitionKey *string `json:"partitionKey"`
}

type IotTopicRuleErrorActionKinesisOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PartitionKey() *string
	SetPartitionKey(val *string)
	PartitionKeyInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StreamName() *string
	SetStreamName(val *string)
	StreamNameInput() *string
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
	ResetPartitionKey()
}

// The jsii proxy struct for IotTopicRuleErrorActionKinesisOutputReference
type jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) PartitionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) PartitionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) StreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionKinesisOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionKinesisOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionKinesisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionKinesisOutputReference_Override(i IotTopicRuleErrorActionKinesisOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionKinesisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetPartitionKey(val *string) {
	_jsii_.Set(
		j,
		"partitionKey",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetStreamName(val *string) {
	_jsii_.Set(
		j,
		"streamName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionKinesisOutputReference) ResetPartitionKey() {
	_jsii_.InvokeVoid(
		i,
		"resetPartitionKey",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionLambda struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#function_arn IotTopicRule#function_arn}.
	FunctionArn *string `json:"functionArn"`
}

type IotTopicRuleErrorActionLambdaOutputReference interface {
	cdktf.ComplexObject
	FunctionArn() *string
	SetFunctionArn(val *string)
	FunctionArnInput() *string
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
}

// The jsii proxy struct for IotTopicRuleErrorActionLambdaOutputReference
type jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) FunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionLambdaOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionLambdaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionLambdaOutputReference_Override(i IotTopicRuleErrorActionLambdaOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) SetFunctionArn(val *string) {
	_jsii_.Set(
		j,
		"functionArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionLambdaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionOutputReference interface {
	cdktf.ComplexObject
	CloudwatchAlarm() IotTopicRuleErrorActionCloudwatchAlarmOutputReference
	CloudwatchAlarmInput() *IotTopicRuleErrorActionCloudwatchAlarm
	CloudwatchMetric() IotTopicRuleErrorActionCloudwatchMetricOutputReference
	CloudwatchMetricInput() *IotTopicRuleErrorActionCloudwatchMetric
	Dynamodb() IotTopicRuleErrorActionDynamodbOutputReference
	DynamodbInput() *IotTopicRuleErrorActionDynamodb
	Dynamodbv2() IotTopicRuleErrorActionDynamodbv2OutputReference
	Dynamodbv2Input() *IotTopicRuleErrorActionDynamodbv2
	Elasticsearch() IotTopicRuleErrorActionElasticsearchOutputReference
	ElasticsearchInput() *IotTopicRuleErrorActionElasticsearch
	Firehose() IotTopicRuleErrorActionFirehoseOutputReference
	FirehoseInput() *IotTopicRuleErrorActionFirehose
	IotAnalytics() IotTopicRuleErrorActionIotAnalyticsOutputReference
	IotAnalyticsInput() *IotTopicRuleErrorActionIotAnalytics
	IotEvents() IotTopicRuleErrorActionIotEventsOutputReference
	IotEventsInput() *IotTopicRuleErrorActionIotEvents
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Kinesis() IotTopicRuleErrorActionKinesisOutputReference
	KinesisInput() *IotTopicRuleErrorActionKinesis
	Lambda() IotTopicRuleErrorActionLambdaOutputReference
	LambdaInput() *IotTopicRuleErrorActionLambda
	Republish() IotTopicRuleErrorActionRepublishOutputReference
	RepublishInput() *IotTopicRuleErrorActionRepublish
	S3() IotTopicRuleErrorActionS3OutputReference
	S3Input() *IotTopicRuleErrorActionS3
	Sns() IotTopicRuleErrorActionSnsOutputReference
	SnsInput() *IotTopicRuleErrorActionSns
	Sqs() IotTopicRuleErrorActionSqsOutputReference
	SqsInput() *IotTopicRuleErrorActionSqs
	StepFunctions() IotTopicRuleErrorActionStepFunctionsOutputReference
	StepFunctionsInput() *IotTopicRuleErrorActionStepFunctions
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
	PutCloudwatchAlarm(value *IotTopicRuleErrorActionCloudwatchAlarm)
	PutCloudwatchMetric(value *IotTopicRuleErrorActionCloudwatchMetric)
	PutDynamodb(value *IotTopicRuleErrorActionDynamodb)
	PutDynamodbv2(value *IotTopicRuleErrorActionDynamodbv2)
	PutElasticsearch(value *IotTopicRuleErrorActionElasticsearch)
	PutFirehose(value *IotTopicRuleErrorActionFirehose)
	PutIotAnalytics(value *IotTopicRuleErrorActionIotAnalytics)
	PutIotEvents(value *IotTopicRuleErrorActionIotEvents)
	PutKinesis(value *IotTopicRuleErrorActionKinesis)
	PutLambda(value *IotTopicRuleErrorActionLambda)
	PutRepublish(value *IotTopicRuleErrorActionRepublish)
	PutS3(value *IotTopicRuleErrorActionS3)
	PutSns(value *IotTopicRuleErrorActionSns)
	PutSqs(value *IotTopicRuleErrorActionSqs)
	PutStepFunctions(value *IotTopicRuleErrorActionStepFunctions)
	ResetCloudwatchAlarm()
	ResetCloudwatchMetric()
	ResetDynamodb()
	ResetDynamodbv2()
	ResetElasticsearch()
	ResetFirehose()
	ResetIotAnalytics()
	ResetIotEvents()
	ResetKinesis()
	ResetLambda()
	ResetRepublish()
	ResetS3()
	ResetSns()
	ResetSqs()
	ResetStepFunctions()
}

// The jsii proxy struct for IotTopicRuleErrorActionOutputReference
type jsiiProxy_IotTopicRuleErrorActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchAlarm() IotTopicRuleErrorActionCloudwatchAlarmOutputReference {
	var returns IotTopicRuleErrorActionCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchAlarmInput() *IotTopicRuleErrorActionCloudwatchAlarm {
	var returns *IotTopicRuleErrorActionCloudwatchAlarm
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchMetric() IotTopicRuleErrorActionCloudwatchMetricOutputReference {
	var returns IotTopicRuleErrorActionCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchMetricInput() *IotTopicRuleErrorActionCloudwatchMetric {
	var returns *IotTopicRuleErrorActionCloudwatchMetric
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodb() IotTopicRuleErrorActionDynamodbOutputReference {
	var returns IotTopicRuleErrorActionDynamodbOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) DynamodbInput() *IotTopicRuleErrorActionDynamodb {
	var returns *IotTopicRuleErrorActionDynamodb
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodbv2() IotTopicRuleErrorActionDynamodbv2OutputReference {
	var returns IotTopicRuleErrorActionDynamodbv2OutputReference
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodbv2Input() *IotTopicRuleErrorActionDynamodbv2 {
	var returns *IotTopicRuleErrorActionDynamodbv2
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Elasticsearch() IotTopicRuleErrorActionElasticsearchOutputReference {
	var returns IotTopicRuleErrorActionElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) ElasticsearchInput() *IotTopicRuleErrorActionElasticsearch {
	var returns *IotTopicRuleErrorActionElasticsearch
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Firehose() IotTopicRuleErrorActionFirehoseOutputReference {
	var returns IotTopicRuleErrorActionFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) FirehoseInput() *IotTopicRuleErrorActionFirehose {
	var returns *IotTopicRuleErrorActionFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotAnalytics() IotTopicRuleErrorActionIotAnalyticsOutputReference {
	var returns IotTopicRuleErrorActionIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotAnalyticsInput() *IotTopicRuleErrorActionIotAnalytics {
	var returns *IotTopicRuleErrorActionIotAnalytics
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotEvents() IotTopicRuleErrorActionIotEventsOutputReference {
	var returns IotTopicRuleErrorActionIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotEventsInput() *IotTopicRuleErrorActionIotEvents {
	var returns *IotTopicRuleErrorActionIotEvents
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Kinesis() IotTopicRuleErrorActionKinesisOutputReference {
	var returns IotTopicRuleErrorActionKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) KinesisInput() *IotTopicRuleErrorActionKinesis {
	var returns *IotTopicRuleErrorActionKinesis
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Lambda() IotTopicRuleErrorActionLambdaOutputReference {
	var returns IotTopicRuleErrorActionLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) LambdaInput() *IotTopicRuleErrorActionLambda {
	var returns *IotTopicRuleErrorActionLambda
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Republish() IotTopicRuleErrorActionRepublishOutputReference {
	var returns IotTopicRuleErrorActionRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) RepublishInput() *IotTopicRuleErrorActionRepublish {
	var returns *IotTopicRuleErrorActionRepublish
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) S3() IotTopicRuleErrorActionS3OutputReference {
	var returns IotTopicRuleErrorActionS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) S3Input() *IotTopicRuleErrorActionS3 {
	var returns *IotTopicRuleErrorActionS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Sns() IotTopicRuleErrorActionSnsOutputReference {
	var returns IotTopicRuleErrorActionSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SnsInput() *IotTopicRuleErrorActionSns {
	var returns *IotTopicRuleErrorActionSns
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Sqs() IotTopicRuleErrorActionSqsOutputReference {
	var returns IotTopicRuleErrorActionSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SqsInput() *IotTopicRuleErrorActionSqs {
	var returns *IotTopicRuleErrorActionSqs
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) StepFunctions() IotTopicRuleErrorActionStepFunctionsOutputReference {
	var returns IotTopicRuleErrorActionStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) StepFunctionsInput() *IotTopicRuleErrorActionStepFunctions {
	var returns *IotTopicRuleErrorActionStepFunctions
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionOutputReference_Override(i IotTopicRuleErrorActionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutCloudwatchAlarm(value *IotTopicRuleErrorActionCloudwatchAlarm) {
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutCloudwatchMetric(value *IotTopicRuleErrorActionCloudwatchMetric) {
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutDynamodb(value *IotTopicRuleErrorActionDynamodb) {
	_jsii_.InvokeVoid(
		i,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutDynamodbv2(value *IotTopicRuleErrorActionDynamodbv2) {
	_jsii_.InvokeVoid(
		i,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutElasticsearch(value *IotTopicRuleErrorActionElasticsearch) {
	_jsii_.InvokeVoid(
		i,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutFirehose(value *IotTopicRuleErrorActionFirehose) {
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutIotAnalytics(value *IotTopicRuleErrorActionIotAnalytics) {
	_jsii_.InvokeVoid(
		i,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutIotEvents(value *IotTopicRuleErrorActionIotEvents) {
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutKinesis(value *IotTopicRuleErrorActionKinesis) {
	_jsii_.InvokeVoid(
		i,
		"putKinesis",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutLambda(value *IotTopicRuleErrorActionLambda) {
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutRepublish(value *IotTopicRuleErrorActionRepublish) {
	_jsii_.InvokeVoid(
		i,
		"putRepublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutS3(value *IotTopicRuleErrorActionS3) {
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutSns(value *IotTopicRuleErrorActionSns) {
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutSqs(value *IotTopicRuleErrorActionSqs) {
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutStepFunctions(value *IotTopicRuleErrorActionStepFunctions) {
	_jsii_.InvokeVoid(
		i,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetDynamodb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionRepublish struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#topic IotTopicRule#topic}.
	Topic *string `json:"topic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#qos IotTopicRule#qos}.
	Qos *float64 `json:"qos"`
}

type IotTopicRuleErrorActionRepublishOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Qos() *float64
	SetQos(val *float64)
	QosInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetQos()
}

// The jsii proxy struct for IotTopicRuleErrorActionRepublishOutputReference
type jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) Qos() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) QosInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionRepublishOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionRepublishOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionRepublishOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionRepublishOutputReference_Override(i IotTopicRuleErrorActionRepublishOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionRepublishOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetQos(val *float64) {
	_jsii_.Set(
		j,
		"qos",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) SetTopic(val *string) {
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionRepublishOutputReference) ResetQos() {
	_jsii_.InvokeVoid(
		i,
		"resetQos",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#bucket_name IotTopicRule#bucket_name}.
	BucketName *string `json:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#key IotTopicRule#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
}

type IotTopicRuleErrorActionS3OutputReference interface {
	cdktf.ComplexObject
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
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
}

// The jsii proxy struct for IotTopicRuleErrorActionS3OutputReference
type jsiiProxy_IotTopicRuleErrorActionS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionS3OutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionS3OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionS3OutputReference_Override(i IotTopicRuleErrorActionS3OutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionSns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#target_arn IotTopicRule#target_arn}.
	TargetArn *string `json:"targetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#message_format IotTopicRule#message_format}.
	MessageFormat *string `json:"messageFormat"`
}

type IotTopicRuleErrorActionSnsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageFormat() *string
	SetMessageFormat(val *string)
	MessageFormatInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TargetArn() *string
	SetTargetArn(val *string)
	TargetArnInput() *string
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
	ResetMessageFormat()
}

// The jsii proxy struct for IotTopicRuleErrorActionSnsOutputReference
type jsiiProxy_IotTopicRuleErrorActionSnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) TargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionSnsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionSnsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionSnsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionSnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionSnsOutputReference_Override(i IotTopicRuleErrorActionSnsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionSnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetMessageFormat(val *string) {
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetTargetArn(val *string) {
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionSnsOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		i,
		"resetMessageFormat",
		nil, // no parameters
	)
}

type IotTopicRuleErrorActionSqs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#queue_url IotTopicRule#queue_url}.
	QueueUrl *string `json:"queueUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#use_base64 IotTopicRule#use_base64}.
	UseBase64 interface{} `json:"useBase64"`
}

type IotTopicRuleErrorActionSqsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueueUrl() *string
	SetQueueUrl(val *string)
	QueueUrlInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UseBase64() interface{}
	SetUseBase64(val interface{})
	UseBase64Input() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for IotTopicRuleErrorActionSqsOutputReference
type jsiiProxy_IotTopicRuleErrorActionSqsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) QueueUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) UseBase64() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) UseBase64Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBase64Input",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionSqsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionSqsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionSqsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionSqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionSqsOutputReference_Override(i IotTopicRuleErrorActionSqsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionSqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetQueueUrl(val *string) {
	_jsii_.Set(
		j,
		"queueUrl",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) SetUseBase64(val interface{}) {
	_jsii_.Set(
		j,
		"useBase64",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionSqsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type IotTopicRuleErrorActionStepFunctions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_machine_name IotTopicRule#state_machine_name}.
	StateMachineName *string `json:"stateMachineName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#execution_name_prefix IotTopicRule#execution_name_prefix}.
	ExecutionNamePrefix *string `json:"executionNamePrefix"`
}

type IotTopicRuleErrorActionStepFunctionsOutputReference interface {
	cdktf.ComplexObject
	ExecutionNamePrefix() *string
	SetExecutionNamePrefix(val *string)
	ExecutionNamePrefixInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StateMachineName() *string
	SetStateMachineName(val *string)
	StateMachineNameInput() *string
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
	ResetExecutionNamePrefix()
}

// The jsii proxy struct for IotTopicRuleErrorActionStepFunctionsOutputReference
type jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) ExecutionNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) ExecutionNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) StateMachineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) StateMachineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionStepFunctionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) IotTopicRuleErrorActionStepFunctionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionStepFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionStepFunctionsOutputReference_Override(i IotTopicRuleErrorActionStepFunctionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.IoT.IotTopicRuleErrorActionStepFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetExecutionNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"executionNamePrefix",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetStateMachineName(val *string) {
	_jsii_.Set(
		j,
		"stateMachineName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionStepFunctionsOutputReference) ResetExecutionNamePrefix() {
	_jsii_.InvokeVoid(
		i,
		"resetExecutionNamePrefix",
		nil, // no parameters
	)
}

type IotTopicRuleFirehose struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#delivery_stream_name IotTopicRule#delivery_stream_name}.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#separator IotTopicRule#separator}.
	Separator *string `json:"separator"`
}

type IotTopicRuleIotAnalytics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#channel_name IotTopicRule#channel_name}.
	ChannelName *string `json:"channelName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
}

type IotTopicRuleIotEvents struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#input_name IotTopicRule#input_name}.
	InputName *string `json:"inputName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#message_id IotTopicRule#message_id}.
	MessageId *string `json:"messageId"`
}

type IotTopicRuleKinesis struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#stream_name IotTopicRule#stream_name}.
	StreamName *string `json:"streamName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#partition_key IotTopicRule#partition_key}.
	PartitionKey *string `json:"partitionKey"`
}

type IotTopicRuleLambda struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#function_arn IotTopicRule#function_arn}.
	FunctionArn *string `json:"functionArn"`
}

type IotTopicRuleRepublish struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#topic IotTopicRule#topic}.
	Topic *string `json:"topic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#qos IotTopicRule#qos}.
	Qos *float64 `json:"qos"`
}

type IotTopicRuleS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#bucket_name IotTopicRule#bucket_name}.
	BucketName *string `json:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#key IotTopicRule#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
}

type IotTopicRuleSns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#target_arn IotTopicRule#target_arn}.
	TargetArn *string `json:"targetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#message_format IotTopicRule#message_format}.
	MessageFormat *string `json:"messageFormat"`
}

type IotTopicRuleSqs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#queue_url IotTopicRule#queue_url}.
	QueueUrl *string `json:"queueUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#use_base64 IotTopicRule#use_base64}.
	UseBase64 interface{} `json:"useBase64"`
}

type IotTopicRuleStepFunctions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#role_arn IotTopicRule#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#state_machine_name IotTopicRule#state_machine_name}.
	StateMachineName *string `json:"stateMachineName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule.html#execution_name_prefix IotTopicRule#execution_name_prefix}.
	ExecutionNamePrefix *string `json:"executionNamePrefix"`
}
