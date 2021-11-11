package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/glue/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/glue_connection.html aws_glue_connection}.
type DataAwsGlueConnection interface {
	cdktf.TerraformDataSource
	Arn() *string
	CatalogId() *string
	CdktfStack() cdktf.TerraformStack
	ConnectionType() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchCriteria() *[]*string
	Name() *string
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
	ConnectionProperties(key *string) *string
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PhysicalConnectionRequirements(index *string) DataAwsGlueConnectionPhysicalConnectionRequirements
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsGlueConnection
type jsiiProxy_DataAwsGlueConnection struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsGlueConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) MatchCriteria() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_connection.html aws_glue_connection} Data Source.
func NewDataAwsGlueConnection(scope constructs.Construct, id *string, config *DataAwsGlueConnectionConfig) DataAwsGlueConnection {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueConnection{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_connection.html aws_glue_connection} Data Source.
func NewDataAwsGlueConnection_Override(d DataAwsGlueConnection, scope constructs.Construct, id *string, config *DataAwsGlueConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueConnection",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnection) SetTags(val interface{}) {
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
func DataAwsGlueConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.DataAwsGlueConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsGlueConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.DataAwsGlueConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsGlueConnection) ConnectionProperties(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"connectionProperties",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueConnection) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsGlueConnection) PhysicalConnectionRequirements(index *string) DataAwsGlueConnectionPhysicalConnectionRequirements {
	var returns DataAwsGlueConnectionPhysicalConnectionRequirements

	_jsii_.Invoke(
		d,
		"physicalConnectionRequirements",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsGlueConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlueConnection) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlueConnection) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsGlueConnection) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsGlueConnection) ToString() *string {
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
func (d *jsiiProxy_DataAwsGlueConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsGlueConnectionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_connection.html#id DataAwsGlueConnection#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_connection.html#tags DataAwsGlueConnection#tags}.
	Tags interface{} `json:"tags"`
}

type DataAwsGlueConnectionPhysicalConnectionRequirements interface {
	cdktf.ComplexComputedList
	AvailabilityZone() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	SecurityGroupIdList() *[]*string
	SubnetId() *string
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

// The jsii proxy struct for DataAwsGlueConnectionPhysicalConnectionRequirements
type jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) SecurityGroupIdList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlueConnectionPhysicalConnectionRequirements(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlueConnectionPhysicalConnectionRequirements {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueConnectionPhysicalConnectionRequirements",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlueConnectionPhysicalConnectionRequirements_Override(d DataAwsGlueConnectionPhysicalConnectionRequirements, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueConnectionPhysicalConnectionRequirements",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueConnectionPhysicalConnectionRequirements) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings}.
type DataAwsGlueDataCatalogEncryptionSettings interface {
	cdktf.TerraformDataSource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
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
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	DataCatalogEncryptionSettings(index *string) DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings
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

// The jsii proxy struct for DataAwsGlueDataCatalogEncryptionSettings
type jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings} Data Source.
func NewDataAwsGlueDataCatalogEncryptionSettings(scope constructs.Construct, id *string, config *DataAwsGlueDataCatalogEncryptionSettingsConfig) DataAwsGlueDataCatalogEncryptionSettings {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings} Data Source.
func NewDataAwsGlueDataCatalogEncryptionSettings_Override(d DataAwsGlueDataCatalogEncryptionSettings, scope constructs.Construct, id *string, config *DataAwsGlueDataCatalogEncryptionSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsGlueDataCatalogEncryptionSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsGlueDataCatalogEncryptionSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) DataCatalogEncryptionSettings(index *string) DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings {
	var returns DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings

	_jsii_.Invoke(
		d,
		"dataCatalogEncryptionSettings",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) ToString() *string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsGlueDataCatalogEncryptionSettingsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_data_catalog_encryption_settings.html#catalog_id DataAwsGlueDataCatalogEncryptionSettings#catalog_id}.
	CatalogId *string `json:"catalogId"`
}

type DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConnectionPasswordEncryption() interface{}
	EncryptionAtRest() interface{}
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

// The jsii proxy struct for DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings
type jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) ConnectionPasswordEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionPasswordEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) EncryptionAtRest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings_Override(d DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption interface {
	cdktf.ComplexComputedList
	AwsKmsKeyId() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ReturnConnectionPasswordEncrypted() interface{}
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

// The jsii proxy struct for DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption
type jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) AwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) ReturnConnectionPasswordEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnConnectionPasswordEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption_Override(d DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest interface {
	cdktf.ComplexComputedList
	CatalogEncryptionMode() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	SseAwsKmsKeyId() *string
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

// The jsii proxy struct for DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest
type jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) CatalogEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) SseAwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest_Override(d DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html aws_glue_script}.
type DataAwsGlueScript interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DagEdge() *[]*DataAwsGlueScriptDagEdge
	SetDagEdge(val *[]*DataAwsGlueScriptDagEdge)
	DagEdgeInput() *[]*DataAwsGlueScriptDagEdge
	DagNode() *[]*DataAwsGlueScriptDagNode
	SetDagNode(val *[]*DataAwsGlueScriptDagNode)
	DagNodeInput() *[]*DataAwsGlueScriptDagNode
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PythonScript() *string
	RawOverrides() interface{}
	ScalaCode() *string
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
	ResetLanguage()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsGlueScript
type jsiiProxy_DataAwsGlueScript struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsGlueScript) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) DagEdge() *[]*DataAwsGlueScriptDagEdge {
	var returns *[]*DataAwsGlueScriptDagEdge
	_jsii_.Get(
		j,
		"dagEdge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) DagEdgeInput() *[]*DataAwsGlueScriptDagEdge {
	var returns *[]*DataAwsGlueScriptDagEdge
	_jsii_.Get(
		j,
		"dagEdgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) DagNode() *[]*DataAwsGlueScriptDagNode {
	var returns *[]*DataAwsGlueScriptDagNode
	_jsii_.Get(
		j,
		"dagNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) DagNodeInput() *[]*DataAwsGlueScriptDagNode {
	var returns *[]*DataAwsGlueScriptDagNode
	_jsii_.Get(
		j,
		"dagNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) PythonScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) ScalaCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalaCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsGlueScript) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html aws_glue_script} Data Source.
func NewDataAwsGlueScript(scope constructs.Construct, id *string, config *DataAwsGlueScriptConfig) DataAwsGlueScript {
	_init_.Initialize()

	j := jsiiProxy_DataAwsGlueScript{}

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueScript",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html aws_glue_script} Data Source.
func NewDataAwsGlueScript_Override(d DataAwsGlueScript, scope constructs.Construct, id *string, config *DataAwsGlueScriptConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.DataAwsGlueScript",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetDagEdge(val *[]*DataAwsGlueScriptDagEdge) {
	_jsii_.Set(
		j,
		"dagEdge",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetDagNode(val *[]*DataAwsGlueScriptDagNode) {
	_jsii_.Set(
		j,
		"dagNode",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetLanguage(val *string) {
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsGlueScript) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsGlueScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.DataAwsGlueScript",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsGlueScript_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.DataAwsGlueScript",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueScript) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsGlueScript) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueScript) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsGlueScript) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsGlueScript) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsGlueScript) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsGlueScript) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsGlueScript) ResetLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsGlueScript) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsGlueScript) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsGlueScript) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsGlueScript) ToString() *string {
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
func (d *jsiiProxy_DataAwsGlueScript) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsGlueScriptConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// dag_edge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#dag_edge DataAwsGlueScript#dag_edge}
	DagEdge *[]*DataAwsGlueScriptDagEdge `json:"dagEdge"`
	// dag_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#dag_node DataAwsGlueScript#dag_node}
	DagNode *[]*DataAwsGlueScriptDagNode `json:"dagNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#language DataAwsGlueScript#language}.
	Language *string `json:"language"`
}

type DataAwsGlueScriptDagEdge struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#source DataAwsGlueScript#source}.
	Source *string `json:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#target DataAwsGlueScript#target}.
	Target *string `json:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#target_parameter DataAwsGlueScript#target_parameter}.
	TargetParameter *string `json:"targetParameter"`
}

type DataAwsGlueScriptDagNode struct {
	// args block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#args DataAwsGlueScript#args}
	Args *[]*DataAwsGlueScriptDagNodeArgs `json:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#id DataAwsGlueScript#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#node_type DataAwsGlueScript#node_type}.
	NodeType *string `json:"nodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#line_number DataAwsGlueScript#line_number}.
	LineNumber *float64 `json:"lineNumber"`
}

type DataAwsGlueScriptDagNodeArgs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#name DataAwsGlueScript#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#value DataAwsGlueScript#value}.
	Value *string `json:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script.html#param DataAwsGlueScript#param}.
	Param interface{} `json:"param"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html aws_glue_catalog_database}.
type GlueCatalogDatabase interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TargetDatabase() GlueCatalogDatabaseTargetDatabaseOutputReference
	TargetDatabaseInput() *GlueCatalogDatabaseTargetDatabase
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
	PutTargetDatabase(value *GlueCatalogDatabaseTargetDatabase)
	ResetCatalogId()
	ResetDescription()
	ResetLocationUri()
	ResetOverrideLogicalId()
	ResetParameters()
	ResetTargetDatabase()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueCatalogDatabase
type jsiiProxy_GlueCatalogDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueCatalogDatabase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) TargetDatabase() GlueCatalogDatabaseTargetDatabaseOutputReference {
	var returns GlueCatalogDatabaseTargetDatabaseOutputReference
	_jsii_.Get(
		j,
		"targetDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) TargetDatabaseInput() *GlueCatalogDatabaseTargetDatabase {
	var returns *GlueCatalogDatabaseTargetDatabase
	_jsii_.Get(
		j,
		"targetDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html aws_glue_catalog_database} Resource.
func NewGlueCatalogDatabase(scope constructs.Construct, id *string, config *GlueCatalogDatabaseConfig) GlueCatalogDatabase {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogDatabase{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html aws_glue_catalog_database} Resource.
func NewGlueCatalogDatabase_Override(g GlueCatalogDatabase, scope constructs.Construct, id *string, config *GlueCatalogDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogDatabase",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetLocationUri(val *string) {
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabase) SetProvider(val cdktf.TerraformProvider) {
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
func GlueCatalogDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueCatalogDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueCatalogDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueCatalogDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueCatalogDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogDatabase) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) PutTargetDatabase(value *GlueCatalogDatabaseTargetDatabase) {
	_jsii_.InvokeVoid(
		g,
		"putTargetDatabase",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) ResetLocationUri() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationUri",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueCatalogDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) ResetTargetDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogDatabase) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueCatalogDatabase) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueCatalogDatabase) ToString() *string {
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
func (g *jsiiProxy_GlueCatalogDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueCatalogDatabaseConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#name GlueCatalogDatabase#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#catalog_id GlueCatalogDatabase#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#description GlueCatalogDatabase#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#location_uri GlueCatalogDatabase#location_uri}.
	LocationUri *string `json:"locationUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#parameters GlueCatalogDatabase#parameters}.
	Parameters interface{} `json:"parameters"`
	// target_database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#target_database GlueCatalogDatabase#target_database}
	TargetDatabase *GlueCatalogDatabaseTargetDatabase `json:"targetDatabase"`
}

type GlueCatalogDatabaseTargetDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#catalog_id GlueCatalogDatabase#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database.html#database_name GlueCatalogDatabase#database_name}.
	DatabaseName *string `json:"databaseName"`
}

type GlueCatalogDatabaseTargetDatabaseOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
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

// The jsii proxy struct for GlueCatalogDatabaseTargetDatabaseOutputReference
type jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogDatabaseTargetDatabaseOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogDatabaseTargetDatabaseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogDatabaseTargetDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogDatabaseTargetDatabaseOutputReference_Override(g GlueCatalogDatabaseTargetDatabaseOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogDatabaseTargetDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogDatabaseTargetDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html aws_glue_catalog_table}.
type GlueCatalogTable interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	PartitionIndex() *[]*GlueCatalogTablePartitionIndex
	SetPartitionIndex(val *[]*GlueCatalogTablePartitionIndex)
	PartitionIndexInput() *[]*GlueCatalogTablePartitionIndex
	PartitionKeys() *[]*GlueCatalogTablePartitionKeys
	SetPartitionKeys(val *[]*GlueCatalogTablePartitionKeys)
	PartitionKeysInput() *[]*GlueCatalogTablePartitionKeys
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Retention() *float64
	SetRetention(val *float64)
	RetentionInput() *float64
	StorageDescriptor() GlueCatalogTableStorageDescriptorOutputReference
	StorageDescriptorInput() *GlueCatalogTableStorageDescriptor
	TableType() *string
	SetTableType(val *string)
	TableTypeInput() *string
	TargetTable() GlueCatalogTableTargetTableOutputReference
	TargetTableInput() *GlueCatalogTableTargetTable
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ViewExpandedText() *string
	SetViewExpandedText(val *string)
	ViewExpandedTextInput() *string
	ViewOriginalText() *string
	SetViewOriginalText(val *string)
	ViewOriginalTextInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutStorageDescriptor(value *GlueCatalogTableStorageDescriptor)
	PutTargetTable(value *GlueCatalogTableTargetTable)
	ResetCatalogId()
	ResetDescription()
	ResetOverrideLogicalId()
	ResetOwner()
	ResetParameters()
	ResetPartitionIndex()
	ResetPartitionKeys()
	ResetRetention()
	ResetStorageDescriptor()
	ResetTableType()
	ResetTargetTable()
	ResetViewExpandedText()
	ResetViewOriginalText()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueCatalogTable
type jsiiProxy_GlueCatalogTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueCatalogTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionIndex() *[]*GlueCatalogTablePartitionIndex {
	var returns *[]*GlueCatalogTablePartitionIndex
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionIndexInput() *[]*GlueCatalogTablePartitionIndex {
	var returns *[]*GlueCatalogTablePartitionIndex
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionKeys() *[]*GlueCatalogTablePartitionKeys {
	var returns *[]*GlueCatalogTablePartitionKeys
	_jsii_.Get(
		j,
		"partitionKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionKeysInput() *[]*GlueCatalogTablePartitionKeys {
	var returns *[]*GlueCatalogTablePartitionKeys
	_jsii_.Get(
		j,
		"partitionKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Retention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) RetentionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) StorageDescriptor() GlueCatalogTableStorageDescriptorOutputReference {
	var returns GlueCatalogTableStorageDescriptorOutputReference
	_jsii_.Get(
		j,
		"storageDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) StorageDescriptorInput() *GlueCatalogTableStorageDescriptor {
	var returns *GlueCatalogTableStorageDescriptor
	_jsii_.Get(
		j,
		"storageDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TargetTable() GlueCatalogTableTargetTableOutputReference {
	var returns GlueCatalogTableTargetTableOutputReference
	_jsii_.Get(
		j,
		"targetTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TargetTableInput() *GlueCatalogTableTargetTable {
	var returns *GlueCatalogTableTargetTable
	_jsii_.Get(
		j,
		"targetTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewExpandedTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewOriginalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalTextInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html aws_glue_catalog_table} Resource.
func NewGlueCatalogTable(scope constructs.Construct, id *string, config *GlueCatalogTableConfig) GlueCatalogTable {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTable{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html aws_glue_catalog_table} Resource.
func NewGlueCatalogTable_Override(g GlueCatalogTable, scope constructs.Construct, id *string, config *GlueCatalogTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTable",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetPartitionIndex(val *[]*GlueCatalogTablePartitionIndex) {
	_jsii_.Set(
		j,
		"partitionIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetPartitionKeys(val *[]*GlueCatalogTablePartitionKeys) {
	_jsii_.Set(
		j,
		"partitionKeys",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetRetention(val *float64) {
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetTableType(val *string) {
	_jsii_.Set(
		j,
		"tableType",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetViewExpandedText(val *string) {
	_jsii_.Set(
		j,
		"viewExpandedText",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable) SetViewOriginalText(val *string) {
	_jsii_.Set(
		j,
		"viewOriginalText",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueCatalogTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueCatalogTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueCatalogTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueCatalogTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTable) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTable) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutStorageDescriptor(value *GlueCatalogTableStorageDescriptor) {
	_jsii_.InvokeVoid(
		g,
		"putStorageDescriptor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutTargetTable(value *GlueCatalogTableTargetTable) {
	_jsii_.InvokeVoid(
		g,
		"putTargetTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueCatalogTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetOwner() {
	_jsii_.InvokeVoid(
		g,
		"resetOwner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetPartitionIndex() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionIndex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetPartitionKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetStorageDescriptor() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageDescriptor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetTableType() {
	_jsii_.InvokeVoid(
		g,
		"resetTableType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetTargetTable() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetViewExpandedText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewExpandedText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetViewOriginalText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewOriginalText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueCatalogTable) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueCatalogTable) ToString() *string {
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
func (g *jsiiProxy_GlueCatalogTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueCatalogTableConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#database_name GlueCatalogTable#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#name GlueCatalogTable#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#catalog_id GlueCatalogTable#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#description GlueCatalogTable#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#owner GlueCatalogTable#owner}.
	Owner *string `json:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#parameters GlueCatalogTable#parameters}.
	Parameters interface{} `json:"parameters"`
	// partition_index block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#partition_index GlueCatalogTable#partition_index}
	PartitionIndex *[]*GlueCatalogTablePartitionIndex `json:"partitionIndex"`
	// partition_keys block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#partition_keys GlueCatalogTable#partition_keys}
	PartitionKeys *[]*GlueCatalogTablePartitionKeys `json:"partitionKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#retention GlueCatalogTable#retention}.
	Retention *float64 `json:"retention"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#storage_descriptor GlueCatalogTable#storage_descriptor}
	StorageDescriptor *GlueCatalogTableStorageDescriptor `json:"storageDescriptor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#table_type GlueCatalogTable#table_type}.
	TableType *string `json:"tableType"`
	// target_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#target_table GlueCatalogTable#target_table}
	TargetTable *GlueCatalogTableTargetTable `json:"targetTable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#view_expanded_text GlueCatalogTable#view_expanded_text}.
	ViewExpandedText *string `json:"viewExpandedText"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#view_original_text GlueCatalogTable#view_original_text}.
	ViewOriginalText *string `json:"viewOriginalText"`
}

type GlueCatalogTablePartitionIndex struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#index_name GlueCatalogTable#index_name}.
	IndexName *string `json:"indexName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#keys GlueCatalogTable#keys}.
	Keys *[]*string `json:"keys"`
}

type GlueCatalogTablePartitionKeys struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#name GlueCatalogTable#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#comment GlueCatalogTable#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#type GlueCatalogTable#type}.
	Type *string `json:"type"`
}

type GlueCatalogTableStorageDescriptor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#bucket_columns GlueCatalogTable#bucket_columns}.
	BucketColumns *[]*string `json:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#columns GlueCatalogTable#columns}
	Columns *[]*GlueCatalogTableStorageDescriptorColumns `json:"columns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#compressed GlueCatalogTable#compressed}.
	Compressed interface{} `json:"compressed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#input_format GlueCatalogTable#input_format}.
	InputFormat *string `json:"inputFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#location GlueCatalogTable#location}.
	Location *string `json:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#number_of_buckets GlueCatalogTable#number_of_buckets}.
	NumberOfBuckets *float64 `json:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#output_format GlueCatalogTable#output_format}.
	OutputFormat *string `json:"outputFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#parameters GlueCatalogTable#parameters}.
	Parameters interface{} `json:"parameters"`
	// schema_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_reference GlueCatalogTable#schema_reference}
	SchemaReference *GlueCatalogTableStorageDescriptorSchemaReference `json:"schemaReference"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#ser_de_info GlueCatalogTable#ser_de_info}
	SerDeInfo *GlueCatalogTableStorageDescriptorSerDeInfo `json:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#skewed_info GlueCatalogTable#skewed_info}
	SkewedInfo *GlueCatalogTableStorageDescriptorSkewedInfo `json:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#sort_columns GlueCatalogTable#sort_columns}
	SortColumns *[]*GlueCatalogTableStorageDescriptorSortColumns `json:"sortColumns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#stored_as_sub_directories GlueCatalogTable#stored_as_sub_directories}.
	StoredAsSubDirectories interface{} `json:"storedAsSubDirectories"`
}

type GlueCatalogTableStorageDescriptorColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#name GlueCatalogTable#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#comment GlueCatalogTable#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#parameters GlueCatalogTable#parameters}.
	Parameters interface{} `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#type GlueCatalogTable#type}.
	Type *string `json:"type"`
}

type GlueCatalogTableStorageDescriptorOutputReference interface {
	cdktf.ComplexObject
	BucketColumns() *[]*string
	SetBucketColumns(val *[]*string)
	BucketColumnsInput() *[]*string
	Columns() *[]*GlueCatalogTableStorageDescriptorColumns
	SetColumns(val *[]*GlueCatalogTableStorageDescriptorColumns)
	ColumnsInput() *[]*GlueCatalogTableStorageDescriptorColumns
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	NumberOfBuckets() *float64
	SetNumberOfBuckets(val *float64)
	NumberOfBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	SchemaReference() GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference
	SchemaReferenceInput() *GlueCatalogTableStorageDescriptorSchemaReference
	SerDeInfo() GlueCatalogTableStorageDescriptorSerDeInfoOutputReference
	SerDeInfoInput() *GlueCatalogTableStorageDescriptorSerDeInfo
	SkewedInfo() GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() *GlueCatalogTableStorageDescriptorSkewedInfo
	SortColumns() *[]*GlueCatalogTableStorageDescriptorSortColumns
	SetSortColumns(val *[]*GlueCatalogTableStorageDescriptorSortColumns)
	SortColumnsInput() *[]*GlueCatalogTableStorageDescriptorSortColumns
	StoredAsSubDirectories() interface{}
	SetStoredAsSubDirectories(val interface{})
	StoredAsSubDirectoriesInput() interface{}
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
	PutSchemaReference(value *GlueCatalogTableStorageDescriptorSchemaReference)
	PutSerDeInfo(value *GlueCatalogTableStorageDescriptorSerDeInfo)
	PutSkewedInfo(value *GlueCatalogTableStorageDescriptorSkewedInfo)
	ResetBucketColumns()
	ResetColumns()
	ResetCompressed()
	ResetInputFormat()
	ResetLocation()
	ResetNumberOfBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSchemaReference()
	ResetSerDeInfo()
	ResetSkewedInfo()
	ResetSortColumns()
	ResetStoredAsSubDirectories()
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Columns() *[]*GlueCatalogTableStorageDescriptorColumns {
	var returns *[]*GlueCatalogTableStorageDescriptorColumns
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ColumnsInput() *[]*GlueCatalogTableStorageDescriptorColumns {
	var returns *[]*GlueCatalogTableStorageDescriptorColumns
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SchemaReference() GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference {
	var returns GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SchemaReferenceInput() *GlueCatalogTableStorageDescriptorSchemaReference {
	var returns *GlueCatalogTableStorageDescriptorSchemaReference
	_jsii_.Get(
		j,
		"schemaReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SerDeInfo() GlueCatalogTableStorageDescriptorSerDeInfoOutputReference {
	var returns GlueCatalogTableStorageDescriptorSerDeInfoOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SerDeInfoInput() *GlueCatalogTableStorageDescriptorSerDeInfo {
	var returns *GlueCatalogTableStorageDescriptorSerDeInfo
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SkewedInfo() GlueCatalogTableStorageDescriptorSkewedInfoOutputReference {
	var returns GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SkewedInfoInput() *GlueCatalogTableStorageDescriptorSkewedInfo {
	var returns *GlueCatalogTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SortColumns() *[]*GlueCatalogTableStorageDescriptorSortColumns {
	var returns *[]*GlueCatalogTableStorageDescriptorSortColumns
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SortColumnsInput() *[]*GlueCatalogTableStorageDescriptorSortColumns {
	var returns *[]*GlueCatalogTableStorageDescriptorSortColumns
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableStorageDescriptorOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableStorageDescriptorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorOutputReference_Override(g GlueCatalogTableStorageDescriptorOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetBucketColumns(val *[]*string) {
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetColumns(val *[]*GlueCatalogTableStorageDescriptorColumns) {
	_jsii_.Set(
		j,
		"columns",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetCompressed(val interface{}) {
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetInputFormat(val *string) {
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetNumberOfBuckets(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetOutputFormat(val *string) {
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetSortColumns(val *[]*GlueCatalogTableStorageDescriptorSortColumns) {
	_jsii_.Set(
		j,
		"sortColumns",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetStoredAsSubDirectories(val interface{}) {
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSchemaReference(value *GlueCatalogTableStorageDescriptorSchemaReference) {
	_jsii_.InvokeVoid(
		g,
		"putSchemaReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSerDeInfo(value *GlueCatalogTableStorageDescriptorSerDeInfo) {
	_jsii_.InvokeVoid(
		g,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSkewedInfo(value *GlueCatalogTableStorageDescriptorSkewedInfo) {
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSchemaReference() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

type GlueCatalogTableStorageDescriptorSchemaReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_version_number GlueCatalogTable#schema_version_number}.
	SchemaVersionNumber *float64 `json:"schemaVersionNumber"`
	// schema_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_id GlueCatalogTable#schema_id}
	SchemaId *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId `json:"schemaId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_version_id GlueCatalogTable#schema_version_id}.
	SchemaVersionId *string `json:"schemaVersionId"`
}

type GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SchemaId() GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference
	SchemaIdInput() *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId
	SchemaVersionId() *string
	SetSchemaVersionId(val *string)
	SchemaVersionIdInput() *string
	SchemaVersionNumber() *float64
	SetSchemaVersionNumber(val *float64)
	SchemaVersionNumberInput() *float64
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
	PutSchemaId(value *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId)
	ResetSchemaId()
	ResetSchemaVersionId()
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaId() GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference {
	var returns GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaIdInput() *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId {
	var returns *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SchemaVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableStorageDescriptorSchemaReferenceOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorSchemaReferenceOutputReference_Override(g GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SetSchemaVersionId(val *string) {
	_jsii_.Set(
		j,
		"schemaVersionId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SetSchemaVersionNumber(val *float64) {
	_jsii_.Set(
		j,
		"schemaVersionNumber",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) PutSchemaId(value *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId) {
	_jsii_.InvokeVoid(
		g,
		"putSchemaId",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference) ResetSchemaVersionId() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaVersionId",
		nil, // no parameters
	)
}

type GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#registry_name GlueCatalogTable#registry_name}.
	RegistryName *string `json:"registryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_arn GlueCatalogTable#schema_arn}.
	SchemaArn *string `json:"schemaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#schema_name GlueCatalogTable#schema_name}.
	SchemaName *string `json:"schemaName"`
}

type GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RegistryName() *string
	SetRegistryName(val *string)
	RegistryNameInput() *string
	SchemaArn() *string
	SetSchemaArn(val *string)
	SchemaArnInput() *string
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
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
	ResetRegistryName()
	ResetSchemaArn()
	ResetSchemaName()
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) RegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) RegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SchemaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SchemaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference_Override(g GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetRegistryName(val *string) {
	_jsii_.Set(
		j,
		"registryName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetSchemaArn(val *string) {
	_jsii_.Set(
		j,
		"schemaArn",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetSchemaName(val *string) {
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) ResetRegistryName() {
	_jsii_.InvokeVoid(
		g,
		"resetRegistryName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) ResetSchemaArn() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSchemaReferenceSchemaIdOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaName",
		nil, // no parameters
	)
}

type GlueCatalogTableStorageDescriptorSerDeInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#name GlueCatalogTable#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#parameters GlueCatalogTable#parameters}.
	Parameters interface{} `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#serialization_library GlueCatalogTable#serialization_library}.
	SerializationLibrary *string `json:"serializationLibrary"`
}

type GlueCatalogTableStorageDescriptorSerDeInfoOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	SerializationLibrary() *string
	SetSerializationLibrary(val *string)
	SerializationLibraryInput() *string
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
	ResetName()
	ResetParameters()
	ResetSerializationLibrary()
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorSerDeInfoOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SerializationLibrary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibrary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SerializationLibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableStorageDescriptorSerDeInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableStorageDescriptorSerDeInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSerDeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorSerDeInfoOutputReference_Override(g GlueCatalogTableStorageDescriptorSerDeInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSerDeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetSerializationLibrary(val *string) {
	_jsii_.Set(
		j,
		"serializationLibrary",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSerDeInfoOutputReference) ResetSerializationLibrary() {
	_jsii_.InvokeVoid(
		g,
		"resetSerializationLibrary",
		nil, // no parameters
	)
}

type GlueCatalogTableStorageDescriptorSkewedInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#skewed_column_names GlueCatalogTable#skewed_column_names}.
	SkewedColumnNames *[]*string `json:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#skewed_column_value_location_maps GlueCatalogTable#skewed_column_value_location_maps}.
	SkewedColumnValueLocationMaps interface{} `json:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#skewed_column_values GlueCatalogTable#skewed_column_values}.
	SkewedColumnValues *[]*string `json:"skewedColumnValues"`
}

type GlueCatalogTableStorageDescriptorSkewedInfoOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SkewedColumnNames() *[]*string
	SetSkewedColumnNames(val *[]*string)
	SkewedColumnNamesInput() *[]*string
	SkewedColumnValueLocationMaps() interface{}
	SetSkewedColumnValueLocationMaps(val interface{})
	SkewedColumnValueLocationMapsInput() interface{}
	SkewedColumnValues() *[]*string
	SetSkewedColumnValues(val *[]*string)
	SkewedColumnValuesInput() *[]*string
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
	ResetSkewedColumnNames()
	ResetSkewedColumnValueLocationMaps()
	ResetSkewedColumnValues()
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMaps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMapsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableStorageDescriptorSkewedInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableStorageDescriptorSkewedInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorSkewedInfoOutputReference_Override(g GlueCatalogTableStorageDescriptorSkewedInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"skewedColumnNames",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnValueLocationMaps(val interface{}) {
	_jsii_.Set(
		j,
		"skewedColumnValueLocationMaps",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnValues(val *[]*string) {
	_jsii_.Set(
		j,
		"skewedColumnValues",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValueLocationMaps() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValueLocationMaps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValues() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValues",
		nil, // no parameters
	)
}

type GlueCatalogTableStorageDescriptorSortColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#column GlueCatalogTable#column}.
	Column *string `json:"column"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#sort_order GlueCatalogTable#sort_order}.
	SortOrder *float64 `json:"sortOrder"`
}

type GlueCatalogTableTargetTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#catalog_id GlueCatalogTable#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#database_name GlueCatalogTable#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table.html#name GlueCatalogTable#name}.
	Name *string `json:"name"`
}

type GlueCatalogTableTargetTableOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
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

// The jsii proxy struct for GlueCatalogTableTargetTableOutputReference
type jsiiProxy_GlueCatalogTableTargetTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCatalogTableTargetTableOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCatalogTableTargetTableOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCatalogTableTargetTableOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableTargetTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCatalogTableTargetTableOutputReference_Override(g GlueCatalogTableTargetTableOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCatalogTableTargetTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableTargetTableOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCatalogTableTargetTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html aws_glue_classifier}.
type GlueClassifier interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CsvClassifier() GlueClassifierCsvClassifierOutputReference
	CsvClassifierInput() *GlueClassifierCsvClassifier
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	GrokClassifier() GlueClassifierGrokClassifierOutputReference
	GrokClassifierInput() *GlueClassifierGrokClassifier
	Id() *string
	JsonClassifier() GlueClassifierJsonClassifierOutputReference
	JsonClassifierInput() *GlueClassifierJsonClassifier
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
	XmlClassifier() GlueClassifierXmlClassifierOutputReference
	XmlClassifierInput() *GlueClassifierXmlClassifier
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCsvClassifier(value *GlueClassifierCsvClassifier)
	PutGrokClassifier(value *GlueClassifierGrokClassifier)
	PutJsonClassifier(value *GlueClassifierJsonClassifier)
	PutXmlClassifier(value *GlueClassifierXmlClassifier)
	ResetCsvClassifier()
	ResetGrokClassifier()
	ResetJsonClassifier()
	ResetOverrideLogicalId()
	ResetXmlClassifier()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueClassifier
type jsiiProxy_GlueClassifier struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueClassifier) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) CsvClassifier() GlueClassifierCsvClassifierOutputReference {
	var returns GlueClassifierCsvClassifierOutputReference
	_jsii_.Get(
		j,
		"csvClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) CsvClassifierInput() *GlueClassifierCsvClassifier {
	var returns *GlueClassifierCsvClassifier
	_jsii_.Get(
		j,
		"csvClassifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) GrokClassifier() GlueClassifierGrokClassifierOutputReference {
	var returns GlueClassifierGrokClassifierOutputReference
	_jsii_.Get(
		j,
		"grokClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) GrokClassifierInput() *GlueClassifierGrokClassifier {
	var returns *GlueClassifierGrokClassifier
	_jsii_.Get(
		j,
		"grokClassifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) JsonClassifier() GlueClassifierJsonClassifierOutputReference {
	var returns GlueClassifierJsonClassifierOutputReference
	_jsii_.Get(
		j,
		"jsonClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) JsonClassifierInput() *GlueClassifierJsonClassifier {
	var returns *GlueClassifierJsonClassifier
	_jsii_.Get(
		j,
		"jsonClassifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) XmlClassifier() GlueClassifierXmlClassifierOutputReference {
	var returns GlueClassifierXmlClassifierOutputReference
	_jsii_.Get(
		j,
		"xmlClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifier) XmlClassifierInput() *GlueClassifierXmlClassifier {
	var returns *GlueClassifierXmlClassifier
	_jsii_.Get(
		j,
		"xmlClassifierInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html aws_glue_classifier} Resource.
func NewGlueClassifier(scope constructs.Construct, id *string, config *GlueClassifierConfig) GlueClassifier {
	_init_.Initialize()

	j := jsiiProxy_GlueClassifier{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifier",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html aws_glue_classifier} Resource.
func NewGlueClassifier_Override(g GlueClassifier, scope constructs.Construct, id *string, config *GlueClassifierConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifier",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueClassifier) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueClassifier) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueClassifier) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueClassifier) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueClassifier) SetProvider(val cdktf.TerraformProvider) {
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
func GlueClassifier_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueClassifier",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueClassifier_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueClassifier",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueClassifier) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueClassifier) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifier) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueClassifier) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueClassifier) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueClassifier) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifier) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueClassifier) PutCsvClassifier(value *GlueClassifierCsvClassifier) {
	_jsii_.InvokeVoid(
		g,
		"putCsvClassifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueClassifier) PutGrokClassifier(value *GlueClassifierGrokClassifier) {
	_jsii_.InvokeVoid(
		g,
		"putGrokClassifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueClassifier) PutJsonClassifier(value *GlueClassifierJsonClassifier) {
	_jsii_.InvokeVoid(
		g,
		"putJsonClassifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueClassifier) PutXmlClassifier(value *GlueClassifierXmlClassifier) {
	_jsii_.InvokeVoid(
		g,
		"putXmlClassifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueClassifier) ResetCsvClassifier() {
	_jsii_.InvokeVoid(
		g,
		"resetCsvClassifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifier) ResetGrokClassifier() {
	_jsii_.InvokeVoid(
		g,
		"resetGrokClassifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifier) ResetJsonClassifier() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonClassifier",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueClassifier) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifier) ResetXmlClassifier() {
	_jsii_.InvokeVoid(
		g,
		"resetXmlClassifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifier) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueClassifier) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueClassifier) ToString() *string {
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
func (g *jsiiProxy_GlueClassifier) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueClassifierConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#name GlueClassifier#name}.
	Name *string `json:"name"`
	// csv_classifier block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#csv_classifier GlueClassifier#csv_classifier}
	CsvClassifier *GlueClassifierCsvClassifier `json:"csvClassifier"`
	// grok_classifier block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#grok_classifier GlueClassifier#grok_classifier}
	GrokClassifier *GlueClassifierGrokClassifier `json:"grokClassifier"`
	// json_classifier block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#json_classifier GlueClassifier#json_classifier}
	JsonClassifier *GlueClassifierJsonClassifier `json:"jsonClassifier"`
	// xml_classifier block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#xml_classifier GlueClassifier#xml_classifier}
	XmlClassifier *GlueClassifierXmlClassifier `json:"xmlClassifier"`
}

type GlueClassifierCsvClassifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#allow_single_column GlueClassifier#allow_single_column}.
	AllowSingleColumn interface{} `json:"allowSingleColumn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#contains_header GlueClassifier#contains_header}.
	ContainsHeader *string `json:"containsHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#delimiter GlueClassifier#delimiter}.
	Delimiter *string `json:"delimiter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#disable_value_trimming GlueClassifier#disable_value_trimming}.
	DisableValueTrimming interface{} `json:"disableValueTrimming"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#header GlueClassifier#header}.
	Header *[]*string `json:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#quote_symbol GlueClassifier#quote_symbol}.
	QuoteSymbol *string `json:"quoteSymbol"`
}

type GlueClassifierCsvClassifierOutputReference interface {
	cdktf.ComplexObject
	AllowSingleColumn() interface{}
	SetAllowSingleColumn(val interface{})
	AllowSingleColumnInput() interface{}
	ContainsHeader() *string
	SetContainsHeader(val *string)
	ContainsHeaderInput() *string
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	DisableValueTrimming() interface{}
	SetDisableValueTrimming(val interface{})
	DisableValueTrimmingInput() interface{}
	Header() *[]*string
	SetHeader(val *[]*string)
	HeaderInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QuoteSymbol() *string
	SetQuoteSymbol(val *string)
	QuoteSymbolInput() *string
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
	ResetAllowSingleColumn()
	ResetContainsHeader()
	ResetDelimiter()
	ResetDisableValueTrimming()
	ResetHeader()
	ResetQuoteSymbol()
}

// The jsii proxy struct for GlueClassifierCsvClassifierOutputReference
type jsiiProxy_GlueClassifierCsvClassifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) AllowSingleColumn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) AllowSingleColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSingleColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ContainsHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ContainsHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DisableValueTrimming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) DisableValueTrimmingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableValueTrimmingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) Header() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) HeaderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) QuoteSymbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) QuoteSymbolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueClassifierCsvClassifierOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueClassifierCsvClassifierOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueClassifierCsvClassifierOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierCsvClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueClassifierCsvClassifierOutputReference_Override(g GlueClassifierCsvClassifierOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierCsvClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetAllowSingleColumn(val interface{}) {
	_jsii_.Set(
		j,
		"allowSingleColumn",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetContainsHeader(val *string) {
	_jsii_.Set(
		j,
		"containsHeader",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetDelimiter(val *string) {
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetDisableValueTrimming(val interface{}) {
	_jsii_.Set(
		j,
		"disableValueTrimming",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetHeader(val *[]*string) {
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetQuoteSymbol(val *string) {
	_jsii_.Set(
		j,
		"quoteSymbol",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierCsvClassifierOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetAllowSingleColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowSingleColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetContainsHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetContainsHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetDisableValueTrimming() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableValueTrimming",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueClassifierCsvClassifierOutputReference) ResetQuoteSymbol() {
	_jsii_.InvokeVoid(
		g,
		"resetQuoteSymbol",
		nil, // no parameters
	)
}

type GlueClassifierGrokClassifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#classification GlueClassifier#classification}.
	Classification *string `json:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#grok_pattern GlueClassifier#grok_pattern}.
	GrokPattern *string `json:"grokPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#custom_patterns GlueClassifier#custom_patterns}.
	CustomPatterns *string `json:"customPatterns"`
}

type GlueClassifierGrokClassifierOutputReference interface {
	cdktf.ComplexObject
	Classification() *string
	SetClassification(val *string)
	ClassificationInput() *string
	CustomPatterns() *string
	SetCustomPatterns(val *string)
	CustomPatternsInput() *string
	GrokPattern() *string
	SetGrokPattern(val *string)
	GrokPatternInput() *string
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
	ResetCustomPatterns()
}

// The jsii proxy struct for GlueClassifierGrokClassifierOutputReference
type jsiiProxy_GlueClassifierGrokClassifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) Classification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) ClassificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) CustomPatterns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) CustomPatternsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GrokPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grokPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GrokPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grokPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueClassifierGrokClassifierOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueClassifierGrokClassifierOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueClassifierGrokClassifierOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierGrokClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueClassifierGrokClassifierOutputReference_Override(g GlueClassifierGrokClassifierOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierGrokClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetClassification(val *string) {
	_jsii_.Set(
		j,
		"classification",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetCustomPatterns(val *string) {
	_jsii_.Set(
		j,
		"customPatterns",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetGrokPattern(val *string) {
	_jsii_.Set(
		j,
		"grokPattern",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierGrokClassifierOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueClassifierGrokClassifierOutputReference) ResetCustomPatterns() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomPatterns",
		nil, // no parameters
	)
}

type GlueClassifierJsonClassifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#json_path GlueClassifier#json_path}.
	JsonPath *string `json:"jsonPath"`
}

type GlueClassifierJsonClassifierOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	JsonPath() *string
	SetJsonPath(val *string)
	JsonPathInput() *string
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

// The jsii proxy struct for GlueClassifierJsonClassifierOutputReference
type jsiiProxy_GlueClassifierJsonClassifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) JsonPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) JsonPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueClassifierJsonClassifierOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueClassifierJsonClassifierOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueClassifierJsonClassifierOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierJsonClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueClassifierJsonClassifierOutputReference_Override(g GlueClassifierJsonClassifierOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierJsonClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) SetJsonPath(val *string) {
	_jsii_.Set(
		j,
		"jsonPath",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierJsonClassifierOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifierJsonClassifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type GlueClassifierXmlClassifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#classification GlueClassifier#classification}.
	Classification *string `json:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier.html#row_tag GlueClassifier#row_tag}.
	RowTag *string `json:"rowTag"`
}

type GlueClassifierXmlClassifierOutputReference interface {
	cdktf.ComplexObject
	Classification() *string
	SetClassification(val *string)
	ClassificationInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RowTag() *string
	SetRowTag(val *string)
	RowTagInput() *string
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

// The jsii proxy struct for GlueClassifierXmlClassifierOutputReference
type jsiiProxy_GlueClassifierXmlClassifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) Classification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) ClassificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) RowTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) RowTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueClassifierXmlClassifierOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueClassifierXmlClassifierOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueClassifierXmlClassifierOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierXmlClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueClassifierXmlClassifierOutputReference_Override(g GlueClassifierXmlClassifierOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueClassifierXmlClassifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) SetClassification(val *string) {
	_jsii_.Set(
		j,
		"classification",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) SetRowTag(val *string) {
	_jsii_.Set(
		j,
		"rowTag",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueClassifierXmlClassifierOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueClassifierXmlClassifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html aws_glue_connection}.
type GlueConnection interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConnectionProperties() interface{}
	SetConnectionProperties(val interface{})
	ConnectionPropertiesInput() interface{}
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchCriteria() *[]*string
	SetMatchCriteria(val *[]*string)
	MatchCriteriaInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PhysicalConnectionRequirements() GlueConnectionPhysicalConnectionRequirementsOutputReference
	PhysicalConnectionRequirementsInput() *GlueConnectionPhysicalConnectionRequirements
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
	PutPhysicalConnectionRequirements(value *GlueConnectionPhysicalConnectionRequirements)
	ResetCatalogId()
	ResetConnectionProperties()
	ResetConnectionType()
	ResetDescription()
	ResetMatchCriteria()
	ResetOverrideLogicalId()
	ResetPhysicalConnectionRequirements()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueConnection
type jsiiProxy_GlueConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) ConnectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) ConnectionPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) MatchCriteria() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) MatchCriteriaInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) PhysicalConnectionRequirements() GlueConnectionPhysicalConnectionRequirementsOutputReference {
	var returns GlueConnectionPhysicalConnectionRequirementsOutputReference
	_jsii_.Get(
		j,
		"physicalConnectionRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) PhysicalConnectionRequirementsInput() *GlueConnectionPhysicalConnectionRequirements {
	var returns *GlueConnectionPhysicalConnectionRequirements
	_jsii_.Get(
		j,
		"physicalConnectionRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html aws_glue_connection} Resource.
func NewGlueConnection(scope constructs.Construct, id *string, config *GlueConnectionConfig) GlueConnection {
	_init_.Initialize()

	j := jsiiProxy_GlueConnection{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html aws_glue_connection} Resource.
func NewGlueConnection_Override(g GlueConnection, scope constructs.Construct, id *string, config *GlueConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueConnection",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueConnection) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetConnectionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"connectionProperties",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetMatchCriteria(val *[]*string) {
	_jsii_.Set(
		j,
		"matchCriteria",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueConnection) SetTagsAll(val interface{}) {
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
func GlueConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueConnection) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueConnection) PutPhysicalConnectionRequirements(value *GlueConnectionPhysicalConnectionRequirements) {
	_jsii_.InvokeVoid(
		g,
		"putPhysicalConnectionRequirements",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnection) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetConnectionType() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetMatchCriteria() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchCriteria",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetPhysicalConnectionRequirements() {
	_jsii_.InvokeVoid(
		g,
		"resetPhysicalConnectionRequirements",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnection) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueConnection) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueConnection) ToString() *string {
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
func (g *jsiiProxy_GlueConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueConnectionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#name GlueConnection#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#catalog_id GlueConnection#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#connection_properties GlueConnection#connection_properties}.
	ConnectionProperties interface{} `json:"connectionProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#connection_type GlueConnection#connection_type}.
	ConnectionType *string `json:"connectionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#description GlueConnection#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#match_criteria GlueConnection#match_criteria}.
	MatchCriteria *[]*string `json:"matchCriteria"`
	// physical_connection_requirements block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#physical_connection_requirements GlueConnection#physical_connection_requirements}
	PhysicalConnectionRequirements *GlueConnectionPhysicalConnectionRequirements `json:"physicalConnectionRequirements"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#tags GlueConnection#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#tags_all GlueConnection#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type GlueConnectionPhysicalConnectionRequirements struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#availability_zone GlueConnection#availability_zone}.
	AvailabilityZone *string `json:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#security_group_id_list GlueConnection#security_group_id_list}.
	SecurityGroupIdList *[]*string `json:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_connection.html#subnet_id GlueConnection#subnet_id}.
	SubnetId *string `json:"subnetId"`
}

type GlueConnectionPhysicalConnectionRequirementsOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIdList() *[]*string
	SetSecurityGroupIdList(val *[]*string)
	SecurityGroupIdListInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	ResetAvailabilityZone()
	ResetSecurityGroupIdList()
	ResetSubnetId()
}

// The jsii proxy struct for GlueConnectionPhysicalConnectionRequirementsOutputReference
type jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SecurityGroupIdList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SecurityGroupIdListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueConnectionPhysicalConnectionRequirementsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueConnectionPhysicalConnectionRequirementsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueConnectionPhysicalConnectionRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueConnectionPhysicalConnectionRequirementsOutputReference_Override(g GlueConnectionPhysicalConnectionRequirementsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueConnectionPhysicalConnectionRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetSecurityGroupIdList(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIdList",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) ResetSecurityGroupIdList() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityGroupIdList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionPhysicalConnectionRequirementsOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetId",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html aws_glue_crawler}.
type GlueCrawler interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogTarget() *[]*GlueCrawlerCatalogTarget
	SetCatalogTarget(val *[]*GlueCrawlerCatalogTarget)
	CatalogTargetInput() *[]*GlueCrawlerCatalogTarget
	CdktfStack() cdktf.TerraformStack
	Classifiers() *[]*string
	SetClassifiers(val *[]*string)
	ClassifiersInput() *[]*string
	Configuration() *string
	SetConfiguration(val *string)
	ConfigurationInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamodbTarget() *[]*GlueCrawlerDynamodbTarget
	SetDynamodbTarget(val *[]*GlueCrawlerDynamodbTarget)
	DynamodbTargetInput() *[]*GlueCrawlerDynamodbTarget
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	JdbcTarget() *[]*GlueCrawlerJdbcTarget
	SetJdbcTarget(val *[]*GlueCrawlerJdbcTarget)
	JdbcTargetInput() *[]*GlueCrawlerJdbcTarget
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LineageConfiguration() GlueCrawlerLineageConfigurationOutputReference
	LineageConfigurationInput() *GlueCrawlerLineageConfiguration
	MongodbTarget() *[]*GlueCrawlerMongodbTarget
	SetMongodbTarget(val *[]*GlueCrawlerMongodbTarget)
	MongodbTargetInput() *[]*GlueCrawlerMongodbTarget
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RecrawlPolicy() GlueCrawlerRecrawlPolicyOutputReference
	RecrawlPolicyInput() *GlueCrawlerRecrawlPolicy
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	S3Target() *[]*GlueCrawlerS3Target
	SetS3Target(val *[]*GlueCrawlerS3Target)
	S3TargetInput() *[]*GlueCrawlerS3Target
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	SchemaChangePolicy() GlueCrawlerSchemaChangePolicyOutputReference
	SchemaChangePolicyInput() *GlueCrawlerSchemaChangePolicy
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	TablePrefix() *string
	SetTablePrefix(val *string)
	TablePrefixInput() *string
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
	PutLineageConfiguration(value *GlueCrawlerLineageConfiguration)
	PutRecrawlPolicy(value *GlueCrawlerRecrawlPolicy)
	PutSchemaChangePolicy(value *GlueCrawlerSchemaChangePolicy)
	ResetCatalogTarget()
	ResetClassifiers()
	ResetConfiguration()
	ResetDescription()
	ResetDynamodbTarget()
	ResetJdbcTarget()
	ResetLineageConfiguration()
	ResetMongodbTarget()
	ResetOverrideLogicalId()
	ResetRecrawlPolicy()
	ResetS3Target()
	ResetSchedule()
	ResetSchemaChangePolicy()
	ResetSecurityConfiguration()
	ResetTablePrefix()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueCrawler
type jsiiProxy_GlueCrawler struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueCrawler) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CatalogTarget() *[]*GlueCrawlerCatalogTarget {
	var returns *[]*GlueCrawlerCatalogTarget
	_jsii_.Get(
		j,
		"catalogTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CatalogTargetInput() *[]*GlueCrawlerCatalogTarget {
	var returns *[]*GlueCrawlerCatalogTarget
	_jsii_.Get(
		j,
		"catalogTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ClassifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DynamodbTarget() *[]*GlueCrawlerDynamodbTarget {
	var returns *[]*GlueCrawlerDynamodbTarget
	_jsii_.Get(
		j,
		"dynamodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) DynamodbTargetInput() *[]*GlueCrawlerDynamodbTarget {
	var returns *[]*GlueCrawlerDynamodbTarget
	_jsii_.Get(
		j,
		"dynamodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) JdbcTarget() *[]*GlueCrawlerJdbcTarget {
	var returns *[]*GlueCrawlerJdbcTarget
	_jsii_.Get(
		j,
		"jdbcTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) JdbcTargetInput() *[]*GlueCrawlerJdbcTarget {
	var returns *[]*GlueCrawlerJdbcTarget
	_jsii_.Get(
		j,
		"jdbcTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) LineageConfiguration() GlueCrawlerLineageConfigurationOutputReference {
	var returns GlueCrawlerLineageConfigurationOutputReference
	_jsii_.Get(
		j,
		"lineageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) LineageConfigurationInput() *GlueCrawlerLineageConfiguration {
	var returns *GlueCrawlerLineageConfiguration
	_jsii_.Get(
		j,
		"lineageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) MongodbTarget() *[]*GlueCrawlerMongodbTarget {
	var returns *[]*GlueCrawlerMongodbTarget
	_jsii_.Get(
		j,
		"mongodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) MongodbTargetInput() *[]*GlueCrawlerMongodbTarget {
	var returns *[]*GlueCrawlerMongodbTarget
	_jsii_.Get(
		j,
		"mongodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RecrawlPolicy() GlueCrawlerRecrawlPolicyOutputReference {
	var returns GlueCrawlerRecrawlPolicyOutputReference
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RecrawlPolicyInput() *GlueCrawlerRecrawlPolicy {
	var returns *GlueCrawlerRecrawlPolicy
	_jsii_.Get(
		j,
		"recrawlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) S3Target() *[]*GlueCrawlerS3Target {
	var returns *[]*GlueCrawlerS3Target
	_jsii_.Get(
		j,
		"s3Target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) S3TargetInput() *[]*GlueCrawlerS3Target {
	var returns *[]*GlueCrawlerS3Target
	_jsii_.Get(
		j,
		"s3TargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SchemaChangePolicy() GlueCrawlerSchemaChangePolicyOutputReference {
	var returns GlueCrawlerSchemaChangePolicyOutputReference
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SchemaChangePolicyInput() *GlueCrawlerSchemaChangePolicy {
	var returns *GlueCrawlerSchemaChangePolicy
	_jsii_.Get(
		j,
		"schemaChangePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawler) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html aws_glue_crawler} Resource.
func NewGlueCrawler(scope constructs.Construct, id *string, config *GlueCrawlerConfig) GlueCrawler {
	_init_.Initialize()

	j := jsiiProxy_GlueCrawler{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html aws_glue_crawler} Resource.
func NewGlueCrawler_Override(g GlueCrawler, scope constructs.Construct, id *string, config *GlueCrawlerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawler",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueCrawler) SetCatalogTarget(val *[]*GlueCrawlerCatalogTarget) {
	_jsii_.Set(
		j,
		"catalogTarget",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetClassifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetConfiguration(val *string) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetDynamodbTarget(val *[]*GlueCrawlerDynamodbTarget) {
	_jsii_.Set(
		j,
		"dynamodbTarget",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetJdbcTarget(val *[]*GlueCrawlerJdbcTarget) {
	_jsii_.Set(
		j,
		"jdbcTarget",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetMongodbTarget(val *[]*GlueCrawlerMongodbTarget) {
	_jsii_.Set(
		j,
		"mongodbTarget",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetS3Target(val *[]*GlueCrawlerS3Target) {
	_jsii_.Set(
		j,
		"s3Target",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetTablePrefix(val *string) {
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueCrawler) SetTagsAll(val interface{}) {
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
func GlueCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueCrawler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueCrawler",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueCrawler) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueCrawler) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCrawler) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCrawler) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCrawler) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCrawler) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCrawler) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueCrawler) PutLineageConfiguration(value *GlueCrawlerLineageConfiguration) {
	_jsii_.InvokeVoid(
		g,
		"putLineageConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutRecrawlPolicy(value *GlueCrawlerRecrawlPolicy) {
	_jsii_.InvokeVoid(
		g,
		"putRecrawlPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) PutSchemaChangePolicy(value *GlueCrawlerSchemaChangePolicy) {
	_jsii_.InvokeVoid(
		g,
		"putSchemaChangePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawler) ResetCatalogTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetClassifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetClassifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetDynamodbTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamodbTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetJdbcTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetJdbcTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetLineageConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetLineageConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetMongodbTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetMongodbTarget",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueCrawler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetRecrawlPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRecrawlPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetS3Target() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Target",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSchemaChangePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaChangePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawler) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueCrawler) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueCrawler) ToString() *string {
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
func (g *jsiiProxy_GlueCrawler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueCrawlerCatalogTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#database_name GlueCrawler#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#tables GlueCrawler#tables}.
	Tables *[]*string `json:"tables"`
}

type GlueCrawlerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#database_name GlueCrawler#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#name GlueCrawler#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#role GlueCrawler#role}.
	Role *string `json:"role"`
	// catalog_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#catalog_target GlueCrawler#catalog_target}
	CatalogTarget *[]*GlueCrawlerCatalogTarget `json:"catalogTarget"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#classifiers GlueCrawler#classifiers}.
	Classifiers *[]*string `json:"classifiers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#configuration GlueCrawler#configuration}.
	Configuration *string `json:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#description GlueCrawler#description}.
	Description *string `json:"description"`
	// dynamodb_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#dynamodb_target GlueCrawler#dynamodb_target}
	DynamodbTarget *[]*GlueCrawlerDynamodbTarget `json:"dynamodbTarget"`
	// jdbc_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#jdbc_target GlueCrawler#jdbc_target}
	JdbcTarget *[]*GlueCrawlerJdbcTarget `json:"jdbcTarget"`
	// lineage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#lineage_configuration GlueCrawler#lineage_configuration}
	LineageConfiguration *GlueCrawlerLineageConfiguration `json:"lineageConfiguration"`
	// mongodb_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#mongodb_target GlueCrawler#mongodb_target}
	MongodbTarget *[]*GlueCrawlerMongodbTarget `json:"mongodbTarget"`
	// recrawl_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#recrawl_policy GlueCrawler#recrawl_policy}
	RecrawlPolicy *GlueCrawlerRecrawlPolicy `json:"recrawlPolicy"`
	// s3_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#s3_target GlueCrawler#s3_target}
	S3Target *[]*GlueCrawlerS3Target `json:"s3Target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#schedule GlueCrawler#schedule}.
	Schedule *string `json:"schedule"`
	// schema_change_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#schema_change_policy GlueCrawler#schema_change_policy}
	SchemaChangePolicy *GlueCrawlerSchemaChangePolicy `json:"schemaChangePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#security_configuration GlueCrawler#security_configuration}.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#table_prefix GlueCrawler#table_prefix}.
	TablePrefix *string `json:"tablePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#tags GlueCrawler#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#tags_all GlueCrawler#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type GlueCrawlerDynamodbTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#path GlueCrawler#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#scan_all GlueCrawler#scan_all}.
	ScanAll interface{} `json:"scanAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#scan_rate GlueCrawler#scan_rate}.
	ScanRate *float64 `json:"scanRate"`
}

type GlueCrawlerJdbcTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `json:"connectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#path GlueCrawler#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#exclusions GlueCrawler#exclusions}.
	Exclusions *[]*string `json:"exclusions"`
}

type GlueCrawlerLineageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#crawler_lineage_settings GlueCrawler#crawler_lineage_settings}.
	CrawlerLineageSettings *string `json:"crawlerLineageSettings"`
}

type GlueCrawlerLineageConfigurationOutputReference interface {
	cdktf.ComplexObject
	CrawlerLineageSettings() *string
	SetCrawlerLineageSettings(val *string)
	CrawlerLineageSettingsInput() *string
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
	ResetCrawlerLineageSettings()
}

// The jsii proxy struct for GlueCrawlerLineageConfigurationOutputReference
type jsiiProxy_GlueCrawlerLineageConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) CrawlerLineageSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerLineageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) CrawlerLineageSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerLineageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCrawlerLineageConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCrawlerLineageConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCrawlerLineageConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerLineageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCrawlerLineageConfigurationOutputReference_Override(g GlueCrawlerLineageConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerLineageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) SetCrawlerLineageSettings(val *string) {
	_jsii_.Set(
		j,
		"crawlerLineageSettings",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerLineageConfigurationOutputReference) ResetCrawlerLineageSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCrawlerLineageSettings",
		nil, // no parameters
	)
}

type GlueCrawlerMongodbTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `json:"connectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#path GlueCrawler#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#scan_all GlueCrawler#scan_all}.
	ScanAll interface{} `json:"scanAll"`
}

type GlueCrawlerRecrawlPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#recrawl_behavior GlueCrawler#recrawl_behavior}.
	RecrawlBehavior *string `json:"recrawlBehavior"`
}

type GlueCrawlerRecrawlPolicyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RecrawlBehavior() *string
	SetRecrawlBehavior(val *string)
	RecrawlBehaviorInput() *string
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
	ResetRecrawlBehavior()
}

// The jsii proxy struct for GlueCrawlerRecrawlPolicyOutputReference
type jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) RecrawlBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recrawlBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) RecrawlBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recrawlBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueCrawlerRecrawlPolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCrawlerRecrawlPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerRecrawlPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCrawlerRecrawlPolicyOutputReference_Override(g GlueCrawlerRecrawlPolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerRecrawlPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) SetRecrawlBehavior(val *string) {
	_jsii_.Set(
		j,
		"recrawlBehavior",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerRecrawlPolicyOutputReference) ResetRecrawlBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetRecrawlBehavior",
		nil, // no parameters
	)
}

type GlueCrawlerS3Target struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#path GlueCrawler#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `json:"connectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#dlq_event_queue_arn GlueCrawler#dlq_event_queue_arn}.
	DlqEventQueueArn *string `json:"dlqEventQueueArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#event_queue_arn GlueCrawler#event_queue_arn}.
	EventQueueArn *string `json:"eventQueueArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#exclusions GlueCrawler#exclusions}.
	Exclusions *[]*string `json:"exclusions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#sample_size GlueCrawler#sample_size}.
	SampleSize *float64 `json:"sampleSize"`
}

type GlueCrawlerSchemaChangePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#delete_behavior GlueCrawler#delete_behavior}.
	DeleteBehavior *string `json:"deleteBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler.html#update_behavior GlueCrawler#update_behavior}.
	UpdateBehavior *string `json:"updateBehavior"`
}

type GlueCrawlerSchemaChangePolicyOutputReference interface {
	cdktf.ComplexObject
	DeleteBehavior() *string
	SetDeleteBehavior(val *string)
	DeleteBehaviorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UpdateBehavior() *string
	SetUpdateBehavior(val *string)
	UpdateBehaviorInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDeleteBehavior()
	ResetUpdateBehavior()
}

// The jsii proxy struct for GlueCrawlerSchemaChangePolicyOutputReference
type jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) DeleteBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) DeleteBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) UpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) UpdateBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateBehaviorInput",
		&returns,
	)
	return returns
}

func NewGlueCrawlerSchemaChangePolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueCrawlerSchemaChangePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerSchemaChangePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueCrawlerSchemaChangePolicyOutputReference_Override(g GlueCrawlerSchemaChangePolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueCrawlerSchemaChangePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) SetDeleteBehavior(val *string) {
	_jsii_.Set(
		j,
		"deleteBehavior",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) SetUpdateBehavior(val *string) {
	_jsii_.Set(
		j,
		"updateBehavior",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) ResetDeleteBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteBehavior",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerSchemaChangePolicyOutputReference) ResetUpdateBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdateBehavior",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings}.
type GlueDataCatalogEncryptionSettings interface {
	cdktf.TerraformResource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DataCatalogEncryptionSettings() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference
	DataCatalogEncryptionSettingsInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings
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
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDataCatalogEncryptionSettings(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings)
	ResetCatalogId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueDataCatalogEncryptionSettings
type jsiiProxy_GlueDataCatalogEncryptionSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) DataCatalogEncryptionSettings() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference {
	var returns GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference
	_jsii_.Get(
		j,
		"dataCatalogEncryptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) DataCatalogEncryptionSettingsInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings {
	var returns *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings
	_jsii_.Get(
		j,
		"dataCatalogEncryptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings} Resource.
func NewGlueDataCatalogEncryptionSettings(scope constructs.Construct, id *string, config *GlueDataCatalogEncryptionSettingsConfig) GlueDataCatalogEncryptionSettings {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogEncryptionSettings{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html aws_glue_data_catalog_encryption_settings} Resource.
func NewGlueDataCatalogEncryptionSettings_Override(g GlueDataCatalogEncryptionSettings, scope constructs.Construct, id *string, config *GlueDataCatalogEncryptionSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettings",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettings) SetProvider(val cdktf.TerraformProvider) {
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
func GlueDataCatalogEncryptionSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueDataCatalogEncryptionSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) PutDataCatalogEncryptionSettings(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings) {
	_jsii_.InvokeVoid(
		g,
		"putDataCatalogEncryptionSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) ToString() *string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueDataCatalogEncryptionSettingsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// data_catalog_encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#data_catalog_encryption_settings GlueDataCatalogEncryptionSettings#data_catalog_encryption_settings}
	DataCatalogEncryptionSettings *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings `json:"dataCatalogEncryptionSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#catalog_id GlueDataCatalogEncryptionSettings#catalog_id}.
	CatalogId *string `json:"catalogId"`
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings struct {
	// connection_password_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#connection_password_encryption GlueDataCatalogEncryptionSettings#connection_password_encryption}
	ConnectionPasswordEncryption *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption `json:"connectionPasswordEncryption"`
	// encryption_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#encryption_at_rest GlueDataCatalogEncryptionSettings#encryption_at_rest}
	EncryptionAtRest *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest `json:"encryptionAtRest"`
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#return_connection_password_encrypted GlueDataCatalogEncryptionSettings#return_connection_password_encrypted}.
	ReturnConnectionPasswordEncrypted interface{} `json:"returnConnectionPasswordEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#aws_kms_key_id GlueDataCatalogEncryptionSettings#aws_kms_key_id}.
	AwsKmsKeyId *string `json:"awsKmsKeyId"`
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference interface {
	cdktf.ComplexObject
	AwsKmsKeyId() *string
	SetAwsKmsKeyId(val *string)
	AwsKmsKeyIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ReturnConnectionPasswordEncrypted() interface{}
	SetReturnConnectionPasswordEncrypted(val interface{})
	ReturnConnectionPasswordEncryptedInput() interface{}
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
	ResetAwsKmsKeyId()
}

// The jsii proxy struct for GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference
type jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) AwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) AwsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) ReturnConnectionPasswordEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnConnectionPasswordEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) ReturnConnectionPasswordEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnConnectionPasswordEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference_Override(g GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) SetAwsKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"awsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) SetReturnConnectionPasswordEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"returnConnectionPasswordEncrypted",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference) ResetAwsKmsKeyId() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsKmsKeyId",
		nil, // no parameters
	)
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#catalog_encryption_mode GlueDataCatalogEncryptionSettings#catalog_encryption_mode}.
	CatalogEncryptionMode *string `json:"catalogEncryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings.html#sse_aws_kms_key_id GlueDataCatalogEncryptionSettings#sse_aws_kms_key_id}.
	SseAwsKmsKeyId *string `json:"sseAwsKmsKeyId"`
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference interface {
	cdktf.ComplexObject
	CatalogEncryptionMode() *string
	SetCatalogEncryptionMode(val *string)
	CatalogEncryptionModeInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SseAwsKmsKeyId() *string
	SetSseAwsKmsKeyId(val *string)
	SseAwsKmsKeyIdInput() *string
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
	ResetSseAwsKmsKeyId()
}

// The jsii proxy struct for GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference
type jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) CatalogEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) CatalogEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SseAwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SseAwsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference_Override(g GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SetCatalogEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"catalogEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SetSseAwsKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"sseAwsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference) ResetSseAwsKmsKeyId() {
	_jsii_.InvokeVoid(
		g,
		"resetSseAwsKmsKeyId",
		nil, // no parameters
	)
}

type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference interface {
	cdktf.ComplexObject
	ConnectionPasswordEncryption() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference
	ConnectionPasswordEncryptionInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption
	EncryptionAtRest() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference
	EncryptionAtRestInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest
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
	PutConnectionPasswordEncryption(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption)
	PutEncryptionAtRest(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest)
}

// The jsii proxy struct for GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference
type jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) ConnectionPasswordEncryption() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference {
	var returns GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionOutputReference
	_jsii_.Get(
		j,
		"connectionPasswordEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) ConnectionPasswordEncryptionInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption {
	var returns *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption
	_jsii_.Get(
		j,
		"connectionPasswordEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) EncryptionAtRest() GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference {
	var returns GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestOutputReference
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) EncryptionAtRestInput() *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest {
	var returns *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest
	_jsii_.Get(
		j,
		"encryptionAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference_Override(g GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) PutConnectionPasswordEncryption(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption) {
	_jsii_.InvokeVoid(
		g,
		"putConnectionPasswordEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsOutputReference) PutEncryptionAtRest(value *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest) {
	_jsii_.InvokeVoid(
		g,
		"putEncryptionAtRest",
		[]interface{}{value},
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html aws_glue_dev_endpoint}.
type GlueDevEndpoint interface {
	cdktf.TerraformResource
	Arguments() interface{}
	SetArguments(val interface{})
	ArgumentsInput() interface{}
	Arn() *string
	AvailabilityZone() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExtraJarsS3Path() *string
	SetExtraJarsS3Path(val *string)
	ExtraJarsS3PathInput() *string
	ExtraPythonLibsS3Path() *string
	SetExtraPythonLibsS3Path(val *string)
	ExtraPythonLibsS3PathInput() *string
	FailureReason() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GlueVersion() *string
	SetGlueVersion(val *string)
	GlueVersionInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfNodesInput() *float64
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	PrivateAddress() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PublicAddress() *string
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	PublicKeys() *[]*string
	SetPublicKeys(val *[]*string)
	PublicKeysInput() *[]*string
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcId() *string
	WorkerType() *string
	SetWorkerType(val *string)
	WorkerTypeInput() *string
	YarnEndpointAddress() *string
	ZeppelinRemoteSparkInterpreterPort() *float64
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetArguments()
	ResetExtraJarsS3Path()
	ResetExtraPythonLibsS3Path()
	ResetGlueVersion()
	ResetNumberOfNodes()
	ResetNumberOfWorkers()
	ResetOverrideLogicalId()
	ResetPublicKey()
	ResetPublicKeys()
	ResetSecurityConfiguration()
	ResetSecurityGroupIds()
	ResetSubnetId()
	ResetTags()
	ResetTagsAll()
	ResetWorkerType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueDevEndpoint
type jsiiProxy_GlueDevEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueDevEndpoint) Arguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ExtraJarsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraJarsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ExtraJarsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraJarsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ExtraPythonLibsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraPythonLibsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ExtraPythonLibsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraPythonLibsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PrivateAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PublicAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) PublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) YarnEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yarnEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDevEndpoint) ZeppelinRemoteSparkInterpreterPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zeppelinRemoteSparkInterpreterPort",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html aws_glue_dev_endpoint} Resource.
func NewGlueDevEndpoint(scope constructs.Construct, id *string, config *GlueDevEndpointConfig) GlueDevEndpoint {
	_init_.Initialize()

	j := jsiiProxy_GlueDevEndpoint{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDevEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html aws_glue_dev_endpoint} Resource.
func NewGlueDevEndpoint_Override(g GlueDevEndpoint, scope constructs.Construct, id *string, config *GlueDevEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueDevEndpoint",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetArguments(val interface{}) {
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetExtraJarsS3Path(val *string) {
	_jsii_.Set(
		j,
		"extraJarsS3Path",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetExtraPythonLibsS3Path(val *string) {
	_jsii_.Set(
		j,
		"extraPythonLibsS3Path",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetNumberOfNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetPublicKey(val *string) {
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetPublicKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"publicKeys",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_GlueDevEndpoint) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueDevEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueDevEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueDevEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueDevEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueDevEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueDevEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDevEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueDevEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueDevEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueDevEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueDevEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetArguments() {
	_jsii_.InvokeVoid(
		g,
		"resetArguments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetExtraJarsS3Path() {
	_jsii_.InvokeVoid(
		g,
		"resetExtraJarsS3Path",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetExtraPythonLibsS3Path() {
	_jsii_.InvokeVoid(
		g,
		"resetExtraPythonLibsS3Path",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueDevEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetPublicKey() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetPublicKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetSubnetId() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) ResetWorkerType() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDevEndpoint) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueDevEndpoint) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueDevEndpoint) ToString() *string {
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
func (g *jsiiProxy_GlueDevEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueDevEndpointConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#name GlueDevEndpoint#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#role_arn GlueDevEndpoint#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#arguments GlueDevEndpoint#arguments}.
	Arguments interface{} `json:"arguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#extra_jars_s3_path GlueDevEndpoint#extra_jars_s3_path}.
	ExtraJarsS3Path *string `json:"extraJarsS3Path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#extra_python_libs_s3_path GlueDevEndpoint#extra_python_libs_s3_path}.
	ExtraPythonLibsS3Path *string `json:"extraPythonLibsS3Path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#glue_version GlueDevEndpoint#glue_version}.
	GlueVersion *string `json:"glueVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#number_of_nodes GlueDevEndpoint#number_of_nodes}.
	NumberOfNodes *float64 `json:"numberOfNodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#number_of_workers GlueDevEndpoint#number_of_workers}.
	NumberOfWorkers *float64 `json:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#public_key GlueDevEndpoint#public_key}.
	PublicKey *string `json:"publicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#public_keys GlueDevEndpoint#public_keys}.
	PublicKeys *[]*string `json:"publicKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#security_configuration GlueDevEndpoint#security_configuration}.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#security_group_ids GlueDevEndpoint#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#subnet_id GlueDevEndpoint#subnet_id}.
	SubnetId *string `json:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#tags GlueDevEndpoint#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#tags_all GlueDevEndpoint#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_dev_endpoint.html#worker_type GlueDevEndpoint#worker_type}.
	WorkerType *string `json:"workerType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html aws_glue_job}.
type GlueJob interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Command() GlueJobCommandOutputReference
	CommandInput() *GlueJobCommand
	Connections() *[]*string
	SetConnections(val *[]*string)
	ConnectionsInput() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultArguments() interface{}
	SetDefaultArguments(val interface{})
	DefaultArgumentsInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExecutionProperty() GlueJobExecutionPropertyOutputReference
	ExecutionPropertyInput() *GlueJobExecutionProperty
	Fqn() *string
	FriendlyUniqueId() *string
	GlueVersion() *string
	SetGlueVersion(val *string)
	GlueVersionInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	NonOverridableArguments() interface{}
	SetNonOverridableArguments(val interface{})
	NonOverridableArgumentsInput() interface{}
	NotificationProperty() GlueJobNotificationPropertyOutputReference
	NotificationPropertyInput() *GlueJobNotificationProperty
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	WorkerType() *string
	SetWorkerType(val *string)
	WorkerTypeInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCommand(value *GlueJobCommand)
	PutExecutionProperty(value *GlueJobExecutionProperty)
	PutNotificationProperty(value *GlueJobNotificationProperty)
	ResetConnections()
	ResetDefaultArguments()
	ResetDescription()
	ResetExecutionProperty()
	ResetGlueVersion()
	ResetMaxCapacity()
	ResetMaxRetries()
	ResetNonOverridableArguments()
	ResetNotificationProperty()
	ResetNumberOfWorkers()
	ResetOverrideLogicalId()
	ResetSecurityConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
	ResetWorkerType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueJob
type jsiiProxy_GlueJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Command() GlueJobCommandOutputReference {
	var returns GlueJobCommandOutputReference
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) CommandInput() *GlueJobCommand {
	var returns *GlueJobCommand
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Connections() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) ConnectionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) DefaultArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) DefaultArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) ExecutionProperty() GlueJobExecutionPropertyOutputReference {
	var returns GlueJobExecutionPropertyOutputReference
	_jsii_.Get(
		j,
		"executionProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) ExecutionPropertyInput() *GlueJobExecutionProperty {
	var returns *GlueJobExecutionProperty
	_jsii_.Get(
		j,
		"executionPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NonOverridableArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonOverridableArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NonOverridableArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonOverridableArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NotificationProperty() GlueJobNotificationPropertyOutputReference {
	var returns GlueJobNotificationPropertyOutputReference
	_jsii_.Get(
		j,
		"notificationProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NotificationPropertyInput() *GlueJobNotificationProperty {
	var returns *GlueJobNotificationProperty
	_jsii_.Get(
		j,
		"notificationPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJob) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html aws_glue_job} Resource.
func NewGlueJob(scope constructs.Construct, id *string, config *GlueJobConfig) GlueJob {
	_init_.Initialize()

	j := jsiiProxy_GlueJob{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html aws_glue_job} Resource.
func NewGlueJob_Override(g GlueJob, scope constructs.Construct, id *string, config *GlueJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJob",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueJob) SetConnections(val *[]*string) {
	_jsii_.Set(
		j,
		"connections",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetDefaultArguments(val interface{}) {
	_jsii_.Set(
		j,
		"defaultArguments",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetNonOverridableArguments(val interface{}) {
	_jsii_.Set(
		j,
		"nonOverridableArguments",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GlueJob) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueJob) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueJob) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueJob) PutCommand(value *GlueJobCommand) {
	_jsii_.InvokeVoid(
		g,
		"putCommand",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueJob) PutExecutionProperty(value *GlueJobExecutionProperty) {
	_jsii_.InvokeVoid(
		g,
		"putExecutionProperty",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueJob) PutNotificationProperty(value *GlueJobNotificationProperty) {
	_jsii_.InvokeVoid(
		g,
		"putNotificationProperty",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueJob) ResetConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetDefaultArguments() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultArguments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetExecutionProperty() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionProperty",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetNonOverridableArguments() {
	_jsii_.InvokeVoid(
		g,
		"resetNonOverridableArguments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetNotificationProperty() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationProperty",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) ResetWorkerType() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJob) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueJob) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueJob) ToString() *string {
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
func (g *jsiiProxy_GlueJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueJobCommand struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#script_location GlueJob#script_location}.
	ScriptLocation *string `json:"scriptLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#name GlueJob#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#python_version GlueJob#python_version}.
	PythonVersion *string `json:"pythonVersion"`
}

type GlueJobCommandOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	ScriptLocation() *string
	SetScriptLocation(val *string)
	ScriptLocationInput() *string
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
	ResetName()
	ResetPythonVersion()
}

// The jsii proxy struct for GlueJobCommandOutputReference
type jsiiProxy_GlueJobCommandOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueJobCommandOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) ScriptLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) ScriptLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobCommandOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueJobCommandOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueJobCommandOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueJobCommandOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobCommandOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueJobCommandOutputReference_Override(g GlueJobCommandOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobCommandOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetPythonVersion(val *string) {
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetScriptLocation(val *string) {
	_jsii_.Set(
		j,
		"scriptLocation",
		val,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueJobCommandOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueJobCommandOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueJobCommandOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueJobCommandOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueJobCommandOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueJobCommandOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueJobCommandOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobCommandOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueJobCommandOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonVersion",
		nil, // no parameters
	)
}

type GlueJobConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// command block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#command GlueJob#command}
	Command *GlueJobCommand `json:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#name GlueJob#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#role_arn GlueJob#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#connections GlueJob#connections}.
	Connections *[]*string `json:"connections"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#default_arguments GlueJob#default_arguments}.
	DefaultArguments interface{} `json:"defaultArguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#description GlueJob#description}.
	Description *string `json:"description"`
	// execution_property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#execution_property GlueJob#execution_property}
	ExecutionProperty *GlueJobExecutionProperty `json:"executionProperty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#glue_version GlueJob#glue_version}.
	GlueVersion *string `json:"glueVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#max_capacity GlueJob#max_capacity}.
	MaxCapacity *float64 `json:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#max_retries GlueJob#max_retries}.
	MaxRetries *float64 `json:"maxRetries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#non_overridable_arguments GlueJob#non_overridable_arguments}.
	NonOverridableArguments interface{} `json:"nonOverridableArguments"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#notification_property GlueJob#notification_property}
	NotificationProperty *GlueJobNotificationProperty `json:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#number_of_workers GlueJob#number_of_workers}.
	NumberOfWorkers *float64 `json:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#security_configuration GlueJob#security_configuration}.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#tags GlueJob#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#tags_all GlueJob#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#timeout GlueJob#timeout}.
	Timeout *float64 `json:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#worker_type GlueJob#worker_type}.
	WorkerType *string `json:"workerType"`
}

type GlueJobExecutionProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#max_concurrent_runs GlueJob#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns"`
}

type GlueJobExecutionPropertyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxConcurrentRuns() *float64
	SetMaxConcurrentRuns(val *float64)
	MaxConcurrentRunsInput() *float64
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
	ResetMaxConcurrentRuns()
}

// The jsii proxy struct for GlueJobExecutionPropertyOutputReference
type jsiiProxy_GlueJobExecutionPropertyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) MaxConcurrentRuns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) MaxConcurrentRunsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueJobExecutionPropertyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueJobExecutionPropertyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueJobExecutionPropertyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobExecutionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueJobExecutionPropertyOutputReference_Override(g GlueJobExecutionPropertyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobExecutionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) SetMaxConcurrentRuns(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentRuns",
		val,
	)
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueJobExecutionPropertyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobExecutionPropertyOutputReference) ResetMaxConcurrentRuns() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentRuns",
		nil, // no parameters
	)
}

type GlueJobNotificationProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_job.html#notify_delay_after GlueJob#notify_delay_after}.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter"`
}

type GlueJobNotificationPropertyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NotifyDelayAfter() *float64
	SetNotifyDelayAfter(val *float64)
	NotifyDelayAfterInput() *float64
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
	ResetNotifyDelayAfter()
}

// The jsii proxy struct for GlueJobNotificationPropertyOutputReference
type jsiiProxy_GlueJobNotificationPropertyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) NotifyDelayAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notifyDelayAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) NotifyDelayAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notifyDelayAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueJobNotificationPropertyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueJobNotificationPropertyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueJobNotificationPropertyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobNotificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueJobNotificationPropertyOutputReference_Override(g GlueJobNotificationPropertyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueJobNotificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) SetNotifyDelayAfter(val *float64) {
	_jsii_.Set(
		j,
		"notifyDelayAfter",
		val,
	)
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueJobNotificationPropertyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobNotificationPropertyOutputReference) ResetNotifyDelayAfter() {
	_jsii_.InvokeVoid(
		g,
		"resetNotifyDelayAfter",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html aws_glue_ml_transform}.
type GlueMlTransform interface {
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
	GlueVersion() *string
	SetGlueVersion(val *string)
	GlueVersionInput() *string
	Id() *string
	InputRecordTables() *[]*GlueMlTransformInputRecordTables
	SetInputRecordTables(val *[]*GlueMlTransformInputRecordTables)
	InputRecordTablesInput() *[]*GlueMlTransformInputRecordTables
	LabelCount() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	Parameters() GlueMlTransformParametersOutputReference
	ParametersInput() *GlueMlTransformParameters
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	WorkerType() *string
	SetWorkerType(val *string)
	WorkerTypeInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutParameters(value *GlueMlTransformParameters)
	ResetDescription()
	ResetGlueVersion()
	ResetMaxCapacity()
	ResetMaxRetries()
	ResetNumberOfWorkers()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
	ResetWorkerType()
	Schema(index *string) GlueMlTransformSchema
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueMlTransform
type jsiiProxy_GlueMlTransform struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueMlTransform) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) InputRecordTables() *[]*GlueMlTransformInputRecordTables {
	var returns *[]*GlueMlTransformInputRecordTables
	_jsii_.Get(
		j,
		"inputRecordTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) InputRecordTablesInput() *[]*GlueMlTransformInputRecordTables {
	var returns *[]*GlueMlTransformInputRecordTables
	_jsii_.Get(
		j,
		"inputRecordTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) LabelCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"labelCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Parameters() GlueMlTransformParametersOutputReference {
	var returns GlueMlTransformParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) ParametersInput() *GlueMlTransformParameters {
	var returns *GlueMlTransformParameters
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransform) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html aws_glue_ml_transform} Resource.
func NewGlueMlTransform(scope constructs.Construct, id *string, config *GlueMlTransformConfig) GlueMlTransform {
	_init_.Initialize()

	j := jsiiProxy_GlueMlTransform{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransform",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html aws_glue_ml_transform} Resource.
func NewGlueMlTransform_Override(g GlueMlTransform, scope constructs.Construct, id *string, config *GlueMlTransformConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransform",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetInputRecordTables(val *[]*GlueMlTransformInputRecordTables) {
	_jsii_.Set(
		j,
		"inputRecordTables",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransform) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueMlTransform_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueMlTransform",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueMlTransform_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueMlTransform",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueMlTransform) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueMlTransform) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueMlTransform) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueMlTransform) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueMlTransform) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueMlTransform) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueMlTransform) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueMlTransform) PutParameters(value *GlueMlTransformParameters) {
	_jsii_.InvokeVoid(
		g,
		"putParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueMlTransform) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) ResetWorkerType() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransform) Schema(index *string) GlueMlTransformSchema {
	var returns GlueMlTransformSchema

	_jsii_.Invoke(
		g,
		"schema",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueMlTransform) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueMlTransform) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueMlTransform) ToString() *string {
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
func (g *jsiiProxy_GlueMlTransform) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueMlTransformConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// input_record_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#input_record_tables GlueMlTransform#input_record_tables}
	InputRecordTables *[]*GlueMlTransformInputRecordTables `json:"inputRecordTables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#name GlueMlTransform#name}.
	Name *string `json:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#parameters GlueMlTransform#parameters}
	Parameters *GlueMlTransformParameters `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#role_arn GlueMlTransform#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#description GlueMlTransform#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#glue_version GlueMlTransform#glue_version}.
	GlueVersion *string `json:"glueVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#max_capacity GlueMlTransform#max_capacity}.
	MaxCapacity *float64 `json:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#max_retries GlueMlTransform#max_retries}.
	MaxRetries *float64 `json:"maxRetries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#number_of_workers GlueMlTransform#number_of_workers}.
	NumberOfWorkers *float64 `json:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#tags GlueMlTransform#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#tags_all GlueMlTransform#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#timeout GlueMlTransform#timeout}.
	Timeout *float64 `json:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#worker_type GlueMlTransform#worker_type}.
	WorkerType *string `json:"workerType"`
}

type GlueMlTransformInputRecordTables struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#database_name GlueMlTransform#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#table_name GlueMlTransform#table_name}.
	TableName *string `json:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#catalog_id GlueMlTransform#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#connection_name GlueMlTransform#connection_name}.
	ConnectionName *string `json:"connectionName"`
}

type GlueMlTransformParameters struct {
	// find_matches_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#find_matches_parameters GlueMlTransform#find_matches_parameters}
	FindMatchesParameters *GlueMlTransformParametersFindMatchesParameters `json:"findMatchesParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#transform_type GlueMlTransform#transform_type}.
	TransformType *string `json:"transformType"`
}

type GlueMlTransformParametersFindMatchesParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#accuracy_cost_trade_off GlueMlTransform#accuracy_cost_trade_off}.
	AccuracyCostTradeOff *float64 `json:"accuracyCostTradeOff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#enforce_provided_labels GlueMlTransform#enforce_provided_labels}.
	EnforceProvidedLabels interface{} `json:"enforceProvidedLabels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#precision_recall_trade_off GlueMlTransform#precision_recall_trade_off}.
	PrecisionRecallTradeOff *float64 `json:"precisionRecallTradeOff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_ml_transform.html#primary_key_column_name GlueMlTransform#primary_key_column_name}.
	PrimaryKeyColumnName *string `json:"primaryKeyColumnName"`
}

type GlueMlTransformParametersFindMatchesParametersOutputReference interface {
	cdktf.ComplexObject
	AccuracyCostTradeOff() *float64
	SetAccuracyCostTradeOff(val *float64)
	AccuracyCostTradeOffInput() *float64
	EnforceProvidedLabels() interface{}
	SetEnforceProvidedLabels(val interface{})
	EnforceProvidedLabelsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PrecisionRecallTradeOff() *float64
	SetPrecisionRecallTradeOff(val *float64)
	PrecisionRecallTradeOffInput() *float64
	PrimaryKeyColumnName() *string
	SetPrimaryKeyColumnName(val *string)
	PrimaryKeyColumnNameInput() *string
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
	ResetAccuracyCostTradeOff()
	ResetEnforceProvidedLabels()
	ResetPrecisionRecallTradeOff()
	ResetPrimaryKeyColumnName()
}

// The jsii proxy struct for GlueMlTransformParametersFindMatchesParametersOutputReference
type jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) AccuracyCostTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) AccuracyCostTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) EnforceProvidedLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) EnforceProvidedLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrecisionRecallTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrecisionRecallTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrimaryKeyColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrimaryKeyColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueMlTransformParametersFindMatchesParametersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueMlTransformParametersFindMatchesParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformParametersFindMatchesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueMlTransformParametersFindMatchesParametersOutputReference_Override(g GlueMlTransformParametersFindMatchesParametersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformParametersFindMatchesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetAccuracyCostTradeOff(val *float64) {
	_jsii_.Set(
		j,
		"accuracyCostTradeOff",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetEnforceProvidedLabels(val interface{}) {
	_jsii_.Set(
		j,
		"enforceProvidedLabels",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetPrecisionRecallTradeOff(val *float64) {
	_jsii_.Set(
		j,
		"precisionRecallTradeOff",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetPrimaryKeyColumnName(val *string) {
	_jsii_.Set(
		j,
		"primaryKeyColumnName",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetAccuracyCostTradeOff() {
	_jsii_.InvokeVoid(
		g,
		"resetAccuracyCostTradeOff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetEnforceProvidedLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceProvidedLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetPrecisionRecallTradeOff() {
	_jsii_.InvokeVoid(
		g,
		"resetPrecisionRecallTradeOff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetPrimaryKeyColumnName() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryKeyColumnName",
		nil, // no parameters
	)
}

type GlueMlTransformParametersOutputReference interface {
	cdktf.ComplexObject
	FindMatchesParameters() GlueMlTransformParametersFindMatchesParametersOutputReference
	FindMatchesParametersInput() *GlueMlTransformParametersFindMatchesParameters
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TransformType() *string
	SetTransformType(val *string)
	TransformTypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutFindMatchesParameters(value *GlueMlTransformParametersFindMatchesParameters)
}

// The jsii proxy struct for GlueMlTransformParametersOutputReference
type jsiiProxy_GlueMlTransformParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) FindMatchesParameters() GlueMlTransformParametersFindMatchesParametersOutputReference {
	var returns GlueMlTransformParametersFindMatchesParametersOutputReference
	_jsii_.Get(
		j,
		"findMatchesParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) FindMatchesParametersInput() *GlueMlTransformParametersFindMatchesParameters {
	var returns *GlueMlTransformParametersFindMatchesParameters
	_jsii_.Get(
		j,
		"findMatchesParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) TransformType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) TransformTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformTypeInput",
		&returns,
	)
	return returns
}

func NewGlueMlTransformParametersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueMlTransformParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueMlTransformParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueMlTransformParametersOutputReference_Override(g GlueMlTransformParametersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersOutputReference) SetTransformType(val *string) {
	_jsii_.Set(
		j,
		"transformType",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueMlTransformParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueMlTransformParametersOutputReference) PutFindMatchesParameters(value *GlueMlTransformParametersFindMatchesParameters) {
	_jsii_.InvokeVoid(
		g,
		"putFindMatchesParameters",
		[]interface{}{value},
	)
}

type GlueMlTransformSchema interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DataType() *string
	Name() *string
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

// The jsii proxy struct for GlueMlTransformSchema
type jsiiProxy_GlueMlTransformSchema struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_GlueMlTransformSchema) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformSchema) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformSchema) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformSchema) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformSchema) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewGlueMlTransformSchema(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) GlueMlTransformSchema {
	_init_.Initialize()

	j := jsiiProxy_GlueMlTransformSchema{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformSchema",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueMlTransformSchema_Override(g GlueMlTransformSchema, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueMlTransformSchema",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		g,
	)
}

func (j *jsiiProxy_GlueMlTransformSchema) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformSchema) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformSchema) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueMlTransformSchema) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueMlTransformSchema) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueMlTransformSchema) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueMlTransformSchema) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueMlTransformSchema) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html aws_glue_partition}.
type GluePartition interface {
	cdktf.TerraformResource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreationTime() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastAccessedTime() *string
	LastAnalyzedTime() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	PartitionValues() *[]*string
	SetPartitionValues(val *[]*string)
	PartitionValuesInput() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StorageDescriptor() GluePartitionStorageDescriptorOutputReference
	StorageDescriptorInput() *GluePartitionStorageDescriptor
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	PutStorageDescriptor(value *GluePartitionStorageDescriptor)
	ResetCatalogId()
	ResetOverrideLogicalId()
	ResetParameters()
	ResetStorageDescriptor()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GluePartition
type jsiiProxy_GluePartition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GluePartition) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) LastAccessedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAccessedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) LastAnalyzedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAnalyzedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) PartitionValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) PartitionValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) StorageDescriptor() GluePartitionStorageDescriptorOutputReference {
	var returns GluePartitionStorageDescriptorOutputReference
	_jsii_.Get(
		j,
		"storageDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) StorageDescriptorInput() *GluePartitionStorageDescriptor {
	var returns *GluePartitionStorageDescriptor
	_jsii_.Get(
		j,
		"storageDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html aws_glue_partition} Resource.
func NewGluePartition(scope constructs.Construct, id *string, config *GluePartitionConfig) GluePartition {
	_init_.Initialize()

	j := jsiiProxy_GluePartition{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html aws_glue_partition} Resource.
func NewGluePartition_Override(g GluePartition, scope constructs.Construct, id *string, config *GluePartitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartition",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GluePartition) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetPartitionValues(val *[]*string) {
	_jsii_.Set(
		j,
		"partitionValues",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GluePartition) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GluePartition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GluePartition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GluePartition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GluePartition",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GluePartition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GluePartition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartition) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartition) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartition) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GluePartition) PutStorageDescriptor(value *GluePartitionStorageDescriptor) {
	_jsii_.InvokeVoid(
		g,
		"putStorageDescriptor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartition) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GluePartition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartition) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartition) ResetStorageDescriptor() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageDescriptor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartition) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GluePartition) ToMetadata() interface{} {
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
func (g *jsiiProxy_GluePartition) ToString() *string {
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
func (g *jsiiProxy_GluePartition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GluePartitionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#database_name GluePartition#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#partition_values GluePartition#partition_values}.
	PartitionValues *[]*string `json:"partitionValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#table_name GluePartition#table_name}.
	TableName *string `json:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#catalog_id GluePartition#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#parameters GluePartition#parameters}.
	Parameters interface{} `json:"parameters"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#storage_descriptor GluePartition#storage_descriptor}
	StorageDescriptor *GluePartitionStorageDescriptor `json:"storageDescriptor"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html aws_glue_partition_index}.
type GluePartitionIndex interface {
	cdktf.TerraformResource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PartitionIndex() GluePartitionIndexPartitionIndexOutputReference
	PartitionIndexInput() *GluePartitionIndexPartitionIndex
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	PutPartitionIndex(value *GluePartitionIndexPartitionIndex)
	ResetCatalogId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GluePartitionIndex
type jsiiProxy_GluePartitionIndex struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GluePartitionIndex) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) PartitionIndex() GluePartitionIndexPartitionIndexOutputReference {
	var returns GluePartitionIndexPartitionIndexOutputReference
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) PartitionIndexInput() *GluePartitionIndexPartitionIndex {
	var returns *GluePartitionIndexPartitionIndex
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html aws_glue_partition_index} Resource.
func NewGluePartitionIndex(scope constructs.Construct, id *string, config *GluePartitionIndexConfig) GluePartitionIndex {
	_init_.Initialize()

	j := jsiiProxy_GluePartitionIndex{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html aws_glue_partition_index} Resource.
func NewGluePartitionIndex_Override(g GluePartitionIndex, scope constructs.Construct, id *string, config *GluePartitionIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionIndex",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndex) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GluePartitionIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GluePartitionIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GluePartitionIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GluePartitionIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GluePartitionIndex) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GluePartitionIndex) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionIndex) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartitionIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartitionIndex) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartitionIndex) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionIndex) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GluePartitionIndex) PutPartitionIndex(value *GluePartitionIndexPartitionIndex) {
	_jsii_.InvokeVoid(
		g,
		"putPartitionIndex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionIndex) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GluePartitionIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionIndex) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GluePartitionIndex) ToMetadata() interface{} {
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
func (g *jsiiProxy_GluePartitionIndex) ToString() *string {
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
func (g *jsiiProxy_GluePartitionIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GluePartitionIndexConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#database_name GluePartitionIndex#database_name}.
	DatabaseName *string `json:"databaseName"`
	// partition_index block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#partition_index GluePartitionIndex#partition_index}
	PartitionIndex *GluePartitionIndexPartitionIndex `json:"partitionIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#table_name GluePartitionIndex#table_name}.
	TableName *string `json:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#catalog_id GluePartitionIndex#catalog_id}.
	CatalogId *string `json:"catalogId"`
}

type GluePartitionIndexPartitionIndex struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#index_name GluePartitionIndex#index_name}.
	IndexName *string `json:"indexName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index.html#keys GluePartitionIndex#keys}.
	Keys *[]*string `json:"keys"`
}

type GluePartitionIndexPartitionIndexOutputReference interface {
	cdktf.ComplexObject
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Keys() *[]*string
	SetKeys(val *[]*string)
	KeysInput() *[]*string
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
	ResetIndexName()
	ResetKeys()
}

// The jsii proxy struct for GluePartitionIndexPartitionIndexOutputReference
type jsiiProxy_GluePartitionIndexPartitionIndexOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) Keys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) KeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGluePartitionIndexPartitionIndexOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GluePartitionIndexPartitionIndexOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GluePartitionIndexPartitionIndexOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionIndexPartitionIndexOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGluePartitionIndexPartitionIndexOutputReference_Override(g GluePartitionIndexPartitionIndexOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionIndexPartitionIndexOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) SetIndexName(val *string) {
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) SetKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		g,
		"resetIndexName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionIndexPartitionIndexOutputReference) ResetKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetKeys",
		nil, // no parameters
	)
}

type GluePartitionStorageDescriptor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#bucket_columns GluePartition#bucket_columns}.
	BucketColumns *[]*string `json:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#columns GluePartition#columns}
	Columns *[]*GluePartitionStorageDescriptorColumns `json:"columns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#compressed GluePartition#compressed}.
	Compressed interface{} `json:"compressed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#input_format GluePartition#input_format}.
	InputFormat *string `json:"inputFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#location GluePartition#location}.
	Location *string `json:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#number_of_buckets GluePartition#number_of_buckets}.
	NumberOfBuckets *float64 `json:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#output_format GluePartition#output_format}.
	OutputFormat *string `json:"outputFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#parameters GluePartition#parameters}.
	Parameters interface{} `json:"parameters"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#ser_de_info GluePartition#ser_de_info}
	SerDeInfo *GluePartitionStorageDescriptorSerDeInfo `json:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#skewed_info GluePartition#skewed_info}
	SkewedInfo *GluePartitionStorageDescriptorSkewedInfo `json:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#sort_columns GluePartition#sort_columns}
	SortColumns *[]*GluePartitionStorageDescriptorSortColumns `json:"sortColumns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#stored_as_sub_directories GluePartition#stored_as_sub_directories}.
	StoredAsSubDirectories interface{} `json:"storedAsSubDirectories"`
}

type GluePartitionStorageDescriptorColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#name GluePartition#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#comment GluePartition#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#type GluePartition#type}.
	Type *string `json:"type"`
}

type GluePartitionStorageDescriptorOutputReference interface {
	cdktf.ComplexObject
	BucketColumns() *[]*string
	SetBucketColumns(val *[]*string)
	BucketColumnsInput() *[]*string
	Columns() *[]*GluePartitionStorageDescriptorColumns
	SetColumns(val *[]*GluePartitionStorageDescriptorColumns)
	ColumnsInput() *[]*GluePartitionStorageDescriptorColumns
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	NumberOfBuckets() *float64
	SetNumberOfBuckets(val *float64)
	NumberOfBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	SerDeInfo() GluePartitionStorageDescriptorSerDeInfoOutputReference
	SerDeInfoInput() *GluePartitionStorageDescriptorSerDeInfo
	SkewedInfo() GluePartitionStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() *GluePartitionStorageDescriptorSkewedInfo
	SortColumns() *[]*GluePartitionStorageDescriptorSortColumns
	SetSortColumns(val *[]*GluePartitionStorageDescriptorSortColumns)
	SortColumnsInput() *[]*GluePartitionStorageDescriptorSortColumns
	StoredAsSubDirectories() interface{}
	SetStoredAsSubDirectories(val interface{})
	StoredAsSubDirectoriesInput() interface{}
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
	PutSerDeInfo(value *GluePartitionStorageDescriptorSerDeInfo)
	PutSkewedInfo(value *GluePartitionStorageDescriptorSkewedInfo)
	ResetBucketColumns()
	ResetColumns()
	ResetCompressed()
	ResetInputFormat()
	ResetLocation()
	ResetNumberOfBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSerDeInfo()
	ResetSkewedInfo()
	ResetSortColumns()
	ResetStoredAsSubDirectories()
}

// The jsii proxy struct for GluePartitionStorageDescriptorOutputReference
type jsiiProxy_GluePartitionStorageDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) Columns() *[]*GluePartitionStorageDescriptorColumns {
	var returns *[]*GluePartitionStorageDescriptorColumns
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ColumnsInput() *[]*GluePartitionStorageDescriptorColumns {
	var returns *[]*GluePartitionStorageDescriptorColumns
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SerDeInfo() GluePartitionStorageDescriptorSerDeInfoOutputReference {
	var returns GluePartitionStorageDescriptorSerDeInfoOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SerDeInfoInput() *GluePartitionStorageDescriptorSerDeInfo {
	var returns *GluePartitionStorageDescriptorSerDeInfo
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SkewedInfo() GluePartitionStorageDescriptorSkewedInfoOutputReference {
	var returns GluePartitionStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SkewedInfoInput() *GluePartitionStorageDescriptorSkewedInfo {
	var returns *GluePartitionStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SortColumns() *[]*GluePartitionStorageDescriptorSortColumns {
	var returns *[]*GluePartitionStorageDescriptorSortColumns
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SortColumnsInput() *[]*GluePartitionStorageDescriptorSortColumns {
	var returns *[]*GluePartitionStorageDescriptorSortColumns
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGluePartitionStorageDescriptorOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GluePartitionStorageDescriptorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GluePartitionStorageDescriptorOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGluePartitionStorageDescriptorOutputReference_Override(g GluePartitionStorageDescriptorOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetBucketColumns(val *[]*string) {
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetColumns(val *[]*GluePartitionStorageDescriptorColumns) {
	_jsii_.Set(
		j,
		"columns",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetCompressed(val interface{}) {
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetInputFormat(val *string) {
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetNumberOfBuckets(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetOutputFormat(val *string) {
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetSortColumns(val *[]*GluePartitionStorageDescriptorSortColumns) {
	_jsii_.Set(
		j,
		"sortColumns",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetStoredAsSubDirectories(val interface{}) {
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) PutSerDeInfo(value *GluePartitionStorageDescriptorSerDeInfo) {
	_jsii_.InvokeVoid(
		g,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) PutSkewedInfo(value *GluePartitionStorageDescriptorSkewedInfo) {
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

type GluePartitionStorageDescriptorSerDeInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#name GluePartition#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#parameters GluePartition#parameters}.
	Parameters interface{} `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#serialization_library GluePartition#serialization_library}.
	SerializationLibrary *string `json:"serializationLibrary"`
}

type GluePartitionStorageDescriptorSerDeInfoOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	SerializationLibrary() *string
	SetSerializationLibrary(val *string)
	SerializationLibraryInput() *string
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
	ResetName()
	ResetParameters()
	ResetSerializationLibrary()
}

// The jsii proxy struct for GluePartitionStorageDescriptorSerDeInfoOutputReference
type jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SerializationLibrary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibrary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SerializationLibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGluePartitionStorageDescriptorSerDeInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GluePartitionStorageDescriptorSerDeInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorSerDeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGluePartitionStorageDescriptorSerDeInfoOutputReference_Override(g GluePartitionStorageDescriptorSerDeInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorSerDeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetSerializationLibrary(val *string) {
	_jsii_.Set(
		j,
		"serializationLibrary",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSerDeInfoOutputReference) ResetSerializationLibrary() {
	_jsii_.InvokeVoid(
		g,
		"resetSerializationLibrary",
		nil, // no parameters
	)
}

type GluePartitionStorageDescriptorSkewedInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#skewed_column_names GluePartition#skewed_column_names}.
	SkewedColumnNames *[]*string `json:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#skewed_column_value_location_maps GluePartition#skewed_column_value_location_maps}.
	SkewedColumnValueLocationMaps interface{} `json:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#skewed_column_values GluePartition#skewed_column_values}.
	SkewedColumnValues *[]*string `json:"skewedColumnValues"`
}

type GluePartitionStorageDescriptorSkewedInfoOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SkewedColumnNames() *[]*string
	SetSkewedColumnNames(val *[]*string)
	SkewedColumnNamesInput() *[]*string
	SkewedColumnValueLocationMaps() interface{}
	SetSkewedColumnValueLocationMaps(val interface{})
	SkewedColumnValueLocationMapsInput() interface{}
	SkewedColumnValues() *[]*string
	SetSkewedColumnValues(val *[]*string)
	SkewedColumnValuesInput() *[]*string
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
	ResetSkewedColumnNames()
	ResetSkewedColumnValueLocationMaps()
	ResetSkewedColumnValues()
}

// The jsii proxy struct for GluePartitionStorageDescriptorSkewedInfoOutputReference
type jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMaps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMapsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SkewedColumnValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGluePartitionStorageDescriptorSkewedInfoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GluePartitionStorageDescriptorSkewedInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGluePartitionStorageDescriptorSkewedInfoOutputReference_Override(g GluePartitionStorageDescriptorSkewedInfoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GluePartitionStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"skewedColumnNames",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnValueLocationMaps(val interface{}) {
	_jsii_.Set(
		j,
		"skewedColumnValueLocationMaps",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetSkewedColumnValues(val *[]*string) {
	_jsii_.Set(
		j,
		"skewedColumnValues",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValueLocationMaps() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValueLocationMaps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValues() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValues",
		nil, // no parameters
	)
}

type GluePartitionStorageDescriptorSortColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#column GluePartition#column}.
	Column *string `json:"column"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition.html#sort_order GluePartition#sort_order}.
	SortOrder *float64 `json:"sortOrder"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html aws_glue_registry}.
type GlueRegistry interface {
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegistryName() *string
	SetRegistryName(val *string)
	RegistryNameInput() *string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueRegistry
type jsiiProxy_GlueRegistry struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueRegistry) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) RegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) RegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueRegistry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html aws_glue_registry} Resource.
func NewGlueRegistry(scope constructs.Construct, id *string, config *GlueRegistryConfig) GlueRegistry {
	_init_.Initialize()

	j := jsiiProxy_GlueRegistry{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueRegistry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html aws_glue_registry} Resource.
func NewGlueRegistry_Override(g GlueRegistry, scope constructs.Construct, id *string, config *GlueRegistryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueRegistry",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueRegistry) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetRegistryName(val *string) {
	_jsii_.Set(
		j,
		"registryName",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueRegistry) SetTagsAll(val interface{}) {
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
func GlueRegistry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueRegistry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueRegistry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueRegistry",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueRegistry) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueRegistry) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueRegistry) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueRegistry) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueRegistry) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueRegistry) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueRegistry) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueRegistry) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueRegistry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueRegistry) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueRegistry) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueRegistry) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueRegistry) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueRegistry) ToString() *string {
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
func (g *jsiiProxy_GlueRegistry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueRegistryConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html#registry_name GlueRegistry#registry_name}.
	RegistryName *string `json:"registryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html#description GlueRegistry#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html#tags GlueRegistry#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_registry.html#tags_all GlueRegistry#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_resource_policy.html aws_glue_resource_policy}.
type GlueResourcePolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnableHybrid() *string
	SetEnableHybrid(val *string)
	EnableHybridInput() *string
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
	ResetEnableHybrid()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueResourcePolicy
type jsiiProxy_GlueResourcePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueResourcePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) EnableHybrid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableHybrid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) EnableHybridInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableHybridInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueResourcePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_resource_policy.html aws_glue_resource_policy} Resource.
func NewGlueResourcePolicy(scope constructs.Construct, id *string, config *GlueResourcePolicyConfig) GlueResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_GlueResourcePolicy{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueResourcePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_resource_policy.html aws_glue_resource_policy} Resource.
func NewGlueResourcePolicy_Override(g GlueResourcePolicy, scope constructs.Construct, id *string, config *GlueResourcePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueResourcePolicy",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetEnableHybrid(val *string) {
	_jsii_.Set(
		j,
		"enableHybrid",
		val,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_GlueResourcePolicy) SetProvider(val cdktf.TerraformProvider) {
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
func GlueResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueResourcePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueResourcePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueResourcePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueResourcePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueResourcePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueResourcePolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueResourcePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueResourcePolicy) ResetEnableHybrid() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableHybrid",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueResourcePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueResourcePolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueResourcePolicy) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueResourcePolicy) ToString() *string {
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
func (g *jsiiProxy_GlueResourcePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueResourcePolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_resource_policy.html#policy GlueResourcePolicy#policy}.
	Policy *string `json:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_resource_policy.html#enable_hybrid GlueResourcePolicy#enable_hybrid}.
	EnableHybrid *string `json:"enableHybrid"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html aws_glue_schema}.
type GlueSchema interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Compatibility() *string
	SetCompatibility(val *string)
	CompatibilityInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LatestSchemaVersion() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NextSchemaVersion() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegistryArn() *string
	SetRegistryArn(val *string)
	RegistryArnInput() *string
	RegistryName() *string
	SchemaCheckpoint() *float64
	SchemaDefinition() *string
	SetSchemaDefinition(val *string)
	SchemaDefinitionInput() *string
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetRegistryArn()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueSchema
type jsiiProxy_GlueSchema struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueSchema) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Compatibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) CompatibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) LatestSchemaVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) NextSchemaVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nextSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) RegistryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) RegistryArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) RegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) SchemaCheckpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaCheckpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) SchemaDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) SchemaDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSchema) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html aws_glue_schema} Resource.
func NewGlueSchema(scope constructs.Construct, id *string, config *GlueSchemaConfig) GlueSchema {
	_init_.Initialize()

	j := jsiiProxy_GlueSchema{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSchema",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html aws_glue_schema} Resource.
func NewGlueSchema_Override(g GlueSchema, scope constructs.Construct, id *string, config *GlueSchemaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSchema",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueSchema) SetCompatibility(val *string) {
	_jsii_.Set(
		j,
		"compatibility",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetDataFormat(val *string) {
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetRegistryArn(val *string) {
	_jsii_.Set(
		j,
		"registryArn",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetSchemaDefinition(val *string) {
	_jsii_.Set(
		j,
		"schemaDefinition",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetSchemaName(val *string) {
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueSchema) SetTagsAll(val interface{}) {
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
func GlueSchema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueSchema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueSchema_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueSchema",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueSchema) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueSchema) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSchema) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSchema) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSchema) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSchema) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSchema) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueSchema) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueSchema) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSchema) ResetRegistryArn() {
	_jsii_.InvokeVoid(
		g,
		"resetRegistryArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSchema) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSchema) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSchema) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueSchema) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueSchema) ToString() *string {
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
func (g *jsiiProxy_GlueSchema) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueSchemaConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#compatibility GlueSchema#compatibility}.
	Compatibility *string `json:"compatibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#data_format GlueSchema#data_format}.
	DataFormat *string `json:"dataFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#schema_definition GlueSchema#schema_definition}.
	SchemaDefinition *string `json:"schemaDefinition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#schema_name GlueSchema#schema_name}.
	SchemaName *string `json:"schemaName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#description GlueSchema#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#registry_arn GlueSchema#registry_arn}.
	RegistryArn *string `json:"registryArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#tags GlueSchema#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_schema.html#tags_all GlueSchema#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html aws_glue_security_configuration}.
type GlueSecurityConfiguration interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EncryptionConfiguration() GlueSecurityConfigurationEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() *GlueSecurityConfigurationEncryptionConfiguration
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
	PutEncryptionConfiguration(value *GlueSecurityConfigurationEncryptionConfiguration)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueSecurityConfiguration
type jsiiProxy_GlueSecurityConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueSecurityConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) EncryptionConfiguration() GlueSecurityConfigurationEncryptionConfigurationOutputReference {
	var returns GlueSecurityConfigurationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) EncryptionConfigurationInput() *GlueSecurityConfigurationEncryptionConfiguration {
	var returns *GlueSecurityConfigurationEncryptionConfiguration
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html aws_glue_security_configuration} Resource.
func NewGlueSecurityConfiguration(scope constructs.Construct, id *string, config *GlueSecurityConfigurationConfig) GlueSecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_GlueSecurityConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html aws_glue_security_configuration} Resource.
func NewGlueSecurityConfiguration_Override(g GlueSecurityConfiguration, scope constructs.Construct, id *string, config *GlueSecurityConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfiguration",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfiguration) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfiguration) SetProvider(val cdktf.TerraformProvider) {
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
func GlueSecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueSecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueSecurityConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueSecurityConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSecurityConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSecurityConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSecurityConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueSecurityConfiguration) PutEncryptionConfiguration(value *GlueSecurityConfigurationEncryptionConfiguration) {
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueSecurityConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfiguration) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueSecurityConfiguration) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueSecurityConfiguration) ToString() *string {
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
func (g *jsiiProxy_GlueSecurityConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueSecurityConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#encryption_configuration GlueSecurityConfiguration#encryption_configuration}
	EncryptionConfiguration *GlueSecurityConfigurationEncryptionConfiguration `json:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#name GlueSecurityConfiguration#name}.
	Name *string `json:"name"`
}

type GlueSecurityConfigurationEncryptionConfiguration struct {
	// cloudwatch_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#cloudwatch_encryption GlueSecurityConfiguration#cloudwatch_encryption}
	CloudwatchEncryption *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption `json:"cloudwatchEncryption"`
	// job_bookmarks_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#job_bookmarks_encryption GlueSecurityConfiguration#job_bookmarks_encryption}
	JobBookmarksEncryption *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption `json:"jobBookmarksEncryption"`
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#s3_encryption GlueSecurityConfiguration#s3_encryption}
	S3Encryption *GlueSecurityConfigurationEncryptionConfigurationS3Encryption `json:"s3Encryption"`
}

type GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#cloudwatch_encryption_mode GlueSecurityConfiguration#cloudwatch_encryption_mode}.
	CloudwatchEncryptionMode *string `json:"cloudwatchEncryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
}

type GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference interface {
	cdktf.ComplexObject
	CloudwatchEncryptionMode() *string
	SetCloudwatchEncryptionMode(val *string)
	CloudwatchEncryptionModeInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	ResetCloudwatchEncryptionMode()
	ResetKmsKeyArn()
}

// The jsii proxy struct for GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference
type jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) CloudwatchEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) CloudwatchEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference_Override(g GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) SetCloudwatchEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"cloudwatchEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) ResetCloudwatchEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudwatchEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

type GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#job_bookmarks_encryption_mode GlueSecurityConfiguration#job_bookmarks_encryption_mode}.
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
}

type GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	JobBookmarksEncryptionMode() *string
	SetJobBookmarksEncryptionMode(val *string)
	JobBookmarksEncryptionModeInput() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	ResetJobBookmarksEncryptionMode()
	ResetKmsKeyArn()
}

// The jsii proxy struct for GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference
type jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) JobBookmarksEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) JobBookmarksEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference_Override(g GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) SetJobBookmarksEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"jobBookmarksEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ResetJobBookmarksEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetJobBookmarksEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

type GlueSecurityConfigurationEncryptionConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchEncryption() GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference
	CloudwatchEncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	JobBookmarksEncryption() GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference
	JobBookmarksEncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption
	S3Encryption() GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference
	S3EncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationS3Encryption
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
	PutCloudwatchEncryption(value *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption)
	PutJobBookmarksEncryption(value *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption)
	PutS3Encryption(value *GlueSecurityConfigurationEncryptionConfigurationS3Encryption)
}

// The jsii proxy struct for GlueSecurityConfigurationEncryptionConfigurationOutputReference
type jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) CloudwatchEncryption() GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference {
	var returns GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryptionOutputReference
	_jsii_.Get(
		j,
		"cloudwatchEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) CloudwatchEncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption {
	var returns *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption
	_jsii_.Get(
		j,
		"cloudwatchEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) JobBookmarksEncryption() GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference {
	var returns GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference
	_jsii_.Get(
		j,
		"jobBookmarksEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) JobBookmarksEncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption {
	var returns *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) S3Encryption() GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference {
	var returns GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference
	_jsii_.Get(
		j,
		"s3Encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) S3EncryptionInput() *GlueSecurityConfigurationEncryptionConfigurationS3Encryption {
	var returns *GlueSecurityConfigurationEncryptionConfigurationS3Encryption
	_jsii_.Get(
		j,
		"s3EncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueSecurityConfigurationEncryptionConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueSecurityConfigurationEncryptionConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueSecurityConfigurationEncryptionConfigurationOutputReference_Override(g GlueSecurityConfigurationEncryptionConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) PutCloudwatchEncryption(value *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption) {
	_jsii_.InvokeVoid(
		g,
		"putCloudwatchEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) PutJobBookmarksEncryption(value *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption) {
	_jsii_.InvokeVoid(
		g,
		"putJobBookmarksEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationOutputReference) PutS3Encryption(value *GlueSecurityConfigurationEncryptionConfigurationS3Encryption) {
	_jsii_.InvokeVoid(
		g,
		"putS3Encryption",
		[]interface{}{value},
	)
}

type GlueSecurityConfigurationEncryptionConfigurationS3Encryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_security_configuration.html#s3_encryption_mode GlueSecurityConfiguration#s3_encryption_mode}.
	S3EncryptionMode *string `json:"s3EncryptionMode"`
}

type GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	S3EncryptionMode() *string
	SetS3EncryptionMode(val *string)
	S3EncryptionModeInput() *string
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
	ResetKmsKeyArn()
	ResetS3EncryptionMode()
}

// The jsii proxy struct for GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference
type jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) S3EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3EncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) S3EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3EncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference_Override(g GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) SetS3EncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"s3EncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationS3EncryptionOutputReference) ResetS3EncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetS3EncryptionMode",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html aws_glue_trigger}.
type GlueTrigger interface {
	cdktf.TerraformResource
	Actions() *[]*GlueTriggerActions
	SetActions(val *[]*GlueTriggerActions)
	ActionsInput() *[]*GlueTriggerActions
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Predicate() GlueTriggerPredicateOutputReference
	PredicateInput() *GlueTriggerPredicate
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
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
	Timeouts() GlueTriggerTimeoutsOutputReference
	TimeoutsInput() *GlueTriggerTimeouts
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WorkflowName() *string
	SetWorkflowName(val *string)
	WorkflowNameInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutPredicate(value *GlueTriggerPredicate)
	PutTimeouts(value *GlueTriggerTimeouts)
	ResetDescription()
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPredicate()
	ResetSchedule()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetWorkflowName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueTrigger
type jsiiProxy_GlueTrigger struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueTrigger) Actions() *[]*GlueTriggerActions {
	var returns *[]*GlueTriggerActions
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) ActionsInput() *[]*GlueTriggerActions {
	var returns *[]*GlueTriggerActions
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Predicate() GlueTriggerPredicateOutputReference {
	var returns GlueTriggerPredicateOutputReference
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) PredicateInput() *GlueTriggerPredicate {
	var returns *GlueTriggerPredicate
	_jsii_.Get(
		j,
		"predicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Timeouts() GlueTriggerTimeoutsOutputReference {
	var returns GlueTriggerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TimeoutsInput() *GlueTriggerTimeouts {
	var returns *GlueTriggerTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTrigger) WorkflowNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html aws_glue_trigger} Resource.
func NewGlueTrigger(scope constructs.Construct, id *string, config *GlueTriggerConfig) GlueTrigger {
	_init_.Initialize()

	j := jsiiProxy_GlueTrigger{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html aws_glue_trigger} Resource.
func NewGlueTrigger_Override(g GlueTrigger, scope constructs.Construct, id *string, config *GlueTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTrigger",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueTrigger) SetActions(val *[]*GlueTriggerActions) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GlueTrigger) SetWorkflowName(val *string) {
	_jsii_.Set(
		j,
		"workflowName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueTrigger) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueTrigger) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueTrigger) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueTrigger) PutPredicate(value *GlueTriggerPredicate) {
	_jsii_.InvokeVoid(
		g,
		"putPredicate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTrigger) PutTimeouts(value *GlueTriggerTimeouts) {
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTrigger) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetPredicate() {
	_jsii_.InvokeVoid(
		g,
		"resetPredicate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) ResetWorkflowName() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflowName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTrigger) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueTrigger) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueTrigger) ToString() *string {
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
func (g *jsiiProxy_GlueTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueTriggerActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#arguments GlueTrigger#arguments}.
	Arguments interface{} `json:"arguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#crawler_name GlueTrigger#crawler_name}.
	CrawlerName *string `json:"crawlerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#job_name GlueTrigger#job_name}.
	JobName *string `json:"jobName"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#notification_property GlueTrigger#notification_property}
	NotificationProperty *GlueTriggerActionsNotificationProperty `json:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#security_configuration GlueTrigger#security_configuration}.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#timeout GlueTrigger#timeout}.
	Timeout *float64 `json:"timeout"`
}

type GlueTriggerActionsNotificationProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#notify_delay_after GlueTrigger#notify_delay_after}.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter"`
}

type GlueTriggerActionsNotificationPropertyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NotifyDelayAfter() *float64
	SetNotifyDelayAfter(val *float64)
	NotifyDelayAfterInput() *float64
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
	ResetNotifyDelayAfter()
}

// The jsii proxy struct for GlueTriggerActionsNotificationPropertyOutputReference
type jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) NotifyDelayAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notifyDelayAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) NotifyDelayAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notifyDelayAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueTriggerActionsNotificationPropertyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueTriggerActionsNotificationPropertyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerActionsNotificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueTriggerActionsNotificationPropertyOutputReference_Override(g GlueTriggerActionsNotificationPropertyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerActionsNotificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) SetNotifyDelayAfter(val *float64) {
	_jsii_.Set(
		j,
		"notifyDelayAfter",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTriggerActionsNotificationPropertyOutputReference) ResetNotifyDelayAfter() {
	_jsii_.InvokeVoid(
		g,
		"resetNotifyDelayAfter",
		nil, // no parameters
	)
}

type GlueTriggerConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#actions GlueTrigger#actions}
	Actions *[]*GlueTriggerActions `json:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#name GlueTrigger#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#type GlueTrigger#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#description GlueTrigger#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#enabled GlueTrigger#enabled}.
	Enabled interface{} `json:"enabled"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#predicate GlueTrigger#predicate}
	Predicate *GlueTriggerPredicate `json:"predicate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#schedule GlueTrigger#schedule}.
	Schedule *string `json:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#tags GlueTrigger#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#tags_all GlueTrigger#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#timeouts GlueTrigger#timeouts}
	Timeouts *GlueTriggerTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#workflow_name GlueTrigger#workflow_name}.
	WorkflowName *string `json:"workflowName"`
}

type GlueTriggerPredicate struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#conditions GlueTrigger#conditions}
	Conditions *[]*GlueTriggerPredicateConditions `json:"conditions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#logical GlueTrigger#logical}.
	Logical *string `json:"logical"`
}

type GlueTriggerPredicateConditions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#crawler_name GlueTrigger#crawler_name}.
	CrawlerName *string `json:"crawlerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#crawl_state GlueTrigger#crawl_state}.
	CrawlState *string `json:"crawlState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#job_name GlueTrigger#job_name}.
	JobName *string `json:"jobName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#logical_operator GlueTrigger#logical_operator}.
	LogicalOperator *string `json:"logicalOperator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#state GlueTrigger#state}.
	State *string `json:"state"`
}

type GlueTriggerPredicateOutputReference interface {
	cdktf.ComplexObject
	Conditions() *[]*GlueTriggerPredicateConditions
	SetConditions(val *[]*GlueTriggerPredicateConditions)
	ConditionsInput() *[]*GlueTriggerPredicateConditions
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Logical() *string
	SetLogical(val *string)
	LogicalInput() *string
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
	ResetLogical()
}

// The jsii proxy struct for GlueTriggerPredicateOutputReference
type jsiiProxy_GlueTriggerPredicateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) Conditions() *[]*GlueTriggerPredicateConditions {
	var returns *[]*GlueTriggerPredicateConditions
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) ConditionsInput() *[]*GlueTriggerPredicateConditions {
	var returns *[]*GlueTriggerPredicateConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) Logical() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) LogicalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueTriggerPredicateOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueTriggerPredicateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueTriggerPredicateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerPredicateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueTriggerPredicateOutputReference_Override(g GlueTriggerPredicateOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerPredicateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) SetConditions(val *[]*GlueTriggerPredicateConditions) {
	_jsii_.Set(
		j,
		"conditions",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) SetLogical(val *string) {
	_jsii_.Set(
		j,
		"logical",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerPredicateOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueTriggerPredicateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTriggerPredicateOutputReference) ResetLogical() {
	_jsii_.InvokeVoid(
		g,
		"resetLogical",
		nil, // no parameters
	)
}

type GlueTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#create GlueTrigger#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger.html#delete GlueTrigger#delete}.
	Delete *string `json:"delete"`
}

type GlueTriggerTimeoutsOutputReference interface {
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
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
}

// The jsii proxy struct for GlueTriggerTimeoutsOutputReference
type jsiiProxy_GlueTriggerTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGlueTriggerTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) GlueTriggerTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_GlueTriggerTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewGlueTriggerTimeoutsOutputReference_Override(g GlueTriggerTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueTriggerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		g,
	)
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTriggerTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTriggerTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetDelete",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html aws_glue_user_defined_function}.
type GlueUserDefinedFunction interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ClassName() *string
	SetClassName(val *string)
	ClassNameInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreateTime() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
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
	OwnerName() *string
	SetOwnerName(val *string)
	OwnerNameInput() *string
	OwnerType() *string
	SetOwnerType(val *string)
	OwnerTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceUris() *[]*GlueUserDefinedFunctionResourceUris
	SetResourceUris(val *[]*GlueUserDefinedFunctionResourceUris)
	ResourceUrisInput() *[]*GlueUserDefinedFunctionResourceUris
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
	ResetCatalogId()
	ResetOverrideLogicalId()
	ResetResourceUris()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueUserDefinedFunction
type jsiiProxy_GlueUserDefinedFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueUserDefinedFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) ClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) OwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) OwnerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) OwnerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) OwnerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) ResourceUris() *[]*GlueUserDefinedFunctionResourceUris {
	var returns *[]*GlueUserDefinedFunctionResourceUris
	_jsii_.Get(
		j,
		"resourceUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) ResourceUrisInput() *[]*GlueUserDefinedFunctionResourceUris {
	var returns *[]*GlueUserDefinedFunctionResourceUris
	_jsii_.Get(
		j,
		"resourceUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueUserDefinedFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html aws_glue_user_defined_function} Resource.
func NewGlueUserDefinedFunction(scope constructs.Construct, id *string, config *GlueUserDefinedFunctionConfig) GlueUserDefinedFunction {
	_init_.Initialize()

	j := jsiiProxy_GlueUserDefinedFunction{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueUserDefinedFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html aws_glue_user_defined_function} Resource.
func NewGlueUserDefinedFunction_Override(g GlueUserDefinedFunction, scope constructs.Construct, id *string, config *GlueUserDefinedFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueUserDefinedFunction",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetClassName(val *string) {
	_jsii_.Set(
		j,
		"className",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetOwnerName(val *string) {
	_jsii_.Set(
		j,
		"ownerName",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetOwnerType(val *string) {
	_jsii_.Set(
		j,
		"ownerType",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueUserDefinedFunction) SetResourceUris(val *[]*GlueUserDefinedFunctionResourceUris) {
	_jsii_.Set(
		j,
		"resourceUris",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GlueUserDefinedFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueUserDefinedFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueUserDefinedFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueUserDefinedFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueUserDefinedFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueUserDefinedFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueUserDefinedFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueUserDefinedFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueUserDefinedFunction) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueUserDefinedFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueUserDefinedFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueUserDefinedFunction) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueUserDefinedFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueUserDefinedFunction) ResetResourceUris() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueUserDefinedFunction) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueUserDefinedFunction) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueUserDefinedFunction) ToString() *string {
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
func (g *jsiiProxy_GlueUserDefinedFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueUserDefinedFunctionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#class_name GlueUserDefinedFunction#class_name}.
	ClassName *string `json:"className"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#database_name GlueUserDefinedFunction#database_name}.
	DatabaseName *string `json:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#name GlueUserDefinedFunction#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#owner_name GlueUserDefinedFunction#owner_name}.
	OwnerName *string `json:"ownerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#owner_type GlueUserDefinedFunction#owner_type}.
	OwnerType *string `json:"ownerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#catalog_id GlueUserDefinedFunction#catalog_id}.
	CatalogId *string `json:"catalogId"`
	// resource_uris block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#resource_uris GlueUserDefinedFunction#resource_uris}
	ResourceUris *[]*GlueUserDefinedFunctionResourceUris `json:"resourceUris"`
}

type GlueUserDefinedFunctionResourceUris struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#resource_type GlueUserDefinedFunction#resource_type}.
	ResourceType *string `json:"resourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_user_defined_function.html#uri GlueUserDefinedFunction#uri}.
	Uri *string `json:"uri"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html aws_glue_workflow}.
type GlueWorkflow interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultRunProperties() interface{}
	SetDefaultRunProperties(val interface{})
	DefaultRunPropertiesInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxConcurrentRuns() *float64
	SetMaxConcurrentRuns(val *float64)
	MaxConcurrentRunsInput() *float64
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
	ResetDefaultRunProperties()
	ResetDescription()
	ResetMaxConcurrentRuns()
	ResetName()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueWorkflow
type jsiiProxy_GlueWorkflow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueWorkflow) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) DefaultRunProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRunProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) DefaultRunPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRunPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) MaxConcurrentRuns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) MaxConcurrentRunsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueWorkflow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html aws_glue_workflow} Resource.
func NewGlueWorkflow(scope constructs.Construct, id *string, config *GlueWorkflowConfig) GlueWorkflow {
	_init_.Initialize()

	j := jsiiProxy_GlueWorkflow{}

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueWorkflow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html aws_glue_workflow} Resource.
func NewGlueWorkflow_Override(g GlueWorkflow, scope constructs.Construct, id *string, config *GlueWorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Glue.GlueWorkflow",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetDefaultRunProperties(val interface{}) {
	_jsii_.Set(
		j,
		"defaultRunProperties",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetMaxConcurrentRuns(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentRuns",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GlueWorkflow) SetTagsAll(val interface{}) {
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
func GlueWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Glue.GlueWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueWorkflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Glue.GlueWorkflow",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (g *jsiiProxy_GlueWorkflow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (g *jsiiProxy_GlueWorkflow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueWorkflow) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (g *jsiiProxy_GlueWorkflow) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (g *jsiiProxy_GlueWorkflow) GetStringAttribute(terraformAttribute *string) *string {
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
func (g *jsiiProxy_GlueWorkflow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (g *jsiiProxy_GlueWorkflow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetDefaultRunProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultRunProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetMaxConcurrentRuns() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentRuns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GlueWorkflow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueWorkflow) SynthesizeAttributes() *map[string]interface{} {
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
func (g *jsiiProxy_GlueWorkflow) ToMetadata() interface{} {
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
func (g *jsiiProxy_GlueWorkflow) ToString() *string {
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
func (g *jsiiProxy_GlueWorkflow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GlueWorkflowConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#default_run_properties GlueWorkflow#default_run_properties}.
	DefaultRunProperties interface{} `json:"defaultRunProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#description GlueWorkflow#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#max_concurrent_runs GlueWorkflow#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#name GlueWorkflow#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#tags GlueWorkflow#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_workflow.html#tags_all GlueWorkflow#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}
