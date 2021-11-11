package mq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/mq/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html aws_mq_broker}.
type DataAwsMqBroker interface {
	cdktf.TerraformDataSource
	Arn() *string
	AuthenticationStrategy() *string
	AutoMinorVersionUpgrade() interface{}
	BrokerId() *string
	SetBrokerId(val *string)
	BrokerIdInput() *string
	BrokerName() *string
	SetBrokerName(val *string)
	BrokerNameInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentMode() *string
	EngineType() *string
	EngineVersion() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HostInstanceType() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PubliclyAccessible() interface{}
	RawOverrides() interface{}
	SecurityGroups() *[]*string
	StorageType() *string
	SubnetIds() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	Configuration(index *string) DataAwsMqBrokerConfiguration
	EncryptionOptions(index *string) DataAwsMqBrokerEncryptionOptions
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	Instances(index *string) DataAwsMqBrokerInstances
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	LdapServerMetadata(index *string) DataAwsMqBrokerLdapServerMetadata
	Logs(index *string) DataAwsMqBrokerLogs
	MaintenanceWindowStartTime(index *string) DataAwsMqBrokerMaintenanceWindowStartTime
	OverrideLogicalId(newLogicalId *string)
	ResetBrokerId()
	ResetBrokerName()
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	User(index *string) DataAwsMqBrokerUser
}

// The jsii proxy struct for DataAwsMqBroker
type jsiiProxy_DataAwsMqBroker struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsMqBroker) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) BrokerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) BrokerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) BrokerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) BrokerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBroker) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html aws_mq_broker} Data Source.
func NewDataAwsMqBroker(scope constructs.Construct, id *string, config *DataAwsMqBrokerConfig) DataAwsMqBroker {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBroker{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBroker",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html aws_mq_broker} Data Source.
func NewDataAwsMqBroker_Override(d DataAwsMqBroker, scope constructs.Construct, id *string, config *DataAwsMqBrokerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBroker",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetBrokerId(val *string) {
	_jsii_.Set(
		j,
		"brokerId",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetBrokerName(val *string) {
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBroker) SetTags(val interface{}) {
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
func DataAwsMqBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.MQ.DataAwsMqBroker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsMqBroker_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.MQ.DataAwsMqBroker",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBroker) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsMqBroker) Configuration(index *string) DataAwsMqBrokerConfiguration {
	var returns DataAwsMqBrokerConfiguration

	_jsii_.Invoke(
		d,
		"configuration",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) EncryptionOptions(index *string) DataAwsMqBrokerEncryptionOptions {
	var returns DataAwsMqBrokerEncryptionOptions

	_jsii_.Invoke(
		d,
		"encryptionOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBroker) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsMqBroker) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBroker) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBroker) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) Instances(index *string) DataAwsMqBrokerInstances {
	var returns DataAwsMqBrokerInstances

	_jsii_.Invoke(
		d,
		"instances",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBroker) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) LdapServerMetadata(index *string) DataAwsMqBrokerLdapServerMetadata {
	var returns DataAwsMqBrokerLdapServerMetadata

	_jsii_.Invoke(
		d,
		"ldapServerMetadata",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) Logs(index *string) DataAwsMqBrokerLogs {
	var returns DataAwsMqBrokerLogs

	_jsii_.Invoke(
		d,
		"logs",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) MaintenanceWindowStartTime(index *string) DataAwsMqBrokerMaintenanceWindowStartTime {
	var returns DataAwsMqBrokerMaintenanceWindowStartTime

	_jsii_.Invoke(
		d,
		"maintenanceWindowStartTime",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsMqBroker) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsMqBroker) ResetBrokerId() {
	_jsii_.InvokeVoid(
		d,
		"resetBrokerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBroker) ResetBrokerName() {
	_jsii_.InvokeVoid(
		d,
		"resetBrokerName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsMqBroker) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBroker) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBroker) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsMqBroker) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsMqBroker) ToString() *string {
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
func (d *jsiiProxy_DataAwsMqBroker) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBroker) User(index *string) DataAwsMqBrokerUser {
	var returns DataAwsMqBrokerUser

	_jsii_.Invoke(
		d,
		"user",
		[]interface{}{index},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html#broker_id DataAwsMqBroker#broker_id}.
	BrokerId *string `json:"brokerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html#broker_name DataAwsMqBroker#broker_name}.
	BrokerName *string `json:"brokerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/mq_broker.html#tags DataAwsMqBroker#tags}.
	Tags interface{} `json:"tags"`
}

type DataAwsMqBrokerConfiguration interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Id() *string
	Revision() *float64
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

// The jsii proxy struct for DataAwsMqBrokerConfiguration
type jsiiProxy_DataAwsMqBrokerConfiguration struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerConfiguration(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerConfiguration {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerConfiguration_Override(d DataAwsMqBrokerConfiguration, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerConfiguration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerConfiguration) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerConfiguration) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerConfiguration) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerEncryptionOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	KmsKeyId() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UseAwsOwnedKey() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsMqBrokerEncryptionOptions
type jsiiProxy_DataAwsMqBrokerEncryptionOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) UseAwsOwnedKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAwsOwnedKey",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerEncryptionOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerEncryptionOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerEncryptionOptions{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerEncryptionOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerEncryptionOptions_Override(d DataAwsMqBrokerEncryptionOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerEncryptionOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerEncryptionOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerEncryptionOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerEncryptionOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerEncryptionOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerEncryptionOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerEncryptionOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerInstances interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConsoleUrl() *string
	Endpoints() *[]*string
	IpAddress() *string
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

// The jsii proxy struct for DataAwsMqBrokerInstances
type jsiiProxy_DataAwsMqBrokerInstances struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) ConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) Endpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerInstances(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerInstances {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerInstances{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerInstances",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerInstances_Override(d DataAwsMqBrokerInstances, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerInstances",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstances) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerInstances) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerInstances) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerInstances) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerInstances) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerInstances) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerLdapServerMetadata interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Hosts() *[]*string
	RoleBase() *string
	RoleName() *string
	RoleSearchMatching() *string
	RoleSearchSubtree() interface{}
	ServiceAccountPassword() *string
	ServiceAccountUsername() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserBase() *string
	UserRoleName() *string
	UserSearchMatching() *string
	UserSearchSubtree() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsMqBrokerLdapServerMetadata
type jsiiProxy_DataAwsMqBrokerLdapServerMetadata struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) RoleBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) RoleSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) RoleSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) ServiceAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) ServiceAccountUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) UserBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) UserRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) UserSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) UserSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtree",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerLdapServerMetadata(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerLdapServerMetadata {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerLdapServerMetadata{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerLdapServerMetadata",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerLdapServerMetadata_Override(d DataAwsMqBrokerLdapServerMetadata, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerLdapServerMetadata",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadata) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerLogs interface {
	cdktf.ComplexComputedList
	Audit() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	General() interface{}
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

// The jsii proxy struct for DataAwsMqBrokerLogs
type jsiiProxy_DataAwsMqBrokerLogs struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) Audit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) General() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"general",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerLogs(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerLogs {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerLogs{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerLogs",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerLogs_Override(d DataAwsMqBrokerLogs, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerLogs",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerLogs) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerLogs) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerLogs) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerLogs) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerLogs) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerLogs) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerMaintenanceWindowStartTime interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DayOfWeek() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TimeOfDay() *string
	TimeZone() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsMqBrokerMaintenanceWindowStartTime
type jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) TimeOfDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerMaintenanceWindowStartTime(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerMaintenanceWindowStartTime {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerMaintenanceWindowStartTime",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerMaintenanceWindowStartTime_Override(d DataAwsMqBrokerMaintenanceWindowStartTime, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerMaintenanceWindowStartTime",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsMqBrokerUser interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConsoleAccess() interface{}
	Groups() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Username() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsMqBrokerUser
type jsiiProxy_DataAwsMqBrokerUser struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsMqBrokerUser) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerUser) ConsoleAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consoleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerUser) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerUser) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerUser) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsMqBrokerUser(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsMqBrokerUser {
	_init_.Initialize()

	j := jsiiProxy_DataAwsMqBrokerUser{}

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerUser",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsMqBrokerUser_Override(d DataAwsMqBrokerUser, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.DataAwsMqBrokerUser",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerUser) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerUser) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerUser) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsMqBrokerUser) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsMqBrokerUser) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsMqBrokerUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsMqBrokerUser) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsMqBrokerUser) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html aws_mq_broker}.
type MqBroker interface {
	cdktf.TerraformResource
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	Arn() *string
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	AuthenticationStrategyInput() *string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	BrokerName() *string
	SetBrokerName(val *string)
	BrokerNameInput() *string
	CdktfStack() cdktf.TerraformStack
	Configuration() MqBrokerConfigurationOutputReference
	ConfigurationInput() *MqBrokerConfiguration
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	DeploymentModeInput() *string
	EncryptionOptions() MqBrokerEncryptionOptionsOutputReference
	EncryptionOptionsInput() *MqBrokerEncryptionOptions
	EngineType() *string
	SetEngineType(val *string)
	EngineTypeInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HostInstanceType() *string
	SetHostInstanceType(val *string)
	HostInstanceTypeInput() *string
	Id() *string
	LdapServerMetadata() MqBrokerLdapServerMetadataOutputReference
	LdapServerMetadataInput() *MqBrokerLdapServerMetadata
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logs() MqBrokerLogsOutputReference
	LogsInput() *MqBrokerLogs
	MaintenanceWindowStartTime() MqBrokerMaintenanceWindowStartTimeOutputReference
	MaintenanceWindowStartTimeInput() *MqBrokerMaintenanceWindowStartTime
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	RawOverrides() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	User() *[]*MqBrokerUser
	SetUser(val *[]*MqBrokerUser)
	UserInput() *[]*MqBrokerUser
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	Instances(index *string) MqBrokerInstances
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutConfiguration(value *MqBrokerConfiguration)
	PutEncryptionOptions(value *MqBrokerEncryptionOptions)
	PutLdapServerMetadata(value *MqBrokerLdapServerMetadata)
	PutLogs(value *MqBrokerLogs)
	PutMaintenanceWindowStartTime(value *MqBrokerMaintenanceWindowStartTime)
	ResetApplyImmediately()
	ResetAuthenticationStrategy()
	ResetAutoMinorVersionUpgrade()
	ResetConfiguration()
	ResetDeploymentMode()
	ResetEncryptionOptions()
	ResetLdapServerMetadata()
	ResetLogs()
	ResetMaintenanceWindowStartTime()
	ResetOverrideLogicalId()
	ResetPubliclyAccessible()
	ResetSecurityGroups()
	ResetStorageType()
	ResetSubnetIds()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MqBroker
type jsiiProxy_MqBroker struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MqBroker) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) AuthenticationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) BrokerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) BrokerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Configuration() MqBrokerConfigurationOutputReference {
	var returns MqBrokerConfigurationOutputReference
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) ConfigurationInput() *MqBrokerConfiguration {
	var returns *MqBrokerConfiguration
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) DeploymentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EncryptionOptions() MqBrokerEncryptionOptionsOutputReference {
	var returns MqBrokerEncryptionOptionsOutputReference
	_jsii_.Get(
		j,
		"encryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EncryptionOptionsInput() *MqBrokerEncryptionOptions {
	var returns *MqBrokerEncryptionOptions
	_jsii_.Get(
		j,
		"encryptionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) HostInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) LdapServerMetadata() MqBrokerLdapServerMetadataOutputReference {
	var returns MqBrokerLdapServerMetadataOutputReference
	_jsii_.Get(
		j,
		"ldapServerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) LdapServerMetadataInput() *MqBrokerLdapServerMetadata {
	var returns *MqBrokerLdapServerMetadata
	_jsii_.Get(
		j,
		"ldapServerMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Logs() MqBrokerLogsOutputReference {
	var returns MqBrokerLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) LogsInput() *MqBrokerLogs {
	var returns *MqBrokerLogs
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) MaintenanceWindowStartTime() MqBrokerMaintenanceWindowStartTimeOutputReference {
	var returns MqBrokerMaintenanceWindowStartTimeOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) MaintenanceWindowStartTimeInput() *MqBrokerMaintenanceWindowStartTime {
	var returns *MqBrokerMaintenanceWindowStartTime
	_jsii_.Get(
		j,
		"maintenanceWindowStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) User() *[]*MqBrokerUser {
	var returns *[]*MqBrokerUser
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBroker) UserInput() *[]*MqBrokerUser {
	var returns *[]*MqBrokerUser
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html aws_mq_broker} Resource.
func NewMqBroker(scope constructs.Construct, id *string, config *MqBrokerConfig) MqBroker {
	_init_.Initialize()

	j := jsiiProxy_MqBroker{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBroker",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html aws_mq_broker} Resource.
func NewMqBroker_Override(m MqBroker, scope constructs.Construct, id *string, config *MqBrokerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBroker",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MqBroker) SetApplyImmediately(val interface{}) {
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetAuthenticationStrategy(val *string) {
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetBrokerName(val *string) {
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetDeploymentMode(val *string) {
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetEngineType(val *string) {
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetHostInstanceType(val *string) {
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_MqBroker) SetUser(val *[]*MqBrokerUser) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func MqBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.MQ.MqBroker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MqBroker_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.MQ.MqBroker",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MqBroker) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBroker) Instances(index *string) MqBrokerInstances {
	var returns MqBrokerInstances

	_jsii_.Invoke(
		m,
		"instances",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MqBroker) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MqBroker) PutConfiguration(value *MqBrokerConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MqBroker) PutEncryptionOptions(value *MqBrokerEncryptionOptions) {
	_jsii_.InvokeVoid(
		m,
		"putEncryptionOptions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MqBroker) PutLdapServerMetadata(value *MqBrokerLdapServerMetadata) {
	_jsii_.InvokeVoid(
		m,
		"putLdapServerMetadata",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MqBroker) PutLogs(value *MqBrokerLogs) {
	_jsii_.InvokeVoid(
		m,
		"putLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MqBroker) PutMaintenanceWindowStartTime(value *MqBrokerMaintenanceWindowStartTime) {
	_jsii_.InvokeVoid(
		m,
		"putMaintenanceWindowStartTime",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MqBroker) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		m,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetAuthenticationStrategy() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthenticationStrategy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetDeploymentMode() {
	_jsii_.InvokeVoid(
		m,
		"resetDeploymentMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetEncryptionOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetLdapServerMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetLdapServerMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetMaintenanceWindowStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowStartTime",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MqBroker) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		m,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetStorageType() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBroker) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBroker) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MqBroker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MqBroker) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MqBrokerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#broker_name MqBroker#broker_name}.
	BrokerName *string `json:"brokerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#engine_type MqBroker#engine_type}.
	EngineType *string `json:"engineType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#engine_version MqBroker#engine_version}.
	EngineVersion *string `json:"engineVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#host_instance_type MqBroker#host_instance_type}.
	HostInstanceType *string `json:"hostInstanceType"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#user MqBroker#user}
	User *[]*MqBrokerUser `json:"user"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#apply_immediately MqBroker#apply_immediately}.
	ApplyImmediately interface{} `json:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#authentication_strategy MqBroker#authentication_strategy}.
	AuthenticationStrategy *string `json:"authenticationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#auto_minor_version_upgrade MqBroker#auto_minor_version_upgrade}.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#configuration MqBroker#configuration}
	Configuration *MqBrokerConfiguration `json:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#deployment_mode MqBroker#deployment_mode}.
	DeploymentMode *string `json:"deploymentMode"`
	// encryption_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#encryption_options MqBroker#encryption_options}
	EncryptionOptions *MqBrokerEncryptionOptions `json:"encryptionOptions"`
	// ldap_server_metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#ldap_server_metadata MqBroker#ldap_server_metadata}
	LdapServerMetadata *MqBrokerLdapServerMetadata `json:"ldapServerMetadata"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#logs MqBroker#logs}
	Logs *MqBrokerLogs `json:"logs"`
	// maintenance_window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#maintenance_window_start_time MqBroker#maintenance_window_start_time}
	MaintenanceWindowStartTime *MqBrokerMaintenanceWindowStartTime `json:"maintenanceWindowStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#publicly_accessible MqBroker#publicly_accessible}.
	PubliclyAccessible interface{} `json:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#security_groups MqBroker#security_groups}.
	SecurityGroups *[]*string `json:"securityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#storage_type MqBroker#storage_type}.
	StorageType *string `json:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#subnet_ids MqBroker#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#tags MqBroker#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#tags_all MqBroker#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type MqBrokerConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#id MqBroker#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#revision MqBroker#revision}.
	Revision *float64 `json:"revision"`
}

type MqBrokerConfigurationOutputReference interface {
	cdktf.ComplexObject
	Id() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Revision() *float64
	SetRevision(val *float64)
	RevisionInput() *float64
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
	ResetRevision()
}

// The jsii proxy struct for MqBrokerConfigurationOutputReference
type jsiiProxy_MqBrokerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) RevisionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMqBrokerConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MqBrokerConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMqBrokerConfigurationOutputReference_Override(m MqBrokerConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) SetRevision(val *float64) {
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerConfigurationOutputReference) ResetRevision() {
	_jsii_.InvokeVoid(
		m,
		"resetRevision",
		nil, // no parameters
	)
}

type MqBrokerEncryptionOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#kms_key_id MqBroker#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#use_aws_owned_key MqBroker#use_aws_owned_key}.
	UseAwsOwnedKey interface{} `json:"useAwsOwnedKey"`
}

type MqBrokerEncryptionOptionsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UseAwsOwnedKey() interface{}
	SetUseAwsOwnedKey(val interface{})
	UseAwsOwnedKeyInput() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetKmsKeyId()
	ResetUseAwsOwnedKey()
}

// The jsii proxy struct for MqBrokerEncryptionOptionsOutputReference
type jsiiProxy_MqBrokerEncryptionOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) UseAwsOwnedKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAwsOwnedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) UseAwsOwnedKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAwsOwnedKeyInput",
		&returns,
	)
	return returns
}

func NewMqBrokerEncryptionOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MqBrokerEncryptionOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerEncryptionOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerEncryptionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMqBrokerEncryptionOptionsOutputReference_Override(m MqBrokerEncryptionOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerEncryptionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) SetUseAwsOwnedKey(val interface{}) {
	_jsii_.Set(
		j,
		"useAwsOwnedKey",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerEncryptionOptionsOutputReference) ResetUseAwsOwnedKey() {
	_jsii_.InvokeVoid(
		m,
		"resetUseAwsOwnedKey",
		nil, // no parameters
	)
}

type MqBrokerInstances interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConsoleUrl() *string
	Endpoints() *[]*string
	IpAddress() *string
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

// The jsii proxy struct for MqBrokerInstances
type jsiiProxy_MqBrokerInstances struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_MqBrokerInstances) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerInstances) ConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerInstances) Endpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerInstances) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerInstances) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerInstances) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewMqBrokerInstances(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) MqBrokerInstances {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerInstances{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerInstances",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewMqBrokerInstances_Override(m MqBrokerInstances, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerInstances",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		m,
	)
}

func (j *jsiiProxy_MqBrokerInstances) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_MqBrokerInstances) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerInstances) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerInstances) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerInstances) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerInstances) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerInstances) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerInstances) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type MqBrokerLdapServerMetadata struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#hosts MqBroker#hosts}.
	Hosts *[]*string `json:"hosts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#role_base MqBroker#role_base}.
	RoleBase *string `json:"roleBase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#role_name MqBroker#role_name}.
	RoleName *string `json:"roleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#role_search_matching MqBroker#role_search_matching}.
	RoleSearchMatching *string `json:"roleSearchMatching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#role_search_subtree MqBroker#role_search_subtree}.
	RoleSearchSubtree interface{} `json:"roleSearchSubtree"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#service_account_password MqBroker#service_account_password}.
	ServiceAccountPassword *string `json:"serviceAccountPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#service_account_username MqBroker#service_account_username}.
	ServiceAccountUsername *string `json:"serviceAccountUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#user_base MqBroker#user_base}.
	UserBase *string `json:"userBase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#user_role_name MqBroker#user_role_name}.
	UserRoleName *string `json:"userRoleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#user_search_matching MqBroker#user_search_matching}.
	UserSearchMatching *string `json:"userSearchMatching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#user_search_subtree MqBroker#user_search_subtree}.
	UserSearchSubtree interface{} `json:"userSearchSubtree"`
}

type MqBrokerLdapServerMetadataOutputReference interface {
	cdktf.ComplexObject
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleBase() *string
	SetRoleBase(val *string)
	RoleBaseInput() *string
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	RoleSearchMatching() *string
	SetRoleSearchMatching(val *string)
	RoleSearchMatchingInput() *string
	RoleSearchSubtree() interface{}
	SetRoleSearchSubtree(val interface{})
	RoleSearchSubtreeInput() interface{}
	ServiceAccountPassword() *string
	SetServiceAccountPassword(val *string)
	ServiceAccountPasswordInput() *string
	ServiceAccountUsername() *string
	SetServiceAccountUsername(val *string)
	ServiceAccountUsernameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserBase() *string
	SetUserBase(val *string)
	UserBaseInput() *string
	UserRoleName() *string
	SetUserRoleName(val *string)
	UserRoleNameInput() *string
	UserSearchMatching() *string
	SetUserSearchMatching(val *string)
	UserSearchMatchingInput() *string
	UserSearchSubtree() interface{}
	SetUserSearchSubtree(val interface{})
	UserSearchSubtreeInput() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetHosts()
	ResetRoleBase()
	ResetRoleName()
	ResetRoleSearchMatching()
	ResetRoleSearchSubtree()
	ResetServiceAccountPassword()
	ResetServiceAccountUsername()
	ResetUserBase()
	ResetUserRoleName()
	ResetUserSearchMatching()
	ResetUserSearchSubtree()
}

// The jsii proxy struct for MqBrokerLdapServerMetadataOutputReference
type jsiiProxy_MqBrokerLdapServerMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtreeInput",
		&returns,
	)
	return returns
}

func NewMqBrokerLdapServerMetadataOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MqBrokerLdapServerMetadataOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerLdapServerMetadataOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerLdapServerMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMqBrokerLdapServerMetadataOutputReference_Override(m MqBrokerLdapServerMetadataOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerLdapServerMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetHosts(val *[]*string) {
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetRoleBase(val *string) {
	_jsii_.Set(
		j,
		"roleBase",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetRoleName(val *string) {
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetRoleSearchMatching(val *string) {
	_jsii_.Set(
		j,
		"roleSearchMatching",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetRoleSearchSubtree(val interface{}) {
	_jsii_.Set(
		j,
		"roleSearchSubtree",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetServiceAccountPassword(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountPassword",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetServiceAccountUsername(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountUsername",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetUserBase(val *string) {
	_jsii_.Set(
		j,
		"userBase",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetUserRoleName(val *string) {
	_jsii_.Set(
		j,
		"userRoleName",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetUserSearchMatching(val *string) {
	_jsii_.Set(
		j,
		"userSearchMatching",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) SetUserSearchSubtree(val interface{}) {
	_jsii_.Set(
		j,
		"userSearchSubtree",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		m,
		"resetHosts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleBase() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleBase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleName() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleSearchMatching() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleSearchMatching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleSearchSubtree() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleSearchSubtree",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetServiceAccountPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccountPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetServiceAccountUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccountUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserBase() {
	_jsii_.InvokeVoid(
		m,
		"resetUserBase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserRoleName() {
	_jsii_.InvokeVoid(
		m,
		"resetUserRoleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserSearchMatching() {
	_jsii_.InvokeVoid(
		m,
		"resetUserSearchMatching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserSearchSubtree() {
	_jsii_.InvokeVoid(
		m,
		"resetUserSearchSubtree",
		nil, // no parameters
	)
}

type MqBrokerLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#audit MqBroker#audit}.
	Audit *string `json:"audit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#general MqBroker#general}.
	General interface{} `json:"general"`
}

type MqBrokerLogsOutputReference interface {
	cdktf.ComplexObject
	Audit() *string
	SetAudit(val *string)
	AuditInput() *string
	General() interface{}
	SetGeneral(val interface{})
	GeneralInput() interface{}
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
	ResetAudit()
	ResetGeneral()
}

// The jsii proxy struct for MqBrokerLogsOutputReference
type jsiiProxy_MqBrokerLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) Audit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) AuditInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) General() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"general",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) GeneralInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMqBrokerLogsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MqBrokerLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMqBrokerLogsOutputReference_Override(m MqBrokerLogsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) SetAudit(val *string) {
	_jsii_.Set(
		j,
		"audit",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) SetGeneral(val interface{}) {
	_jsii_.Set(
		j,
		"general",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLogsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerLogsOutputReference) ResetAudit() {
	_jsii_.InvokeVoid(
		m,
		"resetAudit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLogsOutputReference) ResetGeneral() {
	_jsii_.InvokeVoid(
		m,
		"resetGeneral",
		nil, // no parameters
	)
}

type MqBrokerMaintenanceWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#day_of_week MqBroker#day_of_week}.
	DayOfWeek *string `json:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#time_of_day MqBroker#time_of_day}.
	TimeOfDay *string `json:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#time_zone MqBroker#time_zone}.
	TimeZone *string `json:"timeZone"`
}

type MqBrokerMaintenanceWindowStartTimeOutputReference interface {
	cdktf.ComplexObject
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TimeOfDay() *string
	SetTimeOfDay(val *string)
	TimeOfDayInput() *string
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for MqBrokerMaintenanceWindowStartTimeOutputReference
type jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TimeOfDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TimeOfDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func NewMqBrokerMaintenanceWindowStartTimeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MqBrokerMaintenanceWindowStartTimeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerMaintenanceWindowStartTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMqBrokerMaintenanceWindowStartTimeOutputReference_Override(m MqBrokerMaintenanceWindowStartTimeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqBrokerMaintenanceWindowStartTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetDayOfWeek(val *string) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetTimeOfDay(val *string) {
	_jsii_.Set(
		j,
		"timeOfDay",
		val,
	)
}

func (j *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type MqBrokerUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#password MqBroker#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#username MqBroker#username}.
	Username *string `json:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#console_access MqBroker#console_access}.
	ConsoleAccess interface{} `json:"consoleAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker.html#groups MqBroker#groups}.
	Groups *[]*string `json:"groups"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html aws_mq_configuration}.
type MqConfiguration interface {
	cdktf.TerraformResource
	Arn() *string
	AuthenticationStrategy() *string
	SetAuthenticationStrategy(val *string)
	AuthenticationStrategyInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Data() *string
	SetData(val *string)
	DataInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EngineType() *string
	SetEngineType(val *string)
	EngineTypeInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LatestRevision() *float64
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
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAuthenticationStrategy()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MqConfiguration
type jsiiProxy_MqConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MqConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) AuthenticationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) DataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) LatestRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html aws_mq_configuration} Resource.
func NewMqConfiguration(scope constructs.Construct, id *string, config *MqConfigurationConfig) MqConfiguration {
	_init_.Initialize()

	j := jsiiProxy_MqConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.MQ.MqConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html aws_mq_configuration} Resource.
func NewMqConfiguration_Override(m MqConfiguration, scope constructs.Construct, id *string, config *MqConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.MQ.MqConfiguration",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MqConfiguration) SetAuthenticationStrategy(val *string) {
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetData(val *string) {
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetEngineType(val *string) {
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MqConfiguration) SetTagsAll(val interface{}) {
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
func MqConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.MQ.MqConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MqConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.MQ.MqConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MqConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MqConfiguration) ResetAuthenticationStrategy() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthenticationStrategy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqConfiguration) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MqConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqConfiguration) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MqConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MqConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MqConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MqConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#data MqConfiguration#data}.
	Data *string `json:"data"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#engine_type MqConfiguration#engine_type}.
	EngineType *string `json:"engineType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#engine_version MqConfiguration#engine_version}.
	EngineVersion *string `json:"engineVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#name MqConfiguration#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#authentication_strategy MqConfiguration#authentication_strategy}.
	AuthenticationStrategy *string `json:"authenticationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#description MqConfiguration#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#tags MqConfiguration#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_configuration.html#tags_all MqConfiguration#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}
