package cur

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/cur/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html aws_cur_report_definition}.
type CurReportDefinition interface {
	cdktf.TerraformResource
	AdditionalArtifacts() *[]*string
	SetAdditionalArtifacts(val *[]*string)
	AdditionalArtifactsInput() *[]*string
	AdditionalSchemaElements() *[]*string
	SetAdditionalSchemaElements(val *[]*string)
	AdditionalSchemaElementsInput() *[]*string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RefreshClosedReports() interface{}
	SetRefreshClosedReports(val interface{})
	RefreshClosedReportsInput() interface{}
	ReportName() *string
	SetReportName(val *string)
	ReportNameInput() *string
	ReportVersioning() *string
	SetReportVersioning(val *string)
	ReportVersioningInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Prefix() *string
	SetS3Prefix(val *string)
	S3PrefixInput() *string
	S3Region() *string
	SetS3Region(val *string)
	S3RegionInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TimeUnit() *string
	SetTimeUnit(val *string)
	TimeUnitInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAdditionalArtifacts()
	ResetOverrideLogicalId()
	ResetRefreshClosedReports()
	ResetReportVersioning()
	ResetS3Prefix()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CurReportDefinition
type jsiiProxy_CurReportDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CurReportDefinition) AdditionalArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalSchemaElements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSchemaElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalSchemaElementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSchemaElementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RefreshClosedReports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshClosedReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RefreshClosedReportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshClosedReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnitInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html aws_cur_report_definition} Resource.
func NewCurReportDefinition(scope constructs.Construct, id *string, config *CurReportDefinitionConfig) CurReportDefinition {
	_init_.Initialize()

	j := jsiiProxy_CurReportDefinition{}

	_jsii_.Create(
		"hashicorp_aws.Cur.CurReportDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html aws_cur_report_definition} Resource.
func NewCurReportDefinition_Override(c CurReportDefinition, scope constructs.Construct, id *string, config *CurReportDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cur.CurReportDefinition",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetAdditionalArtifacts(val *[]*string) {
	_jsii_.Set(
		j,
		"additionalArtifacts",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetAdditionalSchemaElements(val *[]*string) {
	_jsii_.Set(
		j,
		"additionalSchemaElements",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetCompression(val *string) {
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetRefreshClosedReports(val interface{}) {
	_jsii_.Set(
		j,
		"refreshClosedReports",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetReportName(val *string) {
	_jsii_.Set(
		j,
		"reportName",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetReportVersioning(val *string) {
	_jsii_.Set(
		j,
		"reportVersioning",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetS3Prefix(val *string) {
	_jsii_.Set(
		j,
		"s3Prefix",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetS3Region(val *string) {
	_jsii_.Set(
		j,
		"s3Region",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition) SetTimeUnit(val *string) {
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CurReportDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cur.CurReportDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CurReportDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cur.CurReportDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CurReportDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetAdditionalArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalArtifacts",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CurReportDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetRefreshClosedReports() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshClosedReports",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetReportVersioning() {
	_jsii_.InvokeVoid(
		c,
		"resetReportVersioning",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetS3Prefix() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Prefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CurReportDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CurReportDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CurReportDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CurReportDefinitionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#additional_schema_elements CurReportDefinition#additional_schema_elements}.
	AdditionalSchemaElements *[]*string `json:"additionalSchemaElements"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#compression CurReportDefinition#compression}.
	Compression *string `json:"compression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#format CurReportDefinition#format}.
	Format *string `json:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#report_name CurReportDefinition#report_name}.
	ReportName *string `json:"reportName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#s3_bucket CurReportDefinition#s3_bucket}.
	S3Bucket *string `json:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#s3_region CurReportDefinition#s3_region}.
	S3Region *string `json:"s3Region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#time_unit CurReportDefinition#time_unit}.
	TimeUnit *string `json:"timeUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#additional_artifacts CurReportDefinition#additional_artifacts}.
	AdditionalArtifacts *[]*string `json:"additionalArtifacts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#refresh_closed_reports CurReportDefinition#refresh_closed_reports}.
	RefreshClosedReports interface{} `json:"refreshClosedReports"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#report_versioning CurReportDefinition#report_versioning}.
	ReportVersioning *string `json:"reportVersioning"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cur_report_definition.html#s3_prefix CurReportDefinition#s3_prefix}.
	S3Prefix *string `json:"s3Prefix"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cur_report_definition.html aws_cur_report_definition}.
type DataAwsCurReportDefinition interface {
	cdktf.TerraformDataSource
	AdditionalArtifacts() *[]*string
	AdditionalSchemaElements() *[]*string
	CdktfStack() cdktf.TerraformStack
	Compression() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Format() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RefreshClosedReports() interface{}
	ReportName() *string
	SetReportName(val *string)
	ReportNameInput() *string
	ReportVersioning() *string
	S3Bucket() *string
	S3Prefix() *string
	S3Region() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TimeUnit() *string
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

// The jsii proxy struct for DataAwsCurReportDefinition
type jsiiProxy_DataAwsCurReportDefinition struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCurReportDefinition) AdditionalArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) AdditionalSchemaElements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSchemaElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) RefreshClosedReports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshClosedReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) ReportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) ReportNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) ReportVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCurReportDefinition) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cur_report_definition.html aws_cur_report_definition} Data Source.
func NewDataAwsCurReportDefinition(scope constructs.Construct, id *string, config *DataAwsCurReportDefinitionConfig) DataAwsCurReportDefinition {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCurReportDefinition{}

	_jsii_.Create(
		"hashicorp_aws.Cur.DataAwsCurReportDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cur_report_definition.html aws_cur_report_definition} Data Source.
func NewDataAwsCurReportDefinition_Override(d DataAwsCurReportDefinition, scope constructs.Construct, id *string, config *DataAwsCurReportDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Cur.DataAwsCurReportDefinition",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCurReportDefinition) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCurReportDefinition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCurReportDefinition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCurReportDefinition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCurReportDefinition) SetReportName(val *string) {
	_jsii_.Set(
		j,
		"reportName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsCurReportDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Cur.DataAwsCurReportDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCurReportDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Cur.DataAwsCurReportDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCurReportDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCurReportDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCurReportDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCurReportDefinition) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) ToString() *string {
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
func (d *jsiiProxy_DataAwsCurReportDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCurReportDefinitionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cur_report_definition.html#report_name DataAwsCurReportDefinition#report_name}.
	ReportName *string `json:"reportName"`
}
