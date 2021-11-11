package globalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/globalaccelerator/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator}.
type DataAwsGlobalacceleratorAccelerator interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DnsName() *string
	Enabled() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	HostedZoneId() *string
	Id() *string
	IpAddressType() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	Attributes(index *string) DataAwsGlobalacceleratorAcceleratorAttributes
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	IpSets(index *string) DataAwsGlobalacceleratorAcceleratorIpSets
	OverrideLogicalId(newLogicalId *string)
	ResetName()
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsGlobalacceleratorAccelerator
type jsiiProxy_DataAwsGlobalacceleratorAccelerator struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator} Data Source.
func NewDataAwsGlobalacceleratorAccelerator(scope constructs.Construct, id *string, config *DataAwsGlobalacceleratorAcceleratorConfig) DataAwsGlobalacceleratorAccelerator {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlobalacceleratorAccelerator{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAccelerator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator} Data Source.
func NewDataAwsGlobalacceleratorAccelerator_Override(d DataAwsGlobalacceleratorAccelerator, scope constructs.Construct, id *string, config *DataAwsGlobalacceleratorAcceleratorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAccelerator",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsGlobalacceleratorAccelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAccelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsGlobalacceleratorAccelerator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAccelerator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) Attributes(index *string) DataAwsGlobalacceleratorAcceleratorAttributes {
	var returns DataAwsGlobalacceleratorAcceleratorAttributes

	_jsii_.Invoke(
		d,
		"attributes",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) IpSets(index *string) DataAwsGlobalacceleratorAcceleratorIpSets {
	var returns DataAwsGlobalacceleratorAcceleratorIpSets

	_jsii_.Invoke(
		d,
		"ipSets",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ToString() *string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAccelerator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsGlobalacceleratorAcceleratorAttributes interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	FlowLogsEnabled() interface{}
	FlowLogsS3Bucket() *string
	FlowLogsS3Prefix() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsGlobalacceleratorAcceleratorAttributes
type jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) FlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) FlowLogsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) FlowLogsS3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlobalacceleratorAcceleratorAttributes(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlobalacceleratorAcceleratorAttributes {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAcceleratorAttributes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlobalacceleratorAcceleratorAttributes_Override(d DataAwsGlobalacceleratorAcceleratorAttributes, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAcceleratorAttributes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorAttributes) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsGlobalacceleratorAcceleratorConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/globalaccelerator_accelerator.html#name DataAwsGlobalacceleratorAccelerator#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/globalaccelerator_accelerator.html#tags DataAwsGlobalacceleratorAccelerator#tags}.
	Tags interface{} `json:"tags"`
}

type DataAwsGlobalacceleratorAcceleratorIpSets interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	IpAddresses() *[]*string
	IpFamily() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsGlobalacceleratorAcceleratorIpSets
type jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) IpFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlobalacceleratorAcceleratorIpSets(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlobalacceleratorAcceleratorIpSets {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAcceleratorIpSets",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlobalacceleratorAcceleratorIpSets_Override(d DataAwsGlobalacceleratorAcceleratorIpSets, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.DataAwsGlobalacceleratorAcceleratorIpSets",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlobalacceleratorAcceleratorIpSets) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator}.
type GlobalacceleratorAccelerator interface {
	cdktf.TerraformResource
	Attributes() GlobalacceleratorAcceleratorAttributesOutputReference
	AttributesInput() *GlobalacceleratorAcceleratorAttributes
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DnsName() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	HostedZoneId() *string
	Id() *string
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
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
	Timeouts() GlobalacceleratorAcceleratorTimeoutsOutputReference
	TimeoutsInput() *GlobalacceleratorAcceleratorTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	IpSets(index *string) GlobalacceleratorAcceleratorIpSets
	OverrideLogicalId(newLogicalId *string)
	PutAttributes(value *GlobalacceleratorAcceleratorAttributes)
	PutTimeouts(value *GlobalacceleratorAcceleratorTimeouts)
	ResetAttributes()
	ResetEnabled()
	ResetIpAddressType()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlobalacceleratorAccelerator
type jsiiProxy_GlobalacceleratorAccelerator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Attributes() GlobalacceleratorAcceleratorAttributesOutputReference {
	var returns GlobalacceleratorAcceleratorAttributesOutputReference
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) AttributesInput() *GlobalacceleratorAcceleratorAttributes {
	var returns *GlobalacceleratorAcceleratorAttributes
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) Timeouts() GlobalacceleratorAcceleratorTimeoutsOutputReference {
	var returns GlobalacceleratorAcceleratorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) TimeoutsInput() *GlobalacceleratorAcceleratorTimeouts {
	var returns *GlobalacceleratorAcceleratorTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator} Resource.
func NewGlobalacceleratorAccelerator(scope constructs.Construct, id *string, config *GlobalacceleratorAcceleratorConfig) GlobalacceleratorAccelerator {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorAccelerator{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAccelerator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html aws_globalaccelerator_accelerator} Resource.
func NewGlobalacceleratorAccelerator_Override(g GlobalacceleratorAccelerator, scope constructs.Construct, id *string, config *GlobalacceleratorAcceleratorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAccelerator",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAccelerator) SetTagsAll(val interface{}) {
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
func GlobalacceleratorAccelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAccelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalacceleratorAccelerator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAccelerator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) IpSets(index *string) GlobalacceleratorAcceleratorIpSets {
	var returns GlobalacceleratorAcceleratorIpSets

	_jsii_.Invoke(
		g,
		"ipSets",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) PutAttributes(value *GlobalacceleratorAcceleratorAttributes) {
	_jsii_.InvokeVoid(
		g,
		"putAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) PutTimeouts(value *GlobalacceleratorAcceleratorTimeouts) {
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddressType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAccelerator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_GlobalacceleratorAccelerator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorAccelerator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlobalacceleratorAcceleratorAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#flow_logs_enabled GlobalacceleratorAccelerator#flow_logs_enabled}.
	FlowLogsEnabled interface{} `json:"flowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#flow_logs_s3_bucket GlobalacceleratorAccelerator#flow_logs_s3_bucket}.
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#flow_logs_s3_prefix GlobalacceleratorAccelerator#flow_logs_s3_prefix}.
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix"`
}

type GlobalacceleratorAcceleratorAttributesOutputReference interface {
	cdktf.ComplexObject
	FlowLogsEnabled() interface{}
	SetFlowLogsEnabled(val interface{})
	FlowLogsEnabledInput() interface{}
	FlowLogsS3Bucket() *string
	SetFlowLogsS3Bucket(val *string)
	FlowLogsS3BucketInput() *string
	FlowLogsS3Prefix() *string
	SetFlowLogsS3Prefix(val *string)
	FlowLogsS3PrefixInput() *string
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
	ResetFlowLogsEnabled()
	ResetFlowLogsS3Bucket()
	ResetFlowLogsS3Prefix()
}

// The jsii proxy struct for GlobalacceleratorAcceleratorAttributesOutputReference
type jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlobalacceleratorAcceleratorAttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlobalacceleratorAcceleratorAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlobalacceleratorAcceleratorAttributesOutputReference_Override(g GlobalacceleratorAcceleratorAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetFlowLogsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"flowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetFlowLogsS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"flowLogsS3Bucket",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetFlowLogsS3Prefix(val *string) {
	_jsii_.Set(
		j,
		"flowLogsS3Prefix",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsS3Bucket() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Bucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsS3Prefix() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Prefix",
		nil, // no parameters
	)
}

type GlobalacceleratorAcceleratorConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#name GlobalacceleratorAccelerator#name}.
	Name *string `json:"name"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#attributes GlobalacceleratorAccelerator#attributes}
	Attributes *GlobalacceleratorAcceleratorAttributes `json:"attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#enabled GlobalacceleratorAccelerator#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#ip_address_type GlobalacceleratorAccelerator#ip_address_type}.
	IpAddressType *string `json:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#tags GlobalacceleratorAccelerator#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#tags_all GlobalacceleratorAccelerator#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#timeouts GlobalacceleratorAccelerator#timeouts}
	Timeouts *GlobalacceleratorAcceleratorTimeouts `json:"timeouts"`
}

type GlobalacceleratorAcceleratorIpSets interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	IpAddresses() *[]*string
	IpFamily() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for GlobalacceleratorAcceleratorIpSets
type jsiiProxy_GlobalacceleratorAcceleratorIpSets struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) IpFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewGlobalacceleratorAcceleratorIpSets(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) GlobalacceleratorAcceleratorIpSets {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorAcceleratorIpSets{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorIpSets",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewGlobalacceleratorAcceleratorIpSets_Override(g GlobalacceleratorAcceleratorIpSets, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorIpSets",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorIpSets) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorIpSets) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorIpSets) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorIpSets) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorIpSets) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorIpSets) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type GlobalacceleratorAcceleratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#create GlobalacceleratorAccelerator#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_accelerator.html#update GlobalacceleratorAccelerator#update}.
	Update *string `json:"update"`
}

type GlobalacceleratorAcceleratorTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
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
	ResetUpdate()
}

// The jsii proxy struct for GlobalacceleratorAcceleratorTimeoutsOutputReference
type jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewGlobalacceleratorAcceleratorTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlobalacceleratorAcceleratorTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlobalacceleratorAcceleratorTimeoutsOutputReference_Override(g GlobalacceleratorAcceleratorTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorAcceleratorTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html aws_globalaccelerator_endpoint_group}.
type GlobalacceleratorEndpointGroup interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EndpointConfiguration() *[]*GlobalacceleratorEndpointGroupEndpointConfiguration
	SetEndpointConfiguration(val *[]*GlobalacceleratorEndpointGroupEndpointConfiguration)
	EndpointConfigurationInput() *[]*GlobalacceleratorEndpointGroupEndpointConfiguration
	EndpointGroupRegion() *string
	SetEndpointGroupRegion(val *string)
	EndpointGroupRegionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckIntervalSecondsInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckPortInput() *float64
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckProtocolInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenerArn() *string
	SetListenerArn(val *string)
	ListenerArnInput() *string
	Node() constructs.Node
	PortOverride() *[]*GlobalacceleratorEndpointGroupPortOverride
	SetPortOverride(val *[]*GlobalacceleratorEndpointGroupPortOverride)
	PortOverrideInput() *[]*GlobalacceleratorEndpointGroupPortOverride
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ThresholdCount() *float64
	SetThresholdCount(val *float64)
	ThresholdCountInput() *float64
	Timeouts() GlobalacceleratorEndpointGroupTimeoutsOutputReference
	TimeoutsInput() *GlobalacceleratorEndpointGroupTimeouts
	TrafficDialPercentage() *float64
	SetTrafficDialPercentage(val *float64)
	TrafficDialPercentageInput() *float64
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *GlobalacceleratorEndpointGroupTimeouts)
	ResetEndpointConfiguration()
	ResetEndpointGroupRegion()
	ResetHealthCheckIntervalSeconds()
	ResetHealthCheckPath()
	ResetHealthCheckPort()
	ResetHealthCheckProtocol()
	ResetOverrideLogicalId()
	ResetPortOverride()
	ResetThresholdCount()
	ResetTimeouts()
	ResetTrafficDialPercentage()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlobalacceleratorEndpointGroup
type jsiiProxy_GlobalacceleratorEndpointGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfiguration() *[]*GlobalacceleratorEndpointGroupEndpointConfiguration {
	var returns *[]*GlobalacceleratorEndpointGroupEndpointConfiguration
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfigurationInput() *[]*GlobalacceleratorEndpointGroupEndpointConfiguration {
	var returns *[]*GlobalacceleratorEndpointGroupEndpointConfiguration
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverride() *[]*GlobalacceleratorEndpointGroupPortOverride {
	var returns *[]*GlobalacceleratorEndpointGroupPortOverride
	_jsii_.Get(
		j,
		"portOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverrideInput() *[]*GlobalacceleratorEndpointGroupPortOverride {
	var returns *[]*GlobalacceleratorEndpointGroupPortOverride
	_jsii_.Get(
		j,
		"portOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Timeouts() GlobalacceleratorEndpointGroupTimeoutsOutputReference {
	var returns GlobalacceleratorEndpointGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TimeoutsInput() *GlobalacceleratorEndpointGroupTimeouts {
	var returns *GlobalacceleratorEndpointGroupTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentageInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html aws_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup(scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) GlobalacceleratorEndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorEndpointGroup{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html aws_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup_Override(g GlobalacceleratorEndpointGroup, scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetEndpointConfiguration(val *[]*GlobalacceleratorEndpointGroupEndpointConfiguration) {
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetEndpointGroupRegion(val *string) {
	_jsii_.Set(
		j,
		"endpointGroupRegion",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetHealthCheckIntervalSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetHealthCheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetHealthCheckProtocol(val *string) {
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetListenerArn(val *string) {
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetPortOverride(val *[]*GlobalacceleratorEndpointGroupPortOverride) {
	_jsii_.Set(
		j,
		"portOverride",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetThresholdCount(val *float64) {
	_jsii_.Set(
		j,
		"thresholdCount",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) SetTrafficDialPercentage(val *float64) {
	_jsii_.Set(
		j,
		"trafficDialPercentage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlobalacceleratorEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalacceleratorEndpointGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutTimeouts(value *GlobalacceleratorEndpointGroupTimeouts) {
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetEndpointGroupRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointGroupRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPort() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckProtocol",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetPortOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetPortOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetThresholdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetTrafficDialPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetTrafficDialPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlobalacceleratorEndpointGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#listener_arn GlobalacceleratorEndpointGroup#listener_arn}.
	ListenerArn *string `json:"listenerArn"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#endpoint_configuration GlobalacceleratorEndpointGroup#endpoint_configuration}
	EndpointConfiguration *[]*GlobalacceleratorEndpointGroupEndpointConfiguration `json:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#endpoint_group_region GlobalacceleratorEndpointGroup#endpoint_group_region}.
	EndpointGroupRegion *string `json:"endpointGroupRegion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#health_check_interval_seconds GlobalacceleratorEndpointGroup#health_check_interval_seconds}.
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#health_check_path GlobalacceleratorEndpointGroup#health_check_path}.
	HealthCheckPath *string `json:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#health_check_port GlobalacceleratorEndpointGroup#health_check_port}.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#health_check_protocol GlobalacceleratorEndpointGroup#health_check_protocol}.
	HealthCheckProtocol *string `json:"healthCheckProtocol"`
	// port_override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#port_override GlobalacceleratorEndpointGroup#port_override}
	PortOverride *[]*GlobalacceleratorEndpointGroupPortOverride `json:"portOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#threshold_count GlobalacceleratorEndpointGroup#threshold_count}.
	ThresholdCount *float64 `json:"thresholdCount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#timeouts GlobalacceleratorEndpointGroup#timeouts}
	Timeouts *GlobalacceleratorEndpointGroupTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#traffic_dial_percentage GlobalacceleratorEndpointGroup#traffic_dial_percentage}.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

type GlobalacceleratorEndpointGroupEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#client_ip_preservation_enabled GlobalacceleratorEndpointGroup#client_ip_preservation_enabled}.
	ClientIpPreservationEnabled interface{} `json:"clientIpPreservationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#endpoint_id GlobalacceleratorEndpointGroup#endpoint_id}.
	EndpointId *string `json:"endpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#weight GlobalacceleratorEndpointGroup#weight}.
	Weight *float64 `json:"weight"`
}

type GlobalacceleratorEndpointGroupPortOverride struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#endpoint_port GlobalacceleratorEndpointGroup#endpoint_port}.
	EndpointPort *float64 `json:"endpointPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#listener_port GlobalacceleratorEndpointGroup#listener_port}.
	ListenerPort *float64 `json:"listenerPort"`
}

type GlobalacceleratorEndpointGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#create GlobalacceleratorEndpointGroup#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#delete GlobalacceleratorEndpointGroup#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group.html#update GlobalacceleratorEndpointGroup#update}.
	Update *string `json:"update"`
}

type GlobalacceleratorEndpointGroupTimeoutsOutputReference interface {
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

// The jsii proxy struct for GlobalacceleratorEndpointGroupTimeoutsOutputReference
type jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewGlobalacceleratorEndpointGroupTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlobalacceleratorEndpointGroupTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlobalacceleratorEndpointGroupTimeoutsOutputReference_Override(g GlobalacceleratorEndpointGroupTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorEndpointGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroupTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html aws_globalaccelerator_listener}.
type GlobalacceleratorListener interface {
	cdktf.TerraformResource
	AcceleratorArn() *string
	SetAcceleratorArn(val *string)
	AcceleratorArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ClientAffinity() *string
	SetClientAffinity(val *string)
	ClientAffinityInput() *string
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
	PortRange() *[]*GlobalacceleratorListenerPortRange
	SetPortRange(val *[]*GlobalacceleratorListenerPortRange)
	PortRangeInput() *[]*GlobalacceleratorListenerPortRange
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() GlobalacceleratorListenerTimeoutsOutputReference
	TimeoutsInput() *GlobalacceleratorListenerTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *GlobalacceleratorListenerTimeouts)
	ResetClientAffinity()
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlobalacceleratorListener
type jsiiProxy_GlobalacceleratorListener struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlobalacceleratorListener) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) AcceleratorArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) ClientAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) ClientAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) PortRange() *[]*GlobalacceleratorListenerPortRange {
	var returns *[]*GlobalacceleratorListenerPortRange
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) PortRangeInput() *[]*GlobalacceleratorListenerPortRange {
	var returns *[]*GlobalacceleratorListenerPortRange
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) Timeouts() GlobalacceleratorListenerTimeoutsOutputReference {
	var returns GlobalacceleratorListenerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListener) TimeoutsInput() *GlobalacceleratorListenerTimeouts {
	var returns *GlobalacceleratorListenerTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html aws_globalaccelerator_listener} Resource.
func NewGlobalacceleratorListener(scope constructs.Construct, id *string, config *GlobalacceleratorListenerConfig) GlobalacceleratorListener {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorListener{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html aws_globalaccelerator_listener} Resource.
func NewGlobalacceleratorListener_Override(g GlobalacceleratorListener, scope constructs.Construct, id *string, config *GlobalacceleratorListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListener",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetAcceleratorArn(val *string) {
	_jsii_.Set(
		j,
		"acceleratorArn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetClientAffinity(val *string) {
	_jsii_.Set(
		j,
		"clientAffinity",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetPortRange(val *[]*GlobalacceleratorListenerPortRange) {
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListener) SetProvider(val cdktf.TerraformProvider) {
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
func GlobalacceleratorListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalacceleratorListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalacceleratorListener) PutTimeouts(value *GlobalacceleratorListenerTimeouts) {
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorListener) ResetClientAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetClientAffinity",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorListener) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_GlobalacceleratorListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (g *jsiiProxy_GlobalacceleratorListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlobalacceleratorListenerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#accelerator_arn GlobalacceleratorListener#accelerator_arn}.
	AcceleratorArn *string `json:"acceleratorArn"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#port_range GlobalacceleratorListener#port_range}
	PortRange *[]*GlobalacceleratorListenerPortRange `json:"portRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#protocol GlobalacceleratorListener#protocol}.
	Protocol *string `json:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#client_affinity GlobalacceleratorListener#client_affinity}.
	ClientAffinity *string `json:"clientAffinity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#timeouts GlobalacceleratorListener#timeouts}
	Timeouts *GlobalacceleratorListenerTimeouts `json:"timeouts"`
}

type GlobalacceleratorListenerPortRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#from_port GlobalacceleratorListener#from_port}.
	FromPort *float64 `json:"fromPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#to_port GlobalacceleratorListener#to_port}.
	ToPort *float64 `json:"toPort"`
}

type GlobalacceleratorListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#create GlobalacceleratorListener#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#delete GlobalacceleratorListener#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_listener.html#update GlobalacceleratorListener#update}.
	Update *string `json:"update"`
}

type GlobalacceleratorListenerTimeoutsOutputReference interface {
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

// The jsii proxy struct for GlobalacceleratorListenerTimeoutsOutputReference
type jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewGlobalacceleratorListenerTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlobalacceleratorListenerTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListenerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlobalacceleratorListenerTimeoutsOutputReference_Override(g GlobalacceleratorListenerTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.GlobalAccelerator.GlobalacceleratorListenerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorListenerTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdate",
		nil, // no parameters
	)
}
