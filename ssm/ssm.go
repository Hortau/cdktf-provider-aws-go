package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/ssm/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/ssm_document aws_ssm_document}.
type DataAwsSsmDocument interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DocumentFormat() *string
	SetDocumentFormat(val *string)
	DocumentFormatInput() *string
	DocumentType() *string
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
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
	ResetDocumentFormat()
	ResetDocumentVersion()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSsmDocument
type jsiiProxy_DataAwsSsmDocument struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSsmDocument) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DocumentFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DocumentFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmDocument) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_document aws_ssm_document} Data Source.
func NewDataAwsSsmDocument(scope constructs.Construct, id *string, config *DataAwsSsmDocumentConfig) DataAwsSsmDocument {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSsmDocument{}

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmDocument",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_document aws_ssm_document} Data Source.
func NewDataAwsSsmDocument_Override(d DataAwsSsmDocument, scope constructs.Construct, id *string, config *DataAwsSsmDocumentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmDocument",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetDocumentFormat(val *string) {
	_jsii_.Set(
		j,
		"documentFormat",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetDocumentVersion(val *string) {
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmDocument) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsSsmDocument_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.DataAwsSsmDocument",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSsmDocument_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.DataAwsSsmDocument",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmDocument) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmDocument) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSsmDocument) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSsmDocument) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmDocument) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSsmDocument) ResetDocumentFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDocumentFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmDocument) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSsmDocument) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmDocument) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmDocument) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSsmDocument) ToString() *string {
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
func (d *jsiiProxy_DataAwsSsmDocument) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type DataAwsSsmDocumentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_document#name DataAwsSsmDocument#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_document#document_format DataAwsSsmDocument#document_format}.
	DocumentFormat *string `json:"documentFormat" yaml:"documentFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_document#document_version DataAwsSsmDocument#document_version}.
	DocumentVersion *string `json:"documentVersion" yaml:"documentVersion"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameter aws_ssm_parameter}.
type DataAwsSsmParameter interface {
	cdktf.TerraformDataSource
	Arn() *string
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
	Type() *string
	Value() *string
	Version() *float64
	WithDecryption() interface{}
	SetWithDecryption(val interface{})
	WithDecryptionInput() interface{}
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
	ResetWithDecryption()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSsmParameter
type jsiiProxy_DataAwsSsmParameter struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSsmParameter) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) WithDecryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDecryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParameter) WithDecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDecryptionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameter aws_ssm_parameter} Data Source.
func NewDataAwsSsmParameter(scope constructs.Construct, id *string, config *DataAwsSsmParameterConfig) DataAwsSsmParameter {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSsmParameter{}

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmParameter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameter aws_ssm_parameter} Data Source.
func NewDataAwsSsmParameter_Override(d DataAwsSsmParameter, scope constructs.Construct, id *string, config *DataAwsSsmParameterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmParameter",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParameter) SetWithDecryption(val interface{}) {
	_jsii_.Set(
		j,
		"withDecryption",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsSsmParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.DataAwsSsmParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSsmParameter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.DataAwsSsmParameter",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmParameter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmParameter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSsmParameter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSsmParameter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmParameter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSsmParameter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmParameter) ResetWithDecryption() {
	_jsii_.InvokeVoid(
		d,
		"resetWithDecryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmParameter) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmParameter) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSsmParameter) ToString() *string {
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
func (d *jsiiProxy_DataAwsSsmParameter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type DataAwsSsmParameterConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameter#name DataAwsSsmParameter#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameter#with_decryption DataAwsSsmParameter#with_decryption}.
	WithDecryption interface{} `json:"withDecryption" yaml:"withDecryption"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path aws_ssm_parameters_by_path}.
type DataAwsSsmParametersByPath interface {
	cdktf.TerraformDataSource
	Arns() *[]*string
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
	Names() *[]*string
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Recursive() interface{}
	SetRecursive(val interface{})
	RecursiveInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Types() *[]*string
	Values() *[]*string
	WithDecryption() interface{}
	SetWithDecryption(val interface{})
	WithDecryptionInput() interface{}
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
	ResetRecursive()
	ResetWithDecryption()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSsmParametersByPath
type jsiiProxy_DataAwsSsmParametersByPath struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Arns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Names() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"names",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Recursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) RecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) WithDecryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDecryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) WithDecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDecryptionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path aws_ssm_parameters_by_path} Data Source.
func NewDataAwsSsmParametersByPath(scope constructs.Construct, id *string, config *DataAwsSsmParametersByPathConfig) DataAwsSsmParametersByPath {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSsmParametersByPath{}

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmParametersByPath",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path aws_ssm_parameters_by_path} Data Source.
func NewDataAwsSsmParametersByPath_Override(d DataAwsSsmParametersByPath, scope constructs.Construct, id *string, config *DataAwsSsmParametersByPathConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmParametersByPath",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetRecursive(val interface{}) {
	_jsii_.Set(
		j,
		"recursive",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmParametersByPath) SetWithDecryption(val interface{}) {
	_jsii_.Set(
		j,
		"withDecryption",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsSsmParametersByPath_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.DataAwsSsmParametersByPath",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSsmParametersByPath_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.DataAwsSsmParametersByPath",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmParametersByPath) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSsmParametersByPath) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmParametersByPath) ResetRecursive() {
	_jsii_.InvokeVoid(
		d,
		"resetRecursive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmParametersByPath) ResetWithDecryption() {
	_jsii_.InvokeVoid(
		d,
		"resetWithDecryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmParametersByPath) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) ToString() *string {
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
func (d *jsiiProxy_DataAwsSsmParametersByPath) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type DataAwsSsmParametersByPathConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path#path DataAwsSsmParametersByPath#path}.
	Path *string `json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path#recursive DataAwsSsmParametersByPath#recursive}.
	Recursive interface{} `json:"recursive" yaml:"recursive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_parameters_by_path#with_decryption DataAwsSsmParametersByPath#with_decryption}.
	WithDecryption interface{} `json:"withDecryption" yaml:"withDecryption"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline aws_ssm_patch_baseline}.
type DataAwsSsmPatchBaseline interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DefaultBaseline() interface{}
	SetDefaultBaseline(val interface{})
	DefaultBaselineInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
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
	ResetDefaultBaseline()
	ResetNamePrefix()
	ResetOperatingSystem()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSsmPatchBaseline
type jsiiProxy_DataAwsSsmPatchBaseline struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) DefaultBaseline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultBaseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) DefaultBaselineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultBaselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline aws_ssm_patch_baseline} Data Source.
func NewDataAwsSsmPatchBaseline(scope constructs.Construct, id *string, config *DataAwsSsmPatchBaselineConfig) DataAwsSsmPatchBaseline {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSsmPatchBaseline{}

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmPatchBaseline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline aws_ssm_patch_baseline} Data Source.
func NewDataAwsSsmPatchBaseline_Override(d DataAwsSsmPatchBaseline, scope constructs.Construct, id *string, config *DataAwsSsmPatchBaselineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.DataAwsSsmPatchBaseline",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetDefaultBaseline(val interface{}) {
	_jsii_.Set(
		j,
		"defaultBaseline",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetOperatingSystem(val *string) {
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataAwsSsmPatchBaseline) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsSsmPatchBaseline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.DataAwsSsmPatchBaseline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSsmPatchBaseline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.DataAwsSsmPatchBaseline",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmPatchBaseline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSsmPatchBaseline) ResetDefaultBaseline() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultBaseline",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmPatchBaseline) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmPatchBaseline) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		d,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSsmPatchBaseline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSsmPatchBaseline) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) ToString() *string {
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
func (d *jsiiProxy_DataAwsSsmPatchBaseline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type DataAwsSsmPatchBaselineConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline#owner DataAwsSsmPatchBaseline#owner}.
	Owner *string `json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline#default_baseline DataAwsSsmPatchBaseline#default_baseline}.
	DefaultBaseline interface{} `json:"defaultBaseline" yaml:"defaultBaseline"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline#name_prefix DataAwsSsmPatchBaseline#name_prefix}.
	NamePrefix *string `json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ssm_patch_baseline#operating_system DataAwsSsmPatchBaseline#operating_system}.
	OperatingSystem *string `json:"operatingSystem" yaml:"operatingSystem"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation aws_ssm_activation}.
type SsmActivation interface {
	cdktf.TerraformResource
	ActivationCode() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	Expired() cdktf.IResolvable
	Fqn() *string
	FriendlyUniqueId() *string
	IamRole() *string
	SetIamRole(val *string)
	IamRoleInput() *string
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
	RegistrationCount() *float64
	RegistrationLimit() *float64
	SetRegistrationLimit(val *float64)
	RegistrationLimitInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
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
	ResetDescription()
	ResetExpirationDate()
	ResetName()
	ResetOverrideLogicalId()
	ResetRegistrationLimit()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmActivation
type jsiiProxy_SsmActivation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmActivation) ActivationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Expired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"expired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) IamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) RegistrationCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"registrationCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) RegistrationLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"registrationLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) RegistrationLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"registrationLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmActivation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation aws_ssm_activation} Resource.
func NewSsmActivation(scope constructs.Construct, id *string, config *SsmActivationConfig) SsmActivation {
	_init_.Initialize()

	j := jsiiProxy_SsmActivation{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmActivation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation aws_ssm_activation} Resource.
func NewSsmActivation_Override(s SsmActivation, scope constructs.Construct, id *string, config *SsmActivationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmActivation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmActivation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetExpirationDate(val *string) {
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetIamRole(val *string) {
	_jsii_.Set(
		j,
		"iamRole",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetRegistrationLimit(val *float64) {
	_jsii_.Set(
		j,
		"registrationLimit",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmActivation) SetTagsAll(val *map[string]*string) {
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
func SsmActivation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmActivation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmActivation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmActivation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmActivation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmActivation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmActivation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmActivation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmActivation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmActivation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmActivation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmActivation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmActivation) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmActivation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmActivation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmActivation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmActivation) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmActivation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) ResetRegistrationLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetRegistrationLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmActivation) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmActivation) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmActivation) ToString() *string {
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
func (s *jsiiProxy_SsmActivation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmActivationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#iam_role SsmActivation#iam_role}.
	IamRole *string `json:"iamRole" yaml:"iamRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#description SsmActivation#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#expiration_date SsmActivation#expiration_date}.
	ExpirationDate *string `json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#name SsmActivation#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#registration_limit SsmActivation#registration_limit}.
	RegistrationLimit *float64 `json:"registrationLimit" yaml:"registrationLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#tags SsmActivation#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_activation#tags_all SsmActivation#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_association aws_ssm_association}.
type SsmAssociation interface {
	cdktf.TerraformResource
	ApplyOnlyAtCronInterval() interface{}
	SetApplyOnlyAtCronInterval(val interface{})
	ApplyOnlyAtCronIntervalInput() interface{}
	AssociationId() *string
	AssociationName() *string
	SetAssociationName(val *string)
	AssociationNameInput() *string
	AutomationTargetParameterName() *string
	SetAutomationTargetParameterName(val *string)
	AutomationTargetParameterNameInput() *string
	CdktfStack() cdktf.TerraformStack
	ComplianceSeverity() *string
	SetComplianceSeverity(val *string)
	ComplianceSeverityInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	MaxConcurrencyInput() *string
	MaxErrors() *string
	SetMaxErrors(val *string)
	MaxErrorsInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OutputLocation() SsmAssociationOutputLocationOutputReference
	OutputLocationInput() *SsmAssociationOutputLocation
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	ScheduleExpressionInput() *string
	Targets() interface{}
	SetTargets(val interface{})
	TargetsInput() interface{}
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
	PutOutputLocation(value *SsmAssociationOutputLocation)
	ResetApplyOnlyAtCronInterval()
	ResetAssociationName()
	ResetAutomationTargetParameterName()
	ResetComplianceSeverity()
	ResetDocumentVersion()
	ResetInstanceId()
	ResetMaxConcurrency()
	ResetMaxErrors()
	ResetOutputLocation()
	ResetOverrideLogicalId()
	ResetParameters()
	ResetScheduleExpression()
	ResetTargets()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmAssociation
type jsiiProxy_SsmAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmAssociation) ApplyOnlyAtCronInterval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyOnlyAtCronInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ApplyOnlyAtCronIntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyOnlyAtCronIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AutomationTargetParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationTargetParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AutomationTargetParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationTargetParameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ComplianceSeverity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ComplianceSeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxConcurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) OutputLocation() SsmAssociationOutputLocationOutputReference {
	var returns SsmAssociationOutputLocationOutputReference
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) OutputLocationInput() *SsmAssociationOutputLocation {
	var returns *SsmAssociationOutputLocation
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_association aws_ssm_association} Resource.
func NewSsmAssociation(scope constructs.Construct, id *string, config *SsmAssociationConfig) SsmAssociation {
	_init_.Initialize()

	j := jsiiProxy_SsmAssociation{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_association aws_ssm_association} Resource.
func NewSsmAssociation_Override(s SsmAssociation, scope constructs.Construct, id *string, config *SsmAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmAssociation) SetApplyOnlyAtCronInterval(val interface{}) {
	_jsii_.Set(
		j,
		"applyOnlyAtCronInterval",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetAssociationName(val *string) {
	_jsii_.Set(
		j,
		"associationName",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetAutomationTargetParameterName(val *string) {
	_jsii_.Set(
		j,
		"automationTargetParameterName",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetComplianceSeverity(val *string) {
	_jsii_.Set(
		j,
		"complianceSeverity",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetDocumentVersion(val *string) {
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetMaxConcurrency(val *string) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetMaxErrors(val *string) {
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetScheduleExpression(val *string) {
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SsmAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmAssociation) PutOutputLocation(value *SsmAssociationOutputLocation) {
	_jsii_.InvokeVoid(
		s,
		"putOutputLocation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmAssociation) ResetApplyOnlyAtCronInterval() {
	_jsii_.InvokeVoid(
		s,
		"resetApplyOnlyAtCronInterval",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetAssociationName() {
	_jsii_.InvokeVoid(
		s,
		"resetAssociationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetAutomationTargetParameterName() {
	_jsii_.InvokeVoid(
		s,
		"resetAutomationTargetParameterName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetComplianceSeverity() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceSeverity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetInstanceId() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetMaxErrors() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxErrors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetOutputLocation() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputLocation",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetScheduleExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmAssociation) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmAssociation) ToString() *string {
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
func (s *jsiiProxy_SsmAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#name SsmAssociation#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#apply_only_at_cron_interval SsmAssociation#apply_only_at_cron_interval}.
	ApplyOnlyAtCronInterval interface{} `json:"applyOnlyAtCronInterval" yaml:"applyOnlyAtCronInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#association_name SsmAssociation#association_name}.
	AssociationName *string `json:"associationName" yaml:"associationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#automation_target_parameter_name SsmAssociation#automation_target_parameter_name}.
	AutomationTargetParameterName *string `json:"automationTargetParameterName" yaml:"automationTargetParameterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#compliance_severity SsmAssociation#compliance_severity}.
	ComplianceSeverity *string `json:"complianceSeverity" yaml:"complianceSeverity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#document_version SsmAssociation#document_version}.
	DocumentVersion *string `json:"documentVersion" yaml:"documentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#instance_id SsmAssociation#instance_id}.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#max_concurrency SsmAssociation#max_concurrency}.
	MaxConcurrency *string `json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#max_errors SsmAssociation#max_errors}.
	MaxErrors *string `json:"maxErrors" yaml:"maxErrors"`
	// output_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#output_location SsmAssociation#output_location}
	OutputLocation *SsmAssociationOutputLocation `json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#parameters SsmAssociation#parameters}.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#schedule_expression SsmAssociation#schedule_expression}.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#targets SsmAssociation#targets}
	Targets interface{} `json:"targets" yaml:"targets"`
}

type SsmAssociationOutputLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#s3_bucket_name SsmAssociation#s3_bucket_name}.
	S3BucketName *string `json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#s3_key_prefix SsmAssociation#s3_key_prefix}.
	S3KeyPrefix *string `json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#s3_region SsmAssociation#s3_region}.
	S3Region *string `json:"s3Region" yaml:"s3Region"`
}

type SsmAssociationOutputLocationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SsmAssociationOutputLocation
	SetInternalValue(val *SsmAssociationOutputLocation)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KeyPrefixInput() *string
	S3Region() *string
	SetS3Region(val *string)
	S3RegionInput() *string
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
	ResetS3KeyPrefix()
	ResetS3Region()
}

// The jsii proxy struct for SsmAssociationOutputLocationOutputReference
type jsiiProxy_SsmAssociationOutputLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) InternalValue() *SsmAssociationOutputLocation {
	var returns *SsmAssociationOutputLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) S3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmAssociationOutputLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmAssociationOutputLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmAssociationOutputLocationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmAssociationOutputLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmAssociationOutputLocationOutputReference_Override(s SsmAssociationOutputLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmAssociationOutputLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetInternalValue(val *SsmAssociationOutputLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetS3Region(val *string) {
	_jsii_.Set(
		j,
		"s3Region",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmAssociationOutputLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociationOutputLocationOutputReference) ResetS3Region() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Region",
		nil, // no parameters
	)
}

type SsmAssociationTargets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#key SsmAssociation#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_association#values SsmAssociation#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_document aws_ssm_document}.
type SsmDocument interface {
	cdktf.TerraformResource
	Arn() *string
	AttachmentsSource() interface{}
	SetAttachmentsSource(val interface{})
	AttachmentsSourceInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	Count() *float64
	SetCount(val *float64)
	CreatedDate() *string
	DefaultVersion() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	DocumentFormat() *string
	SetDocumentFormat(val *string)
	DocumentFormatInput() *string
	DocumentType() *string
	SetDocumentType(val *string)
	DocumentTypeInput() *string
	DocumentVersion() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Hash() *string
	HashType() *string
	Id() *string
	LatestVersion() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Owner() *string
	Permissions() *map[string]*string
	SetPermissions(val *map[string]*string)
	PermissionsInput() *map[string]*string
	PlatformTypes() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SchemaVersion() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VersionName() *string
	SetVersionName(val *string)
	VersionNameInput() *string
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
	Parameter(index *string) SsmDocumentParameter
	ResetAttachmentsSource()
	ResetDocumentFormat()
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetTags()
	ResetTagsAll()
	ResetTargetType()
	ResetVersionName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmDocument
type jsiiProxy_SsmDocument struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmDocument) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) AttachmentsSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) AttachmentsSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DefaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DocumentFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DocumentFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Hash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) HashType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) LatestVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Permissions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) PermissionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) PlatformTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) VersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocument) VersionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_document aws_ssm_document} Resource.
func NewSsmDocument(scope constructs.Construct, id *string, config *SsmDocumentConfig) SsmDocument {
	_init_.Initialize()

	j := jsiiProxy_SsmDocument{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmDocument",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_document aws_ssm_document} Resource.
func NewSsmDocument_Override(s SsmDocument, scope constructs.Construct, id *string, config *SsmDocumentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmDocument",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmDocument) SetAttachmentsSource(val interface{}) {
	_jsii_.Set(
		j,
		"attachmentsSource",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetDocumentFormat(val *string) {
	_jsii_.Set(
		j,
		"documentFormat",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetDocumentType(val *string) {
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetPermissions(val *map[string]*string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_SsmDocument) SetVersionName(val *string) {
	_jsii_.Set(
		j,
		"versionName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SsmDocument_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmDocument",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmDocument_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmDocument",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmDocument) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmDocument) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmDocument) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmDocument) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmDocument) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmDocument) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmDocument) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmDocument) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmDocument) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmDocument) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmDocument) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmDocument) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmDocument) Parameter(index *string) SsmDocumentParameter {
	var returns SsmDocumentParameter

	_jsii_.Invoke(
		s,
		"parameter",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmDocument) ResetAttachmentsSource() {
	_jsii_.InvokeVoid(
		s,
		"resetAttachmentsSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetDocumentFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentFormat",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmDocument) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetPermissions() {
	_jsii_.InvokeVoid(
		s,
		"resetPermissions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetTargetType() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) ResetVersionName() {
	_jsii_.InvokeVoid(
		s,
		"resetVersionName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmDocument) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmDocument) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmDocument) ToString() *string {
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
func (s *jsiiProxy_SsmDocument) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SsmDocumentAttachmentsSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#key SsmDocument#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#values SsmDocument#values}.
	Values *[]*string `json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#name SsmDocument#name}.
	Name *string `json:"name" yaml:"name"`
}

// AWS Systems Manager.
type SsmDocumentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#content SsmDocument#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#document_type SsmDocument#document_type}.
	DocumentType *string `json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#name SsmDocument#name}.
	Name *string `json:"name" yaml:"name"`
	// attachments_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#attachments_source SsmDocument#attachments_source}
	AttachmentsSource interface{} `json:"attachmentsSource" yaml:"attachmentsSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#document_format SsmDocument#document_format}.
	DocumentFormat *string `json:"documentFormat" yaml:"documentFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#permissions SsmDocument#permissions}.
	Permissions *map[string]*string `json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#tags SsmDocument#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#tags_all SsmDocument#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#target_type SsmDocument#target_type}.
	TargetType *string `json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_document#version_name SsmDocument#version_name}.
	VersionName *string `json:"versionName" yaml:"versionName"`
}

type SsmDocumentParameter interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DefaultValue() *string
	Description() *string
	Name() *string
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

// The jsii proxy struct for SsmDocumentParameter
type jsiiProxy_SsmDocumentParameter struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SsmDocumentParameter) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmDocumentParameter) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSsmDocumentParameter(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SsmDocumentParameter {
	_init_.Initialize()

	j := jsiiProxy_SsmDocumentParameter{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmDocumentParameter",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSsmDocumentParameter_Override(s SsmDocumentParameter, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmDocumentParameter",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SsmDocumentParameter) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SsmDocumentParameter) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmDocumentParameter) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SsmDocumentParameter) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmDocumentParameter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmDocumentParameter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmDocumentParameter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmDocumentParameter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmDocumentParameter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmDocumentParameter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmDocumentParameter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmDocumentParameter) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmDocumentParameter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmDocumentParameter) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window aws_ssm_maintenance_window}.
type SsmMaintenanceWindow interface {
	cdktf.TerraformResource
	AllowUnassociatedTargets() interface{}
	SetAllowUnassociatedTargets(val interface{})
	AllowUnassociatedTargetsInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Cutoff() *float64
	SetCutoff(val *float64)
	CutoffInput() *float64
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Duration() *float64
	SetDuration(val *float64)
	DurationInput() *float64
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EndDate() *string
	SetEndDate(val *string)
	EndDateInput() *string
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
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	ScheduleOffset() *float64
	SetScheduleOffset(val *float64)
	ScheduleOffsetInput() *float64
	ScheduleTimezone() *string
	SetScheduleTimezone(val *string)
	ScheduleTimezoneInput() *string
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
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
	ResetAllowUnassociatedTargets()
	ResetDescription()
	ResetEnabled()
	ResetEndDate()
	ResetOverrideLogicalId()
	ResetScheduleOffset()
	ResetScheduleTimezone()
	ResetStartDate()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmMaintenanceWindow
type jsiiProxy_SsmMaintenanceWindow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmMaintenanceWindow) AllowUnassociatedTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnassociatedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) AllowUnassociatedTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnassociatedTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Cutoff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cutoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) CutoffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cutoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) DurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) EndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ScheduleOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ScheduleOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ScheduleTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) ScheduleTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window aws_ssm_maintenance_window} Resource.
func NewSsmMaintenanceWindow(scope constructs.Construct, id *string, config *SsmMaintenanceWindowConfig) SsmMaintenanceWindow {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindow{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window aws_ssm_maintenance_window} Resource.
func NewSsmMaintenanceWindow_Override(s SsmMaintenanceWindow, scope constructs.Construct, id *string, config *SsmMaintenanceWindowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindow",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetAllowUnassociatedTargets(val interface{}) {
	_jsii_.Set(
		j,
		"allowUnassociatedTargets",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetCutoff(val *float64) {
	_jsii_.Set(
		j,
		"cutoff",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetDuration(val *float64) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetEndDate(val *string) {
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetScheduleOffset(val *float64) {
	_jsii_.Set(
		j,
		"scheduleOffset",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetScheduleTimezone(val *string) {
	_jsii_.Set(
		j,
		"scheduleTimezone",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetStartDate(val *string) {
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindow) SetTagsAll(val *map[string]*string) {
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
func SsmMaintenanceWindow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmMaintenanceWindow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmMaintenanceWindow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmMaintenanceWindow",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetAllowUnassociatedTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowUnassociatedTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetEndDate() {
	_jsii_.InvokeVoid(
		s,
		"resetEndDate",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetScheduleOffset() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleOffset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetScheduleTimezone() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleTimezone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetStartDate() {
	_jsii_.InvokeVoid(
		s,
		"resetStartDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindow) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindow) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindow) ToString() *string {
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
func (s *jsiiProxy_SsmMaintenanceWindow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmMaintenanceWindowConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#cutoff SsmMaintenanceWindow#cutoff}.
	Cutoff *float64 `json:"cutoff" yaml:"cutoff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#duration SsmMaintenanceWindow#duration}.
	Duration *float64 `json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#name SsmMaintenanceWindow#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#schedule SsmMaintenanceWindow#schedule}.
	Schedule *string `json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#allow_unassociated_targets SsmMaintenanceWindow#allow_unassociated_targets}.
	AllowUnassociatedTargets interface{} `json:"allowUnassociatedTargets" yaml:"allowUnassociatedTargets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#description SsmMaintenanceWindow#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#enabled SsmMaintenanceWindow#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#end_date SsmMaintenanceWindow#end_date}.
	EndDate *string `json:"endDate" yaml:"endDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#schedule_offset SsmMaintenanceWindow#schedule_offset}.
	ScheduleOffset *float64 `json:"scheduleOffset" yaml:"scheduleOffset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#schedule_timezone SsmMaintenanceWindow#schedule_timezone}.
	ScheduleTimezone *string `json:"scheduleTimezone" yaml:"scheduleTimezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#start_date SsmMaintenanceWindow#start_date}.
	StartDate *string `json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#tags SsmMaintenanceWindow#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window#tags_all SsmMaintenanceWindow#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target aws_ssm_maintenance_window_target}.
type SsmMaintenanceWindowTarget interface {
	cdktf.TerraformResource
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OwnerInformation() *string
	SetOwnerInformation(val *string)
	OwnerInformationInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	Targets() interface{}
	SetTargets(val interface{})
	TargetsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WindowId() *string
	SetWindowId(val *string)
	WindowIdInput() *string
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
	ResetDescription()
	ResetName()
	ResetOverrideLogicalId()
	ResetOwnerInformation()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmMaintenanceWindowTarget
type jsiiProxy_SsmMaintenanceWindowTarget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) OwnerInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) OwnerInformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) WindowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target aws_ssm_maintenance_window_target} Resource.
func NewSsmMaintenanceWindowTarget(scope constructs.Construct, id *string, config *SsmMaintenanceWindowTargetConfig) SsmMaintenanceWindowTarget {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTarget{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target aws_ssm_maintenance_window_target} Resource.
func NewSsmMaintenanceWindowTarget_Override(s SsmMaintenanceWindowTarget, scope constructs.Construct, id *string, config *SsmMaintenanceWindowTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTarget",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetOwnerInformation(val *string) {
	_jsii_.Set(
		j,
		"ownerInformation",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTarget) SetWindowId(val *string) {
	_jsii_.Set(
		j,
		"windowId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SsmMaintenanceWindowTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmMaintenanceWindowTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTarget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTarget) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTarget) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTarget) ResetOwnerInformation() {
	_jsii_.InvokeVoid(
		s,
		"resetOwnerInformation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTarget) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) ToString() *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmMaintenanceWindowTargetConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#resource_type SsmMaintenanceWindowTarget#resource_type}.
	ResourceType *string `json:"resourceType" yaml:"resourceType"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#targets SsmMaintenanceWindowTarget#targets}
	Targets interface{} `json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#window_id SsmMaintenanceWindowTarget#window_id}.
	WindowId *string `json:"windowId" yaml:"windowId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#description SsmMaintenanceWindowTarget#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#name SsmMaintenanceWindowTarget#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#owner_information SsmMaintenanceWindowTarget#owner_information}.
	OwnerInformation *string `json:"ownerInformation" yaml:"ownerInformation"`
}

type SsmMaintenanceWindowTargetTargets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#key SsmMaintenanceWindowTarget#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_target#values SsmMaintenanceWindowTarget#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task aws_ssm_maintenance_window_task}.
type SsmMaintenanceWindowTask interface {
	cdktf.TerraformResource
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
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	MaxConcurrencyInput() *string
	MaxErrors() *string
	SetMaxErrors(val *string)
	MaxErrorsInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	Targets() interface{}
	SetTargets(val interface{})
	TargetsInput() interface{}
	TaskArn() *string
	SetTaskArn(val *string)
	TaskArnInput() *string
	TaskInvocationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference
	TaskInvocationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParameters
	TaskType() *string
	SetTaskType(val *string)
	TaskTypeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WindowId() *string
	SetWindowId(val *string)
	WindowIdInput() *string
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
	PutTaskInvocationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParameters)
	ResetDescription()
	ResetName()
	ResetOverrideLogicalId()
	ResetPriority()
	ResetServiceRoleArn()
	ResetTargets()
	ResetTaskInvocationParameters()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmMaintenanceWindowTask
type jsiiProxy_SsmMaintenanceWindowTask struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) MaxConcurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) MaxErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskInvocationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference
	_jsii_.Get(
		j,
		"taskInvocationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskInvocationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParameters
	_jsii_.Get(
		j,
		"taskInvocationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) WindowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task aws_ssm_maintenance_window_task} Resource.
func NewSsmMaintenanceWindowTask(scope constructs.Construct, id *string, config *SsmMaintenanceWindowTaskConfig) SsmMaintenanceWindowTask {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTask{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task aws_ssm_maintenance_window_task} Resource.
func NewSsmMaintenanceWindowTask_Override(s SsmMaintenanceWindowTask, scope constructs.Construct, id *string, config *SsmMaintenanceWindowTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTask",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetMaxConcurrency(val *string) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetMaxErrors(val *string) {
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetTaskArn(val *string) {
	_jsii_.Set(
		j,
		"taskArn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetTaskType(val *string) {
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTask) SetWindowId(val *string) {
	_jsii_.Set(
		j,
		"windowId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SsmMaintenanceWindowTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmMaintenanceWindowTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) PutTaskInvocationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParameters) {
	_jsii_.InvokeVoid(
		s,
		"putTaskInvocationParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) ResetTaskInvocationParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetTaskInvocationParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTask) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) ToString() *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmMaintenanceWindowTaskConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#max_concurrency SsmMaintenanceWindowTask#max_concurrency}.
	MaxConcurrency *string `json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#max_errors SsmMaintenanceWindowTask#max_errors}.
	MaxErrors *string `json:"maxErrors" yaml:"maxErrors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#task_arn SsmMaintenanceWindowTask#task_arn}.
	TaskArn *string `json:"taskArn" yaml:"taskArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#task_type SsmMaintenanceWindowTask#task_type}.
	TaskType *string `json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#window_id SsmMaintenanceWindowTask#window_id}.
	WindowId *string `json:"windowId" yaml:"windowId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#description SsmMaintenanceWindowTask#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#priority SsmMaintenanceWindowTask#priority}.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#service_role_arn SsmMaintenanceWindowTask#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#targets SsmMaintenanceWindowTask#targets}
	Targets interface{} `json:"targets" yaml:"targets"`
	// task_invocation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#task_invocation_parameters SsmMaintenanceWindowTask#task_invocation_parameters}
	TaskInvocationParameters *SsmMaintenanceWindowTaskTaskInvocationParameters `json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
}

type SsmMaintenanceWindowTaskTargets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#key SsmMaintenanceWindowTask#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#values SsmMaintenanceWindowTask#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

type SsmMaintenanceWindowTaskTaskInvocationParameters struct {
	// automation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#automation_parameters SsmMaintenanceWindowTask#automation_parameters}
	AutomationParameters *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters `json:"automationParameters" yaml:"automationParameters"`
	// lambda_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#lambda_parameters SsmMaintenanceWindowTask#lambda_parameters}
	LambdaParameters *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters `json:"lambdaParameters" yaml:"lambdaParameters"`
	// run_command_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#run_command_parameters SsmMaintenanceWindowTask#run_command_parameters}
	RunCommandParameters *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters `json:"runCommandParameters" yaml:"runCommandParameters"`
	// step_functions_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#step_functions_parameters SsmMaintenanceWindowTask#step_functions_parameters}
	StepFunctionsParameters *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters `json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#document_version SsmMaintenanceWindowTask#document_version}.
	DocumentVersion *string `json:"documentVersion" yaml:"documentVersion"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#parameter SsmMaintenanceWindowTask#parameter}
	Parameter interface{} `json:"parameter" yaml:"parameter"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference interface {
	cdktf.ComplexObject
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Parameter() interface{}
	SetParameter(val interface{})
	ParameterInput() interface{}
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
	ResetDocumentVersion()
	ResetParameter()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) Parameter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetDocumentVersion(val *string) {
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetParameter(val interface{}) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		s,
		"resetParameter",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#values SsmMaintenanceWindowTask#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#client_context SsmMaintenanceWindowTask#client_context}.
	ClientContext *string `json:"clientContext" yaml:"clientContext"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#payload SsmMaintenanceWindowTask#payload}.
	Payload *string `json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#qualifier SsmMaintenanceWindowTask#qualifier}.
	Qualifier *string `json:"qualifier" yaml:"qualifier"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference interface {
	cdktf.ComplexObject
	ClientContext() *string
	SetClientContext(val *string)
	ClientContextInput() *string
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
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
	ResetClientContext()
	ResetPayload()
	ResetQualifier()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) ClientContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) ClientContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetClientContext(val *string) {
	_jsii_.Set(
		j,
		"clientContext",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetPayload(val *string) {
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) ResetClientContext() {
	_jsii_.InvokeVoid(
		s,
		"resetClientContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		s,
		"resetPayload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference) ResetQualifier() {
	_jsii_.InvokeVoid(
		s,
		"resetQualifier",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference interface {
	cdktf.ComplexObject
	AutomationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference
	AutomationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParameters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaParameters() SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference
	LambdaParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	RunCommandParameters() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
	RunCommandParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	StepFunctionsParameters() SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference
	StepFunctionsParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
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
	PutAutomationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters)
	PutLambdaParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters)
	PutRunCommandParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters)
	PutStepFunctionsParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters)
	ResetAutomationParameters()
	ResetLambdaParameters()
	ResetRunCommandParameters()
	ResetStepFunctionsParameters()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) AutomationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference
	_jsii_.Get(
		j,
		"automationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) AutomationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
	_jsii_.Get(
		j,
		"automationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) LambdaParameters() SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) LambdaParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	_jsii_.Get(
		j,
		"lambdaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) RunCommandParameters() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
	_jsii_.Get(
		j,
		"runCommandParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) RunCommandParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	_jsii_.Get(
		j,
		"runCommandParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) StepFunctionsParameters() SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) StepFunctionsParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
	_jsii_.Get(
		j,
		"stepFunctionsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutAutomationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters) {
	_jsii_.InvokeVoid(
		s,
		"putAutomationParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutLambdaParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters) {
	_jsii_.InvokeVoid(
		s,
		"putLambdaParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutRunCommandParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters) {
	_jsii_.InvokeVoid(
		s,
		"putRunCommandParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutStepFunctionsParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters) {
	_jsii_.InvokeVoid(
		s,
		"putStepFunctionsParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetAutomationParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetAutomationParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetLambdaParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetRunCommandParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetRunCommandParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetStepFunctionsParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetStepFunctionsParameters",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#cloudwatch_config SsmMaintenanceWindowTask#cloudwatch_config}
	CloudwatchConfig *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig `json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#comment SsmMaintenanceWindowTask#comment}.
	Comment *string `json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#document_hash SsmMaintenanceWindowTask#document_hash}.
	DocumentHash *string `json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#document_hash_type SsmMaintenanceWindowTask#document_hash_type}.
	DocumentHashType *string `json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#document_version SsmMaintenanceWindowTask#document_version}.
	DocumentVersion *string `json:"documentVersion" yaml:"documentVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_config SsmMaintenanceWindowTask#notification_config}
	NotificationConfig *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig `json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#output_s3_bucket SsmMaintenanceWindowTask#output_s3_bucket}.
	OutputS3Bucket *string `json:"outputS3Bucket" yaml:"outputS3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#output_s3_key_prefix SsmMaintenanceWindowTask#output_s3_key_prefix}.
	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#parameter SsmMaintenanceWindowTask#parameter}
	Parameter interface{} `json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#service_role_arn SsmMaintenanceWindowTask#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#timeout_seconds SsmMaintenanceWindowTask#timeout_seconds}.
	TimeoutSeconds *float64 `json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#cloudwatch_log_group_name SsmMaintenanceWindowTask#cloudwatch_log_group_name}.
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#cloudwatch_output_enabled SsmMaintenanceWindowTask#cloudwatch_output_enabled}.
	CloudwatchOutputEnabled interface{} `json:"cloudwatchOutputEnabled" yaml:"cloudwatchOutputEnabled"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogGroupName() *string
	SetCloudwatchLogGroupName(val *string)
	CloudwatchLogGroupNameInput() *string
	CloudwatchOutputEnabled() interface{}
	SetCloudwatchOutputEnabled(val interface{})
	CloudwatchOutputEnabledInput() interface{}
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig)
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
	ResetCloudwatchLogGroupName()
	ResetCloudwatchOutputEnabled()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) CloudwatchLogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) CloudwatchLogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) CloudwatchOutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchOutputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) CloudwatchOutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchOutputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetCloudwatchLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"cloudwatchLogGroupName",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetCloudwatchOutputEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cloudwatchOutputEnabled",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) ResetCloudwatchLogGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchLogGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference) ResetCloudwatchOutputEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchOutputEnabled",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_arn SsmMaintenanceWindowTask#notification_arn}.
	NotificationArn *string `json:"notificationArn" yaml:"notificationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_events SsmMaintenanceWindowTask#notification_events}.
	NotificationEvents *[]*string `json:"notificationEvents" yaml:"notificationEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_type SsmMaintenanceWindowTask#notification_type}.
	NotificationType *string `json:"notificationType" yaml:"notificationType"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NotificationArn() *string
	SetNotificationArn(val *string)
	NotificationArnInput() *string
	NotificationEvents() *[]*string
	SetNotificationEvents(val *[]*string)
	NotificationEventsInput() *[]*string
	NotificationType() *string
	SetNotificationType(val *string)
	NotificationTypeInput() *string
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
	ResetNotificationArn()
	ResetNotificationEvents()
	ResetNotificationType()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) NotificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetNotificationArn(val *string) {
	_jsii_.Set(
		j,
		"notificationArn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetNotificationEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationEvents",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetNotificationType(val *string) {
	_jsii_.Set(
		j,
		"notificationType",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) ResetNotificationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) ResetNotificationEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference) ResetNotificationType() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationType",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference interface {
	cdktf.ComplexObject
	CloudwatchConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference
	CloudwatchConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	DocumentHash() *string
	SetDocumentHash(val *string)
	DocumentHashInput() *string
	DocumentHashType() *string
	SetDocumentHashType(val *string)
	DocumentHashTypeInput() *string
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NotificationConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference
	NotificationConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	OutputS3Bucket() *string
	SetOutputS3Bucket(val *string)
	OutputS3BucketInput() *string
	OutputS3KeyPrefix() *string
	SetOutputS3KeyPrefix(val *string)
	OutputS3KeyPrefixInput() *string
	Parameter() interface{}
	SetParameter(val interface{})
	ParameterInput() interface{}
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutCloudwatchConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig)
	PutNotificationConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig)
	ResetCloudwatchConfig()
	ResetComment()
	ResetDocumentHash()
	ResetDocumentHashType()
	ResetDocumentVersion()
	ResetNotificationConfig()
	ResetOutputS3Bucket()
	ResetOutputS3KeyPrefix()
	ResetParameter()
	ResetServiceRoleArn()
	ResetTimeoutSeconds()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CloudwatchConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CloudwatchConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) NotificationConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) NotificationConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Parameter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetDocumentHash(val *string) {
	_jsii_.Set(
		j,
		"documentHash",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetDocumentHashType(val *string) {
	_jsii_.Set(
		j,
		"documentHashType",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetDocumentVersion(val *string) {
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetOutputS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"outputS3Bucket",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetOutputS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"outputS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetParameter(val interface{}) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) SetTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) PutCloudwatchConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig) {
	_jsii_.InvokeVoid(
		s,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) PutNotificationConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig) {
	_jsii_.InvokeVoid(
		s,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentHash() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentHash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentHashType() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentHashType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetOutputS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetOutputS3KeyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputS3KeyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		s,
		"resetParameter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#values SsmMaintenanceWindowTask#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#input SsmMaintenanceWindowTask#input}.
	Input *string `json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}.
	Name *string `json:"name" yaml:"name"`
}

type SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference interface {
	cdktf.ComplexObject
	Input() *string
	SetInput(val *string)
	InputInput() *string
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters)
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
	ResetInput()
	ResetName()
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetInput(val *string) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		s,
		"resetInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter aws_ssm_parameter}.
type SsmParameter interface {
	cdktf.TerraformResource
	AllowedPattern() *string
	SetAllowedPattern(val *string)
	AllowedPatternInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Overwrite() interface{}
	SetOverwrite(val interface{})
	OverwriteInput() interface{}
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
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
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	Version() *float64
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
	ResetAllowedPattern()
	ResetDataType()
	ResetDescription()
	ResetKeyId()
	ResetOverrideLogicalId()
	ResetOverwrite()
	ResetTags()
	ResetTagsAll()
	ResetTier()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmParameter
type jsiiProxy_SsmParameter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmParameter) AllowedPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) AllowedPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Overwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) OverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameter) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter aws_ssm_parameter} Resource.
func NewSsmParameter(scope constructs.Construct, id *string, config *SsmParameterConfig) SsmParameter {
	_init_.Initialize()

	j := jsiiProxy_SsmParameter{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmParameter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter aws_ssm_parameter} Resource.
func NewSsmParameter_Override(s SsmParameter, scope constructs.Construct, id *string, config *SsmParameterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmParameter",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmParameter) SetAllowedPattern(val *string) {
	_jsii_.Set(
		j,
		"allowedPattern",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetDataType(val *string) {
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetOverwrite(val interface{}) {
	_jsii_.Set(
		j,
		"overwrite",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SsmParameter) SetValue(val *string) {
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
func SsmParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmParameter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmParameter",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmParameter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmParameter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmParameter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmParameter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmParameter) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmParameter) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmParameter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmParameter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmParameter) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmParameter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmParameter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmParameter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmParameter) ResetAllowedPattern() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedPattern",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetDataType() {
	_jsii_.InvokeVoid(
		s,
		"resetDataType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmParameter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetOverwrite() {
	_jsii_.InvokeVoid(
		s,
		"resetOverwrite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) ResetTier() {
	_jsii_.InvokeVoid(
		s,
		"resetTier",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmParameter) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmParameter) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmParameter) ToString() *string {
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
func (s *jsiiProxy_SsmParameter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmParameterConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#name SsmParameter#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#type SsmParameter#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#value SsmParameter#value}.
	Value *string `json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#allowed_pattern SsmParameter#allowed_pattern}.
	AllowedPattern *string `json:"allowedPattern" yaml:"allowedPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#data_type SsmParameter#data_type}.
	DataType *string `json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#description SsmParameter#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#key_id SsmParameter#key_id}.
	KeyId *string `json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#overwrite SsmParameter#overwrite}.
	Overwrite interface{} `json:"overwrite" yaml:"overwrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#tags SsmParameter#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#tags_all SsmParameter#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_parameter#tier SsmParameter#tier}.
	Tier *string `json:"tier" yaml:"tier"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline aws_ssm_patch_baseline}.
type SsmPatchBaseline interface {
	cdktf.TerraformResource
	ApprovalRule() interface{}
	SetApprovalRule(val interface{})
	ApprovalRuleInput() interface{}
	ApprovedPatches() *[]*string
	SetApprovedPatches(val *[]*string)
	ApprovedPatchesComplianceLevel() *string
	SetApprovedPatchesComplianceLevel(val *string)
	ApprovedPatchesComplianceLevelInput() *string
	ApprovedPatchesEnableNonSecurity() interface{}
	SetApprovedPatchesEnableNonSecurity(val interface{})
	ApprovedPatchesEnableNonSecurityInput() interface{}
	ApprovedPatchesInput() *[]*string
	Arn() *string
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
	GlobalFilter() interface{}
	SetGlobalFilter(val interface{})
	GlobalFilterInput() interface{}
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RejectedPatches() *[]*string
	SetRejectedPatches(val *[]*string)
	RejectedPatchesAction() *string
	SetRejectedPatchesAction(val *string)
	RejectedPatchesActionInput() *string
	RejectedPatchesInput() *[]*string
	Source() interface{}
	SetSource(val interface{})
	SourceInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
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
	ResetApprovalRule()
	ResetApprovedPatches()
	ResetApprovedPatchesComplianceLevel()
	ResetApprovedPatchesEnableNonSecurity()
	ResetDescription()
	ResetGlobalFilter()
	ResetOperatingSystem()
	ResetOverrideLogicalId()
	ResetRejectedPatches()
	ResetRejectedPatchesAction()
	ResetSource()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmPatchBaseline
type jsiiProxy_SsmPatchBaseline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovalRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovalRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approvedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvedPatchesComplianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesComplianceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvedPatchesComplianceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesEnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvedPatchesEnableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesEnableNonSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvedPatchesEnableNonSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approvedPatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) GlobalFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) GlobalFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rejectedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedPatchesAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedPatchesActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rejectedPatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline aws_ssm_patch_baseline} Resource.
func NewSsmPatchBaseline(scope constructs.Construct, id *string, config *SsmPatchBaselineConfig) SsmPatchBaseline {
	_init_.Initialize()

	j := jsiiProxy_SsmPatchBaseline{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmPatchBaseline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline aws_ssm_patch_baseline} Resource.
func NewSsmPatchBaseline_Override(s SsmPatchBaseline, scope constructs.Construct, id *string, config *SsmPatchBaselineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmPatchBaseline",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetApprovalRule(val interface{}) {
	_jsii_.Set(
		j,
		"approvalRule",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetApprovedPatches(val *[]*string) {
	_jsii_.Set(
		j,
		"approvedPatches",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetApprovedPatchesComplianceLevel(val *string) {
	_jsii_.Set(
		j,
		"approvedPatchesComplianceLevel",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetApprovedPatchesEnableNonSecurity(val interface{}) {
	_jsii_.Set(
		j,
		"approvedPatchesEnableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetGlobalFilter(val interface{}) {
	_jsii_.Set(
		j,
		"globalFilter",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetOperatingSystem(val *string) {
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetRejectedPatches(val *[]*string) {
	_jsii_.Set(
		j,
		"rejectedPatches",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetRejectedPatchesAction(val *string) {
	_jsii_.Set(
		j,
		"rejectedPatchesAction",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetSource(val interface{}) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline) SetTagsAll(val *map[string]*string) {
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
func SsmPatchBaseline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmPatchBaseline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmPatchBaseline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmPatchBaseline",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmPatchBaseline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmPatchBaseline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmPatchBaseline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmPatchBaseline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmPatchBaseline) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmPatchBaseline) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmPatchBaseline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmPatchBaseline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmPatchBaseline) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmPatchBaseline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmPatchBaseline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmPatchBaseline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovalRule() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovalRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatches() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatches",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatchesComplianceLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatchesComplianceLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatchesEnableNonSecurity() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatchesEnableNonSecurity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetGlobalFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetGlobalFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmPatchBaseline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetRejectedPatches() {
	_jsii_.InvokeVoid(
		s,
		"resetRejectedPatches",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetRejectedPatchesAction() {
	_jsii_.InvokeVoid(
		s,
		"resetRejectedPatchesAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetSource() {
	_jsii_.InvokeVoid(
		s,
		"resetSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmPatchBaseline) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmPatchBaseline) ToString() *string {
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
func (s *jsiiProxy_SsmPatchBaseline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SsmPatchBaselineApprovalRule struct {
	// patch_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#patch_filter SsmPatchBaseline#patch_filter}
	PatchFilter interface{} `json:"patchFilter" yaml:"patchFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approve_after_days SsmPatchBaseline#approve_after_days}.
	ApproveAfterDays *float64 `json:"approveAfterDays" yaml:"approveAfterDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approve_until_date SsmPatchBaseline#approve_until_date}.
	ApproveUntilDate *string `json:"approveUntilDate" yaml:"approveUntilDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#compliance_level SsmPatchBaseline#compliance_level}.
	ComplianceLevel *string `json:"complianceLevel" yaml:"complianceLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#enable_non_security SsmPatchBaseline#enable_non_security}.
	EnableNonSecurity interface{} `json:"enableNonSecurity" yaml:"enableNonSecurity"`
}

type SsmPatchBaselineApprovalRulePatchFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

// AWS Systems Manager.
type SsmPatchBaselineConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#name SsmPatchBaseline#name}.
	Name *string `json:"name" yaml:"name"`
	// approval_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approval_rule SsmPatchBaseline#approval_rule}
	ApprovalRule interface{} `json:"approvalRule" yaml:"approvalRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approved_patches SsmPatchBaseline#approved_patches}.
	ApprovedPatches *[]*string `json:"approvedPatches" yaml:"approvedPatches"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approved_patches_compliance_level SsmPatchBaseline#approved_patches_compliance_level}.
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#approved_patches_enable_non_security SsmPatchBaseline#approved_patches_enable_non_security}.
	ApprovedPatchesEnableNonSecurity interface{} `json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#description SsmPatchBaseline#description}.
	Description *string `json:"description" yaml:"description"`
	// global_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#global_filter SsmPatchBaseline#global_filter}
	GlobalFilter interface{} `json:"globalFilter" yaml:"globalFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#operating_system SsmPatchBaseline#operating_system}.
	OperatingSystem *string `json:"operatingSystem" yaml:"operatingSystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#rejected_patches SsmPatchBaseline#rejected_patches}.
	RejectedPatches *[]*string `json:"rejectedPatches" yaml:"rejectedPatches"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#rejected_patches_action SsmPatchBaseline#rejected_patches_action}.
	RejectedPatchesAction *string `json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#source SsmPatchBaseline#source}
	Source interface{} `json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#tags SsmPatchBaseline#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#tags_all SsmPatchBaseline#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type SsmPatchBaselineGlobalFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `json:"values" yaml:"values"`
}

type SsmPatchBaselineSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#configuration SsmPatchBaseline#configuration}.
	Configuration *string `json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#name SsmPatchBaseline#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_baseline#products SsmPatchBaseline#products}.
	Products *[]*string `json:"products" yaml:"products"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_group aws_ssm_patch_group}.
type SsmPatchGroup interface {
	cdktf.TerraformResource
	BaselineId() *string
	SetBaselineId(val *string)
	BaselineIdInput() *string
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
	PatchGroup() *string
	SetPatchGroup(val *string)
	PatchGroupInput() *string
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

// The jsii proxy struct for SsmPatchGroup
type jsiiProxy_SsmPatchGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmPatchGroup) BaselineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) BaselineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) PatchGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) PatchGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_group aws_ssm_patch_group} Resource.
func NewSsmPatchGroup(scope constructs.Construct, id *string, config *SsmPatchGroupConfig) SsmPatchGroup {
	_init_.Initialize()

	j := jsiiProxy_SsmPatchGroup{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmPatchGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_group aws_ssm_patch_group} Resource.
func NewSsmPatchGroup_Override(s SsmPatchGroup, scope constructs.Construct, id *string, config *SsmPatchGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmPatchGroup",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetBaselineId(val *string) {
	_jsii_.Set(
		j,
		"baselineId",
		val,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetPatchGroup(val *string) {
	_jsii_.Set(
		j,
		"patchGroup",
		val,
	)
}

func (j *jsiiProxy_SsmPatchGroup) SetProvider(val cdktf.TerraformProvider) {
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
func SsmPatchGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmPatchGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmPatchGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmPatchGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmPatchGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmPatchGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmPatchGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmPatchGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmPatchGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmPatchGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmPatchGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmPatchGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmPatchGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmPatchGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmPatchGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmPatchGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmPatchGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmPatchGroup) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmPatchGroup) ToString() *string {
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
func (s *jsiiProxy_SsmPatchGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmPatchGroupConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_group#baseline_id SsmPatchGroup#baseline_id}.
	BaselineId *string `json:"baselineId" yaml:"baselineId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_patch_group#patch_group SsmPatchGroup#patch_group}.
	PatchGroup *string `json:"patchGroup" yaml:"patchGroup"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync aws_ssm_resource_data_sync}.
type SsmResourceDataSync interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3Destination() SsmResourceDataSyncS3DestinationOutputReference
	S3DestinationInput() *SsmResourceDataSyncS3Destination
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
	PutS3Destination(value *SsmResourceDataSyncS3Destination)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmResourceDataSync
type jsiiProxy_SsmResourceDataSync struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmResourceDataSync) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) S3Destination() SsmResourceDataSyncS3DestinationOutputReference {
	var returns SsmResourceDataSyncS3DestinationOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) S3DestinationInput() *SsmResourceDataSyncS3Destination {
	var returns *SsmResourceDataSyncS3Destination
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSync) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync aws_ssm_resource_data_sync} Resource.
func NewSsmResourceDataSync(scope constructs.Construct, id *string, config *SsmResourceDataSyncConfig) SsmResourceDataSync {
	_init_.Initialize()

	j := jsiiProxy_SsmResourceDataSync{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmResourceDataSync",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync aws_ssm_resource_data_sync} Resource.
func NewSsmResourceDataSync_Override(s SsmResourceDataSync, scope constructs.Construct, id *string, config *SsmResourceDataSyncConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmResourceDataSync",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmResourceDataSync) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSync) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSync) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSync) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSync) SetProvider(val cdktf.TerraformProvider) {
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
func SsmResourceDataSync_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ssm.SsmResourceDataSync",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmResourceDataSync_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ssm.SsmResourceDataSync",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SsmResourceDataSync) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SsmResourceDataSync) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmResourceDataSync) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmResourceDataSync) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmResourceDataSync) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmResourceDataSync) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmResourceDataSync) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmResourceDataSync) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmResourceDataSync) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmResourceDataSync) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmResourceDataSync) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmResourceDataSync) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmResourceDataSync) PutS3Destination(value *SsmResourceDataSyncS3Destination) {
	_jsii_.InvokeVoid(
		s,
		"putS3Destination",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SsmResourceDataSync) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmResourceDataSync) SynthesizeAttributes() *map[string]interface{} {
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
func (s *jsiiProxy_SsmResourceDataSync) ToMetadata() interface{} {
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
func (s *jsiiProxy_SsmResourceDataSync) ToString() *string {
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
func (s *jsiiProxy_SsmResourceDataSync) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Systems Manager.
type SsmResourceDataSyncConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#name SsmResourceDataSync#name}.
	Name *string `json:"name" yaml:"name"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#s3_destination SsmResourceDataSync#s3_destination}
	S3Destination *SsmResourceDataSyncS3Destination `json:"s3Destination" yaml:"s3Destination"`
}

type SsmResourceDataSyncS3Destination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#bucket_name SsmResourceDataSync#bucket_name}.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#region SsmResourceDataSync#region}.
	Region *string `json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#kms_key_arn SsmResourceDataSync#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#prefix SsmResourceDataSync#prefix}.
	Prefix *string `json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_resource_data_sync#sync_format SsmResourceDataSync#sync_format}.
	SyncFormat *string `json:"syncFormat" yaml:"syncFormat"`
}

type SsmResourceDataSyncS3DestinationOutputReference interface {
	cdktf.ComplexObject
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	InternalValue() *SsmResourceDataSyncS3Destination
	SetInternalValue(val *SsmResourceDataSyncS3Destination)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SyncFormat() *string
	SetSyncFormat(val *string)
	SyncFormatInput() *string
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
	ResetKmsKeyArn()
	ResetPrefix()
	ResetSyncFormat()
}

// The jsii proxy struct for SsmResourceDataSyncS3DestinationOutputReference
type jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) InternalValue() *SsmResourceDataSyncS3Destination {
	var returns *SsmResourceDataSyncS3Destination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SyncFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SyncFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmResourceDataSyncS3DestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SsmResourceDataSyncS3DestinationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmResourceDataSyncS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSsmResourceDataSyncS3DestinationOutputReference_Override(s SsmResourceDataSyncS3DestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ssm.SsmResourceDataSyncS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetInternalValue(val *SsmResourceDataSyncS3Destination) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetSyncFormat(val *string) {
	_jsii_.Set(
		j,
		"syncFormat",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmResourceDataSyncS3DestinationOutputReference) ResetSyncFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncFormat",
		nil, // no parameters
	)
}
