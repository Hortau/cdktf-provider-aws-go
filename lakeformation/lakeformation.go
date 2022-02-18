package lakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/lakeformation/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings}.
type DataAwsLakeformationDataLakeSettings interface {
	cdktf.TerraformDataSource
	Admins() *[]*string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	TrustedResourceOwners() *[]*string
	AddOverride(path *string, value interface{})
	CreateDatabaseDefaultPermissions(index *string) DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions
	CreateTableDefaultPermissions(index *string) DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions
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
	ResetCatalogId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLakeformationDataLakeSettings
type jsiiProxy_DataAwsLakeformationDataLakeSettings struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Admins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Data Source.
func NewDataAwsLakeformationDataLakeSettings(scope constructs.Construct, id *string, config *DataAwsLakeformationDataLakeSettingsConfig) DataAwsLakeformationDataLakeSettings {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationDataLakeSettings{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Data Source.
func NewDataAwsLakeformationDataLakeSettings_Override(d DataAwsLakeformationDataLakeSettings, scope constructs.Construct, id *string, config *DataAwsLakeformationDataLakeSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettings) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsLakeformationDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLakeformationDataLakeSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) CreateDatabaseDefaultPermissions(index *string) DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions {
	var returns DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions

	_jsii_.Invoke(
		d,
		"createDatabaseDefaultPermissions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) CreateTableDefaultPermissions(index *string) DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions {
	var returns DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions

	_jsii_.Invoke(
		d,
		"createTableDefaultPermissions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) ToString() *string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type DataAwsLakeformationDataLakeSettingsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_data_lake_settings#catalog_id DataAwsLakeformationDataLakeSettings#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
}

type DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Permissions() *[]*string
	Principal() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions
type jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions_Override(d DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateDatabaseDefaultPermissions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Permissions() *[]*string
	Principal() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions
type jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions_Override(d DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationDataLakeSettingsCreateTableDefaultPermissions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions aws_lakeformation_permissions}.
type DataAwsLakeformationPermissions interface {
	cdktf.TerraformDataSource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CatalogResource() interface{}
	SetCatalogResource(val interface{})
	CatalogResourceInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Database() DataAwsLakeformationPermissionsDatabaseOutputReference
	DatabaseInput() *DataAwsLakeformationPermissionsDatabase
	DataLocation() DataAwsLakeformationPermissionsDataLocationOutputReference
	DataLocationInput() *DataAwsLakeformationPermissionsDataLocation
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Permissions() *[]*string
	PermissionsWithGrantOption() *[]*string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Table() DataAwsLakeformationPermissionsTableOutputReference
	TableInput() *DataAwsLakeformationPermissionsTable
	TableWithColumns() DataAwsLakeformationPermissionsTableWithColumnsOutputReference
	TableWithColumnsInput() *DataAwsLakeformationPermissionsTableWithColumns
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
	PutDatabase(value *DataAwsLakeformationPermissionsDatabase)
	PutDataLocation(value *DataAwsLakeformationPermissionsDataLocation)
	PutTable(value *DataAwsLakeformationPermissionsTable)
	PutTableWithColumns(value *DataAwsLakeformationPermissionsTableWithColumns)
	ResetCatalogId()
	ResetCatalogResource()
	ResetDatabase()
	ResetDataLocation()
	ResetOverrideLogicalId()
	ResetTable()
	ResetTableWithColumns()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsLakeformationPermissions
type jsiiProxy_DataAwsLakeformationPermissions struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) CatalogResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) CatalogResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Database() DataAwsLakeformationPermissionsDatabaseOutputReference {
	var returns DataAwsLakeformationPermissionsDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) DatabaseInput() *DataAwsLakeformationPermissionsDatabase {
	var returns *DataAwsLakeformationPermissionsDatabase
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) DataLocation() DataAwsLakeformationPermissionsDataLocationOutputReference {
	var returns DataAwsLakeformationPermissionsDataLocationOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) DataLocationInput() *DataAwsLakeformationPermissionsDataLocation {
	var returns *DataAwsLakeformationPermissionsDataLocation
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) Table() DataAwsLakeformationPermissionsTableOutputReference {
	var returns DataAwsLakeformationPermissionsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TableInput() *DataAwsLakeformationPermissionsTable {
	var returns *DataAwsLakeformationPermissionsTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TableWithColumns() DataAwsLakeformationPermissionsTableWithColumnsOutputReference {
	var returns DataAwsLakeformationPermissionsTableWithColumnsOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TableWithColumnsInput() *DataAwsLakeformationPermissionsTableWithColumns {
	var returns *DataAwsLakeformationPermissionsTableWithColumns
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions aws_lakeformation_permissions} Data Source.
func NewDataAwsLakeformationPermissions(scope constructs.Construct, id *string, config *DataAwsLakeformationPermissionsConfig) DataAwsLakeformationPermissions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationPermissions{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions aws_lakeformation_permissions} Data Source.
func NewDataAwsLakeformationPermissions_Override(d DataAwsLakeformationPermissions, scope constructs.Construct, id *string, config *DataAwsLakeformationPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissions",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetCatalogResource(val interface{}) {
	_jsii_.Set(
		j,
		"catalogResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissions) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsLakeformationPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLakeformationPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) PutDatabase(value *DataAwsLakeformationPermissionsDatabase) {
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) PutDataLocation(value *DataAwsLakeformationPermissionsDataLocation) {
	_jsii_.InvokeVoid(
		d,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) PutTable(value *DataAwsLakeformationPermissionsTable) {
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) PutTableWithColumns(value *DataAwsLakeformationPermissionsTableWithColumns) {
	_jsii_.InvokeVoid(
		d,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetCatalogResource() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogResource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetDataLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetDataLocation",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissions) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) ToString() *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type DataAwsLakeformationPermissionsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#principal DataAwsLakeformationPermissions#principal}.
	Principal *string `json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_resource DataAwsLakeformationPermissions#catalog_resource}.
	CatalogResource interface{} `json:"catalogResource" yaml:"catalogResource"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#database DataAwsLakeformationPermissions#database}
	Database *DataAwsLakeformationPermissionsDatabase `json:"database" yaml:"database"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#data_location DataAwsLakeformationPermissions#data_location}
	DataLocation *DataAwsLakeformationPermissionsDataLocation `json:"dataLocation" yaml:"dataLocation"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#table DataAwsLakeformationPermissions#table}
	Table *DataAwsLakeformationPermissionsTable `json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#table_with_columns DataAwsLakeformationPermissions#table_with_columns}
	TableWithColumns *DataAwsLakeformationPermissionsTableWithColumns `json:"tableWithColumns" yaml:"tableWithColumns"`
}

type DataAwsLakeformationPermissionsDataLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#arn DataAwsLakeformationPermissions#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
}

type DataAwsLakeformationPermissionsDataLocationOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	InternalValue() *DataAwsLakeformationPermissionsDataLocation
	SetInternalValue(val *DataAwsLakeformationPermissionsDataLocation)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
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
	ResetCatalogId()
}

// The jsii proxy struct for DataAwsLakeformationPermissionsDataLocationOutputReference
type jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) InternalValue() *DataAwsLakeformationPermissionsDataLocation {
	var returns *DataAwsLakeformationPermissionsDataLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataAwsLakeformationPermissionsDataLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DataAwsLakeformationPermissionsDataLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsDataLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDataAwsLakeformationPermissionsDataLocationOutputReference_Override(d DataAwsLakeformationPermissionsDataLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsDataLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetInternalValue(val *DataAwsLakeformationPermissionsDataLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsDataLocationOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

type DataAwsLakeformationPermissionsDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#name DataAwsLakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
}

type DataAwsLakeformationPermissionsDatabaseOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	InternalValue() *DataAwsLakeformationPermissionsDatabase
	SetInternalValue(val *DataAwsLakeformationPermissionsDatabase)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetCatalogId()
}

// The jsii proxy struct for DataAwsLakeformationPermissionsDatabaseOutputReference
type jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) InternalValue() *DataAwsLakeformationPermissionsDatabase {
	var returns *DataAwsLakeformationPermissionsDatabase
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataAwsLakeformationPermissionsDatabaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DataAwsLakeformationPermissionsDatabaseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDataAwsLakeformationPermissionsDatabaseOutputReference_Override(d DataAwsLakeformationPermissionsDatabaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetInternalValue(val *DataAwsLakeformationPermissionsDatabase) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsDatabaseOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

type DataAwsLakeformationPermissionsTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#database_name DataAwsLakeformationPermissions#database_name}.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#name DataAwsLakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#wildcard DataAwsLakeformationPermissions#wildcard}.
	Wildcard interface{} `json:"wildcard" yaml:"wildcard"`
}

type DataAwsLakeformationPermissionsTableOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	InternalValue() *DataAwsLakeformationPermissionsTable
	SetInternalValue(val *DataAwsLakeformationPermissionsTable)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	ResetCatalogId()
	ResetName()
	ResetWildcard()
}

// The jsii proxy struct for DataAwsLakeformationPermissionsTableOutputReference
type jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) InternalValue() *DataAwsLakeformationPermissionsTable {
	var returns *DataAwsLakeformationPermissionsTable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}

func NewDataAwsLakeformationPermissionsTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DataAwsLakeformationPermissionsTableOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDataAwsLakeformationPermissionsTableOutputReference_Override(d DataAwsLakeformationPermissionsTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetInternalValue(val *DataAwsLakeformationPermissionsTable) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) SetWildcard(val interface{}) {
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		d,
		"resetWildcard",
		nil, // no parameters
	)
}

type DataAwsLakeformationPermissionsTableWithColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#database_name DataAwsLakeformationPermissions#database_name}.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#name DataAwsLakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#column_names DataAwsLakeformationPermissions#column_names}.
	ColumnNames *[]*string `json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#excluded_column_names DataAwsLakeformationPermissions#excluded_column_names}.
	ExcludedColumnNames *[]*string `json:"excludedColumnNames" yaml:"excludedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#wildcard DataAwsLakeformationPermissions#wildcard}.
	Wildcard interface{} `json:"wildcard" yaml:"wildcard"`
}

type DataAwsLakeformationPermissionsTableWithColumnsOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	ColumnNames() *[]*string
	SetColumnNames(val *[]*string)
	ColumnNamesInput() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	ExcludedColumnNames() *[]*string
	SetExcludedColumnNames(val *[]*string)
	ExcludedColumnNamesInput() *[]*string
	InternalValue() *DataAwsLakeformationPermissionsTableWithColumns
	SetInternalValue(val *DataAwsLakeformationPermissionsTableWithColumns)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	ResetCatalogId()
	ResetColumnNames()
	ResetExcludedColumnNames()
	ResetWildcard()
}

// The jsii proxy struct for DataAwsLakeformationPermissionsTableWithColumnsOutputReference
type jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ExcludedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ExcludedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) InternalValue() *DataAwsLakeformationPermissionsTableWithColumns {
	var returns *DataAwsLakeformationPermissionsTableWithColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}

func NewDataAwsLakeformationPermissionsTableWithColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) DataAwsLakeformationPermissionsTableWithColumnsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsTableWithColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDataAwsLakeformationPermissionsTableWithColumnsOutputReference_Override(d DataAwsLakeformationPermissionsTableWithColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationPermissionsTableWithColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"columnNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetExcludedColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedColumnNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetInternalValue(val *DataAwsLakeformationPermissionsTableWithColumns) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) SetWildcard(val interface{}) {
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ResetColumnNames() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ResetExcludedColumnNames() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedColumnNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationPermissionsTableWithColumnsOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		d,
		"resetWildcard",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_resource aws_lakeformation_resource}.
type DataAwsLakeformationResource interface {
	cdktf.TerraformDataSource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastModified() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
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

// The jsii proxy struct for DataAwsLakeformationResource
type jsiiProxy_DataAwsLakeformationResource struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsLakeformationResource) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLakeformationResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_resource aws_lakeformation_resource} Data Source.
func NewDataAwsLakeformationResource(scope constructs.Construct, id *string, config *DataAwsLakeformationResourceConfig) DataAwsLakeformationResource {
	_init_.Initialize()

	j := jsiiProxy_DataAwsLakeformationResource{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_resource aws_lakeformation_resource} Data Source.
func NewDataAwsLakeformationResource_Override(d DataAwsLakeformationResource, scope constructs.Construct, id *string, config *DataAwsLakeformationResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.DataAwsLakeformationResource",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsLakeformationResource) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsLakeformationResource) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsLakeformationResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.DataAwsLakeformationResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsLakeformationResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.DataAwsLakeformationResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsLakeformationResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsLakeformationResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsLakeformationResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsLakeformationResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsLakeformationResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsLakeformationResource) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationResource) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsLakeformationResource) ToString() *string {
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
func (d *jsiiProxy_DataAwsLakeformationResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type DataAwsLakeformationResourceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_resource#arn DataAwsLakeformationResource#arn}.
	Arn *string `json:"arn" yaml:"arn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings}.
type LakeformationDataLakeSettings interface {
	cdktf.TerraformResource
	Admins() *[]*string
	SetAdmins(val *[]*string)
	AdminsInput() *[]*string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreateDatabaseDefaultPermissions() interface{}
	SetCreateDatabaseDefaultPermissions(val interface{})
	CreateDatabaseDefaultPermissionsInput() interface{}
	CreateTableDefaultPermissions() interface{}
	SetCreateTableDefaultPermissions(val interface{})
	CreateTableDefaultPermissionsInput() interface{}
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
	TrustedResourceOwners() *[]*string
	SetTrustedResourceOwners(val *[]*string)
	TrustedResourceOwnersInput() *[]*string
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
	ResetAdmins()
	ResetCatalogId()
	ResetCreateDatabaseDefaultPermissions()
	ResetCreateTableDefaultPermissions()
	ResetOverrideLogicalId()
	ResetTrustedResourceOwners()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LakeformationDataLakeSettings
type jsiiProxy_LakeformationDataLakeSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Admins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AdminsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateDatabaseDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateDatabaseDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateTableDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateTableDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TrustedResourceOwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwnersInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
func NewLakeformationDataLakeSettings(scope constructs.Construct, id *string, config *LakeformationDataLakeSettingsConfig) LakeformationDataLakeSettings {
	_init_.Initialize()

	j := jsiiProxy_LakeformationDataLakeSettings{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
func NewLakeformationDataLakeSettings_Override(l LakeformationDataLakeSettings, scope constructs.Construct, id *string, config *LakeformationDataLakeSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetAdmins(val *[]*string) {
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetCreateDatabaseDefaultPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"createDatabaseDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetCreateTableDefaultPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"createTableDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings) SetTrustedResourceOwners(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedResourceOwners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LakeformationDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.LakeformationDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakeformationDataLakeSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.LakeformationDataLakeSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LakeformationDataLakeSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationDataLakeSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetAdmins() {
	_jsii_.InvokeVoid(
		l,
		"resetAdmins",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCreateDatabaseDefaultPermissions() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateDatabaseDefaultPermissions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCreateTableDefaultPermissions() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateTableDefaultPermissions",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LakeformationDataLakeSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetTrustedResourceOwners() {
	_jsii_.InvokeVoid(
		l,
		"resetTrustedResourceOwners",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) SynthesizeAttributes() *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) ToMetadata() interface{} {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) ToString() *string {
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
func (l *jsiiProxy_LakeformationDataLakeSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type LakeformationDataLakeSettingsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#admins LakeformationDataLakeSettings#admins}.
	Admins *[]*string `json:"admins" yaml:"admins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#catalog_id LakeformationDataLakeSettings#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// create_database_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#create_database_default_permissions LakeformationDataLakeSettings#create_database_default_permissions}
	CreateDatabaseDefaultPermissions interface{} `json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// create_table_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#create_table_default_permissions LakeformationDataLakeSettings#create_table_default_permissions}
	CreateTableDefaultPermissions interface{} `json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#trusted_resource_owners LakeformationDataLakeSettings#trusted_resource_owners}.
	TrustedResourceOwners *[]*string `json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

type LakeformationDataLakeSettingsCreateDatabaseDefaultPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#permissions LakeformationDataLakeSettings#permissions}.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#principal LakeformationDataLakeSettings#principal}.
	Principal *string `json:"principal" yaml:"principal"`
}

type LakeformationDataLakeSettingsCreateTableDefaultPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#permissions LakeformationDataLakeSettings#permissions}.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_data_lake_settings#principal LakeformationDataLakeSettings#principal}.
	Principal *string `json:"principal" yaml:"principal"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions aws_lakeformation_permissions}.
type LakeformationPermissions interface {
	cdktf.TerraformResource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CatalogResource() interface{}
	SetCatalogResource(val interface{})
	CatalogResourceInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Database() LakeformationPermissionsDatabaseOutputReference
	DatabaseInput() *LakeformationPermissionsDatabase
	DataLocation() LakeformationPermissionsDataLocationOutputReference
	DataLocationInput() *LakeformationPermissionsDataLocation
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	PermissionsInput() *[]*string
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	PermissionsWithGrantOptionInput() *[]*string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Table() LakeformationPermissionsTableOutputReference
	TableInput() *LakeformationPermissionsTable
	TableWithColumns() LakeformationPermissionsTableWithColumnsOutputReference
	TableWithColumnsInput() *LakeformationPermissionsTableWithColumns
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
	PutDatabase(value *LakeformationPermissionsDatabase)
	PutDataLocation(value *LakeformationPermissionsDataLocation)
	PutTable(value *LakeformationPermissionsTable)
	PutTableWithColumns(value *LakeformationPermissionsTableWithColumns)
	ResetCatalogId()
	ResetCatalogResource()
	ResetDatabase()
	ResetDataLocation()
	ResetOverrideLogicalId()
	ResetPermissionsWithGrantOption()
	ResetTable()
	ResetTableWithColumns()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LakeformationPermissions
type jsiiProxy_LakeformationPermissions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LakeformationPermissions) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Database() LakeformationPermissionsDatabaseOutputReference {
	var returns LakeformationPermissionsDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DatabaseInput() *LakeformationPermissionsDatabase {
	var returns *LakeformationPermissionsDatabase
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataLocation() LakeformationPermissionsDataLocationOutputReference {
	var returns LakeformationPermissionsDataLocationOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataLocationInput() *LakeformationPermissionsDataLocation {
	var returns *LakeformationPermissionsDataLocation
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsWithGrantOptionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Table() LakeformationPermissionsTableOutputReference {
	var returns LakeformationPermissionsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableInput() *LakeformationPermissionsTable {
	var returns *LakeformationPermissionsTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableWithColumns() LakeformationPermissionsTableWithColumnsOutputReference {
	var returns LakeformationPermissionsTableWithColumnsOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableWithColumnsInput() *LakeformationPermissionsTableWithColumns {
	var returns *LakeformationPermissionsTableWithColumns
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions aws_lakeformation_permissions} Resource.
func NewLakeformationPermissions(scope constructs.Construct, id *string, config *LakeformationPermissionsConfig) LakeformationPermissions {
	_init_.Initialize()

	j := jsiiProxy_LakeformationPermissions{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions aws_lakeformation_permissions} Resource.
func NewLakeformationPermissions_Override(l LakeformationPermissions, scope constructs.Construct, id *string, config *LakeformationPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissions",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetCatalogResource(val interface{}) {
	_jsii_.Set(
		j,
		"catalogResource",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetPermissions(val *[]*string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetPermissionsWithGrantOption(val *[]*string) {
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions) SetProvider(val cdktf.TerraformProvider) {
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
func LakeformationPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.LakeformationPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakeformationPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.LakeformationPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationPermissions) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationPermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutDatabase(value *LakeformationPermissionsDatabase) {
	_jsii_.InvokeVoid(
		l,
		"putDatabase",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutDataLocation(value *LakeformationPermissionsDataLocation) {
	_jsii_.InvokeVoid(
		l,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutTable(value *LakeformationPermissionsTable) {
	_jsii_.InvokeVoid(
		l,
		"putTable",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutTableWithColumns(value *LakeformationPermissionsTableWithColumns) {
	_jsii_.InvokeVoid(
		l,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetCatalogResource() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogResource",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetDatabase() {
	_jsii_.InvokeVoid(
		l,
		"resetDatabase",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetDataLocation() {
	_jsii_.InvokeVoid(
		l,
		"resetDataLocation",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LakeformationPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetPermissionsWithGrantOption() {
	_jsii_.InvokeVoid(
		l,
		"resetPermissionsWithGrantOption",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetTable() {
	_jsii_.InvokeVoid(
		l,
		"resetTable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		l,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) SynthesizeAttributes() *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissions) ToMetadata() interface{} {
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
func (l *jsiiProxy_LakeformationPermissions) ToString() *string {
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
func (l *jsiiProxy_LakeformationPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type LakeformationPermissionsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#permissions LakeformationPermissions#permissions}.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#principal LakeformationPermissions#principal}.
	Principal *string `json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_resource LakeformationPermissions#catalog_resource}.
	CatalogResource interface{} `json:"catalogResource" yaml:"catalogResource"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#database LakeformationPermissions#database}
	Database *LakeformationPermissionsDatabase `json:"database" yaml:"database"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#data_location LakeformationPermissions#data_location}
	DataLocation *LakeformationPermissionsDataLocation `json:"dataLocation" yaml:"dataLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#permissions_with_grant_option LakeformationPermissions#permissions_with_grant_option}.
	PermissionsWithGrantOption *[]*string `json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#table LakeformationPermissions#table}
	Table *LakeformationPermissionsTable `json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#table_with_columns LakeformationPermissions#table_with_columns}
	TableWithColumns *LakeformationPermissionsTableWithColumns `json:"tableWithColumns" yaml:"tableWithColumns"`
}

type LakeformationPermissionsDataLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#arn LakeformationPermissions#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
}

type LakeformationPermissionsDataLocationOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	InternalValue() *LakeformationPermissionsDataLocation
	SetInternalValue(val *LakeformationPermissionsDataLocation)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
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
	ResetCatalogId()
}

// The jsii proxy struct for LakeformationPermissionsDataLocationOutputReference
type jsiiProxy_LakeformationPermissionsDataLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) InternalValue() *LakeformationPermissionsDataLocation {
	var returns *LakeformationPermissionsDataLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLakeformationPermissionsDataLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LakeformationPermissionsDataLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LakeformationPermissionsDataLocationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsDataLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLakeformationPermissionsDataLocationOutputReference_Override(l LakeformationPermissionsDataLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsDataLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetInternalValue(val *LakeformationPermissionsDataLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissionsDataLocationOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

type LakeformationPermissionsDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
}

type LakeformationPermissionsDatabaseOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	InternalValue() *LakeformationPermissionsDatabase
	SetInternalValue(val *LakeformationPermissionsDatabase)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetCatalogId()
}

// The jsii proxy struct for LakeformationPermissionsDatabaseOutputReference
type jsiiProxy_LakeformationPermissionsDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) InternalValue() *LakeformationPermissionsDatabase {
	var returns *LakeformationPermissionsDatabase
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLakeformationPermissionsDatabaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LakeformationPermissionsDatabaseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LakeformationPermissionsDatabaseOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLakeformationPermissionsDatabaseOutputReference_Override(l LakeformationPermissionsDatabaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetInternalValue(val *LakeformationPermissionsDatabase) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissionsDatabaseOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

type LakeformationPermissionsTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#database_name LakeformationPermissions#database_name}.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#wildcard LakeformationPermissions#wildcard}.
	Wildcard interface{} `json:"wildcard" yaml:"wildcard"`
}

type LakeformationPermissionsTableOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	InternalValue() *LakeformationPermissionsTable
	SetInternalValue(val *LakeformationPermissionsTable)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	ResetCatalogId()
	ResetName()
	ResetWildcard()
}

// The jsii proxy struct for LakeformationPermissionsTableOutputReference
type jsiiProxy_LakeformationPermissionsTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) InternalValue() *LakeformationPermissionsTable {
	var returns *LakeformationPermissionsTable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}

func NewLakeformationPermissionsTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LakeformationPermissionsTableOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LakeformationPermissionsTableOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLakeformationPermissionsTableOutputReference_Override(l LakeformationPermissionsTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetInternalValue(val *LakeformationPermissionsTable) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableOutputReference) SetWildcard(val interface{}) {
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissionsTableOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		l,
		"resetWildcard",
		nil, // no parameters
	)
}

type LakeformationPermissionsTableWithColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#database_name LakeformationPermissions#database_name}.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#column_names LakeformationPermissions#column_names}.
	ColumnNames *[]*string `json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#excluded_column_names LakeformationPermissions#excluded_column_names}.
	ExcludedColumnNames *[]*string `json:"excludedColumnNames" yaml:"excludedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#wildcard LakeformationPermissions#wildcard}.
	Wildcard interface{} `json:"wildcard" yaml:"wildcard"`
}

type LakeformationPermissionsTableWithColumnsOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	ColumnNames() *[]*string
	SetColumnNames(val *[]*string)
	ColumnNamesInput() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	ExcludedColumnNames() *[]*string
	SetExcludedColumnNames(val *[]*string)
	ExcludedColumnNamesInput() *[]*string
	InternalValue() *LakeformationPermissionsTableWithColumns
	SetInternalValue(val *LakeformationPermissionsTableWithColumns)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	ResetCatalogId()
	ResetColumnNames()
	ResetExcludedColumnNames()
	ResetWildcard()
}

// The jsii proxy struct for LakeformationPermissionsTableWithColumnsOutputReference
type jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ExcludedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ExcludedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) InternalValue() *LakeformationPermissionsTableWithColumns {
	var returns *LakeformationPermissionsTableWithColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}

func NewLakeformationPermissionsTableWithColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) LakeformationPermissionsTableWithColumnsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsTableWithColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewLakeformationPermissionsTableWithColumnsOutputReference_Override(l LakeformationPermissionsTableWithColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationPermissionsTableWithColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"columnNames",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetExcludedColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedColumnNames",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetInternalValue(val *LakeformationPermissionsTableWithColumns) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) SetWildcard(val interface{}) {
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ResetColumnNames() {
	_jsii_.InvokeVoid(
		l,
		"resetColumnNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ResetExcludedColumnNames() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludedColumnNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissionsTableWithColumnsOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		l,
		"resetWildcard",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource aws_lakeformation_resource}.
type LakeformationResource interface {
	cdktf.TerraformResource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastModified() *string
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
	ResetRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LakeformationResource
type jsiiProxy_LakeformationResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LakeformationResource) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource aws_lakeformation_resource} Resource.
func NewLakeformationResource(scope constructs.Construct, id *string, config *LakeformationResourceConfig) LakeformationResource {
	_init_.Initialize()

	j := jsiiProxy_LakeformationResource{}

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource aws_lakeformation_resource} Resource.
func NewLakeformationResource_Override(l LakeformationResource, scope constructs.Construct, id *string, config *LakeformationResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.lakeformation.LakeformationResource",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakeformationResource) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_LakeformationResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakeformationResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakeformationResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakeformationResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LakeformationResource) SetRoleArn(val *string) {
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
func LakeformationResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.lakeformation.LakeformationResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakeformationResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.lakeformation.LakeformationResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LakeformationResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LakeformationResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (l *jsiiProxy_LakeformationResource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LakeformationResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LakeformationResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (l *jsiiProxy_LakeformationResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (l *jsiiProxy_LakeformationResource) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LakeformationResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (l *jsiiProxy_LakeformationResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LakeformationResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LakeformationResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationResource) ResetRoleArn() {
	_jsii_.InvokeVoid(
		l,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationResource) SynthesizeAttributes() *map[string]interface{} {
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
func (l *jsiiProxy_LakeformationResource) ToMetadata() interface{} {
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
func (l *jsiiProxy_LakeformationResource) ToString() *string {
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
func (l *jsiiProxy_LakeformationResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Lake Formation.
type LakeformationResourceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource#arn LakeformationResource#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource#role_arn LakeformationResource#role_arn}.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}
