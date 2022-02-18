package servicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/servicecatalog/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint aws_servicecatalog_constraint}.
type DataAwsServicecatalogConstraint interface {
	cdktf.TerraformDataSource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Owner() *string
	Parameters() *string
	PortfolioId() *string
	ProductId() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
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
	ResetAcceptLanguage()
	ResetDescription()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsServicecatalogConstraint
type jsiiProxy_DataAwsServicecatalogConstraint struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint aws_servicecatalog_constraint} Data Source.
func NewDataAwsServicecatalogConstraint(scope constructs.Construct, id *string, config *DataAwsServicecatalogConstraintConfig) DataAwsServicecatalogConstraint {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogConstraint{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogConstraint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint aws_servicecatalog_constraint} Data Source.
func NewDataAwsServicecatalogConstraint_Override(d DataAwsServicecatalogConstraint, scope constructs.Construct, id *string, config *DataAwsServicecatalogConstraintConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogConstraint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogConstraint) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsServicecatalogConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsServicecatalogConstraint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogConstraint",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogConstraint) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogConstraint) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogConstraint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogConstraint) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) ToString() *string {
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
func (d *jsiiProxy_DataAwsServicecatalogConstraint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type DataAwsServicecatalogConstraintConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint#id DataAwsServicecatalogConstraint#id}.
	Id *string `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint#accept_language DataAwsServicecatalogConstraint#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_constraint#description DataAwsServicecatalogConstraint#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_launch_paths aws_servicecatalog_launch_paths}.
type DataAwsServicecatalogLaunchPaths interface {
	cdktf.TerraformDataSource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	Summaries(index *string) DataAwsServicecatalogLaunchPathsSummaries
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsServicecatalogLaunchPaths
type jsiiProxy_DataAwsServicecatalogLaunchPaths struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_launch_paths aws_servicecatalog_launch_paths} Data Source.
func NewDataAwsServicecatalogLaunchPaths(scope constructs.Construct, id *string, config *DataAwsServicecatalogLaunchPathsConfig) DataAwsServicecatalogLaunchPaths {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogLaunchPaths{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPaths",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_launch_paths aws_servicecatalog_launch_paths} Data Source.
func NewDataAwsServicecatalogLaunchPaths_Override(d DataAwsServicecatalogLaunchPaths, scope constructs.Construct, id *string, config *DataAwsServicecatalogLaunchPathsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPaths",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPaths) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsServicecatalogLaunchPaths_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPaths",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsServicecatalogLaunchPaths_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPaths",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) Summaries(index *string) DataAwsServicecatalogLaunchPathsSummaries {
	var returns DataAwsServicecatalogLaunchPathsSummaries

	_jsii_.Invoke(
		d,
		"summaries",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) ToString() *string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPaths) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type DataAwsServicecatalogLaunchPathsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_launch_paths#product_id DataAwsServicecatalogLaunchPaths#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_launch_paths#accept_language DataAwsServicecatalogLaunchPaths#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
}

type DataAwsServicecatalogLaunchPathsSummaries interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConstraintSummaries() cdktf.IResolvable
	Name() *string
	PathId() *string
	Tags() *map[string]*string
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

// The jsii proxy struct for DataAwsServicecatalogLaunchPathsSummaries
type jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) ConstraintSummaries() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"constraintSummaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsServicecatalogLaunchPathsSummaries(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsServicecatalogLaunchPathsSummaries {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPathsSummaries",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsServicecatalogLaunchPathsSummaries_Override(d DataAwsServicecatalogLaunchPathsSummaries, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPathsSummaries",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummaries) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Description() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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

// The jsii proxy struct for DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries
type jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsServicecatalogLaunchPathsSummariesConstraintSummaries(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsServicecatalogLaunchPathsSummariesConstraintSummaries_Override(d DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesConstraintSummaries) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio aws_servicecatalog_portfolio}.
type DataAwsServicecatalogPortfolio interface {
	cdktf.TerraformDataSource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
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
	Name() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProviderName() *string
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsServicecatalogPortfolio
type jsiiProxy_DataAwsServicecatalogPortfolio struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio aws_servicecatalog_portfolio} Data Source.
func NewDataAwsServicecatalogPortfolio(scope constructs.Construct, id *string, config *DataAwsServicecatalogPortfolioConfig) DataAwsServicecatalogPortfolio {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogPortfolio{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolio",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio aws_servicecatalog_portfolio} Data Source.
func NewDataAwsServicecatalogPortfolio_Override(d DataAwsServicecatalogPortfolio, scope constructs.Construct, id *string, config *DataAwsServicecatalogPortfolioConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolio",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolio) SetTags(val *map[string]*string) {
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
func DataAwsServicecatalogPortfolio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsServicecatalogPortfolio_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolio",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolio) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ToString() *string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolio) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type DataAwsServicecatalogPortfolioConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio#id DataAwsServicecatalogPortfolio#id}.
	Id *string `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio#accept_language DataAwsServicecatalogPortfolio#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio#tags DataAwsServicecatalogPortfolio#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints aws_servicecatalog_portfolio_constraints}.
type DataAwsServicecatalogPortfolioConstraints interface {
	cdktf.TerraformDataSource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	Details(index *string) DataAwsServicecatalogPortfolioConstraintsDetails
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetProductId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsServicecatalogPortfolioConstraints
type jsiiProxy_DataAwsServicecatalogPortfolioConstraints struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints aws_servicecatalog_portfolio_constraints} Data Source.
func NewDataAwsServicecatalogPortfolioConstraints(scope constructs.Construct, id *string, config *DataAwsServicecatalogPortfolioConstraintsConfig) DataAwsServicecatalogPortfolioConstraints {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogPortfolioConstraints{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraints",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints aws_servicecatalog_portfolio_constraints} Data Source.
func NewDataAwsServicecatalogPortfolioConstraints_Override(d DataAwsServicecatalogPortfolioConstraints, scope constructs.Construct, id *string, config *DataAwsServicecatalogPortfolioConstraintsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraints",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsServicecatalogPortfolioConstraints_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraints",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsServicecatalogPortfolioConstraints_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraints",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) Details(index *string) DataAwsServicecatalogPortfolioConstraintsDetails {
	var returns DataAwsServicecatalogPortfolioConstraintsDetails

	_jsii_.Invoke(
		d,
		"details",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ResetProductId() {
	_jsii_.InvokeVoid(
		d,
		"resetProductId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ToString() *string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraints) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type DataAwsServicecatalogPortfolioConstraintsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints#portfolio_id DataAwsServicecatalogPortfolioConstraints#portfolio_id}.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints#accept_language DataAwsServicecatalogPortfolioConstraints#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_portfolio_constraints#product_id DataAwsServicecatalogPortfolioConstraints#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
}

type DataAwsServicecatalogPortfolioConstraintsDetails interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ConstraintId() *string
	Description() *string
	Owner() *string
	PortfolioId() *string
	ProductId() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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

// The jsii proxy struct for DataAwsServicecatalogPortfolioConstraintsDetails
type jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) ConstraintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsServicecatalogPortfolioConstraintsDetails(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsServicecatalogPortfolioConstraintsDetails {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraintsDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsServicecatalogPortfolioConstraintsDetails_Override(d DataAwsServicecatalogPortfolioConstraintsDetails, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogPortfolioConstraintsDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogPortfolioConstraintsDetails) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product aws_servicecatalog_product}.
type DataAwsServicecatalogProduct interface {
	cdktf.TerraformDataSource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Distributor() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HasDefaultPath() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	Owner() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	SupportDescription() *string
	SupportEmail() *string
	SupportUrl() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsServicecatalogProduct
type jsiiProxy_DataAwsServicecatalogProduct struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Distributor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) HasDefaultPath() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasDefaultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SupportDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SupportEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product aws_servicecatalog_product} Data Source.
func NewDataAwsServicecatalogProduct(scope constructs.Construct, id *string, config *DataAwsServicecatalogProductConfig) DataAwsServicecatalogProduct {
	_init_.Initialize()

	j := jsiiProxy_DataAwsServicecatalogProduct{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product aws_servicecatalog_product} Data Source.
func NewDataAwsServicecatalogProduct_Override(d DataAwsServicecatalogProduct, scope constructs.Construct, id *string, config *DataAwsServicecatalogProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogProduct",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsServicecatalogProduct) SetTags(val *map[string]*string) {
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
func DataAwsServicecatalogProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsServicecatalogProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.DataAwsServicecatalogProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogProduct) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsServicecatalogProduct) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsServicecatalogProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogProduct) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsServicecatalogProduct) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) ToString() *string {
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
func (d *jsiiProxy_DataAwsServicecatalogProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type DataAwsServicecatalogProductConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product#id DataAwsServicecatalogProduct#id}.
	Id *string `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product#accept_language DataAwsServicecatalogProduct#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/servicecatalog_product#tags DataAwsServicecatalogProduct#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_budget_resource_association aws_servicecatalog_budget_resource_association}.
type ServicecatalogBudgetResourceAssociation interface {
	cdktf.TerraformResource
	BudgetName() *string
	SetBudgetName(val *string)
	BudgetNameInput() *string
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
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
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

// The jsii proxy struct for ServicecatalogBudgetResourceAssociation
type jsiiProxy_ServicecatalogBudgetResourceAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) BudgetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) BudgetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_budget_resource_association aws_servicecatalog_budget_resource_association} Resource.
func NewServicecatalogBudgetResourceAssociation(scope constructs.Construct, id *string, config *ServicecatalogBudgetResourceAssociationConfig) ServicecatalogBudgetResourceAssociation {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogBudgetResourceAssociation{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogBudgetResourceAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_budget_resource_association aws_servicecatalog_budget_resource_association} Resource.
func NewServicecatalogBudgetResourceAssociation_Override(s ServicecatalogBudgetResourceAssociation, scope constructs.Construct, id *string, config *ServicecatalogBudgetResourceAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogBudgetResourceAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetBudgetName(val *string) {
	_jsii_.Set(
		j,
		"budgetName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogBudgetResourceAssociation) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogBudgetResourceAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogBudgetResourceAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogBudgetResourceAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogBudgetResourceAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogBudgetResourceAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogBudgetResourceAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_budget_resource_association#budget_name ServicecatalogBudgetResourceAssociation#budget_name}.
	BudgetName *string `json:"budgetName" yaml:"budgetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_budget_resource_association#resource_id ServicecatalogBudgetResourceAssociation#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint aws_servicecatalog_constraint}.
type ServicecatalogConstraint interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	Owner() *string
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAcceptLanguage()
	ResetDescription()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogConstraint
type jsiiProxy_ServicecatalogConstraint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogConstraint) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint aws_servicecatalog_constraint} Resource.
func NewServicecatalogConstraint(scope constructs.Construct, id *string, config *ServicecatalogConstraintConfig) ServicecatalogConstraint {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogConstraint{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogConstraint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint aws_servicecatalog_constraint} Resource.
func NewServicecatalogConstraint_Override(s ServicecatalogConstraint, scope constructs.Construct, id *string, config *ServicecatalogConstraintConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogConstraint",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogConstraint) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogConstraint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogConstraint",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogConstraint) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogConstraint) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogConstraint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogConstraint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogConstraintConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#parameters ServicecatalogConstraint#parameters}.
	Parameters *string `json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#portfolio_id ServicecatalogConstraint#portfolio_id}.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#product_id ServicecatalogConstraint#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#type ServicecatalogConstraint#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#accept_language ServicecatalogConstraint#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#description ServicecatalogConstraint#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_organizations_access aws_servicecatalog_organizations_access}.
type ServicecatalogOrganizationsAccess interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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

// The jsii proxy struct for ServicecatalogOrganizationsAccess
type jsiiProxy_ServicecatalogOrganizationsAccess struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_organizations_access aws_servicecatalog_organizations_access} Resource.
func NewServicecatalogOrganizationsAccess(scope constructs.Construct, id *string, config *ServicecatalogOrganizationsAccessConfig) ServicecatalogOrganizationsAccess {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogOrganizationsAccess{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogOrganizationsAccess",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_organizations_access aws_servicecatalog_organizations_access} Resource.
func NewServicecatalogOrganizationsAccess_Override(s ServicecatalogOrganizationsAccess, scope constructs.Construct, id *string, config *ServicecatalogOrganizationsAccessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogOrganizationsAccess",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogOrganizationsAccess) SetProvider(val cdktf.TerraformProvider) {
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
func ServicecatalogOrganizationsAccess_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogOrganizationsAccess",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogOrganizationsAccess_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogOrganizationsAccess",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogOrganizationsAccess) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogOrganizationsAccess) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogOrganizationsAccessConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_organizations_access#enabled ServicecatalogOrganizationsAccess#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio aws_servicecatalog_portfolio}.
type ServicecatalogPortfolio interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
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
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ServicecatalogPortfolioTimeoutsOutputReference
	TimeoutsInput() *ServicecatalogPortfolioTimeouts
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
	PutTimeouts(value *ServicecatalogPortfolioTimeouts)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogPortfolio
type jsiiProxy_ServicecatalogPortfolio struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogPortfolio) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) Timeouts() ServicecatalogPortfolioTimeoutsOutputReference {
	var returns ServicecatalogPortfolioTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolio) TimeoutsInput() *ServicecatalogPortfolioTimeouts {
	var returns *ServicecatalogPortfolioTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio aws_servicecatalog_portfolio} Resource.
func NewServicecatalogPortfolio(scope constructs.Construct, id *string, config *ServicecatalogPortfolioConfig) ServicecatalogPortfolio {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogPortfolio{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolio",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio aws_servicecatalog_portfolio} Resource.
func NewServicecatalogPortfolio_Override(s ServicecatalogPortfolio, scope constructs.Construct, id *string, config *ServicecatalogPortfolioConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolio",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetProviderName(val *string) {
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolio) SetTagsAll(val *map[string]*string) {
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
func ServicecatalogPortfolio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogPortfolio_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolio",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) PutTimeouts(value *ServicecatalogPortfolioTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolio) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogPortfolio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolio) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogPortfolioConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#name ServicecatalogPortfolio#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#provider_name ServicecatalogPortfolio#provider_name}.
	ProviderName *string `json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#description ServicecatalogPortfolio#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#tags ServicecatalogPortfolio#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#tags_all ServicecatalogPortfolio#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#timeouts ServicecatalogPortfolio#timeouts}
	Timeouts *ServicecatalogPortfolioTimeouts `json:"timeouts" yaml:"timeouts"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share}.
type ServicecatalogPortfolioShare interface {
	cdktf.TerraformResource
	Accepted() cdktf.IResolvable
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	PrincipalId() *string
	SetPrincipalId(val *string)
	PrincipalIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ShareTagOptions() interface{}
	SetShareTagOptions(val interface{})
	ShareTagOptionsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WaitForAcceptance() interface{}
	SetWaitForAcceptance(val interface{})
	WaitForAcceptanceInput() interface{}
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetShareTagOptions()
	ResetWaitForAcceptance()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogPortfolioShare
type jsiiProxy_ServicecatalogPortfolioShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Accepted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"accepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) PrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ShareTagOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareTagOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) ShareTagOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareTagOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) WaitForAcceptance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForAcceptance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) WaitForAcceptanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForAcceptanceInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share} Resource.
func NewServicecatalogPortfolioShare(scope constructs.Construct, id *string, config *ServicecatalogPortfolioShareConfig) ServicecatalogPortfolioShare {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogPortfolioShare{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share aws_servicecatalog_portfolio_share} Resource.
func NewServicecatalogPortfolioShare_Override(s ServicecatalogPortfolioShare, scope constructs.Construct, id *string, config *ServicecatalogPortfolioShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetPrincipalId(val *string) {
	_jsii_.Set(
		j,
		"principalId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetShareTagOptions(val interface{}) {
	_jsii_.Set(
		j,
		"shareTagOptions",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioShare) SetWaitForAcceptance(val interface{}) {
	_jsii_.Set(
		j,
		"waitForAcceptance",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogPortfolioShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogPortfolioShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetShareTagOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetShareTagOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) ResetWaitForAcceptance() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForAcceptance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogPortfolioShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogPortfolioShareConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#portfolio_id ServicecatalogPortfolioShare#portfolio_id}.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#principal_id ServicecatalogPortfolioShare#principal_id}.
	PrincipalId *string `json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#type ServicecatalogPortfolioShare#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#accept_language ServicecatalogPortfolioShare#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#share_tag_options ServicecatalogPortfolioShare#share_tag_options}.
	ShareTagOptions interface{} `json:"shareTagOptions" yaml:"shareTagOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio_share#wait_for_acceptance ServicecatalogPortfolioShare#wait_for_acceptance}.
	WaitForAcceptance interface{} `json:"waitForAcceptance" yaml:"waitForAcceptance"`
}

type ServicecatalogPortfolioTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#create ServicecatalogPortfolio#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#delete ServicecatalogPortfolio#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_portfolio#update ServicecatalogPortfolio#update}.
	Update *string `json:"update" yaml:"update"`
}

type ServicecatalogPortfolioTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *ServicecatalogPortfolioTimeouts
	SetInternalValue(val *ServicecatalogPortfolioTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for ServicecatalogPortfolioTimeoutsOutputReference
type jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) InternalValue() *ServicecatalogPortfolioTimeouts {
	var returns *ServicecatalogPortfolioTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewServicecatalogPortfolioTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ServicecatalogPortfolioTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewServicecatalogPortfolioTimeoutsOutputReference_Override(s ServicecatalogPortfolioTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPortfolioTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetInternalValue(val *ServicecatalogPortfolioTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPortfolioTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association aws_servicecatalog_principal_portfolio_association}.
type ServicecatalogPrincipalPortfolioAssociation interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	PrincipalArn() *string
	SetPrincipalArn(val *string)
	PrincipalArnInput() *string
	PrincipalType() *string
	SetPrincipalType(val *string)
	PrincipalTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetPrincipalType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogPrincipalPortfolioAssociation
type jsiiProxy_ServicecatalogPrincipalPortfolioAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PrincipalArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PrincipalArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PrincipalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) PrincipalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association aws_servicecatalog_principal_portfolio_association} Resource.
func NewServicecatalogPrincipalPortfolioAssociation(scope constructs.Construct, id *string, config *ServicecatalogPrincipalPortfolioAssociationConfig) ServicecatalogPrincipalPortfolioAssociation {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogPrincipalPortfolioAssociation{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPrincipalPortfolioAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association aws_servicecatalog_principal_portfolio_association} Resource.
func NewServicecatalogPrincipalPortfolioAssociation_Override(s ServicecatalogPrincipalPortfolioAssociation, scope constructs.Construct, id *string, config *ServicecatalogPrincipalPortfolioAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogPrincipalPortfolioAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetPrincipalArn(val *string) {
	_jsii_.Set(
		j,
		"principalArn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetPrincipalType(val *string) {
	_jsii_.Set(
		j,
		"principalType",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SetProvider(val cdktf.TerraformProvider) {
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
func ServicecatalogPrincipalPortfolioAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogPrincipalPortfolioAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogPrincipalPortfolioAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogPrincipalPortfolioAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ResetPrincipalType() {
	_jsii_.InvokeVoid(
		s,
		"resetPrincipalType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogPrincipalPortfolioAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogPrincipalPortfolioAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#portfolio_id ServicecatalogPrincipalPortfolioAssociation#portfolio_id}.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#principal_arn ServicecatalogPrincipalPortfolioAssociation#principal_arn}.
	PrincipalArn *string `json:"principalArn" yaml:"principalArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#accept_language ServicecatalogPrincipalPortfolioAssociation#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#principal_type ServicecatalogPrincipalPortfolioAssociation#principal_type}.
	PrincipalType *string `json:"principalType" yaml:"principalType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product aws_servicecatalog_product}.
type ServicecatalogProduct interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Distributor() *string
	SetDistributor(val *string)
	DistributorInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HasDefaultPath() cdktf.IResolvable
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
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProvisioningArtifactParameters() ServicecatalogProductProvisioningArtifactParametersOutputReference
	ProvisioningArtifactParametersInput() *ServicecatalogProductProvisioningArtifactParameters
	RawOverrides() interface{}
	Status() *string
	SupportDescription() *string
	SetSupportDescription(val *string)
	SupportDescriptionInput() *string
	SupportEmail() *string
	SetSupportEmail(val *string)
	SupportEmailInput() *string
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutProvisioningArtifactParameters(value *ServicecatalogProductProvisioningArtifactParameters)
	ResetAcceptLanguage()
	ResetDescription()
	ResetDistributor()
	ResetOverrideLogicalId()
	ResetSupportDescription()
	ResetSupportEmail()
	ResetSupportUrl()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogProduct
type jsiiProxy_ServicecatalogProduct struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Distributor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) DistributorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) HasDefaultPath() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasDefaultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) ProvisioningArtifactParameters() ServicecatalogProductProvisioningArtifactParametersOutputReference {
	var returns ServicecatalogProductProvisioningArtifactParametersOutputReference
	_jsii_.Get(
		j,
		"provisioningArtifactParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) ProvisioningArtifactParametersInput() *ServicecatalogProductProvisioningArtifactParameters {
	var returns *ServicecatalogProductProvisioningArtifactParameters
	_jsii_.Get(
		j,
		"provisioningArtifactParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProduct) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product aws_servicecatalog_product} Resource.
func NewServicecatalogProduct(scope constructs.Construct, id *string, config *ServicecatalogProductConfig) ServicecatalogProduct {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProduct{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product aws_servicecatalog_product} Resource.
func NewServicecatalogProduct_Override(s ServicecatalogProduct, scope constructs.Construct, id *string, config *ServicecatalogProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProduct",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetDistributor(val *string) {
	_jsii_.Set(
		j,
		"distributor",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetSupportDescription(val *string) {
	_jsii_.Set(
		j,
		"supportDescription",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetSupportEmail(val *string) {
	_jsii_.Set(
		j,
		"supportEmail",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetSupportUrl(val *string) {
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProduct) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogProduct) PutProvisioningArtifactParameters(value *ServicecatalogProductProvisioningArtifactParameters) {
	_jsii_.InvokeVoid(
		s,
		"putProvisioningArtifactParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetDistributor() {
	_jsii_.InvokeVoid(
		s,
		"resetDistributor",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetSupportDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetSupportEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportEmail",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogProductConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#name ServicecatalogProduct#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#owner ServicecatalogProduct#owner}.
	Owner *string `json:"owner" yaml:"owner"`
	// provisioning_artifact_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#provisioning_artifact_parameters ServicecatalogProduct#provisioning_artifact_parameters}
	ProvisioningArtifactParameters *ServicecatalogProductProvisioningArtifactParameters `json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#type ServicecatalogProduct#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#accept_language ServicecatalogProduct#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#description ServicecatalogProduct#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#distributor ServicecatalogProduct#distributor}.
	Distributor *string `json:"distributor" yaml:"distributor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#support_description ServicecatalogProduct#support_description}.
	SupportDescription *string `json:"supportDescription" yaml:"supportDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#support_email ServicecatalogProduct#support_email}.
	SupportEmail *string `json:"supportEmail" yaml:"supportEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#support_url ServicecatalogProduct#support_url}.
	SupportUrl *string `json:"supportUrl" yaml:"supportUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#tags ServicecatalogProduct#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#tags_all ServicecatalogProduct#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association aws_servicecatalog_product_portfolio_association}.
type ServicecatalogProductPortfolioAssociation interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
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
	PortfolioId() *string
	SetPortfolioId(val *string)
	PortfolioIdInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SourcePortfolioId() *string
	SetSourcePortfolioId(val *string)
	SourcePortfolioIdInput() *string
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
	ResetAcceptLanguage()
	ResetOverrideLogicalId()
	ResetSourcePortfolioId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogProductPortfolioAssociation
type jsiiProxy_ServicecatalogProductPortfolioAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) PortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SourcePortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SourcePortfolioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortfolioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association aws_servicecatalog_product_portfolio_association} Resource.
func NewServicecatalogProductPortfolioAssociation(scope constructs.Construct, id *string, config *ServicecatalogProductPortfolioAssociationConfig) ServicecatalogProductPortfolioAssociation {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProductPortfolioAssociation{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProductPortfolioAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association aws_servicecatalog_product_portfolio_association} Resource.
func NewServicecatalogProductPortfolioAssociation_Override(s ServicecatalogProductPortfolioAssociation, scope constructs.Construct, id *string, config *ServicecatalogProductPortfolioAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProductPortfolioAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductPortfolioAssociation) SetSourcePortfolioId(val *string) {
	_jsii_.Set(
		j,
		"sourcePortfolioId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogProductPortfolioAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogProductPortfolioAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogProductPortfolioAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogProductPortfolioAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ResetSourcePortfolioId() {
	_jsii_.InvokeVoid(
		s,
		"resetSourcePortfolioId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogProductPortfolioAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogProductPortfolioAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association#portfolio_id ServicecatalogProductPortfolioAssociation#portfolio_id}.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association#product_id ServicecatalogProductPortfolioAssociation#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association#accept_language ServicecatalogProductPortfolioAssociation#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product_portfolio_association#source_portfolio_id ServicecatalogProductPortfolioAssociation#source_portfolio_id}.
	SourcePortfolioId *string `json:"sourcePortfolioId" yaml:"sourcePortfolioId"`
}

type ServicecatalogProductProvisioningArtifactParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#description ServicecatalogProduct#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#disable_template_validation ServicecatalogProduct#disable_template_validation}.
	DisableTemplateValidation interface{} `json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#name ServicecatalogProduct#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#template_physical_id ServicecatalogProduct#template_physical_id}.
	TemplatePhysicalId *string `json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#template_url ServicecatalogProduct#template_url}.
	TemplateUrl *string `json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_product#type ServicecatalogProduct#type}.
	Type *string `json:"type" yaml:"type"`
}

type ServicecatalogProductProvisioningArtifactParametersOutputReference interface {
	cdktf.ComplexObject
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableTemplateValidation() interface{}
	SetDisableTemplateValidation(val interface{})
	DisableTemplateValidationInput() interface{}
	InternalValue() *ServicecatalogProductProvisioningArtifactParameters
	SetInternalValue(val *ServicecatalogProductProvisioningArtifactParameters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	TemplatePhysicalId() *string
	SetTemplatePhysicalId(val *string)
	TemplatePhysicalIdInput() *string
	TemplateUrl() *string
	SetTemplateUrl(val *string)
	TemplateUrlInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDescription()
	ResetDisableTemplateValidation()
	ResetName()
	ResetTemplatePhysicalId()
	ResetTemplateUrl()
	ResetType()
}

// The jsii proxy struct for ServicecatalogProductProvisioningArtifactParametersOutputReference
type jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) DisableTemplateValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) DisableTemplateValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) InternalValue() *ServicecatalogProductProvisioningArtifactParameters {
	var returns *ServicecatalogProductProvisioningArtifactParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TemplatePhysicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TemplatePhysicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TemplateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewServicecatalogProductProvisioningArtifactParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ServicecatalogProductProvisioningArtifactParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProductProvisioningArtifactParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewServicecatalogProductProvisioningArtifactParametersOutputReference_Override(s ServicecatalogProductProvisioningArtifactParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProductProvisioningArtifactParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetDisableTemplateValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableTemplateValidation",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetInternalValue(val *ServicecatalogProductProvisioningArtifactParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetTemplatePhysicalId(val *string) {
	_jsii_.Set(
		j,
		"templatePhysicalId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetTemplateUrl(val *string) {
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetDisableTemplateValidation() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableTemplateValidation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetTemplatePhysicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplatePhysicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetTemplateUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplateUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProductProvisioningArtifactParametersOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product}.
type ServicecatalogProvisionedProduct interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CloudwatchDashboardNames() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IgnoreErrors() interface{}
	SetIgnoreErrors(val interface{})
	IgnoreErrorsInput() interface{}
	LastProvisioningRecordId() *string
	LastRecordId() *string
	LastSuccessfulProvisioningRecordId() *string
	LaunchRoleArn() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	NotificationArnsInput() *[]*string
	PathId() *string
	SetPathId(val *string)
	PathIdInput() *string
	PathName() *string
	SetPathName(val *string)
	PathNameInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	ProductName() *string
	SetProductName(val *string)
	ProductNameInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	ProvisioningArtifactIdInput() *string
	ProvisioningArtifactName() *string
	SetProvisioningArtifactName(val *string)
	ProvisioningArtifactNameInput() *string
	ProvisioningParameters() interface{}
	SetProvisioningParameters(val interface{})
	ProvisioningParametersInput() interface{}
	RawOverrides() interface{}
	RetainPhysicalResources() interface{}
	SetRetainPhysicalResources(val interface{})
	RetainPhysicalResourcesInput() interface{}
	StackSetProvisioningPreferences() ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
	StackSetProvisioningPreferencesInput() *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	Status() *string
	StatusMessage() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ServicecatalogProvisionedProductTimeoutsOutputReference
	TimeoutsInput() *ServicecatalogProvisionedProductTimeouts
	Type() *string
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
	PutStackSetProvisioningPreferences(value *ServicecatalogProvisionedProductStackSetProvisioningPreferences)
	PutTimeouts(value *ServicecatalogProvisionedProductTimeouts)
	ResetAcceptLanguage()
	ResetIgnoreErrors()
	ResetNotificationArns()
	ResetOverrideLogicalId()
	ResetPathId()
	ResetPathName()
	ResetProductId()
	ResetProductName()
	ResetProvisioningArtifactId()
	ResetProvisioningArtifactName()
	ResetProvisioningParameters()
	ResetRetainPhysicalResources()
	ResetStackSetProvisioningPreferences()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogProvisionedProduct
type jsiiProxy_ServicecatalogProvisionedProduct struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CloudwatchDashboardNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudwatchDashboardNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) IgnoreErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) IgnoreErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastProvisioningRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProvisioningRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LastSuccessfulProvisioningRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSuccessfulProvisioningRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) LaunchRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) NotificationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) PathNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningArtifactNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) ProvisioningParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RetainPhysicalResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainPhysicalResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) RetainPhysicalResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainPhysicalResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StackSetProvisioningPreferences() ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference {
	var returns ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
	_jsii_.Get(
		j,
		"stackSetProvisioningPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StackSetProvisioningPreferencesInput() *ServicecatalogProvisionedProductStackSetProvisioningPreferences {
	var returns *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	_jsii_.Get(
		j,
		"stackSetProvisioningPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Timeouts() ServicecatalogProvisionedProductTimeoutsOutputReference {
	var returns ServicecatalogProvisionedProductTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) TimeoutsInput() *ServicecatalogProvisionedProductTimeouts {
	var returns *ServicecatalogProvisionedProductTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product} Resource.
func NewServicecatalogProvisionedProduct(scope constructs.Construct, id *string, config *ServicecatalogProvisionedProductConfig) ServicecatalogProvisionedProduct {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProvisionedProduct{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product aws_servicecatalog_provisioned_product} Resource.
func NewServicecatalogProvisionedProduct_Override(s ServicecatalogProvisionedProduct, scope constructs.Construct, id *string, config *ServicecatalogProvisionedProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProduct",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetIgnoreErrors(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreErrors",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetPathId(val *string) {
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetPathName(val *string) {
	_jsii_.Set(
		j,
		"pathName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProductName(val *string) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProvisioningArtifactId(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProvisioningArtifactName(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactName",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetProvisioningParameters(val interface{}) {
	_jsii_.Set(
		j,
		"provisioningParameters",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetRetainPhysicalResources(val interface{}) {
	_jsii_.Set(
		j,
		"retainPhysicalResources",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProduct) SetTagsAll(val *map[string]*string) {
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
func ServicecatalogProvisionedProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogProvisionedProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) PutStackSetProvisioningPreferences(value *ServicecatalogProvisionedProductStackSetProvisioningPreferences) {
	_jsii_.InvokeVoid(
		s,
		"putStackSetProvisioningPreferences",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) PutTimeouts(value *ServicecatalogProvisionedProductTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetIgnoreErrors() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreErrors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetNotificationArns() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationArns",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetPathId() {
	_jsii_.InvokeVoid(
		s,
		"resetPathId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetPathName() {
	_jsii_.InvokeVoid(
		s,
		"resetPathName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProductId() {
	_jsii_.InvokeVoid(
		s,
		"resetProductId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProductName() {
	_jsii_.InvokeVoid(
		s,
		"resetProductName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningArtifactId() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningArtifactName() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningArtifactName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetProvisioningParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetRetainPhysicalResources() {
	_jsii_.InvokeVoid(
		s,
		"resetRetainPhysicalResources",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetStackSetProvisioningPreferences() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetProvisioningPreferences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogProvisionedProductConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#name ServicecatalogProvisionedProduct#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#accept_language ServicecatalogProvisionedProduct#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#ignore_errors ServicecatalogProvisionedProduct#ignore_errors}.
	IgnoreErrors interface{} `json:"ignoreErrors" yaml:"ignoreErrors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#notification_arns ServicecatalogProvisionedProduct#notification_arns}.
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#path_id ServicecatalogProvisionedProduct#path_id}.
	PathId *string `json:"pathId" yaml:"pathId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#path_name ServicecatalogProvisionedProduct#path_name}.
	PathName *string `json:"pathName" yaml:"pathName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#product_id ServicecatalogProvisionedProduct#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#product_name ServicecatalogProvisionedProduct#product_name}.
	ProductName *string `json:"productName" yaml:"productName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#provisioning_artifact_id ServicecatalogProvisionedProduct#provisioning_artifact_id}.
	ProvisioningArtifactId *string `json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#provisioning_artifact_name ServicecatalogProvisionedProduct#provisioning_artifact_name}.
	ProvisioningArtifactName *string `json:"provisioningArtifactName" yaml:"provisioningArtifactName"`
	// provisioning_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#provisioning_parameters ServicecatalogProvisionedProduct#provisioning_parameters}
	ProvisioningParameters interface{} `json:"provisioningParameters" yaml:"provisioningParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#retain_physical_resources ServicecatalogProvisionedProduct#retain_physical_resources}.
	RetainPhysicalResources interface{} `json:"retainPhysicalResources" yaml:"retainPhysicalResources"`
	// stack_set_provisioning_preferences block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#stack_set_provisioning_preferences ServicecatalogProvisionedProduct#stack_set_provisioning_preferences}
	StackSetProvisioningPreferences *ServicecatalogProvisionedProductStackSetProvisioningPreferences `json:"stackSetProvisioningPreferences" yaml:"stackSetProvisioningPreferences"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#tags ServicecatalogProvisionedProduct#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#tags_all ServicecatalogProvisionedProduct#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#timeouts ServicecatalogProvisionedProduct#timeouts}
	Timeouts *ServicecatalogProvisionedProductTimeouts `json:"timeouts" yaml:"timeouts"`
}

type ServicecatalogProvisionedProductProvisioningParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#key ServicecatalogProvisionedProduct#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#use_previous_value ServicecatalogProvisionedProduct#use_previous_value}.
	UsePreviousValue interface{} `json:"usePreviousValue" yaml:"usePreviousValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#value ServicecatalogProvisionedProduct#value}.
	Value *string `json:"value" yaml:"value"`
}

type ServicecatalogProvisionedProductStackSetProvisioningPreferences struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#accounts ServicecatalogProvisionedProduct#accounts}.
	Accounts *[]*string `json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#failure_tolerance_count ServicecatalogProvisionedProduct#failure_tolerance_count}.
	FailureToleranceCount *float64 `json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#failure_tolerance_percentage ServicecatalogProvisionedProduct#failure_tolerance_percentage}.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#max_concurrency_count ServicecatalogProvisionedProduct#max_concurrency_count}.
	MaxConcurrencyCount *float64 `json:"maxConcurrencyCount" yaml:"maxConcurrencyCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#max_concurrency_percentage ServicecatalogProvisionedProduct#max_concurrency_percentage}.
	MaxConcurrencyPercentage *float64 `json:"maxConcurrencyPercentage" yaml:"maxConcurrencyPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#regions ServicecatalogProvisionedProduct#regions}.
	Regions *[]*string `json:"regions" yaml:"regions"`
}

type ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference interface {
	cdktf.ComplexObject
	Accounts() *[]*string
	SetAccounts(val *[]*string)
	AccountsInput() *[]*string
	FailureToleranceCount() *float64
	SetFailureToleranceCount(val *float64)
	FailureToleranceCountInput() *float64
	FailureTolerancePercentage() *float64
	SetFailureTolerancePercentage(val *float64)
	FailureTolerancePercentageInput() *float64
	InternalValue() *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	SetInternalValue(val *ServicecatalogProvisionedProductStackSetProvisioningPreferences)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxConcurrencyCount() *float64
	SetMaxConcurrencyCount(val *float64)
	MaxConcurrencyCountInput() *float64
	MaxConcurrencyPercentage() *float64
	SetMaxConcurrencyPercentage(val *float64)
	MaxConcurrencyPercentageInput() *float64
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
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
	ResetAccounts()
	ResetFailureToleranceCount()
	ResetFailureTolerancePercentage()
	ResetMaxConcurrencyCount()
	ResetMaxConcurrencyPercentage()
	ResetRegions()
}

// The jsii proxy struct for ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
type jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InternalValue() *ServicecatalogProvisionedProductStackSetProvisioningPreferences {
	var returns *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference_Override(s ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetFailureToleranceCount(val *float64) {
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetFailureTolerancePercentage(val *float64) {
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetInternalValue(val *ServicecatalogProvisionedProductStackSetProvisioningPreferences) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetMaxConcurrencyCount(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrencyCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetMaxConcurrencyPercentage(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrencyPercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		s,
		"resetAccounts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetMaxConcurrencyCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrencyCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetMaxConcurrencyPercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrencyPercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetRegions",
		nil, // no parameters
	)
}

type ServicecatalogProvisionedProductTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#create ServicecatalogProvisionedProduct#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#delete ServicecatalogProvisionedProduct#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioned_product#update ServicecatalogProvisionedProduct#update}.
	Update *string `json:"update" yaml:"update"`
}

type ServicecatalogProvisionedProductTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *ServicecatalogProvisionedProductTimeouts
	SetInternalValue(val *ServicecatalogProvisionedProductTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for ServicecatalogProvisionedProductTimeoutsOutputReference
type jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) InternalValue() *ServicecatalogProvisionedProductTimeouts {
	var returns *ServicecatalogProvisionedProductTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewServicecatalogProvisionedProductTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ServicecatalogProvisionedProductTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProductTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewServicecatalogProvisionedProductTimeoutsOutputReference_Override(s ServicecatalogProvisionedProductTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisionedProductTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetInternalValue(val *ServicecatalogProvisionedProductTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact aws_servicecatalog_provisioning_artifact}.
type ServicecatalogProvisioningArtifact interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableTemplateValidation() interface{}
	SetDisableTemplateValidation(val interface{})
	DisableTemplateValidationInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Guidance() *string
	SetGuidance(val *string)
	GuidanceInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TemplatePhysicalId() *string
	SetTemplatePhysicalId(val *string)
	TemplatePhysicalIdInput() *string
	TemplateUrl() *string
	SetTemplateUrl(val *string)
	TemplateUrlInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAcceptLanguage()
	ResetActive()
	ResetDescription()
	ResetDisableTemplateValidation()
	ResetGuidance()
	ResetName()
	ResetOverrideLogicalId()
	ResetTemplatePhysicalId()
	ResetTemplateUrl()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogProvisioningArtifact
type jsiiProxy_ServicecatalogProvisioningArtifact struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) DisableTemplateValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) DisableTemplateValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Guidance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) GuidanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guidanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TemplatePhysicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TemplatePhysicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TemplateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact aws_servicecatalog_provisioning_artifact} Resource.
func NewServicecatalogProvisioningArtifact(scope constructs.Construct, id *string, config *ServicecatalogProvisioningArtifactConfig) ServicecatalogProvisioningArtifact {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogProvisioningArtifact{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisioningArtifact",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact aws_servicecatalog_provisioning_artifact} Resource.
func NewServicecatalogProvisioningArtifact_Override(s ServicecatalogProvisioningArtifact, scope constructs.Construct, id *string, config *ServicecatalogProvisioningArtifactConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisioningArtifact",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetDisableTemplateValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableTemplateValidation",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetGuidance(val *string) {
	_jsii_.Set(
		j,
		"guidance",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetTemplatePhysicalId(val *string) {
	_jsii_.Set(
		j,
		"templatePhysicalId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetTemplateUrl(val *string) {
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisioningArtifact) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogProvisioningArtifact_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisioningArtifact",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogProvisioningArtifact_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogProvisioningArtifact",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetActive() {
	_jsii_.InvokeVoid(
		s,
		"resetActive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetDisableTemplateValidation() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableTemplateValidation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetGuidance() {
	_jsii_.InvokeVoid(
		s,
		"resetGuidance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetTemplatePhysicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplatePhysicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetTemplateUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplateUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisioningArtifact) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogProvisioningArtifact) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogProvisioningArtifactConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#product_id ServicecatalogProvisioningArtifact#product_id}.
	ProductId *string `json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#accept_language ServicecatalogProvisioningArtifact#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#active ServicecatalogProvisioningArtifact#active}.
	Active interface{} `json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#description ServicecatalogProvisioningArtifact#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#disable_template_validation ServicecatalogProvisioningArtifact#disable_template_validation}.
	DisableTemplateValidation interface{} `json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#guidance ServicecatalogProvisioningArtifact#guidance}.
	Guidance *string `json:"guidance" yaml:"guidance"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#name ServicecatalogProvisioningArtifact#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#template_physical_id ServicecatalogProvisioningArtifact#template_physical_id}.
	TemplatePhysicalId *string `json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#template_url ServicecatalogProvisioningArtifact#template_url}.
	TemplateUrl *string `json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_provisioning_artifact#type ServicecatalogProvisioningArtifact#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action aws_servicecatalog_service_action}.
type ServicecatalogServiceAction interface {
	cdktf.TerraformResource
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AcceptLanguageInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Definition() ServicecatalogServiceActionDefinitionOutputReference
	DefinitionInput() *ServicecatalogServiceActionDefinition
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
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
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
	PutDefinition(value *ServicecatalogServiceActionDefinition)
	ResetAcceptLanguage()
	ResetDescription()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogServiceAction
type jsiiProxy_ServicecatalogServiceAction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogServiceAction) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) AcceptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Definition() ServicecatalogServiceActionDefinitionOutputReference {
	var returns ServicecatalogServiceActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) DefinitionInput() *ServicecatalogServiceActionDefinition {
	var returns *ServicecatalogServiceActionDefinition
	_jsii_.Get(
		j,
		"definitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceAction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action aws_servicecatalog_service_action} Resource.
func NewServicecatalogServiceAction(scope constructs.Construct, id *string, config *ServicecatalogServiceActionConfig) ServicecatalogServiceAction {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogServiceAction{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceAction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action aws_servicecatalog_service_action} Resource.
func NewServicecatalogServiceAction_Override(s ServicecatalogServiceAction, scope constructs.Construct, id *string, config *ServicecatalogServiceActionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceAction",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceAction) SetProvider(val cdktf.TerraformProvider) {
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
func ServicecatalogServiceAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogServiceAction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceAction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogServiceAction) PutDefinition(value *ServicecatalogServiceActionDefinition) {
	_jsii_.InvokeVoid(
		s,
		"putDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicecatalogServiceAction) ResetAcceptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogServiceAction) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogServiceAction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogServiceAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogServiceAction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogServiceActionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#definition ServicecatalogServiceAction#definition}
	Definition *ServicecatalogServiceActionDefinition `json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#name ServicecatalogServiceAction#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#accept_language ServicecatalogServiceAction#accept_language}.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#description ServicecatalogServiceAction#description}.
	Description *string `json:"description" yaml:"description"`
}

type ServicecatalogServiceActionDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#name ServicecatalogServiceAction#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#version ServicecatalogServiceAction#version}.
	Version *string `json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#assume_role ServicecatalogServiceAction#assume_role}.
	AssumeRole *string `json:"assumeRole" yaml:"assumeRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#parameters ServicecatalogServiceAction#parameters}.
	Parameters *string `json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_service_action#type ServicecatalogServiceAction#type}.
	Type *string `json:"type" yaml:"type"`
}

type ServicecatalogServiceActionDefinitionOutputReference interface {
	cdktf.ComplexObject
	AssumeRole() *string
	SetAssumeRole(val *string)
	AssumeRoleInput() *string
	InternalValue() *ServicecatalogServiceActionDefinition
	SetInternalValue(val *ServicecatalogServiceActionDefinition)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetAssumeRole()
	ResetParameters()
	ResetType()
}

// The jsii proxy struct for ServicecatalogServiceActionDefinitionOutputReference
type jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) AssumeRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) AssumeRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) InternalValue() *ServicecatalogServiceActionDefinition {
	var returns *ServicecatalogServiceActionDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func NewServicecatalogServiceActionDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ServicecatalogServiceActionDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewServicecatalogServiceActionDefinitionOutputReference_Override(s ServicecatalogServiceActionDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogServiceActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetAssumeRole(val *string) {
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetInternalValue(val *ServicecatalogServiceActionDefinition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		s,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogServiceActionDefinitionOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option aws_servicecatalog_tag_option}.
type ServicecatalogTagOption interface {
	cdktf.TerraformResource
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Owner() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetActive()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicecatalogTagOption
type jsiiProxy_ServicecatalogTagOption struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogTagOption) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOption) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option aws_servicecatalog_tag_option} Resource.
func NewServicecatalogTagOption(scope constructs.Construct, id *string, config *ServicecatalogTagOptionConfig) ServicecatalogTagOption {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogTagOption{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOption",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option aws_servicecatalog_tag_option} Resource.
func NewServicecatalogTagOption_Override(s ServicecatalogTagOption, scope constructs.Construct, id *string, config *ServicecatalogTagOptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOption",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOption) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogTagOption_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOption",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogTagOption_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOption",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicecatalogTagOption) ResetActive() {
	_jsii_.InvokeVoid(
		s,
		"resetActive",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogTagOption) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogTagOption) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOption) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogTagOptionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option#key ServicecatalogTagOption#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option#value ServicecatalogTagOption#value}.
	Value *string `json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option#active ServicecatalogTagOption#active}.
	Active interface{} `json:"active" yaml:"active"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option_resource_association aws_servicecatalog_tag_option_resource_association}.
type ServicecatalogTagOptionResourceAssociation interface {
	cdktf.TerraformResource
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
	ResourceArn() *string
	ResourceCreatedTime() *string
	ResourceDescription() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceName() *string
	TagOptionId() *string
	SetTagOptionId(val *string)
	TagOptionIdInput() *string
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

// The jsii proxy struct for ServicecatalogTagOptionResourceAssociation
type jsiiProxy_ServicecatalogTagOptionResourceAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) TagOptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagOptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) TagOptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagOptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option_resource_association aws_servicecatalog_tag_option_resource_association} Resource.
func NewServicecatalogTagOptionResourceAssociation(scope constructs.Construct, id *string, config *ServicecatalogTagOptionResourceAssociationConfig) ServicecatalogTagOptionResourceAssociation {
	_init_.Initialize()

	j := jsiiProxy_ServicecatalogTagOptionResourceAssociation{}

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOptionResourceAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option_resource_association aws_servicecatalog_tag_option_resource_association} Resource.
func NewServicecatalogTagOptionResourceAssociation_Override(s ServicecatalogTagOptionResourceAssociation, scope constructs.Construct, id *string, config *ServicecatalogTagOptionResourceAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOptionResourceAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SetTagOptionId(val *string) {
	_jsii_.Set(
		j,
		"tagOptionId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServicecatalogTagOptionResourceAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOptionResourceAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicecatalogTagOptionResourceAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.servicecatalog.ServicecatalogTagOptionResourceAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_ServicecatalogTagOptionResourceAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Service Catalog.
type ServicecatalogTagOptionResourceAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option_resource_association#resource_id ServicecatalogTagOptionResourceAssociation#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_tag_option_resource_association#tag_option_id ServicecatalogTagOptionResourceAssociation#tag_option_id}.
	TagOptionId *string `json:"tagOptionId" yaml:"tagOptionId"`
}
