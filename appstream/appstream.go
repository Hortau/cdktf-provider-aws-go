package appstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/appstream/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html aws_appstream_fleet}.
type AppstreamFleet interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeCapacity() AppstreamFleetComputeCapacityOutputReference
	ComputeCapacityInput() *AppstreamFleetComputeCapacity
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisconnectTimeoutInSeconds() *float64
	SetDisconnectTimeoutInSeconds(val *float64)
	DisconnectTimeoutInSecondsInput() *float64
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference
	DomainJoinInfoInput() *AppstreamFleetDomainJoinInfo
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	IdleDisconnectTimeoutInSeconds() *float64
	SetIdleDisconnectTimeoutInSeconds(val *float64)
	IdleDisconnectTimeoutInSecondsInput() *float64
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxUserDurationInSeconds() *float64
	SetMaxUserDurationInSeconds(val *float64)
	MaxUserDurationInSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	StreamView() *string
	SetStreamView(val *string)
	StreamViewInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() AppstreamFleetVpcConfigOutputReference
	VpcConfigInput() *AppstreamFleetVpcConfig
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutComputeCapacity(value *AppstreamFleetComputeCapacity)
	PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo)
	PutVpcConfig(value *AppstreamFleetVpcConfig)
	ResetDescription()
	ResetDisconnectTimeoutInSeconds()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetFleetType()
	ResetIamRoleArn()
	ResetIdleDisconnectTimeoutInSeconds()
	ResetImageArn()
	ResetImageName()
	ResetMaxUserDurationInSeconds()
	ResetOverrideLogicalId()
	ResetStreamView()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamFleet
type jsiiProxy_AppstreamFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacity() AppstreamFleetComputeCapacityOutputReference {
	var returns AppstreamFleetComputeCapacityOutputReference
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacityInput() *AppstreamFleetComputeCapacity {
	var returns *AppstreamFleetComputeCapacity
	_jsii_.Get(
		j,
		"computeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference {
	var returns AppstreamFleetDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfoInput() *AppstreamFleetDomainJoinInfo {
	var returns *AppstreamFleetDomainJoinInfo
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfig() AppstreamFleetVpcConfigOutputReference {
	var returns AppstreamFleetVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfigInput() *AppstreamFleetVpcConfig {
	var returns *AppstreamFleetVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html aws_appstream_fleet} Resource.
func NewAppstreamFleet(scope constructs.Construct, id *string, config *AppstreamFleetConfig) AppstreamFleet {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleet{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html aws_appstream_fleet} Resource.
func NewAppstreamFleet_Override(a AppstreamFleet, scope constructs.Construct, id *string, config *AppstreamFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleet",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"disconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetEnableDefaultInternetAccess(val interface{}) {
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetFleetType(val *string) {
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetIdleDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetImageArn(val *string) {
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetMaxUserDurationInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxUserDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetStreamView(val *string) {
	_jsii_.Set(
		j,
		"streamView",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetTagsAll(val interface{}) {
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
func AppstreamFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppStream.AppstreamFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppStream.AppstreamFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppstreamFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutComputeCapacity(value *AppstreamFleetComputeCapacity) {
	_jsii_.InvokeVoid(
		a,
		"putComputeCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo) {
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutVpcConfig(value *AppstreamFleetVpcConfig) {
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIdleDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetMaxUserDurationInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUserDurationInSeconds",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetStreamView() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamView",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppstreamFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppstreamFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamFleetComputeCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#desired_instances AppstreamFleet#desired_instances}.
	DesiredInstances *float64 `json:"desiredInstances"`
}

type AppstreamFleetComputeCapacityOutputReference interface {
	cdktf.ComplexObject
	DesiredInstances() *float64
	SetDesiredInstances(val *float64)
	DesiredInstancesInput() *float64
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

// The jsii proxy struct for AppstreamFleetComputeCapacityOutputReference
type jsiiProxy_AppstreamFleetComputeCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) DesiredInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) DesiredInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetComputeCapacityOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamFleetComputeCapacityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetComputeCapacityOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetComputeCapacityOutputReference_Override(a AppstreamFleetComputeCapacityOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetDesiredInstances(val *float64) {
	_jsii_.Set(
		j,
		"desiredInstances",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type AppstreamFleetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// compute_capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#compute_capacity AppstreamFleet#compute_capacity}
	ComputeCapacity *AppstreamFleetComputeCapacity `json:"computeCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#instance_type AppstreamFleet#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#name AppstreamFleet#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#description AppstreamFleet#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#disconnect_timeout_in_seconds AppstreamFleet#disconnect_timeout_in_seconds}.
	DisconnectTimeoutInSeconds *float64 `json:"disconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#display_name AppstreamFleet#display_name}.
	DisplayName *string `json:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#domain_join_info AppstreamFleet#domain_join_info}
	DomainJoinInfo *AppstreamFleetDomainJoinInfo `json:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#enable_default_internet_access AppstreamFleet#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `json:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#fleet_type AppstreamFleet#fleet_type}.
	FleetType *string `json:"fleetType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#iam_role_arn AppstreamFleet#iam_role_arn}.
	IamRoleArn *string `json:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#idle_disconnect_timeout_in_seconds AppstreamFleet#idle_disconnect_timeout_in_seconds}.
	IdleDisconnectTimeoutInSeconds *float64 `json:"idleDisconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#image_arn AppstreamFleet#image_arn}.
	ImageArn *string `json:"imageArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#image_name AppstreamFleet#image_name}.
	ImageName *string `json:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#max_user_duration_in_seconds AppstreamFleet#max_user_duration_in_seconds}.
	MaxUserDurationInSeconds *float64 `json:"maxUserDurationInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#stream_view AppstreamFleet#stream_view}.
	StreamView *string `json:"streamView"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#tags AppstreamFleet#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#tags_all AppstreamFleet#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#vpc_config AppstreamFleet#vpc_config}
	VpcConfig *AppstreamFleetVpcConfig `json:"vpcConfig"`
}

type AppstreamFleetDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#directory_name AppstreamFleet#directory_name}.
	DirectoryName *string `json:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#organizational_unit_distinguished_name AppstreamFleet#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName"`
}

type AppstreamFleetDomainJoinInfoOutputReference interface {
	cdktf.ComplexObject
	DirectoryName() *string
	SetDirectoryName(val *string)
	DirectoryNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
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
	ResetDirectoryName()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for AppstreamFleetDomainJoinInfoOutputReference
type jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) DirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetDomainJoinInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamFleetDomainJoinInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetDomainJoinInfoOutputReference_Override(a AppstreamFleetDomainJoinInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"directoryName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

type AppstreamFleetVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#security_group_ids AppstreamFleet#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet.html#subnet_ids AppstreamFleet#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
}

type AppstreamFleetVpcConfigOutputReference interface {
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
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for AppstreamFleetVpcConfigOutputReference
type jsiiProxy_AppstreamFleetVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetVpcConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamFleetVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetVpcConfigOutputReference_Override(a AppstreamFleetVpcConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamFleetVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html aws_appstream_image_builder}.
type AppstreamImageBuilder interface {
	cdktf.TerraformResource
	AccessEndpoint() *[]*AppstreamImageBuilderAccessEndpoint
	SetAccessEndpoint(val *[]*AppstreamImageBuilderAccessEndpoint)
	AccessEndpointInput() *[]*AppstreamImageBuilderAccessEndpoint
	AppstreamAgentVersion() *string
	SetAppstreamAgentVersion(val *string)
	AppstreamAgentVersionInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference
	DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() AppstreamImageBuilderVpcConfigOutputReference
	VpcConfigInput() *AppstreamImageBuilderVpcConfig
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo)
	PutVpcConfig(value *AppstreamImageBuilderVpcConfig)
	ResetAccessEndpoint()
	ResetAppstreamAgentVersion()
	ResetDescription()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetIamRoleArn()
	ResetImageArn()
	ResetImageName()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamImageBuilder
type jsiiProxy_AppstreamImageBuilder struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpoint() *[]*AppstreamImageBuilderAccessEndpoint {
	var returns *[]*AppstreamImageBuilderAccessEndpoint
	_jsii_.Get(
		j,
		"accessEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpointInput() *[]*AppstreamImageBuilderAccessEndpoint {
	var returns *[]*AppstreamImageBuilderAccessEndpoint
	_jsii_.Get(
		j,
		"accessEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference {
	var returns AppstreamImageBuilderDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo {
	var returns *AppstreamImageBuilderDomainJoinInfo
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfig() AppstreamImageBuilderVpcConfigOutputReference {
	var returns AppstreamImageBuilderVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfigInput() *AppstreamImageBuilderVpcConfig {
	var returns *AppstreamImageBuilderVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder(scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) AppstreamImageBuilder {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilder{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder_Override(a AppstreamImageBuilder, scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetAccessEndpoint(val *[]*AppstreamImageBuilderAccessEndpoint) {
	_jsii_.Set(
		j,
		"accessEndpoint",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetAppstreamAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"appstreamAgentVersion",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetEnableDefaultInternetAccess(val interface{}) {
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetImageArn(val *string) {
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetTagsAll(val interface{}) {
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
func AppstreamImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppStream.AppstreamImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamImageBuilder_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppStream.AppstreamImageBuilder",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo) {
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutVpcConfig(value *AppstreamImageBuilderVpcConfig) {
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAccessEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAppstreamAgentVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAppstreamAgentVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppstreamImageBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamImageBuilderAccessEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#endpoint_type AppstreamImageBuilder#endpoint_type}.
	EndpointType *string `json:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#vpce_id AppstreamImageBuilder#vpce_id}.
	VpceId *string `json:"vpceId"`
}

type AppstreamImageBuilderConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#instance_type AppstreamImageBuilder#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#name AppstreamImageBuilder#name}.
	Name *string `json:"name"`
	// access_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#access_endpoint AppstreamImageBuilder#access_endpoint}
	AccessEndpoint *[]*AppstreamImageBuilderAccessEndpoint `json:"accessEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#appstream_agent_version AppstreamImageBuilder#appstream_agent_version}.
	AppstreamAgentVersion *string `json:"appstreamAgentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#description AppstreamImageBuilder#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#display_name AppstreamImageBuilder#display_name}.
	DisplayName *string `json:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#domain_join_info AppstreamImageBuilder#domain_join_info}
	DomainJoinInfo *AppstreamImageBuilderDomainJoinInfo `json:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#enable_default_internet_access AppstreamImageBuilder#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `json:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#iam_role_arn AppstreamImageBuilder#iam_role_arn}.
	IamRoleArn *string `json:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#image_arn AppstreamImageBuilder#image_arn}.
	ImageArn *string `json:"imageArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#image_name AppstreamImageBuilder#image_name}.
	ImageName *string `json:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#tags AppstreamImageBuilder#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#tags_all AppstreamImageBuilder#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#vpc_config AppstreamImageBuilder#vpc_config}
	VpcConfig *AppstreamImageBuilderVpcConfig `json:"vpcConfig"`
}

type AppstreamImageBuilderDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#directory_name AppstreamImageBuilder#directory_name}.
	DirectoryName *string `json:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#organizational_unit_distinguished_name AppstreamImageBuilder#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName"`
}

type AppstreamImageBuilderDomainJoinInfoOutputReference interface {
	cdktf.ComplexObject
	DirectoryName() *string
	SetDirectoryName(val *string)
	DirectoryNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
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
	ResetDirectoryName()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for AppstreamImageBuilderDomainJoinInfoOutputReference
type jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) DirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamImageBuilderDomainJoinInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamImageBuilderDomainJoinInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilderDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamImageBuilderDomainJoinInfoOutputReference_Override(a AppstreamImageBuilderDomainJoinInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilderDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"directoryName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

type AppstreamImageBuilderVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#security_group_ids AppstreamImageBuilder#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder.html#subnet_ids AppstreamImageBuilder#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
}

type AppstreamImageBuilderVpcConfigOutputReference interface {
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
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for AppstreamImageBuilderVpcConfigOutputReference
type jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamImageBuilderVpcConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamImageBuilderVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilderVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamImageBuilderVpcConfigOutputReference_Override(a AppstreamImageBuilderVpcConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamImageBuilderVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html aws_appstream_stack}.
type AppstreamStack interface {
	cdktf.TerraformResource
	AccessEndpoints() *[]*AppstreamStackAccessEndpoints
	SetAccessEndpoints(val *[]*AppstreamStackAccessEndpoints)
	AccessEndpointsInput() *[]*AppstreamStackAccessEndpoints
	ApplicationSettings() AppstreamStackApplicationSettingsOutputReference
	ApplicationSettingsInput() *AppstreamStackApplicationSettings
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EmbedHostDomains() *[]*string
	SetEmbedHostDomains(val *[]*string)
	EmbedHostDomainsInput() *[]*string
	FeedbackUrl() *string
	SetFeedbackUrl(val *string)
	FeedbackUrlInput() *string
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
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlInput() *string
	StorageConnectors() *[]*AppstreamStackStorageConnectors
	SetStorageConnectors(val *[]*AppstreamStackStorageConnectors)
	StorageConnectorsInput() *[]*AppstreamStackStorageConnectors
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserSettings() *[]*AppstreamStackUserSettings
	SetUserSettings(val *[]*AppstreamStackUserSettings)
	UserSettingsInput() *[]*AppstreamStackUserSettings
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutApplicationSettings(value *AppstreamStackApplicationSettings)
	ResetAccessEndpoints()
	ResetApplicationSettings()
	ResetDescription()
	ResetDisplayName()
	ResetEmbedHostDomains()
	ResetFeedbackUrl()
	ResetOverrideLogicalId()
	ResetRedirectUrl()
	ResetStorageConnectors()
	ResetTags()
	ResetTagsAll()
	ResetUserSettings()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamStack
type jsiiProxy_AppstreamStack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamStack) AccessEndpoints() *[]*AppstreamStackAccessEndpoints {
	var returns *[]*AppstreamStackAccessEndpoints
	_jsii_.Get(
		j,
		"accessEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) AccessEndpointsInput() *[]*AppstreamStackAccessEndpoints {
	var returns *[]*AppstreamStackAccessEndpoints
	_jsii_.Get(
		j,
		"accessEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettings() AppstreamStackApplicationSettingsOutputReference {
	var returns AppstreamStackApplicationSettingsOutputReference
	_jsii_.Get(
		j,
		"applicationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettingsInput() *AppstreamStackApplicationSettings {
	var returns *AppstreamStackApplicationSettings
	_jsii_.Get(
		j,
		"applicationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectors() *[]*AppstreamStackStorageConnectors {
	var returns *[]*AppstreamStackStorageConnectors
	_jsii_.Get(
		j,
		"storageConnectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectorsInput() *[]*AppstreamStackStorageConnectors {
	var returns *[]*AppstreamStackStorageConnectors
	_jsii_.Get(
		j,
		"storageConnectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettings() *[]*AppstreamStackUserSettings {
	var returns *[]*AppstreamStackUserSettings
	_jsii_.Get(
		j,
		"userSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettingsInput() *[]*AppstreamStackUserSettings {
	var returns *[]*AppstreamStackUserSettings
	_jsii_.Get(
		j,
		"userSettingsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html aws_appstream_stack} Resource.
func NewAppstreamStack(scope constructs.Construct, id *string, config *AppstreamStackConfig) AppstreamStack {
	_init_.Initialize()

	j := jsiiProxy_AppstreamStack{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html aws_appstream_stack} Resource.
func NewAppstreamStack_Override(a AppstreamStack, scope constructs.Construct, id *string, config *AppstreamStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamStack",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamStack) SetAccessEndpoints(val *[]*AppstreamStackAccessEndpoints) {
	_jsii_.Set(
		j,
		"accessEndpoints",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetEmbedHostDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"embedHostDomains",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetFeedbackUrl(val *string) {
	_jsii_.Set(
		j,
		"feedbackUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetStorageConnectors(val *[]*AppstreamStackStorageConnectors) {
	_jsii_.Set(
		j,
		"storageConnectors",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetUserSettings(val *[]*AppstreamStackUserSettings) {
	_jsii_.Set(
		j,
		"userSettings",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppstreamStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppStream.AppstreamStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamStack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppStream.AppstreamStack",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppstreamStack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamStack) PutApplicationSettings(value *AppstreamStackApplicationSettings) {
	_jsii_.InvokeVoid(
		a,
		"putApplicationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) ResetAccessEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetApplicationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetEmbedHostDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbedHostDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetFeedbackUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetFeedbackUrl",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamStack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetStorageConnectors() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConnectors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetUserSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUserSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppstreamStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppstreamStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamStackAccessEndpoints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#endpoint_type AppstreamStack#endpoint_type}.
	EndpointType *string `json:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#vpce_id AppstreamStack#vpce_id}.
	VpceId *string `json:"vpceId"`
}

type AppstreamStackApplicationSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#enabled AppstreamStack#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#settings_group AppstreamStack#settings_group}.
	SettingsGroup *string `json:"settingsGroup"`
}

type AppstreamStackApplicationSettingsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SettingsGroup() *string
	SetSettingsGroup(val *string)
	SettingsGroupInput() *string
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
	ResetEnabled()
	ResetSettingsGroup()
}

// The jsii proxy struct for AppstreamStackApplicationSettingsOutputReference
type jsiiProxy_AppstreamStackApplicationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SettingsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SettingsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamStackApplicationSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AppstreamStackApplicationSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamStackApplicationSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamStackApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamStackApplicationSettingsOutputReference_Override(a AppstreamStackApplicationSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppStream.AppstreamStackApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetSettingsGroup(val *string) {
	_jsii_.Set(
		j,
		"settingsGroup",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) ResetSettingsGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetSettingsGroup",
		nil, // no parameters
	)
}

type AppstreamStackConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#name AppstreamStack#name}.
	Name *string `json:"name"`
	// access_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#access_endpoints AppstreamStack#access_endpoints}
	AccessEndpoints *[]*AppstreamStackAccessEndpoints `json:"accessEndpoints"`
	// application_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#application_settings AppstreamStack#application_settings}
	ApplicationSettings *AppstreamStackApplicationSettings `json:"applicationSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#description AppstreamStack#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#display_name AppstreamStack#display_name}.
	DisplayName *string `json:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#embed_host_domains AppstreamStack#embed_host_domains}.
	EmbedHostDomains *[]*string `json:"embedHostDomains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#feedback_url AppstreamStack#feedback_url}.
	FeedbackUrl *string `json:"feedbackUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#redirect_url AppstreamStack#redirect_url}.
	RedirectUrl *string `json:"redirectUrl"`
	// storage_connectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#storage_connectors AppstreamStack#storage_connectors}
	StorageConnectors *[]*AppstreamStackStorageConnectors `json:"storageConnectors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#tags AppstreamStack#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#tags_all AppstreamStack#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// user_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#user_settings AppstreamStack#user_settings}
	UserSettings *[]*AppstreamStackUserSettings `json:"userSettings"`
}

type AppstreamStackStorageConnectors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#connector_type AppstreamStack#connector_type}.
	ConnectorType *string `json:"connectorType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#domains AppstreamStack#domains}.
	Domains *[]*string `json:"domains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#resource_identifier AppstreamStack#resource_identifier}.
	ResourceIdentifier *string `json:"resourceIdentifier"`
}

type AppstreamStackUserSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#action AppstreamStack#action}.
	Action *string `json:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack.html#permission AppstreamStack#permission}.
	Permission *string `json:"permission"`
}
